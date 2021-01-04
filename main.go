package main

import (
	"time"

	"github.com/NebulaTrade/config"
	"github.com/NebulaTrade/exchanges"
	"github.com/NebulaTrade/trading"
	"github.com/NebulaTrade/wallet"
)

func main() {

	for {

		w := wallet.ReadWallet()
		w.Balance = exchanges.GetBinanceWalletBNB()
		w.Symbol = config.CURRENCY
		if w.Status == "BUY" || w.Status == "BUY ORDER" {
			w.Timer++
		}
		w.WriteInWallet()
		trading.ExecuteMarket(&w)
		time.Sleep(1 * time.Second)
	}

}
