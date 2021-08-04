package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	token "./model"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func amain() {

	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s\n", name)
	fmt.Printf("symbol:%s\n", symbol)
	fmt.Printf("decimals:%s\n", decimals)

	fmt.Printf("wei:%s\n", bal)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance:%f", value)
}
