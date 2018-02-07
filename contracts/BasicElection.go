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

// BasicElectionABI is the input ABI used to generate the binding from.
const BasicElectionABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"ballotExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionPhase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createdBy\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"addAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionType\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"privateKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"election\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteId\",\"type\":\"bytes32\"},{\"name\":\"vote\",\"type\":\"string\"},{\"name\":\"passphrase\",\"type\":\"string\"}],\"name\":\"castVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeSelf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entry\",\"type\":\"bytes32\"}],\"name\":\"removeAuthorized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVoteAt\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBallotCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"metadataLocation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowVoteUpdates\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"removeBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteId\",\"type\":\"bytes32\"},{\"name\":\"vote\",\"type\":\"string\"},{\"name\":\"passphrase\",\"type\":\"string\"}],\"name\":\"updateVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"getBallotIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deductVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"setPrivateKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVoteCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"address\"}],\"name\":\"addBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"hashedUserId\",\"type\":\"bytes32\"},{\"name\":\"allowanceAddress\",\"type\":\"address\"},{\"name\":\"ownerOfAllowance\",\"type\":\"address\"},{\"name\":\"allowUpdates\",\"type\":\"bool\"},{\"name\":\"revealerAddress\",\"type\":\"address\"},{\"name\":\"location\",\"type\":\"string\"},{\"name\":\"gatewayAddress\",\"type\":\"address\"},{\"name\":\"autoActivate\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"KeyReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteId\",\"type\":\"bytes32\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteId\",\"type\":\"bytes32\"}],\"name\":\"UpdateVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BasicElection is an auto generated Go binding around an Ethereum contract.
type BasicElection struct {
	BasicElectionCaller     // Read-only binding to the contract
	BasicElectionTransactor // Write-only binding to the contract
}

// BasicElectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicElectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicElectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicElectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicElectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicElectionSession struct {
	Contract     *BasicElection    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicElectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicElectionCallerSession struct {
	Contract *BasicElectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BasicElectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicElectionTransactorSession struct {
	Contract     *BasicElectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BasicElectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicElectionRaw struct {
	Contract *BasicElection // Generic contract binding to access the raw methods on
}

// BasicElectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicElectionCallerRaw struct {
	Contract *BasicElectionCaller // Generic read-only contract binding to access the raw methods on
}

// BasicElectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicElectionTransactorRaw struct {
	Contract *BasicElectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicElection creates a new instance of BasicElection, bound to a specific deployed contract.
func NewBasicElection(address common.Address, backend bind.ContractBackend) (*BasicElection, error) {
	contract, err := bindBasicElection(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicElection{BasicElectionCaller: BasicElectionCaller{contract: contract}, BasicElectionTransactor: BasicElectionTransactor{contract: contract}}, nil
}

// NewBasicElectionCaller creates a new read-only instance of BasicElection, bound to a specific deployed contract.
func NewBasicElectionCaller(address common.Address, caller bind.ContractCaller) (*BasicElectionCaller, error) {
	contract, err := bindBasicElection(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BasicElectionCaller{contract: contract}, nil
}

// NewBasicElectionTransactor creates a new write-only instance of BasicElection, bound to a specific deployed contract.
func NewBasicElectionTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicElectionTransactor, error) {
	contract, err := bindBasicElection(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BasicElectionTransactor{contract: contract}, nil
}

// bindBasicElection binds a generic wrapper to an already deployed contract.
func bindBasicElection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicElectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicElection *BasicElectionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicElection.Contract.BasicElectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicElection *BasicElectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.Contract.BasicElectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicElection *BasicElectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicElection.Contract.BasicElectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicElection *BasicElectionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicElection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicElection *BasicElectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicElection *BasicElectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicElection.Contract.contract.Transact(opts, method, params...)
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_BasicElection *BasicElectionCaller) AllowVoteUpdates(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "allowVoteUpdates")
	return *ret0, err
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_BasicElection *BasicElectionSession) AllowVoteUpdates() (bool, error) {
	return _BasicElection.Contract.AllowVoteUpdates(&_BasicElection.CallOpts)
}

// AllowVoteUpdates is a free data retrieval call binding the contract method 0xbf41bf36.
//
// Solidity: function allowVoteUpdates() constant returns(bool)
func (_BasicElection *BasicElectionCallerSession) AllowVoteUpdates() (bool, error) {
	return _BasicElection.Contract.AllowVoteUpdates(&_BasicElection.CallOpts)
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_BasicElection *BasicElectionCaller) BallotExists(opts *bind.CallOpts, b common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "ballotExists", b)
	return *ret0, err
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_BasicElection *BasicElectionSession) BallotExists(b common.Address) (bool, error) {
	return _BasicElection.Contract.BallotExists(&_BasicElection.CallOpts, b)
}

