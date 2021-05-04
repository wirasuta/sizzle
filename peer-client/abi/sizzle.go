// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sizzle

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

// SizzleCertMetadata is an auto generated low-level Go binding around an user-defined struct.
type SizzleCertMetadata struct {
	Owner         common.Address
	Domain        string
	PubKey        string
	Reputation    *big.Int
	ReputationMax *big.Int
	Valid         bool
}

// SizzlePeerMetadata is an auto generated low-level Go binding around an user-defined struct.
type SizzlePeerMetadata struct {
	Addr       common.Address
	Reputation *big.Int
}

// SizzleABI is the input ABI used to generate the binding from.
const SizzleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"peer\",\"type\":\"address\"}],\"name\":\"CertDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"peer\",\"type\":\"address\"}],\"name\":\"CertEndorsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"CertPublishRequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"reputation\",\"type\":\"int256\"}],\"name\":\"CertValid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"certPublishRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"certEndorseByPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"certDenyByPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"certEndorseByUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"certDenyByUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"certQuery\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"reputation\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"reputationMax\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structSizzle.CertMetadata\",\"name\":\"cert\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peerRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"peerQuery\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"reputation\",\"type\":\"int256\"}],\"internalType\":\"structSizzle.PeerMetadata\",\"name\":\"peer\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Sizzle is an auto generated Go binding around an Ethereum contract.
type Sizzle struct {
	SizzleCaller     // Read-only binding to the contract
	SizzleTransactor // Write-only binding to the contract
	SizzleFilterer   // Log filterer for contract events
}

// SizzleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SizzleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SizzleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SizzleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SizzleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SizzleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SizzleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SizzleSession struct {
	Contract     *Sizzle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SizzleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SizzleCallerSession struct {
	Contract *SizzleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SizzleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SizzleTransactorSession struct {
	Contract     *SizzleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SizzleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SizzleRaw struct {
	Contract *Sizzle // Generic contract binding to access the raw methods on
}

// SizzleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SizzleCallerRaw struct {
	Contract *SizzleCaller // Generic read-only contract binding to access the raw methods on
}

// SizzleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SizzleTransactorRaw struct {
	Contract *SizzleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSizzle creates a new instance of Sizzle, bound to a specific deployed contract.
func NewSizzle(address common.Address, backend bind.ContractBackend) (*Sizzle, error) {
	contract, err := bindSizzle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sizzle{SizzleCaller: SizzleCaller{contract: contract}, SizzleTransactor: SizzleTransactor{contract: contract}, SizzleFilterer: SizzleFilterer{contract: contract}}, nil
}

// NewSizzleCaller creates a new read-only instance of Sizzle, bound to a specific deployed contract.
func NewSizzleCaller(address common.Address, caller bind.ContractCaller) (*SizzleCaller, error) {
	contract, err := bindSizzle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SizzleCaller{contract: contract}, nil
}

// NewSizzleTransactor creates a new write-only instance of Sizzle, bound to a specific deployed contract.
func NewSizzleTransactor(address common.Address, transactor bind.ContractTransactor) (*SizzleTransactor, error) {
	contract, err := bindSizzle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SizzleTransactor{contract: contract}, nil
}

// NewSizzleFilterer creates a new log filterer instance of Sizzle, bound to a specific deployed contract.
func NewSizzleFilterer(address common.Address, filterer bind.ContractFilterer) (*SizzleFilterer, error) {
	contract, err := bindSizzle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SizzleFilterer{contract: contract}, nil
}

// bindSizzle binds a generic wrapper to an already deployed contract.
func bindSizzle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SizzleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sizzle *SizzleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sizzle.Contract.SizzleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sizzle *SizzleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sizzle.Contract.SizzleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sizzle *SizzleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sizzle.Contract.SizzleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sizzle *SizzleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sizzle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sizzle *SizzleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sizzle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sizzle *SizzleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sizzle.Contract.contract.Transact(opts, method, params...)
}

