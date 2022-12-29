package main

import (
	"context"
	"math/big"

	"github.com/adshao/go-binance/v2"
	"github.com/pkg/errors"
)

type BinanceClient struct {
	service *binance.ListPricesService
}

func NewBinanceClient() *BinanceClient {
	client := binance.NewClient("lol", "kek") // because we use public api
	service := client.NewListPricesService()
	return &BinanceClient{
		service: service,
	}
}

func (b *BinanceClient) GetPrice(symbol string) (*big.Float, error) {
	b.service.Symbol(symbol)

	result, err := b.service.
		Symbol(symbol).
		Do(context.Background())

	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("Can't get price, result array is empty")
	}
	price, ok := (&big.Float{}).SetString(result[0].Price)
	if !ok {
		return nil, errors.New("can't parse price")
	}
	return price, nil
}
