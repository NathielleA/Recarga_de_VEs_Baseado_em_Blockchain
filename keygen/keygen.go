// keygen/keygen.go
package main

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet" // A BIBLIOTECA CORRETA
)

const mnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"

func main() {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatalf("Falha ao criar carteira: %v", err)
	}

	fmt.Println("Frase Mnemônica:", mnemonic)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("Use estes endereços no seu arquivo genesis.json")
	fmt.Println("-----------------------------------------------------------------")

	// Gera e imprime os 2 primeiros endereços
	for i := 0; i < 2; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
		account, err := wallet.Derive(path, true)
		if err != nil {
			log.Fatalf("Falha ao derivar conta %d: %v", i, err)
		}
		fmt.Printf("Conta de Índice %d: %s\n", i, account.Address.Hex())
	}
}
