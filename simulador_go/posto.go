package main

import (
	"context"
	"log"
	"strings"

	"ChargingStation/bindings" // O caminho de importação permanece o mesmo

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Station struct {
	Client  *ethclient.Client
	Address common.Address
}

func NewStation(client *ethclient.Client, ownerPrivateKey string, contractAddressHex string) (*Station, error) {
	address := common.HexToAddress(contractAddressHex)
	return &Station{Client: client, Address: address}, nil
}

func (s *Station) MonitorEvents() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{s.Address},
	}

	logs := make(chan types.Log)
	sub, err := s.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("POSTO: Erro ao subscrever eventos: %v", err)
	}

	log.Println("POSTO: Monitorando eventos do contrato...")

	// Corrigido: O ABI está dentro de BindingsMetaData.ABI
	contractAbi, err := abi.JSON(strings.NewReader(string(bindings.BindingsMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Printf("POSTO: Erro na subscricao: %v", err)
		case vLog := <-logs:
			event, err := contractAbi.EventByID(vLog.Topics[0])
			if err != nil {
				log.Printf("POSTO: Nao foi possivel encontrar evento com ID: %v", err)
				continue
			}

			switch event.Name {
			case "StationReserved":
				log.Printf("POSTO: Evento recebido! Estacao reservada. Fisicamente, a estacao seria bloqueada agora.")
			case "ChargingStarted":
				log.Printf("POSTO: Evento recebido! Recarga iniciada. Liberando energia...")
			case "ChargingCompleted":
				// Corrigido: Usa BindingsChargingCompleted
				var completedEvent bindings.BindingsChargingCompleted
				_ = contractAbi.UnpackIntoInterface(&completedEvent, event.Name, vLog.Data)

				// O acesso aos campos do evento (ChargeTime, Cost) pode precisar de ajuste
				// se a estrutura de unpacking mudar. Com base na sua lógica, vamos
				// manter como está por enquanto, mas fique atento a isso.
				log.Printf("POSTO: Evento recebido! Recarga finalizada. Tempo: %d s, Custo: %d wei. Preparando para pagamento.", completedEvent.ChargeTime, completedEvent.Cost)

			case "PaymentReceived":
				log.Printf("POSTO: Evento recebido! Pagamento confirmado. Obrigado!")
			}
		}
	}
}
