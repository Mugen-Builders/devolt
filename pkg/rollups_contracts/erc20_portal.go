// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package json_rpc

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
	_ = abi.ConvertType
)

// ERC20PortalMetaData contains all meta data concerning the ERC20Portal contract.
var ERC20PortalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIInputBox\",\"name\":\"_inputBox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dapp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_execLayerData\",\"type\":\"bytes\"}],\"name\":\"depositERC20Tokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInputBox\",\"outputs\":[{\"internalType\":\"contractIInputBox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC20PortalABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20PortalMetaData.ABI instead.
var ERC20PortalABI = ERC20PortalMetaData.ABI

// ERC20Portal is an auto generated Go binding around an Ethereum contract.
type ERC20Portal struct {
	ERC20PortalCaller     // Read-only binding to the contract
	ERC20PortalTransactor // Write-only binding to the contract
	ERC20PortalFilterer   // Log filterer for contract events
}

// ERC20PortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PortalSession struct {
	Contract     *ERC20Portal      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20PortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PortalCallerSession struct {
	Contract *ERC20PortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20PortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PortalTransactorSession struct {
	Contract     *ERC20PortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20PortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PortalRaw struct {
	Contract *ERC20Portal // Generic contract binding to access the raw methods on
}

// ERC20PortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PortalCallerRaw struct {
	Contract *ERC20PortalCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PortalTransactorRaw struct {
	Contract *ERC20PortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Portal creates a new instance of ERC20Portal, bound to a specific deployed contract.
func NewERC20Portal(address common.Address, backend bind.ContractBackend) (*ERC20Portal, error) {
	contract, err := bindERC20Portal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Portal{ERC20PortalCaller: ERC20PortalCaller{contract: contract}, ERC20PortalTransactor: ERC20PortalTransactor{contract: contract}, ERC20PortalFilterer: ERC20PortalFilterer{contract: contract}}, nil
}

// NewERC20PortalCaller creates a new read-only instance of ERC20Portal, bound to a specific deployed contract.
func NewERC20PortalCaller(address common.Address, caller bind.ContractCaller) (*ERC20PortalCaller, error) {
	contract, err := bindERC20Portal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PortalCaller{contract: contract}, nil
}

// NewERC20PortalTransactor creates a new write-only instance of ERC20Portal, bound to a specific deployed contract.
func NewERC20PortalTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PortalTransactor, error) {
	contract, err := bindERC20Portal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PortalTransactor{contract: contract}, nil
}

// NewERC20PortalFilterer creates a new log filterer instance of ERC20Portal, bound to a specific deployed contract.
func NewERC20PortalFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PortalFilterer, error) {
	contract, err := bindERC20Portal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PortalFilterer{contract: contract}, nil
}

// bindERC20Portal binds a generic wrapper to an already deployed contract.
func bindERC20Portal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20PortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Portal *ERC20PortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Portal.Contract.ERC20PortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Portal *ERC20PortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Portal.Contract.ERC20PortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Portal *ERC20PortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Portal.Contract.ERC20PortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Portal *ERC20PortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Portal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Portal *ERC20PortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Portal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Portal *ERC20PortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Portal.Contract.contract.Transact(opts, method, params...)
}

// GetInputBox is a free data retrieval call binding the contract method 0x00aace9a.
//
// Solidity: function getInputBox() view returns(address)
func (_ERC20Portal *ERC20PortalCaller) GetInputBox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Portal.contract.Call(opts, &out, "getInputBox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInputBox is a free data retrieval call binding the contract method 0x00aace9a.
//
// Solidity: function getInputBox() view returns(address)
func (_ERC20Portal *ERC20PortalSession) GetInputBox() (common.Address, error) {
	return _ERC20Portal.Contract.GetInputBox(&_ERC20Portal.CallOpts)
}

// GetInputBox is a free data retrieval call binding the contract method 0x00aace9a.
//
// Solidity: function getInputBox() view returns(address)
func (_ERC20Portal *ERC20PortalCallerSession) GetInputBox() (common.Address, error) {
	return _ERC20Portal.Contract.GetInputBox(&_ERC20Portal.CallOpts)
}

// DepositERC20Tokens is a paid mutator transaction binding the contract method 0x95854b81.
//
// Solidity: function depositERC20Tokens(address _token, address _dapp, uint256 _amount, bytes _execLayerData) returns()
func (_ERC20Portal *ERC20PortalTransactor) DepositERC20Tokens(opts *bind.TransactOpts, _token common.Address, _dapp common.Address, _amount *big.Int, _execLayerData []byte) (*types.Transaction, error) {
	return _ERC20Portal.contract.Transact(opts, "depositERC20Tokens", _token, _dapp, _amount, _execLayerData)
}

// DepositERC20Tokens is a paid mutator transaction binding the contract method 0x95854b81.
//
// Solidity: function depositERC20Tokens(address _token, address _dapp, uint256 _amount, bytes _execLayerData) returns()
func (_ERC20Portal *ERC20PortalSession) DepositERC20Tokens(_token common.Address, _dapp common.Address, _amount *big.Int, _execLayerData []byte) (*types.Transaction, error) {
	return _ERC20Portal.Contract.DepositERC20Tokens(&_ERC20Portal.TransactOpts, _token, _dapp, _amount, _execLayerData)
}

// DepositERC20Tokens is a paid mutator transaction binding the contract method 0x95854b81.
//
// Solidity: function depositERC20Tokens(address _token, address _dapp, uint256 _amount, bytes _execLayerData) returns()
func (_ERC20Portal *ERC20PortalTransactorSession) DepositERC20Tokens(_token common.Address, _dapp common.Address, _amount *big.Int, _execLayerData []byte) (*types.Transaction, error) {
	return _ERC20Portal.Contract.DepositERC20Tokens(&_ERC20Portal.TransactOpts, _token, _dapp, _amount, _execLayerData)
}
