// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gen

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

// NonProfitABI is the input ABI used to generate the binding from.
const NonProfitABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"Id\",\"type\":\"address\"}],\"name\":\"getRequestsByDonorId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"donorId\",\"type\":\"address\"}],\"name\":\"completeRequest\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"Id\",\"type\":\"address\"}],\"name\":\"getRequestsByVendorId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vendorId\",\"type\":\"address\"},{\"name\":\"nonProfitId\",\"type\":\"address\"}],\"name\":\"createRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"Id\",\"type\":\"address\"}],\"name\":\"getAvailableRequestsByNonProfitId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"Id\",\"type\":\"address\"}],\"name\":\"getRequestsByNonProfitId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRequests\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"Id\",\"type\":\"address\"}],\"name\":\"getClosedRequestsByNonProfitId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RequestCreated\",\"type\":\"event\"}]"

// NonProfitBin is the compiled bytecode used for deploying new contracts.
const NonProfitBin = `0x608060405234801561001057600080fd5b50600080556110d2806100246000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663453f594b811461009d5780634f21a3251461022657806352c92f5e1461023f5780635951a0e11461026057806363849c27146102995780637b4bee53146103dd5780638aea61dc146103fe578063b3685e7414610413578063c6a023e914610434575b600080fd5b3480156100a957600080fd5b506100be600160a060020a036004351661044f565b60405180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b8381101561010e5781810151838201526020016100f6565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b8381101561014d578181015183820152602001610135565b50505050905001868103845289818151815260200191508051906020019060200280838360005b8381101561018c578181015183820152602001610174565b50505050905001868103835288818151815260200191508051906020019060200280838360005b838110156101cb5781810151838201526020016101b3565b50505050905001868103825287818151815260200191508051906020019060200280838360005b8381101561020a5781810151838201526020016101f2565b505050509050019a505050505050505050505060405180910390f35b61023d600435600160a060020a03602435166106b8565b005b34801561024b57600080fd5b506100be600160a060020a036004351661074b565b34801561026c57600080fd5b50610287600160a060020a036004358116906024351661099f565b60408051918252519081900360200190f35b3480156102a557600080fd5b506102ba600160a060020a0360043516610a2d565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b838110156103065781810151838201526020016102ee565b50505050905001858103845288818151815260200191508051906020019060200280838360005b8381101561034557818101518382015260200161032d565b50505050905001858103835287818151815260200191508051906020019060200280838360005b8381101561038457818101518382015260200161036c565b50505050905001858103825286818151815260200191508051906020019060200280838360005b838110156103c35781810151838201526020016103ab565b505050509050019850505050505050505060405180910390f35b3480156103e957600080fd5b506100be600160a060020a0360043516610c3a565b34801561040a57600080fd5b50610287610e8e565b34801561041f57600080fd5b506102ba600160a060020a0360043516610e94565b34801561044057600080fd5b5061023d600435602435611091565b60608060608060608060608060608060008060005460405190808252806020026020018201604052801561048d578160200160208202803883390190505b5096506000546040519080825280602002602001820160405280156104bc578160200160208202803883390190505b5095506000546040519080825280602002602001820160405280156104eb578160200160208202803883390190505b50945060005460405190808252806020026020018201604052801561051a578160200160208202803883390190505b509350600054604051908082528060200260200182016040528015610549578160200160208202803883390190505b50925060009150600090505b6000548110156106a55760008181526001602081905260409091200154600160a060020a038e81169116141561069d5780878381518110151561059457fe5b60209081029091018101919091526000828152600190915260409020548651600160a060020a03909116908790849081106105cb57fe5b600160a060020a039283166020918202909201810191909152600083815260019091526040902060020154865191169086908490811061060757fe5b600160a060020a039092166020928302909101820152600082815260019091526040902060030154845185908490811061063d57fe5b6020908102909101810191909152600082815260019091526040902060040154835160ff9091169084908490811061067157fe5b90602001906020020190600181111561068657fe5b9081600181111561069357fe5b9052506001909101905b600101610555565b50949b939a509198509650945092505050565b600082815260016020819052604080832060048101805460ff191684179055918201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386811691909117909155825460039093015491519216926108fc670de0b6b3a76400009092028015929092029290818181858888f19350505050158015610746573d6000803e3d6000fd5b505050565b606080606080606080606080606080600080600054604051908082528060200260200182016040528015610789578160200160208202803883390190505b5096506000546040519080825280602002602001820160405280156107b8578160200160208202803883390190505b5095506000546040519080825280602002602001820160405280156107e7578160200160208202803883390190505b509450600054604051908082528060200260200182016040528015610816578160200160208202803883390190505b509350600054604051908082528060200260200182016040528015610845578160200160208202803883390190505b50925060009150600090505b6000548110156106a557600081815260016020526040902054600160a060020a038e8116911614156109975780878381518110151561088c57fe5b60209081029091018101919091526000828152600190915260409020600201548651600160a060020a03909116908790849081106108c657fe5b600160a060020a0392831660209182029092018101919091526000838152600191829052604090200154865191169086908490811061090157fe5b600160a060020a039092166020928302909101820152600082815260019091526040902060030154845185908490811061093757fe5b6020908102909101810191909152600082815260019091526040902060040154835160ff9091169084908490811061096b57fe5b90602001906020020190600181111561098057fe5b9081600181111561098d57fe5b9052506001909101905b600101610851565b6000805460018082018355818352602090815260408084208054600160a060020a0380891673ffffffffffffffffffffffffffffffffffffffff199283161783556002909201805492881692909116919091179055805183815290517f97890542421e06ad80deca7ba945fd211a48ec6d89da2510960446025b78e324929181900390910190a19392505050565b606080606080606080606080600080600054604051908082528060200260200182016040528015610a68578160200160208202803883390190505b509550600054604051908082528060200260200182016040528015610a97578160200160208202803883390190505b509450600054604051908082528060200260200182016040528015610ac6578160200160208202803883390190505b509350600054604051908082528060200260200182016040528015610af5578160200160208202803883390190505b50925060009150600090505b600054811015610c2957600081815260016020526040902060020154600160a060020a038c811691161480610b545750600081815260016020819052604082206004015460ff1690811115610b5257fe5b145b15610c2157808683815181101515610b6857fe5b60209081029091018101919091526000828152600190915260409020548551600160a060020a0390911690869084908110610b9f57fe5b600160a060020a03928316602091820290920181019190915260008381526001918290526040902001548551911690859084908110610bda57fe5b600160a060020a0390921660209283029091018201526000828152600190915260409020600301548351849084908110610c1057fe5b602090810290910101526001909101905b600101610b01565b509399929850909650945092505050565b606080606080606080606080606080600080600054604051908082528060200260200182016040528015610c78578160200160208202803883390190505b509650600054604051908082528060200260200182016040528015610ca7578160200160208202803883390190505b509550600054604051908082528060200260200182016040528015610cd6578160200160208202803883390190505b509450600054604051908082528060200260200182016040528015610d05578160200160208202803883390190505b509350600054604051908082528060200260200182016040528015610d34578160200160208202803883390190505b50925060009150600090505b6000548110156106a557600081815260016020526040902060020154600160a060020a038e811691161415610e8657808783815181101515610d7e57fe5b60209081029091018101919091526000828152600190915260409020548651600160a060020a0390911690879084908110610db557fe5b600160a060020a03928316602091820290920181019190915260008381526001918290526040902001548651911690869084908110610df057fe5b600160a060020a0390921660209283029091018201526000828152600190915260409020600301548451859084908110610e2657fe5b6020908102909101810191909152600082815260019091526040902060040154835160ff90911690849084908110610e5a57fe5b906020019060200201906001811115610e6f57fe5b90816001811115610e7c57fe5b9052506001909101905b600101610d40565b60005481565b606080606080606080606080600080600054604051908082528060200260200182016040528015610ecf578160200160208202803883390190505b509550600054604051908082528060200260200182016040528015610efe578160200160208202803883390190505b509450600054604051908082528060200260200182016040528015610f2d578160200160208202803883390190505b509350600054604051908082528060200260200182016040528015610f5c578160200160208202803883390190505b50925060009150600090505b600054811015610c2957600081815260016020526040902060020154600160a060020a038c811691161480610fbc575060008181526001602081905260409091206004015460ff1681811115610fba57fe5b145b1561108957808683815181101515610fd057fe5b60209081029091018101919091526000828152600190915260409020548551600160a060020a039091169086908490811061100757fe5b600160a060020a0392831660209182029092018101919091526000838152600191829052604090200154855191169085908490811061104257fe5b600160a060020a039092166020928302909101820152600082815260019091526040902060030154835184908490811061107857fe5b602090810290910101526001909101905b600101610f68565b600091825260016020526040909120600301555600a165627a7a72305820c59a3384b52573571d9f62b726d4b2fa6a17430c2c32d4a54df3c0749a2b38590029`

