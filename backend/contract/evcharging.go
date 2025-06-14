// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// EVChargingReserva is an auto generated low-level Go binding around an user-defined struct.
type EVChargingReserva struct {
	Id        *big.Int
	Usuario   common.Address
	VeiculoId string
	PontoId   string
	Inicio    string
	Fim       string
	Preco     *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"}],\"name\":\"NovaReserva\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contador\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"veiculoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pontoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inicio\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fim\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"preco\",\"type\":\"uint256\"}],\"name\":\"criarReserva\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getReserva\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"veiculoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pontoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inicio\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fim\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"preco\",\"type\":\"uint256\"}],\"internalType\":\"structEVCharging.Reserva\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reservas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"veiculoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pontoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inicio\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fim\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"preco\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f5f553480156011575f5ffd5b506110a48061001f5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806353d27cc21461004e5780635ff64a7f1461006c5780636d461b7a1461009c5780639f436853146100b8575b5f5ffd5b6100566100ee565b6040516100639190610809565b60405180910390f35b6100866004803603810190610081919061085d565b6100f3565b60405161009391906109f4565b60405180910390f35b6100b660048036038101906100b19190610b40565b6103c6565b005b6100d260048036038101906100cd919061085d565b61052c565b6040516100e59796959493929190610c7e565b60405180910390f35b5f5481565b6100fb6107a1565b60015f8381526020019081526020015f206040518060e00160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461018590610d34565b80601f01602080910402602001604051908101604052809291908181526020018280546101b190610d34565b80156101fc5780601f106101d3576101008083540402835291602001916101fc565b820191905f5260205f20905b8154815290600101906020018083116101df57829003601f168201915b5050505050815260200160038201805461021590610d34565b80601f016020809104026020016040519081016040528092919081815260200182805461024190610d34565b801561028c5780601f106102635761010080835404028352916020019161028c565b820191905f5260205f20905b81548152906001019060200180831161026f57829003601f168201915b505050505081526020016004820180546102a590610d34565b80601f01602080910402602001604051908101604052809291908181526020018280546102d190610d34565b801561031c5780601f106102f35761010080835404028352916020019161031c565b820191905f5260205f20905b8154815290600101906020018083116102ff57829003601f168201915b5050505050815260200160058201805461033590610d34565b80601f016020809104026020016040519081016040528092919081815260200182805461036190610d34565b80156103ac5780601f10610383576101008083540402835291602001916103ac565b820191905f5260205f20905b81548152906001019060200180831161038f57829003601f168201915b505050505081526020016006820154815250509050919050565b6040518060e001604052805f5481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020018481526020018381526020018281525060015f5f5481526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190816104859190610f04565b50606082015181600301908161049b9190610f04565b5060808201518160040190816104b19190610f04565b5060a08201518160050190816104c79190610f04565b5060c082015181600601559050507f634884580d31ee8959030b0f7f35d6096bb76cd2e4ce002f9a5fc57d995783c35f5433604051610507929190610fd3565b60405180910390a15f5f81548092919061052090611027565b91905055505050505050565b6001602052805f5260405f205f91509050805f015490806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600201805461057690610d34565b80601f01602080910402602001604051908101604052809291908181526020018280546105a290610d34565b80156105ed5780601f106105c4576101008083540402835291602001916105ed565b820191905f5260205f20905b8154815290600101906020018083116105d057829003601f168201915b50505050509080600301805461060290610d34565b80601f016020809104026020016040519081016040528092919081815260200182805461062e90610d34565b80156106795780601f1061065057610100808354040283529160200191610679565b820191905f5260205f20905b81548152906001019060200180831161065c57829003601f168201915b50505050509080600401805461068e90610d34565b80601f01602080910402602001604051908101604052809291908181526020018280546106ba90610d34565b80156107055780601f106106dc57610100808354040283529160200191610705565b820191905f5260205f20905b8154815290600101906020018083116106e857829003601f168201915b50505050509080600501805461071a90610d34565b80601f016020809104026020016040519081016040528092919081815260200182805461074690610d34565b80156107915780601f1061076857610100808354040283529160200191610791565b820191905f5260205f20905b81548152906001019060200180831161077457829003601f168201915b5050505050908060060154905087565b6040518060e001604052805f81526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001606081526020015f81525090565b5f819050919050565b610803816107f1565b82525050565b5f60208201905061081c5f8301846107fa565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b61083c816107f1565b8114610846575f5ffd5b50565b5f8135905061085781610833565b92915050565b5f602082840312156108725761087161082b565b5b5f61087f84828501610849565b91505092915050565b610891816107f1565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108c082610897565b9050919050565b6108d0816108b6565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610918826108d6565b61092281856108e0565b93506109328185602086016108f0565b61093b816108fe565b840191505092915050565b5f60e083015f83015161095b5f860182610888565b50602083015161096e60208601826108c7565b5060408301518482036040860152610986828261090e565b915050606083015184820360608601526109a0828261090e565b915050608083015184820360808601526109ba828261090e565b91505060a083015184820360a08601526109d4828261090e565b91505060c08301516109e960c0860182610888565b508091505092915050565b5f6020820190508181035f830152610a0c8184610946565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610a52826108fe565b810181811067ffffffffffffffff82111715610a7157610a70610a1c565b5b80604052505050565b5f610a83610822565b9050610a8f8282610a49565b919050565b5f67ffffffffffffffff821115610aae57610aad610a1c565b5b610ab7826108fe565b9050602081019050919050565b828183375f83830152505050565b5f610ae4610adf84610a94565b610a7a565b905082815260208101848484011115610b0057610aff610a18565b5b610b0b848285610ac4565b509392505050565b5f82601f830112610b2757610b26610a14565b5b8135610b37848260208601610ad2565b91505092915050565b5f5f5f5f5f60a08688031215610b5957610b5861082b565b5b5f86013567ffffffffffffffff811115610b7657610b7561082f565b5b610b8288828901610b13565b955050602086013567ffffffffffffffff811115610ba357610ba261082f565b5b610baf88828901610b13565b945050604086013567ffffffffffffffff811115610bd057610bcf61082f565b5b610bdc88828901610b13565b935050606086013567ffffffffffffffff811115610bfd57610bfc61082f565b5b610c0988828901610b13565b9250506080610c1a88828901610849565b9150509295509295909350565b610c30816108b6565b82525050565b5f82825260208201905092915050565b5f610c50826108d6565b610c5a8185610c36565b9350610c6a8185602086016108f0565b610c73816108fe565b840191505092915050565b5f60e082019050610c915f83018a6107fa565b610c9e6020830189610c27565b8181036040830152610cb08188610c46565b90508181036060830152610cc48187610c46565b90508181036080830152610cd88186610c46565b905081810360a0830152610cec8185610c46565b9050610cfb60c08301846107fa565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610d4b57607f821691505b602082108103610d5e57610d5d610d07565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610dc07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d85565b610dca8683610d85565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610e05610e00610dfb846107f1565b610de2565b6107f1565b9050919050565b5f819050919050565b610e1e83610deb565b610e32610e2a82610e0c565b848454610d91565b825550505050565b5f5f905090565b610e49610e3a565b610e54818484610e15565b505050565b5b81811015610e7757610e6c5f82610e41565b600181019050610e5a565b5050565b601f821115610ebc57610e8d81610d64565b610e9684610d76565b81016020851015610ea5578190505b610eb9610eb185610d76565b830182610e59565b50505b505050565b5f82821c905092915050565b5f610edc5f1984600802610ec1565b1980831691505092915050565b5f610ef48383610ecd565b9150826002028217905092915050565b610f0d826108d6565b67ffffffffffffffff811115610f2657610f25610a1c565b5b610f308254610d34565b610f3b828285610e7b565b5f60209050601f831160018114610f6c575f8415610f5a578287015190505b610f648582610ee9565b865550610fcb565b601f198416610f7a86610d64565b5f5b82811015610fa157848901518255600182019150602085019450602081019050610f7c565b86831015610fbe5784890151610fba601f891682610ecd565b8355505b6001600288020188555050505b505050505050565b5f604082019050610fe65f8301856107fa565b610ff36020830184610c27565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611031826107f1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361106357611062610ffa565b5b60018201905091905056fea2646970667358221220f69afd18538960877558876c1a737abdaddef828ca5646116efbc7617af09f3464736f6c634300081e0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
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
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// Contador is a free data retrieval call binding the contract method 0x53d27cc2.
//
// Solidity: function contador() view returns(uint256)
func (_Contract *ContractCaller) Contador(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "contador")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contador is a free data retrieval call binding the contract method 0x53d27cc2.
//
// Solidity: function contador() view returns(uint256)
func (_Contract *ContractSession) Contador() (*big.Int, error) {
	return _Contract.Contract.Contador(&_Contract.CallOpts)
}

// Contador is a free data retrieval call binding the contract method 0x53d27cc2.
//
// Solidity: function contador() view returns(uint256)
func (_Contract *ContractCallerSession) Contador() (*big.Int, error) {
	return _Contract.Contract.Contador(&_Contract.CallOpts)
}

// GetReserva is a free data retrieval call binding the contract method 0x5ff64a7f.
//
// Solidity: function getReserva(uint256 id) view returns((uint256,address,string,string,string,string,uint256))
func (_Contract *ContractCaller) GetReserva(opts *bind.CallOpts, id *big.Int) (EVChargingReserva, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getReserva", id)

	if err != nil {
		return *new(EVChargingReserva), err
	}

	out0 := *abi.ConvertType(out[0], new(EVChargingReserva)).(*EVChargingReserva)

	return out0, err

}

// GetReserva is a free data retrieval call binding the contract method 0x5ff64a7f.
//
// Solidity: function getReserva(uint256 id) view returns((uint256,address,string,string,string,string,uint256))
func (_Contract *ContractSession) GetReserva(id *big.Int) (EVChargingReserva, error) {
	return _Contract.Contract.GetReserva(&_Contract.CallOpts, id)
}

// GetReserva is a free data retrieval call binding the contract method 0x5ff64a7f.
//
// Solidity: function getReserva(uint256 id) view returns((uint256,address,string,string,string,string,uint256))
func (_Contract *ContractCallerSession) GetReserva(id *big.Int) (EVChargingReserva, error) {
	return _Contract.Contract.GetReserva(&_Contract.CallOpts, id)
}

// Reservas is a free data retrieval call binding the contract method 0x9f436853.
//
// Solidity: function reservas(uint256 ) view returns(uint256 id, address usuario, string veiculoId, string pontoId, string inicio, string fim, uint256 preco)
func (_Contract *ContractCaller) Reservas(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Usuario   common.Address
	VeiculoId string
	PontoId   string
	Inicio    string
	Fim       string
	Preco     *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reservas", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Usuario   common.Address
		VeiculoId string
		PontoId   string
		Inicio    string
		Fim       string
		Preco     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Usuario = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VeiculoId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PontoId = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Inicio = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Fim = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Preco = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Reservas is a free data retrieval call binding the contract method 0x9f436853.
//
// Solidity: function reservas(uint256 ) view returns(uint256 id, address usuario, string veiculoId, string pontoId, string inicio, string fim, uint256 preco)
func (_Contract *ContractSession) Reservas(arg0 *big.Int) (struct {
	Id        *big.Int
	Usuario   common.Address
	VeiculoId string
	PontoId   string
	Inicio    string
	Fim       string
	Preco     *big.Int
}, error) {
	return _Contract.Contract.Reservas(&_Contract.CallOpts, arg0)
}

// Reservas is a free data retrieval call binding the contract method 0x9f436853.
//
// Solidity: function reservas(uint256 ) view returns(uint256 id, address usuario, string veiculoId, string pontoId, string inicio, string fim, uint256 preco)
func (_Contract *ContractCallerSession) Reservas(arg0 *big.Int) (struct {
	Id        *big.Int
	Usuario   common.Address
	VeiculoId string
	PontoId   string
	Inicio    string
	Fim       string
	Preco     *big.Int
}, error) {
	return _Contract.Contract.Reservas(&_Contract.CallOpts, arg0)
}

// CriarReserva is a paid mutator transaction binding the contract method 0x6d461b7a.
//
// Solidity: function criarReserva(string veiculoId, string pontoId, string inicio, string fim, uint256 preco) returns()
func (_Contract *ContractTransactor) CriarReserva(opts *bind.TransactOpts, veiculoId string, pontoId string, inicio string, fim string, preco *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "criarReserva", veiculoId, pontoId, inicio, fim, preco)
}

// CriarReserva is a paid mutator transaction binding the contract method 0x6d461b7a.
//
// Solidity: function criarReserva(string veiculoId, string pontoId, string inicio, string fim, uint256 preco) returns()
func (_Contract *ContractSession) CriarReserva(veiculoId string, pontoId string, inicio string, fim string, preco *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CriarReserva(&_Contract.TransactOpts, veiculoId, pontoId, inicio, fim, preco)
}

// CriarReserva is a paid mutator transaction binding the contract method 0x6d461b7a.
//
// Solidity: function criarReserva(string veiculoId, string pontoId, string inicio, string fim, uint256 preco) returns()
func (_Contract *ContractTransactorSession) CriarReserva(veiculoId string, pontoId string, inicio string, fim string, preco *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CriarReserva(&_Contract.TransactOpts, veiculoId, pontoId, inicio, fim, preco)
}

// ContractNovaReservaIterator is returned from FilterNovaReserva and is used to iterate over the raw logs and unpacked data for NovaReserva events raised by the Contract contract.
type ContractNovaReservaIterator struct {
	Event *ContractNovaReserva // Event containing the contract specifics and raw log

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
func (it *ContractNovaReservaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNovaReserva)
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
		it.Event = new(ContractNovaReserva)
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
func (it *ContractNovaReservaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNovaReservaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNovaReserva represents a NovaReserva event raised by the Contract contract.
type ContractNovaReserva struct {
	Id      *big.Int
	Usuario common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNovaReserva is a free log retrieval operation binding the contract event 0x634884580d31ee8959030b0f7f35d6096bb76cd2e4ce002f9a5fc57d995783c3.
//
// Solidity: event NovaReserva(uint256 id, address usuario)
func (_Contract *ContractFilterer) FilterNovaReserva(opts *bind.FilterOpts) (*ContractNovaReservaIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NovaReserva")
	if err != nil {
		return nil, err
	}
	return &ContractNovaReservaIterator{contract: _Contract.contract, event: "NovaReserva", logs: logs, sub: sub}, nil
}

// WatchNovaReserva is a free log subscription operation binding the contract event 0x634884580d31ee8959030b0f7f35d6096bb76cd2e4ce002f9a5fc57d995783c3.
//
// Solidity: event NovaReserva(uint256 id, address usuario)
func (_Contract *ContractFilterer) WatchNovaReserva(opts *bind.WatchOpts, sink chan<- *ContractNovaReserva) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NovaReserva")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNovaReserva)
				if err := _Contract.contract.UnpackLog(event, "NovaReserva", log); err != nil {
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

// ParseNovaReserva is a log parse operation binding the contract event 0x634884580d31ee8959030b0f7f35d6096bb76cd2e4ce002f9a5fc57d995783c3.
//
// Solidity: event NovaReserva(uint256 id, address usuario)
func (_Contract *ContractFilterer) ParseNovaReserva(log types.Log) (*ContractNovaReserva, error) {
	event := new(ContractNovaReserva)
	if err := _Contract.contract.UnpackLog(event, "NovaReserva", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
