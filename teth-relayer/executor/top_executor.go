package executor

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/teth-relayer/common"
	config "github.com/ethereum/go-ethereum/teth-relayer/config"
)

type TopClient struct {
	TopClient     *http.Client
	Provider      string
	CurrentHeight int64
	UpdatedAt     time.Time
}

type TopExecutor struct {
	mutex         sync.RWMutex
	clientIdx     int
	TopClients    []*TopClient
	Config        *config.Config
	sourceChainID common.CrossChainID
	destChainID   common.CrossChainID
}

func initTopClients(providers []string, network config.ChainNetwork) []*TopClient {
	topClients := make([]*TopClient, 0)
	if len(providers[0]) == 0 {
		return topClients
	}
	for _, provider := range providers {
		rpcClient := &http.Client{
			// Transport: &http.Transport{
			// 	Proxy: http.ProxyFromEnvironment,
			// 	DialContext: (&net.Dialer{
			// 		Timeout:   30 * time.Second,
			// 		KeepAlive: 30 * time.Second,
			// 	}).DialContext,
			// 	MaxIdleConns:        100,
			// 	MaxIdleConnsPerHost: 100,
			// 	IdleConnTimeout:     time.Duration(90) * time.Second,
			// },
			// Timeout: 20 * time.Second,
		}
		topClients = append(topClients, &TopClient{
			TopClient: rpcClient,
			Provider:  provider,
			UpdatedAt: time.Now(),
		})
	}
	return topClients
}

func NewTopExecutor(cfg *config.Config, networkType config.ChainNetwork) (*TopExecutor, error) {
	return &TopExecutor{
		clientIdx:     0,
		TopClients:    initTopClients(cfg.TopConfig.RpcAddrs, networkType),
		Config:        cfg,
		sourceChainID: common.CrossChainID(cfg.CrossChainConfig.SourceChainID),
		destChainID:   common.CrossChainID(cfg.CrossChainConfig.DestChainID),
	}, nil
}

func (executor *TopExecutor) GetClient() *http.Client {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.TopClients[executor.clientIdx].TopClient
}

func (executor *TopExecutor) GetProvider() string {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.TopClients[executor.clientIdx].Provider
}

func (executor *TopExecutor) SwitchTopClient() {
	executor.mutex.Lock()
	defer executor.mutex.Unlock()
	executor.clientIdx++
	if executor.clientIdx >= len(executor.TopClients) {
		executor.clientIdx = 0
	}
	common.Logger.Infof("Switch to RPC endpoint: %s", executor.Config.TopConfig.RpcAddrs[executor.clientIdx])
}

func (executor *TopExecutor) GetLatestBlockHeight(address string) (uint64, error) {
	data := "target_account_addr=" + address + "&body={\"params\":{\"account_addr\":\"" + address + "\",\"height\":\"latest\"}}&method=getBlock&sequence_id=1&identity_token=&version=1.0"
	request, err := http.NewRequest("POST", executor.GetProvider(), strings.NewReader(data))
	if err != nil {
		common.Logger.Errorf("GetLatestBlockHeight NewRequest error, err=%s", err.Error())
		return 0, err
	}
	request.Close = true
	response, err := executor.GetClient().Do(request)
	if err != nil {
		common.Logger.Errorf("GetLatestBlockHeight Do request error, err=%s", err.Error())
		return 0, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		common.Logger.Errorf("GetLatestBlockHeight ReadAll error, err=%s", err.Error())
		return 0, err
	}

	var unknown interface{}
	err = json.Unmarshal(body, &unknown)
	if err != nil {
		common.Logger.Errorf("GetLatestBlockHeight Unmarshal error, err=%s", err.Error())
		return 0, err
	}
	m := unknown.(map[string]interface{})
	errnointerface, isErrnoPresent := m["errno"]
	errmsg, isErrmsgPresent := m["errmsg"]

	var errno int64 = 0
	if isErrnoPresent {
		errnof, _ := errnointerface.(float64)
		errno = int64(errnof)
		if errno != 0 {
			if isErrmsgPresent && errmsg != "OK" {
				common.Logger.Errorf("GetLatestBlockHeight error, err=%d, errmsg=%s", errno, errmsg)
			} else {
				common.Logger.Errorf("GetLatestBlockHeight error, err=%d", errno)
			}
			return 0, common.ErrRpcReturnError
		}
	}
	if isErrmsgPresent && errmsg != "OK" {
		common.Logger.Errorf("GetLatestBlockHeight error, err=%d, errmsg=%s", errno, errmsg)
		return 0, common.ErrRpcReturnError
	}
	if data, isPresent := m["data"]; isPresent {
		d := data.(map[string]interface{})
		if value, isPresent := d["value"]; isPresent {
			v := value.(map[string]interface{})
			if height, isPresent := v["height"]; isPresent {
				heightf := height.(float64)
				return uint64(heightf), nil
			}
		}
	}
	// common.Logger.Errorf("GetLatestBlockHeight height not found")
	return 0, common.ErrRpcValueNotfound
}

func (executor *TopExecutor) GetBlockByHeight(address string, height uint64) ([]byte, error) {
	data := "target_account_addr=" + address + "&body={\"params\":{\"account_addr\":\"" + address + "\",\"height\":" + strconv.FormatUint(height, 10) + "}}&method=getRawBlock&sequence_id=1&identity_token=&version=1.0"
	var empty []byte
	request, err := http.NewRequest("POST", executor.GetProvider(), strings.NewReader(data))
	if err != nil {
		common.Logger.Errorf("GetBlockByHeight NewRequest error, err=%s", err.Error())
		return empty, err
	}
	response, err := executor.GetClient().Do(request)
	if err != nil {
		common.Logger.Errorf("GetBlockByHeight Do request error, err=%s", err.Error())
		return empty, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		common.Logger.Errorf("GetBlockByHeight ReadAll error, err=%s", err.Error())
		return empty, err
	}

	var unknown interface{}
	err = json.Unmarshal(body, &unknown)
	if err != nil {
		common.Logger.Errorf("GetBlockByHeight Unmarshal error, err=%s", err.Error())
		return empty, err
	}
	m := unknown.(map[string]interface{})
	errnointerface, isErrnoPresent := m["errno"]
	errmsg, isErrmsgPresent := m["errmsg"]

	var errno int64 = 0
	if isErrnoPresent {
		errnof, _ := errnointerface.(float64)
		errno = int64(errnof)
		if errno != 0 {
			if isErrmsgPresent && errmsg != "OK" {
				common.Logger.Errorf("GetBlockByHeight error, err=%d, errmsg=%s", errno, errmsg)
			} else {
				common.Logger.Errorf("GetBlockByHeight error, err=%d", errno)
			}
			return empty, common.ErrRpcReturnError
		}
	}
	if isErrmsgPresent && errmsg != "OK" {
		common.Logger.Errorf("GetBlockByHeight error, err=%d, errmsg=%s", errno, errmsg)
		return empty, common.ErrRpcReturnError
	}
	if data, isPresent := m["data"]; isPresent {
		d := data.(map[string]interface{})
		if value, isPresent := d["value"]; isPresent {
			return []byte(value.(string)), nil
		}
	}
	// common.Logger.Errorf("GetBlockByHeight height not found")
	return empty, common.ErrRpcValueNotfound
}
