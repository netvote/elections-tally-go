// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BaseBallotABI is the input ABI used to generate the binding from.
const BaseBallotABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionPhase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeSelf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"metadataLocation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"location\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BaseBallot is an auto generated Go binding around an Ethereum contract.
type BaseBallot struct {
	BaseBallotCaller     // Read-only binding to the contract
	BaseBallotTransactor // Write-only binding to the contract
}

// BaseBallotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseBallotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseBallotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseBallotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseBallotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseBallotSession struct {
	Contract     *BaseBallot       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseBallotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseBallotCallerSession struct {
	Contract *BaseBallotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BaseBallotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseBallotTransactorSession struct {
	Contract     *BaseBallotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BaseBallotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseBallotRaw struct {
	Contract *BaseBallot // Generic contract binding to access the raw methods on
}

// BaseBallotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseBallotCallerRaw struct {
	Contract *BaseBallotCaller // Generic read-only contract binding to access the raw methods on
}

// BaseBallotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseBallotTransactorRaw struct {
	Contract *BaseBallotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseBallot creates a new instance of BaseBallot, bound to a specific deployed contract.
func NewBaseBallot(address common.Address, backend bind.ContractBackend) (*BaseBallot, error) {
	contract, err := bindBaseBallot(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseBallot{BaseBallotCaller: BaseBallotCaller{contract: contract}, BaseBallotTransactor: BaseBallotTransactor{contract: contract}}, nil
}

// NewBaseBallotCaller creates a new read-only instance of BaseBallot, bound to a specific deployed contract.
func NewBaseBallotCaller(address common.Address, caller bind.ContractCaller) (*BaseBallotCaller, error) {
	contract, err := bindBaseBallot(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BaseBallotCaller{contract: contract}, nil
}

// NewBaseBallotTransactor creates a new write-only instance of BaseBallot, bound to a specific deployed contract.
func NewBaseBallotTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseBallotTransactor, error) {
	contract, err := bindBaseBallot(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BaseBallotTransactor{contract: contract}, nil
}

// bindBaseBallot binds a generic wrapper to an already deployed contract.
func bindBaseBallot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseBallotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseBallot *BaseBallotRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseBallot.Contract.BaseBallotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseBallot *BaseBallotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBallot.Contract.BaseBallotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseBallot *BaseBallotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseBallot.Contract.BaseBallotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseBallot *BaseBallotCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseBallot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseBallot *BaseBallotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBallot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseBallot *BaseBallotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseBallot.Contract.contract.Transact(opts, method, params...)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BaseBallot *BaseBallotCaller) CheckConfig(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseBallot.contract.Call(opts, out, "checkConfig")
	return *ret0, err
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BaseBallot *BaseBallotSession) CheckConfig() (bool, error) {
	return _BaseBallot.Contract.CheckConfig(&_BaseBallot.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BaseBallot *BaseBallotCallerSession) CheckConfig() (bool, error) {
	return _BaseBallot.Contract.CheckConfig(&_BaseBallot.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BaseBallot *BaseBallotCaller) ElectionPhase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BaseBallot.contract.Call(opts, out, "electionPhase")
	return *ret0, err
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BaseBallot *BaseBallotSession) ElectionPhase() (uint8, error) {
	return _BaseBallot.Contract.ElectionPhase(&_BaseBallot.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BaseBallot *BaseBallotCallerSession) ElectionPhase() (uint8, error) {
	return _BaseBallot.Contract.ElectionPhase(&_BaseBallot.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BaseBallot *BaseBallotCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseBallot.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BaseBallot *BaseBallotSession) IsAdmin(addr common.Address) (bool, error) {
	return _BaseBallot.Contract.IsAdmin(&_BaseBallot.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BaseBallot *BaseBallotCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _BaseBallot.Contract.IsAdmin(&_BaseBallot.CallOpts, addr)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BaseBallot *BaseBallotCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseBallot.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BaseBallot *BaseBallotSession) IsClosed() (bool, error) {
	return _BaseBallot.Contract.IsClosed(&_BaseBallot.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BaseBallot *BaseBallotCallerSession) IsClosed() (bool, error) {
	return _BaseBallot.Contract.IsClosed(&_BaseBallot.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BaseBallot *BaseBallotCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseBallot.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BaseBallot *BaseBallotSession) IsLocked() (bool, error) {
	return _BaseBallot.Contract.IsLocked(&_BaseBallot.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BaseBallot *BaseBallotCallerSession) IsLocked() (bool, error) {
	return _BaseBallot.Contract.IsLocked(&_BaseBallot.CallOpts)
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_BaseBallot *BaseBallotCaller) MetadataLocation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseBallot.contract.Call(opts, out, "metadataLocation")
	return *ret0, err
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_BaseBallot *BaseBallotSession) MetadataLocation() (string, error) {
	return _BaseBallot.Contract.MetadataLocation(&_BaseBallot.CallOpts)
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_BaseBallot *BaseBallotCallerSession) MetadataLocation() (string, error) {
	return _BaseBallot.Contract.MetadataLocation(&_BaseBallot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseBallot *BaseBallotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseBallot.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseBallot *BaseBallotSession) Owner() (common.Address, error) {
	return _BaseBallot.Contract.Owner(&_BaseBallot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseBallot *BaseBallotCallerSession) Owner() (common.Address, error) {
	return _BaseBallot.Contract.Owner(&_BaseBallot.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BaseBallot *BaseBallotTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BaseBallot *BaseBallotSession) Abort() (*types.Transaction, error) {
	return _BaseBallot.Contract.Abort(&_BaseBallot.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BaseBallot *BaseBallotTransactorSession) Abort() (*types.Transaction, error) {
	return _BaseBallot.Contract.Abort(&_BaseBallot.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BaseBallot *BaseBallotTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BaseBallot *BaseBallotSession) Activate() (*types.Transaction, error) {
	return _BaseBallot.Contract.Activate(&_BaseBallot.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BaseBallot *BaseBallotTransactorSession) Activate() (*types.Transaction, error) {
	return _BaseBallot.Contract.Activate(&_BaseBallot.TransactOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BaseBallot *BaseBallotTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BaseBallot *BaseBallotSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _BaseBallot.Contract.AddAdmin(&_BaseBallot.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BaseBallot *BaseBallotTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _BaseBallot.Contract.AddAdmin(&_BaseBallot.TransactOpts, addr)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BaseBallot *BaseBallotTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BaseBallot *BaseBallotSession) Close() (*types.Transaction, error) {
	return _BaseBallot.Contract.Close(&_BaseBallot.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BaseBallot *BaseBallotTransactorSession) Close() (*types.Transaction, error) {
	return _BaseBallot.Contract.Close(&_BaseBallot.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BaseBallot *BaseBallotTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BaseBallot *BaseBallotSession) Lock() (*types.Transaction, error) {
	return _BaseBallot.Contract.Lock(&_BaseBallot.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BaseBallot *BaseBallotTransactorSession) Lock() (*types.Transaction, error) {
	return _BaseBallot.Contract.Lock(&_BaseBallot.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BaseBallot *BaseBallotTransactor) RemoveAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "removeAdmin", addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BaseBallot *BaseBallotSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _BaseBallot.Contract.RemoveAdmin(&_BaseBallot.TransactOpts, addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BaseBallot *BaseBallotTransactorSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _BaseBallot.Contract.RemoveAdmin(&_BaseBallot.TransactOpts, addr)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BaseBallot *BaseBallotTransactor) RemoveSelf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "removeSelf")
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BaseBallot *BaseBallotSession) RemoveSelf() (*types.Transaction, error) {
	return _BaseBallot.Contract.RemoveSelf(&_BaseBallot.TransactOpts)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BaseBallot *BaseBallotTransactorSession) RemoveSelf() (*types.Transaction, error) {
	return _BaseBallot.Contract.RemoveSelf(&_BaseBallot.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseBallot *BaseBallotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseBallot *BaseBallotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseBallot.Contract.TransferOwnership(&_BaseBallot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseBallot *BaseBallotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseBallot.Contract.TransferOwnership(&_BaseBallot.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BaseBallot *BaseBallotTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBallot.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BaseBallot *BaseBallotSession) Unlock() (*types.Transaction, error) {
	return _BaseBallot.Contract.Unlock(&_BaseBallot.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BaseBallot *BaseBallotTransactorSession) Unlock() (*types.Transaction, error) {
	return _BaseBallot.Contract.Unlock(&_BaseBallot.TransactOpts)
}
