package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Crypto []struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

func main() {
	Info1()
	Info2()
	Info3()
	Info4()
	Info5()
}

func Info1() {
	url := "http://api.coinmarketcap.com/v1/ticker/bitcoin"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	var crypto Crypto
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}

	/*switch info2 {
	case "price":*/
	fmt.Printf("The price of %v today is %v usd \n", crypto[0].Name, crypto[0].PriceUsd)
}

func Info2() {
		url := "http://api.coinmarketcap.com/v1/ticker/dogecoin"
		resp, err := http.Get(url)
		if err != nil {
			panic(err.Error())
		}
		defer resp.Body.Close()

		var crypto Crypto
		if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
			return
		}
		fmt.Printf("The price of %v today is %v usd \n", crypto[0].Name, crypto[0].PriceUsd)
}

func Info3() {
	url := "http://api.coinmarketcap.com/v1/ticker/monero"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	var crypto Crypto
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}
	fmt.Printf("The price of %v today is %v usd \n", crypto[0].Name, crypto[0].PriceUsd)
}

func Info4() {
	url := "http://api.coinmarketcap.com/v1/ticker/ethereum"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	var crypto Crypto
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}
	fmt.Printf("The price of %v today is %v usd \n", crypto[0].Name, crypto[0].PriceUsd)
}

func Info5() {
	url := "http://api.coinmarketcap.com/v1/ticker/litecoin"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	var crypto Crypto
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}
	fmt.Printf("The price of %v today is %v usd \n", crypto[0].Name, crypto[0].PriceUsd)
}