// BallotExists is a free data retrieval call binding the contract method 0x138b1fa0.
//
// Solidity: function ballotExists(b address) constant returns(bool)
func (_BasicElection *BasicElectionCallerSession) BallotExists(b common.Address) (bool, error) {
	return _BasicElection.Contract.BallotExists(&_BasicElection.CallOpts, b)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BasicElection *BasicElectionCaller) CheckConfig(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "checkConfig")
	return *ret0, err
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BasicElection *BasicElectionSession) CheckConfig() (bool, error) {
	return _BasicElection.Contract.CheckConfig(&_BasicElection.CallOpts)
}

// CheckConfig is a free data retrieval call binding the contract method 0xf098fe47.
//
// Solidity: function checkConfig() constant returns(bool)
func (_BasicElection *BasicElectionCallerSession) CheckConfig() (bool, error) {
	return _BasicElection.Contract.CheckConfig(&_BasicElection.CallOpts)
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_BasicElection *BasicElectionCaller) CreatedBy(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "createdBy")
	return *ret0, err
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_BasicElection *BasicElectionSession) CreatedBy() (string, error) {
	return _BasicElection.Contract.CreatedBy(&_BasicElection.CallOpts)
}

// CreatedBy is a free data retrieval call binding the contract method 0x3a5673a4.
//
// Solidity: function createdBy() constant returns(string)
func (_BasicElection *BasicElectionCallerSession) CreatedBy() (string, error) {
	return _BasicElection.Contract.CreatedBy(&_BasicElection.CallOpts)
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_BasicElection *BasicElectionCaller) Election(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "election")
	return *ret0, err
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_BasicElection *BasicElectionSession) Election() (common.Address, error) {
	return _BasicElection.Contract.Election(&_BasicElection.CallOpts)
}

// Election is a free data retrieval call binding the contract method 0x4b1d3ede.
//
// Solidity: function election() constant returns(address)
func (_BasicElection *BasicElectionCallerSession) Election() (common.Address, error) {
	return _BasicElection.Contract.Election(&_BasicElection.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BasicElection *BasicElectionCaller) ElectionPhase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "electionPhase")
	return *ret0, err
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BasicElection *BasicElectionSession) ElectionPhase() (uint8, error) {
	return _BasicElection.Contract.ElectionPhase(&_BasicElection.CallOpts)
}

// ElectionPhase is a free data retrieval call binding the contract method 0x265050b6.
//
// Solidity: function electionPhase() constant returns(uint8)
func (_BasicElection *BasicElectionCallerSession) ElectionPhase() (uint8, error) {
	return _BasicElection.Contract.ElectionPhase(&_BasicElection.CallOpts)
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_BasicElection *BasicElectionCaller) ElectionType(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "electionType")
	return *ret0, err
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_BasicElection *BasicElectionSession) ElectionType() (string, error) {
	return _BasicElection.Contract.ElectionType(&_BasicElection.CallOpts)
}

