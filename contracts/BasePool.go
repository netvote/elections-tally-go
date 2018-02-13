// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BasePoolABI is the input ABI used to generate the binding from.
const BasePoolABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"ballotExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionPhase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createdBy\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"addAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"election\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteId\",\"type\":\"bytes32\"},{\"name\":\"vote\",\"type\":\"string\"},{\"name\":\"passphrase\",\"type\":\"string\"}],\"name\":\"castVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeSelf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"removeAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVoteAt\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBallotCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"removeBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteId\",\"type\":\"bytes32\"},{\"name\":\"vote\",\"type\":\"string\"},{\"name\":\"passphrase\",\"type\":\"string\"}],\"name\":\"updateVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"getBallotIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVoteCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"addBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"hashedUserId\",\"type\":\"bytes32\"},{\"name\":\"el\",\"type\":\"address\"},{\"name\":\"gw\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteId\",\"type\":\"bytes32\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteId\",\"type\":\"bytes32\"}],\"name\":\"UpdateVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BasePool is an auto generated Go binding around an Ethereum contract.
type BasePool struct {
	BasePoolCaller     // Read-only binding to the contract
	BasePoolTransactor // Write-only binding to the contract
}

// BasePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasePoolSession struct {
	Contract     *BasePool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasePoolCallerSession struct {
	Contract *BasePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BasePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasePoolTransactorSession struct {
	Contract     *BasePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BasePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasePoolRaw struct {
	Contract *BasePool // Generic contract binding to access the raw methods on
}

// BasePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasePoolCallerRaw struct {
	Contract *BasePoolCaller // Generic read-only contract binding to access the raw methods on
}

// BasePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasePoolTransactorRaw struct {
	Contract *BasePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasePool creates a new instance of BasePool, bound to a specific deployed contract.
func NewBasePool(address common.Address, backend bind.ContractBackend) (*BasePool, error) {
	contract, err := bindBasePool(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasePool{BasePoolCaller: BasePoolCaller{contract: contract}, BasePoolTransactor: BasePoolTransactor{contract: contract}}, nil
}

// NewBasePoolCaller creates a new read-only instance of BasePool, bound to a specific deployed contract.
func NewBasePoolCaller(address common.Address, caller bind.ContractCaller) (*BasePoolCaller, error) {
	contract, err := bindBasePool(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BasePoolCaller{contract: contract}, nil
}

// NewBasePoolTransactor creates a new write-only instance of BasePool, bound to a specific deployed contract.
func NewBasePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BasePoolTransactor, error) {
	contract, err := bindBasePool(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BasePoolTransactor{contract: contract}, nil
}

// bindBasePool binds a generic wrapper to an already deployed contract.
func bindBasePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasePool *BasePoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasePool.Contract.BasePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasePool *BasePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePool.Contract.BasePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasePool *BasePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasePool.Contract.BasePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasePool *BasePoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasePool *BasePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasePool *BasePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasePool.Contract.contract.Transact(opts, method, params...)
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_BasePool *BasePoolCaller) BallotExists(opts *bind.CallOpts, b common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "ballotExists", b)
	return *ret0, err
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_BasePool *BasePoolSession) BallotExists(b common.Address) (bool, error) {
	return _BasePool.Contract.BallotExists(&_BasePool.CallOpts, b)
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_BasePool *BasePoolCallerSession) BallotExists(b common.Address) (bool, error) {
	return _BasePool.Contract.BallotExists(&_BasePool.CallOpts, b)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BasePool *BasePoolCaller) CheckConfig(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "checkConfig")
	return *ret0, err
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BasePool *BasePoolSession) CheckConfig() (bool, error) {
	return _BasePool.Contract.CheckConfig(&_BasePool.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BasePool *BasePoolCallerSession) CheckConfig() (bool, error) {
	return _BasePool.Contract.CheckConfig(&_BasePool.CallOpts)
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_BasePool *BasePoolCaller) CreatedBy(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "createdBy")
	return *ret0, err
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_BasePool *BasePoolSession) CreatedBy() (string, error) {
	return _BasePool.Contract.CreatedBy(&_BasePool.CallOpts)
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_BasePool *BasePoolCallerSession) CreatedBy() (string, error) {
	return _BasePool.Contract.CreatedBy(&_BasePool.CallOpts)
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_BasePool *BasePoolCaller) Election(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "election")
	return *ret0, err
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_BasePool *BasePoolSession) Election() (common.Address, error) {
	return _BasePool.Contract.Election(&_BasePool.CallOpts)
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_BasePool *BasePoolCallerSession) Election() (common.Address, error) {
	return _BasePool.Contract.Election(&_BasePool.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BasePool *BasePoolCaller) ElectionPhase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "electionPhase")
	return *ret0, err
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BasePool *BasePoolSession) ElectionPhase() (uint8, error) {
	return _BasePool.Contract.ElectionPhase(&_BasePool.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BasePool *BasePoolCallerSession) ElectionPhase() (uint8, error) {
	return _BasePool.Contract.ElectionPhase(&_BasePool.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_BasePool *BasePoolCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "gateway")
	return *ret0, err
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_BasePool *BasePoolSession) Gateway() (common.Address, error) {
	return _BasePool.Contract.Gateway(&_BasePool.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_BasePool *BasePoolCallerSession) Gateway() (common.Address, error) {
	return _BasePool.Contract.Gateway(&_BasePool.CallOpts)
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_BasePool *BasePoolCaller) GetBallot(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "getBallot", index)
	return *ret0, err
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_BasePool *BasePoolSession) GetBallot(index *big.Int) (common.Address, error) {
	return _BasePool.Contract.GetBallot(&_BasePool.CallOpts, index)
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_BasePool *BasePoolCallerSession) GetBallot(index *big.Int) (common.Address, error) {
	return _BasePool.Contract.GetBallot(&_BasePool.CallOpts, index)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_BasePool *BasePoolCaller) GetBallotCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "getBallotCount")
	return *ret0, err
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_BasePool *BasePoolSession) GetBallotCount() (*big.Int, error) {
	return _BasePool.Contract.GetBallotCount(&_BasePool.CallOpts)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_BasePool *BasePoolCallerSession) GetBallotCount() (*big.Int, error) {
	return _BasePool.Contract.GetBallotCount(&_BasePool.CallOpts)
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_BasePool *BasePoolCaller) GetBallotIndex(opts *bind.CallOpts, b common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "getBallotIndex", b)
	return *ret0, err
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_BasePool *BasePoolSession) GetBallotIndex(b common.Address) (*big.Int, error) {
	return _BasePool.Contract.GetBallotIndex(&_BasePool.CallOpts, b)
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_BasePool *BasePoolCallerSession) GetBallotIndex(b common.Address) (*big.Int, error) {
	return _BasePool.Contract.GetBallotIndex(&_BasePool.CallOpts, b)
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_BasePool *BasePoolCaller) GetVoteAt(opts *bind.CallOpts, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "getVoteAt", index)
	return *ret0, err
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_BasePool *BasePoolSession) GetVoteAt(index *big.Int) (string, error) {
	return _BasePool.Contract.GetVoteAt(&_BasePool.CallOpts, index)
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_BasePool *BasePoolCallerSession) GetVoteAt(index *big.Int) (string, error) {
	return _BasePool.Contract.GetVoteAt(&_BasePool.CallOpts, index)
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_BasePool *BasePoolCaller) GetVoteCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "getVoteCount")
	return *ret0, err
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_BasePool *BasePoolSession) GetVoteCount() (*big.Int, error) {
	return _BasePool.Contract.GetVoteCount(&_BasePool.CallOpts)
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_BasePool *BasePoolCallerSession) GetVoteCount() (*big.Int, error) {
	return _BasePool.Contract.GetVoteCount(&_BasePool.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BasePool *BasePoolCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BasePool *BasePoolSession) IsAdmin(addr common.Address) (bool, error) {
	return _BasePool.Contract.IsAdmin(&_BasePool.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BasePool *BasePoolCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _BasePool.Contract.IsAdmin(&_BasePool.CallOpts, addr)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BasePool *BasePoolCaller) IsAuthorized(opts *bind.CallOpts, entry [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "isAuthorized", entry)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BasePool *BasePoolSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _BasePool.Contract.IsAuthorized(&_BasePool.CallOpts, entry)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BasePool *BasePoolCallerSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _BasePool.Contract.IsAuthorized(&_BasePool.CallOpts, entry)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BasePool *BasePoolCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BasePool *BasePoolSession) IsClosed() (bool, error) {
	return _BasePool.Contract.IsClosed(&_BasePool.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BasePool *BasePoolCallerSession) IsClosed() (bool, error) {
	return _BasePool.Contract.IsClosed(&_BasePool.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BasePool *BasePoolCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BasePool *BasePoolSession) IsLocked() (bool, error) {
	return _BasePool.Contract.IsLocked(&_BasePool.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BasePool *BasePoolCallerSession) IsLocked() (bool, error) {
	return _BasePool.Contract.IsLocked(&_BasePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasePool *BasePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasePool *BasePoolSession) Owner() (common.Address, error) {
	return _BasePool.Contract.Owner(&_BasePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasePool *BasePoolCallerSession) Owner() (common.Address, error) {
	return _BasePool.Contract.Owner(&_BasePool.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_BasePool *BasePoolCaller) Votes(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasePool.contract.Call(opts, out, "votes", arg0)
	return *ret0, err
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_BasePool *BasePoolSession) Votes(arg0 [32]byte) (string, error) {
	return _BasePool.Contract.Votes(&_BasePool.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_BasePool *BasePoolCallerSession) Votes(arg0 [32]byte) (string, error) {
	return _BasePool.Contract.Votes(&_BasePool.CallOpts, arg0)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BasePool *BasePoolTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BasePool *BasePoolSession) Abort() (*types.Transaction, error) {
	return _BasePool.Contract.Abort(&_BasePool.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BasePool *BasePoolTransactorSession) Abort() (*types.Transaction, error) {
	return _BasePool.Contract.Abort(&_BasePool.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BasePool *BasePoolTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BasePool *BasePoolSession) Activate() (*types.Transaction, error) {
	return _BasePool.Contract.Activate(&_BasePool.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BasePool *BasePoolTransactorSession) Activate() (*types.Transaction, error) {
	return _BasePool.Contract.Activate(&_BasePool.TransactOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BasePool *BasePoolTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BasePool *BasePoolSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.AddAdmin(&_BasePool.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BasePool *BasePoolTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.AddAdmin(&_BasePool.TransactOpts, addr)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BasePool *BasePoolTransactor) AddAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "addAuthorized", entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BasePool *BasePoolSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BasePool.Contract.AddAuthorized(&_BasePool.TransactOpts, entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BasePool *BasePoolTransactorSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BasePool.Contract.AddAuthorized(&_BasePool.TransactOpts, entry)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_BasePool *BasePoolTransactor) AddBallot(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "addBallot", b)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_BasePool *BasePoolSession) AddBallot(b common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.AddBallot(&_BasePool.TransactOpts, b)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_BasePool *BasePoolTransactorSession) AddBallot(b common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.AddBallot(&_BasePool.TransactOpts, b)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasePool *BasePoolTransactor) CastVote(opts *bind.TransactOpts, voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "castVote", voteId, vote, passphrase)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasePool *BasePoolSession) CastVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasePool.Contract.CastVote(&_BasePool.TransactOpts, voteId, vote, passphrase)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasePool *BasePoolTransactorSession) CastVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasePool.Contract.CastVote(&_BasePool.TransactOpts, voteId, vote, passphrase)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BasePool *BasePoolTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BasePool *BasePoolSession) Close() (*types.Transaction, error) {
	return _BasePool.Contract.Close(&_BasePool.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BasePool *BasePoolTransactorSession) Close() (*types.Transaction, error) {
	return _BasePool.Contract.Close(&_BasePool.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BasePool *BasePoolTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BasePool *BasePoolSession) Lock() (*types.Transaction, error) {
	return _BasePool.Contract.Lock(&_BasePool.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BasePool *BasePoolTransactorSession) Lock() (*types.Transaction, error) {
	return _BasePool.Contract.Lock(&_BasePool.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BasePool *BasePoolTransactor) RemoveAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "removeAdmin", addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BasePool *BasePoolSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.RemoveAdmin(&_BasePool.TransactOpts, addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BasePool *BasePoolTransactorSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.RemoveAdmin(&_BasePool.TransactOpts, addr)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BasePool *BasePoolTransactor) RemoveAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "removeAuthorized", entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BasePool *BasePoolSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BasePool.Contract.RemoveAuthorized(&_BasePool.TransactOpts, entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BasePool *BasePoolTransactorSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BasePool.Contract.RemoveAuthorized(&_BasePool.TransactOpts, entry)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_BasePool *BasePoolTransactor) RemoveBallot(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "removeBallot", b)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_BasePool *BasePoolSession) RemoveBallot(b common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.RemoveBallot(&_BasePool.TransactOpts, b)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_BasePool *BasePoolTransactorSession) RemoveBallot(b common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.RemoveBallot(&_BasePool.TransactOpts, b)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BasePool *BasePoolTransactor) RemoveSelf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "removeSelf")
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BasePool *BasePoolSession) RemoveSelf() (*types.Transaction, error) {
	return _BasePool.Contract.RemoveSelf(&_BasePool.TransactOpts)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BasePool *BasePoolTransactorSession) RemoveSelf() (*types.Transaction, error) {
	return _BasePool.Contract.RemoveSelf(&_BasePool.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasePool *BasePoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasePool *BasePoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.TransferOwnership(&_BasePool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasePool *BasePoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasePool.Contract.TransferOwnership(&_BasePool.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BasePool *BasePoolTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BasePool *BasePoolSession) Unlock() (*types.Transaction, error) {
	return _BasePool.Contract.Unlock(&_BasePool.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BasePool *BasePoolTransactorSession) Unlock() (*types.Transaction, error) {
	return _BasePool.Contract.Unlock(&_BasePool.TransactOpts)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasePool *BasePoolTransactor) UpdateVote(opts *bind.TransactOpts, voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasePool.contract.Transact(opts, "updateVote", voteId, vote, passphrase)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasePool *BasePoolSession) UpdateVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasePool.Contract.UpdateVote(&_BasePool.TransactOpts, voteId, vote, passphrase)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasePool *BasePoolTransactorSession) UpdateVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasePool.Contract.UpdateVote(&_BasePool.TransactOpts, voteId, vote, passphrase)
}
