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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"crossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"packageDecodeError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"}],\"name\":\"receivePackage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAIL_ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYN_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"SendSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_EXCHANGE_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"encodePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"}],\"name\":\"handleCCPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"handleContractMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"sendCCPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b0355f5b": "ACK_PACKAGE()",
		"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
		"8cc8f561": "FAIL_ACK_PACKAGE()",
		"05e68258": "SYN_PACKAGE()",
		"8420782f": "SendSequenceMap(uint8)",
		"0e77db70": "TOKEN_EXCHANGE_CONTRACT_ADDR()",
		"4fa689d8": "TRANSFER_IN_HANDLE()",
		"b6d28d2e": "TRANSFER_OUT_HANDLE()",
		"a78abc16": "alreadyInit()",
		"3bdc47a6": "encodePayload(uint8,uint256,bytes)",
		"82fa3ab1": "handleCCPackage(bytes,uint64,uint8)",
		"460c8eb5": "handleContractMap(uint8)",
		"e1c7392a": "init()",
		"27018549": "sendCCPackage(uint8,bytes,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610b3b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806382fa3ab11161008c578063a78abc1611610066578063a78abc16146103c9578063b0355f5b146102f3578063b6d28d2e146103c1578063e1c7392a146103e5576100ea565b806382fa3ab1146103035780638420782f146103845780638cc8f561146103c1576100ea565b80633bdc47a6116100c85780633bdc47a6146101ab578063460c8eb5146102d35780634fa689d8146102f357806351e80672146102fb576100ea565b806305e68258146100ef5780630e77db701461010d5780632701854914610131575b600080fd5b6100f76103ed565b6040805160ff9092168252519081900360200190f35b6101156103f2565b604080516001600160a01b039092168252519081900360200190f35b6101a96004803603606081101561014757600080fd5b60ff8235169190810190604081016020820135600160201b81111561016b57600080fd5b82018360208201111561017d57600080fd5b803590602001918460018302840111600160201b8311171561019e57600080fd5b9193509150356103f8565b005b61025e600480360360608110156101c157600080fd5b60ff82351691602081013591810190606081016040820135600160201b8111156101ea57600080fd5b8201836020820111156101fc57600080fd5b803590602001918460018302840111600160201b8311171561021d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104ee945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610298578181015183820152602001610280565b50505050905090810190601f1680156102c55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610115600480360360208110156102e957600080fd5b503560ff16610504565b6100f761051f565b610115610524565b6101a96004803603606081101561031957600080fd5b810190602081018135600160201b81111561033357600080fd5b82018360208201111561034557600080fd5b803590602001918460018302840111600160201b8311171561036657600080fd5b9193509150803567ffffffffffffffff16906020013560ff1661052a565b6103a46004803603602081101561039a57600080fd5b503560ff1661092b565b6040805167ffffffffffffffff9092168252519081900360200190f35b6100f7610947565b6103d161094c565b604080519115158252519081900360200190f35b6101a9610955565b600081565b61020081565b60005460ff16610447576040805162461bcd60e51b815260206004820152601560248201527410dbdb9d1c9858dd081b9bdd081a5b9a5d081e595d605a1b604482015290519081900360640190fd5b60ff84166000908152600260209081526040808320548151601f870184900484028101840190925285825267ffffffffffffffff16926104b292849289926104ad92909188918b908b90819084018382808284376000920191909152506104ee92505050565b6109f1565b60ff949094166000908152600260205260409020805467ffffffffffffffff1916600190950167ffffffffffffffff1694909417909355505050565b6040805160008152602081019091529392505050565b6001602052600090815260409020546001600160a01b031681565b600181565b61010081565b60005460ff16610579576040805162461bcd60e51b815260206004820152601560248201527410dbdb9d1c9858dd081b9bdd081a5b9a5d081e595d605a1b604482015290519081900360640190fd5b600080600060606105bf88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610aec92505050565b935093509350935083610644578460ff168667ffffffffffffffff167f997d91a02d5e696ab94a30c03b00d5ada3207f80818db98c37f3f5bfdeda00128a8a60405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a350505050610925565b6040805160ff858116825291519187169167ffffffffffffffff8916917f591bae32fbc801fec820cde2a1ac597eb40c47c67d087a10cc51417dad79e1a8919081900360200190a360ff83166109205760ff85166000818152600160209081526040808320548151631182b87560e01b815260048101958652602481019283528651604482015286516001600160a01b03909216958695631182b875958d958a9593949093606490910192918601918190849084905b838110156107125781810151838201526020016106fa565b50505050905090810190601f16801561073f5780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561075f57600080fd5b505af192505050801561084357506040513d6000823e601f3d908101601f19168201604052602081101561079257600080fd5b8101908080516040519392919084600160201b8211156107b157600080fd5b9083019060208201858111156107c657600080fd5b8251600160201b8111828201881017156107df57600080fd5b82525081516020918201929091019080838360005b8381101561080c5781810151838201526020016107f4565b50505050905090810190601f1680156108395780820380516001836020036101000a031916815260200191505b5060405250505060015b61091c576040516000815260443d101561085f575060006108fc565b60046000803e60005160e01c6308c379a081146108805760009150506108fc565b60043d036004833e81513d602482011167ffffffffffffffff821117156108ac576000925050506108fc565b808301805167ffffffffffffffff8111156108ce5760009450505050506108fc565b8060208301013d86018111156108ec576000955050505050506108fc565b601f01601f191660405250925050505b80610907575061090d565b50610917565b3d6000803e3d6000fd5b61091e565b505b505b505050505b50505050565b60026020526000908152604090205467ffffffffffffffff1681565b600281565b60005460ff1681565b60005460ff16156109a5576040805162461bcd60e51b815260206004820152601560248201527410dbdb9d1c9858dd081b9bdd081a5b9a5d081e595d605a1b604482015290519081900360640190fd5b6001600081815260208290527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f80546001600160a01b031916610200179055805460ff19169091179055565b8160ff168367ffffffffffffffff167f7ddc2710ee05e4aecc08ed38e678b1a9d30b350f364a5c933bae2e9443961d0361026884604051808361ffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a6c578181015183820152602001610a54565b50505050905090810190601f168015610a995780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3505050565b60208110610acb578251825260209283019290910190601f1901610aac565b915181516020939093036101000a6000190180199091169216919091179052565b506040805160008082526020820190925260019282919056fea26469706673582212206328a5f37d246630cc615a7434c52e7e3f44d2c6b40c3904124b2b483136a62264736f6c63430006040033",
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

