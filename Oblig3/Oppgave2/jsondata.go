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

var catch Crypto

func main() {
	//http.HandleFunc("/",template1)
	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		Info1()
	})
	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		Info2()
	})
	http.HandleFunc("/3", func(w http.ResponseWriter, r *http.Request) {
		Info3()
	})
	http.HandleFunc("/4", func(w http.ResponseWriter, r *http.Request) {
		Info4()
	})
	http.HandleFunc("/5", func(w http.ResponseWriter, r *http.Request) {
		Info5()
	})
	http.ListenAndServe(":8080", nil)
	/*
	Info1()
	Info2()
	Info3()
	Info4()
	Info5()*/
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
	fmt.Printf("The Symbol of  %v is %v  \n", crypto[0].Name, crypto[0].Symbol)
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
	fmt.Printf("The Symbol of  %v is %v  \n", crypto[0].Name, crypto[0].Symbol)
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
	fmt.Printf("The Symbol of  %v is %v  \n", crypto[0].Name, crypto[0].Symbol)
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
	fmt.Printf("The Symbol of  %v is %v  \n", crypto[0].Name, crypto[0].Symbol)
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
	fmt.Printf("The Symbol of  %v is %v  \n", crypto[0].Name, crypto[0].Symbol)
}
