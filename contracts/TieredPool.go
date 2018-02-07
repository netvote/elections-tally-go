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

// TieredPoolABI is the input ABI used to generate the binding from.
const TieredPoolABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"ballotExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionPhase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createdBy\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkBallots\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"addAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"election\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteId\",\"type\":\"bytes32\"},{\"name\":\"vote\",\"type\":\"string\"},{\"name\":\"passphrase\",\"type\":\"string\"}],\"name\":\"castVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeSelf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"removeAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVoteAt\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBallotCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"removeBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteId\",\"type\":\"bytes32\"},{\"name\":\"vote\",\"type\":\"string\"},{\"name\":\"passphrase\",\"type\":\"string\"}],\"name\":\"updateVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"getBallotIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVoteCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"addBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"hashedUserId\",\"type\":\"bytes32\"},{\"name\":\"el\",\"type\":\"address\"},{\"name\":\"gw\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteId\",\"type\":\"bytes32\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteId\",\"type\":\"bytes32\"}],\"name\":\"UpdateVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TieredPool is an auto generated Go binding around an Ethereum contract.
type TieredPool struct {
	TieredPoolCaller     // Read-only binding to the contract
	TieredPoolTransactor // Write-only binding to the contract
}

// TieredPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type TieredPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TieredPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TieredPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TieredPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TieredPoolSession struct {
	Contract     *TieredPool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TieredPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TieredPoolCallerSession struct {
	Contract *TieredPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TieredPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TieredPoolTransactorSession struct {
	Contract     *TieredPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TieredPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type TieredPoolRaw struct {
	Contract *TieredPool // Generic contract binding to access the raw methods on
}

// TieredPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TieredPoolCallerRaw struct {
	Contract *TieredPoolCaller // Generic read-only contract binding to access the raw methods on
}

// TieredPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TieredPoolTransactorRaw struct {
	Contract *TieredPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTieredPool creates a new instance of TieredPool, bound to a specific deployed contract.
func NewTieredPool(address common.Address, backend bind.ContractBackend) (*TieredPool, error) {
	contract, err := bindTieredPool(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TieredPool{TieredPoolCaller: TieredPoolCaller{contract: contract}, TieredPoolTransactor: TieredPoolTransactor{contract: contract}}, nil
}

// NewTieredPoolCaller creates a new read-only instance of TieredPool, bound to a specific deployed contract.
func NewTieredPoolCaller(address common.Address, caller bind.ContractCaller) (*TieredPoolCaller, error) {
	contract, err := bindTieredPool(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TieredPoolCaller{contract: contract}, nil
}

// NewTieredPoolTransactor creates a new write-only instance of TieredPool, bound to a specific deployed contract.
func NewTieredPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*TieredPoolTransactor, error) {
	contract, err := bindTieredPool(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TieredPoolTransactor{contract: contract}, nil
}

// bindTieredPool binds a generic wrapper to an already deployed contract.
func bindTieredPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TieredPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TieredPool *TieredPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TieredPool.Contract.TieredPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TieredPool *TieredPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredPool.Contract.TieredPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TieredPool *TieredPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TieredPool.Contract.TieredPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TieredPool *TieredPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TieredPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TieredPool *TieredPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TieredPool *TieredPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TieredPool.Contract.contract.Transact(opts, method, params...)
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_TieredPool *TieredPoolCaller) BallotExists(opts *bind.CallOpts, b common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "ballotExists", b)
	return *ret0, err
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_TieredPool *TieredPoolSession) BallotExists(b common.Address) (bool, error) {
	return _TieredPool.Contract.BallotExists(&_TieredPool.CallOpts, b)
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_TieredPool *TieredPoolCallerSession) BallotExists(b common.Address) (bool, error) {
	return _TieredPool.Contract.BallotExists(&_TieredPool.CallOpts, b)
}

// CheckBallots is a free data retrieval call binding the contract method 0x3abfb2e9.
//
// Solidity: function checkBallots() constant returns(bool)
func (_TieredPool *TieredPoolCaller) CheckBallots(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "checkBallots")
	return *ret0, err
}