// SendSequenceMap is a free data retrieval call binding the contract method 0x8420782f.
//
// Solidity: function SendSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainCaller) SendSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "SendSequenceMap", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SendSequenceMap is a free data retrieval call binding the contract method 0x8420782f.
//
// Solidity: function SendSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainSession) SendSequenceMap(arg0 uint8) (uint64, error) {
	return _CrossChain.Contract.SendSequenceMap(&_CrossChain.CallOpts, arg0)
}

// SendSequenceMap is a free data retrieval call binding the contract method 0x8420782f.
//
// Solidity: function SendSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainCallerSession) SendSequenceMap(arg0 uint8) (uint64, error) {
	return _CrossChain.Contract.SendSequenceMap(&_CrossChain.CallOpts, arg0)
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

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) pure returns(bytes)
func (_CrossChain *CrossChainCaller) EncodePayload(opts *bind.CallOpts, packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "encodePayload", packageType, relayFee, msgBytes)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) pure returns(bytes)
func (_CrossChain *CrossChainSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _CrossChain.Contract.EncodePayload(&_CrossChain.CallOpts, packageType, relayFee, msgBytes)
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) pure returns(bytes)
func (_CrossChain *CrossChainCallerSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _CrossChain.Contract.EncodePayload(&_CrossChain.CallOpts, packageType, relayFee, msgBytes)
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

// SendCCPackage is a paid mutator transaction binding the contract method 0x27018549.
//
// Solidity: function sendCCPackage(uint8 handleId, bytes msgBytes, uint256 relayFee) returns()
func (_CrossChain *CrossChainTransactor) SendCCPackage(opts *bind.TransactOpts, handleId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _CrossChain.contract.Transact(opts, "sendCCPackage", handleId, msgBytes, relayFee)
}

// SendCCPackage is a paid mutator transaction binding the contract method 0x27018549.
//
// Solidity: function sendCCPackage(uint8 handleId, bytes msgBytes, uint256 relayFee) returns()
func (_CrossChain *CrossChainSession) SendCCPackage(handleId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _CrossChain.Contract.SendCCPackage(&_CrossChain.TransactOpts, handleId, msgBytes, relayFee)
}

// SendCCPackage is a paid mutator transaction binding the contract method 0x27018549.
//
// Solidity: function sendCCPackage(uint8 handleId, bytes msgBytes, uint256 relayFee) returns()
func (_CrossChain *CrossChainTransactorSession) SendCCPackage(handleId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _CrossChain.Contract.SendCCPackage(&_CrossChain.TransactOpts, handleId, msgBytes, relayFee)
}

// CrossChainCrossChainPackageIterator is returned from FilterCrossChainPackage and is used to iterate over the raw logs and unpacked data for CrossChainPackage events raised by the CrossChain contract.
type CrossChainCrossChainPackageIterator struct {
	Event *CrossChainCrossChainPackage // Event containing the contract specifics and raw log

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
func (it *CrossChainCrossChainPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainCrossChainPackage)
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
		it.Event = new(CrossChainCrossChainPackage)
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
func (it *CrossChainCrossChainPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainCrossChainPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainCrossChainPackage represents a CrossChainPackage event raised by the CrossChain contract.
type CrossChainCrossChainPackage struct {
	ChainId         uint16
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCrossChainPackage is a free log retrieval operation binding the contract event 0x7ddc2710ee05e4aecc08ed38e678b1a9d30b350f364a5c933bae2e9443961d03.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) FilterCrossChainPackage(opts *bind.FilterOpts, packageSequence []uint64, channelId []uint8) (*CrossChainCrossChainPackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _CrossChain.contract.FilterLogs(opts, "crossChainPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainCrossChainPackageIterator{contract: _CrossChain.contract, event: "crossChainPackage", logs: logs, sub: sub}, nil
}

// WatchCrossChainPackage is a free log subscription operation binding the contract event 0x7ddc2710ee05e4aecc08ed38e678b1a9d30b350f364a5c933bae2e9443961d03.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) WatchCrossChainPackage(opts *bind.WatchOpts, sink chan<- *CrossChainCrossChainPackage, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _CrossChain.contract.WatchLogs(opts, "crossChainPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainCrossChainPackage)
				if err := _CrossChain.contract.UnpackLog(event, "crossChainPackage", log); err != nil {
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

// ParseCrossChainPackage is a log parse operation binding the contract event 0x7ddc2710ee05e4aecc08ed38e678b1a9d30b350f364a5c933bae2e9443961d03.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) ParseCrossChainPackage(log types.Log) (*CrossChainCrossChainPackage, error) {
	event := new(CrossChainCrossChainPackage)
	if err := _CrossChain.contract.UnpackLog(event, "crossChainPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ICrossChainMetaData contains all meta data concerning the ICrossChain contract.
var ICrossChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"handle_id\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"sendCCPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"27018549": "sendCCPackage(uint8,bytes,uint256)",
	},
}

// ICrossChainABI is the input ABI used to generate the binding from.
// Deprecated: Use ICrossChainMetaData.ABI instead.
var ICrossChainABI = ICrossChainMetaData.ABI

// Deprecated: Use ICrossChainMetaData.Sigs instead.
// ICrossChainFuncSigs maps the 4-byte function signature to its string representation.
var ICrossChainFuncSigs = ICrossChainMetaData.Sigs

// ICrossChain is an auto generated Go binding around an Ethereum contract.
type ICrossChain struct {
	ICrossChainCaller     // Read-only binding to the contract
	ICrossChainTransactor // Write-only binding to the contract
	ICrossChainFilterer   // Log filterer for contract events
}

// ICrossChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICrossChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICrossChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICrossChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICrossChainSession struct {
	Contract     *ICrossChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICrossChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICrossChainCallerSession struct {
	Contract *ICrossChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ICrossChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICrossChainTransactorSession struct {
	Contract     *ICrossChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ICrossChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICrossChainRaw struct {
	Contract *ICrossChain // Generic contract binding to access the raw methods on
}

// ICrossChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICrossChainCallerRaw struct {
	Contract *ICrossChainCaller // Generic read-only contract binding to access the raw methods on
}

// ICrossChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICrossChainTransactorRaw struct {
	Contract *ICrossChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICrossChain creates a new instance of ICrossChain, bound to a specific deployed contract.
func NewICrossChain(address common.Address, backend bind.ContractBackend) (*ICrossChain, error) {
	contract, err := bindICrossChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICrossChain{ICrossChainCaller: ICrossChainCaller{contract: contract}, ICrossChainTransactor: ICrossChainTransactor{contract: contract}, ICrossChainFilterer: ICrossChainFilterer{contract: contract}}, nil
}

// NewICrossChainCaller creates a new read-only instance of ICrossChain, bound to a specific deployed contract.
func NewICrossChainCaller(address common.Address, caller bind.ContractCaller) (*ICrossChainCaller, error) {
	contract, err := bindICrossChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossChainCaller{contract: contract}, nil
}

// NewICrossChainTransactor creates a new write-only instance of ICrossChain, bound to a specific deployed contract.
func NewICrossChainTransactor(address common.Address, transactor bind.ContractTransactor) (*ICrossChainTransactor, error) {
	contract, err := bindICrossChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossChainTransactor{contract: contract}, nil
}

// NewICrossChainFilterer creates a new log filterer instance of ICrossChain, bound to a specific deployed contract.
func NewICrossChainFilterer(address common.Address, filterer bind.ContractFilterer) (*ICrossChainFilterer, error) {
	contract, err := bindICrossChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICrossChainFilterer{contract: contract}, nil
}

// bindICrossChain binds a generic wrapper to an already deployed contract.
func bindICrossChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICrossChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossChain *ICrossChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossChain.Contract.ICrossChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossChain *ICrossChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChain.Contract.ICrossChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossChain *ICrossChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossChain.Contract.ICrossChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossChain *ICrossChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossChain *ICrossChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossChain *ICrossChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossChain.Contract.contract.Transact(opts, method, params...)
}

// SendCCPackage is a paid mutator transaction binding the contract method 0x27018549.
//
// Solidity: function sendCCPackage(uint8 handle_id, bytes msgBytes, uint256 fee) returns()
func (_ICrossChain *ICrossChainTransactor) SendCCPackage(opts *bind.TransactOpts, handle_id uint8, msgBytes []byte, fee *big.Int) (*types.Transaction, error) {
	return _ICrossChain.contract.Transact(opts, "sendCCPackage", handle_id, msgBytes, fee)
}

// SendCCPackage is a paid mutator transaction binding the contract method 0x27018549.
//
// Solidity: function sendCCPackage(uint8 handle_id, bytes msgBytes, uint256 fee) returns()
func (_ICrossChain *ICrossChainSession) SendCCPackage(handle_id uint8, msgBytes []byte, fee *big.Int) (*types.Transaction, error) {
	return _ICrossChain.Contract.SendCCPackage(&_ICrossChain.TransactOpts, handle_id, msgBytes, fee)
}

// SendCCPackage is a paid mutator transaction binding the contract method 0x27018549.
//
// Solidity: function sendCCPackage(uint8 handle_id, bytes msgBytes, uint256 fee) returns()
func (_ICrossChain *ICrossChainTransactorSession) SendCCPackage(handle_id uint8, msgBytes []byte, fee *big.Int) (*types.Transaction, error) {
	return _ICrossChain.Contract.SendCCPackage(&_ICrossChain.TransactOpts, handle_id, msgBytes, fee)
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
