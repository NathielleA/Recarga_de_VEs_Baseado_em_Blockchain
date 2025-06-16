// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pricePerSecond\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chargeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"ChargingCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ChargingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StationReserved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"chargeStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentStatus\",\"outputs\":[{\"internalType\":\"enumChargingStation.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endCharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chargeTime\",\"type\":\"uint256\"}],\"name\":\"payForCharge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startCharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalChargedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// ChargeStartTime is a free data retrieval call binding the contract method 0xe9f161df.
//
// Solidity: function chargeStartTime() view returns(uint256)
func (_Bindings *BindingsCaller) ChargeStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "chargeStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChargeStartTime is a free data retrieval call binding the contract method 0xe9f161df.
//
// Solidity: function chargeStartTime() view returns(uint256)
func (_Bindings *BindingsSession) ChargeStartTime() (*big.Int, error) {
	return _Bindings.Contract.ChargeStartTime(&_Bindings.CallOpts)
}

// ChargeStartTime is a free data retrieval call binding the contract method 0xe9f161df.
//
// Solidity: function chargeStartTime() view returns(uint256)
func (_Bindings *BindingsCallerSession) ChargeStartTime() (*big.Int, error) {
	return _Bindings.Contract.ChargeStartTime(&_Bindings.CallOpts)
}

// CurrentStatus is a free data retrieval call binding the contract method 0xef8a9235.
//
// Solidity: function currentStatus() view returns(uint8)
func (_Bindings *BindingsCaller) CurrentStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "currentStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurrentStatus is a free data retrieval call binding the contract method 0xef8a9235.
//
// Solidity: function currentStatus() view returns(uint8)
func (_Bindings *BindingsSession) CurrentStatus() (uint8, error) {
	return _Bindings.Contract.CurrentStatus(&_Bindings.CallOpts)
}

// CurrentStatus is a free data retrieval call binding the contract method 0xef8a9235.
//
// Solidity: function currentStatus() view returns(uint8)
func (_Bindings *BindingsCallerSession) CurrentStatus() (uint8, error) {
	return _Bindings.Contract.CurrentStatus(&_Bindings.CallOpts)
}

// CurrentUser is a free data retrieval call binding the contract method 0x92ee0334.
//
// Solidity: function currentUser() view returns(address)
func (_Bindings *BindingsCaller) CurrentUser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "currentUser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentUser is a free data retrieval call binding the contract method 0x92ee0334.
//
// Solidity: function currentUser() view returns(address)
func (_Bindings *BindingsSession) CurrentUser() (common.Address, error) {
	return _Bindings.Contract.CurrentUser(&_Bindings.CallOpts)
}