// DeployNonProfit deploys a new Ethereum contract, binding an instance of NonProfit to it.
func DeployNonProfit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NonProfit, error) {
	parsed, err := abi.JSON(strings.NewReader(NonProfitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NonProfitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NonProfit{NonProfitCaller: NonProfitCaller{contract: contract}, NonProfitTransactor: NonProfitTransactor{contract: contract}, NonProfitFilterer: NonProfitFilterer{contract: contract}}, nil
}

// NonProfit is an auto generated Go binding around an Ethereum contract.
type NonProfit struct {
	NonProfitCaller     // Read-only binding to the contract
	NonProfitTransactor // Write-only binding to the contract
	NonProfitFilterer   // Log filterer for contract events
}

// NonProfitCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonProfitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonProfitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonProfitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonProfitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonProfitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonProfitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonProfitSession struct {
	Contract     *NonProfit        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NonProfitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonProfitCallerSession struct {
	Contract *NonProfitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NonProfitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonProfitTransactorSession struct {
	Contract     *NonProfitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NonProfitRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonProfitRaw struct {
	Contract *NonProfit // Generic contract binding to access the raw methods on
}

// NonProfitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonProfitCallerRaw struct {
	Contract *NonProfitCaller // Generic read-only contract binding to access the raw methods on
}

// NonProfitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonProfitTransactorRaw struct {
	Contract *NonProfitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonProfit creates a new instance of NonProfit, bound to a specific deployed contract.
func NewNonProfit(address common.Address, backend bind.ContractBackend) (*NonProfit, error) {
	contract, err := bindNonProfit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonProfit{NonProfitCaller: NonProfitCaller{contract: contract}, NonProfitTransactor: NonProfitTransactor{contract: contract}, NonProfitFilterer: NonProfitFilterer{contract: contract}}, nil
}

// NewNonProfitCaller creates a new read-only instance of NonProfit, bound to a specific deployed contract.
func NewNonProfitCaller(address common.Address, caller bind.ContractCaller) (*NonProfitCaller, error) {
	contract, err := bindNonProfit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonProfitCaller{contract: contract}, nil
}

// NewNonProfitTransactor creates a new write-only instance of NonProfit, bound to a specific deployed contract.
func NewNonProfitTransactor(address common.Address, transactor bind.ContractTransactor) (*NonProfitTransactor, error) {
	contract, err := bindNonProfit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonProfitTransactor{contract: contract}, nil
}

// NewNonProfitFilterer creates a new log filterer instance of NonProfit, bound to a specific deployed contract.
func NewNonProfitFilterer(address common.Address, filterer bind.ContractFilterer) (*NonProfitFilterer, error) {
	contract, err := bindNonProfit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonProfitFilterer{contract: contract}, nil
}

// bindNonProfit binds a generic wrapper to an already deployed contract.
func bindNonProfit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NonProfitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonProfit *NonProfitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NonProfit.Contract.NonProfitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonProfit *NonProfitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonProfit.Contract.NonProfitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonProfit *NonProfitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonProfit.Contract.NonProfitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonProfit *NonProfitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NonProfit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonProfit *NonProfitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonProfit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonProfit *NonProfitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonProfit.Contract.contract.Transact(opts, method, params...)
}

// GetAvailableRequestsByNonProfitId is a free data retrieval call binding the contract method 0x63849c27.
//
// Solidity: function getAvailableRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[])
func (_NonProfit *NonProfitCaller) GetAvailableRequestsByNonProfitId(opts *bind.CallOpts, Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
		ret3 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _NonProfit.contract.Call(opts, out, "getAvailableRequestsByNonProfitId", Id)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetAvailableRequestsByNonProfitId is a free data retrieval call binding the contract method 0x63849c27.
//
// Solidity: function getAvailableRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[])
func (_NonProfit *NonProfitSession) GetAvailableRequestsByNonProfitId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, error) {
	return _NonProfit.Contract.GetAvailableRequestsByNonProfitId(&_NonProfit.CallOpts, Id)
}

// GetAvailableRequestsByNonProfitId is a free data retrieval call binding the contract method 0x63849c27.
//
// Solidity: function getAvailableRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[])
func (_NonProfit *NonProfitCallerSession) GetAvailableRequestsByNonProfitId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, error) {
	return _NonProfit.Contract.GetAvailableRequestsByNonProfitId(&_NonProfit.CallOpts, Id)
}

// GetClosedRequestsByNonProfitId is a free data retrieval call binding the contract method 0xb3685e74.
//
// Solidity: function getClosedRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[])
func (_NonProfit *NonProfitCaller) GetClosedRequestsByNonProfitId(opts *bind.CallOpts, Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
		ret3 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _NonProfit.contract.Call(opts, out, "getClosedRequestsByNonProfitId", Id)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetClosedRequestsByNonProfitId is a free data retrieval call binding the contract method 0xb3685e74.
//
// Solidity: function getClosedRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[])
func (_NonProfit *NonProfitSession) GetClosedRequestsByNonProfitId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, error) {
	return _NonProfit.Contract.GetClosedRequestsByNonProfitId(&_NonProfit.CallOpts, Id)
}

// GetClosedRequestsByNonProfitId is a free data retrieval call binding the contract method 0xb3685e74.
//
// Solidity: function getClosedRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[])
func (_NonProfit *NonProfitCallerSession) GetClosedRequestsByNonProfitId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, error) {
	return _NonProfit.Contract.GetClosedRequestsByNonProfitId(&_NonProfit.CallOpts, Id)
}

