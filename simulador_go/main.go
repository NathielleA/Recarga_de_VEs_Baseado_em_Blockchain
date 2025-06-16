package main

import (
	"log"
	"time"

	// Importe o pacote gerado pelo abigen.
	// Substitua "NOME_DO_SEU_MODULO_GO/simulador_go/recargave_contract" pelo caminho correto
	// onde NOME_DO_SEU_MODULO_GO é o nome do seu módulo Go (definido no go.mod).
	// Se RecargaVEAbigen.go estiver no mesmo diretório e com `package main`,
	// você não precisaria de um import separado para ele, mas usar um pacote distinto é mais limpo.
	// Para este exemplo, vamos assumir que você criou um pacote `recargave_contract`:
	"charging-system/blockchain/contracts" // <<< SUBSTITUA ESTE CAMINHO

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// Endereço do seu contrato que foi deployado com sucesso
	contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"

	// Chave privada da conta do CARRO (Account #1 do seu log)
	carPrivateKey = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

<<<<<<< HEAD
	// Chave privada da conta do DONO DO POSTO (Account #0 do seu log)
	stationOwnerPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
=======
	// Chave privada da conta que será o proprietário do posto (pré-financiada)
	// !!! NUNCA FAÇA HARDCODE DE CHAVES PRIVADAS EM PRODUÇÃO !!!
	postoOwnerPrivateKeyHex = "2307470b0ba60fd8410dcab163c45c513dc1bf44b7b1c1f5a0a94c2ed572c2ac" // <<< SUBSTITUA (ex: "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")

	// Chave privada da conta que será o VE (pré-financiada)
	// !!! NUNCA FAÇA HARDCODE DE CHAVES PRIVADAS EM PRODUÇÃO !!!
	veOwnerPrivateKeyHex = "3874744cc85007e018e4f22b0b66c310793196755fee92e87878aa3ea4c983e0" // <<< SUBSTITUA (ex: "df57089feb5f7fbf800908df002191758743875e5c39e8799d681825e68369f9")

	// Chain ID da sua rede Geth privada (o mesmo do genesis.json)
	chainIDValue = 13370 // <<< SUBSTITUA se o seu ChainID for diferente
>>>>>>> b22cec106ccd5462efbe31c439795a75e686cbb9
)

func main() {
	client, err := ethclient.Dial("ws://127.0.0.1:8545") //era http e trnasformei em websocket
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Criando a entidade Carro
	car, err := NewCar(client, carPrivateKey)
	if err != nil {
		log.Fatalf("Failed to create car: %v", err)
	}
	log.Printf("Carro criado com o endereco: %s", car.Address.Hex())

	// Criando a entidade Posto
	station, err := NewStation(client, stationOwnerPrivateKey, contractAddress)
	if err != nil {
		log.Fatalf("Failed to create station: %v", err)
	}
	log.Printf("Posto monitorando o contrato: %s", station.Address.Hex())

	// Iniciar o monitoramento de eventos do posto em uma go-routine
	go station.MonitorEvents()

	// --- Início da Simulação ---
	log.Println("--- SIMULACAO: INICIANDO ---")

	// 1. Carro reserva a estação
	if err := car.ReserveStation(contractAddress); err != nil {
		log.Fatalf("Erro na reserva: %v", err)
	}
	time.Sleep(5 * time.Second) // Espera para o evento ser processado

	// 2. Carro inicia a recarga
	if err := car.StartCharge(contractAddress); err != nil {
		log.Fatalf("Erro ao iniciar recarga: %v", err)
	}
	time.Sleep(5 * time.Second) // Espera

	// 3. Simula tempo de recarga
	log.Println("... Carro carregando por 10 segundos ...")
	time.Sleep(10 * time.Second)

	// 4. Carro finaliza a recarga
	if err := car.EndCharge(contractAddress); err != nil {
		log.Fatalf("Erro ao finalizar recarga: %v", err)
	}
	time.Sleep(5 * time.Second)

	// 5. Carro efetua o pagamento
	const chargeTimeForPayment = 10 // segundos
	if err := car.Pay(contractAddress, chargeTimeForPayment); err != nil {
		log.Fatalf("Erro no pagamento: %v", err)
	}

	log.Println("--- SIMULACAO: FINALIZADA COM SUCESSO ---")
}
