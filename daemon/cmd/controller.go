// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"core\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"name\":\"defund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liqdebt\",\"type\":\"uint256\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"}],\"name\":\"setCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setCollateralPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"core\",\"type\":\"address\"}],\"name\":\"setCore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200155438038062001554833981810160405281019062000037919062000319565b620000576200004b6200008160201b60201c565b6200008960201b60201c565b62000068826200014d60201b60201c565b6200007981620001a160201b60201c565b5050620003e3565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6200015d620001f560201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b620001b1620001f560201b60201c565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b620002056200008160201b60201c565b73ffffffffffffffffffffffffffffffffffffffff166200022b6200028660201b60201c565b73ffffffffffffffffffffffffffffffffffffffff161462000284576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200027b90620003c1565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002e182620002b4565b9050919050565b620002f381620002d4565b8114620002ff57600080fd5b50565b6000815190506200031381620002e8565b92915050565b60008060408385031215620003335762000332620002af565b5b6000620003438582860162000302565b9250506020620003568582860162000302565b9150509250929050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000620003a960208362000360565b9150620003b68262000371565b602082019050919050565b60006020820190508181036000830152620003dc816200039a565b9050919050565b61116180620003f36000396000f3fe6080604052600436106100865760003560e01c80638da5cb5b116100595780638da5cb5b14610110578063a1865cb51461013b578063c397ae3914610164578063ca1d209d14610180578063f2fde38b1461019c57610086565b80630203d8fb1461008b578063715018a6146100a757806380009630146100be57806382e46b75146100e7575b600080fd5b6100a560048036038101906100a09190610c47565b6101c5565b005b3480156100b357600080fd5b506100bc6104d4565b005b3480156100ca57600080fd5b506100e560048036038101906100e09190610c87565b6104e8565b005b3480156100f357600080fd5b5061010e60048036038101906101099190610c87565b610534565b005b34801561011c57600080fd5b50610125610580565b6040516101329190610cc3565b60405180910390f35b34801561014757600080fd5b50610162600480360381019061015d9190610cde565b6105a9565b005b61017e60048036038101906101799190610d0b565b610639565b005b61019a60048036038101906101959190610cde565b6108a0565b005b3480156101a857600080fd5b506101c360048036038101906101be9190610c87565b6109e1565b005b60003390506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff1660e01b81526004016102279190610cc3565b602060405180830381865afa158015610244573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102689190610d60565b9050828110156102ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a490610dea565b60405180910390fd5b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166326c013038689896040518463ffffffff1660e01b815260040161031193929190610e19565b6060604051808303816000875af1158015610330573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103549190610e50565b809350819450829550505050808610156103a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039a90610eef565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156103e9573d6000803e3d6000fd5b50600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac8683896104359190610f3e565b6040518363ffffffff1660e01b8152600401610452929190610f72565b600060405180830381600087803b15801561046c57600080fd5b505af1158015610480573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156104ca573d6000803e3d6000fd5b5050505050505050565b6104dc610a64565b6104e66000610ae2565b565b6104f0610a64565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61053c610a64565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a1865cb5826040518263ffffffff1660e01b81526004016106049190610f9b565b600060405180830381600087803b15801561061e57600080fd5b505af1158015610632573d6000803e3d6000fd5b5050505050565b60003390506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff1660e01b815260040161069b9190610cc3565b602060405180830381865afa1580156106b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106dc9190610d60565b905082811015610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071890610dea565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635ff657ed8386866040518463ffffffff1660e01b815260040161078093929190610fb6565b60408051808303816000875af115801561079e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c29190610fed565b5050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac83856040518363ffffffff1660e01b8152600401610821929190610f72565b600060405180830381600087803b15801561083b57600080fd5b505af115801561084f573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f19350505050158015610899573d6000803e3d6000fd5b5050505050565b60003390506000349050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630af6ce858383866040518463ffffffff1660e01b815260040161090993929190610fb6565b60408051808303816000875af1158015610927573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094b9190610fed565b5050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983856040518363ffffffff1660e01b81526004016109aa929190610f72565b600060405180830381600087803b1580156109c457600080fd5b505af11580156109d8573d6000803e3d6000fd5b50505050505050565b6109e9610a64565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4f9061109f565b60405180910390fd5b610a6181610ae2565b50565b610a6c610ba6565b73ffffffffffffffffffffffffffffffffffffffff16610a8a610580565b73ffffffffffffffffffffffffffffffffffffffff1614610ae0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad79061110b565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610bde82610bb3565b9050919050565b610bee81610bd3565b8114610bf957600080fd5b50565b600081359050610c0b81610be5565b92915050565b6000819050919050565b610c2481610c11565b8114610c2f57600080fd5b50565b600081359050610c4181610c1b565b92915050565b60008060408385031215610c5e57610c5d610bae565b5b6000610c6c85828601610bfc565b9250506020610c7d85828601610c32565b9150509250929050565b600060208284031215610c9d57610c9c610bae565b5b6000610cab84828501610bfc565b91505092915050565b610cbd81610bd3565b82525050565b6000602082019050610cd86000830184610cb4565b92915050565b600060208284031215610cf457610cf3610bae565b5b6000610d0284828501610c32565b91505092915050565b60008060408385031215610d2257610d21610bae565b5b6000610d3085828601610c32565b9250506020610d4185828601610c32565b9150509250929050565b600081519050610d5a81610c1b565b92915050565b600060208284031215610d7657610d75610bae565b5b6000610d8484828501610d4b565b91505092915050565b600082825260208201905092915050565b7f436f6e74726f6c6c65722f496e73756666696369656e742062616c616e636500600082015250565b6000610dd4601f83610d8d565b9150610ddf82610d9e565b602082019050919050565b60006020820190508181036000830152610e0381610dc7565b9050919050565b610e1381610c11565b82525050565b6000606082019050610e2e6000830186610cb4565b610e3b6020830185610cb4565b610e486040830184610e0a565b949350505050565b600080600060608486031215610e6957610e68610bae565b5b6000610e7786828701610d4b565b9350506020610e8886828701610d4b565b9250506040610e9986828701610d4b565b9150509250925092565b7f436f6e74726f6c6c65722f4e6567617469766520646562746f72000000000000600082015250565b6000610ed9601a83610d8d565b9150610ee482610ea3565b602082019050919050565b60006020820190508181036000830152610f0881610ecc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610f4982610c11565b9150610f5483610c11565b9250828203905081811115610f6c57610f6b610f0f565b5b92915050565b6000604082019050610f876000830185610cb4565b610f946020830184610e0a565b9392505050565b6000602082019050610fb06000830184610e0a565b92915050565b6000606082019050610fcb6000830186610cb4565b610fd86020830185610e0a565b610fe56040830184610e0a565b949350505050565b6000806040838503121561100457611003610bae565b5b600061101285828601610d4b565b925050602061102385828601610d4b565b9150509250929050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611089602683610d8d565b91506110948261102d565b604082019050919050565b600060208201905081810360008301526110b88161107c565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006110f5602083610d8d565b9150611100826110bf565b602082019050919050565b60006020820190508181036000830152611124816110e8565b905091905056fea2646970667358221220fe3ab8fbac9cda8760e9799d325df0b43295bc31a59615870ed989a34a791dd364736f6c63430008110033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, core common.Address, coin common.Address) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, core, coin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Defund is a paid mutator transaction binding the contract method 0xc397ae39.
//
// Solidity: function defund(uint256 collateral, uint256 debt) payable returns()
func (_Contract *ContractTransactor) Defund(opts *bind.TransactOpts, collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "defund", collateral, debt)
}

