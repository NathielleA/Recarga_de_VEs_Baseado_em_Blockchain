package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"ve-backend/contract" // caminho do binding gerado (ajustado ao go.mod)
)

func main() {
	// Espera o nó Ethereum iniciar (tentativas com intervalo)
	var client *ethclient.Client
	var err error
	for i := 0; i < 10; i++ {
		client, err = ethclient.Dial("http://ethereum-node:8545")
		if err == nil {
			break
		}
		log.Println("Tentando conectar ao nó Ethereum...")
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("Erro ao conectar com Ethereum: %v", err)
	}

	// Carrega a chave privada via variável de ambiente
	privKeyHex := os.Getenv("PRIVATE_KEY")
	if privKeyHex == "" {
		log.Fatal("PRIVATE_KEY não definida no ambiente")
	}

	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		log.Fatalf("Erro ao carregar chave privada: %v", err)
	}

	// Cria transação autenticada com o ID da rede local (ex: 2025 ou 1337)
	chainID := big.NewInt(1337) // use 1337 se estiver no modo --dev
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Erro ao criar transactor: %v", err)
	}

	// Faz deploy do contrato
	address, tx, instance, err := contract.DeployEVCharging(auth, client)
	if err != nil {
		log.Fatalf("Erro ao fazer deploy do contrato: %v", err)
	}

	fmt.Println("✅ Contrato deployado em:", address.Hex())
	fmt.Println("🔗 Transação de deploy:", tx.Hash().Hex())

	_ = instance // você pode usar essa instância para chamar métodos depois
}