// CurrentUser is a free data retrieval call binding the contract method 0x92ee0334.
//
// Solidity: function currentUser() view returns(address)
func (_Bindings *BindingsCallerSession) CurrentUser() (common.Address, error) {
	return _Bindings.Contract.CurrentUser(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// PricePerSecond is a free data retrieval call binding the contract method 0x3804fd1f.
//
// Solidity: function pricePerSecond() view returns(uint256)
func (_Bindings *BindingsCaller) PricePerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pricePerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerSecond is a free data retrieval call binding the contract method 0x3804fd1f.
//
// Solidity: function pricePerSecond() view returns(uint256)
func (_Bindings *BindingsSession) PricePerSecond() (*big.Int, error) {
	return _Bindings.Contract.PricePerSecond(&_Bindings.CallOpts)
}

// PricePerSecond is a free data retrieval call binding the contract method 0x3804fd1f.
//
// Solidity: function pricePerSecond() view returns(uint256)
func (_Bindings *BindingsCallerSession) PricePerSecond() (*big.Int, error) {
	return _Bindings.Contract.PricePerSecond(&_Bindings.CallOpts)
}

// ReservationTime is a free data retrieval call binding the contract method 0xbc9071b1.
//
// Solidity: function reservationTime() view returns(uint256)
func (_Bindings *BindingsCaller) ReservationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "reservationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservationTime is a free data retrieval call binding the contract method 0xbc9071b1.
//
// Solidity: function reservationTime() view returns(uint256)
func (_Bindings *BindingsSession) ReservationTime() (*big.Int, error) {
	return _Bindings.Contract.ReservationTime(&_Bindings.CallOpts)
}

// ReservationTime is a free data retrieval call binding the contract method 0xbc9071b1.
//
// Solidity: function reservationTime() view returns(uint256)
func (_Bindings *BindingsCallerSession) ReservationTime() (*big.Int, error) {
	return _Bindings.Contract.ReservationTime(&_Bindings.CallOpts)
}

// TotalChargedTime is a free data retrieval call binding the contract method 0xfbff897d.
//
// Solidity: function totalChargedTime() view returns(uint256)
func (_Bindings *BindingsCaller) TotalChargedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "totalChargedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalChargedTime is a free data retrieval call binding the contract method 0xfbff897d.
//
// Solidity: function totalChargedTime() view returns(uint256)
func (_Bindings *BindingsSession) TotalChargedTime() (*big.Int, error) {
	return _Bindings.Contract.TotalChargedTime(&_Bindings.CallOpts)
}

// TotalChargedTime is a free data retrieval call binding the contract method 0xfbff897d.
//
// Solidity: function totalChargedTime() view returns(uint256)
func (_Bindings *BindingsCallerSession) TotalChargedTime() (*big.Int, error) {
	return _Bindings.Contract.TotalChargedTime(&_Bindings.CallOpts)
}

// EndCharge is a paid mutator transaction binding the contract method 0x7deb31dc.
//
// Solidity: function endCharge() returns()
func (_Bindings *BindingsTransactor) EndCharge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "endCharge")
}

// EndCharge is a paid mutator transaction binding the contract method 0x7deb31dc.
//
// Solidity: function endCharge() returns()
func (_Bindings *BindingsSession) EndCharge() (*types.Transaction, error) {
	return _Bindings.Contract.EndCharge(&_Bindings.TransactOpts)
}

// EndCharge is a paid mutator transaction binding the contract method 0x7deb31dc.
//
// Solidity: function endCharge() returns()
func (_Bindings *BindingsTransactorSession) EndCharge() (*types.Transaction, error) {
	return _Bindings.Contract.EndCharge(&_Bindings.TransactOpts)
}

// PayForCharge is a paid mutator transaction binding the contract method 0xf667a752.
//
// Solidity: function payForCharge(uint256 _chargeTime) payable returns()
func (_Bindings *BindingsTransactor) PayForCharge(opts *bind.TransactOpts, _chargeTime *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "payForCharge", _chargeTime)
}

// PayForCharge is a paid mutator transaction binding the contract method 0xf667a752.
//
// Solidity: function payForCharge(uint256 _chargeTime) payable returns()
func (_Bindings *BindingsSession) PayForCharge(_chargeTime *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PayForCharge(&_Bindings.TransactOpts, _chargeTime)
}

// PayForCharge is a paid mutator transaction binding the contract method 0xf667a752.
//
// Solidity: function payForCharge(uint256 _chargeTime) payable returns()
func (_Bindings *BindingsTransactorSession) PayForCharge(_chargeTime *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PayForCharge(&_Bindings.TransactOpts, _chargeTime)
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_Bindings *BindingsTransactor) Reserve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "reserve")
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_Bindings *BindingsSession) Reserve() (*types.Transaction, error) {
	return _Bindings.Contract.Reserve(&_Bindings.TransactOpts)
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_Bindings *BindingsTransactorSession) Reserve() (*types.Transaction, error) {
	return _Bindings.Contract.Reserve(&_Bindings.TransactOpts)
}

// StartCharge is a paid mutator transaction binding the contract method 0x1eb467ba.
//
// Solidity: function startCharge() returns()
func (_Bindings *BindingsTransactor) StartCharge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "startCharge")
}

// StartCharge is a paid mutator transaction binding the contract method 0x1eb467ba.
//
// Solidity: function startCharge() returns()
func (_Bindings *BindingsSession) StartCharge() (*types.Transaction, error) {
	return _Bindings.Contract.StartCharge(&_Bindings.TransactOpts)
}

