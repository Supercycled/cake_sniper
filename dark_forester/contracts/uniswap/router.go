// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakeFactoryABI is the input ABI used to generate the binding from.
const IPancakeFactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INIT_CODE_PAIR_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPancakeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IPancakeFactoryFuncSigs = map[string]string{
	"5855a25a": "INIT_CODE_PAIR_HASH()",
	"1e3dd18b": "allPairs(uint256)",
	"574f2ba3": "allPairsLength()",
	"c9c65396": "createPair(address,address)",
	"017e7e58": "feeTo()",
	"094b7415": "feeToSetter()",
	"e6a43905": "getPair(address,address)",
	"f46901ed": "setFeeTo(address)",
	"a2e74af6": "setFeeToSetter(address)",
}

// IPancakeFactory is an auto generated Go binding around an Ethereum contract.
type IPancakeFactory struct {
	IPancakeFactoryCaller     // Read-only binding to the contract
	IPancakeFactoryTransactor // Write-only binding to the contract
	IPancakeFactoryFilterer   // Log filterer for contract events
}

// IPancakeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeFactorySession struct {
	Contract     *IPancakeFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeFactoryCallerSession struct {
	Contract *IPancakeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IPancakeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeFactoryTransactorSession struct {
	Contract     *IPancakeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPancakeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeFactoryRaw struct {
	Contract *IPancakeFactory // Generic contract binding to access the raw methods on
}

// IPancakeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeFactoryCallerRaw struct {
	Contract *IPancakeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IPancakeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeFactoryTransactorRaw struct {
	Contract *IPancakeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeFactory creates a new instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactory(address common.Address, backend bind.ContractBackend) (*IPancakeFactory, error) {
	contract, err := bindIPancakeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactory{IPancakeFactoryCaller: IPancakeFactoryCaller{contract: contract}, IPancakeFactoryTransactor: IPancakeFactoryTransactor{contract: contract}, IPancakeFactoryFilterer: IPancakeFactoryFilterer{contract: contract}}, nil
}

// NewIPancakeFactoryCaller creates a new read-only instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryCaller(address common.Address, caller bind.ContractCaller) (*IPancakeFactoryCaller, error) {
	contract, err := bindIPancakeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryCaller{contract: contract}, nil
}

// NewIPancakeFactoryTransactor creates a new write-only instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeFactoryTransactor, error) {
	contract, err := bindIPancakeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryTransactor{contract: contract}, nil
}

// NewIPancakeFactoryFilterer creates a new log filterer instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeFactoryFilterer, error) {
	contract, err := bindIPancakeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryFilterer{contract: contract}, nil
}

