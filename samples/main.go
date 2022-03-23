package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/eosrpc.go"
)

func main() {
	eos := eosrpc.New("https://wax.greymass.com")
	chain := eos.NewChainAPI()

	ch, _ := chain.GetInfo()
	fmt.Println(ch)

	data, err := chain.GetCurrencyBalance(eosrpc.GetCurrencyBalanceProps{
		Code:    "eosio.token",
		Account: "5g2vm.wam",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
