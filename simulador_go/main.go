package main

import (
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddress        = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
	stationOwnerPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

var carPrivateKeys = []string{
	"59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	"5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
	"7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
}

func main() {
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Erro ao conectar ao Ethereum: %v", err)
	}

	station, err := NewStation(client, stationOwnerPrivateKey, contractAddress)
	if err != nil {
		log.Fatalf("Erro ao criar o posto: %v", err)
	}
	log.Printf("Posto criado: %s", station.Address.Hex())
	go station.MonitorEvents()

	var cars []*Car
	for i, key := range carPrivateKeys {
		car, err := NewCar(client, key)
		if err != nil {
			log.Fatalf("Erro ao criar o carro %d: %v", i+1, err)
		}
		log.Printf("Carro %d: %s", i+1, car.Address.Hex())
		cars = append(cars, car)
	}

	log.Println("--- SIMULAÇÃO DE CONCORRÊNCIA INICIADA ---")

	var wg sync.WaitGroup

	for i, car := range cars {
		wg.Add(1)
		go func(i int, car *Car) {
			defer wg.Done()

			for {
				log.Printf("[Carro %d] Tentando reservar o posto...", i+1)
				err := car.ReserveStation(contractAddress)
				if err == nil {
					log.Printf("[Carro %d] ✅ Reserva realizada com sucesso!", i+1)
					break
				}
				log.Printf("[Carro %d] Posto ocupado. Tentando novamente em 2s...", i+1)
				time.Sleep(2 * time.Second)
			}

			time.Sleep(2 * time.Second)

			if err := car.StartCharge(contractAddress); err != nil {
				log.Printf("[Carro %d] Erro ao iniciar carga: %v", i+1, err)
				return
			}
			log.Printf("[Carro %d] 🔌 Carga iniciada", i+1)

			time.Sleep(5 * time.Second)

			if err := car.EndCharge(contractAddress); err != nil {
				log.Printf("[Carro %d] Erro ao finalizar carga: %v", i+1, err)
				return
			}
			log.Printf("[Carro %d] ⚡ Carga finalizada", i+1)

			if err := car.Pay(contractAddress, 5); err != nil {
				log.Printf("[Carro %d] Erro no pagamento: %v", i+1, err)
				return
			}
			log.Printf("[Carro %d] 💰 Pagamento realizado", i+1)

			log.Printf("[Carro %d] ✅ Processo concluído.", i+1)
		}(i, car)
	}

	wg.Wait()

	log.Println("--- SIMULAÇÃO DE CONCORRÊNCIA FINALIZADA ---")
}
