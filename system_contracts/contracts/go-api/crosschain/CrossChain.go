// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crosschain

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CrossChainMetaData contains all meta data concerning the CrossChain contract.
var CrossChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"packageDecodeError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"}],\"name\":\"receivePackage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAIL_ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYN_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_EXCHANGE_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"}],\"name\":\"handleCCPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"handleContractMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b0355f5b": "ACK_PACKAGE()",
		"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
		"8cc8f561": "FAIL_ACK_PACKAGE()",
		"05e68258": "SYN_PACKAGE()",
		"0e77db70": "TOKEN_EXCHANGE_CONTRACT_ADDR()",
		"4fa689d8": "TRANSFER_IN_HANDLE()",
		"b6d28d2e": "TRANSFER_OUT_HANDLE()",
		"a78abc16": "alreadyInit()",
		"82fa3ab1": "handleCCPackage(bytes,uint64,uint8)",
		"460c8eb5": "handleContractMap(uint8)",
		"e1c7392a": "init()",
	},
	Bin: "0x608060405234801561001057600080fd5b5061073e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806382fa3ab11161007157806382fa3ab1146101205780638cc8f561146101a5578063a78abc16146101ad578063b0355f5b14610110578063b6d28d2e146101a5578063e1c7392a146101c9576100a9565b806305e68258146100ae5780630e77db70146100cc578063460c8eb5146100f05780634fa689d81461011057806351e8067214610118575b600080fd5b6100b66101d1565b6040805160ff9092168252519081900360200190f35b6100d46101d6565b604080516001600160a01b039092168252519081900360200190f35b6100d46004803603602081101561010657600080fd5b503560ff166101dc565b6100b66101f7565b6100d46101fc565b6101a36004803603606081101561013657600080fd5b81019060208101813564010000000081111561015157600080fd5b82018360208201111561016357600080fd5b8035906020019184600183028401116401000000008311171561018557600080fd5b9193509150803567ffffffffffffffff16906020013560ff16610202565b005b6100b6610605565b6101b561060a565b604080519115158252519081900360200190f35b6101a3610613565b600081565b61020081565b6001602052600090815260409020546001600160a01b031681565b600181565b61010081565b60005460ff16610251576040805162461bcd60e51b815260206004820152601560248201527410dbdb9d1c9858dd081b9bdd081a5b9a5d081e595d605a1b604482015290519081900360640190fd5b6000806000606061029788888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506106af92505050565b93509350935093508361031c578460ff168667ffffffffffffffff167f997d91a02d5e696ab94a30c03b00d5ada3207f80818db98c37f3f5bfdeda00128a8a60405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a3505050506105ff565b6040805160ff858116825291519187169167ffffffffffffffff8916917f591bae32fbc801fec820cde2a1ac597eb40c47c67d087a10cc51417dad79e1a8919081900360200190a360ff83166105fa5760ff85166000818152600160209081526040808320548151631182b87560e01b815260048101958652602481019283528651604482015286516001600160a01b03909216958695631182b875958d958a9593949093606490910192918601918190849084905b838110156103ea5781810151838201526020016103d2565b50505050905090810190601f1680156104175780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561043757600080fd5b505af192505050801561051d57506040513d6000823e601f3d908101601f19168201604052602081101561046a57600080fd5b810190808051604051939291908464010000000082111561048a57600080fd5b90830190602082018581111561049f57600080fd5b82516401000000008111828201881017156104b957600080fd5b82525081516020918201929091019080838360005b838110156104e65781810151838201526020016104ce565b50505050905090810190601f1680156105135780820380516001836020036101000a031916815260200191505b5060405250505060015b6105f6576040516000815260443d1015610539575060006105d6565b60046000803e60005160e01c6308c379a0811461055a5760009150506105d6565b60043d036004833e81513d602482011167ffffffffffffffff82111715610586576000925050506105d6565b808301805167ffffffffffffffff8111156105a85760009450505050506105d6565b8060208301013d86018111156105c6576000955050505050506105d6565b601f01601f191660405250925050505b806105e157506105e7565b506105f1565b3d6000803e3d6000fd5b6105f8565b505b505b505050505b50505050565b600281565b60005460ff1681565b60005460ff1615610663576040805162461bcd60e51b815260206004820152601560248201527410dbdb9d1c9858dd081b9bdd081a5b9a5d081e595d605a1b604482015290519081900360640190fd5b6001600081815260208290527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f80546001600160a01b031916610200179055805460ff19169091179055565b5060408051600080825260208201909252600192829190565b602081106106e7578251825260209283019290910190601f19016106c8565b915181516020939093036101000a600019018019909116921691909117905256fea26469706673582212204259d42de28a8744f2bbbd4fbc4ae7c2e603b47e245f64b15b05f45cba94e4e564736f6c63430006040033",
}

// CrossChainABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossChainMetaData.ABI instead.
var CrossChainABI = CrossChainMetaData.ABI

// Deprecated: Use CrossChainMetaData.Sigs instead.
// CrossChainFuncSigs maps the 4-byte function signature to its string representation.
var CrossChainFuncSigs = CrossChainMetaData.Sigs

// CrossChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossChainMetaData.Bin instead.
var CrossChainBin = CrossChainMetaData.Bin

// DeployCrossChain deploys a new Ethereum contract, binding an instance of CrossChain to it.
func DeployCrossChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossChain, error) {
	parsed, err := CrossChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossChain{CrossChainCaller: CrossChainCaller{contract: contract}, CrossChainTransactor: CrossChainTransactor{contract: contract}, CrossChainFilterer: CrossChainFilterer{contract: contract}}, nil
}

// CrossChain is an auto generated Go binding around an Ethereum contract.
type CrossChain struct {
	CrossChainCaller     // Read-only binding to the contract
	CrossChainTransactor // Write-only binding to the contract
	CrossChainFilterer   // Log filterer for contract events
}

// CrossChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainSession struct {
	Contract     *CrossChain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainCallerSession struct {
	Contract *CrossChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrossChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainTransactorSession struct {
	Contract     *CrossChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrossChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainRaw struct {
	Contract *CrossChain // Generic contract binding to access the raw methods on
}

// CrossChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainCallerRaw struct {
	Contract *CrossChainCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainTransactorRaw struct {
	Contract *CrossChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChain creates a new instance of CrossChain, bound to a specific deployed contract.
func NewCrossChain(address common.Address, backend bind.ContractBackend) (*CrossChain, error) {
	contract, err := bindCrossChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChain{CrossChainCaller: CrossChainCaller{contract: contract}, CrossChainTransactor: CrossChainTransactor{contract: contract}, CrossChainFilterer: CrossChainFilterer{contract: contract}}, nil
}

// NewCrossChainCaller creates a new read-only instance of CrossChain, bound to a specific deployed contract.
func NewCrossChainCaller(address common.Address, caller bind.ContractCaller) (*CrossChainCaller, error) {
	contract, err := bindCrossChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainCaller{contract: contract}, nil
}

// NewCrossChainTransactor creates a new write-only instance of CrossChain, bound to a specific deployed contract.
func NewCrossChainTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainTransactor, error) {
	contract, err := bindCrossChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainTransactor{contract: contract}, nil
}

// NewCrossChainFilterer creates a new log filterer instance of CrossChain, bound to a specific deployed contract.
func NewCrossChainFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainFilterer, error) {
	contract, err := bindCrossChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainFilterer{contract: contract}, nil
}

// bindCrossChain binds a generic wrapper to an already deployed contract.
func bindCrossChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChain *CrossChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChain.Contract.CrossChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChain *CrossChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChain.Contract.CrossChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChain *CrossChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChain.Contract.CrossChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChain *CrossChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChain *CrossChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChain *CrossChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChain.Contract.contract.Transact(opts, method, params...)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCaller) ACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "ACK_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainSession) ACKPACKAGE() (uint8, error) {
	return _CrossChain.Contract.ACKPACKAGE(&_CrossChain.CallOpts)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCallerSession) ACKPACKAGE() (uint8, error) {
	return _CrossChain.Contract.ACKPACKAGE(&_CrossChain.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_CrossChain *CrossChainCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_CrossChain *CrossChainSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _CrossChain.Contract.CROSSCHAINCONTRACTADDR(&_CrossChain.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_CrossChain *CrossChainCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _CrossChain.Contract.CROSSCHAINCONTRACTADDR(&_CrossChain.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCaller) FAILACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "FAIL_ACK_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainSession) FAILACKPACKAGE() (uint8, error) {
	return _CrossChain.Contract.FAILACKPACKAGE(&_CrossChain.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCallerSession) FAILACKPACKAGE() (uint8, error) {
	return _CrossChain.Contract.FAILACKPACKAGE(&_CrossChain.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCaller) SYNPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "SYN_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainSession) SYNPACKAGE() (uint8, error) {
	return _CrossChain.Contract.SYNPACKAGE(&_CrossChain.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCallerSession) SYNPACKAGE() (uint8, error) {
	return _CrossChain.Contract.SYNPACKAGE(&_CrossChain.CallOpts)
}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_CrossChain *CrossChainCaller) TOKENEXCHANGECONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "TOKEN_EXCHANGE_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_CrossChain *CrossChainSession) TOKENEXCHANGECONTRACTADDR() (common.Address, error) {
	return _CrossChain.Contract.TOKENEXCHANGECONTRACTADDR(&_CrossChain.CallOpts)
}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_CrossChain *CrossChainCallerSession) TOKENEXCHANGECONTRACTADDR() (common.Address, error) {
	return _CrossChain.Contract.TOKENEXCHANGECONTRACTADDR(&_CrossChain.CallOpts)
}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_CrossChain *CrossChainCaller) TRANSFERINHANDLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "TRANSFER_IN_HANDLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_CrossChain *CrossChainSession) TRANSFERINHANDLE() (uint8, error) {
	return _CrossChain.Contract.TRANSFERINHANDLE(&_CrossChain.CallOpts)
}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_CrossChain *CrossChainCallerSession) TRANSFERINHANDLE() (uint8, error) {
	return _CrossChain.Contract.TRANSFERINHANDLE(&_CrossChain.CallOpts)
}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_CrossChain *CrossChainCaller) TRANSFEROUTHANDLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "TRANSFER_OUT_HANDLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_CrossChain *CrossChainSession) TRANSFEROUTHANDLE() (uint8, error) {
	return _CrossChain.Contract.TRANSFEROUTHANDLE(&_CrossChain.CallOpts)
}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_CrossChain *CrossChainCallerSession) TRANSFEROUTHANDLE() (uint8, error) {
	return _CrossChain.Contract.TRANSFEROUTHANDLE(&_CrossChain.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_CrossChain *CrossChainCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_CrossChain *CrossChainSession) AlreadyInit() (bool, error) {
	return _CrossChain.Contract.AlreadyInit(&_CrossChain.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_CrossChain *CrossChainCallerSession) AlreadyInit() (bool, error) {
	return _CrossChain.Contract.AlreadyInit(&_CrossChain.CallOpts)
}

// HandleContractMap is a free data retrieval call binding the contract method 0x460c8eb5.
//
// Solidity: function handleContractMap(uint8 ) view returns(address)
func (_CrossChain *CrossChainCaller) HandleContractMap(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "handleContractMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HandleContractMap is a free data retrieval call binding the contract method 0x460c8eb5.
//
// Solidity: function handleContractMap(uint8 ) view returns(address)
func (_CrossChain *CrossChainSession) HandleContractMap(arg0 uint8) (common.Address, error) {
	return _CrossChain.Contract.HandleContractMap(&_CrossChain.CallOpts, arg0)
}

// HandleContractMap is a free data retrieval call binding the contract method 0x460c8eb5.
//
// Solidity: function handleContractMap(uint8 ) view returns(address)
func (_CrossChain *CrossChainCallerSession) HandleContractMap(arg0 uint8) (common.Address, error) {
	return _CrossChain.Contract.HandleContractMap(&_CrossChain.CallOpts, arg0)
}

// HandleCCPackage is a paid mutator transaction binding the contract method 0x82fa3ab1.
//
// Solidity: function handleCCPackage(bytes payload, uint64 packageSequence, uint8 handleId) returns()
func (_CrossChain *CrossChainTransactor) HandleCCPackage(opts *bind.TransactOpts, payload []byte, packageSequence uint64, handleId uint8) (*types.Transaction, error) {
	return _CrossChain.contract.Transact(opts, "handleCCPackage", payload, packageSequence, handleId)
}

// HandleCCPackage is a paid mutator transaction binding the contract method 0x82fa3ab1.
//
// Solidity: function handleCCPackage(bytes payload, uint64 packageSequence, uint8 handleId) returns()
func (_CrossChain *CrossChainSession) HandleCCPackage(payload []byte, packageSequence uint64, handleId uint8) (*types.Transaction, error) {
	return _CrossChain.Contract.HandleCCPackage(&_CrossChain.TransactOpts, payload, packageSequence, handleId)
}

// HandleCCPackage is a paid mutator transaction binding the contract method 0x82fa3ab1.
//
// Solidity: function handleCCPackage(bytes payload, uint64 packageSequence, uint8 handleId) returns()
func (_CrossChain *CrossChainTransactorSession) HandleCCPackage(payload []byte, packageSequence uint64, handleId uint8) (*types.Transaction, error) {
	return _CrossChain.Contract.HandleCCPackage(&_CrossChain.TransactOpts, payload, packageSequence, handleId)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_CrossChain *CrossChainTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChain.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_CrossChain *CrossChainSession) Init() (*types.Transaction, error) {
	return _CrossChain.Contract.Init(&_CrossChain.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_CrossChain *CrossChainTransactorSession) Init() (*types.Transaction, error) {
	return _CrossChain.Contract.Init(&_CrossChain.TransactOpts)
}

// CrossChainPackageDecodeErrorIterator is returned from FilterPackageDecodeError and is used to iterate over the raw logs and unpacked data for PackageDecodeError events raised by the CrossChain contract.
type CrossChainPackageDecodeErrorIterator struct {
	Event *CrossChainPackageDecodeError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossChainPackageDecodeErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPackageDecodeError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossChainPackageDecodeError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossChainPackageDecodeErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainPackageDecodeErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainPackageDecodeError represents a PackageDecodeError event raised by the CrossChain contract.
type CrossChainPackageDecodeError struct {
	PackageSequence uint64
	HandleId        uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPackageDecodeError is a free log retrieval operation binding the contract event 0x997d91a02d5e696ab94a30c03b00d5ada3207f80818db98c37f3f5bfdeda0012.
//
// Solidity: event packageDecodeError(uint64 indexed packageSequence, uint8 indexed handleId, bytes payload)
func (_CrossChain *CrossChainFilterer) FilterPackageDecodeError(opts *bind.FilterOpts, packageSequence []uint64, handleId []uint8) (*CrossChainPackageDecodeErrorIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var handleIdRule []interface{}
	for _, handleIdItem := range handleId {
		handleIdRule = append(handleIdRule, handleIdItem)
	}

	logs, sub, err := _CrossChain.contract.FilterLogs(opts, "packageDecodeError", packageSequenceRule, handleIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPackageDecodeErrorIterator{contract: _CrossChain.contract, event: "packageDecodeError", logs: logs, sub: sub}, nil
}

// WatchPackageDecodeError is a free log subscription operation binding the contract event 0x997d91a02d5e696ab94a30c03b00d5ada3207f80818db98c37f3f5bfdeda0012.
//
// Solidity: event packageDecodeError(uint64 indexed packageSequence, uint8 indexed handleId, bytes payload)
func (_CrossChain *CrossChainFilterer) WatchPackageDecodeError(opts *bind.WatchOpts, sink chan<- *CrossChainPackageDecodeError, packageSequence []uint64, handleId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var handleIdRule []interface{}
	for _, handleIdItem := range handleId {
		handleIdRule = append(handleIdRule, handleIdItem)
	}

	logs, sub, err := _CrossChain.contract.WatchLogs(opts, "packageDecodeError", packageSequenceRule, handleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainPackageDecodeError)
				if err := _CrossChain.contract.UnpackLog(event, "packageDecodeError", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePackageDecodeError is a log parse operation binding the contract event 0x997d91a02d5e696ab94a30c03b00d5ada3207f80818db98c37f3f5bfdeda0012.
//
// Solidity: event packageDecodeError(uint64 indexed packageSequence, uint8 indexed handleId, bytes payload)
func (_CrossChain *CrossChainFilterer) ParsePackageDecodeError(log types.Log) (*CrossChainPackageDecodeError, error) {
	event := new(CrossChainPackageDecodeError)
	if err := _CrossChain.contract.UnpackLog(event, "packageDecodeError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainReceivePackageIterator is returned from FilterReceivePackage and is used to iterate over the raw logs and unpacked data for ReceivePackage events raised by the CrossChain contract.
type CrossChainReceivePackageIterator struct {
	Event *CrossChainReceivePackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossChainReceivePackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainReceivePackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossChainReceivePackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossChainReceivePackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainReceivePackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainReceivePackage represents a ReceivePackage event raised by the CrossChain contract.
type CrossChainReceivePackage struct {
	PackageType     uint8
	PackageSequence uint64
	HandleId        uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReceivePackage is a free log retrieval operation binding the contract event 0x591bae32fbc801fec820cde2a1ac597eb40c47c67d087a10cc51417dad79e1a8.
//
// Solidity: event receivePackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed handleId)
func (_CrossChain *CrossChainFilterer) FilterReceivePackage(opts *bind.FilterOpts, packageSequence []uint64, handleId []uint8) (*CrossChainReceivePackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var handleIdRule []interface{}
	for _, handleIdItem := range handleId {
		handleIdRule = append(handleIdRule, handleIdItem)
	}

	logs, sub, err := _CrossChain.contract.FilterLogs(opts, "receivePackage", packageSequenceRule, handleIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainReceivePackageIterator{contract: _CrossChain.contract, event: "receivePackage", logs: logs, sub: sub}, nil
}

// WatchReceivePackage is a free log subscription operation binding the contract event 0x591bae32fbc801fec820cde2a1ac597eb40c47c67d087a10cc51417dad79e1a8.
//
// Solidity: event receivePackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed handleId)
func (_CrossChain *CrossChainFilterer) WatchReceivePackage(opts *bind.WatchOpts, sink chan<- *CrossChainReceivePackage, packageSequence []uint64, handleId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var handleIdRule []interface{}
	for _, handleIdItem := range handleId {
		handleIdRule = append(handleIdRule, handleIdItem)
	}

	logs, sub, err := _CrossChain.contract.WatchLogs(opts, "receivePackage", packageSequenceRule, handleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainReceivePackage)
				if err := _CrossChain.contract.UnpackLog(event, "receivePackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivePackage is a log parse operation binding the contract event 0x591bae32fbc801fec820cde2a1ac597eb40c47c67d087a10cc51417dad79e1a8.
//
// Solidity: event receivePackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed handleId)
func (_CrossChain *CrossChainFilterer) ParseReceivePackage(log types.Log) (*CrossChainReceivePackage, error) {
	event := new(CrossChainReceivePackage)
	if err := _CrossChain.contract.UnpackLog(event, "receivePackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAppMetaData contains all meta data concerning the IApp contract.
var IAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1182b875": "handleSynPackage(uint8,bytes)",
	},
}

// IAppABI is the input ABI used to generate the binding from.
// Deprecated: Use IAppMetaData.ABI instead.
var IAppABI = IAppMetaData.ABI

// Deprecated: Use IAppMetaData.Sigs instead.
// IAppFuncSigs maps the 4-byte function signature to its string representation.
var IAppFuncSigs = IAppMetaData.Sigs

// IApp is an auto generated Go binding around an Ethereum contract.
type IApp struct {
	IAppCaller     // Read-only binding to the contract
	IAppTransactor // Write-only binding to the contract
	IAppFilterer   // Log filterer for contract events
}

// IAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAppSession struct {
	Contract     *IApp             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAppCallerSession struct {
	Contract *IAppCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAppTransactorSession struct {
	Contract     *IAppTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAppRaw struct {
	Contract *IApp // Generic contract binding to access the raw methods on
}

// IAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAppCallerRaw struct {
	Contract *IAppCaller // Generic read-only contract binding to access the raw methods on
}

// IAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAppTransactorRaw struct {
	Contract *IAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIApp creates a new instance of IApp, bound to a specific deployed contract.
func NewIApp(address common.Address, backend bind.ContractBackend) (*IApp, error) {
	contract, err := bindIApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IApp{IAppCaller: IAppCaller{contract: contract}, IAppTransactor: IAppTransactor{contract: contract}, IAppFilterer: IAppFilterer{contract: contract}}, nil
}

// NewIAppCaller creates a new read-only instance of IApp, bound to a specific deployed contract.
func NewIAppCaller(address common.Address, caller bind.ContractCaller) (*IAppCaller, error) {
	contract, err := bindIApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAppCaller{contract: contract}, nil
}

// NewIAppTransactor creates a new write-only instance of IApp, bound to a specific deployed contract.
func NewIAppTransactor(address common.Address, transactor bind.ContractTransactor) (*IAppTransactor, error) {
	contract, err := bindIApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAppTransactor{contract: contract}, nil
}

// NewIAppFilterer creates a new log filterer instance of IApp, bound to a specific deployed contract.
func NewIAppFilterer(address common.Address, filterer bind.ContractFilterer) (*IAppFilterer, error) {
	contract, err := bindIApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAppFilterer{contract: contract}, nil
}

// bindIApp binds a generic wrapper to an already deployed contract.
func bindIApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApp *IAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApp.Contract.IAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApp *IAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApp.Contract.IAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApp *IAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApp.Contract.IAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApp *IAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApp *IAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApp *IAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApp.Contract.contract.Transact(opts, method, params...)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 handleId, bytes msgBytes) returns(bytes responsePayload)
func (_IApp *IAppTransactor) HandleSynPackage(opts *bind.TransactOpts, handleId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApp.contract.Transact(opts, "handleSynPackage", handleId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 handleId, bytes msgBytes) returns(bytes responsePayload)
func (_IApp *IAppSession) HandleSynPackage(handleId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApp.Contract.HandleSynPackage(&_IApp.TransactOpts, handleId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 handleId, bytes msgBytes) returns(bytes responsePayload)
func (_IApp *IAppTransactorSession) HandleSynPackage(handleId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApp.Contract.HandleSynPackage(&_IApp.TransactOpts, handleId, msgBytes)
}

// MemoryMetaData contains all meta data concerning the Memory contract.
var MemoryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201e6c18961e25356acf663626a5c21d4f8fa6784c42a0e19910e97a034ebbdbfb64736f6c63430006040033",
}

// MemoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MemoryMetaData.ABI instead.
var MemoryABI = MemoryMetaData.ABI

// MemoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MemoryMetaData.Bin instead.
var MemoryBin = MemoryMetaData.Bin

// DeployMemory deploys a new Ethereum contract, binding an instance of Memory to it.
func DeployMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Memory, error) {
	parsed, err := MemoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// Memory is an auto generated Go binding around an Ethereum contract.
type Memory struct {
	MemoryCaller     // Read-only binding to the contract
	MemoryTransactor // Write-only binding to the contract
	MemoryFilterer   // Log filterer for contract events
}

// MemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemorySession struct {
	Contract     *Memory           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemoryCallerSession struct {
	Contract *MemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemoryTransactorSession struct {
	Contract     *MemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemoryRaw struct {
	Contract *Memory // Generic contract binding to access the raw methods on
}

// MemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemoryCallerRaw struct {
	Contract *MemoryCaller // Generic read-only contract binding to access the raw methods on
}

// MemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemoryTransactorRaw struct {
	Contract *MemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemory creates a new instance of Memory, bound to a specific deployed contract.
func NewMemory(address common.Address, backend bind.ContractBackend) (*Memory, error) {
	contract, err := bindMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// NewMemoryCaller creates a new read-only instance of Memory, bound to a specific deployed contract.
func NewMemoryCaller(address common.Address, caller bind.ContractCaller) (*MemoryCaller, error) {
	contract, err := bindMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryCaller{contract: contract}, nil
}

// NewMemoryTransactor creates a new write-only instance of Memory, bound to a specific deployed contract.
func NewMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MemoryTransactor, error) {
	contract, err := bindMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryTransactor{contract: contract}, nil
}

// NewMemoryFilterer creates a new log filterer instance of Memory, bound to a specific deployed contract.
func NewMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MemoryFilterer, error) {
	contract, err := bindMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemoryFilterer{contract: contract}, nil
}

// bindMemory binds a generic wrapper to an already deployed contract.
func bindMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MemoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.MemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transact(opts, method, params...)
}

// SystemMetaData contains all meta data concerning the System contract.
var SystemMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_EXCHANGE_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
		"0e77db70": "TOKEN_EXCHANGE_CONTRACT_ADDR()",
		"4fa689d8": "TRANSFER_IN_HANDLE()",
		"b6d28d2e": "TRANSFER_OUT_HANDLE()",
		"a78abc16": "alreadyInit()",
	},
	Bin: "0x608060405234801561001057600080fd5b5061010e806100206000396000f3fe6080604052348015600f57600080fd5b506004361060505760003560e01c80630e77db701460555780634fa689d814607757806351e80672146093578063a78abc16146099578063b6d28d2e1460b3575b600080fd5b605b60b9565b604080516001600160a01b039092168252519081900360200190f35b607d60bf565b6040805160ff9092168252519081900360200190f35b605b60c4565b609f60ca565b604080519115158252519081900360200190f35b607d60d3565b61020081565b600181565b61010081565b60005460ff1681565b60028156fea26469706673582212205b76cf7c6f7a0ecc58d246f4d78dbf67e39b2fd779a337b119ab2ec53df3f13d64736f6c63430006040033",
}

// SystemABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemMetaData.ABI instead.
var SystemABI = SystemMetaData.ABI

// Deprecated: Use SystemMetaData.Sigs instead.
// SystemFuncSigs maps the 4-byte function signature to its string representation.
var SystemFuncSigs = SystemMetaData.Sigs

// SystemBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemMetaData.Bin instead.
var SystemBin = SystemMetaData.Bin

// DeploySystem deploys a new Ethereum contract, binding an instance of System to it.
func DeploySystem(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *System, error) {
	parsed, err := SystemMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// System is an auto generated Go binding around an Ethereum contract.
type System struct {
	SystemCaller     // Read-only binding to the contract
	SystemTransactor // Write-only binding to the contract
	SystemFilterer   // Log filterer for contract events
}

// SystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemSession struct {
	Contract     *System           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemCallerSession struct {
	Contract *SystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemTransactorSession struct {
	Contract     *SystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRaw struct {
	Contract *System // Generic contract binding to access the raw methods on
}

// SystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemCallerRaw struct {
	Contract *SystemCaller // Generic read-only contract binding to access the raw methods on
}

// SystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemTransactorRaw struct {
	Contract *SystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystem creates a new instance of System, bound to a specific deployed contract.
func NewSystem(address common.Address, backend bind.ContractBackend) (*System, error) {
	contract, err := bindSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// NewSystemCaller creates a new read-only instance of System, bound to a specific deployed contract.
func NewSystemCaller(address common.Address, caller bind.ContractCaller) (*SystemCaller, error) {
	contract, err := bindSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemCaller{contract: contract}, nil
}

// NewSystemTransactor creates a new write-only instance of System, bound to a specific deployed contract.
func NewSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemTransactor, error) {
	contract, err := bindSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemTransactor{contract: contract}, nil
}

// NewSystemFilterer creates a new log filterer instance of System, bound to a specific deployed contract.
func NewSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemFilterer, error) {
	contract, err := bindSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemFilterer{contract: contract}, nil
}

// bindSystem binds a generic wrapper to an already deployed contract.
func bindSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.SystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.contract.Transact(opts, method, params...)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _System.Contract.CROSSCHAINCONTRACTADDR(&_System.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _System.Contract.CROSSCHAINCONTRACTADDR(&_System.CallOpts)
}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) TOKENEXCHANGECONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TOKEN_EXCHANGE_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) TOKENEXCHANGECONTRACTADDR() (common.Address, error) {
	return _System.Contract.TOKENEXCHANGECONTRACTADDR(&_System.CallOpts)
}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) TOKENEXCHANGECONTRACTADDR() (common.Address, error) {
	return _System.Contract.TOKENEXCHANGECONTRACTADDR(&_System.CallOpts)
}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_System *SystemCaller) TRANSFERINHANDLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TRANSFER_IN_HANDLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_System *SystemSession) TRANSFERINHANDLE() (uint8, error) {
	return _System.Contract.TRANSFERINHANDLE(&_System.CallOpts)
}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_System *SystemCallerSession) TRANSFERINHANDLE() (uint8, error) {
	return _System.Contract.TRANSFERINHANDLE(&_System.CallOpts)
}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_System *SystemCaller) TRANSFEROUTHANDLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TRANSFER_OUT_HANDLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_System *SystemSession) TRANSFEROUTHANDLE() (uint8, error) {
	return _System.Contract.TRANSFEROUTHANDLE(&_System.CallOpts)
}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_System *SystemCallerSession) TRANSFEROUTHANDLE() (uint8, error) {
	return _System.Contract.TRANSFEROUTHANDLE(&_System.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemSession) AlreadyInit() (bool, error) {
	return _System.Contract.AlreadyInit(&_System.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemCallerSession) AlreadyInit() (bool, error) {
	return _System.Contract.AlreadyInit(&_System.CallOpts)
}
