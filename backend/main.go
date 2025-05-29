package main

import (
    "context"
    "fmt"
    "log"
	
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

    blockNumber, err := client.BlockNumber(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Bloco atual: %d\n", blockNumber)
}
