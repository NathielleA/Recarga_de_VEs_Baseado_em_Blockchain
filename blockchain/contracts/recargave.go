// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package recargave

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

// RecargaVESimplificadoSessao is an auto generated low-level Go binding around an user-defined struct.
type RecargaVESimplificadoSessao struct {
	IdSessao              *big.Int
	IdPosto               *big.Int
	EnderecoVE            common.Address
	HorarioAgendadoInicio *big.Int
	HorarioAgendadoFim    *big.Int
	HorarioRealInicio     *big.Int
	EnergiaConsumida      *big.Int
	CustoTotal            *big.Int
	Status                uint8
}

// RecargaveMetaData contains all meta data concerning the Recargave contract.
var RecargaveMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idPosto\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proprietario\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nome\",\"type\":\"string\"}],\"name\":\"PostoRegistrado\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idSessao\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"energiaConsumida\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"custoTotal\",\"type\":\"uint256\"}],\"name\":\"RecargaFinalizada\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idSessao\",\"type\":\"uint256\"}],\"name\":\"RecargaIniciada\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idSessao\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idPosto\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"enderecoVE\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"horarioInicio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"horarioFim\",\"type\":\"uint256\"}],\"name\":\"SessaoAgendada\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idSessao\",\"type\":\"uint256\"}],\"name\":\"SessaoCancelada\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idPosto\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_horarioInicio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duracaoMinutos\",\"type\":\"uint256\"}],\"name\":\"agendarSessao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idSessao\",\"type\":\"uint256\"}],\"name\":\"cancelarSessao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idSessao\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_energiaConsumida\",\"type\":\"uint256\"}],\"name\":\"finalizarRecarga\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idSessao\",\"type\":\"uint256\"}],\"name\":\"getSessao\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"idSessao\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idPosto\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"enderecoVE\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"horarioAgendadoInicio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"horarioAgendadoFim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"horarioRealInicio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"energiaConsumida\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"custoTotal\",\"type\":\"uint256\"},{\"internalType\":\"enumRecargaVESimplificado.StatusSessao\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structRecargaVESimplificado.Sessao\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idSessao\",\"type\":\"uint256\"}],\"name\":\"iniciarRecarga\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"postos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"idPosto\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proprietario\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nome\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"localizacao\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"precoPorUnidadeEnergia\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proximoIdPosto\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proximoIdSessao\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nome\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_localizacao\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_precoPorUnidadeEnergia\",\"type\":\"uint256\"}],\"name\":\"registrarPosto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessoes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"idSessao\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idPosto\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"enderecoVE\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"horarioAgendadoInicio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"horarioAgendadoFim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"horarioRealInicio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"energiaConsumida\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"custoTotal\",\"type\":\"uint256\"},{\"internalType\":\"enumRecargaVESimplificado.StatusSessao\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessoesPorPosto\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f55600180553480156016575f5ffd5b50611da4806100245f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c80638aeedc251161006f5780638aeedc2514610179578063a718ae9614610195578063a97ab930146101cd578063b53d3e25146101e9578063c8f947c114610205578063cd90c06f14610223576100a7565b80630a915a10146100ab5780631735301e146100c75780631da099e6146100f7578063241cb514146101155780632ab5096914610145575b5f5ffd5b6100c560048036038101906100c0919061109c565b61023f565b005b6100e160048036038101906100dc91906110da565b610408565b6040516100ee919061127b565b60405180910390f35b6100ff610508565b60405161010c91906112a4565b60405180910390f35b61012f600480360381019061012a919061109c565b61050d565b60405161013c91906112a4565b60405180910390f35b61015f600480360381019061015a91906110da565b610538565b60405161017095949392919061133c565b60405180910390f35b610193600480360381019061018e91906110da565b610695565b005b6101af60048036038101906101aa91906110da565b61086d565b6040516101c4999897969594939291906113aa565b60405180910390f35b6101e760048036038101906101e291906110da565b6108e2565b005b61020360048036038101906101fe9190611435565b610a5a565b005b61020d610e24565b60405161021a91906112a4565b60405180910390f35b61023d600480360381019061023891906114e6565b610e2a565b005b5f60035f8481526020019081526020015f2090505f60025f836001015481526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16826002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f3906115c1565b60405180910390fd5b600160038111156103105761030f611153565b5b826008015f9054906101000a900460ff16600381111561033357610332611153565b5b14610373576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036a90611629565b60405180910390fd5b82826006018190555080600401548361038c9190611674565b82600701819055506002826008015f6101000a81548160ff021916908360038111156103bb576103ba611153565b5b0217905550837fa7a1530868b4f4e1e8a987d23325f7bfc66dfaf15b746dd9bfc11c0b917c1865836006015484600701546040516103fa9291906116b5565b60405180910390a250505050565b610410610ff6565b60035f8381526020019081526020015f20604051806101200160405290815f820154815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152602001600882015f9054906101000a900460ff1660038111156104eb576104ea611153565b5b60038111156104fd576104fc611153565b5b815250509050919050565b5f5481565b6004602052815f5260405f208181548110610526575f80fd5b905f5260205f20015f91509150505481565b6002602052805f5260405f205f91509050805f015490806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600201805461058290611709565b80601f01602080910402602001604051908101604052809291908181526020018280546105ae90611709565b80156105f95780601f106105d0576101008083540402835291602001916105f9565b820191905f5260205f20905b8154815290600101906020018083116105dc57829003601f168201915b50505050509080600301805461060e90611709565b80601f016020809104026020016040519081016040528092919081815260200182805461063a90611709565b80156106855780601f1061065c57610100808354040283529160200191610685565b820191905f5260205f20905b81548152906001019060200180831161066857829003601f168201915b5050505050908060040154905085565b5f60035f8381526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16816002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461073a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610731906115c1565b60405180910390fd5b5f600381111561074d5761074c611153565b5b816008015f9054906101000a900460ff1660038111156107705761076f611153565b5b146107b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a790611783565b60405180910390fd5b806003015442101580156107c8575080600401544211155b610807576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fe906117eb565b60405180910390fd5b6001816008015f6101000a81548160ff0219169083600381111561082e5761082d611153565b5b0217905550428160050181905550817f41a1f1d6bcb91629ef55145436e5e4b93c0b96ecb4d80c51f27b6b8ceeb7857660405160405180910390a25050565b6003602052805f5260405f205f91509050805f015490806001015490806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003015490806004015490806005015490806006015490806007015490806008015f9054906101000a900460ff16905089565b5f60035f8381526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16816002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610987576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097e906115c1565b60405180910390fd5b5f600381111561099a57610999611153565b5b816008015f9054906101000a900460ff1660038111156109bd576109bc611153565b5b146109fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f490611879565b60405180910390fd5b6003816008015f6101000a81548160ff02191690836003811115610a2457610a23611153565b5b0217905550817fa26ee77645ce5364f41aa0394e99dab9dc93eb68e80e149e6ce3a3394c686b2860405160405180910390a25050565b5f60025f8581526020019081526020015f205f015403610aaf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa6906118e1565b60405180910390fd5b5f603c82610abd9190611674565b83610ac891906118ff565b9050828111610b0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b039061197c565b60405180910390fd5b5f60045f8681526020019081526020015f2090505f5f90505b8180549050811015610c45575f60035f848481548110610b4857610b4761199a565b5b905f5260205f20015481526020019081526020015f2090505f6003811115610b7357610b72611153565b5b816008015f9054906101000a900460ff166003811115610b9657610b95611153565b5b1480610bd6575060016003811115610bb157610bb0611153565b5b816008015f9054906101000a900460ff166003811115610bd457610bd3611153565b5b145b15610c37575f816004015487108015610bf25750816003015485115b90508015610c35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2c90611a11565b60405180910390fd5b505b508080600101915050610b25565b505f60015f815480929190610c5990611a2f565b9190505590506040518061012001604052808281526020018781526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018481526020015f81526020015f81526020015f81526020015f6003811115610cc457610cc3611153565b5b81525060035f8381526020019081526020015f205f820151815f0155602082015181600101556040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070155610100820151816008015f6101000a81548160ff02191690836003811115610d8d57610d8c611153565b5b021790555090505060045f8781526020019081526020015f2081908060018154018082558091505060019003905f5260205f20015f90919091909150553373ffffffffffffffffffffffffffffffffffffffff1686827f6c55c2668ba44bf59ad5f2dab96ddc6d4537864bdfa5895056ecb11b55d3bb848887604051610e149291906116b5565b60405180910390a4505050505050565b60015481565b5f5f5f815480929190610e3c90611a2f565b9190505590506040518060a001604052808281526020013373ffffffffffffffffffffffffffffffffffffffff16815260200187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018381525060025f8381526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610f799190611c43565b506060820151816003019081610f8f9190611c43565b50608082015181600401559050503373ffffffffffffffffffffffffffffffffffffffff16817fc462caf9f0889b3d4dfe9d08c0f9cd1f1e64c5bb5f500d061e5d9475a13cdd908888604051610fe6929190611d4c565b60405180910390a3505050505050565b6040518061012001604052805f81526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f600381111561105b5761105a611153565b5b81525090565b5f5ffd5b5f5ffd5b5f819050919050565b61107b81611069565b8114611085575f5ffd5b50565b5f8135905061109681611072565b92915050565b5f5f604083850312156110b2576110b1611061565b5b5f6110bf85828601611088565b92505060206110d085828601611088565b9150509250929050565b5f602082840312156110ef576110ee611061565b5b5f6110fc84828501611088565b91505092915050565b61110e81611069565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61113d82611114565b9050919050565b61114d81611133565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6004811061119157611190611153565b5b50565b5f8190506111a182611180565b919050565b5f6111b082611194565b9050919050565b6111c0816111a6565b82525050565b61012082015f8201516111db5f850182611105565b5060208201516111ee6020850182611105565b5060408201516112016040850182611144565b5060608201516112146060850182611105565b5060808201516112276080850182611105565b5060a082015161123a60a0850182611105565b5060c082015161124d60c0850182611105565b5060e082015161126060e0850182611105565b506101008201516112756101008501826111b7565b50505050565b5f6101208201905061128f5f8301846111c6565b92915050565b61129e81611069565b82525050565b5f6020820190506112b75f830184611295565b92915050565b6112c681611133565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61130e826112cc565b61131881856112d6565b93506113288185602086016112e6565b611331816112f4565b840191505092915050565b5f60a08201905061134f5f830188611295565b61135c60208301876112bd565b818103604083015261136e8186611304565b905081810360608301526113828185611304565b90506113916080830184611295565b9695505050505050565b6113a4816111a6565b82525050565b5f610120820190506113be5f83018c611295565b6113cb602083018b611295565b6113d8604083018a6112bd565b6113e56060830189611295565b6113f26080830188611295565b6113ff60a0830187611295565b61140c60c0830186611295565b61141960e0830185611295565b61142761010083018461139b565b9a9950505050505050505050565b5f5f5f6060848603121561144c5761144b611061565b5b5f61145986828701611088565b935050602061146a86828701611088565b925050604061147b86828701611088565b9150509250925092565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126114a6576114a5611485565b5b8235905067ffffffffffffffff8111156114c3576114c2611489565b5b6020830191508360018202830111156114df576114de61148d565b5b9250929050565b5f5f5f5f5f606086880312156114ff576114fe611061565b5b5f86013567ffffffffffffffff81111561151c5761151b611065565b5b61152888828901611491565b9550955050602086013567ffffffffffffffff81111561154b5761154a611065565b5b61155788828901611491565b9350935050604061156a88828901611088565b9150509295509295909350565b7f4e616f206175746f72697a61646f2e00000000000000000000000000000000005f82015250565b5f6115ab600f836112d6565b91506115b682611577565b602082019050919050565b5f6020820190508181035f8301526115d88161159f565b9050919050565b7f52656361726761206e616f20696e6963696164612e00000000000000000000005f82015250565b5f6116136015836112d6565b915061161e826115df565b602082019050919050565b5f6020820190508181035f83015261164081611607565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61167e82611069565b915061168983611069565b925082820261169781611069565b915082820484148315176116ae576116ad611647565b5b5092915050565b5f6040820190506116c85f830185611295565b6116d56020830184611295565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061172057607f821691505b602082108103611733576117326116dc565b5b50919050565b7f53657373616f206e616f2065737461207265736572766164612e0000000000005f82015250565b5f61176d601a836112d6565b915061177882611739565b602082019050919050565b5f6020820190508181035f83015261179a81611761565b9050919050565b7f466f726120646f20686f726172696f206167656e6461646f2e000000000000005f82015250565b5f6117d56019836112d6565b91506117e0826117a1565b602082019050919050565b5f6020820190508181035f830152611802816117c9565b9050919050565b7f536f20706f64652063616e63656c617220736573736f657320726573657276615f8201527f6461732e00000000000000000000000000000000000000000000000000000000602082015250565b5f6118636024836112d6565b915061186e82611809565b604082019050919050565b5f6020820190508181035f83015261189081611857565b9050919050565b7f506f73746f206e616f206578697374652e0000000000000000000000000000005f82015250565b5f6118cb6011836112d6565b91506118d682611897565b602082019050919050565b5f6020820190508181035f8301526118f8816118bf565b9050919050565b5f61190982611069565b915061191483611069565b925082820190508082111561192c5761192b611647565b5b92915050565b7f4475726163616f20696e76616c6964612e0000000000000000000000000000005f82015250565b5f6119666011836112d6565b915061197182611932565b602082019050919050565b5f6020820190508181035f8301526119938161195a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f486f726172696f20656d20636f6e666c69746f2e0000000000000000000000005f82015250565b5f6119fb6014836112d6565b9150611a06826119c7565b602082019050919050565b5f6020820190508181035f830152611a28816119ef565b9050919050565b5f611a3982611069565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a6b57611a6a611647565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611aff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611ac4565b611b098683611ac4565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611b44611b3f611b3a84611069565b611b21565b611069565b9050919050565b5f819050919050565b611b5d83611b2a565b611b71611b6982611b4b565b848454611ad0565b825550505050565b5f5f905090565b611b88611b79565b611b93818484611b54565b505050565b5b81811015611bb657611bab5f82611b80565b600181019050611b99565b5050565b601f821115611bfb57611bcc81611aa3565b611bd584611ab5565b81016020851015611be4578190505b611bf8611bf085611ab5565b830182611b98565b50505b505050565b5f82821c905092915050565b5f611c1b5f1984600802611c00565b1980831691505092915050565b5f611c338383611c0c565b9150826002028217905092915050565b611c4c826112cc565b67ffffffffffffffff811115611c6557611c64611a76565b5b611c6f8254611709565b611c7a828285611bba565b5f60209050601f831160018114611cab575f8415611c99578287015190505b611ca38582611c28565b865550611d0a565b601f198416611cb986611aa3565b5f5b82811015611ce057848901518255600182019150602085019450602081019050611cbb565b86831015611cfd5784890151611cf9601f891682611c0c565b8355505b6001600288020188555050505b505050505050565b828183375f83830152505050565b5f611d2b83856112d6565b9350611d38838584611d12565b611d41836112f4565b840190509392505050565b5f6020820190508181035f830152611d65818486611d20565b9050939250505056fea26469706673582212207c43ac13d28d8045c498868d38338464e88d65dd19d2ae7775dcba2f41ddc2fe64736f6c634300081e0033",
}

