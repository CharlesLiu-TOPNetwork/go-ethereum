package executor

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"

	relayercommon "github.com/telosprotocol/teth-relayer/common"
	config "github.com/telosprotocol/teth-relayer/config"
	"github.com/telosprotocol/teth-relayer/executor/crosschain"
)

type TethClient struct {
	TethClient    *ethclient.Client
	Provider      string
	CurrentHeight int64
	UpdatedAt     time.Time
}

type TethExecutor struct {
	mutex         sync.RWMutex
	db            *gorm.DB
	clientIdx     int
	tethClients   []*TethClient
	sourceChainID relayercommon.CrossChainID
	destChainID   relayercommon.CrossChainID
	privateKey    *ecdsa.PrivateKey
	txSender      common.Address
	cfg           *config.Config
}

func getPrivateKey(cfg *config.TethConfig) (*ecdsa.PrivateKey, error) {
	privKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func initClients(providers []string) []*TethClient {
	clients := make([]*TethClient, 0)

	for _, provider := range providers {
		client, err := ethclient.Dial(provider)
		if err != nil {
			panic("new eth client error")
		}
		clients = append(clients, &TethClient{
			TethClient: client,
			Provider:   provider,
			UpdatedAt:  time.Now(),
		})
	}

	return clients
}

func NewTethExecutor(db *gorm.DB, cfg *config.Config) (*TethExecutor, error) {
	privKey, err := getPrivateKey(&cfg.TethConfig)
	if err != nil {
		return nil, err
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("get public key error")
	}
	txSender := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &TethExecutor{
		db:            db,
		clientIdx:     0,
		tethClients:   initClients(cfg.TethConfig.Providers),
		privateKey:    privKey,
		txSender:      txSender,
		sourceChainID: relayercommon.CrossChainID(cfg.CrossChainConfig.SourceChainID),
		destChainID:   relayercommon.CrossChainID(cfg.CrossChainConfig.DestChainID),
		cfg:           cfg,
	}, nil
}

func (executor *TethExecutor) GetClient() *ethclient.Client {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.tethClients[executor.clientIdx].TethClient
}

func (executor *TethExecutor) SwitchTethClient() {
	executor.mutex.Lock()
	defer executor.mutex.Unlock()
	executor.clientIdx++
	if executor.clientIdx >= len(executor.tethClients) {
		executor.clientIdx = 0
	}
	relayercommon.Logger.Infof("Switch to provider: %s", executor.cfg.TethConfig.Providers[executor.clientIdx])
}

func (executor *TethExecutor) GetLatestBlockHeight(client *ethclient.Client) (int64, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	block, err := client.BlockByNumber(ctxWithTimeout, nil)
	if err != nil {
		return 0, err
	}
	return block.Number().Int64(), nil
}

func (executor *TethExecutor) getTransactor(nonce uint64) (*bind.TransactOpts, error) {
	txOpts := bind.NewKeyedTransactor(executor.privateKey)
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	txOpts.GasLimit = executor.cfg.TethConfig.GasLimit
	if executor.cfg.TethConfig.GasPrice == 0 {
		txOpts.GasPrice = big.NewInt(DefaultGasPrice)
	} else {
		txOpts.GasPrice = big.NewInt(int64(executor.cfg.TethConfig.GasPrice))
	}
	return txOpts, nil
}

func (executor *TethExecutor) getCallOpts() (*bind.CallOpts, error) {
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}
	return callOpts, nil
}

func (executor *TethExecutor) CallBuildInSystemContract(channelID relayercommon.CrossChainChannelID, height, sequence uint64, msgBytes, proofBytes []byte, nonce uint64) (common.Hash, error) {
	txOpts, err := executor.getTransactor(nonce)
	if err != nil {
		return common.Hash{}, err
	}

	crossChainInstance, err := crosschain.NewCrossChain(crossChainContractAddr, executor.GetClient())
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := crossChainInstance.HandleCCPackage(txOpts, msgBytes, sequence, uint8(channelID))
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (executor *TethExecutor) RelayCrossChainPackage(channelID relayercommon.CrossChainChannelID, sequence, height uint64) (common.Hash, error) {
	nonce, err := executor.GetClient().PendingNonceAt(context.Background(), executor.txSender)
	if err != nil {
		return common.Hash{}, err
	}
	var empty []byte
	tx, err := executor.CallBuildInSystemContract(channelID, height, sequence, empty, empty, nonce)
	if err != nil {
		return common.Hash{}, err
	}

	relayercommon.Logger.Infof("channelID: %d, sequence: %d, txHash: %s", channelID, sequence, tx.String())
	return tx, nil
}
