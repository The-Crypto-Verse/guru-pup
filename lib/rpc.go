package lib

import "github.com/World-of-Cryptopups/eosrpc.go"

// creates a new chain api instance
func ChainAPI() *eosrpc.ChainAPI {
	rpc := eosrpc.New("https://wax.greymass.com")

	return rpc.NewChainAPI()
}

type BalanceProps struct {
	Kasu string
	Wax  string
}

// fetches the wax and kasu token wallet balances
func GetTokenBalance(wallet string) (*BalanceProps, error) {
	chain := ChainAPI()

	waxtoken, err := chain.GetCurrencyBalance(eosrpc.GetCurrencyBalanceProps{
		Code:    "eosio.token",
		Account: wallet,
	})
	if err != nil {
		return nil, err
	}

	kasutoken, err := chain.GetCurrencyBalance(eosrpc.GetCurrencyBalanceProps{
		Code:    "thekasutoken",
		Account: wallet,
	})
	if err != nil {
		return nil, err
	}

	wax := "0 WAX"
	if len(waxtoken) > 0 {
		wax = waxtoken[0]
	}

	kasu := "0 KASU"
	if len(kasutoken) > 0 {
		kasu = kasutoken[0]
	}

	return &BalanceProps{
		Kasu: kasu,
		Wax:  wax,
	}, nil

}
