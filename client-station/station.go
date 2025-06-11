package main

import (
	"fmt"
)

var clientID int

func main() {

	fmt.Printf("Selecione o Servidor:\n")
	fmt.Printf("1 - Servidor 1\n")
	fmt.Printf("2 - Servidor 2\n")
	var serverChoice int
	fmt.Scanf("%d", &serverChoice)
	if serverChoice < 1 || serverChoice > 2 {
		fmt.Println("Escolha inválida. Por favor, selecione 1 ou 2.")
		return
	}

	fmt.Printf("Selecione o Posto:\n")
	fmt.Printf("1 - Posto 1\n")
	fmt.Printf("2 - Posto 2\n")
	var stationChoice int
	fmt.Scanf("%d", &stationChoice)
	if stationChoice < 1 || stationChoice > 2 {
		fmt.Println("Escolha inválida. Por favor, selecione 1 ou 2.")
		return
	}

	//fmt.Printf("Posto ID: %d\n", stationID)

	//id := fmt.Sprintf("%d", clientID)
}
