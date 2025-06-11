package main

import (
	"context"
	"fmt"
	"log"

	contracts "charging-system/blockchain/contracts" // THIS IS THE CORRECT PATH // Importa o pacote do contrato

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// --- CONFIGURAÇÕES ---
const (
	// URL do seu nó Geth privado (do docker-compose)
	gethRPCURL = "http://localhost:8545"

	// Chave privada da conta que vai implantar o contrato.
	// Use uma das chaves das contas definidas no seu genesis.json.
	deployerPrivateKeyHex = "8969d74267c878e17205a3acca2641c17093366ec3cdc68e0bc411733e429cde"
)

func main() {
	// Conecta ao nó Ethereum
	client, err := ethclient.Dial(gethRPCURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao nó Ethereum: %v", err)
	}
	fmt.Println("Conectado ao nó Ethereum.")

	// Carrega a chave privada do implantador
	privateKey, err := crypto.HexToECDSA(deployerPrivateKeyHex)
	if err != nil {
		log.Fatalf("Erro ao carregar chave privada: %v", err)
	}

	// Configura a autorização para a transação (auth)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Falha ao obter chainID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Falha ao criar autorização: %v", err)
	}

	// Implanta o contrato na blockchain
	fmt.Println("Implantando contrato...")
	address, tx, _, err := contracts.DeployRecargave(auth, client)
	if err != nil {
		log.Fatalf("Falha ao implantar o contrato: %v", err)
	}

	// Imprime o endereço do contrato para ser usado no main.go
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("Contrato implantado com sucesso!")
	fmt.Println("Endereço do Contrato:", address.Hex())
	fmt.Println("Hash da Transação:   ", tx.Hash().Hex())
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("COPIE O ENDEREÇO DO CONTRATO E COLE NO SEU ARQUIVO main.go")
}
