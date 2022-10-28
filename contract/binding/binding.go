// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

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
)

// BindingMetaData contains all meta data concerning the Binding contract.
var BindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"withdrawal_credentials\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"amount\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"index\",\"type\":\"bytes\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"withdrawal_credentials\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"deposit_data_root\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_deposit_count\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_deposit_root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060005b601f8110156101025760026021826020811061002c57fe5b01546021836020811061003b57fe5b015460405160200180838152602001828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106100925780518252601f199092019160209182019101610073565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156100d1573d6000803e3d6000fd5b5050506040513d60208110156100e657600080fd5b5051602160018301602081106100f857fe5b0155600101610014565b50611322806101126000396000f3fe60806040526004361061003f5760003560e01c806301ffc9a714610044578063228951181461008c578063621fd130146101a2578063c5f2892f1461022c575b600080fd5b34801561005057600080fd5b506100786004803603602081101561006757600080fd5b50356001600160e01b031916610253565b604080519115158252519081900360200190f35b6101a0600480360360808110156100a257600080fd5b8101906020810181356401000000008111156100bd57600080fd5b8201836020820111156100cf57600080fd5b803590602001918460018302840111640100000000831117156100f157600080fd5b91939092909160208101903564010000000081111561010f57600080fd5b82018360208201111561012157600080fd5b8035906020019184600183028401116401000000008311171561014357600080fd5b91939092909160208101903564010000000081111561016157600080fd5b82018360208201111561017357600080fd5b8035906020019184600183028401116401000000008311171561019557600080fd5b91935091503561028a565b005b3480156101ae57600080fd5b506101b7610ce6565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101f15781810151838201526020016101d9565b50505050905090810190601f16801561021e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023857600080fd5b50610241610cf8565b60408051918252519081900360200190f35b60006001600160e01b031982166301ffc9a760e01b148061028457506001600160e01b03198216638564090760e01b145b92915050565b603086146102c95760405162461bcd60e51b81526004018080602001828103825260268152602001806112516026913960400191505060405180910390fd5b602084146103085760405162461bcd60e51b81526004018080602001828103825260368152602001806111e86036913960400191505060405180910390fd5b606082146103475760405162461bcd60e51b81526004018080602001828103825260298152602001806112c46029913960400191505060405180910390fd5b670de0b6b3a764000034101561038e5760405162461bcd60e51b815260040180806020018281038252602681526020018061129e6026913960400191505060405180910390fd5b633b9aca003406156103d15760405162461bcd60e51b815260040180806020018281038252603381526020018061121e6033913960400191505060405180910390fd5b633b9aca00340467ffffffffffffffff81111561041f5760405162461bcd60e51b81526004018080602001828103825260278152602001806112776027913960400191505060405180910390fd5b606061042a82610fc6565b90507f649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c589898989858a8a61045f602054610fc6565b6040805160a0808252810189905290819060208201908201606083016080840160c085018e8e80828437600083820152601f01601f191690910187810386528c815260200190508c8c808284376000838201819052601f909101601f191690920188810386528c5181528c51602091820193918e019250908190849084905b838110156104f65781810151838201526020016104de565b50505050905090810190601f1680156105235780820380516001836020036101000a031916815260200191505b5086810383528881526020018989808284376000838201819052601f909101601f19169092018881038452895181528951602091820193918b019250908190849084905b8381101561057f578181015183820152602001610567565b50505050905090810190601f1680156105ac5780820380516001836020036101000a031916815260200191505b509d505050505050505050505050505060405180910390a1600060028a8a600060801b604051602001808484808284376fffffffffffffffffffffffffffffffff199094169190930190815260408051600f19818403018152601090920190819052815191955093508392506020850191508083835b602083106106415780518252601f199092019160209182019101610622565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610680573d6000803e3d6000fd5b5050506040513d602081101561069557600080fd5b5051905060006002806106ab6040848a8c61114a565b6040516020018083838082843780830192505050925050506040516020818303038152906040526040518082805190602001908083835b602083106107015780518252601f1990920191602091820191016106e2565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610740573d6000803e3d6000fd5b5050506040513d602081101561075557600080fd5b50516002610766896040818d61114a565b60405160009060200180848480828437919091019283525050604080518083038152602092830191829052805190945090925082918401908083835b602083106107c15780518252601f1990920191602091820191016107a2565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610800573d6000803e3d6000fd5b5050506040513d602081101561081557600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019081905281519192909182918401908083835b6020831061086b5780518252601f19909201916020918201910161084c565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156108aa573d6000803e3d6000fd5b5050506040513d60208110156108bf57600080fd5b50516040805160208101858152929350600092600292839287928f928f92018383808284378083019250505093505050506040516020818303038152906040526040518082805190602001908083835b6020831061092e5780518252601f19909201916020918201910161090f565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561096d573d6000803e3d6000fd5b5050506040513d602081101561098257600080fd5b50516040518651600291889160009188916020918201918291908601908083835b602083106109c25780518252601f1990920191602091820191016109a3565b6001836020036101000a0380198251168184511680821785525050505050509050018367ffffffffffffffff191667ffffffffffffffff1916815260180182815260200193505050506040516020818303038152906040526040518082805190602001908083835b60208310610a495780518252601f199092019160209182019101610a2a565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610a88573d6000803e3d6000fd5b5050506040513d6020811015610a9d57600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019081905281519192909182918401908083835b60208310610af35780518252601f199092019160209182019101610ad4565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610b32573d6000803e3d6000fd5b5050506040513d6020811015610b4757600080fd5b50519050858114610b895760405162461bcd60e51b81526004018080602001828103825260548152602001806111946054913960600191505060405180910390fd5b60205463ffffffff11610bcd5760405162461bcd60e51b81526004018080602001828103825260218152602001806111736021913960400191505060405180910390fd5b602080546001019081905560005b6020811015610cda578160011660011415610c0d578260008260208110610bfe57fe5b015550610cdd95505050505050565b600260008260208110610c1c57fe5b01548460405160200180838152602001828152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310610c745780518252601f199092019160209182019101610c55565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610cb3573d6000803e3d6000fd5b5050506040513d6020811015610cc857600080fd5b50519250600282049150600101610bdb565b50fe5b50505050505050565b6060610cf3602054610fc6565b905090565b6020546000908190815b6020811015610ea9578160011660011415610ddb57600260008260208110610d2657fe5b01548460405160200180838152602001828152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310610d7e5780518252601f199092019160209182019101610d5f565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610dbd573d6000803e3d6000fd5b5050506040513d6020811015610dd257600080fd5b50519250610e9b565b60028360218360208110610deb57fe5b015460405160200180838152602001828152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310610e425780518252601f199092019160209182019101610e23565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610e81573d6000803e3d6000fd5b5050506040513d6020811015610e9657600080fd5b505192505b600282049150600101610d02565b50600282610eb8602054610fc6565b600060401b6040516020018084815260200183805190602001908083835b60208310610ef55780518252601f199092019160209182019101610ed6565b51815160209384036101000a600019018019909216911617905267ffffffffffffffff199590951692019182525060408051808303600719018152601890920190819052815191955093508392850191508083835b60208310610f695780518252601f199092019160209182019101610f4a565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610fa8573d6000803e3d6000fd5b5050506040513d6020811015610fbd57600080fd5b50519250505090565b60408051600880825281830190925260609160208201818036833701905050905060c082901b8060071a60f81b8260008151811061100057fe5b60200101906001600160f81b031916908160001a9053508060061a60f81b8260018151811061102b57fe5b60200101906001600160f81b031916908160001a9053508060051a60f81b8260028151811061105657fe5b60200101906001600160f81b031916908160001a9053508060041a60f81b8260038151811061108157fe5b60200101906001600160f81b031916908160001a9053508060031a60f81b826004815181106110ac57fe5b60200101906001600160f81b031916908160001a9053508060021a60f81b826005815181106110d757fe5b60200101906001600160f81b031916908160001a9053508060011a60f81b8260068151811061110257fe5b60200101906001600160f81b031916908160001a9053508060001a60f81b8260078151811061112d57fe5b60200101906001600160f81b031916908160001a90535050919050565b60008085851115611159578182fd5b83861115611165578182fd5b505082019391909203915056fe4465706f736974436f6e74726163743a206d65726b6c6520747265652066756c6c4465706f736974436f6e74726163743a207265636f6e7374727563746564204465706f7369744461746120646f6573206e6f74206d6174636820737570706c696564206465706f7369745f646174615f726f6f744465706f736974436f6e74726163743a20696e76616c6964207769746864726177616c5f63726564656e7469616c73206c656e6774684465706f736974436f6e74726163743a206465706f7369742076616c7565206e6f74206d756c7469706c65206f6620677765694465706f736974436f6e74726163743a20696e76616c6964207075626b6579206c656e6774684465706f736974436f6e74726163743a206465706f7369742076616c756520746f6f20686967684465706f736974436f6e74726163743a206465706f7369742076616c756520746f6f206c6f774465706f736974436f6e74726163743a20696e76616c6964207369676e6174757265206c656e677468a2646970667358221220187e8867b451f9e0a5b7ac4b7a81e2f993d02f830006f0cc4e28d5ad9ac13baa64736f6c634300060b0033",
}

// BindingABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingMetaData.ABI instead.
var BindingABI = BindingMetaData.ABI

// BindingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingMetaData.Bin instead.
var BindingBin = BindingMetaData.Bin

// DeployBinding deploys a new Ethereum contract, binding an instance of Binding to it.
func DeployBinding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Binding, error) {
	parsed, err := BindingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Binding{BindingCaller: BindingCaller{contract: contract}, BindingTransactor: BindingTransactor{contract: contract}, BindingFilterer: BindingFilterer{contract: contract}}, nil
}

// Binding is an auto generated Go binding around an Ethereum contract.
type Binding struct {
	BindingCaller     // Read-only binding to the contract
	BindingTransactor // Write-only binding to the contract
	BindingFilterer   // Log filterer for contract events
}

// BindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingSession struct {
	Contract     *Binding          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingCallerSession struct {
	Contract *BindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingTransactorSession struct {
	Contract     *BindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingRaw struct {
	Contract *Binding // Generic contract binding to access the raw methods on
}

// BindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingCallerRaw struct {
	Contract *BindingCaller // Generic read-only contract binding to access the raw methods on
}

// BindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingTransactorRaw struct {
	Contract *BindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBinding creates a new instance of Binding, bound to a specific deployed contract.
func NewBinding(address common.Address, backend bind.ContractBackend) (*Binding, error) {
	contract, err := bindBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Binding{BindingCaller: BindingCaller{contract: contract}, BindingTransactor: BindingTransactor{contract: contract}, BindingFilterer: BindingFilterer{contract: contract}}, nil
}

// NewBindingCaller creates a new read-only instance of Binding, bound to a specific deployed contract.
func NewBindingCaller(address common.Address, caller bind.ContractCaller) (*BindingCaller, error) {
	contract, err := bindBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingCaller{contract: contract}, nil
}

// NewBindingTransactor creates a new write-only instance of Binding, bound to a specific deployed contract.
func NewBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingTransactor, error) {
	contract, err := bindBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingTransactor{contract: contract}, nil
}

