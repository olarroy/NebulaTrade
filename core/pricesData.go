package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//BitstampCoin - latest info about the coin
type BitstampCoin struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Vwap      string `json:"vwap"`
	Volume    string `json:"volume"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Open      string `json:"open"`
}

//BitfinexCoin -
type BitfinexCoin struct {
	Symbol              string
	Bid                 string
	BidSize             string
	Asl                 string
	AskSize             string
	DailyChange         string
	DailyChangeRelative string
	LastPrice           string
	Volume              string
	High                string
	Low                 string
}

//BitfinexPrice -
func BitfinexPrice(exchange string) BitfinexCoin {

	//tBTCUSD
	req, err := http.Get("https://api-pub.bitfinex.com/v2/tickers?symbols=" + exchange)
	if err != nil {
		log.Panic(err)
	}
	defer req.Body.Close()
	resp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panic(err)
	}

	/*
		Once we get the data, split and get data
	*/

	formattedResponse := string(resp[2 : len(string(resp))-2])
	allValues := strings.Split(formattedResponse, ",")

	bfCoin := BitfinexCoin{
		Symbol:              allValues[0],
		Bid:                 allValues[1],
		BidSize:             allValues[2],
		Asl:                 allValues[3],
		AskSize:             allValues[4],
		DailyChange:         allValues[5],
		DailyChangeRelative: allValues[6],
		LastPrice:           allValues[7],
		Volume:              allValues[8],
		High:                allValues[9],
		Low:                 allValues[10],
	}

	log.Println(bfCoin.LastPrice)

	return bfCoin
}

//BitstampPrice - get latest data of X coin
func BitstampPrice(exchange string) BitstampCoin {

	var CoinData BitstampCoin
	resp, err := http.Get("https://www.bitstamp.net/api/v2/ticker/" + exchange)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}
	err = json.Unmarshal(data, &CoinData)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	return CoinData

}