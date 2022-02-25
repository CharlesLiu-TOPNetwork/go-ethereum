package executor

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	DefaultGasPrice = 20000000000 // 20 GWei
)

var (
	crossChainContractAddr = common.HexToAddress("0x0000000000000000000000000000000000000100")
)
