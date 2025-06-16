package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"ChargingStation/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Car struct {
	Client  *ethclient.Client
	Auth    *bind.TransactOpts
	Address common.Address
}

func NewCar(client *ethclient.Client, privateKeyHex string) (*Car, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	return &Car{Client: client, Auth: auth, Address: address}, nil
}

// CORRIGIDO: O tipo de retorno agora é *bindings.Bindings
func (c *Car) getInstance(contractAddressHex string) (*bindings.Bindings, error) {
	address := common.HexToAddress(contractAddressHex)
	// CORRIGIDO: A função chamada agora é NewBindings
	return bindings.NewBindings(address, c.Client)
}

func (c *Car) ReserveStation(contractAddress string) error {
	instance, err := c.getInstance(contractAddress)
	if err != nil {
		return err
	}

	tx, err := instance.Reserve(c.Auth)
	if err != nil {
		return err
	}
	log.Printf("CARRO: Funcao Reserve() chamada. Tx: %s", tx.Hash().Hex())
	return nil
}

func (c *Car) StartCharge(contractAddress string) error {
	instance, err := c.getInstance(contractAddress)
	if err != nil {
		return err
	}

	tx, err := instance.StartCharge(c.Auth)
	if err != nil {
		return err
	}
	log.Printf("CARRO: Funcao StartCharge() chamada. Tx: %s", tx.Hash().Hex())
	return nil
}

func (c *Car) EndCharge(contractAddress string) error {
	instance, err := c.getInstance(contractAddress)
	if err != nil {
		return err
	}

	tx, err := instance.EndCharge(c.Auth)
	if err != nil {
		return err
	}
	log.Printf("CARRO: Funcao EndCharge() chamada. Tx: %s", tx.Hash().Hex())
	return nil
}

func (c *Car) Pay(contractAddress string, chargeTime int64) error {
	instance, err := c.getInstance(contractAddress)
	if err != nil {
		return err
	}

	pricePerSecond, err := instance.PricePerSecond(nil)
	if err != nil {
		return err
	}

	totalCost := new(big.Int).Mul(big.NewInt(chargeTime), pricePerSecond)
	c.Auth.Value = totalCost

	tx, err := instance.PayForCharge(c.Auth, big.NewInt(chargeTime))
	if err != nil {
		return err
	}
	log.Printf("CARRO: Funcao Pay() chamada com %d wei. Tx: %s", totalCost, tx.Hash().Hex())
	return nil
}