// Defund is a paid mutator transaction binding the contract method 0xc397ae39.
//
// Solidity: function defund(uint256 collateral, uint256 debt) payable returns()
func (_Contract *ContractSession) Defund(collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Defund(&_Contract.TransactOpts, collateral, debt)
}

// Defund is a paid mutator transaction binding the contract method 0xc397ae39.
//
// Solidity: function defund(uint256 collateral, uint256 debt) payable returns()
func (_Contract *ContractTransactorSession) Defund(collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Defund(&_Contract.TransactOpts, collateral, debt)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 debt) payable returns()
func (_Contract *ContractTransactor) Fund(opts *bind.TransactOpts, debt *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fund", debt)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 debt) payable returns()
func (_Contract *ContractSession) Fund(debt *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Fund(&_Contract.TransactOpts, debt)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 debt) payable returns()
func (_Contract *ContractTransactorSession) Fund(debt *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Fund(&_Contract.TransactOpts, debt)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address user, uint256 liqdebt) payable returns()
func (_Contract *ContractTransactor) Kick(opts *bind.TransactOpts, user common.Address, liqdebt *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "kick", user, liqdebt)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address user, uint256 liqdebt) payable returns()
func (_Contract *ContractSession) Kick(user common.Address, liqdebt *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Kick(&_Contract.TransactOpts, user, liqdebt)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address user, uint256 liqdebt) payable returns()
func (_Contract *ContractTransactorSession) Kick(user common.Address, liqdebt *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Kick(&_Contract.TransactOpts, user, liqdebt)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetCoin is a paid mutator transaction binding the contract method 0x82e46b75.
//
// Solidity: function setCoin(address coin) returns()
func (_Contract *ContractTransactor) SetCoin(opts *bind.TransactOpts, coin common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCoin", coin)
}

// SetCoin is a paid mutator transaction binding the contract method 0x82e46b75.
//
// Solidity: function setCoin(address coin) returns()
func (_Contract *ContractSession) SetCoin(coin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCoin(&_Contract.TransactOpts, coin)
}

// SetCoin is a paid mutator transaction binding the contract method 0x82e46b75.
//
// Solidity: function setCoin(address coin) returns()
func (_Contract *ContractTransactorSession) SetCoin(coin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCoin(&_Contract.TransactOpts, coin)
}

// SetCollateralPrice is a paid mutator transaction binding the contract method 0xa1865cb5.
//
// Solidity: function setCollateralPrice(uint256 price) returns()
func (_Contract *ContractTransactor) SetCollateralPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCollateralPrice", price)
}

// SetCollateralPrice is a paid mutator transaction binding the contract method 0xa1865cb5.
//
// Solidity: function setCollateralPrice(uint256 price) returns()
func (_Contract *ContractSession) SetCollateralPrice(price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetCollateralPrice(&_Contract.TransactOpts, price)
}

// SetCollateralPrice is a paid mutator transaction binding the contract method 0xa1865cb5.
//
// Solidity: function setCollateralPrice(uint256 price) returns()
func (_Contract *ContractTransactorSession) SetCollateralPrice(price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetCollateralPrice(&_Contract.TransactOpts, price)
}

// SetCore is a paid mutator transaction binding the contract method 0x80009630.
//
// Solidity: function setCore(address core) returns()
func (_Contract *ContractTransactor) SetCore(opts *bind.TransactOpts, core common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCore", core)
}

// SetCore is a paid mutator transaction binding the contract method 0x80009630.
//
// Solidity: function setCore(address core) returns()
func (_Contract *ContractSession) SetCore(core common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCore(&_Contract.TransactOpts, core)
}

// SetCore is a paid mutator transaction binding the contract method 0x80009630.
//
// Solidity: function setCore(address core) returns()
func (_Contract *ContractTransactorSession) SetCore(core common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCore(&_Contract.TransactOpts, core)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
