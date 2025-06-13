package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	recargave "recargave" // ajuste o caminho do pacote conforme necessário

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Conexão com o nó local da blockchain
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Erro ao conectar no nó: %v", err)
	}
	defer client.Close()

	// Chave privada do carro (exemplo, substitua pela sua)
	privateKeyHex := "SUA_CHAVE_PRIVADA"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Transactor
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	chainID := big.NewInt(1337) // exemplo, use o correto para sua rede
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// Endereço do contrato já implantado
	contractAddress := common.HexToAddress("ENDERECO_DO_CONTRATO")
	contract, err := recargave.NewRecargave(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Etapa 1: Agendar sessão de recarga
	idPosto := big.NewInt(0)
	horarioInicio := big.NewInt(time.Now().Add(time.Minute * 5).Unix()) // 5 min após o atual
	duracao := big.NewInt(30) // 30 minutos

	tx, err := contract.AgendarSessao(auth, idPosto, horarioInicio, duracao)
	if err != nil {
		log.Fatalf("Erro ao agendar sessão: %v", err)
	}
	fmt.Println("Sessão agendada! Hash da transação:", tx.Hash().Hex())

	// Esperar alguns segundos (simulação do tempo até a recarga)
	time.Sleep(10 * time.Second)

	// Obter próximo ID da sessão
	idSessao, err := contract.ProximoIdSessao(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Erro ao obter ID da sessão: %v", err)
	}
	idSessao = new(big.Int).Sub(idSessao, big.NewInt(1)) // última sessão criada

	// Etapa 2: Iniciar recarga
	tx, err = contract.IniciarRecarga(auth, idSessao)
	if err != nil {
		log.Fatalf("Erro ao iniciar recarga: %v", err)
	}
	fmt.Println("Recarga iniciada! Hash da transação:", tx.Hash().Hex())

	// Simular recarga
	time.Sleep(5 * time.Second)

	// Etapa 3: Finalizar recarga
	energiaConsumida := big.NewInt(15) // kWh
	tx, err = contract.FinalizarRecarga(auth, idSessao, energiaConsumida)
	if err != nil {
		log.Fatalf("Erro ao finalizar recarga: %v", err)
	}
	fmt.Println("Recarga finalizada! Hash da transação:", tx.Hash().Hex())

	// Etapa 4: Consultar dados da sessão
	sessao, err := contract.GetSessao(&bind.CallOpts{}, idSessao)
	if err != nil {
		log.Fatalf("Erro ao buscar sessão: %v", err)
	}

	fmt.Println("Resumo da Sessão:")
	fmt.Println("ID:", sessao.IdSessao)
	fmt.Println("Posto:", sessao.IdPosto)
	fmt.Println("Energia:", sessao.EnergiaConsumida)
	fmt.Println("Custo:", sessao.CustoTotal)
	fmt.Println("Status:", sessao.Status)
}