// NewBindingFilterer creates a new log filterer instance of Binding, bound to a specific deployed contract.
func NewBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingFilterer, error) {
	contract, err := bindBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingFilterer{contract: contract}, nil
}

// bindBinding binds a generic wrapper to an already deployed contract.
func bindBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BindingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Binding *BindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Binding.Contract.BindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Binding *BindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.Contract.BindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Binding *BindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Binding.Contract.BindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Binding *BindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Binding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Binding *BindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Binding *BindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Binding.Contract.contract.Transact(opts, method, params...)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_Binding *BindingCaller) GetDepositCount(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "get_deposit_count")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_Binding *BindingSession) GetDepositCount() ([]byte, error) {
	return _Binding.Contract.GetDepositCount(&_Binding.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_Binding *BindingCallerSession) GetDepositCount() ([]byte, error) {
	return _Binding.Contract.GetDepositCount(&_Binding.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_Binding *BindingCaller) GetDepositRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "get_deposit_root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_Binding *BindingSession) GetDepositRoot() ([32]byte, error) {
	return _Binding.Contract.GetDepositRoot(&_Binding.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_Binding *BindingCallerSession) GetDepositRoot() ([32]byte, error) {
	return _Binding.Contract.GetDepositRoot(&_Binding.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Binding *BindingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Binding *BindingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Binding.Contract.SupportsInterface(&_Binding.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Binding *BindingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Binding.Contract.SupportsInterface(&_Binding.CallOpts, interfaceId)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_Binding *BindingTransactor) Deposit(opts *bind.TransactOpts, pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "deposit", pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_Binding *BindingSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Binding.Contract.Deposit(&_Binding.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_Binding *BindingTransactorSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Binding.Contract.Deposit(&_Binding.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// BindingDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the Binding contract.
type BindingDepositEventIterator struct {
	Event *BindingDepositEvent // Event containing the contract specifics and raw log

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
func (it *BindingDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingDepositEvent)
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
		it.Event = new(BindingDepositEvent)
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
func (it *BindingDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingDepositEvent represents a DepositEvent event raised by the Binding contract.
type BindingDepositEvent struct {
	Pubkey                []byte
	WithdrawalCredentials []byte
	Amount                []byte
	Signature             []byte
	Index                 []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Binding *BindingFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*BindingDepositEventIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &BindingDepositEventIterator{contract: _Binding.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Binding *BindingFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *BindingDepositEvent) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingDepositEvent)
				if err := _Binding.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Binding *BindingFilterer) ParseDepositEvent(log types.Log) (*BindingDepositEvent, error) {
	event := new(BindingDepositEvent)
	if err := _Binding.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
