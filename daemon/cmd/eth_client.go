package main

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthClient struct {
	pk       *ecdsa.PrivateKey
	contract *Contract
}

func NewEthClient(pk *ecdsa.PrivateKey, contract *Contract) EthClient {
	return EthClient{
		pk:       pk,
		contract: contract,
	}
}

func (c *EthClient) buildOpts() *bind.TransactOpts {
	opt := bind.NewKeyedTransactor(c.pk)
	opt.GasPrice = big.NewInt(469667619) // because ganache doesn't support new ethereum API
	return opt
}

func (c *EthClient) SetPrice(price *big.Int) (*types.Transaction, error) {
	return c.contract.SetCollateralPrice(c.buildOpts(), price)
}