// CheckBallots is a free data retrieval call binding the contract method 0x3abfb2e9.
//
// Solidity: function checkBallots() constant returns(bool)
func (_TieredPool *TieredPoolSession) CheckBallots() (bool, error) {
	return _TieredPool.Contract.CheckBallots(&_TieredPool.CallOpts)
}

// CheckBallots is a free data retrieval call binding the contract method 0x3abfb2e9.
//
// Solidity: function checkBallots() constant returns(bool)
func (_TieredPool *TieredPoolCallerSession) CheckBallots() (bool, error) {
	return _TieredPool.Contract.CheckBallots(&_TieredPool.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredPool *TieredPoolCaller) CheckConfig(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "checkConfig")
	return *ret0, err
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredPool *TieredPoolSession) CheckConfig() (bool, error) {
	return _TieredPool.Contract.CheckConfig(&_TieredPool.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_TieredPool *TieredPoolCallerSession) CheckConfig() (bool, error) {
	return _TieredPool.Contract.CheckConfig(&_TieredPool.CallOpts)
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_TieredPool *TieredPoolCaller) CreatedBy(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "createdBy")
	return *ret0, err
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_TieredPool *TieredPoolSession) CreatedBy() (string, error) {
	return _TieredPool.Contract.CreatedBy(&_TieredPool.CallOpts)
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_TieredPool *TieredPoolCallerSession) CreatedBy() (string, error) {
	return _TieredPool.Contract.CreatedBy(&_TieredPool.CallOpts)
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_TieredPool *TieredPoolCaller) Election(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "election")
	return *ret0, err
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_TieredPool *TieredPoolSession) Election() (common.Address, error) {
	return _TieredPool.Contract.Election(&_TieredPool.CallOpts)
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_TieredPool *TieredPoolCallerSession) Election() (common.Address, error) {
	return _TieredPool.Contract.Election(&_TieredPool.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredPool *TieredPoolCaller) ElectionPhase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "electionPhase")
	return *ret0, err
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredPool *TieredPoolSession) ElectionPhase() (uint8, error) {
	return _TieredPool.Contract.ElectionPhase(&_TieredPool.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_TieredPool *TieredPoolCallerSession) ElectionPhase() (uint8, error) {
	return _TieredPool.Contract.ElectionPhase(&_TieredPool.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_TieredPool *TieredPoolCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "gateway")
	return *ret0, err
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_TieredPool *TieredPoolSession) Gateway() (common.Address, error) {
	return _TieredPool.Contract.Gateway(&_TieredPool.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_TieredPool *TieredPoolCallerSession) Gateway() (common.Address, error) {
	return _TieredPool.Contract.Gateway(&_TieredPool.CallOpts)
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_TieredPool *TieredPoolCaller) GetBallot(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "getBallot", index)
	return *ret0, err
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_TieredPool *TieredPoolSession) GetBallot(index *big.Int) (common.Address, error) {
	return _TieredPool.Contract.GetBallot(&_TieredPool.CallOpts, index)
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_TieredPool *TieredPoolCallerSession) GetBallot(index *big.Int) (common.Address, error) {
	return _TieredPool.Contract.GetBallot(&_TieredPool.CallOpts, index)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_TieredPool *TieredPoolCaller) GetBallotCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "getBallotCount")
	return *ret0, err
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_TieredPool *TieredPoolSession) GetBallotCount() (*big.Int, error) {
	return _TieredPool.Contract.GetBallotCount(&_TieredPool.CallOpts)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_TieredPool *TieredPoolCallerSession) GetBallotCount() (*big.Int, error) {
	return _TieredPool.Contract.GetBallotCount(&_TieredPool.CallOpts)
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_TieredPool *TieredPoolCaller) GetBallotIndex(opts *bind.CallOpts, b common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "getBallotIndex", b)
	return *ret0, err
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_TieredPool *TieredPoolSession) GetBallotIndex(b common.Address) (*big.Int, error) {
	return _TieredPool.Contract.GetBallotIndex(&_TieredPool.CallOpts, b)
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_TieredPool *TieredPoolCallerSession) GetBallotIndex(b common.Address) (*big.Int, error) {
	return _TieredPool.Contract.GetBallotIndex(&_TieredPool.CallOpts, b)
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_TieredPool *TieredPoolCaller) GetVoteAt(opts *bind.CallOpts, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "getVoteAt", index)
	return *ret0, err
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_TieredPool *TieredPoolSession) GetVoteAt(index *big.Int) (string, error) {
	return _TieredPool.Contract.GetVoteAt(&_TieredPool.CallOpts, index)
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_TieredPool *TieredPoolCallerSession) GetVoteAt(index *big.Int) (string, error) {
	return _TieredPool.Contract.GetVoteAt(&_TieredPool.CallOpts, index)
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_TieredPool *TieredPoolCaller) GetVoteCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "getVoteCount")
	return *ret0, err
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_TieredPool *TieredPoolSession) GetVoteCount() (*big.Int, error) {
	return _TieredPool.Contract.GetVoteCount(&_TieredPool.CallOpts)
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_TieredPool *TieredPoolCallerSession) GetVoteCount() (*big.Int, error) {
	return _TieredPool.Contract.GetVoteCount(&_TieredPool.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredPool *TieredPoolCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredPool *TieredPoolSession) IsAdmin(addr common.Address) (bool, error) {
	return _TieredPool.Contract.IsAdmin(&_TieredPool.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_TieredPool *TieredPoolCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _TieredPool.Contract.IsAdmin(&_TieredPool.CallOpts, addr)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_TieredPool *TieredPoolCaller) IsAuthorized(opts *bind.CallOpts, entry [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "isAuthorized", entry)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_TieredPool *TieredPoolSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _TieredPool.Contract.IsAuthorized(&_TieredPool.CallOpts, entry)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_TieredPool *TieredPoolCallerSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _TieredPool.Contract.IsAuthorized(&_TieredPool.CallOpts, entry)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredPool *TieredPoolCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredPool *TieredPoolSession) IsClosed() (bool, error) {
	return _TieredPool.Contract.IsClosed(&_TieredPool.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_TieredPool *TieredPoolCallerSession) IsClosed() (bool, error) {
	return _TieredPool.Contract.IsClosed(&_TieredPool.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredPool *TieredPoolCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredPool *TieredPoolSession) IsLocked() (bool, error) {
	return _TieredPool.Contract.IsLocked(&_TieredPool.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_TieredPool *TieredPoolCallerSession) IsLocked() (bool, error) {
	return _TieredPool.Contract.IsLocked(&_TieredPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredPool *TieredPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredPool *TieredPoolSession) Owner() (common.Address, error) {
	return _TieredPool.Contract.Owner(&_TieredPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TieredPool *TieredPoolCallerSession) Owner() (common.Address, error) {
	return _TieredPool.Contract.Owner(&_TieredPool.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_TieredPool *TieredPoolCaller) Votes(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TieredPool.contract.Call(opts, out, "votes", arg0)
	return *ret0, err
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_TieredPool *TieredPoolSession) Votes(arg0 [32]byte) (string, error) {
	return _TieredPool.Contract.Votes(&_TieredPool.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_TieredPool *TieredPoolCallerSession) Votes(arg0 [32]byte) (string, error) {
	return _TieredPool.Contract.Votes(&_TieredPool.CallOpts, arg0)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredPool *TieredPoolTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredPool *TieredPoolSession) Abort() (*types.Transaction, error) {
	return _TieredPool.Contract.Abort(&_TieredPool.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_TieredPool *TieredPoolTransactorSession) Abort() (*types.Transaction, error) {
	return _TieredPool.Contract.Abort(&_TieredPool.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredPool *TieredPoolTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredPool *TieredPoolSession) Activate() (*types.Transaction, error) {
	return _TieredPool.Contract.Activate(&_TieredPool.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_TieredPool *TieredPoolTransactorSession) Activate() (*types.Transaction, error) {
	return _TieredPool.Contract.Activate(&_TieredPool.TransactOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredPool *TieredPoolTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredPool *TieredPoolSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.AddAdmin(&_TieredPool.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_TieredPool *TieredPoolTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.AddAdmin(&_TieredPool.TransactOpts, addr)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_TieredPool *TieredPoolTransactor) AddAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "addAuthorized", entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_TieredPool *TieredPoolSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _TieredPool.Contract.AddAuthorized(&_TieredPool.TransactOpts, entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_TieredPool *TieredPoolTransactorSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _TieredPool.Contract.AddAuthorized(&_TieredPool.TransactOpts, entry)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_TieredPool *TieredPoolTransactor) AddBallot(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "addBallot", b)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_TieredPool *TieredPoolSession) AddBallot(b common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.AddBallot(&_TieredPool.TransactOpts, b)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_TieredPool *TieredPoolTransactorSession) AddBallot(b common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.AddBallot(&_TieredPool.TransactOpts, b)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_TieredPool *TieredPoolTransactor) CastVote(opts *bind.TransactOpts, voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "castVote", voteId, vote, passphrase)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_TieredPool *TieredPoolSession) CastVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _TieredPool.Contract.CastVote(&_TieredPool.TransactOpts, voteId, vote, passphrase)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_TieredPool *TieredPoolTransactorSession) CastVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _TieredPool.Contract.CastVote(&_TieredPool.TransactOpts, voteId, vote, passphrase)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredPool *TieredPoolTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredPool *TieredPoolSession) Close() (*types.Transaction, error) {
	return _TieredPool.Contract.Close(&_TieredPool.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_TieredPool *TieredPoolTransactorSession) Close() (*types.Transaction, error) {
	return _TieredPool.Contract.Close(&_TieredPool.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredPool *TieredPoolTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredPool *TieredPoolSession) Lock() (*types.Transaction, error) {
	return _TieredPool.Contract.Lock(&_TieredPool.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TieredPool *TieredPoolTransactorSession) Lock() (*types.Transaction, error) {
	return _TieredPool.Contract.Lock(&_TieredPool.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredPool *TieredPoolTransactor) RemoveAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "removeAdmin", addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredPool *TieredPoolSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.RemoveAdmin(&_TieredPool.TransactOpts, addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_TieredPool *TieredPoolTransactorSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.RemoveAdmin(&_TieredPool.TransactOpts, addr)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_TieredPool *TieredPoolTransactor) RemoveAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "removeAuthorized", entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_TieredPool *TieredPoolSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _TieredPool.Contract.RemoveAuthorized(&_TieredPool.TransactOpts, entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_TieredPool *TieredPoolTransactorSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _TieredPool.Contract.RemoveAuthorized(&_TieredPool.TransactOpts, entry)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_TieredPool *TieredPoolTransactor) RemoveBallot(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "removeBallot", b)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_TieredPool *TieredPoolSession) RemoveBallot(b common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.RemoveBallot(&_TieredPool.TransactOpts, b)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_TieredPool *TieredPoolTransactorSession) RemoveBallot(b common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.RemoveBallot(&_TieredPool.TransactOpts, b)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredPool *TieredPoolTransactor) RemoveSelf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "removeSelf")
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredPool *TieredPoolSession) RemoveSelf() (*types.Transaction, error) {
	return _TieredPool.Contract.RemoveSelf(&_TieredPool.TransactOpts)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_TieredPool *TieredPoolTransactorSession) RemoveSelf() (*types.Transaction, error) {
	return _TieredPool.Contract.RemoveSelf(&_TieredPool.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredPool *TieredPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredPool *TieredPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.TransferOwnership(&_TieredPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TieredPool *TieredPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TieredPool.Contract.TransferOwnership(&_TieredPool.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredPool *TieredPoolTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredPool *TieredPoolSession) Unlock() (*types.Transaction, error) {
	return _TieredPool.Contract.Unlock(&_TieredPool.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TieredPool *TieredPoolTransactorSession) Unlock() (*types.Transaction, error) {
	return _TieredPool.Contract.Unlock(&_TieredPool.TransactOpts)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_TieredPool *TieredPoolTransactor) UpdateVote(opts *bind.TransactOpts, voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _TieredPool.contract.Transact(opts, "updateVote", voteId, vote, passphrase)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_TieredPool *TieredPoolSession) UpdateVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _TieredPool.Contract.UpdateVote(&_TieredPool.TransactOpts, voteId, vote, passphrase)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_TieredPool *TieredPoolTransactorSession) UpdateVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _TieredPool.Contract.UpdateVote(&_TieredPool.TransactOpts, voteId, vote, passphrase)
}
