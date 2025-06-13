package main

import (
    "context"
    "fmt"
    "log"
    "math/big"

    contracts "charging-system/blockchain/contracts" // ajuste o import se necessário

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

const (
    gethRPCURL           = "http://localhost:8545"
    stationPrivateKeyHex = "8969d74267c878e17205a3acca2641c17093366ec3cdc68e0bc411733e429cde" // Troque para a chave do proprietário do posto
    contractAddressHex   = "charge-system/contracts/RecargaVE.sol" // Endereço do contrato implantado (substitua pelo endereço real após implantar o contrato
)

func main() {
    client, err := ethclient.Dial(gethRPCURL)
    if err != nil {
        log.Fatalf("Falha ao conectar ao nó Ethereum: %v", err)
    }

    privateKey, err := crypto.HexToECDSA(stationPrivateKeyHex)
    if err != nil {
        log.Fatalf("Erro ao carregar chave privada: %v", err)
    }

    chainID, err := client.ChainID(context.Background())
    if err != nil {
        log.Fatalf("Falha ao obter chainID: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        log.Fatalf("Falha ao criar autorização: %v", err)
    }

    contractAddress := common.HexToAddress(contractAddressHex)
    contract, err := contracts.NewRecargave(contractAddress, client)
    if err != nil {
        log.Fatalf("Erro ao instanciar contrato: %v", err)
    }

    postos := []struct {
        Nome     string
        Local    string
        PrecoWei int64
    }{
        {"Posto Centro", "Rua A, 100", 100},
        {"Posto Norte", "Av. B, 200", 120},
        {"Posto Sul", "Rua C, 300", 90},
        {"Posto Leste", "Av. D, 400", 110},
    }

    for i, p := range postos {
        tx, err := contract.RegistrarPosto(auth, p.Nome, p.Local, big.NewInt(p.PrecoWei))
        if err != nil {
            log.Fatalf("Erro ao registrar posto %s: %v", p.Nome, err)
        }
        fmt.Printf("Posto %s registrado! TX: %s\n", p.Nome, tx.Hash().Hex())

        // Consulta o posto recém-criado (IDs começam em 1)
        posto, err := contract.Postos(nil, big.NewInt(int64(i+1)))
        if err != nil {
            log.Fatalf("Erro ao consultar posto %d: %v", i+1, err)
        }
        fmt.Printf("ID: %d, Nome: %s, Localização: %s, Preço: %d\n", i+1, posto.Nome, posto.Localizacao, posto.PrecoPorUnidadeEnergia)
    }
}