// RecargaveABI is the input ABI used to generate the binding from.
// Deprecated: Use RecargaveMetaData.ABI instead.
var RecargaveABI = RecargaveMetaData.ABI

// RecargaveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecargaveMetaData.Bin instead.
var RecargaveBin = RecargaveMetaData.Bin

// DeployRecargave deploys a new Ethereum contract, binding an instance of Recargave to it.
func DeployRecargave(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recargave, error) {
	parsed, err := RecargaveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecargaveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recargave{RecargaveCaller: RecargaveCaller{contract: contract}, RecargaveTransactor: RecargaveTransactor{contract: contract}, RecargaveFilterer: RecargaveFilterer{contract: contract}}, nil
}

// Recargave is an auto generated Go binding around an Ethereum contract.
type Recargave struct {
	RecargaveCaller     // Read-only binding to the contract
	RecargaveTransactor // Write-only binding to the contract
	RecargaveFilterer   // Log filterer for contract events
}

// RecargaveCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecargaveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecargaveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecargaveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecargaveSession struct {
	Contract     *Recargave        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecargaveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecargaveCallerSession struct {
	Contract *RecargaveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RecargaveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecargaveTransactorSession struct {
	Contract     *RecargaveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RecargaveRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecargaveRaw struct {
	Contract *Recargave // Generic contract binding to access the raw methods on
}

// RecargaveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecargaveCallerRaw struct {
	Contract *RecargaveCaller // Generic read-only contract binding to access the raw methods on
}

// RecargaveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecargaveTransactorRaw struct {
	Contract *RecargaveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecargave creates a new instance of Recargave, bound to a specific deployed contract.
func NewRecargave(address common.Address, backend bind.ContractBackend) (*Recargave, error) {
	contract, err := bindRecargave(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recargave{RecargaveCaller: RecargaveCaller{contract: contract}, RecargaveTransactor: RecargaveTransactor{contract: contract}, RecargaveFilterer: RecargaveFilterer{contract: contract}}, nil
}

// NewRecargaveCaller creates a new read-only instance of Recargave, bound to a specific deployed contract.
func NewRecargaveCaller(address common.Address, caller bind.ContractCaller) (*RecargaveCaller, error) {
	contract, err := bindRecargave(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecargaveCaller{contract: contract}, nil
}

// NewRecargaveTransactor creates a new write-only instance of Recargave, bound to a specific deployed contract.
func NewRecargaveTransactor(address common.Address, transactor bind.ContractTransactor) (*RecargaveTransactor, error) {
	contract, err := bindRecargave(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecargaveTransactor{contract: contract}, nil
}

// NewRecargaveFilterer creates a new log filterer instance of Recargave, bound to a specific deployed contract.
func NewRecargaveFilterer(address common.Address, filterer bind.ContractFilterer) (*RecargaveFilterer, error) {
	contract, err := bindRecargave(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecargaveFilterer{contract: contract}, nil
}

// bindRecargave binds a generic wrapper to an already deployed contract.
func bindRecargave(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecargaveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recargave *RecargaveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recargave.Contract.RecargaveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recargave *RecargaveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recargave.Contract.RecargaveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recargave *RecargaveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recargave.Contract.RecargaveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recargave *RecargaveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recargave.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recargave *RecargaveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recargave.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recargave *RecargaveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recargave.Contract.contract.Transact(opts, method, params...)
}

// GetSessao is a free data retrieval call binding the contract method 0x1735301e.
//
// Solidity: function getSessao(uint256 _idSessao) view returns((uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8))
func (_Recargave *RecargaveCaller) GetSessao(opts *bind.CallOpts, _idSessao *big.Int) (RecargaVESimplificadoSessao, error) {
	var out []interface{}
	err := _Recargave.contract.Call(opts, &out, "getSessao", _idSessao)

	if err != nil {
		return *new(RecargaVESimplificadoSessao), err
	}

	out0 := *abi.ConvertType(out[0], new(RecargaVESimplificadoSessao)).(*RecargaVESimplificadoSessao)

	return out0, err

}

// GetSessao is a free data retrieval call binding the contract method 0x1735301e.
//
// Solidity: function getSessao(uint256 _idSessao) view returns((uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8))
func (_Recargave *RecargaveSession) GetSessao(_idSessao *big.Int) (RecargaVESimplificadoSessao, error) {
	return _Recargave.Contract.GetSessao(&_Recargave.CallOpts, _idSessao)
}

// GetSessao is a free data retrieval call binding the contract method 0x1735301e.
//
// Solidity: function getSessao(uint256 _idSessao) view returns((uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8))
func (_Recargave *RecargaveCallerSession) GetSessao(_idSessao *big.Int) (RecargaVESimplificadoSessao, error) {
	return _Recargave.Contract.GetSessao(&_Recargave.CallOpts, _idSessao)
}

// Postos is a free data retrieval call binding the contract method 0x2ab50969.
//
// Solidity: function postos(uint256 ) view returns(uint256 idPosto, address proprietario, string nome, string localizacao, uint256 precoPorUnidadeEnergia)
func (_Recargave *RecargaveCaller) Postos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IdPosto                *big.Int
	Proprietario           common.Address
	Nome                   string
	Localizacao            string
	PrecoPorUnidadeEnergia *big.Int
}, error) {
	var out []interface{}
	err := _Recargave.contract.Call(opts, &out, "postos", arg0)

	outstruct := new(struct {
		IdPosto                *big.Int
		Proprietario           common.Address
		Nome                   string
		Localizacao            string
		PrecoPorUnidadeEnergia *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IdPosto = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Proprietario = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Nome = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Localizacao = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.PrecoPorUnidadeEnergia = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Postos is a free data retrieval call binding the contract method 0x2ab50969.
//
// Solidity: function postos(uint256 ) view returns(uint256 idPosto, address proprietario, string nome, string localizacao, uint256 precoPorUnidadeEnergia)
func (_Recargave *RecargaveSession) Postos(arg0 *big.Int) (struct {
	IdPosto                *big.Int
	Proprietario           common.Address
	Nome                   string
	Localizacao            string
	PrecoPorUnidadeEnergia *big.Int
}, error) {
	return _Recargave.Contract.Postos(&_Recargave.CallOpts, arg0)
}

// Postos is a free data retrieval call binding the contract method 0x2ab50969.
//
// Solidity: function postos(uint256 ) view returns(uint256 idPosto, address proprietario, string nome, string localizacao, uint256 precoPorUnidadeEnergia)
func (_Recargave *RecargaveCallerSession) Postos(arg0 *big.Int) (struct {
	IdPosto                *big.Int
	Proprietario           common.Address
	Nome                   string
	Localizacao            string
	PrecoPorUnidadeEnergia *big.Int
}, error) {
	return _Recargave.Contract.Postos(&_Recargave.CallOpts, arg0)
}

// ProximoIdPosto is a free data retrieval call binding the contract method 0x1da099e6.
//
// Solidity: function proximoIdPosto() view returns(uint256)
func (_Recargave *RecargaveCaller) ProximoIdPosto(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Recargave.contract.Call(opts, &out, "proximoIdPosto")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProximoIdPosto is a free data retrieval call binding the contract method 0x1da099e6.
//
// Solidity: function proximoIdPosto() view returns(uint256)
func (_Recargave *RecargaveSession) ProximoIdPosto() (*big.Int, error) {
	return _Recargave.Contract.ProximoIdPosto(&_Recargave.CallOpts)
}

// ProximoIdPosto is a free data retrieval call binding the contract method 0x1da099e6.
//
// Solidity: function proximoIdPosto() view returns(uint256)
func (_Recargave *RecargaveCallerSession) ProximoIdPosto() (*big.Int, error) {
	return _Recargave.Contract.ProximoIdPosto(&_Recargave.CallOpts)
}

// ProximoIdSessao is a free data retrieval call binding the contract method 0xc8f947c1.
//
// Solidity: function proximoIdSessao() view returns(uint256)
func (_Recargave *RecargaveCaller) ProximoIdSessao(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Recargave.contract.Call(opts, &out, "proximoIdSessao")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProximoIdSessao is a free data retrieval call binding the contract method 0xc8f947c1.
//
// Solidity: function proximoIdSessao() view returns(uint256)
func (_Recargave *RecargaveSession) ProximoIdSessao() (*big.Int, error) {
	return _Recargave.Contract.ProximoIdSessao(&_Recargave.CallOpts)
}

// ProximoIdSessao is a free data retrieval call binding the contract method 0xc8f947c1.
//
// Solidity: function proximoIdSessao() view returns(uint256)
func (_Recargave *RecargaveCallerSession) ProximoIdSessao() (*big.Int, error) {
	return _Recargave.Contract.ProximoIdSessao(&_Recargave.CallOpts)
}

// Sessoes is a free data retrieval call binding the contract method 0xa718ae96.
//
// Solidity: function sessoes(uint256 ) view returns(uint256 idSessao, uint256 idPosto, address enderecoVE, uint256 horarioAgendadoInicio, uint256 horarioAgendadoFim, uint256 horarioRealInicio, uint256 energiaConsumida, uint256 custoTotal, uint8 status)
func (_Recargave *RecargaveCaller) Sessoes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IdSessao              *big.Int
	IdPosto               *big.Int
	EnderecoVE            common.Address
	HorarioAgendadoInicio *big.Int
	HorarioAgendadoFim    *big.Int
	HorarioRealInicio     *big.Int
	EnergiaConsumida      *big.Int
	CustoTotal            *big.Int
	Status                uint8
}, error) {
	var out []interface{}
	err := _Recargave.contract.Call(opts, &out, "sessoes", arg0)

	outstruct := new(struct {
		IdSessao              *big.Int
		IdPosto               *big.Int
		EnderecoVE            common.Address
		HorarioAgendadoInicio *big.Int
		HorarioAgendadoFim    *big.Int
		HorarioRealInicio     *big.Int
		EnergiaConsumida      *big.Int
		CustoTotal            *big.Int
		Status                uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IdSessao = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IdPosto = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EnderecoVE = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.HorarioAgendadoInicio = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HorarioAgendadoFim = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HorarioRealInicio = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EnergiaConsumida = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CustoTotal = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[8], new(uint8)).(*uint8)

	return *outstruct, err

}

// Sessoes is a free data retrieval call binding the contract method 0xa718ae96.
//
// Solidity: function sessoes(uint256 ) view returns(uint256 idSessao, uint256 idPosto, address enderecoVE, uint256 horarioAgendadoInicio, uint256 horarioAgendadoFim, uint256 horarioRealInicio, uint256 energiaConsumida, uint256 custoTotal, uint8 status)
func (_Recargave *RecargaveSession) Sessoes(arg0 *big.Int) (struct {
	IdSessao              *big.Int
	IdPosto               *big.Int
	EnderecoVE            common.Address
	HorarioAgendadoInicio *big.Int
	HorarioAgendadoFim    *big.Int
	HorarioRealInicio     *big.Int
	EnergiaConsumida      *big.Int
	CustoTotal            *big.Int
	Status                uint8
}, error) {
	return _Recargave.Contract.Sessoes(&_Recargave.CallOpts, arg0)
}

// Sessoes is a free data retrieval call binding the contract method 0xa718ae96.
//
// Solidity: function sessoes(uint256 ) view returns(uint256 idSessao, uint256 idPosto, address enderecoVE, uint256 horarioAgendadoInicio, uint256 horarioAgendadoFim, uint256 horarioRealInicio, uint256 energiaConsumida, uint256 custoTotal, uint8 status)
func (_Recargave *RecargaveCallerSession) Sessoes(arg0 *big.Int) (struct {
	IdSessao              *big.Int
	IdPosto               *big.Int
	EnderecoVE            common.Address
	HorarioAgendadoInicio *big.Int
	HorarioAgendadoFim    *big.Int
	HorarioRealInicio     *big.Int
	EnergiaConsumida      *big.Int
	CustoTotal            *big.Int
	Status                uint8
}, error) {
	return _Recargave.Contract.Sessoes(&_Recargave.CallOpts, arg0)
}

// SessoesPorPosto is a free data retrieval call binding the contract method 0x241cb514.
//
// Solidity: function sessoesPorPosto(uint256 , uint256 ) view returns(uint256)
func (_Recargave *RecargaveCaller) SessoesPorPosto(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Recargave.contract.Call(opts, &out, "sessoesPorPosto", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SessoesPorPosto is a free data retrieval call binding the contract method 0x241cb514.
//
// Solidity: function sessoesPorPosto(uint256 , uint256 ) view returns(uint256)
func (_Recargave *RecargaveSession) SessoesPorPosto(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Recargave.Contract.SessoesPorPosto(&_Recargave.CallOpts, arg0, arg1)
}

// SessoesPorPosto is a free data retrieval call binding the contract method 0x241cb514.
//
// Solidity: function sessoesPorPosto(uint256 , uint256 ) view returns(uint256)
func (_Recargave *RecargaveCallerSession) SessoesPorPosto(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Recargave.Contract.SessoesPorPosto(&_Recargave.CallOpts, arg0, arg1)
}

// AgendarSessao is a paid mutator transaction binding the contract method 0xb53d3e25.
//
// Solidity: function agendarSessao(uint256 _idPosto, uint256 _horarioInicio, uint256 _duracaoMinutos) returns()
func (_Recargave *RecargaveTransactor) AgendarSessao(opts *bind.TransactOpts, _idPosto *big.Int, _horarioInicio *big.Int, _duracaoMinutos *big.Int) (*types.Transaction, error) {
	return _Recargave.contract.Transact(opts, "agendarSessao", _idPosto, _horarioInicio, _duracaoMinutos)
}

// AgendarSessao is a paid mutator transaction binding the contract method 0xb53d3e25.
//
// Solidity: function agendarSessao(uint256 _idPosto, uint256 _horarioInicio, uint256 _duracaoMinutos) returns()
func (_Recargave *RecargaveSession) AgendarSessao(_idPosto *big.Int, _horarioInicio *big.Int, _duracaoMinutos *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.AgendarSessao(&_Recargave.TransactOpts, _idPosto, _horarioInicio, _duracaoMinutos)
}

// AgendarSessao is a paid mutator transaction binding the contract method 0xb53d3e25.
//
// Solidity: function agendarSessao(uint256 _idPosto, uint256 _horarioInicio, uint256 _duracaoMinutos) returns()
func (_Recargave *RecargaveTransactorSession) AgendarSessao(_idPosto *big.Int, _horarioInicio *big.Int, _duracaoMinutos *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.AgendarSessao(&_Recargave.TransactOpts, _idPosto, _horarioInicio, _duracaoMinutos)
}

// CancelarSessao is a paid mutator transaction binding the contract method 0xa97ab930.
//
// Solidity: function cancelarSessao(uint256 _idSessao) returns()
func (_Recargave *RecargaveTransactor) CancelarSessao(opts *bind.TransactOpts, _idSessao *big.Int) (*types.Transaction, error) {
	return _Recargave.contract.Transact(opts, "cancelarSessao", _idSessao)
}

// CancelarSessao is a paid mutator transaction binding the contract method 0xa97ab930.
//
// Solidity: function cancelarSessao(uint256 _idSessao) returns()
func (_Recargave *RecargaveSession) CancelarSessao(_idSessao *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.CancelarSessao(&_Recargave.TransactOpts, _idSessao)
}

// CancelarSessao is a paid mutator transaction binding the contract method 0xa97ab930.
//
// Solidity: function cancelarSessao(uint256 _idSessao) returns()
func (_Recargave *RecargaveTransactorSession) CancelarSessao(_idSessao *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.CancelarSessao(&_Recargave.TransactOpts, _idSessao)
}

// FinalizarRecarga is a paid mutator transaction binding the contract method 0x0a915a10.
//
// Solidity: function finalizarRecarga(uint256 _idSessao, uint256 _energiaConsumida) returns()
func (_Recargave *RecargaveTransactor) FinalizarRecarga(opts *bind.TransactOpts, _idSessao *big.Int, _energiaConsumida *big.Int) (*types.Transaction, error) {
	return _Recargave.contract.Transact(opts, "finalizarRecarga", _idSessao, _energiaConsumida)
}

// FinalizarRecarga is a paid mutator transaction binding the contract method 0x0a915a10.
//
// Solidity: function finalizarRecarga(uint256 _idSessao, uint256 _energiaConsumida) returns()
func (_Recargave *RecargaveSession) FinalizarRecarga(_idSessao *big.Int, _energiaConsumida *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.FinalizarRecarga(&_Recargave.TransactOpts, _idSessao, _energiaConsumida)
}

// FinalizarRecarga is a paid mutator transaction binding the contract method 0x0a915a10.
//
// Solidity: function finalizarRecarga(uint256 _idSessao, uint256 _energiaConsumida) returns()
func (_Recargave *RecargaveTransactorSession) FinalizarRecarga(_idSessao *big.Int, _energiaConsumida *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.FinalizarRecarga(&_Recargave.TransactOpts, _idSessao, _energiaConsumida)
}

// IniciarRecarga is a paid mutator transaction binding the contract method 0x8aeedc25.
//
// Solidity: function iniciarRecarga(uint256 _idSessao) returns()
func (_Recargave *RecargaveTransactor) IniciarRecarga(opts *bind.TransactOpts, _idSessao *big.Int) (*types.Transaction, error) {
	return _Recargave.contract.Transact(opts, "iniciarRecarga", _idSessao)
}

// IniciarRecarga is a paid mutator transaction binding the contract method 0x8aeedc25.
//
// Solidity: function iniciarRecarga(uint256 _idSessao) returns()
func (_Recargave *RecargaveSession) IniciarRecarga(_idSessao *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.IniciarRecarga(&_Recargave.TransactOpts, _idSessao)
}

// IniciarRecarga is a paid mutator transaction binding the contract method 0x8aeedc25.
//
// Solidity: function iniciarRecarga(uint256 _idSessao) returns()
func (_Recargave *RecargaveTransactorSession) IniciarRecarga(_idSessao *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.IniciarRecarga(&_Recargave.TransactOpts, _idSessao)
}

// RegistrarPosto is a paid mutator transaction binding the contract method 0xcd90c06f.
//
// Solidity: function registrarPosto(string _nome, string _localizacao, uint256 _precoPorUnidadeEnergia) returns()
func (_Recargave *RecargaveTransactor) RegistrarPosto(opts *bind.TransactOpts, _nome string, _localizacao string, _precoPorUnidadeEnergia *big.Int) (*types.Transaction, error) {
	return _Recargave.contract.Transact(opts, "registrarPosto", _nome, _localizacao, _precoPorUnidadeEnergia)
}

// RegistrarPosto is a paid mutator transaction binding the contract method 0xcd90c06f.
//
// Solidity: function registrarPosto(string _nome, string _localizacao, uint256 _precoPorUnidadeEnergia) returns()
func (_Recargave *RecargaveSession) RegistrarPosto(_nome string, _localizacao string, _precoPorUnidadeEnergia *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.RegistrarPosto(&_Recargave.TransactOpts, _nome, _localizacao, _precoPorUnidadeEnergia)
}

// RegistrarPosto is a paid mutator transaction binding the contract method 0xcd90c06f.
//
// Solidity: function registrarPosto(string _nome, string _localizacao, uint256 _precoPorUnidadeEnergia) returns()
func (_Recargave *RecargaveTransactorSession) RegistrarPosto(_nome string, _localizacao string, _precoPorUnidadeEnergia *big.Int) (*types.Transaction, error) {
	return _Recargave.Contract.RegistrarPosto(&_Recargave.TransactOpts, _nome, _localizacao, _precoPorUnidadeEnergia)
}

// RecargavePostoRegistradoIterator is returned from FilterPostoRegistrado and is used to iterate over the raw logs and unpacked data for PostoRegistrado events raised by the Recargave contract.
type RecargavePostoRegistradoIterator struct {
	Event *RecargavePostoRegistrado // Event containing the contract specifics and raw log

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
func (it *RecargavePostoRegistradoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecargavePostoRegistrado)
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
		it.Event = new(RecargavePostoRegistrado)
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
func (it *RecargavePostoRegistradoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecargavePostoRegistradoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecargavePostoRegistrado represents a PostoRegistrado event raised by the Recargave contract.
type RecargavePostoRegistrado struct {
	IdPosto      *big.Int
	Proprietario common.Address
	Nome         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostoRegistrado is a free log retrieval operation binding the contract event 0xc462caf9f0889b3d4dfe9d08c0f9cd1f1e64c5bb5f500d061e5d9475a13cdd90.
//
// Solidity: event PostoRegistrado(uint256 indexed idPosto, address indexed proprietario, string nome)
func (_Recargave *RecargaveFilterer) FilterPostoRegistrado(opts *bind.FilterOpts, idPosto []*big.Int, proprietario []common.Address) (*RecargavePostoRegistradoIterator, error) {

	var idPostoRule []interface{}
	for _, idPostoItem := range idPosto {
		idPostoRule = append(idPostoRule, idPostoItem)
	}
	var proprietarioRule []interface{}
	for _, proprietarioItem := range proprietario {
		proprietarioRule = append(proprietarioRule, proprietarioItem)
	}

	logs, sub, err := _Recargave.contract.FilterLogs(opts, "PostoRegistrado", idPostoRule, proprietarioRule)
	if err != nil {
		return nil, err
	}
	return &RecargavePostoRegistradoIterator{contract: _Recargave.contract, event: "PostoRegistrado", logs: logs, sub: sub}, nil
}

// WatchPostoRegistrado is a free log subscription operation binding the contract event 0xc462caf9f0889b3d4dfe9d08c0f9cd1f1e64c5bb5f500d061e5d9475a13cdd90.
//
// Solidity: event PostoRegistrado(uint256 indexed idPosto, address indexed proprietario, string nome)
func (_Recargave *RecargaveFilterer) WatchPostoRegistrado(opts *bind.WatchOpts, sink chan<- *RecargavePostoRegistrado, idPosto []*big.Int, proprietario []common.Address) (event.Subscription, error) {

	var idPostoRule []interface{}
	for _, idPostoItem := range idPosto {
		idPostoRule = append(idPostoRule, idPostoItem)
	}
	var proprietarioRule []interface{}
	for _, proprietarioItem := range proprietario {
		proprietarioRule = append(proprietarioRule, proprietarioItem)
	}

	logs, sub, err := _Recargave.contract.WatchLogs(opts, "PostoRegistrado", idPostoRule, proprietarioRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecargavePostoRegistrado)
				if err := _Recargave.contract.UnpackLog(event, "PostoRegistrado", log); err != nil {
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

// ParsePostoRegistrado is a log parse operation binding the contract event 0xc462caf9f0889b3d4dfe9d08c0f9cd1f1e64c5bb5f500d061e5d9475a13cdd90.
//
// Solidity: event PostoRegistrado(uint256 indexed idPosto, address indexed proprietario, string nome)
func (_Recargave *RecargaveFilterer) ParsePostoRegistrado(log types.Log) (*RecargavePostoRegistrado, error) {
	event := new(RecargavePostoRegistrado)
	if err := _Recargave.contract.UnpackLog(event, "PostoRegistrado", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecargaveRecargaFinalizadaIterator is returned from FilterRecargaFinalizada and is used to iterate over the raw logs and unpacked data for RecargaFinalizada events raised by the Recargave contract.
type RecargaveRecargaFinalizadaIterator struct {
	Event *RecargaveRecargaFinalizada // Event containing the contract specifics and raw log

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
func (it *RecargaveRecargaFinalizadaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecargaveRecargaFinalizada)
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
		it.Event = new(RecargaveRecargaFinalizada)
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
func (it *RecargaveRecargaFinalizadaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecargaveRecargaFinalizadaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecargaveRecargaFinalizada represents a RecargaFinalizada event raised by the Recargave contract.
type RecargaveRecargaFinalizada struct {
	IdSessao         *big.Int
	EnergiaConsumida *big.Int
	CustoTotal       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRecargaFinalizada is a free log retrieval operation binding the contract event 0xa7a1530868b4f4e1e8a987d23325f7bfc66dfaf15b746dd9bfc11c0b917c1865.
//
// Solidity: event RecargaFinalizada(uint256 indexed idSessao, uint256 energiaConsumida, uint256 custoTotal)
func (_Recargave *RecargaveFilterer) FilterRecargaFinalizada(opts *bind.FilterOpts, idSessao []*big.Int) (*RecargaveRecargaFinalizadaIterator, error) {

	var idSessaoRule []interface{}
	for _, idSessaoItem := range idSessao {
		idSessaoRule = append(idSessaoRule, idSessaoItem)
	}

	logs, sub, err := _Recargave.contract.FilterLogs(opts, "RecargaFinalizada", idSessaoRule)
	if err != nil {
		return nil, err
	}
	return &RecargaveRecargaFinalizadaIterator{contract: _Recargave.contract, event: "RecargaFinalizada", logs: logs, sub: sub}, nil
}

// WatchRecargaFinalizada is a free log subscription operation binding the contract event 0xa7a1530868b4f4e1e8a987d23325f7bfc66dfaf15b746dd9bfc11c0b917c1865.
//
// Solidity: event RecargaFinalizada(uint256 indexed idSessao, uint256 energiaConsumida, uint256 custoTotal)
func (_Recargave *RecargaveFilterer) WatchRecargaFinalizada(opts *bind.WatchOpts, sink chan<- *RecargaveRecargaFinalizada, idSessao []*big.Int) (event.Subscription, error) {

	var idSessaoRule []interface{}
	for _, idSessaoItem := range idSessao {
		idSessaoRule = append(idSessaoRule, idSessaoItem)
	}

	logs, sub, err := _Recargave.contract.WatchLogs(opts, "RecargaFinalizada", idSessaoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecargaveRecargaFinalizada)
				if err := _Recargave.contract.UnpackLog(event, "RecargaFinalizada", log); err != nil {
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

// ParseRecargaFinalizada is a log parse operation binding the contract event 0xa7a1530868b4f4e1e8a987d23325f7bfc66dfaf15b746dd9bfc11c0b917c1865.
//
// Solidity: event RecargaFinalizada(uint256 indexed idSessao, uint256 energiaConsumida, uint256 custoTotal)
func (_Recargave *RecargaveFilterer) ParseRecargaFinalizada(log types.Log) (*RecargaveRecargaFinalizada, error) {
	event := new(RecargaveRecargaFinalizada)
	if err := _Recargave.contract.UnpackLog(event, "RecargaFinalizada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecargaveRecargaIniciadaIterator is returned from FilterRecargaIniciada and is used to iterate over the raw logs and unpacked data for RecargaIniciada events raised by the Recargave contract.
type RecargaveRecargaIniciadaIterator struct {
	Event *RecargaveRecargaIniciada // Event containing the contract specifics and raw log

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
func (it *RecargaveRecargaIniciadaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecargaveRecargaIniciada)
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
		it.Event = new(RecargaveRecargaIniciada)
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
func (it *RecargaveRecargaIniciadaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecargaveRecargaIniciadaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecargaveRecargaIniciada represents a RecargaIniciada event raised by the Recargave contract.
type RecargaveRecargaIniciada struct {
	IdSessao *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRecargaIniciada is a free log retrieval operation binding the contract event 0x41a1f1d6bcb91629ef55145436e5e4b93c0b96ecb4d80c51f27b6b8ceeb78576.
//
// Solidity: event RecargaIniciada(uint256 indexed idSessao)
func (_Recargave *RecargaveFilterer) FilterRecargaIniciada(opts *bind.FilterOpts, idSessao []*big.Int) (*RecargaveRecargaIniciadaIterator, error) {

	var idSessaoRule []interface{}
	for _, idSessaoItem := range idSessao {
		idSessaoRule = append(idSessaoRule, idSessaoItem)
	}

	logs, sub, err := _Recargave.contract.FilterLogs(opts, "RecargaIniciada", idSessaoRule)
	if err != nil {
		return nil, err
	}
	return &RecargaveRecargaIniciadaIterator{contract: _Recargave.contract, event: "RecargaIniciada", logs: logs, sub: sub}, nil
}

// WatchRecargaIniciada is a free log subscription operation binding the contract event 0x41a1f1d6bcb91629ef55145436e5e4b93c0b96ecb4d80c51f27b6b8ceeb78576.
//
// Solidity: event RecargaIniciada(uint256 indexed idSessao)
func (_Recargave *RecargaveFilterer) WatchRecargaIniciada(opts *bind.WatchOpts, sink chan<- *RecargaveRecargaIniciada, idSessao []*big.Int) (event.Subscription, error) {

	var idSessaoRule []interface{}
	for _, idSessaoItem := range idSessao {
		idSessaoRule = append(idSessaoRule, idSessaoItem)
	}

	logs, sub, err := _Recargave.contract.WatchLogs(opts, "RecargaIniciada", idSessaoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecargaveRecargaIniciada)
				if err := _Recargave.contract.UnpackLog(event, "RecargaIniciada", log); err != nil {
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

// ParseRecargaIniciada is a log parse operation binding the contract event 0x41a1f1d6bcb91629ef55145436e5e4b93c0b96ecb4d80c51f27b6b8ceeb78576.
//
// Solidity: event RecargaIniciada(uint256 indexed idSessao)
func (_Recargave *RecargaveFilterer) ParseRecargaIniciada(log types.Log) (*RecargaveRecargaIniciada, error) {
	event := new(RecargaveRecargaIniciada)
	if err := _Recargave.contract.UnpackLog(event, "RecargaIniciada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecargaveSessaoAgendadaIterator is returned from FilterSessaoAgendada and is used to iterate over the raw logs and unpacked data for SessaoAgendada events raised by the Recargave contract.
type RecargaveSessaoAgendadaIterator struct {
	Event *RecargaveSessaoAgendada // Event containing the contract specifics and raw log

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
func (it *RecargaveSessaoAgendadaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecargaveSessaoAgendada)
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
		it.Event = new(RecargaveSessaoAgendada)
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
func (it *RecargaveSessaoAgendadaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecargaveSessaoAgendadaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecargaveSessaoAgendada represents a SessaoAgendada event raised by the Recargave contract.
type RecargaveSessaoAgendada struct {
	IdSessao      *big.Int
	IdPosto       *big.Int
	EnderecoVE    common.Address
	HorarioInicio *big.Int
	HorarioFim    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSessaoAgendada is a free log retrieval operation binding the contract event 0x6c55c2668ba44bf59ad5f2dab96ddc6d4537864bdfa5895056ecb11b55d3bb84.
//
// Solidity: event SessaoAgendada(uint256 indexed idSessao, uint256 indexed idPosto, address indexed enderecoVE, uint256 horarioInicio, uint256 horarioFim)
func (_Recargave *RecargaveFilterer) FilterSessaoAgendada(opts *bind.FilterOpts, idSessao []*big.Int, idPosto []*big.Int, enderecoVE []common.Address) (*RecargaveSessaoAgendadaIterator, error) {

	var idSessaoRule []interface{}
	for _, idSessaoItem := range idSessao {
		idSessaoRule = append(idSessaoRule, idSessaoItem)
	}
	var idPostoRule []interface{}
	for _, idPostoItem := range idPosto {
		idPostoRule = append(idPostoRule, idPostoItem)
	}
	var enderecoVERule []interface{}
	for _, enderecoVEItem := range enderecoVE {
		enderecoVERule = append(enderecoVERule, enderecoVEItem)
	}

	logs, sub, err := _Recargave.contract.FilterLogs(opts, "SessaoAgendada", idSessaoRule, idPostoRule, enderecoVERule)
	if err != nil {
		return nil, err
	}
	return &RecargaveSessaoAgendadaIterator{contract: _Recargave.contract, event: "SessaoAgendada", logs: logs, sub: sub}, nil
}

// WatchSessaoAgendada is a free log subscription operation binding the contract event 0x6c55c2668ba44bf59ad5f2dab96ddc6d4537864bdfa5895056ecb11b55d3bb84.
//
// Solidity: event SessaoAgendada(uint256 indexed idSessao, uint256 indexed idPosto, address indexed enderecoVE, uint256 horarioInicio, uint256 horarioFim)
func (_Recargave *RecargaveFilterer) WatchSessaoAgendada(opts *bind.WatchOpts, sink chan<- *RecargaveSessaoAgendada, idSessao []*big.Int, idPosto []*big.Int, enderecoVE []common.Address) (event.Subscription, error) {

	var idSessaoRule []interface{}
	for _, idSessaoItem := range idSessao {
		idSessaoRule = append(idSessaoRule, idSessaoItem)
	}
	var idPostoRule []interface{}
	for _, idPostoItem := range idPosto {
		idPostoRule = append(idPostoRule, idPostoItem)
	}
	var enderecoVERule []interface{}
	for _, enderecoVEItem := range enderecoVE {
		enderecoVERule = append(enderecoVERule, enderecoVEItem)
	}

	logs, sub, err := _Recargave.contract.WatchLogs(opts, "SessaoAgendada", idSessaoRule, idPostoRule, enderecoVERule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecargaveSessaoAgendada)
				if err := _Recargave.contract.UnpackLog(event, "SessaoAgendada", log); err != nil {
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

// ParseSessaoAgendada is a log parse operation binding the contract event 0x6c55c2668ba44bf59ad5f2dab96ddc6d4537864bdfa5895056ecb11b55d3bb84.
//
// Solidity: event SessaoAgendada(uint256 indexed idSessao, uint256 indexed idPosto, address indexed enderecoVE, uint256 horarioInicio, uint256 horarioFim)
func (_Recargave *RecargaveFilterer) ParseSessaoAgendada(log types.Log) (*RecargaveSessaoAgendada, error) {
	event := new(RecargaveSessaoAgendada)
	if err := _Recargave.contract.UnpackLog(event, "SessaoAgendada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecargaveSessaoCanceladaIterator is returned from FilterSessaoCancelada and is used to iterate over the raw logs and unpacked data for SessaoCancelada events raised by the Recargave contract.
type RecargaveSessaoCanceladaIterator struct {
	Event *RecargaveSessaoCancelada // Event containing the contract specifics and raw log

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
func (it *RecargaveSessaoCanceladaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecargaveSessaoCancelada)
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
		it.Event = new(RecargaveSessaoCancelada)
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
func (it *RecargaveSessaoCanceladaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecargaveSessaoCanceladaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecargaveSessaoCancelada represents a SessaoCancelada event raised by the Recargave contract.
type RecargaveSessaoCancelada struct {
	IdSessao *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSessaoCancelada is a free log retrieval operation binding the contract event 0xa26ee77645ce5364f41aa0394e99dab9dc93eb68e80e149e6ce3a3394c686b28.
//
// Solidity: event SessaoCancelada(uint256 indexed idSessao)
func (_Recargave *RecargaveFilterer) FilterSessaoCancelada(opts *bind.FilterOpts, idSessao []*big.Int) (*RecargaveSessaoCanceladaIterator, error) {

	var idSessaoRule []interface{}
	for _, idSessaoItem := range idSessao {
		idSessaoRule = append(idSessaoRule, idSessaoItem)
	}

	logs, sub, err := _Recargave.contract.FilterLogs(opts, "SessaoCancelada", idSessaoRule)
	if err != nil {
		return nil, err
	}
	return &RecargaveSessaoCanceladaIterator{contract: _Recargave.contract, event: "SessaoCancelada", logs: logs, sub: sub}, nil
}

// WatchSessaoCancelada is a free log subscription operation binding the contract event 0xa26ee77645ce5364f41aa0394e99dab9dc93eb68e80e149e6ce3a3394c686b28.
//
// Solidity: event SessaoCancelada(uint256 indexed idSessao)
func (_Recargave *RecargaveFilterer) WatchSessaoCancelada(opts *bind.WatchOpts, sink chan<- *RecargaveSessaoCancelada, idSessao []*big.Int) (event.Subscription, error) {

	var idSessaoRule []interface{}
	for _, idSessaoItem := range idSessao {
		idSessaoRule = append(idSessaoRule, idSessaoItem)
	}

	logs, sub, err := _Recargave.contract.WatchLogs(opts, "SessaoCancelada", idSessaoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecargaveSessaoCancelada)
				if err := _Recargave.contract.UnpackLog(event, "SessaoCancelada", log); err != nil {
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

// ParseSessaoCancelada is a log parse operation binding the contract event 0xa26ee77645ce5364f41aa0394e99dab9dc93eb68e80e149e6ce3a3394c686b28.
//
// Solidity: event SessaoCancelada(uint256 indexed idSessao)
func (_Recargave *RecargaveFilterer) ParseSessaoCancelada(log types.Log) (*RecargaveSessaoCancelada, error) {
	event := new(RecargaveSessaoCancelada)
	if err := _Recargave.contract.UnpackLog(event, "SessaoCancelada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
