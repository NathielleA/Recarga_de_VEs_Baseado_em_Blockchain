package main

import (
	"fmt"
	"math/rand"
	"time"
)

var clientID int

func generatePosition() (int, int, int, int) {
	originX := rand.Intn(1000)
	originY := rand.Intn(1000)
	destinationX := rand.Intn(1000)
	destinationY := rand.Intn(1000)

	return originX, originY, destinationX, destinationY
}

func startCarLoop() {
	for {
		origX, origY, destX, destY := generatePosition()

		fmt.Print("%d, %d, %d, %d, %d", clientID, origX, origY, destX, destY)

		time.Sleep(2 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	clientID = rand.Intn(1000)

	fmt.Printf("Carro ID: %d\n", clientID)

	// Iniciando o loop do carro
	startCarLoop()

	//id := fmt.Sprintf("%d", clientID)
}
