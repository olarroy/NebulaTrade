package main

import (
	"github.com/NebulaTrade/config"
	"github.com/NebulaTrade/trading"
)

func main() {

	/*
		Geeting all the flags and setting them to the config file
	*/

	// BinanceWallet := exchanges.GetBinanceWalletBNB()

	// configCurrency := flag.String("currency", "1234", "Currency")
	// configLimit := flag.Float64("limit", 1, "Limit coins")
	// configProfit := flag.Float64("profit", 0.00002, "Profit to make")
	// configCounter := flag.Int("counter", 1000, "counter")

	// flag.Parse()

	// config.CURRENCY = *configCurrency
	// limit := *configLimit
	// config.PROFIT = *configProfit
	// config.COUNTER = *configCounter

	// /*
	// 	Set all the initial config data to start initial buy
	// */
	// w := wallet.ReadWallet()
	// w.Limit = limit
	// w.Symbol = config.CURRENCY
	// w.LastSell = 5000
	// w.Transactions = 0
	// w.Status = "BUY"
	// w.Available = BinanceWallet / 3
	// w.Balance = BinanceWallet
	// w.Profit = config.PROFIT
	// w.WriteInWallet()

	// for {

	// 	trading.SellingPositions()
	// 	w := wallet.ReadWallet()

	// 	if w.Status == "BUY" || w.Status == "BUY ORDER" {
	// 		w.Timer++
	// 	}
	// 	w.WriteInWallet()
	// 	trading.ExecuteMarket(&w)
	// 	time.Sleep(2 * time.Second)
	// }
	config.CURRENCY = "XRPBNB"
	trading.SellingPositions()
}