// ElectionType is a free data retrieval call binding the contract method 0x4885c72a.
//
// Solidity: function electionType() constant returns(string)
func (_BasicElection *BasicElectionCallerSession) ElectionType() (string, error) {
	return _BasicElection.Contract.ElectionType(&_BasicElection.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_BasicElection *BasicElectionCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "gateway")
	return *ret0, err
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_BasicElection *BasicElectionSession) Gateway() (common.Address, error) {
	return _BasicElection.Contract.Gateway(&_BasicElection.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_BasicElection *BasicElectionCallerSession) Gateway() (common.Address, error) {
	return _BasicElection.Contract.Gateway(&_BasicElection.CallOpts)
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_BasicElection *BasicElectionCaller) GetBallot(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "getBallot", index)
	return *ret0, err
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_BasicElection *BasicElectionSession) GetBallot(index *big.Int) (common.Address, error) {
	return _BasicElection.Contract.GetBallot(&_BasicElection.CallOpts, index)
}

// GetBallot is a free data retrieval call binding the contract method 0xf9d5ee75.
//
// Solidity: function getBallot(index uint256) constant returns(address)
func (_BasicElection *BasicElectionCallerSession) GetBallot(index *big.Int) (common.Address, error) {
	return _BasicElection.Contract.GetBallot(&_BasicElection.CallOpts, index)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_BasicElection *BasicElectionCaller) GetBallotCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "getBallotCount")
	return *ret0, err
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_BasicElection *BasicElectionSession) GetBallotCount() (*big.Int, error) {
	return _BasicElection.Contract.GetBallotCount(&_BasicElection.CallOpts)
}

// GetBallotCount is a free data retrieval call binding the contract method 0xb4741495.
//
// Solidity: function getBallotCount() constant returns(uint256)
func (_BasicElection *BasicElectionCallerSession) GetBallotCount() (*big.Int, error) {
	return _BasicElection.Contract.GetBallotCount(&_BasicElection.CallOpts)
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_BasicElection *BasicElectionCaller) GetBallotIndex(opts *bind.CallOpts, b common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "getBallotIndex", b)
	return *ret0, err
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_BasicElection *BasicElectionSession) GetBallotIndex(b common.Address) (*big.Int, error) {
	return _BasicElection.Contract.GetBallotIndex(&_BasicElection.CallOpts, b)
}

// GetBallotIndex is a free data retrieval call binding the contract method 0xdb30eafb.
//
// Solidity: function getBallotIndex(b address) constant returns(uint256)
func (_BasicElection *BasicElectionCallerSession) GetBallotIndex(b common.Address) (*big.Int, error) {
	return _BasicElection.Contract.GetBallotIndex(&_BasicElection.CallOpts, b)
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_BasicElection *BasicElectionCaller) GetVoteAt(opts *bind.CallOpts, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "getVoteAt", index)
	return *ret0, err
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_BasicElection *BasicElectionSession) GetVoteAt(index *big.Int) (string, error) {
	return _BasicElection.Contract.GetVoteAt(&_BasicElection.CallOpts, index)
}

// GetVoteAt is a free data retrieval call binding the contract method 0x9cc8f2b3.
//
// Solidity: function getVoteAt(index uint256) constant returns(string)
func (_BasicElection *BasicElectionCallerSession) GetVoteAt(index *big.Int) (string, error) {
	return _BasicElection.Contract.GetVoteAt(&_BasicElection.CallOpts, index)
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_BasicElection *BasicElectionCaller) GetVoteCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "getVoteCount")
	return *ret0, err
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_BasicElection *BasicElectionSession) GetVoteCount() (*big.Int, error) {
	return _BasicElection.Contract.GetVoteCount(&_BasicElection.CallOpts)
}

