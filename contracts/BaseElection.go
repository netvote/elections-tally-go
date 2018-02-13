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

// BaseElectionABI is the input ABI used to generate the binding from.
const BaseElectionABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionPhase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"addAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionType\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"privateKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeSelf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"removeAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowVoteUpdates\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deductVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"setPrivateKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"hashedUserId\",\"type\":\"bytes32\"},{\"name\":\"allowanceAddress\",\"type\":\"address\"},{\"name\":\"acct\",\"type\":\"address\"},{\"name\":\"allowUpdates\",\"type\":\"bool\"},{\"name\":\"revealer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"KeyReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BaseElection is an auto generated Go binding around an Ethereum contract.
type BaseElection struct {
	BaseElectionCaller     // Read-only binding to the contract
	BaseElectionTransactor // Write-only binding to the contract
}

// BaseElectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseElectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseElectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseElectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseElectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseElectionSession struct {
	Contract     *BaseElection     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseElectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseElectionCallerSession struct {
	Contract *BaseElectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BaseElectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseElectionTransactorSession struct {
	Contract     *BaseElectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseElectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseElectionRaw struct {
	Contract *BaseElection // Generic contract binding to access the raw methods on
}

// BaseElectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseElectionCallerRaw struct {
	Contract *BaseElectionCaller // Generic read-only contract binding to access the raw methods on
}

// BaseElectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseElectionTransactorRaw struct {
	Contract *BaseElectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseElection creates a new instance of BaseElection, bound to a specific deployed contract.
func NewBaseElection(address common.Address, backend bind.ContractBackend) (*BaseElection, error) {
	contract, err := bindBaseElection(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseElection{BaseElectionCaller: BaseElectionCaller{contract: contract}, BaseElectionTransactor: BaseElectionTransactor{contract: contract}}, nil
}

// NewBaseElectionCaller creates a new read-only instance of BaseElection, bound to a specific deployed contract.
func NewBaseElectionCaller(address common.Address, caller bind.ContractCaller) (*BaseElectionCaller, error) {
	contract, err := bindBaseElection(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BaseElectionCaller{contract: contract}, nil
}

// NewBaseElectionTransactor creates a new write-only instance of BaseElection, bound to a specific deployed contract.
func NewBaseElectionTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseElectionTransactor, error) {
	contract, err := bindBaseElection(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BaseElectionTransactor{contract: contract}, nil
}

// bindBaseElection binds a generic wrapper to an already deployed contract.
func bindBaseElection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseElectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseElection *BaseElectionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseElection.Contract.BaseElectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseElection *BaseElectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.Contract.BaseElectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseElection *BaseElectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseElection.Contract.BaseElectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseElection *BaseElectionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseElection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseElection *BaseElectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseElection *BaseElectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseElection.Contract.contract.Transact(opts, method, params...)
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_BaseElection *BaseElectionCaller) AllowVoteUpdates(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "allowVoteUpdates")
	return *ret0, err
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_BaseElection *BaseElectionSession) AllowVoteUpdates() (bool, error) {
	return _BaseElection.Contract.AllowVoteUpdates(&_BaseElection.CallOpts)
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_BaseElection *BaseElectionCallerSession) AllowVoteUpdates() (bool, error) {
	return _BaseElection.Contract.AllowVoteUpdates(&_BaseElection.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BaseElection *BaseElectionCaller) CheckConfig(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "checkConfig")
	return *ret0, err
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BaseElection *BaseElectionSession) CheckConfig() (bool, error) {
	return _BaseElection.Contract.CheckConfig(&_BaseElection.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BaseElection *BaseElectionCallerSession) CheckConfig() (bool, error) {
	return _BaseElection.Contract.CheckConfig(&_BaseElection.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BaseElection *BaseElectionCaller) ElectionPhase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "electionPhase")
	return *ret0, err
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BaseElection *BaseElectionSession) ElectionPhase() (uint8, error) {
	return _BaseElection.Contract.ElectionPhase(&_BaseElection.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BaseElection *BaseElectionCallerSession) ElectionPhase() (uint8, error) {
	return _BaseElection.Contract.ElectionPhase(&_BaseElection.CallOpts)
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_BaseElection *BaseElectionCaller) ElectionType(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "electionType")
	return *ret0, err
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_BaseElection *BaseElectionSession) ElectionType() (string, error) {
	return _BaseElection.Contract.ElectionType(&_BaseElection.CallOpts)
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_BaseElection *BaseElectionCallerSession) ElectionType() (string, error) {
	return _BaseElection.Contract.ElectionType(&_BaseElection.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BaseElection *BaseElectionCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BaseElection *BaseElectionSession) IsAdmin(addr common.Address) (bool, error) {
	return _BaseElection.Contract.IsAdmin(&_BaseElection.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BaseElection *BaseElectionCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _BaseElection.Contract.IsAdmin(&_BaseElection.CallOpts, addr)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BaseElection *BaseElectionCaller) IsAuthorized(opts *bind.CallOpts, entry [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "isAuthorized", entry)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BaseElection *BaseElectionSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _BaseElection.Contract.IsAuthorized(&_BaseElection.CallOpts, entry)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BaseElection *BaseElectionCallerSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _BaseElection.Contract.IsAuthorized(&_BaseElection.CallOpts, entry)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BaseElection *BaseElectionCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BaseElection *BaseElectionSession) IsClosed() (bool, error) {
	return _BaseElection.Contract.IsClosed(&_BaseElection.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BaseElection *BaseElectionCallerSession) IsClosed() (bool, error) {
	return _BaseElection.Contract.IsClosed(&_BaseElection.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BaseElection *BaseElectionCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BaseElection *BaseElectionSession) IsLocked() (bool, error) {
	return _BaseElection.Contract.IsLocked(&_BaseElection.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BaseElection *BaseElectionCallerSession) IsLocked() (bool, error) {
	return _BaseElection.Contract.IsLocked(&_BaseElection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseElection *BaseElectionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseElection *BaseElectionSession) Owner() (common.Address, error) {
	return _BaseElection.Contract.Owner(&_BaseElection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseElection *BaseElectionCallerSession) Owner() (common.Address, error) {
	return _BaseElection.Contract.Owner(&_BaseElection.CallOpts)
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_BaseElection *BaseElectionCaller) PrivateKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "privateKey")
	return *ret0, err
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_BaseElection *BaseElectionSession) PrivateKey() (string, error) {
	return _BaseElection.Contract.PrivateKey(&_BaseElection.CallOpts)
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_BaseElection *BaseElectionCallerSession) PrivateKey() (string, error) {
	return _BaseElection.Contract.PrivateKey(&_BaseElection.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_BaseElection *BaseElectionCaller) PublicKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseElection.contract.Call(opts, out, "publicKey")
	return *ret0, err
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_BaseElection *BaseElectionSession) PublicKey() (string, error) {
	return _BaseElection.Contract.PublicKey(&_BaseElection.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_BaseElection *BaseElectionCallerSession) PublicKey() (string, error) {
	return _BaseElection.Contract.PublicKey(&_BaseElection.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BaseElection *BaseElectionTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BaseElection *BaseElectionSession) Abort() (*types.Transaction, error) {
	return _BaseElection.Contract.Abort(&_BaseElection.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BaseElection *BaseElectionTransactorSession) Abort() (*types.Transaction, error) {
	return _BaseElection.Contract.Abort(&_BaseElection.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BaseElection *BaseElectionTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BaseElection *BaseElectionSession) Activate() (*types.Transaction, error) {
	return _BaseElection.Contract.Activate(&_BaseElection.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BaseElection *BaseElectionTransactorSession) Activate() (*types.Transaction, error) {
	return _BaseElection.Contract.Activate(&_BaseElection.TransactOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BaseElection *BaseElectionTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BaseElection *BaseElectionSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _BaseElection.Contract.AddAdmin(&_BaseElection.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BaseElection *BaseElectionTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _BaseElection.Contract.AddAdmin(&_BaseElection.TransactOpts, addr)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BaseElection *BaseElectionTransactor) AddAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "addAuthorized", entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BaseElection *BaseElectionSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BaseElection.Contract.AddAuthorized(&_BaseElection.TransactOpts, entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BaseElection *BaseElectionTransactorSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BaseElection.Contract.AddAuthorized(&_BaseElection.TransactOpts, entry)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BaseElection *BaseElectionTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BaseElection *BaseElectionSession) Close() (*types.Transaction, error) {
	return _BaseElection.Contract.Close(&_BaseElection.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BaseElection *BaseElectionTransactorSession) Close() (*types.Transaction, error) {
	return _BaseElection.Contract.Close(&_BaseElection.TransactOpts)
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_BaseElection *BaseElectionTransactor) DeductVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "deductVote")
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_BaseElection *BaseElectionSession) DeductVote() (*types.Transaction, error) {
	return _BaseElection.Contract.DeductVote(&_BaseElection.TransactOpts)
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_BaseElection *BaseElectionTransactorSession) DeductVote() (*types.Transaction, error) {
	return _BaseElection.Contract.DeductVote(&_BaseElection.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BaseElection *BaseElectionTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BaseElection *BaseElectionSession) Lock() (*types.Transaction, error) {
	return _BaseElection.Contract.Lock(&_BaseElection.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BaseElection *BaseElectionTransactorSession) Lock() (*types.Transaction, error) {
	return _BaseElection.Contract.Lock(&_BaseElection.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BaseElection *BaseElectionTransactor) RemoveAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "removeAdmin", addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BaseElection *BaseElectionSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _BaseElection.Contract.RemoveAdmin(&_BaseElection.TransactOpts, addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BaseElection *BaseElectionTransactorSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _BaseElection.Contract.RemoveAdmin(&_BaseElection.TransactOpts, addr)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BaseElection *BaseElectionTransactor) RemoveAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "removeAuthorized", entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BaseElection *BaseElectionSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BaseElection.Contract.RemoveAuthorized(&_BaseElection.TransactOpts, entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BaseElection *BaseElectionTransactorSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BaseElection.Contract.RemoveAuthorized(&_BaseElection.TransactOpts, entry)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BaseElection *BaseElectionTransactor) RemoveSelf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "removeSelf")
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BaseElection *BaseElectionSession) RemoveSelf() (*types.Transaction, error) {
	return _BaseElection.Contract.RemoveSelf(&_BaseElection.TransactOpts)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BaseElection *BaseElectionTransactorSession) RemoveSelf() (*types.Transaction, error) {
	return _BaseElection.Contract.RemoveSelf(&_BaseElection.TransactOpts)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_BaseElection *BaseElectionTransactor) SetPrivateKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "setPrivateKey", key)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_BaseElection *BaseElectionSession) SetPrivateKey(key string) (*types.Transaction, error) {
	return _BaseElection.Contract.SetPrivateKey(&_BaseElection.TransactOpts, key)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_BaseElection *BaseElectionTransactorSession) SetPrivateKey(key string) (*types.Transaction, error) {
	return _BaseElection.Contract.SetPrivateKey(&_BaseElection.TransactOpts, key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_BaseElection *BaseElectionTransactor) SetPublicKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "setPublicKey", key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_BaseElection *BaseElectionSession) SetPublicKey(key string) (*types.Transaction, error) {
	return _BaseElection.Contract.SetPublicKey(&_BaseElection.TransactOpts, key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_BaseElection *BaseElectionTransactorSession) SetPublicKey(key string) (*types.Transaction, error) {
	return _BaseElection.Contract.SetPublicKey(&_BaseElection.TransactOpts, key)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseElection *BaseElectionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseElection *BaseElectionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseElection.Contract.TransferOwnership(&_BaseElection.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseElection *BaseElectionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseElection.Contract.TransferOwnership(&_BaseElection.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BaseElection *BaseElectionTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseElection.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BaseElection *BaseElectionSession) Unlock() (*types.Transaction, error) {
	return _BaseElection.Contract.Unlock(&_BaseElection.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BaseElection *BaseElectionTransactorSession) Unlock() (*types.Transaction, error) {
	return _BaseElection.Contract.Unlock(&_BaseElection.TransactOpts)
}