// GetRequestsByDonorId is a free data retrieval call binding the contract method 0x453f594b.
//
// Solidity: function getRequestsByDonorId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitCaller) GetRequestsByDonorId(opts *bind.CallOpts, Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
		ret3 = new([]*big.Int)
		ret4 = new([]uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _NonProfit.contract.Call(opts, out, "getRequestsByDonorId", Id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetRequestsByDonorId is a free data retrieval call binding the contract method 0x453f594b.
//
// Solidity: function getRequestsByDonorId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitSession) GetRequestsByDonorId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	return _NonProfit.Contract.GetRequestsByDonorId(&_NonProfit.CallOpts, Id)
}

// GetRequestsByDonorId is a free data retrieval call binding the contract method 0x453f594b.
//
// Solidity: function getRequestsByDonorId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitCallerSession) GetRequestsByDonorId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	return _NonProfit.Contract.GetRequestsByDonorId(&_NonProfit.CallOpts, Id)
}

// GetRequestsByNonProfitId is a free data retrieval call binding the contract method 0x7b4bee53.
//
// Solidity: function getRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitCaller) GetRequestsByNonProfitId(opts *bind.CallOpts, Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
		ret3 = new([]*big.Int)
		ret4 = new([]uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _NonProfit.contract.Call(opts, out, "getRequestsByNonProfitId", Id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetRequestsByNonProfitId is a free data retrieval call binding the contract method 0x7b4bee53.
//
// Solidity: function getRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitSession) GetRequestsByNonProfitId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	return _NonProfit.Contract.GetRequestsByNonProfitId(&_NonProfit.CallOpts, Id)
}

// GetRequestsByNonProfitId is a free data retrieval call binding the contract method 0x7b4bee53.
//
// Solidity: function getRequestsByNonProfitId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitCallerSession) GetRequestsByNonProfitId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	return _NonProfit.Contract.GetRequestsByNonProfitId(&_NonProfit.CallOpts, Id)
}