// GetVoteCount is a free data retrieval call binding the contract method 0xe7b3387c.
//
// Solidity: function getVoteCount() constant returns(uint256)
func (_BasicElection *BasicElectionCallerSession) GetVoteCount() (*big.Int, error) {
	return _BasicElection.Contract.GetVoteCount(&_BasicElection.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BasicElection *BasicElectionCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BasicElection *BasicElectionSession) IsAdmin(addr common.Address) (bool, error) {
	return _BasicElection.Contract.IsAdmin(&_BasicElection.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_BasicElection *BasicElectionCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _BasicElection.Contract.IsAdmin(&_BasicElection.CallOpts, addr)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BasicElection *BasicElectionCaller) IsAuthorized(opts *bind.CallOpts, entry [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "isAuthorized", entry)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BasicElection *BasicElectionSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _BasicElection.Contract.IsAuthorized(&_BasicElection.CallOpts, entry)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x62f4ed90.
//
// Solidity: function isAuthorized(entry bytes32) constant returns(bool)
func (_BasicElection *BasicElectionCallerSession) IsAuthorized(entry [32]byte) (bool, error) {
	return _BasicElection.Contract.IsAuthorized(&_BasicElection.CallOpts, entry)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BasicElection *BasicElectionCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BasicElection *BasicElectionSession) IsClosed() (bool, error) {
	return _BasicElection.Contract.IsClosed(&_BasicElection.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BasicElection *BasicElectionCallerSession) IsClosed() (bool, error) {
	return _BasicElection.Contract.IsClosed(&_BasicElection.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BasicElection *BasicElectionCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BasicElection *BasicElectionSession) IsLocked() (bool, error) {
	return _BasicElection.Contract.IsLocked(&_BasicElection.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BasicElection *BasicElectionCallerSession) IsLocked() (bool, error) {
	return _BasicElection.Contract.IsLocked(&_BasicElection.CallOpts)
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_BasicElection *BasicElectionCaller) MetadataLocation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "metadataLocation")
	return *ret0, err
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_BasicElection *BasicElectionSession) MetadataLocation() (string, error) {
	return _BasicElection.Contract.MetadataLocation(&_BasicElection.CallOpts)
}

// MetadataLocation is a free data retrieval call binding the contract method 0xb86901af.
//
// Solidity: function metadataLocation() constant returns(string)
func (_BasicElection *BasicElectionCallerSession) MetadataLocation() (string, error) {
	return _BasicElection.Contract.MetadataLocation(&_BasicElection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasicElection *BasicElectionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasicElection *BasicElectionSession) Owner() (common.Address, error) {
	return _BasicElection.Contract.Owner(&_BasicElection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasicElection *BasicElectionCallerSession) Owner() (common.Address, error) {
	return _BasicElection.Contract.Owner(&_BasicElection.CallOpts)
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_BasicElection *BasicElectionCaller) PrivateKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "privateKey")
	return *ret0, err
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_BasicElection *BasicElectionSession) PrivateKey() (string, error) {
	return _BasicElection.Contract.PrivateKey(&_BasicElection.CallOpts)
}

// PrivateKey is a free data retrieval call binding the contract method 0x49da5a0f.
//
// Solidity: function privateKey() constant returns(string)
func (_BasicElection *BasicElectionCallerSession) PrivateKey() (string, error) {
	return _BasicElection.Contract.PrivateKey(&_BasicElection.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_BasicElection *BasicElectionCaller) PublicKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "publicKey")
	return *ret0, err
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_BasicElection *BasicElectionSession) PublicKey() (string, error) {
	return _BasicElection.Contract.PublicKey(&_BasicElection.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_BasicElection *BasicElectionCallerSession) PublicKey() (string, error) {
	return _BasicElection.Contract.PublicKey(&_BasicElection.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_BasicElection *BasicElectionCaller) Votes(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BasicElection.contract.Call(opts, out, "votes", arg0)
	return *ret0, err
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_BasicElection *BasicElectionSession) Votes(arg0 [32]byte) (string, error) {
	return _BasicElection.Contract.Votes(&_BasicElection.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes( bytes32) constant returns(string)
func (_BasicElection *BasicElectionCallerSession) Votes(arg0 [32]byte) (string, error) {
	return _BasicElection.Contract.Votes(&_BasicElection.CallOpts, arg0)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BasicElection *BasicElectionTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BasicElection *BasicElectionSession) Abort() (*types.Transaction, error) {
	return _BasicElection.Contract.Abort(&_BasicElection.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_BasicElection *BasicElectionTransactorSession) Abort() (*types.Transaction, error) {
	return _BasicElection.Contract.Abort(&_BasicElection.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BasicElection *BasicElectionTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BasicElection *BasicElectionSession) Activate() (*types.Transaction, error) {
	return _BasicElection.Contract.Activate(&_BasicElection.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_BasicElection *BasicElectionTransactorSession) Activate() (*types.Transaction, error) {
	return _BasicElection.Contract.Activate(&_BasicElection.TransactOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BasicElection *BasicElectionTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BasicElection *BasicElectionSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.AddAdmin(&_BasicElection.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(addr address) returns()
func (_BasicElection *BasicElectionTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.AddAdmin(&_BasicElection.TransactOpts, addr)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BasicElection *BasicElectionTransactor) AddAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "addAuthorized", entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BasicElection *BasicElectionSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BasicElection.Contract.AddAuthorized(&_BasicElection.TransactOpts, entry)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0x41af099d.
//
// Solidity: function addAuthorized(entry bytes32) returns()
func (_BasicElection *BasicElectionTransactorSession) AddAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BasicElection.Contract.AddAuthorized(&_BasicElection.TransactOpts, entry)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_BasicElection *BasicElectionTransactor) AddBallot(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "addBallot", b)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_BasicElection *BasicElectionSession) AddBallot(b common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.AddBallot(&_BasicElection.TransactOpts, b)
}

// AddBallot is a paid mutator transaction binding the contract method 0xf8c74afa.
//
// Solidity: function addBallot(b address) returns()
func (_BasicElection *BasicElectionTransactorSession) AddBallot(b common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.AddBallot(&_BasicElection.TransactOpts, b)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasicElection *BasicElectionTransactor) CastVote(opts *bind.TransactOpts, voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "castVote", voteId, vote, passphrase)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasicElection *BasicElectionSession) CastVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasicElection.Contract.CastVote(&_BasicElection.TransactOpts, voteId, vote, passphrase)
}

// CastVote is a paid mutator transaction binding the contract method 0x552c3a02.
//
// Solidity: function castVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasicElection *BasicElectionTransactorSession) CastVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasicElection.Contract.CastVote(&_BasicElection.TransactOpts, voteId, vote, passphrase)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BasicElection *BasicElectionTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BasicElection *BasicElectionSession) Close() (*types.Transaction, error) {
	return _BasicElection.Contract.Close(&_BasicElection.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BasicElection *BasicElectionTransactorSession) Close() (*types.Transaction, error) {
	return _BasicElection.Contract.Close(&_BasicElection.TransactOpts)
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_BasicElection *BasicElectionTransactor) DeductVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "deductVote")
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_BasicElection *BasicElectionSession) DeductVote() (*types.Transaction, error) {
	return _BasicElection.Contract.DeductVote(&_BasicElection.TransactOpts)
}

// DeductVote is a paid mutator transaction binding the contract method 0xde61a732.
//
// Solidity: function deductVote() returns()
func (_BasicElection *BasicElectionTransactorSession) DeductVote() (*types.Transaction, error) {
	return _BasicElection.Contract.DeductVote(&_BasicElection.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BasicElection *BasicElectionTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BasicElection *BasicElectionSession) Lock() (*types.Transaction, error) {
	return _BasicElection.Contract.Lock(&_BasicElection.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_BasicElection *BasicElectionTransactorSession) Lock() (*types.Transaction, error) {
	return _BasicElection.Contract.Lock(&_BasicElection.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BasicElection *BasicElectionTransactor) RemoveAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "removeAdmin", addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BasicElection *BasicElectionSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.RemoveAdmin(&_BasicElection.TransactOpts, addr)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(addr address) returns()
func (_BasicElection *BasicElectionTransactorSession) RemoveAdmin(addr common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.RemoveAdmin(&_BasicElection.TransactOpts, addr)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BasicElection *BasicElectionTransactor) RemoveAuthorized(opts *bind.TransactOpts, entry [32]byte) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "removeAuthorized", entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BasicElection *BasicElectionSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BasicElection.Contract.RemoveAuthorized(&_BasicElection.TransactOpts, entry)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x72e789b3.
//
// Solidity: function removeAuthorized(entry bytes32) returns()
func (_BasicElection *BasicElectionTransactorSession) RemoveAuthorized(entry [32]byte) (*types.Transaction, error) {
	return _BasicElection.Contract.RemoveAuthorized(&_BasicElection.TransactOpts, entry)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_BasicElection *BasicElectionTransactor) RemoveBallot(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "removeBallot", b)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_BasicElection *BasicElectionSession) RemoveBallot(b common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.RemoveBallot(&_BasicElection.TransactOpts, b)
}

// RemoveBallot is a paid mutator transaction binding the contract method 0xd71c6772.
//
// Solidity: function removeBallot(b address) returns()
func (_BasicElection *BasicElectionTransactorSession) RemoveBallot(b common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.RemoveBallot(&_BasicElection.TransactOpts, b)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BasicElection *BasicElectionTransactor) RemoveSelf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "removeSelf")
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BasicElection *BasicElectionSession) RemoveSelf() (*types.Transaction, error) {
	return _BasicElection.Contract.RemoveSelf(&_BasicElection.TransactOpts)
}

// RemoveSelf is a paid mutator transaction binding the contract method 0x5e898dac.
//
// Solidity: function removeSelf() returns()
func (_BasicElection *BasicElectionTransactorSession) RemoveSelf() (*types.Transaction, error) {
	return _BasicElection.Contract.RemoveSelf(&_BasicElection.TransactOpts)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_BasicElection *BasicElectionTransactor) SetPrivateKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "setPrivateKey", key)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_BasicElection *BasicElectionSession) SetPrivateKey(key string) (*types.Transaction, error) {
	return _BasicElection.Contract.SetPrivateKey(&_BasicElection.TransactOpts, key)
}

// SetPrivateKey is a paid mutator transaction binding the contract method 0xe0d4f017.
//
// Solidity: function setPrivateKey(key string) returns()
func (_BasicElection *BasicElectionTransactorSession) SetPrivateKey(key string) (*types.Transaction, error) {
	return _BasicElection.Contract.SetPrivateKey(&_BasicElection.TransactOpts, key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_BasicElection *BasicElectionTransactor) SetPublicKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "setPublicKey", key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_BasicElection *BasicElectionSession) SetPublicKey(key string) (*types.Transaction, error) {
	return _BasicElection.Contract.SetPublicKey(&_BasicElection.TransactOpts, key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x6f6fc077.
//
// Solidity: function setPublicKey(key string) returns()
func (_BasicElection *BasicElectionTransactorSession) SetPublicKey(key string) (*types.Transaction, error) {
	return _BasicElection.Contract.SetPublicKey(&_BasicElection.TransactOpts, key)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasicElection *BasicElectionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasicElection *BasicElectionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.TransferOwnership(&_BasicElection.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasicElection *BasicElectionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasicElection.Contract.TransferOwnership(&_BasicElection.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BasicElection *BasicElectionTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BasicElection *BasicElectionSession) Unlock() (*types.Transaction, error) {
	return _BasicElection.Contract.Unlock(&_BasicElection.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_BasicElection *BasicElectionTransactorSession) Unlock() (*types.Transaction, error) {
	return _BasicElection.Contract.Unlock(&_BasicElection.TransactOpts)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasicElection *BasicElectionTransactor) UpdateVote(opts *bind.TransactOpts, voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasicElection.contract.Transact(opts, "updateVote", voteId, vote, passphrase)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasicElection *BasicElectionSession) UpdateVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasicElection.Contract.UpdateVote(&_BasicElection.TransactOpts, voteId, vote, passphrase)
}

// UpdateVote is a paid mutator transaction binding the contract method 0xd8b5708c.
//
// Solidity: function updateVote(voteId bytes32, vote string, passphrase string) returns()
func (_BasicElection *BasicElectionTransactorSession) UpdateVote(voteId [32]byte, vote string, passphrase string) (*types.Transaction, error) {
	return _BasicElection.Contract.UpdateVote(&_BasicElection.TransactOpts, voteId, vote, passphrase)
}
