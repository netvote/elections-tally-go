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

// TieredElectionABI is the input ABI used to generate the binding from.
const TieredElectionABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"ballotExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"address\"}],\"name\":\"poolExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionPhase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"p\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"addAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionType\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"privateKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeSelf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"removeAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"address\"}],\"name\":\"getPoolIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBallotCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowVoteUpdates\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"removeBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"p\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"getBallotIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deductVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"setPrivateKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"addBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"hashedUserId\",\"type\":\"bytes32\"},{\"name\":\"allowanceAddress\",\"type\":\"address\"},{\"name\":\"acct\",\"type\":\"address\"},{\"name\":\"allowUpdates\",\"type\":\"bool\"},{\"name\":\"revealer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"KeyReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TieredElection is an auto generated Go binding around an Ethereum contract.
type TieredElection struct {
	TieredElectionCaller     // Read-only binding to the contract
	TieredElectionTransactor // Write-only binding to the contract
}

// TieredElectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TieredElectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TieredElectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TieredElectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TieredElectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TieredElectionSession struct {
	Contract     *TieredElection   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TieredElectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TieredElectionCallerSession struct {
	Contract *TieredElectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TieredElectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TieredElectionTransactorSession struct {
	Contract     *TieredElectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TieredElectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TieredElectionRaw struct {
	Contract *TieredElection // Generic contract binding to access the raw methods on
}

// TieredElectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TieredElectionCallerRaw struct {
	Contract *TieredElectionCaller // Generic read-only contract binding to access the raw methods on
}

// TieredElectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TieredElectionTransactorRaw struct {
	Contract *TieredElectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTieredElection creates a new instance of TieredElection, bound to a specific deployed contract.
func NewTieredElection(address common.Address, backend bind.ContractBackend) (*TieredElection, error) {
	contract, err := bindTieredElection(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TieredElection{TieredElectionCaller: TieredElectionCaller{contract: contract}, TieredElectionTransactor: TieredElectionTransactor{contract: contract}}, nil
}

// NewTieredElectionCaller creates a new read-only instance of TieredElection, bound to a specific deployed contract.
func NewTieredElectionCaller(address common.Address, caller bind.ContractCaller) (*TieredElectionCaller, error) {
	contract, err := bindTieredElection(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TieredElectionCaller{contract: contract}, nil
}

// NewTieredElectionTransactor creates a new write-only instance of TieredElection, bound to a specific deployed contract.
func NewTieredElectionTransactor(address common.Address, transactor bind.ContractTransactor) (*TieredElectionTransactor, error) {
	contract, err := bindTieredElection(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TieredElectionTransactor{contract: contract}, nil
}

// bindTieredElection binds a generic wrapper to an already deployed contract.
func bindTieredElection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TieredElectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TieredElection *TieredElectionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TieredElection.Contract.TieredElectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TieredElection *TieredElectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.Contract.TieredElectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TieredElection *TieredElectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TieredElection.Contract.TieredElectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TieredElection *TieredElectionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TieredElection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TieredElection *TieredElectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TieredElection *TieredElectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TieredElection.Contract.contract.Transact(opts, method, params...)
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_TieredElection *TieredElectionCaller) AllowVoteUpdates(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "allowVoteUpdates")
	return *ret0, err
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_TieredElection *TieredElectionSession) AllowVoteUpdates() (bool, error) {
	return _TieredElection.Contract.AllowVoteUpdates(&_TieredElection.CallOpts)
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_TieredElection *TieredElectionCallerSession) AllowVoteUpdates() (bool, error) {
	return _TieredElection.Contract.AllowVoteUpdates(&_TieredElection.CallOpts)
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_TieredElection *TieredElectionCaller) BallotExists(opts *bind.CallOpts, b common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "ballotExists", b)
	return *ret0, err
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_TieredElection *TieredElectionSession) BallotExists(b common.Address) (bool, error) {
	return _TieredElection.Contract.BallotExists(&_TieredElection.CallOpts, b)
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_TieredElection *TieredElectionCallerSession) BallotExists(b common.Address) (bool, error) {
	return _TieredElection.Contract.BallotExists(&_TieredElection.CallOpts, b)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredElection *TieredElectionCaller) CheckConfig(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "checkConfig")
	return *ret0, err
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredElection *TieredElectionSession) CheckConfig() (bool, error) {
	return _TieredElection.Contract.CheckConfig(&_TieredElection.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredElection *TieredElectionCallerSession) CheckConfig() (bool, error) {
	return _TieredElection.Contract.CheckConfig(&_TieredElection.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredElection *TieredElectionCaller) ElectionPhase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "electionPhase")
	return *ret0, err
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredElection *TieredElectionSession) ElectionPhase() (uint8, error) {
	return _TieredElection.Contract.ElectionPhase(&_TieredElection.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredElection *TieredElectionCallerSession) ElectionPhase() (uint8, error) {
	return _TieredElection.Contract.ElectionPhase(&_TieredElection.CallOpts)
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_TieredElection *TieredElectionCaller) ElectionType(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "electionType")
	return *ret0, err
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_TieredElection *TieredElectionSession) ElectionType() (string, error) {
	return _TieredElection.Contract.ElectionType(&_TieredElection.CallOpts)
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_TieredElection *TieredElectionCallerSession) ElectionType() (string, error) {
	return _TieredElection.Contract.ElectionType(&_TieredElection.CallOpts)
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_TieredElection *TieredElectionCaller) GetBallot(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "getBallot", index)
	return *ret0, err
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_TieredElection *TieredElectionSession) GetBallot(index *big.Int) (common.Address, error) {
	return _TieredElection.Contract.GetBallot(&_TieredElection.CallOpts, index)
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_TieredElection *TieredElectionCallerSession) GetBallot(index *big.Int) (common.Address, error) {
	return _TieredElection.Contract.GetBallot(&_TieredElection.CallOpts, index)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_TieredElection *TieredElectionCaller) GetBallotCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "getBallotCount")
	return *ret0, err
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_TieredElection *TieredElectionSession) GetBallotCount() (*big.Int, error) {
	return _TieredElection.Contract.GetBallotCount(&_TieredElection.CallOpts)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_TieredElection *TieredElectionCallerSession) GetBallotCount() (*big.Int, error) {
	return _TieredElection.Contract.GetBallotCount(&_TieredElection.CallOpts)
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_TieredElection *TieredElectionCaller) GetBallotIndex(opts *bind.CallOpts, b common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "getBallotIndex", b)
	return *ret0, err
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_TieredElection *TieredElectionSession) GetBallotIndex(b common.Address) (*big.Int, error) {
	return _TieredElection.Contract.GetBallotIndex(&_TieredElection.CallOpts, b)
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_TieredElection *TieredElectionCallerSession) GetBallotIndex(b common.Address) (*big.Int, error) {
	return _TieredElection.Contract.GetBallotIndex(&_TieredElection.CallOpts, b)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(index uint256) constant returns(address)
func (_TieredElection *TieredElectionCaller) GetPool(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "getPool", index)
	return *ret0, err
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(index uint256) constant returns(address)
func (_TieredElection *TieredElectionSession) GetPool(index *big.Int) (common.Address, error) {
	return _TieredElection.Contract.GetPool(&_TieredElection.CallOpts, index)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(index uint256) constant returns(address)
func (_TieredElection *TieredElectionCallerSession) GetPool(index *big.Int) (common.Address, error) {
	return _TieredElection.Contract.GetPool(&_TieredElection.CallOpts, index)
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() constant returns(uint256)
func (_TieredElection *TieredElectionCaller) GetPoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "getPoolCount")
	return *ret0, err
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() constant returns(uint256)
func (_TieredElection *TieredElectionSession) GetPoolCount() (*big.Int, error) {
	return _TieredElection.Contract.GetPoolCount(&_TieredElection.CallOpts)
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() constant returns(uint256)
func (_TieredElection *TieredElectionCallerSession) GetPoolCount() (*big.Int, error) {
	return _TieredElection.Contract.GetPoolCount(&_TieredElection.CallOpts)
}

// GetPoolIndex is a free data retrieval call binding the contract method 0x87063850.
//
// Solidity: function getPoolIndex(p address) constant returns(uint256)
func (_TieredElection *TieredElectionCaller) GetPoolIndex(opts *bind.CallOpts, p common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "getPoolIndex", p)
	return *ret0, err
}

// GetPoolIndex is a free data retrieval call binding the contract method 0x87063850.
//
// Solidity: function getPoolIndex(p address) constant returns(uint256)
func (_TieredElection *TieredElectionSession) GetPoolIndex(p common.Address) (*big.Int, error) {
	return _TieredElection.Contract.GetPoolIndex(&_TieredElection.CallOpts, p)
}

// GetPoolIndex is a free data retrieval call binding the contract method 0x87063850.
//
// Solidity: function getPoolIndex(p address) constant returns(uint256)
func (_TieredElection *TieredElectionCallerSession) GetPoolIndex(p common.Address) (*big.Int, error) {
	return _TieredElection.Contract.GetPoolIndex(&_TieredElection.CallOpts, p)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredElection *TieredElectionCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredElection *TieredElectionSession) IsAdmin(addr common.Address) (bool, error) {
	return _TieredElection.Contract.IsAdmin(&_TieredElection.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredElection *TieredElectionCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _TieredElection.Contract.IsAdmin(&_TieredElection.CallOpts, addr)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_TieredElection *TieredElectionCaller) IsAuthorized(opts *bind.CallOpts, entry [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "isAuthorized", entry)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_TieredElection *TieredElectionSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _TieredElection.Contract.IsAuthorized(&_TieredElection.CallOpts, entry)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_TieredElection *TieredElectionCallerSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _TieredElection.Contract.IsAuthorized(&_TieredElection.CallOpts, entry)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredElection *TieredElectionCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredElection *TieredElectionSession) IsClosed() (bool, error) {
	return _TieredElection.Contract.IsClosed(&_TieredElection.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredElection *TieredElectionCallerSession) IsClosed() (bool, error) {
	return _TieredElection.Contract.IsClosed(&_TieredElection.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredElection *TieredElectionCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredElection *TieredElectionSession) IsLocked() (bool, error) {
	return _TieredElection.Contract.IsLocked(&_TieredElection.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredElection *TieredElectionCallerSession) IsLocked() (bool, error) {
	return _TieredElection.Contract.IsLocked(&_TieredElection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredElection *TieredElectionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredElection *TieredElectionSession) Owner() (common.Address, error) {
	return _TieredElection.Contract.Owner(&_TieredElection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredElection *TieredElectionCallerSession) Owner() (common.Address, error) {
	return _TieredElection.Contract.Owner(&_TieredElection.CallOpts)
}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(p address) constant returns(bool)
func (_TieredElection *TieredElectionCaller) PoolExists(opts *bind.CallOpts, p common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "poolExists", p)
	return *ret0, err
}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(p address) constant returns(bool)
func (_TieredElection *TieredElectionSession) PoolExists(p common.Address) (bool, error) {
	return _TieredElection.Contract.PoolExists(&_TieredElection.CallOpts, p)
}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(p address) constant returns(bool)
func (_TieredElection *TieredElectionCallerSession) PoolExists(p common.Address) (bool, error) {
	return _TieredElection.Contract.PoolExists(&_TieredElection.CallOpts, p)
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_TieredElection *TieredElectionCaller) PrivateKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "privateKey")
	return *ret0, err
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_TieredElection *TieredElectionSession) PrivateKey() (string, error) {
	return _TieredElection.Contract.PrivateKey(&_TieredElection.CallOpts)
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_TieredElection *TieredElectionCallerSession) PrivateKey() (string, error) {
	return _TieredElection.Contract.PrivateKey(&_TieredElection.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_TieredElection *TieredElectionCaller) PublicKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TieredElection.contract.Call(opts, out, "publicKey")
	return *ret0, err
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_TieredElection *TieredElectionSession) PublicKey() (string, error) {
	return _TieredElection.Contract.PublicKey(&_TieredElection.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_TieredElection *TieredElectionCallerSession) PublicKey() (string, error) {
	return _TieredElection.Contract.PublicKey(&_TieredElection.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredElection *TieredElectionTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredElection *TieredElectionSession) Abort() (*types.Transaction, error) {
	return _TieredElection.Contract.Abort(&_TieredElection.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredElection *TieredElectionTransactorSession) Abort() (*types.Transaction, error) {
	return _TieredElection.Contract.Abort(&_TieredElection.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredElection *TieredElectionTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredElection *TieredElectionSession) Activate() (*types.Transaction, error) {
	return _TieredElection.Contract.Activate(&_TieredElection.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredElection *TieredElectionTransactorSession) Activate() (*types.Transaction, error) {
	return _TieredElection.Contract.Activate(&_TieredElection.TransactOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredElection *TieredElectionTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredElection *TieredElectionSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.AddAdmin(&_TieredElection.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredElection *TieredElectionTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.AddAdmin(&_TieredElection.TransactOpts, addr)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_TieredElection *TieredElectionTransactor) AddAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "addAuthorized", entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_TieredElection *TieredElectionSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _TieredElection.Contract.AddAuthorized(&_TieredElection.TransactOpts, entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_TieredElection *TieredElectionTransactorSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _TieredElection.Contract.AddAuthorized(&_TieredElection.TransactOpts, entry)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_TieredElection *TieredElectionTransactor) AddBallot(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "addBallot", b)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_TieredElection *TieredElectionSession) AddBallot(b common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.AddBallot(&_TieredElection.TransactOpts, b)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_TieredElection *TieredElectionTransactorSession) AddBallot(b common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.AddBallot(&_TieredElection.TransactOpts, b)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(p address) returns()
func (_TieredElection *TieredElectionTransactor) AddPool(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "addPool", p)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(p address) returns()
func (_TieredElection *TieredElectionSession) AddPool(p common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.AddPool(&_TieredElection.TransactOpts, p)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(p address) returns()
func (_TieredElection *TieredElectionTransactorSession) AddPool(p common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.AddPool(&_TieredElection.TransactOpts, p)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredElection *TieredElectionTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredElection *TieredElectionSession) Close() (*types.Transaction, error) {
	return _TieredElection.Contract.Close(&_TieredElection.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredElection *TieredElectionTransactorSession) Close() (*types.Transaction, error) {
	return _TieredElection.Contract.Close(&_TieredElection.TransactOpts)
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_TieredElection *TieredElectionTransactor) DeductVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "deductVote")
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_TieredElection *TieredElectionSession) DeductVote() (*types.Transaction, error) {
	return _TieredElection.Contract.DeductVote(&_TieredElection.TransactOpts)
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_TieredElection *TieredElectionTransactorSession) DeductVote() (*types.Transaction, error) {
	return _TieredElection.Contract.DeductVote(&_TieredElection.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredElection *TieredElectionTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredElection *TieredElectionSession) Lock() (*types.Transaction, error) {
	return _TieredElection.Contract.Lock(&_TieredElection.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredElection *TieredElectionTransactorSession) Lock() (*types.Transaction, error) {
	return _TieredElection.Contract.Lock(&_TieredElection.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredElection *TieredElectionTransactor) RemoveAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "removeAdmin", addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredElection *TieredElectionSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.RemoveAdmin(&_TieredElection.TransactOpts, addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredElection *TieredElectionTransactorSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.RemoveAdmin(&_TieredElection.TransactOpts, addr)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_TieredElection *TieredElectionTransactor) RemoveAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "removeAuthorized", entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_TieredElection *TieredElectionSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _TieredElection.Contract.RemoveAuthorized(&_TieredElection.TransactOpts, entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_TieredElection *TieredElectionTransactorSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _TieredElection.Contract.RemoveAuthorized(&_TieredElection.TransactOpts, entry)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_TieredElection *TieredElectionTransactor) RemoveBallot(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "removeBallot", b)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_TieredElection *TieredElectionSession) RemoveBallot(b common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.RemoveBallot(&_TieredElection.TransactOpts, b)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_TieredElection *TieredElectionTransactorSession) RemoveBallot(b common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.RemoveBallot(&_TieredElection.TransactOpts, b)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(p address) returns()
func (_TieredElection *TieredElectionTransactor) RemovePool(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "removePool", p)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(p address) returns()
func (_TieredElection *TieredElectionSession) RemovePool(p common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.RemovePool(&_TieredElection.TransactOpts, p)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(p address) returns()
func (_TieredElection *TieredElectionTransactorSession) RemovePool(p common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.RemovePool(&_TieredElection.TransactOpts, p)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredElection *TieredElectionTransactor) RemoveSelf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "removeSelf")
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredElection *TieredElectionSession) RemoveSelf() (*types.Transaction, error) {
	return _TieredElection.Contract.RemoveSelf(&_TieredElection.TransactOpts)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredElection *TieredElectionTransactorSession) RemoveSelf() (*types.Transaction, error) {
	return _TieredElection.Contract.RemoveSelf(&_TieredElection.TransactOpts)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_TieredElection *TieredElectionTransactor) SetPrivateKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "setPrivateKey", key)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_TieredElection *TieredElectionSession) SetPrivateKey(key string) (*types.Transaction, error) {
	return _TieredElection.Contract.SetPrivateKey(&_TieredElection.TransactOpts, key)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_TieredElection *TieredElectionTransactorSession) SetPrivateKey(key string) (*types.Transaction, error) {
	return _TieredElection.Contract.SetPrivateKey(&_TieredElection.TransactOpts, key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_TieredElection *TieredElectionTransactor) SetPublicKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "setPublicKey", key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_TieredElection *TieredElectionSession) SetPublicKey(key string) (*types.Transaction, error) {
	return _TieredElection.Contract.SetPublicKey(&_TieredElection.TransactOpts, key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_TieredElection *TieredElectionTransactorSession) SetPublicKey(key string) (*types.Transaction, error) {
	return _TieredElection.Contract.SetPublicKey(&_TieredElection.TransactOpts, key)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredElection *TieredElectionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredElection *TieredElectionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.TransferOwnership(&_TieredElection.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredElection *TieredElectionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TieredElection.Contract.TransferOwnership(&_TieredElection.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredElection *TieredElectionTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredElection.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredElection *TieredElectionSession) Unlock() (*types.Transaction, error) {
	return _TieredElection.Contract.Unlock(&_TieredElection.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredElection *TieredElectionTransactorSession) Unlock() (*types.Transaction, error) {
	return _TieredElection.Contract.Unlock(&_TieredElection.TransactOpts)
}