// GetRequestsByVendorId is a free data retrieval call binding the contract method 0x52c92f5e.
//
// Solidity: function getRequestsByVendorId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitCaller) GetRequestsByVendorId(opts *bind.CallOpts, Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
		ret3 = new([]*big.Int)
		ret4 = new([]uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _NonProfit.contract.Call(opts, out, "getRequestsByVendorId", Id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetRequestsByVendorId is a free data retrieval call binding the contract method 0x52c92f5e.
//
// Solidity: function getRequestsByVendorId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitSession) GetRequestsByVendorId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	return _NonProfit.Contract.GetRequestsByVendorId(&_NonProfit.CallOpts, Id)
}

// GetRequestsByVendorId is a free data retrieval call binding the contract method 0x52c92f5e.
//
// Solidity: function getRequestsByVendorId(Id address) constant returns(uint256[], address[], address[], uint256[], uint8[])
func (_NonProfit *NonProfitCallerSession) GetRequestsByVendorId(Id common.Address) ([]*big.Int, []common.Address, []common.Address, []*big.Int, []uint8, error) {
	return _NonProfit.Contract.GetRequestsByVendorId(&_NonProfit.CallOpts, Id)
}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() constant returns(uint256)
func (_NonProfit *NonProfitCaller) TotalRequests(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NonProfit.contract.Call(opts, out, "totalRequests")
	return *ret0, err
}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() constant returns(uint256)
func (_NonProfit *NonProfitSession) TotalRequests() (*big.Int, error) {
	return _NonProfit.Contract.TotalRequests(&_NonProfit.CallOpts)
}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() constant returns(uint256)
func (_NonProfit *NonProfitCallerSession) TotalRequests() (*big.Int, error) {
	return _NonProfit.Contract.TotalRequests(&_NonProfit.CallOpts)
}

// CompleteRequest is a paid mutator transaction binding the contract method 0x4f21a325.
//
// Solidity: function completeRequest(id uint256, donorId address) returns()
func (_NonProfit *NonProfitTransactor) CompleteRequest(opts *bind.TransactOpts, id *big.Int, donorId common.Address) (*types.Transaction, error) {
	return _NonProfit.contract.Transact(opts, "completeRequest", id, donorId)
}

// CompleteRequest is a paid mutator transaction binding the contract method 0x4f21a325.
//
// Solidity: function completeRequest(id uint256, donorId address) returns()
func (_NonProfit *NonProfitSession) CompleteRequest(id *big.Int, donorId common.Address) (*types.Transaction, error) {
	return _NonProfit.Contract.CompleteRequest(&_NonProfit.TransactOpts, id, donorId)
}

// CompleteRequest is a paid mutator transaction binding the contract method 0x4f21a325.
//
// Solidity: function completeRequest(id uint256, donorId address) returns()
func (_NonProfit *NonProfitTransactorSession) CompleteRequest(id *big.Int, donorId common.Address) (*types.Transaction, error) {
	return _NonProfit.Contract.CompleteRequest(&_NonProfit.TransactOpts, id, donorId)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x5951a0e1.
//
// Solidity: function createRequest(vendorId address, nonProfitId address) returns(uint256)
func (_NonProfit *NonProfitTransactor) CreateRequest(opts *bind.TransactOpts, vendorId common.Address, nonProfitId common.Address) (*types.Transaction, error) {
	return _NonProfit.contract.Transact(opts, "createRequest", vendorId, nonProfitId)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x5951a0e1.
//
// Solidity: function createRequest(vendorId address, nonProfitId address) returns(uint256)
func (_NonProfit *NonProfitSession) CreateRequest(vendorId common.Address, nonProfitId common.Address) (*types.Transaction, error) {
	return _NonProfit.Contract.CreateRequest(&_NonProfit.TransactOpts, vendorId, nonProfitId)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x5951a0e1.
//
// Solidity: function createRequest(vendorId address, nonProfitId address) returns(uint256)
func (_NonProfit *NonProfitTransactorSession) CreateRequest(vendorId common.Address, nonProfitId common.Address) (*types.Transaction, error) {
	return _NonProfit.Contract.CreateRequest(&_NonProfit.TransactOpts, vendorId, nonProfitId)
}

// SetAmount is a paid mutator transaction binding the contract method 0xc6a023e9.
//
// Solidity: function setAmount(id uint256, amount uint256) returns()
func (_NonProfit *NonProfitTransactor) SetAmount(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NonProfit.contract.Transact(opts, "setAmount", id, amount)
}

// SetAmount is a paid mutator transaction binding the contract method 0xc6a023e9.
//
// Solidity: function setAmount(id uint256, amount uint256) returns()
func (_NonProfit *NonProfitSession) SetAmount(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NonProfit.Contract.SetAmount(&_NonProfit.TransactOpts, id, amount)
}

// SetAmount is a paid mutator transaction binding the contract method 0xc6a023e9.
//
// Solidity: function setAmount(id uint256, amount uint256) returns()
func (_NonProfit *NonProfitTransactorSession) SetAmount(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NonProfit.Contract.SetAmount(&_NonProfit.TransactOpts, id, amount)
}

// NonProfitRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the NonProfit contract.
type NonProfitRequestCreatedIterator struct {
	Event *NonProfitRequestCreated // Event containing the contract specifics and raw log

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
func (it *NonProfitRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonProfitRequestCreated)
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
		it.Event = new(NonProfitRequestCreated)
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
func (it *NonProfitRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonProfitRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonProfitRequestCreated represents a RequestCreated event raised by the NonProfit contract.
type NonProfitRequestCreated struct {
	*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x97890542421e06ad80deca7ba945fd211a48ec6d89da2510960446025b78e324.
//
// Solidity: e RequestCreated( uint256)
func (_NonProfit *NonProfitFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*NonProfitRequestCreatedIterator, error) {

	logs, sub, err := _NonProfit.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &NonProfitRequestCreatedIterator{contract: _NonProfit.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x97890542421e06ad80deca7ba945fd211a48ec6d89da2510960446025b78e324.
//
// Solidity: e RequestCreated( uint256)
func (_NonProfit *NonProfitFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *NonProfitRequestCreated) (event.Subscription, error) {

	logs, sub, err := _NonProfit.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonProfitRequestCreated)
				if err := _NonProfit.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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