// bindIPancakeFactory binds a generic wrapper to an already deployed contract.
func bindIPancakeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeFactory *IPancakeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeFactory.Contract.IPancakeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeFactory *IPancakeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.IPancakeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeFactory *IPancakeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.IPancakeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeFactory *IPancakeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeFactory *IPancakeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeFactory *IPancakeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.contract.Transact(opts, method, params...)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_IPancakeFactory *IPancakeFactoryCaller) INITCODEPAIRHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "INIT_CODE_PAIR_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_IPancakeFactory *IPancakeFactorySession) INITCODEPAIRHASH() ([32]byte, error) {
	return _IPancakeFactory.Contract.INITCODEPAIRHASH(&_IPancakeFactory.CallOpts)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_IPancakeFactory *IPancakeFactoryCallerSession) INITCODEPAIRHASH() ([32]byte, error) {
	return _IPancakeFactory.Contract.INITCODEPAIRHASH(&_IPancakeFactory.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPancakeFactory *IPancakeFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IPancakeFactory.Contract.AllPairs(&_IPancakeFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IPancakeFactory.Contract.AllPairs(&_IPancakeFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPancakeFactory *IPancakeFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPancakeFactory *IPancakeFactorySession) AllPairsLength() (*big.Int, error) {
	return _IPancakeFactory.Contract.AllPairsLength(&_IPancakeFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPancakeFactory *IPancakeFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _IPancakeFactory.Contract.AllPairsLength(&_IPancakeFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPancakeFactory *IPancakeFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPancakeFactory *IPancakeFactorySession) FeeTo() (common.Address, error) {
	return _IPancakeFactory.Contract.FeeTo(&_IPancakeFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPancakeFactory *IPancakeFactoryCallerSession) FeeTo() (common.Address, error) {
	return _IPancakeFactory.Contract.FeeTo(&_IPancakeFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPancakeFactory *IPancakeFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPancakeFactory *IPancakeFactorySession) FeeToSetter() (common.Address, error) {
	return _IPancakeFactory.Contract.FeeToSetter(&_IPancakeFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPancakeFactory *IPancakeFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _IPancakeFactory.Contract.FeeToSetter(&_IPancakeFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IPancakeFactory.Contract.GetPair(&_IPancakeFactory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IPancakeFactory.Contract.GetPair(&_IPancakeFactory.CallOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPancakeFactory *IPancakeFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPancakeFactory *IPancakeFactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.CreatePair(&_IPancakeFactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPancakeFactory *IPancakeFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.CreatePair(&_IPancakeFactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPancakeFactory *IPancakeFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPancakeFactory *IPancakeFactorySession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.SetFeeTo(&_IPancakeFactory.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPancakeFactory *IPancakeFactoryTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.SetFeeTo(&_IPancakeFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPancakeFactory *IPancakeFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPancakeFactory *IPancakeFactorySession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.SetFeeToSetter(&_IPancakeFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPancakeFactory *IPancakeFactoryTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.SetFeeToSetter(&_IPancakeFactory.TransactOpts, arg0)
}

// IPancakeFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the IPancakeFactory contract.
type IPancakeFactoryPairCreatedIterator struct {
	Event *IPancakeFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *IPancakeFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakeFactoryPairCreated)
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
		it.Event = new(IPancakeFactoryPairCreated)
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
func (it *IPancakeFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakeFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakeFactoryPairCreated represents a PairCreated event raised by the IPancakeFactory contract.
type IPancakeFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPancakeFactory *IPancakeFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*IPancakeFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IPancakeFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryPairCreatedIterator{contract: _IPancakeFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPancakeFactory *IPancakeFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *IPancakeFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IPancakeFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakeFactoryPairCreated)
				if err := _IPancakeFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPancakeFactory *IPancakeFactoryFilterer) ParsePairCreated(log types.Log) (*IPancakeFactoryPairCreated, error) {
	event := new(IPancakeFactoryPairCreated)
	if err := _IPancakeFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairABI is the input ABI used to generate the binding from.
const IPancakePairABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPancakePairFuncSigs maps the 4-byte function signature to its string representation.
var IPancakePairFuncSigs = map[string]string{
	"3644e515": "DOMAIN_SEPARATOR()",
	"ba9a7a56": "MINIMUM_LIQUIDITY()",
	"30adf81f": "PERMIT_TYPEHASH()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"89afcb44": "burn(address)",
	"313ce567": "decimals()",
	"c45a0155": "factory()",
	"0902f1ac": "getReserves()",
	"485cc955": "initialize(address,address)",
	"7464fc3d": "kLast()",
	"6a627842": "mint(address)",
	"06fdde03": "name()",
	"7ecebe00": "nonces(address)",
	"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
	"5909c0d5": "price0CumulativeLast()",
	"5a3d5493": "price1CumulativeLast()",
	"bc25cf77": "skim(address)",
	"022c0d9f": "swap(uint256,uint256,address,bytes)",
	"95d89b41": "symbol()",
	"fff6cae9": "sync()",
	"0dfe1681": "token0()",
	"d21220a7": "token1()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IPancakePair is an auto generated Go binding around an Ethereum contract.
type IPancakePair struct {
	IPancakePairCaller     // Read-only binding to the contract
	IPancakePairTransactor // Write-only binding to the contract
	IPancakePairFilterer   // Log filterer for contract events
}

// IPancakePairCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakePairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakePairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakePairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakePairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakePairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakePairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakePairSession struct {
	Contract     *IPancakePair     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakePairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakePairCallerSession struct {
	Contract *IPancakePairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPancakePairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakePairTransactorSession struct {
	Contract     *IPancakePairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPancakePairRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakePairRaw struct {
	Contract *IPancakePair // Generic contract binding to access the raw methods on
}

// IPancakePairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakePairCallerRaw struct {
	Contract *IPancakePairCaller // Generic read-only contract binding to access the raw methods on
}

// IPancakePairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakePairTransactorRaw struct {
	Contract *IPancakePairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakePair creates a new instance of IPancakePair, bound to a specific deployed contract.
func NewIPancakePair(address common.Address, backend bind.ContractBackend) (*IPancakePair, error) {
	contract, err := bindIPancakePair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakePair{IPancakePairCaller: IPancakePairCaller{contract: contract}, IPancakePairTransactor: IPancakePairTransactor{contract: contract}, IPancakePairFilterer: IPancakePairFilterer{contract: contract}}, nil
}

// NewIPancakePairCaller creates a new read-only instance of IPancakePair, bound to a specific deployed contract.
func NewIPancakePairCaller(address common.Address, caller bind.ContractCaller) (*IPancakePairCaller, error) {
	contract, err := bindIPancakePair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakePairCaller{contract: contract}, nil
}

// NewIPancakePairTransactor creates a new write-only instance of IPancakePair, bound to a specific deployed contract.
func NewIPancakePairTransactor(address common.Address, transactor bind.ContractTransactor) (*IPancakePairTransactor, error) {
	contract, err := bindIPancakePair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakePairTransactor{contract: contract}, nil
}

// NewIPancakePairFilterer creates a new log filterer instance of IPancakePair, bound to a specific deployed contract.
func NewIPancakePairFilterer(address common.Address, filterer bind.ContractFilterer) (*IPancakePairFilterer, error) {
	contract, err := bindIPancakePair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakePairFilterer{contract: contract}, nil
}

// bindIPancakePair binds a generic wrapper to an already deployed contract.
func bindIPancakePair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakePairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakePair *IPancakePairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakePair.Contract.IPancakePairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakePair *IPancakePairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakePair.Contract.IPancakePairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakePair *IPancakePairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakePair.Contract.IPancakePairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakePair *IPancakePairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakePair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakePair *IPancakePairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakePair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakePair *IPancakePairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakePair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPancakePair *IPancakePairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPancakePair *IPancakePairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IPancakePair.Contract.DOMAINSEPARATOR(&_IPancakePair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPancakePair *IPancakePairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IPancakePair.Contract.DOMAINSEPARATOR(&_IPancakePair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_IPancakePair *IPancakePairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_IPancakePair *IPancakePairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _IPancakePair.Contract.MINIMUMLIQUIDITY(&_IPancakePair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _IPancakePair.Contract.MINIMUMLIQUIDITY(&_IPancakePair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IPancakePair *IPancakePairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IPancakePair *IPancakePairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IPancakePair.Contract.PERMITTYPEHASH(&_IPancakePair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IPancakePair *IPancakePairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IPancakePair.Contract.PERMITTYPEHASH(&_IPancakePair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IPancakePair *IPancakePairCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IPancakePair *IPancakePairSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.Allowance(&_IPancakePair.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.Allowance(&_IPancakePair.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.BalanceOf(&_IPancakePair.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.BalanceOf(&_IPancakePair.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IPancakePair *IPancakePairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IPancakePair *IPancakePairSession) Decimals() (uint8, error) {
	return _IPancakePair.Contract.Decimals(&_IPancakePair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IPancakePair *IPancakePairCallerSession) Decimals() (uint8, error) {
	return _IPancakePair.Contract.Decimals(&_IPancakePair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPancakePair *IPancakePairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPancakePair *IPancakePairSession) Factory() (common.Address, error) {
	return _IPancakePair.Contract.Factory(&_IPancakePair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPancakePair *IPancakePairCallerSession) Factory() (common.Address, error) {
	return _IPancakePair.Contract.Factory(&_IPancakePair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IPancakePair *IPancakePairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IPancakePair *IPancakePairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _IPancakePair.Contract.GetReserves(&_IPancakePair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IPancakePair *IPancakePairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _IPancakePair.Contract.GetReserves(&_IPancakePair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_IPancakePair *IPancakePairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_IPancakePair *IPancakePairSession) KLast() (*big.Int, error) {
	return _IPancakePair.Contract.KLast(&_IPancakePair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) KLast() (*big.Int, error) {
	return _IPancakePair.Contract.KLast(&_IPancakePair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IPancakePair *IPancakePairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IPancakePair *IPancakePairSession) Name() (string, error) {
	return _IPancakePair.Contract.Name(&_IPancakePair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IPancakePair *IPancakePairCallerSession) Name() (string, error) {
	return _IPancakePair.Contract.Name(&_IPancakePair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.Nonces(&_IPancakePair.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.Nonces(&_IPancakePair.CallOpts, owner)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairSession) Price0CumulativeLast() (*big.Int, error) {
	return _IPancakePair.Contract.Price0CumulativeLast(&_IPancakePair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _IPancakePair.Contract.Price0CumulativeLast(&_IPancakePair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairSession) Price1CumulativeLast() (*big.Int, error) {
	return _IPancakePair.Contract.Price1CumulativeLast(&_IPancakePair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _IPancakePair.Contract.Price1CumulativeLast(&_IPancakePair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IPancakePair *IPancakePairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IPancakePair *IPancakePairSession) Symbol() (string, error) {
	return _IPancakePair.Contract.Symbol(&_IPancakePair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IPancakePair *IPancakePairCallerSession) Symbol() (string, error) {
	return _IPancakePair.Contract.Symbol(&_IPancakePair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPancakePair *IPancakePairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPancakePair *IPancakePairSession) Token0() (common.Address, error) {
	return _IPancakePair.Contract.Token0(&_IPancakePair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPancakePair *IPancakePairCallerSession) Token0() (common.Address, error) {
	return _IPancakePair.Contract.Token0(&_IPancakePair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPancakePair *IPancakePairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPancakePair *IPancakePairSession) Token1() (common.Address, error) {
	return _IPancakePair.Contract.Token1(&_IPancakePair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPancakePair *IPancakePairCallerSession) Token1() (common.Address, error) {
	return _IPancakePair.Contract.Token1(&_IPancakePair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IPancakePair *IPancakePairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IPancakePair *IPancakePairSession) TotalSupply() (*big.Int, error) {
	return _IPancakePair.Contract.TotalSupply(&_IPancakePair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) TotalSupply() (*big.Int, error) {
	return _IPancakePair.Contract.TotalSupply(&_IPancakePair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.Approve(&_IPancakePair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.Approve(&_IPancakePair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Burn(&_IPancakePair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Burn(&_IPancakePair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IPancakePair *IPancakePairTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "initialize", arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IPancakePair *IPancakePairSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Initialize(&_IPancakePair.TransactOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IPancakePair *IPancakePairTransactorSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Initialize(&_IPancakePair.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPancakePair *IPancakePairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPancakePair *IPancakePairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Mint(&_IPancakePair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPancakePair *IPancakePairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Mint(&_IPancakePair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPancakePair *IPancakePairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPancakePair *IPancakePairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakePair.Contract.Permit(&_IPancakePair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPancakePair *IPancakePairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakePair.Contract.Permit(&_IPancakePair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPancakePair *IPancakePairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPancakePair *IPancakePairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Skim(&_IPancakePair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPancakePair *IPancakePairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Skim(&_IPancakePair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPancakePair *IPancakePairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPancakePair *IPancakePairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPancakePair.Contract.Swap(&_IPancakePair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPancakePair *IPancakePairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPancakePair.Contract.Swap(&_IPancakePair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IPancakePair *IPancakePairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IPancakePair *IPancakePairSession) Sync() (*types.Transaction, error) {
	return _IPancakePair.Contract.Sync(&_IPancakePair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IPancakePair *IPancakePairTransactorSession) Sync() (*types.Transaction, error) {
	return _IPancakePair.Contract.Sync(&_IPancakePair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.Transfer(&_IPancakePair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.Transfer(&_IPancakePair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.TransferFrom(&_IPancakePair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.TransferFrom(&_IPancakePair.TransactOpts, from, to, value)
}

// IPancakePairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IPancakePair contract.
type IPancakePairApprovalIterator struct {
	Event *IPancakePairApproval // Event containing the contract specifics and raw log

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
func (it *IPancakePairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairApproval)
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
		it.Event = new(IPancakePairApproval)
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
func (it *IPancakePairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairApproval represents a Approval event raised by the IPancakePair contract.
type IPancakePairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IPancakePair *IPancakePairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IPancakePairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairApprovalIterator{contract: _IPancakePair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IPancakePair *IPancakePairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IPancakePairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairApproval)
				if err := _IPancakePair.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IPancakePair *IPancakePairFilterer) ParseApproval(log types.Log) (*IPancakePairApproval, error) {
	event := new(IPancakePairApproval)
	if err := _IPancakePair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IPancakePair contract.
type IPancakePairBurnIterator struct {
	Event *IPancakePairBurn // Event containing the contract specifics and raw log

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
func (it *IPancakePairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairBurn)
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
		it.Event = new(IPancakePairBurn)
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
func (it *IPancakePairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairBurn represents a Burn event raised by the IPancakePair contract.
type IPancakePairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_IPancakePair *IPancakePairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IPancakePairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairBurnIterator{contract: _IPancakePair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_IPancakePair *IPancakePairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IPancakePairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairBurn)
				if err := _IPancakePair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_IPancakePair *IPancakePairFilterer) ParseBurn(log types.Log) (*IPancakePairBurn, error) {
	event := new(IPancakePairBurn)
	if err := _IPancakePair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IPancakePair contract.
type IPancakePairMintIterator struct {
	Event *IPancakePairMint // Event containing the contract specifics and raw log

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
func (it *IPancakePairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairMint)
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
		it.Event = new(IPancakePairMint)
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
func (it *IPancakePairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairMint represents a Mint event raised by the IPancakePair contract.
type IPancakePairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*IPancakePairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairMintIterator{contract: _IPancakePair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IPancakePairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairMint)
				if err := _IPancakePair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairFilterer) ParseMint(log types.Log) (*IPancakePairMint, error) {
	event := new(IPancakePairMint)
	if err := _IPancakePair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IPancakePair contract.
type IPancakePairSwapIterator struct {
	Event *IPancakePairSwap // Event containing the contract specifics and raw log

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
func (it *IPancakePairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairSwap)
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
		it.Event = new(IPancakePairSwap)
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
func (it *IPancakePairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairSwap represents a Swap event raised by the IPancakePair contract.
type IPancakePairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_IPancakePair *IPancakePairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IPancakePairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairSwapIterator{contract: _IPancakePair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_IPancakePair *IPancakePairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IPancakePairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairSwap)
				if err := _IPancakePair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_IPancakePair *IPancakePairFilterer) ParseSwap(log types.Log) (*IPancakePairSwap, error) {
	event := new(IPancakePairSwap)
	if err := _IPancakePair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the IPancakePair contract.
type IPancakePairSyncIterator struct {
	Event *IPancakePairSync // Event containing the contract specifics and raw log

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
func (it *IPancakePairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairSync)
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
		it.Event = new(IPancakePairSync)
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
func (it *IPancakePairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairSync represents a Sync event raised by the IPancakePair contract.
type IPancakePairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_IPancakePair *IPancakePairFilterer) FilterSync(opts *bind.FilterOpts) (*IPancakePairSyncIterator, error) {

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &IPancakePairSyncIterator{contract: _IPancakePair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_IPancakePair *IPancakePairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *IPancakePairSync) (event.Subscription, error) {

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairSync)
				if err := _IPancakePair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_IPancakePair *IPancakePairFilterer) ParseSync(log types.Log) (*IPancakePairSync, error) {
	event := new(IPancakePairSync)
	if err := _IPancakePair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IPancakePair contract.
type IPancakePairTransferIterator struct {
	Event *IPancakePairTransfer // Event containing the contract specifics and raw log

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
func (it *IPancakePairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairTransfer)
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
		it.Event = new(IPancakePairTransfer)
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
func (it *IPancakePairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairTransfer represents a Transfer event raised by the IPancakePair contract.
type IPancakePairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IPancakePair *IPancakePairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IPancakePairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairTransferIterator{contract: _IPancakePair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IPancakePair *IPancakePairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IPancakePairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairTransfer)
				if err := _IPancakePair.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IPancakePair *IPancakePairFilterer) ParseTransfer(log types.Log) (*IPancakePairTransfer, error) {
	event := new(IPancakePairTransfer)
	if err := _IPancakePair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakeRouter01ABI is the input ABI used to generate the binding from.
const IPancakeRouter01ABI = "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPancakeRouter01FuncSigs maps the 4-byte function signature to its string representation.
var IPancakeRouter01FuncSigs = map[string]string{
	"ad5c4648": "WETH()",
	"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
	"f305d719": "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
	"c45a0155": "factory()",
	"85f8c259": "getAmountIn(uint256,uint256,uint256)",
	"054d50d4": "getAmountOut(uint256,uint256,uint256)",
	"1f00ca74": "getAmountsIn(uint256,address[])",
	"d06ca61f": "getAmountsOut(uint256,address[])",
	"ad615dec": "quote(uint256,uint256,uint256)",
	"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
	"02751cec": "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
	"ded9382a": "removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"fb3bdb41": "swapETHForExactTokens(uint256,address[],address,uint256)",
	"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
	"18cbafe5": "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
	"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
	"4a25d94a": "swapTokensForExactETH(uint256,uint256,address[],address,uint256)",
	"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
}

// IPancakeRouter01 is an auto generated Go binding around an Ethereum contract.
type IPancakeRouter01 struct {
	IPancakeRouter01Caller     // Read-only binding to the contract
	IPancakeRouter01Transactor // Write-only binding to the contract
	IPancakeRouter01Filterer   // Log filterer for contract events
}

// IPancakeRouter01Caller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeRouter01Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeRouter01Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeRouter01Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeRouter01Session struct {
	Contract     *IPancakeRouter01 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeRouter01CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeRouter01CallerSession struct {
	Contract *IPancakeRouter01Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPancakeRouter01TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeRouter01TransactorSession struct {
	Contract     *IPancakeRouter01Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPancakeRouter01Raw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeRouter01Raw struct {
	Contract *IPancakeRouter01 // Generic contract binding to access the raw methods on
}

// IPancakeRouter01CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeRouter01CallerRaw struct {
	Contract *IPancakeRouter01Caller // Generic read-only contract binding to access the raw methods on
}

// IPancakeRouter01TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeRouter01TransactorRaw struct {
	Contract *IPancakeRouter01Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeRouter01 creates a new instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01(address common.Address, backend bind.ContractBackend) (*IPancakeRouter01, error) {
	contract, err := bindIPancakeRouter01(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01{IPancakeRouter01Caller: IPancakeRouter01Caller{contract: contract}, IPancakeRouter01Transactor: IPancakeRouter01Transactor{contract: contract}, IPancakeRouter01Filterer: IPancakeRouter01Filterer{contract: contract}}, nil
}

// NewIPancakeRouter01Caller creates a new read-only instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Caller(address common.Address, caller bind.ContractCaller) (*IPancakeRouter01Caller, error) {
	contract, err := bindIPancakeRouter01(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Caller{contract: contract}, nil
}

// NewIPancakeRouter01Transactor creates a new write-only instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Transactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeRouter01Transactor, error) {
	contract, err := bindIPancakeRouter01(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Transactor{contract: contract}, nil
}

// NewIPancakeRouter01Filterer creates a new log filterer instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Filterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeRouter01Filterer, error) {
	contract, err := bindIPancakeRouter01(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Filterer{contract: contract}, nil
}

// bindIPancakeRouter01 binds a generic wrapper to an already deployed contract.
func bindIPancakeRouter01(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakeRouter01ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter01.Contract.IPancakeRouter01Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.IPancakeRouter01Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.IPancakeRouter01Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter01 *IPancakeRouter01CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter01.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter01 *IPancakeRouter01TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter01 *IPancakeRouter01TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01Session) WETH() (common.Address, error) {
	return _IPancakeRouter01.Contract.WETH(&_IPancakeRouter01.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) WETH() (common.Address, error) {
	return _IPancakeRouter01.Contract.WETH(&_IPancakeRouter01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01Session) Factory() (common.Address, error) {
	return _IPancakeRouter01.Contract.Factory(&_IPancakeRouter01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) Factory() (common.Address, error) {
	return _IPancakeRouter01.Contract.Factory(&_IPancakeRouter01.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountIn(&_IPancakeRouter01.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountIn(&_IPancakeRouter01.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountOut(&_IPancakeRouter01.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountOut(&_IPancakeRouter01.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsIn(&_IPancakeRouter01.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsIn(&_IPancakeRouter01.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsOut(&_IPancakeRouter01.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsOut(&_IPancakeRouter01.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.Quote(&_IPancakeRouter01.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.Quote(&_IPancakeRouter01.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.AddLiquidity(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.AddLiquidity(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01Session) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.AddLiquidityETH(&_IPancakeRouter01.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.AddLiquidityETH(&_IPancakeRouter01.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidity(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidity(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01Session) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityETH(&_IPancakeRouter01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityETH(&_IPancakeRouter01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01Session) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityETHWithPermit(&_IPancakeRouter01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityETHWithPermit(&_IPancakeRouter01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityWithPermit(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityWithPermit(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapETHForExactTokens(&_IPancakeRouter01.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapETHForExactTokens(&_IPancakeRouter01.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactETHForTokens(&_IPancakeRouter01.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactETHForTokens(&_IPancakeRouter01.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactTokensForETH(&_IPancakeRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactTokensForETH(&_IPancakeRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactTokensForTokens(&_IPancakeRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactTokensForTokens(&_IPancakeRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapTokensForExactETH(&_IPancakeRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapTokensForExactETH(&_IPancakeRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapTokensForExactTokens(&_IPancakeRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapTokensForExactTokens(&_IPancakeRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// IPancakeRouter02ABI is the input ABI used to generate the binding from.
const IPancakeRouter02ABI = "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPancakeRouter02FuncSigs maps the 4-byte function signature to its string representation.
var IPancakeRouter02FuncSigs = map[string]string{
	"ad5c4648": "WETH()",
	"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
	"f305d719": "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
	"c45a0155": "factory()",
	"85f8c259": "getAmountIn(uint256,uint256,uint256)",
	"054d50d4": "getAmountOut(uint256,uint256,uint256)",
	"1f00ca74": "getAmountsIn(uint256,address[])",
	"d06ca61f": "getAmountsOut(uint256,address[])",
	"ad615dec": "quote(uint256,uint256,uint256)",
	"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
	"02751cec": "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
	"af2979eb": "removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)",
	"ded9382a": "removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"5b0d5984": "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"fb3bdb41": "swapETHForExactTokens(uint256,address[],address,uint256)",
	"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
	"b6f9de95": "swapExactETHForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256)",
	"18cbafe5": "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
	"791ac947": "swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
	"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
	"5c11d795": "swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
	"4a25d94a": "swapTokensForExactETH(uint256,uint256,address[],address,uint256)",
	"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
}

// IPancakeRouter02 is an auto generated Go binding around an Ethereum contract.
type IPancakeRouter02 struct {
	IPancakeRouter02Caller     // Read-only binding to the contract
	IPancakeRouter02Transactor // Write-only binding to the contract
	IPancakeRouter02Filterer   // Log filterer for contract events
}

// IPancakeRouter02Caller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeRouter02Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter02Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeRouter02Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter02Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeRouter02Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter02Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeRouter02Session struct {
	Contract     *IPancakeRouter02 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeRouter02CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeRouter02CallerSession struct {
	Contract *IPancakeRouter02Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPancakeRouter02TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeRouter02TransactorSession struct {
	Contract     *IPancakeRouter02Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPancakeRouter02Raw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeRouter02Raw struct {
	Contract *IPancakeRouter02 // Generic contract binding to access the raw methods on
}

// IPancakeRouter02CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeRouter02CallerRaw struct {
	Contract *IPancakeRouter02Caller // Generic read-only contract binding to access the raw methods on
}

// IPancakeRouter02TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeRouter02TransactorRaw struct {
	Contract *IPancakeRouter02Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeRouter02 creates a new instance of IPancakeRouter02, bound to a specific deployed contract.
func NewIPancakeRouter02(address common.Address, backend bind.ContractBackend) (*IPancakeRouter02, error) {
	contract, err := bindIPancakeRouter02(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter02{IPancakeRouter02Caller: IPancakeRouter02Caller{contract: contract}, IPancakeRouter02Transactor: IPancakeRouter02Transactor{contract: contract}, IPancakeRouter02Filterer: IPancakeRouter02Filterer{contract: contract}}, nil
}

// NewIPancakeRouter02Caller creates a new read-only instance of IPancakeRouter02, bound to a specific deployed contract.
func NewIPancakeRouter02Caller(address common.Address, caller bind.ContractCaller) (*IPancakeRouter02Caller, error) {
	contract, err := bindIPancakeRouter02(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter02Caller{contract: contract}, nil
}

// NewIPancakeRouter02Transactor creates a new write-only instance of IPancakeRouter02, bound to a specific deployed contract.
func NewIPancakeRouter02Transactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeRouter02Transactor, error) {
	contract, err := bindIPancakeRouter02(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter02Transactor{contract: contract}, nil
}

// NewIPancakeRouter02Filterer creates a new log filterer instance of IPancakeRouter02, bound to a specific deployed contract.
func NewIPancakeRouter02Filterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeRouter02Filterer, error) {
	contract, err := bindIPancakeRouter02(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter02Filterer{contract: contract}, nil
}

// bindIPancakeRouter02 binds a generic wrapper to an already deployed contract.
func bindIPancakeRouter02(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakeRouter02ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter02 *IPancakeRouter02Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter02.Contract.IPancakeRouter02Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter02 *IPancakeRouter02Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.IPancakeRouter02Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter02 *IPancakeRouter02Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.IPancakeRouter02Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter02 *IPancakeRouter02CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter02.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter02 *IPancakeRouter02TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter02 *IPancakeRouter02TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02Session) WETH() (common.Address, error) {
	return _IPancakeRouter02.Contract.WETH(&_IPancakeRouter02.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) WETH() (common.Address, error) {
	return _IPancakeRouter02.Contract.WETH(&_IPancakeRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02Session) Factory() (common.Address, error) {
	return _IPancakeRouter02.Contract.Factory(&_IPancakeRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) Factory() (common.Address, error) {
	return _IPancakeRouter02.Contract.Factory(&_IPancakeRouter02.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter02 *IPancakeRouter02Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter02 *IPancakeRouter02Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountIn(&_IPancakeRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountIn(&_IPancakeRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter02 *IPancakeRouter02Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter02 *IPancakeRouter02Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountOut(&_IPancakeRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountOut(&_IPancakeRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountsIn(&_IPancakeRouter02.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountsIn(&_IPancakeRouter02.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountsOut(&_IPancakeRouter02.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountsOut(&_IPancakeRouter02.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.Quote(&_IPancakeRouter02.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.Quote(&_IPancakeRouter02.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.AddLiquidity(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.AddLiquidity(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02Session) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.AddLiquidityETH(&_IPancakeRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.AddLiquidityETH(&_IPancakeRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidity(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidity(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETH(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETH(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHWithPermit(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHWithPermit(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityWithPermit(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityWithPermit(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapETHForExactTokens(&_IPancakeRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapETHForExactTokens(&_IPancakeRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactETHForTokens(&_IPancakeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactETHForTokens(&_IPancakeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForETH(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForETH(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapTokensForExactETH(&_IPancakeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapTokensForExactETH(&_IPancakeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapTokensForExactTokens(&_IPancakeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapTokensForExactTokens(&_IPancakeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// IWETHABI is the input ABI used to generate the binding from.
const IWETHABI = "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWETHFuncSigs maps the 4-byte function signature to its string representation.
var IWETHFuncSigs = map[string]string{
	"d0e30db0": "deposit()",
	"a9059cbb": "transfer(address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// IWETH is an auto generated Go binding around an Ethereum contract.
type IWETH struct {
	IWETHCaller     // Read-only binding to the contract
	IWETHTransactor // Write-only binding to the contract
	IWETHFilterer   // Log filterer for contract events
}

// IWETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETHSession struct {
	Contract     *IWETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETHCallerSession struct {
	Contract *IWETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETHTransactorSession struct {
	Contract     *IWETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWETHRaw struct {
	Contract *IWETH // Generic contract binding to access the raw methods on
}

// IWETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETHCallerRaw struct {
	Contract *IWETHCaller // Generic read-only contract binding to access the raw methods on
}

// IWETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETHTransactorRaw struct {
	Contract *IWETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH creates a new instance of IWETH, bound to a specific deployed contract.
func NewIWETH(address common.Address, backend bind.ContractBackend) (*IWETH, error) {
	contract, err := bindIWETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH{IWETHCaller: IWETHCaller{contract: contract}, IWETHTransactor: IWETHTransactor{contract: contract}, IWETHFilterer: IWETHFilterer{contract: contract}}, nil
}

// NewIWETHCaller creates a new read-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHCaller(address common.Address, caller bind.ContractCaller) (*IWETHCaller, error) {
	contract, err := bindIWETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHCaller{contract: contract}, nil
}

// NewIWETHTransactor creates a new write-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*IWETHTransactor, error) {
	contract, err := bindIWETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHTransactor{contract: contract}, nil
}

// NewIWETHFilterer creates a new log filterer instance of IWETH, bound to a specific deployed contract.
func NewIWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*IWETHFilterer, error) {
	contract, err := bindIWETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETHFilterer{contract: contract}, nil
}

// bindIWETH binds a generic wrapper to an already deployed contract.
func bindIWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.IWETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}

// PancakeLibraryABI is the input ABI used to generate the binding from.
const PancakeLibraryABI = "[]"

// PancakeLibraryBin is the compiled bytecode used for deploying new contracts.
var PancakeLibraryBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202e29f17138b1d353873f45111d23748df4fc4c110619ae9b5097c61a251e7d9b64736f6c634300060c0033"

// DeployPancakeLibrary deploys a new Ethereum contract, binding an instance of PancakeLibrary to it.
func DeployPancakeLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PancakeLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PancakeLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PancakeLibrary{PancakeLibraryCaller: PancakeLibraryCaller{contract: contract}, PancakeLibraryTransactor: PancakeLibraryTransactor{contract: contract}, PancakeLibraryFilterer: PancakeLibraryFilterer{contract: contract}}, nil
}

// PancakeLibrary is an auto generated Go binding around an Ethereum contract.
type PancakeLibrary struct {
	PancakeLibraryCaller     // Read-only binding to the contract
	PancakeLibraryTransactor // Write-only binding to the contract
	PancakeLibraryFilterer   // Log filterer for contract events
}

// PancakeLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeLibrarySession struct {
	Contract     *PancakeLibrary   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeLibraryCallerSession struct {
	Contract *PancakeLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PancakeLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeLibraryTransactorSession struct {
	Contract     *PancakeLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PancakeLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeLibraryRaw struct {
	Contract *PancakeLibrary // Generic contract binding to access the raw methods on
}

// PancakeLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeLibraryCallerRaw struct {
	Contract *PancakeLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeLibraryTransactorRaw struct {
	Contract *PancakeLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeLibrary creates a new instance of PancakeLibrary, bound to a specific deployed contract.
func NewPancakeLibrary(address common.Address, backend bind.ContractBackend) (*PancakeLibrary, error) {
	contract, err := bindPancakeLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeLibrary{PancakeLibraryCaller: PancakeLibraryCaller{contract: contract}, PancakeLibraryTransactor: PancakeLibraryTransactor{contract: contract}, PancakeLibraryFilterer: PancakeLibraryFilterer{contract: contract}}, nil
}

// NewPancakeLibraryCaller creates a new read-only instance of PancakeLibrary, bound to a specific deployed contract.
func NewPancakeLibraryCaller(address common.Address, caller bind.ContractCaller) (*PancakeLibraryCaller, error) {
	contract, err := bindPancakeLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeLibraryCaller{contract: contract}, nil
}

// NewPancakeLibraryTransactor creates a new write-only instance of PancakeLibrary, bound to a specific deployed contract.
func NewPancakeLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeLibraryTransactor, error) {
	contract, err := bindPancakeLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeLibraryTransactor{contract: contract}, nil
}

// NewPancakeLibraryFilterer creates a new log filterer instance of PancakeLibrary, bound to a specific deployed contract.
func NewPancakeLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeLibraryFilterer, error) {
	contract, err := bindPancakeLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeLibraryFilterer{contract: contract}, nil
}

// bindPancakeLibrary binds a generic wrapper to an already deployed contract.
func bindPancakeLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeLibrary *PancakeLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeLibrary.Contract.PancakeLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeLibrary *PancakeLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeLibrary.Contract.PancakeLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeLibrary *PancakeLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeLibrary.Contract.PancakeLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeLibrary *PancakeLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeLibrary *PancakeLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeLibrary *PancakeLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeLibrary.Contract.contract.Transact(opts, method, params...)
}

// PancakeRouterABI is the input ABI used to generate the binding from.
const PancakeRouterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// PancakeRouterFuncSigs maps the 4-byte function signature to its string representation.
var PancakeRouterFuncSigs = map[string]string{
	"ad5c4648": "WETH()",
	"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
	"f305d719": "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
	"c45a0155": "factory()",
	"85f8c259": "getAmountIn(uint256,uint256,uint256)",
	"054d50d4": "getAmountOut(uint256,uint256,uint256)",
	"1f00ca74": "getAmountsIn(uint256,address[])",
	"d06ca61f": "getAmountsOut(uint256,address[])",
	"ad615dec": "quote(uint256,uint256,uint256)",
	"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
	"02751cec": "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
	"af2979eb": "removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)",
	"ded9382a": "removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"5b0d5984": "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"fb3bdb41": "swapETHForExactTokens(uint256,address[],address,uint256)",
	"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
	"b6f9de95": "swapExactETHForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256)",
	"18cbafe5": "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
	"791ac947": "swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
	"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
	"5c11d795": "swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
	"4a25d94a": "swapTokensForExactETH(uint256,uint256,address[],address,uint256)",
	"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
}

// PancakeRouterBin is the compiled bytecode used for deploying new contracts.
var PancakeRouterBin = "0x60c060405234801561001057600080fd5b50604051620046e0380380620046e08339818101604052604081101561003557600080fd5b5080516020909101516001600160601b0319606092831b8116608052911b1660a05260805160601c60a05160601c61455b620001856000398061015f5280610ce45280610d1f5280610e16528061103452806113be528061152452806118d352806119ae5280611a885280611b565280611c9c5280611d245280611f605280611fdb528061208f528061215b52806121f05280612264528061274752806129b15280612a075280612a3b5280612aaf5280612c3d5280612d805280612e08525080610ea45280610f7b52806110fa5280611133528061126e528061144c528061150252806116725280611be95280611d565280611eb0528061229652806124d452806126cc52806126f55280612725528061289252806129e55280612cd05280612e3a52806136bf528061370252806139dc5280613b525280613f435280613ffc52806140af525061455b6000f3fe60806040526004361061014f5760003560e01c80638803dbee116100b6578063c45a01551161006f578063c45a015514610a10578063d06ca61f14610a25578063ded9382a14610ada578063e8e3370014610b4d578063f305d71914610bcd578063fb3bdb4114610c1357610188565b80638803dbee146107df578063ad5c464814610875578063ad615dec146108a6578063af2979eb146108dc578063b6f9de951461092f578063baa2abde146109b357610188565b80634a25d94a116101085780634a25d94a146104f05780635b0d5984146105865780635c11d795146105f9578063791ac9471461068f5780637ff36ab51461072557806385f8c259146107a957610188565b806302751cec1461018d578063054d50d4146101f957806318cbafe5146102415780631f00ca74146103275780632195995c146103dc57806338ed17391461045a57610188565b3661018857336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461018657fe5b005b600080fd5b34801561019957600080fd5b506101e0600480360360c08110156101b057600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a00135610c97565b6040805192835260208301919091528051918290030190f35b34801561020557600080fd5b5061022f6004803603606081101561021c57600080fd5b5080359060208101359060400135610db1565b60408051918252519081900360200190f35b34801561024d57600080fd5b506102d7600480360360a081101561026457600080fd5b813591602081013591810190606081016040820135600160201b81111561028a57600080fd5b82018360208201111561029c57600080fd5b803590602001918460208302840111600160201b831117156102bd57600080fd5b91935091506001600160a01b038135169060200135610dc6565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103135781810151838201526020016102fb565b505050509050019250505060405180910390f35b34801561033357600080fd5b506102d76004803603604081101561034a57600080fd5b81359190810190604081016020820135600160201b81111561036b57600080fd5b82018360208201111561037d57600080fd5b803590602001918460208302840111600160201b8311171561039e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506110f3945050505050565b3480156103e857600080fd5b506101e0600480360361016081101561040057600080fd5b506001600160a01b038135811691602081013582169160408201359160608101359160808201359160a08101359091169060c08101359060e081013515159060ff6101008201351690610120810135906101400135611129565b34801561046657600080fd5b506102d7600480360360a081101561047d57600080fd5b813591602081013591810190606081016040820135600160201b8111156104a357600080fd5b8201836020820111156104b557600080fd5b803590602001918460208302840111600160201b831117156104d657600080fd5b91935091506001600160a01b038135169060200135611223565b3480156104fc57600080fd5b506102d7600480360360a081101561051357600080fd5b813591602081013591810190606081016040820135600160201b81111561053957600080fd5b82018360208201111561054b57600080fd5b803590602001918460208302840111600160201b8311171561056c57600080fd5b91935091506001600160a01b03813516906020013561136e565b34801561059257600080fd5b5061022f60048036036101408110156105aa57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a08101359060c081013515159060ff60e082013516906101008101359061012001356114fa565b34801561060557600080fd5b50610186600480360360a081101561061c57600080fd5b813591602081013591810190606081016040820135600160201b81111561064257600080fd5b82018360208201111561065457600080fd5b803590602001918460208302840111600160201b8311171561067557600080fd5b91935091506001600160a01b038135169060200135611608565b34801561069b57600080fd5b50610186600480360360a08110156106b257600080fd5b813591602081013591810190606081016040820135600160201b8111156106d857600080fd5b8201836020820111156106ea57600080fd5b803590602001918460208302840111600160201b8311171561070b57600080fd5b91935091506001600160a01b038135169060200135611885565b6102d76004803603608081101561073b57600080fd5b81359190810190604081016020820135600160201b81111561075c57600080fd5b82018360208201111561076e57600080fd5b803590602001918460208302840111600160201b8311171561078f57600080fd5b91935091506001600160a01b038135169060200135611b0e565b3480156107b557600080fd5b5061022f600480360360608110156107cc57600080fd5b5080359060208101359060400135611e58565b3480156107eb57600080fd5b506102d7600480360360a081101561080257600080fd5b813591602081013591810190606081016040820135600160201b81111561082857600080fd5b82018360208201111561083a57600080fd5b803590602001918460208302840111600160201b8311171561085b57600080fd5b91935091506001600160a01b038135169060200135611e65565b34801561088157600080fd5b5061088a611f5e565b604080516001600160a01b039092168252519081900360200190f35b3480156108b257600080fd5b5061022f600480360360608110156108c957600080fd5b5080359060208101359060400135611f82565b3480156108e857600080fd5b5061022f600480360360c08110156108ff57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a00135611f8f565b6101866004803603608081101561094557600080fd5b81359190810190604081016020820135600160201b81111561096657600080fd5b82018360208201111561097857600080fd5b803590602001918460208302840111600160201b8311171561099957600080fd5b91935091506001600160a01b038135169060200135612115565b3480156109bf57600080fd5b506101e0600480360360e08110156109d657600080fd5b506001600160a01b038135811691602081013582169160408201359160608101359160808201359160a08101359091169060c00135612486565b348015610a1c57600080fd5b5061088a6126ca565b348015610a3157600080fd5b506102d760048036036040811015610a4857600080fd5b81359190810190604081016020820135600160201b811115610a6957600080fd5b820183602082011115610a7b57600080fd5b803590602001918460208302840111600160201b83111715610a9c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506126ee945050505050565b348015610ae657600080fd5b506101e06004803603610140811015610afe57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a08101359060c081013515159060ff60e0820135169061010081013590610120013561271b565b348015610b5957600080fd5b50610baf6004803603610100811015610b7157600080fd5b506001600160a01b038135811691602081013582169160408201359160608101359160808201359160a08101359160c0820135169060e0013561282f565b60408051938452602084019290925282820152519081900360600190f35b610baf600480360360c0811015610be357600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a00135612962565b6102d760048036036080811015610c2957600080fd5b81359190810190604081016020820135600160201b811115610c4a57600080fd5b820183602082011115610c5c57600080fd5b803590602001918460208302840111600160201b83111715610c7d57600080fd5b91935091506001600160a01b038135169060200135612bf5565b6000808242811015610cde576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b610d0d897f00000000000000000000000000000000000000000000000000000000000000008a8a8a308a612486565b9093509150610d1d898685612f6e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d836040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015610d8357600080fd5b505af1158015610d97573d6000803e3d6000fd5b50505050610da585836130d8565b50965096945050505050565b6000610dbe8484846131d0565b949350505050565b60608142811015610e0c576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001686866000198101818110610e4657fe5b905060200201356001600160a01b03166001600160a01b031614610e9f576040805162461bcd60e51b815260206004820152601b6024820152600080516020614409833981519152604482015290519081900360640190fd5b610efd7f0000000000000000000000000000000000000000000000000000000000000000898888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506132a892505050565b91508682600184510381518110610f1057fe5b60200260200101511015610f555760405162461bcd60e51b81526004018080602001828103825260298152602001806143bd6029913960400191505060405180910390fd5b610ff386866000818110610f6557fe5b905060200201356001600160a01b031633610fd97f00000000000000000000000000000000000000000000000000000000000000008a8a6000818110610fa757fe5b905060200201356001600160a01b03168b8b6001818110610fc457fe5b905060200201356001600160a01b03166133f4565b85600081518110610fe657fe5b60200260200101516134b3565b61103282878780806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250309250613610915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d8360018551038151811061107157fe5b60200260200101516040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156110af57600080fd5b505af11580156110c3573d6000803e3d6000fd5b505050506110e884836001855103815181106110db57fe5b60200260200101516130d8565b509695505050505050565b60606111207f0000000000000000000000000000000000000000000000000000000000000000848461384d565b90505b92915050565b60008060006111597f00000000000000000000000000000000000000000000000000000000000000008f8f6133f4565b9050600087611168578c61116c565b6000195b6040805163d505accf60e01b815233600482015230602482015260448101839052606481018c905260ff8a16608482015260a4810189905260c4810188905290519192506001600160a01b0384169163d505accf9160e48082019260009290919082900301818387803b1580156111e257600080fd5b505af11580156111f6573d6000803e3d6000fd5b505050506112098f8f8f8f8f8f8f612486565b809450819550505050509b509b9950505050505050505050565b60608142811015611269576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b6112c77f0000000000000000000000000000000000000000000000000000000000000000898888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506132a892505050565b915086826001845103815181106112da57fe5b6020026020010151101561131f5760405162461bcd60e51b81526004018080602001828103825260298152602001806143bd6029913960400191505060405180910390fd5b61132f86866000818110610f6557fe5b6110e882878780806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250899250613610915050565b606081428110156113b4576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016868660001981018181106113ee57fe5b905060200201356001600160a01b03166001600160a01b031614611447576040805162461bcd60e51b815260206004820152601b6024820152600080516020614409833981519152604482015290519081900360640190fd5b6114a57f00000000000000000000000000000000000000000000000000000000000000008988888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061384d92505050565b915086826000815181106114b557fe5b60200260200101511115610f555760405162461bcd60e51b81526004018080602001828103825260258152602001806144936025913960400191505060405180910390fd5b6000806115487f00000000000000000000000000000000000000000000000000000000000000008d7f00000000000000000000000000000000000000000000000000000000000000006133f4565b9050600086611557578b61155b565b6000195b6040805163d505accf60e01b815233600482015230602482015260448101839052606481018b905260ff8916608482015260a4810188905260c4810187905290519192506001600160a01b0384169163d505accf9160e48082019260009290919082900301818387803b1580156115d157600080fd5b505af11580156115e5573d6000803e3d6000fd5b505050506115f78d8d8d8d8d8d611f8f565b9d9c50505050505050505050505050565b804281101561164c576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b6116c18585600081811061165c57fe5b905060200201356001600160a01b0316336116bb7f00000000000000000000000000000000000000000000000000000000000000008989600081811061169e57fe5b905060200201356001600160a01b03168a8a6001818110610fc457fe5b8a6134b3565b6000858560001981018181106116d357fe5b905060200201356001600160a01b03166001600160a01b03166370a08231856040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561172f57600080fd5b505afa158015611743573d6000803e3d6000fd5b505050506040513d602081101561175957600080fd5b5051604080516020888102828101820190935288825292935061179b929091899189918291850190849080828437600092019190915250889250613985915050565b8661183e82888860001981018181106117b057fe5b905060200201356001600160a01b03166001600160a01b03166370a08231886040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561180c57600080fd5b505afa158015611820573d6000803e3d6000fd5b505050506040513d602081101561183657600080fd5b505190613c87565b101561187b5760405162461bcd60e51b81526004018080602001828103825260298152602001806143bd6029913960400191505060405180910390fd5b5050505050505050565b80428110156118c9576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168585600019810181811061190357fe5b905060200201356001600160a01b03166001600160a01b03161461195c576040805162461bcd60e51b815260206004820152601b6024820152600080516020614409833981519152604482015290519081900360640190fd5b61196c8585600081811061165c57fe5b6119aa858580806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250309250613985915050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611a1957600080fd5b505afa158015611a2d573d6000803e3d6000fd5b505050506040513d6020811015611a4357600080fd5b5051905086811015611a865760405162461bcd60e51b81526004018080602001828103825260298152602001806143bd6029913960400191505060405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015611aec57600080fd5b505af1158015611b00573d6000803e3d6000fd5b5050505061187b84826130d8565b60608142811015611b54576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031686866000818110611b8b57fe5b905060200201356001600160a01b03166001600160a01b031614611be4576040805162461bcd60e51b815260206004820152601b6024820152600080516020614409833981519152604482015290519081900360640190fd5b611c427f0000000000000000000000000000000000000000000000000000000000000000348888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506132a892505050565b91508682600184510381518110611c5557fe5b60200260200101511015611c9a5760405162461bcd60e51b81526004018080602001828103825260298152602001806143bd6029913960400191505060405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db083600081518110611cd657fe5b60200260200101516040518263ffffffff1660e01b81526004016000604051808303818588803b158015611d0957600080fd5b505af1158015611d1d573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb611d827f00000000000000000000000000000000000000000000000000000000000000008989600081811061169e57fe5b84600081518110611d8f57fe5b60200260200101516040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015611ddd57600080fd5b505af1158015611df1573d6000803e3d6000fd5b505050506040513d6020811015611e0757600080fd5b5051611e0f57fe5b611e4e82878780806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250899250613610915050565b5095945050505050565b6000610dbe848484613cd7565b60608142811015611eab576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b611f097f00000000000000000000000000000000000000000000000000000000000000008988888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061384d92505050565b91508682600081518110611f1957fe5b6020026020010151111561131f5760405162461bcd60e51b81526004018080602001828103825260258152602001806144936025913960400191505060405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000081565b6000610dbe848484613daf565b60008142811015611fd5576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b612004887f00000000000000000000000000000000000000000000000000000000000000008989893089612486565b90508092505061208d88858a6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561205c57600080fd5b505afa158015612070573d6000803e3d6000fd5b505050506040513d602081101561208657600080fd5b5051612f6e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d836040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156120f357600080fd5b505af1158015612107573d6000803e3d6000fd5b505050506110e884836130d8565b8042811015612159576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168585600081811061219057fe5b905060200201356001600160a01b03166001600160a01b0316146121e9576040805162461bcd60e51b815260206004820152601b6024820152600080516020614409833981519152604482015290519081900360640190fd5b60003490507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561224957600080fd5b505af115801561225d573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb6122c27f00000000000000000000000000000000000000000000000000000000000000008989600081811061169e57fe5b836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561230957600080fd5b505af115801561231d573d6000803e3d6000fd5b505050506040513d602081101561233357600080fd5b505161233b57fe5b60008686600019810181811061234d57fe5b905060200201356001600160a01b03166001600160a01b03166370a08231866040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156123a957600080fd5b505afa1580156123bd573d6000803e3d6000fd5b505050506040513d60208110156123d357600080fd5b505160408051602089810282810182019093528982529293506124159290918a918a918291850190849080828437600092019190915250899250613985915050565b8761183e828989600019810181811061242a57fe5b905060200201356001600160a01b03166001600160a01b03166370a08231896040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561180c57600080fd5b60008082428110156124cd576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b60006124fa7f00000000000000000000000000000000000000000000000000000000000000008c8c6133f4565b604080516323b872dd60e01b81523360048201526001600160a01b03831660248201819052604482018d9052915192935090916323b872dd916064808201926020929091908290030181600087803b15801561255557600080fd5b505af1158015612569573d6000803e3d6000fd5b505050506040513d602081101561257f57600080fd5b50506040805163226bf2d160e21b81526001600160a01b03888116600483015282516000938493928616926389afcb44926024808301939282900301818787803b1580156125cc57600080fd5b505af11580156125e0573d6000803e3d6000fd5b505050506040513d60408110156125f657600080fd5b508051602090910151909250905060006126108e8e613e55565b509050806001600160a01b03168e6001600160a01b031614612633578183612636565b82825b90975095508a87101561267a5760405162461bcd60e51b81526004018080602001828103825260248152602001806144296024913960400191505060405180910390fd5b898610156126b95760405162461bcd60e51b81526004018080602001828103825260248152602001806145026024913960400191505060405180910390fd5b505050505097509795505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60606111207f000000000000000000000000000000000000000000000000000000000000000084846132a8565b600080600061276b7f00000000000000000000000000000000000000000000000000000000000000008e7f00000000000000000000000000000000000000000000000000000000000000006133f4565b905060008761277a578c61277e565b6000195b6040805163d505accf60e01b815233600482015230602482015260448101839052606481018c905260ff8a16608482015260a4810189905260c4810188905290519192506001600160a01b0384169163d505accf9160e48082019260009290919082900301818387803b1580156127f457600080fd5b505af1158015612808573d6000803e3d6000fd5b5050505061281a8e8e8e8e8e8e610c97565b909f909e509c50505050505050505050505050565b60008060008342811015612878576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b6128868c8c8c8c8c8c613f33565b909450925060006128b87f00000000000000000000000000000000000000000000000000000000000000008e8e6133f4565b90506128c68d3383886134b3565b6128d28c3383876134b3565b806001600160a01b0316636a627842886040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050602060405180830381600087803b15801561292157600080fd5b505af1158015612935573d6000803e3d6000fd5b505050506040513d602081101561294b57600080fd5b5051949d939c50939a509198505050505050505050565b600080600083428110156129ab576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b6129d98a7f00000000000000000000000000000000000000000000000000000000000000008b348c8c613f33565b90945092506000612a2b7f00000000000000000000000000000000000000000000000000000000000000008c7f00000000000000000000000000000000000000000000000000000000000000006133f4565b9050612a398b3383886134b3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b158015612a9457600080fd5b505af1158015612aa8573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb82866040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015612b2457600080fd5b505af1158015612b38573d6000803e3d6000fd5b505050506040513d6020811015612b4e57600080fd5b5051612b5657fe5b806001600160a01b0316636a627842886040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050602060405180830381600087803b158015612ba557600080fd5b505af1158015612bb9573d6000803e3d6000fd5b505050506040513d6020811015612bcf57600080fd5b5051925034841015612be757612be7338534036130d8565b505096509650969350505050565b60608142811015612c3b576040805162461bcd60e51b8152602060048201526016602482015260008051602061439d833981519152604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031686866000818110612c7257fe5b905060200201356001600160a01b03166001600160a01b031614612ccb576040805162461bcd60e51b815260206004820152601b6024820152600080516020614409833981519152604482015290519081900360640190fd5b612d297f00000000000000000000000000000000000000000000000000000000000000008888888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061384d92505050565b91503482600081518110612d3957fe5b60200260200101511115612d7e5760405162461bcd60e51b81526004018080602001828103825260258152602001806144936025913960400191505060405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db083600081518110612dba57fe5b60200260200101516040518263ffffffff1660e01b81526004016000604051808303818588803b158015612ded57600080fd5b505af1158015612e01573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb612e667f00000000000000000000000000000000000000000000000000000000000000008989600081811061169e57fe5b84600081518110612e7357fe5b60200260200101516040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015612ec157600080fd5b505af1158015612ed5573d6000803e3d6000fd5b505050506040513d6020811015612eeb57600080fd5b5051612ef357fe5b612f3282878780806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250899250613610915050565b81600081518110612f3f57fe5b6020026020010151341115611e4e57611e4e3383600081518110612f5f57fe5b602002602001015134036130d8565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b178152925182516000946060949389169392918291908083835b60208310612feb5780518252601f199092019160209182019101612fcc565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461304d576040519150601f19603f3d011682016040523d82523d6000602084013e613052565b606091505b5091509150818015613080575080511580613080575080806020019051602081101561307d57600080fd5b50515b6130d1576040805162461bcd60e51b815260206004820152601f60248201527f5472616e7366657248656c7065723a205452414e534645525f4641494c454400604482015290519081900360640190fd5b5050505050565b604080516000808252602082019092526001600160a01b0384169083906040518082805190602001908083835b602083106131245780518252601f199092019160209182019101613105565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114613186576040519150601f19603f3d011682016040523d82523d6000602084013e61318b565b606091505b50509050806131cb5760405162461bcd60e51b81526004018080602001828103825260238152602001806144706023913960400191505060405180910390fd5b505050565b60008084116132105760405162461bcd60e51b815260040180806020018281038252602981526020018061434a6029913960400191505060405180910390fd5b6000831180156132205750600082115b61325b5760405162461bcd60e51b81526004018080602001828103825260268152602001806144b86026913960400191505060405180910390fd5b6000613269856126f76141c4565b9050600061327782856141c4565b905060006132918361328b886127106141c4565b90614227565b905080828161329c57fe5b04979650505050505050565b6060600282511015613301576040805162461bcd60e51b815260206004820152601c60248201527f50616e63616b654c6962726172793a20494e56414c49445f5041544800000000604482015290519081900360640190fd5b815167ffffffffffffffff8111801561331957600080fd5b50604051908082528060200260200182016040528015613343578160200160208202803683370190505b509050828160008151811061335457fe5b60200260200101818152505060005b60018351038110156133ec576000806133a68786858151811061338257fe5b602002602001015187866001018151811061339957fe5b6020026020010151614276565b915091506133c88484815181106133b957fe5b602002602001015183836131d0565b8484600101815181106133d757fe5b60209081029190910101525050600101613363565b509392505050565b60008060006134038585613e55565b604080516bffffffffffffffffffffffff19606094851b811660208084019190915293851b81166034830152825160288184030181526048830184528051908501206001600160f81b031960688401529a90941b9093166069840152607d8301989098527efb7f630766e6a796048ea87d01acd3068e8ff67d078148a3fa3f4a84f69bd5609d808401919091528851808403909101815260bd909201909752805196019590952095945050505050565b604080516001600160a01b0385811660248301528481166044830152606480830185905283518084039091018152608490920183526020820180516001600160e01b03166323b872dd60e01b17815292518251600094606094938a169392918291908083835b602083106135385780518252601f199092019160209182019101613519565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461359a576040519150601f19603f3d011682016040523d82523d6000602084013e61359f565b606091505b50915091508180156135cd5750805115806135cd57508080602001905160208110156135ca57600080fd5b50515b6136085760405162461bcd60e51b81526004018080602001828103825260248152602001806144de6024913960400191505060405180910390fd5b505050505050565b60005b60018351038110156138475760008084838151811061362e57fe5b602002602001015185846001018151811061364557fe5b602002602001015191509150600061365d8383613e55565b509050600087856001018151811061367157fe5b60200260200101519050600080836001600160a01b0316866001600160a01b03161461369f578260006136a3565b6000835b91509150600060028a510388106136ba57886136fb565b6136fb7f0000000000000000000000000000000000000000000000000000000000000000878c8b600201815181106136ee57fe5b60200260200101516133f4565b90506137287f000000000000000000000000000000000000000000000000000000000000000088886133f4565b6001600160a01b031663022c0d9f84848460006040519080825280601f01601f191660200182016040528015613765576020820181803683370190505b506040518563ffffffff1660e01b815260040180858152602001848152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156137cd5781810151838201526020016137b5565b50505050905090810190601f1680156137fa5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561381c57600080fd5b505af1158015613830573d6000803e3d6000fd5b505060019099019850613613975050505050505050565b50505050565b60606002825110156138a6576040805162461bcd60e51b815260206004820152601c60248201527f50616e63616b654c6962726172793a20494e56414c49445f5041544800000000604482015290519081900360640190fd5b815167ffffffffffffffff811180156138be57600080fd5b506040519080825280602002602001820160405280156138e8578160200160208202803683370190505b50905082816001835103815181106138fc57fe5b60209081029190910101528151600019015b80156133ec5760008061393e8786600186038151811061392a57fe5b602002602001015187868151811061339957fe5b9150915061396084848151811061395157fe5b60200260200101518383613cd7565b84600185038151811061396f57fe5b602090810291909101015250506000190161390e565b60005b60018351038110156131cb576000808483815181106139a357fe5b60200260200101518584600101815181106139ba57fe5b60200260200101519150915060006139d28383613e55565b5090506000613a027f000000000000000000000000000000000000000000000000000000000000000085856133f4565b9050600080600080846001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015613a4357600080fd5b505afa158015613a57573d6000803e3d6000fd5b505050506040513d6060811015613a6d57600080fd5b5080516020909101516001600160701b0391821693501690506000806001600160a01b038a811690891614613aa3578284613aa6565b83835b91509150613afb828b6001600160a01b03166370a082318a6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561180c57600080fd5b9550613b088683836131d0565b945050505050600080856001600160a01b0316886001600160a01b031614613b3257826000613b36565b6000835b91509150600060028c51038a10613b4d578a613b81565b613b817f0000000000000000000000000000000000000000000000000000000000000000898e8d600201815181106136ee57fe5b604080516000808252602082019283905263022c0d9f60e01b835260248201878152604483018790526001600160a01b038086166064850152608060848501908152845160a48601819052969750908c169563022c0d9f958a958a958a9591949193919260c486019290918190849084905b83811015613c0b578181015183820152602001613bf3565b50505050905090810190601f168015613c385780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015613c5a57600080fd5b505af1158015613c6e573d6000803e3d6000fd5b50506001909b019a506139889950505050505050505050565b80820382811115611123576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6000808411613d175760405162461bcd60e51b815260040180806020018281038252602a815260200180614373602a913960400191505060405180910390fd5b600083118015613d275750600082115b613d625760405162461bcd60e51b81526004018080602001828103825260268152602001806144b86026913960400191505060405180910390fd5b6000613d7a612710613d7486886141c4565b906141c4565b90506000613d8e6126f7613d748689613c87565b9050613da56001828481613d9e57fe5b0490614227565b9695505050505050565b6000808411613def5760405162461bcd60e51b815260040180806020018281038252602381526020018061444d6023913960400191505060405180910390fd5b600083118015613dff5750600082115b613e3a5760405162461bcd60e51b81526004018080602001828103825260268152602001806144b86026913960400191505060405180910390fd5b82613e4585846141c4565b81613e4c57fe5b04949350505050565b600080826001600160a01b0316846001600160a01b03161415613ea95760405162461bcd60e51b81526004018080602001828103825260238152602001806143e66023913960400191505060405180910390fd5b826001600160a01b0316846001600160a01b031610613ec9578284613ecc565b83835b90925090506001600160a01b038216613f2c576040805162461bcd60e51b815260206004820152601c60248201527f50616e63616b654c6962726172793a205a45524f5f4144445245535300000000604482015290519081900360640190fd5b9250929050565b60008060006001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e6a439058a8a6040518363ffffffff1660e01b815260040180836001600160a01b03168152602001826001600160a01b031681526020019250505060206040518083038186803b158015613fbf57600080fd5b505afa158015613fd3573d6000803e3d6000fd5b505050506040513d6020811015613fe957600080fd5b50516001600160a01b031614156140a7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c9c6539689896040518363ffffffff1660e01b815260040180836001600160a01b03168152602001826001600160a01b0316815260200192505050602060405180830381600087803b15801561407a57600080fd5b505af115801561408e573d6000803e3d6000fd5b505050506040513d60208110156140a457600080fd5b50505b6000806140d57f00000000000000000000000000000000000000000000000000000000000000008b8b614276565b915091508160001480156140e7575080155b156140f7578793508692506141b7565b6000614104898484613daf565b9050878111614157578581101561414c5760405162461bcd60e51b81526004018080602001828103825260248152602001806145026024913960400191505060405180910390fd5b8894509250826141b5565b6000614164898486613daf565b90508981111561417057fe5b878110156141af5760405162461bcd60e51b81526004018080602001828103825260248152602001806144296024913960400191505060405180910390fd5b94508793505b505b5050965096945050505050565b60008115806141df575050808202828282816141dc57fe5b04145b611123576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b80820182811015611123576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fd5b60008060006142858585613e55565b5090506142938686866133f4565b506000806142a28888886133f4565b6001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b1580156142da57600080fd5b505afa1580156142ee573d6000803e3d6000fd5b505050506040513d606081101561430457600080fd5b5080516020909101516001600160701b0391821693501690506001600160a01b038781169084161461433757808261433a565b81815b9099909850965050505050505056fe50616e63616b654c6962726172793a20494e53554646494349454e545f494e5055545f414d4f554e5450616e63616b654c6962726172793a20494e53554646494349454e545f4f55545055545f414d4f554e5450616e63616b65526f757465723a20455850495245440000000000000000000050616e63616b65526f757465723a20494e53554646494349454e545f4f55545055545f414d4f554e5450616e63616b654c6962726172793a204944454e544943414c5f41444452455353455350616e63616b65526f757465723a20494e56414c49445f50415448000000000050616e63616b65526f757465723a20494e53554646494349454e545f415f414d4f554e5450616e63616b654c6962726172793a20494e53554646494349454e545f414d4f554e545472616e7366657248656c7065723a204554485f5452414e534645525f4641494c454450616e63616b65526f757465723a204558434553534956455f494e5055545f414d4f554e5450616e63616b654c6962726172793a20494e53554646494349454e545f4c49515549444954595472616e7366657248656c7065723a205452414e534645525f46524f4d5f4641494c454450616e63616b65526f757465723a20494e53554646494349454e545f425f414d4f554e54a26469706673582212202fa1cef8ff40cf3d090e251b105101ff3139841fba2033131fd521ba7f6eef0a64736f6c634300060c0033"

// DeployPancakeRouter deploys a new Ethereum contract, binding an instance of PancakeRouter to it.
func DeployPancakeRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _WETH common.Address) (common.Address, *types.Transaction, *PancakeRouter, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeRouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PancakeRouterBin), backend, _factory, _WETH)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PancakeRouter{PancakeRouterCaller: PancakeRouterCaller{contract: contract}, PancakeRouterTransactor: PancakeRouterTransactor{contract: contract}, PancakeRouterFilterer: PancakeRouterFilterer{contract: contract}}, nil
}

// PancakeRouter is an auto generated Go binding around an Ethereum contract.
type PancakeRouter struct {
	PancakeRouterCaller     // Read-only binding to the contract
	PancakeRouterTransactor // Write-only binding to the contract
	PancakeRouterFilterer   // Log filterer for contract events
}

// PancakeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeRouterSession struct {
	Contract     *PancakeRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeRouterCallerSession struct {
	Contract *PancakeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PancakeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeRouterTransactorSession struct {
	Contract     *PancakeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PancakeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeRouterRaw struct {
	Contract *PancakeRouter // Generic contract binding to access the raw methods on
}

// PancakeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeRouterCallerRaw struct {
	Contract *PancakeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeRouterTransactorRaw struct {
	Contract *PancakeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeRouter creates a new instance of PancakeRouter, bound to a specific deployed contract.
func NewPancakeRouter(address common.Address, backend bind.ContractBackend) (*PancakeRouter, error) {
	contract, err := bindPancakeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeRouter{PancakeRouterCaller: PancakeRouterCaller{contract: contract}, PancakeRouterTransactor: PancakeRouterTransactor{contract: contract}, PancakeRouterFilterer: PancakeRouterFilterer{contract: contract}}, nil
}

// NewPancakeRouterCaller creates a new read-only instance of PancakeRouter, bound to a specific deployed contract.
func NewPancakeRouterCaller(address common.Address, caller bind.ContractCaller) (*PancakeRouterCaller, error) {
	contract, err := bindPancakeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeRouterCaller{contract: contract}, nil
}

// NewPancakeRouterTransactor creates a new write-only instance of PancakeRouter, bound to a specific deployed contract.
func NewPancakeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeRouterTransactor, error) {
	contract, err := bindPancakeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeRouterTransactor{contract: contract}, nil
}

// NewPancakeRouterFilterer creates a new log filterer instance of PancakeRouter, bound to a specific deployed contract.
func NewPancakeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeRouterFilterer, error) {
	contract, err := bindPancakeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeRouterFilterer{contract: contract}, nil
}

// bindPancakeRouter binds a generic wrapper to an already deployed contract.
func bindPancakeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeRouter *PancakeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeRouter.Contract.PancakeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeRouter *PancakeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeRouter.Contract.PancakeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeRouter *PancakeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeRouter.Contract.PancakeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeRouter *PancakeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeRouter *PancakeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeRouter *PancakeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeRouter.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_PancakeRouter *PancakeRouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeRouter.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_PancakeRouter *PancakeRouterSession) WETH() (common.Address, error) {
	return _PancakeRouter.Contract.WETH(&_PancakeRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_PancakeRouter *PancakeRouterCallerSession) WETH() (common.Address, error) {
	return _PancakeRouter.Contract.WETH(&_PancakeRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeRouter *PancakeRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeRouter *PancakeRouterSession) Factory() (common.Address, error) {
	return _PancakeRouter.Contract.Factory(&_PancakeRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeRouter *PancakeRouterCallerSession) Factory() (common.Address, error) {
	return _PancakeRouter.Contract.Factory(&_PancakeRouter.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_PancakeRouter *PancakeRouterCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeRouter.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_PancakeRouter *PancakeRouterSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _PancakeRouter.Contract.GetAmountIn(&_PancakeRouter.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_PancakeRouter *PancakeRouterCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _PancakeRouter.Contract.GetAmountIn(&_PancakeRouter.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_PancakeRouter *PancakeRouterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeRouter.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_PancakeRouter *PancakeRouterSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _PancakeRouter.Contract.GetAmountOut(&_PancakeRouter.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_PancakeRouter *PancakeRouterCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _PancakeRouter.Contract.GetAmountOut(&_PancakeRouter.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _PancakeRouter.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _PancakeRouter.Contract.GetAmountsIn(&_PancakeRouter.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _PancakeRouter.Contract.GetAmountsIn(&_PancakeRouter.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _PancakeRouter.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _PancakeRouter.Contract.GetAmountsOut(&_PancakeRouter.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _PancakeRouter.Contract.GetAmountsOut(&_PancakeRouter.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_PancakeRouter *PancakeRouterCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeRouter.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_PancakeRouter *PancakeRouterSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _PancakeRouter.Contract.Quote(&_PancakeRouter.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_PancakeRouter *PancakeRouterCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _PancakeRouter.Contract.Quote(&_PancakeRouter.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_PancakeRouter *PancakeRouterTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_PancakeRouter *PancakeRouterSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.AddLiquidity(&_PancakeRouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_PancakeRouter *PancakeRouterTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.AddLiquidity(&_PancakeRouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_PancakeRouter *PancakeRouterTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_PancakeRouter *PancakeRouterSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.AddLiquidityETH(&_PancakeRouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_PancakeRouter *PancakeRouterTransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.AddLiquidityETH(&_PancakeRouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_PancakeRouter *PancakeRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_PancakeRouter *PancakeRouterSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidity(&_PancakeRouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_PancakeRouter *PancakeRouterTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidity(&_PancakeRouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_PancakeRouter *PancakeRouterTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_PancakeRouter *PancakeRouterSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityETH(&_PancakeRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_PancakeRouter *PancakeRouterTransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityETH(&_PancakeRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_PancakeRouter *PancakeRouterTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_PancakeRouter *PancakeRouterSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_PancakeRouter *PancakeRouterTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_PancakeRouter *PancakeRouterTransactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_PancakeRouter *PancakeRouterSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityETHWithPermit(&_PancakeRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_PancakeRouter *PancakeRouterTransactorSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityETHWithPermit(&_PancakeRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_PancakeRouter *PancakeRouterTransactor) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_PancakeRouter *PancakeRouterSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_PancakeRouter *PancakeRouterTransactorSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_PancakeRouter *PancakeRouterTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_PancakeRouter *PancakeRouterSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityWithPermit(&_PancakeRouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_PancakeRouter *PancakeRouterTransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeRouter.Contract.RemoveLiquidityWithPermit(&_PancakeRouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapETHForExactTokens(&_PancakeRouter.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapETHForExactTokens(&_PancakeRouter.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactETHForTokens(&_PancakeRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactETHForTokens(&_PancakeRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_PancakeRouter *PancakeRouterTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_PancakeRouter *PancakeRouterSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_PancakeRouter *PancakeRouterTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactTokensForETH(&_PancakeRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactTokensForETH(&_PancakeRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_PancakeRouter *PancakeRouterTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_PancakeRouter *PancakeRouterSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_PancakeRouter *PancakeRouterTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactTokensForTokens(&_PancakeRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactTokensForTokens(&_PancakeRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_PancakeRouter *PancakeRouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_PancakeRouter *PancakeRouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_PancakeRouter *PancakeRouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_PancakeRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapTokensForExactETH(&_PancakeRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapTokensForExactETH(&_PancakeRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapTokensForExactTokens(&_PancakeRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_PancakeRouter *PancakeRouterTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.SwapTokensForExactTokens(&_PancakeRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PancakeRouter *PancakeRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PancakeRouter *PancakeRouterSession) Receive() (*types.Transaction, error) {
	return _PancakeRouter.Contract.Receive(&_PancakeRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PancakeRouter *PancakeRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _PancakeRouter.Contract.Receive(&_PancakeRouter.TransactOpts)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122091f288bb807e5d98b1c59712d537917fd1137c95d69950e13f2d428ecc0faa6d64736f6c634300060c0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TransferHelperABI is the input ABI used to generate the binding from.
const TransferHelperABI = "[]"

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
var TransferHelperBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122048285055e743e41c4b5586eed03165024d7101dd207f00372c4630fb8211867d64736f6c634300060c0033"

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}