// CertQuery is a free data retrieval call binding the contract method 0x55476f9d.
//
// Solidity: function certQuery(string domain) view returns((address,string,string,int256,int256,bool) cert)
func (_Sizzle *SizzleCaller) CertQuery(opts *bind.CallOpts, domain string) (SizzleCertMetadata, error) {
	var out []interface{}
	err := _Sizzle.contract.Call(opts, &out, "certQuery", domain)

	if err != nil {
		return *new(SizzleCertMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(SizzleCertMetadata)).(*SizzleCertMetadata)

	return out0, err

}

// CertQuery is a free data retrieval call binding the contract method 0x55476f9d.
//
// Solidity: function certQuery(string domain) view returns((address,string,string,int256,int256,bool) cert)
func (_Sizzle *SizzleSession) CertQuery(domain string) (SizzleCertMetadata, error) {
	return _Sizzle.Contract.CertQuery(&_Sizzle.CallOpts, domain)
}

// CertQuery is a free data retrieval call binding the contract method 0x55476f9d.
//
// Solidity: function certQuery(string domain) view returns((address,string,string,int256,int256,bool) cert)
func (_Sizzle *SizzleCallerSession) CertQuery(domain string) (SizzleCertMetadata, error) {
	return _Sizzle.Contract.CertQuery(&_Sizzle.CallOpts, domain)
}

// PeerQuery is a free data retrieval call binding the contract method 0xf4dcc32d.
//
// Solidity: function peerQuery(address addr) view returns((address,int256) peer)
func (_Sizzle *SizzleCaller) PeerQuery(opts *bind.CallOpts, addr common.Address) (SizzlePeerMetadata, error) {
	var out []interface{}
	err := _Sizzle.contract.Call(opts, &out, "peerQuery", addr)

	if err != nil {
		return *new(SizzlePeerMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(SizzlePeerMetadata)).(*SizzlePeerMetadata)

	return out0, err

}

// PeerQuery is a free data retrieval call binding the contract method 0xf4dcc32d.
//
// Solidity: function peerQuery(address addr) view returns((address,int256) peer)
func (_Sizzle *SizzleSession) PeerQuery(addr common.Address) (SizzlePeerMetadata, error) {
	return _Sizzle.Contract.PeerQuery(&_Sizzle.CallOpts, addr)
}

// PeerQuery is a free data retrieval call binding the contract method 0xf4dcc32d.
//
// Solidity: function peerQuery(address addr) view returns((address,int256) peer)
func (_Sizzle *SizzleCallerSession) PeerQuery(addr common.Address) (SizzlePeerMetadata, error) {
	return _Sizzle.Contract.PeerQuery(&_Sizzle.CallOpts, addr)
}

// CertDenyByPeer is a paid mutator transaction binding the contract method 0x0f873b70.
//
// Solidity: function certDenyByPeer(string domain) returns()
func (_Sizzle *SizzleTransactor) CertDenyByPeer(opts *bind.TransactOpts, domain string) (*types.Transaction, error) {
	return _Sizzle.contract.Transact(opts, "certDenyByPeer", domain)
}

// CertDenyByPeer is a paid mutator transaction binding the contract method 0x0f873b70.
//
// Solidity: function certDenyByPeer(string domain) returns()
func (_Sizzle *SizzleSession) CertDenyByPeer(domain string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertDenyByPeer(&_Sizzle.TransactOpts, domain)
}

// CertDenyByPeer is a paid mutator transaction binding the contract method 0x0f873b70.
//
// Solidity: function certDenyByPeer(string domain) returns()
func (_Sizzle *SizzleTransactorSession) CertDenyByPeer(domain string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertDenyByPeer(&_Sizzle.TransactOpts, domain)
}

// CertDenyByUser is a paid mutator transaction binding the contract method 0x65eea1df.
//
// Solidity: function certDenyByUser(string domain) returns()
func (_Sizzle *SizzleTransactor) CertDenyByUser(opts *bind.TransactOpts, domain string) (*types.Transaction, error) {
	return _Sizzle.contract.Transact(opts, "certDenyByUser", domain)
}

// CertDenyByUser is a paid mutator transaction binding the contract method 0x65eea1df.
//
// Solidity: function certDenyByUser(string domain) returns()
func (_Sizzle *SizzleSession) CertDenyByUser(domain string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertDenyByUser(&_Sizzle.TransactOpts, domain)
}

// CertDenyByUser is a paid mutator transaction binding the contract method 0x65eea1df.
//
// Solidity: function certDenyByUser(string domain) returns()
func (_Sizzle *SizzleTransactorSession) CertDenyByUser(domain string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertDenyByUser(&_Sizzle.TransactOpts, domain)
}

// CertEndorseByPeer is a paid mutator transaction binding the contract method 0x33b60a8f.
//
// Solidity: function certEndorseByPeer(string domain) returns()
func (_Sizzle *SizzleTransactor) CertEndorseByPeer(opts *bind.TransactOpts, domain string) (*types.Transaction, error) {
	return _Sizzle.contract.Transact(opts, "certEndorseByPeer", domain)
}

// CertEndorseByPeer is a paid mutator transaction binding the contract method 0x33b60a8f.
//
// Solidity: function certEndorseByPeer(string domain) returns()
func (_Sizzle *SizzleSession) CertEndorseByPeer(domain string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertEndorseByPeer(&_Sizzle.TransactOpts, domain)
}

// CertEndorseByPeer is a paid mutator transaction binding the contract method 0x33b60a8f.
//
// Solidity: function certEndorseByPeer(string domain) returns()
func (_Sizzle *SizzleTransactorSession) CertEndorseByPeer(domain string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertEndorseByPeer(&_Sizzle.TransactOpts, domain)
}

// CertEndorseByUser is a paid mutator transaction binding the contract method 0xd6695b52.
//
// Solidity: function certEndorseByUser(string domain) returns()
func (_Sizzle *SizzleTransactor) CertEndorseByUser(opts *bind.TransactOpts, domain string) (*types.Transaction, error) {
	return _Sizzle.contract.Transact(opts, "certEndorseByUser", domain)
}

// CertEndorseByUser is a paid mutator transaction binding the contract method 0xd6695b52.
//
// Solidity: function certEndorseByUser(string domain) returns()
func (_Sizzle *SizzleSession) CertEndorseByUser(domain string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertEndorseByUser(&_Sizzle.TransactOpts, domain)
}

// CertEndorseByUser is a paid mutator transaction binding the contract method 0xd6695b52.
//
// Solidity: function certEndorseByUser(string domain) returns()
func (_Sizzle *SizzleTransactorSession) CertEndorseByUser(domain string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertEndorseByUser(&_Sizzle.TransactOpts, domain)
}

// CertPublishRequest is a paid mutator transaction binding the contract method 0xe5c98d62.
//
// Solidity: function certPublishRequest(string domain, string pubKey) returns()
func (_Sizzle *SizzleTransactor) CertPublishRequest(opts *bind.TransactOpts, domain string, pubKey string) (*types.Transaction, error) {
	return _Sizzle.contract.Transact(opts, "certPublishRequest", domain, pubKey)
}

// CertPublishRequest is a paid mutator transaction binding the contract method 0xe5c98d62.
//
// Solidity: function certPublishRequest(string domain, string pubKey) returns()
func (_Sizzle *SizzleSession) CertPublishRequest(domain string, pubKey string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertPublishRequest(&_Sizzle.TransactOpts, domain, pubKey)
}

// CertPublishRequest is a paid mutator transaction binding the contract method 0xe5c98d62.
//
// Solidity: function certPublishRequest(string domain, string pubKey) returns()
func (_Sizzle *SizzleTransactorSession) CertPublishRequest(domain string, pubKey string) (*types.Transaction, error) {
	return _Sizzle.Contract.CertPublishRequest(&_Sizzle.TransactOpts, domain, pubKey)
}

// PeerRegister is a paid mutator transaction binding the contract method 0xa4394810.
//
// Solidity: function peerRegister() returns()
func (_Sizzle *SizzleTransactor) PeerRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sizzle.contract.Transact(opts, "peerRegister")
}

// PeerRegister is a paid mutator transaction binding the contract method 0xa4394810.
//
// Solidity: function peerRegister() returns()
func (_Sizzle *SizzleSession) PeerRegister() (*types.Transaction, error) {
	return _Sizzle.Contract.PeerRegister(&_Sizzle.TransactOpts)
}

// PeerRegister is a paid mutator transaction binding the contract method 0xa4394810.
//
// Solidity: function peerRegister() returns()
func (_Sizzle *SizzleTransactorSession) PeerRegister() (*types.Transaction, error) {
	return _Sizzle.Contract.PeerRegister(&_Sizzle.TransactOpts)
}

// SizzleCertDeniedIterator is returned from FilterCertDenied and is used to iterate over the raw logs and unpacked data for CertDenied events raised by the Sizzle contract.
type SizzleCertDeniedIterator struct {
	Event *SizzleCertDenied // Event containing the contract specifics and raw log

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
func (it *SizzleCertDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SizzleCertDenied)
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
		it.Event = new(SizzleCertDenied)
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
func (it *SizzleCertDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SizzleCertDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SizzleCertDenied represents a CertDenied event raised by the Sizzle contract.
type SizzleCertDenied struct {
	Domain string
	Peer   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCertDenied is a free log retrieval operation binding the contract event 0x976eb8888f24e203d6ec895956dce7a61216e6e3cf8a9151d2d5f6327154b562.
//
// Solidity: event CertDenied(string domain, address peer)
func (_Sizzle *SizzleFilterer) FilterCertDenied(opts *bind.FilterOpts) (*SizzleCertDeniedIterator, error) {

	logs, sub, err := _Sizzle.contract.FilterLogs(opts, "CertDenied")
	if err != nil {
		return nil, err
	}
	return &SizzleCertDeniedIterator{contract: _Sizzle.contract, event: "CertDenied", logs: logs, sub: sub}, nil
}

// WatchCertDenied is a free log subscription operation binding the contract event 0x976eb8888f24e203d6ec895956dce7a61216e6e3cf8a9151d2d5f6327154b562.
//
// Solidity: event CertDenied(string domain, address peer)
func (_Sizzle *SizzleFilterer) WatchCertDenied(opts *bind.WatchOpts, sink chan<- *SizzleCertDenied) (event.Subscription, error) {

	logs, sub, err := _Sizzle.contract.WatchLogs(opts, "CertDenied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SizzleCertDenied)
				if err := _Sizzle.contract.UnpackLog(event, "CertDenied", log); err != nil {
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

// ParseCertDenied is a log parse operation binding the contract event 0x976eb8888f24e203d6ec895956dce7a61216e6e3cf8a9151d2d5f6327154b562.
//
// Solidity: event CertDenied(string domain, address peer)
func (_Sizzle *SizzleFilterer) ParseCertDenied(log types.Log) (*SizzleCertDenied, error) {
	event := new(SizzleCertDenied)
	if err := _Sizzle.contract.UnpackLog(event, "CertDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SizzleCertEndorsedIterator is returned from FilterCertEndorsed and is used to iterate over the raw logs and unpacked data for CertEndorsed events raised by the Sizzle contract.
type SizzleCertEndorsedIterator struct {
	Event *SizzleCertEndorsed // Event containing the contract specifics and raw log

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
func (it *SizzleCertEndorsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SizzleCertEndorsed)
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
		it.Event = new(SizzleCertEndorsed)
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
func (it *SizzleCertEndorsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SizzleCertEndorsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SizzleCertEndorsed represents a CertEndorsed event raised by the Sizzle contract.
type SizzleCertEndorsed struct {
	Domain string
	Peer   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCertEndorsed is a free log retrieval operation binding the contract event 0x49fb6adaa570bbf0c679498e3eed668725b969b9367d983820ef7ed87747e444.
//
// Solidity: event CertEndorsed(string domain, address peer)
func (_Sizzle *SizzleFilterer) FilterCertEndorsed(opts *bind.FilterOpts) (*SizzleCertEndorsedIterator, error) {

	logs, sub, err := _Sizzle.contract.FilterLogs(opts, "CertEndorsed")
	if err != nil {
		return nil, err
	}
	return &SizzleCertEndorsedIterator{contract: _Sizzle.contract, event: "CertEndorsed", logs: logs, sub: sub}, nil
}

// WatchCertEndorsed is a free log subscription operation binding the contract event 0x49fb6adaa570bbf0c679498e3eed668725b969b9367d983820ef7ed87747e444.
//
// Solidity: event CertEndorsed(string domain, address peer)
func (_Sizzle *SizzleFilterer) WatchCertEndorsed(opts *bind.WatchOpts, sink chan<- *SizzleCertEndorsed) (event.Subscription, error) {

	logs, sub, err := _Sizzle.contract.WatchLogs(opts, "CertEndorsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SizzleCertEndorsed)
				if err := _Sizzle.contract.UnpackLog(event, "CertEndorsed", log); err != nil {
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

// ParseCertEndorsed is a log parse operation binding the contract event 0x49fb6adaa570bbf0c679498e3eed668725b969b9367d983820ef7ed87747e444.
//
// Solidity: event CertEndorsed(string domain, address peer)
func (_Sizzle *SizzleFilterer) ParseCertEndorsed(log types.Log) (*SizzleCertEndorsed, error) {
	event := new(SizzleCertEndorsed)
	if err := _Sizzle.contract.UnpackLog(event, "CertEndorsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SizzleCertPublishRequestCreatedIterator is returned from FilterCertPublishRequestCreated and is used to iterate over the raw logs and unpacked data for CertPublishRequestCreated events raised by the Sizzle contract.
type SizzleCertPublishRequestCreatedIterator struct {
	Event *SizzleCertPublishRequestCreated // Event containing the contract specifics and raw log

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
func (it *SizzleCertPublishRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SizzleCertPublishRequestCreated)
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
		it.Event = new(SizzleCertPublishRequestCreated)
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
func (it *SizzleCertPublishRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SizzleCertPublishRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SizzleCertPublishRequestCreated represents a CertPublishRequestCreated event raised by the Sizzle contract.
type SizzleCertPublishRequestCreated struct {
	Owner  common.Address
	Domain string
	PubKey string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCertPublishRequestCreated is a free log retrieval operation binding the contract event 0x685016f6d66a8b93123f62b2fcaaa502109fc3f76d89572c201a6f55ba8f9650.
//
// Solidity: event CertPublishRequestCreated(address owner, string domain, string pubKey)
func (_Sizzle *SizzleFilterer) FilterCertPublishRequestCreated(opts *bind.FilterOpts) (*SizzleCertPublishRequestCreatedIterator, error) {

	logs, sub, err := _Sizzle.contract.FilterLogs(opts, "CertPublishRequestCreated")
	if err != nil {
		return nil, err
	}
	return &SizzleCertPublishRequestCreatedIterator{contract: _Sizzle.contract, event: "CertPublishRequestCreated", logs: logs, sub: sub}, nil
}

// WatchCertPublishRequestCreated is a free log subscription operation binding the contract event 0x685016f6d66a8b93123f62b2fcaaa502109fc3f76d89572c201a6f55ba8f9650.
//
// Solidity: event CertPublishRequestCreated(address owner, string domain, string pubKey)
func (_Sizzle *SizzleFilterer) WatchCertPublishRequestCreated(opts *bind.WatchOpts, sink chan<- *SizzleCertPublishRequestCreated) (event.Subscription, error) {

	logs, sub, err := _Sizzle.contract.WatchLogs(opts, "CertPublishRequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SizzleCertPublishRequestCreated)
				if err := _Sizzle.contract.UnpackLog(event, "CertPublishRequestCreated", log); err != nil {
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

// ParseCertPublishRequestCreated is a log parse operation binding the contract event 0x685016f6d66a8b93123f62b2fcaaa502109fc3f76d89572c201a6f55ba8f9650.
//
// Solidity: event CertPublishRequestCreated(address owner, string domain, string pubKey)
func (_Sizzle *SizzleFilterer) ParseCertPublishRequestCreated(log types.Log) (*SizzleCertPublishRequestCreated, error) {
	event := new(SizzleCertPublishRequestCreated)
	if err := _Sizzle.contract.UnpackLog(event, "CertPublishRequestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SizzleCertValidIterator is returned from FilterCertValid and is used to iterate over the raw logs and unpacked data for CertValid events raised by the Sizzle contract.
type SizzleCertValidIterator struct {
	Event *SizzleCertValid // Event containing the contract specifics and raw log

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
func (it *SizzleCertValidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SizzleCertValid)
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
		it.Event = new(SizzleCertValid)
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
func (it *SizzleCertValidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SizzleCertValidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SizzleCertValid represents a CertValid event raised by the Sizzle contract.
type SizzleCertValid struct {
	Owner      common.Address
	Domain     string
	PubKey     string
	Reputation *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCertValid is a free log retrieval operation binding the contract event 0xf7f4c9af4aeb3c95f7dfe47c6d9b659084dae8e7cb54948065187000c1eae418.
//
// Solidity: event CertValid(address owner, string domain, string pubKey, int256 reputation)
func (_Sizzle *SizzleFilterer) FilterCertValid(opts *bind.FilterOpts) (*SizzleCertValidIterator, error) {

	logs, sub, err := _Sizzle.contract.FilterLogs(opts, "CertValid")
	if err != nil {
		return nil, err
	}
	return &SizzleCertValidIterator{contract: _Sizzle.contract, event: "CertValid", logs: logs, sub: sub}, nil
}

// WatchCertValid is a free log subscription operation binding the contract event 0xf7f4c9af4aeb3c95f7dfe47c6d9b659084dae8e7cb54948065187000c1eae418.
//
// Solidity: event CertValid(address owner, string domain, string pubKey, int256 reputation)
func (_Sizzle *SizzleFilterer) WatchCertValid(opts *bind.WatchOpts, sink chan<- *SizzleCertValid) (event.Subscription, error) {

	logs, sub, err := _Sizzle.contract.WatchLogs(opts, "CertValid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SizzleCertValid)
				if err := _Sizzle.contract.UnpackLog(event, "CertValid", log); err != nil {
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

// ParseCertValid is a log parse operation binding the contract event 0xf7f4c9af4aeb3c95f7dfe47c6d9b659084dae8e7cb54948065187000c1eae418.
//
// Solidity: event CertValid(address owner, string domain, string pubKey, int256 reputation)
func (_Sizzle *SizzleFilterer) ParseCertValid(log types.Log) (*SizzleCertValid, error) {
	event := new(SizzleCertValid)
	if err := _Sizzle.contract.UnpackLog(event, "CertValid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
