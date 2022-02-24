// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenexchange

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

// SystemMetaData contains all meta data concerning the System contract.
var SystemMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_EXCHANGE_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
		"0e77db70": "TOKEN_EXCHANGE_CONTRACT_ADDR()",
		"4fa689d8": "TRANSFER_IN_HANDLE()",
		"b6d28d2e": "TRANSFER_OUT_HANDLE()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060e18061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060465760003560e01c80630e77db7014604b5780634fa689d814606d57806351e80672146089578063b6d28d2e14608f575b600080fd5b60516095565b604080516001600160a01b039092168252519081900360200190f35b6073609b565b6040805160ff9092168252519081900360200190f35b605160a0565b607360a6565b61020081565b600181565b61010081565b60028156fea26469706673582212201367dd7187c6089feb9b3212cdaccb894ec0daafb988606762c8813df5cd72c964736f6c63430006040033",
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

// TokenExchangeMetaData contains all meta data concerning the TokenExchange contract.
var TokenExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_EXCHANGE_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_HANDLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"handleId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
		"0e77db70": "TOKEN_EXCHANGE_CONTRACT_ADDR()",
		"4fa689d8": "TRANSFER_IN_HANDLE()",
		"b6d28d2e": "TRANSFER_OUT_HANDLE()",
		"1182b875": "handleSynPackage(uint8,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b506103aa806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630e77db701461005c5780631182b875146100805780634fa689d81461016f57806351e806721461018d578063b6d28d2e14610195575b600080fd5b61006461019d565b604080516001600160a01b039092168252519081900360200190f35b6100fa6004803603604081101561009657600080fd5b60ff82351691908101906040810160208201356401000000008111156100bb57600080fd5b8201836020820111156100cd57600080fd5b803590602001918460018302840111640100000000831117156100ef57600080fd5b5090925090506101a3565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561013457818101518382015260200161011c565b50505050905090810190601f1680156101615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610177610231565b6040805160ff9092168252519081900360200190f35b610064610236565b61017761023c565b61020081565b606060ff8416600114156101f7576101f083838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061024192505050565b905061022a565b6040805162461bcd60e51b815260206004820152600360248201526262756760e81b604482015290519081900360640190fd5b9392505050565b600181565b61010081565b600281565b6060600061024d610354565b610256846102c7565b91509150816102a3576040805162461bcd60e51b81526020600482015260146024820152731b5cd9d09e5d195cc8191958dbd9194819985a5b60621b604482015290519081900360640190fd5b60006102ae826102e4565b5050604080516000815260208101909152949350505050565b60006102d1610354565b6102d9610354565b600192509050915091565b6040516000907396932b7a373d8586c4a2d3c98517803ff2818cec906301312d009062bc614e9084818181858888f193505050503d8060008114610344576040519150601f19603f3d011682016040523d82523d6000602084013e610349565b606091505b506000949350505050565b60408051606081018252600080825260208201819052918101919091529056fea26469706673582212202dccea49a1d11ac552b7e785f6b1ee21c0249001fda5248831da8fdf992fd78964736f6c63430006040033",
}

// TokenExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenExchangeMetaData.ABI instead.
var TokenExchangeABI = TokenExchangeMetaData.ABI

// Deprecated: Use TokenExchangeMetaData.Sigs instead.
// TokenExchangeFuncSigs maps the 4-byte function signature to its string representation.
var TokenExchangeFuncSigs = TokenExchangeMetaData.Sigs

// TokenExchangeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenExchangeMetaData.Bin instead.
var TokenExchangeBin = TokenExchangeMetaData.Bin

// DeployTokenExchange deploys a new Ethereum contract, binding an instance of TokenExchange to it.
func DeployTokenExchange(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenExchange, error) {
	parsed, err := TokenExchangeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenExchangeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenExchange{TokenExchangeCaller: TokenExchangeCaller{contract: contract}, TokenExchangeTransactor: TokenExchangeTransactor{contract: contract}, TokenExchangeFilterer: TokenExchangeFilterer{contract: contract}}, nil
}

// TokenExchange is an auto generated Go binding around an Ethereum contract.
type TokenExchange struct {
	TokenExchangeCaller     // Read-only binding to the contract
	TokenExchangeTransactor // Write-only binding to the contract
	TokenExchangeFilterer   // Log filterer for contract events
}

// TokenExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenExchangeSession struct {
	Contract     *TokenExchange    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenExchangeCallerSession struct {
	Contract *TokenExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenExchangeTransactorSession struct {
	Contract     *TokenExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenExchangeRaw struct {
	Contract *TokenExchange // Generic contract binding to access the raw methods on
}

// TokenExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenExchangeCallerRaw struct {
	Contract *TokenExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// TokenExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenExchangeTransactorRaw struct {
	Contract *TokenExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenExchange creates a new instance of TokenExchange, bound to a specific deployed contract.
func NewTokenExchange(address common.Address, backend bind.ContractBackend) (*TokenExchange, error) {
	contract, err := bindTokenExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenExchange{TokenExchangeCaller: TokenExchangeCaller{contract: contract}, TokenExchangeTransactor: TokenExchangeTransactor{contract: contract}, TokenExchangeFilterer: TokenExchangeFilterer{contract: contract}}, nil
}

// NewTokenExchangeCaller creates a new read-only instance of TokenExchange, bound to a specific deployed contract.
func NewTokenExchangeCaller(address common.Address, caller bind.ContractCaller) (*TokenExchangeCaller, error) {
	contract, err := bindTokenExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeCaller{contract: contract}, nil
}

// NewTokenExchangeTransactor creates a new write-only instance of TokenExchange, bound to a specific deployed contract.
func NewTokenExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenExchangeTransactor, error) {
	contract, err := bindTokenExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeTransactor{contract: contract}, nil
}

// NewTokenExchangeFilterer creates a new log filterer instance of TokenExchange, bound to a specific deployed contract.
func NewTokenExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenExchangeFilterer, error) {
	contract, err := bindTokenExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeFilterer{contract: contract}, nil
}

// bindTokenExchange binds a generic wrapper to an already deployed contract.
func bindTokenExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenExchange *TokenExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenExchange.Contract.TokenExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenExchange *TokenExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.Contract.TokenExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenExchange *TokenExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenExchange.Contract.TokenExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenExchange *TokenExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenExchange *TokenExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenExchange *TokenExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenExchange.Contract.contract.Transact(opts, method, params...)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_TokenExchange *TokenExchangeCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_TokenExchange *TokenExchangeSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _TokenExchange.Contract.CROSSCHAINCONTRACTADDR(&_TokenExchange.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _TokenExchange.Contract.CROSSCHAINCONTRACTADDR(&_TokenExchange.CallOpts)
}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_TokenExchange *TokenExchangeCaller) TOKENEXCHANGECONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "TOKEN_EXCHANGE_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_TokenExchange *TokenExchangeSession) TOKENEXCHANGECONTRACTADDR() (common.Address, error) {
	return _TokenExchange.Contract.TOKENEXCHANGECONTRACTADDR(&_TokenExchange.CallOpts)
}

// TOKENEXCHANGECONTRACTADDR is a free data retrieval call binding the contract method 0x0e77db70.
//
// Solidity: function TOKEN_EXCHANGE_CONTRACT_ADDR() view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) TOKENEXCHANGECONTRACTADDR() (common.Address, error) {
	return _TokenExchange.Contract.TOKENEXCHANGECONTRACTADDR(&_TokenExchange.CallOpts)
}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_TokenExchange *TokenExchangeCaller) TRANSFERINHANDLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "TRANSFER_IN_HANDLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_TokenExchange *TokenExchangeSession) TRANSFERINHANDLE() (uint8, error) {
	return _TokenExchange.Contract.TRANSFERINHANDLE(&_TokenExchange.CallOpts)
}

// TRANSFERINHANDLE is a free data retrieval call binding the contract method 0x4fa689d8.
//
// Solidity: function TRANSFER_IN_HANDLE() view returns(uint8)
func (_TokenExchange *TokenExchangeCallerSession) TRANSFERINHANDLE() (uint8, error) {
	return _TokenExchange.Contract.TRANSFERINHANDLE(&_TokenExchange.CallOpts)
}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_TokenExchange *TokenExchangeCaller) TRANSFEROUTHANDLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "TRANSFER_OUT_HANDLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_TokenExchange *TokenExchangeSession) TRANSFEROUTHANDLE() (uint8, error) {
	return _TokenExchange.Contract.TRANSFEROUTHANDLE(&_TokenExchange.CallOpts)
}

// TRANSFEROUTHANDLE is a free data retrieval call binding the contract method 0xb6d28d2e.
//
// Solidity: function TRANSFER_OUT_HANDLE() view returns(uint8)
func (_TokenExchange *TokenExchangeCallerSession) TRANSFEROUTHANDLE() (uint8, error) {
	return _TokenExchange.Contract.TRANSFEROUTHANDLE(&_TokenExchange.CallOpts)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 handleId, bytes msgBytes) returns(bytes)
func (_TokenExchange *TokenExchangeTransactor) HandleSynPackage(opts *bind.TransactOpts, handleId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "handleSynPackage", handleId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 handleId, bytes msgBytes) returns(bytes)
func (_TokenExchange *TokenExchangeSession) HandleSynPackage(handleId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _TokenExchange.Contract.HandleSynPackage(&_TokenExchange.TransactOpts, handleId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 handleId, bytes msgBytes) returns(bytes)
func (_TokenExchange *TokenExchangeTransactorSession) HandleSynPackage(handleId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _TokenExchange.Contract.HandleSynPackage(&_TokenExchange.TransactOpts, handleId, msgBytes)
}
