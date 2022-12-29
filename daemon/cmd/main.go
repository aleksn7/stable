package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	TOKEN_PRECISION    = (&big.Int{}).Exp(big.NewInt(10), big.NewInt(18), nil)
	SIGNER_PRIVATE_KEY = os.Getenv("PRIVATE_KEY")
	RPC_HOST           = os.Getenv("RPC_HOST")
	CONTROLLER_ADDRESS = os.Getenv("CONTROLLER_ADDRESS")
)

const (
	SYMBOL = "ETHRUB"
)

func CreateEthClient() EthClient {
	client, err := ethclient.Dial(RPC_HOST)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(CONTROLLER_ADDRESS)
	instance, err := NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	pk, err := crypto.HexToECDSA(SIGNER_PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	return NewEthClient(pk, instance)
}

func main() {
	ethClient := CreateEthClient()
	binanceClient := NewBinanceClient()

	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		symbolPrice, err := binanceClient.GetPrice(SYMBOL)
		if err != nil {
			fmt.Println("Can't get price from binance", err)
			continue
		}

		price, _ := symbolPrice.Int(nil)
		price.Mul(price, TOKEN_PRECISION)

		_, err = ethClient.SetPrice(price)
		if err != nil {
			fmt.Println("Can't set price to contract", err)
			continue
		}

		fmt.Println("Set price", price)
	}
}
