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

// TieredBallotABI is the input ABI used to generate the binding from.
const TieredBallotABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"group\",\"type\":\"bytes32\"}],\"name\":\"groupPoolCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGroupCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"address\"}],\"name\":\"poolExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionPhase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"p\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteId\",\"type\":\"bytes32\"}],\"name\":\"castVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"bytes32\"}],\"name\":\"removeGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"election\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPoolGroupAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeSelf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getPoolVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"group\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGroupPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"address\"}],\"name\":\"getPoolIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkPools\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolVoterCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"metadataLocation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkElection\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"p\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"group\",\"type\":\"bytes32\"}],\"name\":\"addPoolToGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolGroupCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"bytes32\"}],\"name\":\"addGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"electionAddress\",\"type\":\"address\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"location\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"voteId\",\"type\":\"bytes32\"}],\"name\":\"BallotVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TieredBallot is an auto generated Go binding around an Ethereum contract.
type TieredBallot struct {
	TieredBallotCaller     // Read-only binding to the contract
	TieredBallotTransactor // Write-only binding to the contract
}

// TieredBallotCaller is an auto generated read-only Go binding around an Ethereum contract.
type TieredBallotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TieredBallotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TieredBallotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TieredBallotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TieredBallotSession struct {
	Contract     *TieredBallot     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TieredBallotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TieredBallotCallerSession struct {
	Contract *TieredBallotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TieredBallotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TieredBallotTransactorSession struct {
	Contract     *TieredBallotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TieredBallotRaw is an auto generated low-level Go binding around an Ethereum contract.
type TieredBallotRaw struct {
	Contract *TieredBallot // Generic contract binding to access the raw methods on
}

// TieredBallotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TieredBallotCallerRaw struct {
	Contract *TieredBallotCaller // Generic read-only contract binding to access the raw methods on
}

// TieredBallotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TieredBallotTransactorRaw struct {
	Contract *TieredBallotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTieredBallot creates a new instance of TieredBallot, bound to a specific deployed contract.
func NewTieredBallot(address common.Address, backend bind.ContractBackend) (*TieredBallot, error) {
	contract, err := bindTieredBallot(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TieredBallot{TieredBallotCaller: TieredBallotCaller{contract: contract}, TieredBallotTransactor: TieredBallotTransactor{contract: contract}}, nil
}

// NewTieredBallotCaller creates a new read-only instance of TieredBallot, bound to a specific deployed contract.
func NewTieredBallotCaller(address common.Address, caller bind.ContractCaller) (*TieredBallotCaller, error) {
	contract, err := bindTieredBallot(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TieredBallotCaller{contract: contract}, nil
}

// NewTieredBallotTransactor creates a new write-only instance of TieredBallot, bound to a specific deployed contract.
func NewTieredBallotTransactor(address common.Address, transactor bind.ContractTransactor) (*TieredBallotTransactor, error) {
	contract, err := bindTieredBallot(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TieredBallotTransactor{contract: contract}, nil
}

// bindTieredBallot binds a generic wrapper to an already deployed contract.
func bindTieredBallot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TieredBallotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TieredBallot *TieredBallotRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TieredBallot.Contract.TieredBallotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TieredBallot *TieredBallotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredBallot.Contract.TieredBallotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TieredBallot *TieredBallotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TieredBallot.Contract.TieredBallotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TieredBallot *TieredBallotCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TieredBallot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TieredBallot *TieredBallotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredBallot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TieredBallot *TieredBallotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TieredBallot.Contract.contract.Transact(opts, method, params...)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredBallot *TieredBallotCaller) CheckConfig(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "checkConfig")
	return *ret0, err
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredBallot *TieredBallotSession) CheckConfig() (bool, error) {
	return _TieredBallot.Contract.CheckConfig(&_TieredBallot.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredBallot *TieredBallotCallerSession) CheckConfig() (bool, error) {
	return _TieredBallot.Contract.CheckConfig(&_TieredBallot.CallOpts)
}

// CheckElection is a free data retrieval call binding the contract method 0xd1bdd4d3.
//
// Solidity: function checkElection() constant returns(bool)
func (_TieredBallot *TieredBallotCaller) CheckElection(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "checkElection")
	return *ret0, err
}

// CheckElection is a free data retrieval call binding the contract method 0xd1bdd4d3.
//
// Solidity: function checkElection() constant returns(bool)
func (_TieredBallot *TieredBallotSession) CheckElection() (bool, error) {
	return _TieredBallot.Contract.CheckElection(&_TieredBallot.CallOpts)
}

// CheckElection is a free data retrieval call binding the contract method 0xd1bdd4d3.
//
// Solidity: function checkElection() constant returns(bool)
func (_TieredBallot *TieredBallotCallerSession) CheckElection() (bool, error) {
	return _TieredBallot.Contract.CheckElection(&_TieredBallot.CallOpts)
}

// CheckPools is a free data retrieval call binding the contract method 0x8e3d89cb.
//
// Solidity: function checkPools() constant returns(bool)
func (_TieredBallot *TieredBallotCaller) CheckPools(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "checkPools")
	return *ret0, err
}

// CheckPools is a free data retrieval call binding the contract method 0x8e3d89cb.
//
// Solidity: function checkPools() constant returns(bool)
func (_TieredBallot *TieredBallotSession) CheckPools() (bool, error) {
	return _TieredBallot.Contract.CheckPools(&_TieredBallot.CallOpts)
}

// CheckPools is a free data retrieval call binding the contract method 0x8e3d89cb.
//
// Solidity: function checkPools() constant returns(bool)
func (_TieredBallot *TieredBallotCallerSession) CheckPools() (bool, error) {
	return _TieredBallot.Contract.CheckPools(&_TieredBallot.CallOpts)
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_TieredBallot *TieredBallotCaller) Election(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "election")
	return *ret0, err
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_TieredBallot *TieredBallotSession) Election() (common.Address, error) {
	return _TieredBallot.Contract.Election(&_TieredBallot.CallOpts)
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_TieredBallot *TieredBallotCallerSession) Election() (common.Address, error) {
	return _TieredBallot.Contract.Election(&_TieredBallot.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredBallot *TieredBallotCaller) ElectionPhase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "electionPhase")
	return *ret0, err
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredBallot *TieredBallotSession) ElectionPhase() (uint8, error) {
	return _TieredBallot.Contract.ElectionPhase(&_TieredBallot.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredBallot *TieredBallotCallerSession) ElectionPhase() (uint8, error) {
	return _TieredBallot.Contract.ElectionPhase(&_TieredBallot.CallOpts)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(index uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotCaller) GetGroup(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getGroup", index)
	return *ret0, err
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(index uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotSession) GetGroup(index *big.Int) ([32]byte, error) {
	return _TieredBallot.Contract.GetGroup(&_TieredBallot.CallOpts, index)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(index uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotCallerSession) GetGroup(index *big.Int) ([32]byte, error) {
	return _TieredBallot.Contract.GetGroup(&_TieredBallot.CallOpts, index)
}

// GetGroupCount is a free data retrieval call binding the contract method 0x06545a93.
//
// Solidity: function getGroupCount() constant returns(uint256)
func (_TieredBallot *TieredBallotCaller) GetGroupCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getGroupCount")
	return *ret0, err
}

// GetGroupCount is a free data retrieval call binding the contract method 0x06545a93.
//
// Solidity: function getGroupCount() constant returns(uint256)
func (_TieredBallot *TieredBallotSession) GetGroupCount() (*big.Int, error) {
	return _TieredBallot.Contract.GetGroupCount(&_TieredBallot.CallOpts)
}

// GetGroupCount is a free data retrieval call binding the contract method 0x06545a93.
//
// Solidity: function getGroupCount() constant returns(uint256)
func (_TieredBallot *TieredBallotCallerSession) GetGroupCount() (*big.Int, error) {
	return _TieredBallot.Contract.GetGroupCount(&_TieredBallot.CallOpts)
}

// GetGroupPool is a free data retrieval call binding the contract method 0x7a4e367c.
//
// Solidity: function getGroupPool(group bytes32, index uint256) constant returns(address)
func (_TieredBallot *TieredBallotCaller) GetGroupPool(opts *bind.CallOpts, group [32]byte, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getGroupPool", group, index)
	return *ret0, err
}

// GetGroupPool is a free data retrieval call binding the contract method 0x7a4e367c.
//
// Solidity: function getGroupPool(group bytes32, index uint256) constant returns(address)
func (_TieredBallot *TieredBallotSession) GetGroupPool(group [32]byte, index *big.Int) (common.Address, error) {
	return _TieredBallot.Contract.GetGroupPool(&_TieredBallot.CallOpts, group, index)
}

// GetGroupPool is a free data retrieval call binding the contract method 0x7a4e367c.
//
// Solidity: function getGroupPool(group bytes32, index uint256) constant returns(address)
func (_TieredBallot *TieredBallotCallerSession) GetGroupPool(group [32]byte, index *big.Int) (common.Address, error) {
	return _TieredBallot.Contract.GetGroupPool(&_TieredBallot.CallOpts, group, index)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(index uint256) constant returns(address)
func (_TieredBallot *TieredBallotCaller) GetPool(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getPool", index)
	return *ret0, err
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(index uint256) constant returns(address)
func (_TieredBallot *TieredBallotSession) GetPool(index *big.Int) (common.Address, error) {
	return _TieredBallot.Contract.GetPool(&_TieredBallot.CallOpts, index)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(index uint256) constant returns(address)
func (_TieredBallot *TieredBallotCallerSession) GetPool(index *big.Int) (common.Address, error) {
	return _TieredBallot.Contract.GetPool(&_TieredBallot.CallOpts, index)
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() constant returns(uint256)
func (_TieredBallot *TieredBallotCaller) GetPoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getPoolCount")
	return *ret0, err
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() constant returns(uint256)
func (_TieredBallot *TieredBallotSession) GetPoolCount() (*big.Int, error) {
	return _TieredBallot.Contract.GetPoolCount(&_TieredBallot.CallOpts)
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() constant returns(uint256)
func (_TieredBallot *TieredBallotCallerSession) GetPoolCount() (*big.Int, error) {
	return _TieredBallot.Contract.GetPoolCount(&_TieredBallot.CallOpts)
}

// GetPoolGroupAt is a free data retrieval call binding the contract method 0x597c92f7.
//
// Solidity: function getPoolGroupAt(pool address, index uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotCaller) GetPoolGroupAt(opts *bind.CallOpts, pool common.Address, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getPoolGroupAt", pool, index)
	return *ret0, err
}

// GetPoolGroupAt is a free data retrieval call binding the contract method 0x597c92f7.
//
// Solidity: function getPoolGroupAt(pool address, index uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotSession) GetPoolGroupAt(pool common.Address, index *big.Int) ([32]byte, error) {
	return _TieredBallot.Contract.GetPoolGroupAt(&_TieredBallot.CallOpts, pool, index)
}

// GetPoolGroupAt is a free data retrieval call binding the contract method 0x597c92f7.
//
// Solidity: function getPoolGroupAt(pool address, index uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotCallerSession) GetPoolGroupAt(pool common.Address, index *big.Int) ([32]byte, error) {
	return _TieredBallot.Contract.GetPoolGroupAt(&_TieredBallot.CallOpts, pool, index)
}

// GetPoolGroupCount is a free data retrieval call binding the contract method 0xe5de025a.
//
// Solidity: function getPoolGroupCount(pool address) constant returns(uint256)
func (_TieredBallot *TieredBallotCaller) GetPoolGroupCount(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getPoolGroupCount", pool)
	return *ret0, err
}

// GetPoolGroupCount is a free data retrieval call binding the contract method 0xe5de025a.
//
// Solidity: function getPoolGroupCount(pool address) constant returns(uint256)
func (_TieredBallot *TieredBallotSession) GetPoolGroupCount(pool common.Address) (*big.Int, error) {
	return _TieredBallot.Contract.GetPoolGroupCount(&_TieredBallot.CallOpts, pool)
}

// GetPoolGroupCount is a free data retrieval call binding the contract method 0xe5de025a.
//
// Solidity: function getPoolGroupCount(pool address) constant returns(uint256)
func (_TieredBallot *TieredBallotCallerSession) GetPoolGroupCount(pool common.Address) (*big.Int, error) {
	return _TieredBallot.Contract.GetPoolGroupCount(&_TieredBallot.CallOpts, pool)
}

// GetPoolIndex is a free data retrieval call binding the contract method 0x87063850.
//
// Solidity: function getPoolIndex(p address) constant returns(uint256)
func (_TieredBallot *TieredBallotCaller) GetPoolIndex(opts *bind.CallOpts, p common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getPoolIndex", p)
	return *ret0, err
}

// GetPoolIndex is a free data retrieval call binding the contract method 0x87063850.
//
// Solidity: function getPoolIndex(p address) constant returns(uint256)
func (_TieredBallot *TieredBallotSession) GetPoolIndex(p common.Address) (*big.Int, error) {
	return _TieredBallot.Contract.GetPoolIndex(&_TieredBallot.CallOpts, p)
}

// GetPoolIndex is a free data retrieval call binding the contract method 0x87063850.
//
// Solidity: function getPoolIndex(p address) constant returns(uint256)
func (_TieredBallot *TieredBallotCallerSession) GetPoolIndex(p common.Address) (*big.Int, error) {
	return _TieredBallot.Contract.GetPoolIndex(&_TieredBallot.CallOpts, p)
}

// GetPoolVoter is a free data retrieval call binding the contract method 0x68e3e6b6.
//
// Solidity: function getPoolVoter(pool address, i uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotCaller) GetPoolVoter(opts *bind.CallOpts, pool common.Address, i *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getPoolVoter", pool, i)
	return *ret0, err
}

// GetPoolVoter is a free data retrieval call binding the contract method 0x68e3e6b6.
//
// Solidity: function getPoolVoter(pool address, i uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotSession) GetPoolVoter(pool common.Address, i *big.Int) ([32]byte, error) {
	return _TieredBallot.Contract.GetPoolVoter(&_TieredBallot.CallOpts, pool, i)
}

// GetPoolVoter is a free data retrieval call binding the contract method 0x68e3e6b6.
//
// Solidity: function getPoolVoter(pool address, i uint256) constant returns(bytes32)
func (_TieredBallot *TieredBallotCallerSession) GetPoolVoter(pool common.Address, i *big.Int) ([32]byte, error) {
	return _TieredBallot.Contract.GetPoolVoter(&_TieredBallot.CallOpts, pool, i)
}

// GetPoolVoterCount is a free data retrieval call binding the contract method 0x978daca7.
//
// Solidity: function getPoolVoterCount(pool address) constant returns(uint256)
func (_TieredBallot *TieredBallotCaller) GetPoolVoterCount(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "getPoolVoterCount", pool)
	return *ret0, err
}

// GetPoolVoterCount is a free data retrieval call binding the contract method 0x978daca7.
//
// Solidity: function getPoolVoterCount(pool address) constant returns(uint256)
func (_TieredBallot *TieredBallotSession) GetPoolVoterCount(pool common.Address) (*big.Int, error) {
	return _TieredBallot.Contract.GetPoolVoterCount(&_TieredBallot.CallOpts, pool)
}

// GetPoolVoterCount is a free data retrieval call binding the contract method 0x978daca7.
//
// Solidity: function getPoolVoterCount(pool address) constant returns(uint256)
func (_TieredBallot *TieredBallotCallerSession) GetPoolVoterCount(pool common.Address) (*big.Int, error) {
	return _TieredBallot.Contract.GetPoolVoterCount(&_TieredBallot.CallOpts, pool)
}

// GroupPoolCount is a free data retrieval call binding the contract method 0x027defc1.
//
// Solidity: function groupPoolCount(group bytes32) constant returns(uint256)
func (_TieredBallot *TieredBallotCaller) GroupPoolCount(opts *bind.CallOpts, group [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "groupPoolCount", group)
	return *ret0, err
}

// GroupPoolCount is a free data retrieval call binding the contract method 0x027defc1.
//
// Solidity: function groupPoolCount(group bytes32) constant returns(uint256)
func (_TieredBallot *TieredBallotSession) GroupPoolCount(group [32]byte) (*big.Int, error) {
	return _TieredBallot.Contract.GroupPoolCount(&_TieredBallot.CallOpts, group)
}

// GroupPoolCount is a free data retrieval call binding the contract method 0x027defc1.
//
// Solidity: function groupPoolCount(group bytes32) constant returns(uint256)
func (_TieredBallot *TieredBallotCallerSession) GroupPoolCount(group [32]byte) (*big.Int, error) {
	return _TieredBallot.Contract.GroupPoolCount(&_TieredBallot.CallOpts, group)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredBallot *TieredBallotCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredBallot *TieredBallotSession) IsAdmin(addr common.Address) (bool, error) {
	return _TieredBallot.Contract.IsAdmin(&_TieredBallot.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredBallot *TieredBallotCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _TieredBallot.Contract.IsAdmin(&_TieredBallot.CallOpts, addr)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredBallot *TieredBallotCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredBallot *TieredBallotSession) IsClosed() (bool, error) {
	return _TieredBallot.Contract.IsClosed(&_TieredBallot.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredBallot *TieredBallotCallerSession) IsClosed() (bool, error) {
	return _TieredBallot.Contract.IsClosed(&_TieredBallot.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredBallot *TieredBallotCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredBallot *TieredBallotSession) IsLocked() (bool, error) {
	return _TieredBallot.Contract.IsLocked(&_TieredBallot.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredBallot *TieredBallotCallerSession) IsLocked() (bool, error) {
	return _TieredBallot.Contract.IsLocked(&_TieredBallot.CallOpts)
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_TieredBallot *TieredBallotCaller) MetadataLocation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "metadataLocation")
	return *ret0, err
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_TieredBallot *TieredBallotSession) MetadataLocation() (string, error) {
	return _TieredBallot.Contract.MetadataLocation(&_TieredBallot.CallOpts)
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_TieredBallot *TieredBallotCallerSession) MetadataLocation() (string, error) {
	return _TieredBallot.Contract.MetadataLocation(&_TieredBallot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredBallot *TieredBallotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredBallot *TieredBallotSession) Owner() (common.Address, error) {
	return _TieredBallot.Contract.Owner(&_TieredBallot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredBallot *TieredBallotCallerSession) Owner() (common.Address, error) {
	return _TieredBallot.Contract.Owner(&_TieredBallot.CallOpts)
}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(p address) constant returns(bool)
func (_TieredBallot *TieredBallotCaller) PoolExists(opts *bind.CallOpts, p common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredBallot.contract.Call(opts, out, "poolExists", p)
	return *ret0, err
}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(p address) constant returns(bool)
func (_TieredBallot *TieredBallotSession) PoolExists(p common.Address) (bool, error) {
	return _TieredBallot.Contract.PoolExists(&_TieredBallot.CallOpts, p)
}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(p address) constant returns(bool)
func (_TieredBallot *TieredBallotCallerSession) PoolExists(p common.Address) (bool, error) {
	return _TieredBallot.Contract.PoolExists(&_TieredBallot.CallOpts, p)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredBallot *TieredBallotTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredBallot *TieredBallotSession) Abort() (*types.Transaction, error) {
	return _TieredBallot.Contract.Abort(&_TieredBallot.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredBallot *TieredBallotTransactorSession) Abort() (*types.Transaction, error) {
	return _TieredBallot.Contract.Abort(&_TieredBallot.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredBallot *TieredBallotTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredBallot *TieredBallotSession) Activate() (*types.Transaction, error) {
	return _TieredBallot.Contract.Activate(&_TieredBallot.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredBallot *TieredBallotTransactorSession) Activate() (*types.Transaction, error) {
	return _TieredBallot.Contract.Activate(&_TieredBallot.TransactOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredBallot *TieredBallotTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredBallot *TieredBallotSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.AddAdmin(&_TieredBallot.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredBallot *TieredBallotTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.AddAdmin(&_TieredBallot.TransactOpts, addr)
}

// AddGroup is a paid mutator transaction binding the contract method 0xfe2b699a.
//
// Solidity: function addGroup(group bytes32) returns()
func (_TieredBallot *TieredBallotTransactor) AddGroup(opts *bind.TransactOpts, group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "addGroup", group)
}

// AddGroup is a paid mutator transaction binding the contract method 0xfe2b699a.
//
// Solidity: function addGroup(group bytes32) returns()
func (_TieredBallot *TieredBallotSession) AddGroup(group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.Contract.AddGroup(&_TieredBallot.TransactOpts, group)
}

// AddGroup is a paid mutator transaction binding the contract method 0xfe2b699a.
//
// Solidity: function addGroup(group bytes32) returns()
func (_TieredBallot *TieredBallotTransactorSession) AddGroup(group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.Contract.AddGroup(&_TieredBallot.TransactOpts, group)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(p address) returns()
func (_TieredBallot *TieredBallotTransactor) AddPool(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "addPool", p)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(p address) returns()
func (_TieredBallot *TieredBallotSession) AddPool(p common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.AddPool(&_TieredBallot.TransactOpts, p)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(p address) returns()
func (_TieredBallot *TieredBallotTransactorSession) AddPool(p common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.AddPool(&_TieredBallot.TransactOpts, p)
}

// AddPoolToGroup is a paid mutator transaction binding the contract method 0xe5d3b258.
//
// Solidity: function addPoolToGroup(pool address, group bytes32) returns()
func (_TieredBallot *TieredBallotTransactor) AddPoolToGroup(opts *bind.TransactOpts, pool common.Address, group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "addPoolToGroup", pool, group)
}

// AddPoolToGroup is a paid mutator transaction binding the contract method 0xe5d3b258.
//
// Solidity: function addPoolToGroup(pool address, group bytes32) returns()
func (_TieredBallot *TieredBallotSession) AddPoolToGroup(pool common.Address, group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.Contract.AddPoolToGroup(&_TieredBallot.TransactOpts, pool, group)
}

// AddPoolToGroup is a paid mutator transaction binding the contract method 0xe5d3b258.
//
// Solidity: function addPoolToGroup(pool address, group bytes32) returns()
func (_TieredBallot *TieredBallotTransactorSession) AddPoolToGroup(pool common.Address, group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.Contract.AddPoolToGroup(&_TieredBallot.TransactOpts, pool, group)
}

// CastVote is a paid mutator transaction binding the contract method 0x3efafb6e.
//
// Solidity: function castVote(voteId bytes32) returns()
func (_TieredBallot *TieredBallotTransactor) CastVote(opts *bind.TransactOpts, voteId [32]byte) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "castVote", voteId)
}

// CastVote is a paid mutator transaction binding the contract method 0x3efafb6e.
//
// Solidity: function castVote(voteId bytes32) returns()
func (_TieredBallot *TieredBallotSession) CastVote(voteId [32]byte) (*types.Transaction, error) {
	return _TieredBallot.Contract.CastVote(&_TieredBallot.TransactOpts, voteId)
}

// CastVote is a paid mutator transaction binding the contract method 0x3efafb6e.
//
// Solidity: function castVote(voteId bytes32) returns()
func (_TieredBallot *TieredBallotTransactorSession) CastVote(voteId [32]byte) (*types.Transaction, error) {
	return _TieredBallot.Contract.CastVote(&_TieredBallot.TransactOpts, voteId)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredBallot *TieredBallotTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredBallot *TieredBallotSession) Close() (*types.Transaction, error) {
	return _TieredBallot.Contract.Close(&_TieredBallot.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredBallot *TieredBallotTransactorSession) Close() (*types.Transaction, error) {
	return _TieredBallot.Contract.Close(&_TieredBallot.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredBallot *TieredBallotTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredBallot *TieredBallotSession) Lock() (*types.Transaction, error) {
	return _TieredBallot.Contract.Lock(&_TieredBallot.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredBallot *TieredBallotTransactorSession) Lock() (*types.Transaction, error) {
	return _TieredBallot.Contract.Lock(&_TieredBallot.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredBallot *TieredBallotTransactor) RemoveAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "removeAdmin", addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredBallot *TieredBallotSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.RemoveAdmin(&_TieredBallot.TransactOpts, addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredBallot *TieredBallotTransactorSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.RemoveAdmin(&_TieredBallot.TransactOpts, addr)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x43273213.
//
// Solidity: function removeGroup(group bytes32) returns()
func (_TieredBallot *TieredBallotTransactor) RemoveGroup(opts *bind.TransactOpts, group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "removeGroup", group)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x43273213.
//
// Solidity: function removeGroup(group bytes32) returns()
func (_TieredBallot *TieredBallotSession) RemoveGroup(group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.Contract.RemoveGroup(&_TieredBallot.TransactOpts, group)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x43273213.
//
// Solidity: function removeGroup(group bytes32) returns()
func (_TieredBallot *TieredBallotTransactorSession) RemoveGroup(group [32]byte) (*types.Transaction, error) {
	return _TieredBallot.Contract.RemoveGroup(&_TieredBallot.TransactOpts, group)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(p address) returns()
func (_TieredBallot *TieredBallotTransactor) RemovePool(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "removePool", p)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(p address) returns()
func (_TieredBallot *TieredBallotSession) RemovePool(p common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.RemovePool(&_TieredBallot.TransactOpts, p)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(p address) returns()
func (_TieredBallot *TieredBallotTransactorSession) RemovePool(p common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.RemovePool(&_TieredBallot.TransactOpts, p)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredBallot *TieredBallotTransactor) RemoveSelf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "removeSelf")
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredBallot *TieredBallotSession) RemoveSelf() (*types.Transaction, error) {
	return _TieredBallot.Contract.RemoveSelf(&_TieredBallot.TransactOpts)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredBallot *TieredBallotTransactorSession) RemoveSelf() (*types.Transaction, error) {
	return _TieredBallot.Contract.RemoveSelf(&_TieredBallot.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredBallot *TieredBallotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredBallot *TieredBallotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.TransferOwnership(&_TieredBallot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredBallot *TieredBallotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TieredBallot.Contract.TransferOwnership(&_TieredBallot.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredBallot *TieredBallotTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredBallot.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredBallot *TieredBallotSession) Unlock() (*types.Transaction, error) {
	return _TieredBallot.Contract.Unlock(&_TieredBallot.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredBallot *TieredBallotTransactorSession) Unlock() (*types.Transaction, error) {
	return _TieredBallot.Contract.Unlock(&_TieredBallot.TransactOpts)
}