// StartCharge is a paid mutator transaction binding the contract method 0x1eb467ba.
//
// Solidity: function startCharge() returns()
func (_Bindings *BindingsTransactorSession) StartCharge() (*types.Transaction, error) {
	return _Bindings.Contract.StartCharge(&_Bindings.TransactOpts)
}

// BindingsChargingCompletedIterator is returned from FilterChargingCompleted and is used to iterate over the raw logs and unpacked data for ChargingCompleted events raised by the Bindings contract.
type BindingsChargingCompletedIterator struct {
	Event *BindingsChargingCompleted // Event containing the contract specifics and raw log

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
func (it *BindingsChargingCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsChargingCompleted)
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
		it.Event = new(BindingsChargingCompleted)
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
func (it *BindingsChargingCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsChargingCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsChargingCompleted represents a ChargingCompleted event raised by the Bindings contract.
type BindingsChargingCompleted struct {
	User       common.Address
	ChargeTime *big.Int
	Cost       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChargingCompleted is a free log retrieval operation binding the contract event 0xd634322d22e7e91580cc86a4d1a431bf2b499fbdc65a723c7da4d64aafff3d10.
//
// Solidity: event ChargingCompleted(address indexed user, uint256 chargeTime, uint256 cost)
func (_Bindings *BindingsFilterer) FilterChargingCompleted(opts *bind.FilterOpts, user []common.Address) (*BindingsChargingCompletedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ChargingCompleted", userRule)
	if err != nil {
		return nil, err
	}
	return &BindingsChargingCompletedIterator{contract: _Bindings.contract, event: "ChargingCompleted", logs: logs, sub: sub}, nil
}

// WatchChargingCompleted is a free log subscription operation binding the contract event 0xd634322d22e7e91580cc86a4d1a431bf2b499fbdc65a723c7da4d64aafff3d10.
//
// Solidity: event ChargingCompleted(address indexed user, uint256 chargeTime, uint256 cost)
func (_Bindings *BindingsFilterer) WatchChargingCompleted(opts *bind.WatchOpts, sink chan<- *BindingsChargingCompleted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ChargingCompleted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsChargingCompleted)
				if err := _Bindings.contract.UnpackLog(event, "ChargingCompleted", log); err != nil {
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

// ParseChargingCompleted is a log parse operation binding the contract event 0xd634322d22e7e91580cc86a4d1a431bf2b499fbdc65a723c7da4d64aafff3d10.
//
// Solidity: event ChargingCompleted(address indexed user, uint256 chargeTime, uint256 cost)
func (_Bindings *BindingsFilterer) ParseChargingCompleted(log types.Log) (*BindingsChargingCompleted, error) {
	event := new(BindingsChargingCompleted)
	if err := _Bindings.contract.UnpackLog(event, "ChargingCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsChargingStartedIterator is returned from FilterChargingStarted and is used to iterate over the raw logs and unpacked data for ChargingStarted events raised by the Bindings contract.
type BindingsChargingStartedIterator struct {
	Event *BindingsChargingStarted // Event containing the contract specifics and raw log

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
func (it *BindingsChargingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsChargingStarted)
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
		it.Event = new(BindingsChargingStarted)
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
func (it *BindingsChargingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsChargingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsChargingStarted represents a ChargingStarted event raised by the Bindings contract.
type BindingsChargingStarted struct {
	User      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChargingStarted is a free log retrieval operation binding the contract event 0xd00e45240cc6bfd6a05c69740bd0308f288041905cc603e4f0e9004714823815.
//
// Solidity: event ChargingStarted(address indexed user, uint256 timestamp)
func (_Bindings *BindingsFilterer) FilterChargingStarted(opts *bind.FilterOpts, user []common.Address) (*BindingsChargingStartedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ChargingStarted", userRule)
	if err != nil {
		return nil, err
	}
	return &BindingsChargingStartedIterator{contract: _Bindings.contract, event: "ChargingStarted", logs: logs, sub: sub}, nil
}

// WatchChargingStarted is a free log subscription operation binding the contract event 0xd00e45240cc6bfd6a05c69740bd0308f288041905cc603e4f0e9004714823815.
//
// Solidity: event ChargingStarted(address indexed user, uint256 timestamp)
func (_Bindings *BindingsFilterer) WatchChargingStarted(opts *bind.WatchOpts, sink chan<- *BindingsChargingStarted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ChargingStarted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsChargingStarted)
				if err := _Bindings.contract.UnpackLog(event, "ChargingStarted", log); err != nil {
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

// ParseChargingStarted is a log parse operation binding the contract event 0xd00e45240cc6bfd6a05c69740bd0308f288041905cc603e4f0e9004714823815.
//
// Solidity: event ChargingStarted(address indexed user, uint256 timestamp)
func (_Bindings *BindingsFilterer) ParseChargingStarted(log types.Log) (*BindingsChargingStarted, error) {
	event := new(BindingsChargingStarted)
	if err := _Bindings.contract.UnpackLog(event, "ChargingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the Bindings contract.
type BindingsPaymentReceivedIterator struct {
	Event *BindingsPaymentReceived // Event containing the contract specifics and raw log

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
func (it *BindingsPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPaymentReceived)
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
		it.Event = new(BindingsPaymentReceived)
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
func (it *BindingsPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPaymentReceived represents a PaymentReceived event raised by the Bindings contract.
type BindingsPaymentReceived struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address indexed user, uint256 amount)
func (_Bindings *BindingsFilterer) FilterPaymentReceived(opts *bind.FilterOpts, user []common.Address) (*BindingsPaymentReceivedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PaymentReceived", userRule)
	if err != nil {
		return nil, err
	}
	return &BindingsPaymentReceivedIterator{contract: _Bindings.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address indexed user, uint256 amount)
func (_Bindings *BindingsFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *BindingsPaymentReceived, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PaymentReceived", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPaymentReceived)
				if err := _Bindings.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
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

// ParsePaymentReceived is a log parse operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address indexed user, uint256 amount)
func (_Bindings *BindingsFilterer) ParsePaymentReceived(log types.Log) (*BindingsPaymentReceived, error) {
	event := new(BindingsPaymentReceived)
	if err := _Bindings.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsStationReservedIterator is returned from FilterStationReserved and is used to iterate over the raw logs and unpacked data for StationReserved events raised by the Bindings contract.
type BindingsStationReservedIterator struct {
	Event *BindingsStationReserved // Event containing the contract specifics and raw log

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
func (it *BindingsStationReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsStationReserved)
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
		it.Event = new(BindingsStationReserved)
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
func (it *BindingsStationReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsStationReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsStationReserved represents a StationReserved event raised by the Bindings contract.
type BindingsStationReserved struct {
	User      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStationReserved is a free log retrieval operation binding the contract event 0x5240346fd2f36f311dab71002c255639afbf28d09bd63cfb5d549f28927993ba.
//
// Solidity: event StationReserved(address indexed user, uint256 timestamp)
func (_Bindings *BindingsFilterer) FilterStationReserved(opts *bind.FilterOpts, user []common.Address) (*BindingsStationReservedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "StationReserved", userRule)
	if err != nil {
		return nil, err
	}
	return &BindingsStationReservedIterator{contract: _Bindings.contract, event: "StationReserved", logs: logs, sub: sub}, nil
}

// WatchStationReserved is a free log subscription operation binding the contract event 0x5240346fd2f36f311dab71002c255639afbf28d09bd63cfb5d549f28927993ba.
//
// Solidity: event StationReserved(address indexed user, uint256 timestamp)
func (_Bindings *BindingsFilterer) WatchStationReserved(opts *bind.WatchOpts, sink chan<- *BindingsStationReserved, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "StationReserved", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsStationReserved)
				if err := _Bindings.contract.UnpackLog(event, "StationReserved", log); err != nil {
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

// ParseStationReserved is a log parse operation binding the contract event 0x5240346fd2f36f311dab71002c255639afbf28d09bd63cfb5d549f28927993ba.
//
// Solidity: event StationReserved(address indexed user, uint256 timestamp)
func (_Bindings *BindingsFilterer) ParseStationReserved(log types.Log) (*BindingsStationReserved, error) {
	event := new(BindingsStationReserved)
	if err := _Bindings.contract.UnpackLog(event, "StationReserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
