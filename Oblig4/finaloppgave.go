package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"html/template"
)

//type structure 1:  https://api.coinmarketcap.com/v1/ticker/bitcoin/?convert=NOK
type Crypto4 []struct {
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
	PriceNok         string `json:"price_nok"`
	Two4HVolumeNok   string `json:"24h_volume_nok"`
	MarketCapNok     string `json:"market_cap_nok"`
}

// type structure 2: https://api.cryptonator.com/api/full/btc-usd
type Crypto5 struct {
	Ticker struct {
		Base    string `json:"base"`
		Target  string `json:"target"`
		Price   string `json:"price"`
		Volume  string `json:"volume"`
		Change  string `json:"change"`
		Markets []struct {
			Market string  `json:"market"`
			Price  string  `json:"price"`
			Volume float64 `json:"volume"`
		} `json:"markets"`
	} `json:"ticker"`
	Timestamp int    `json:"timestamp"`
	Success   bool   `json:"success"`
	Error     string `json:"error"`
}

func main()  {
	//call those functions through handleFunc by putting numbers to differentiate.
	http.HandleFunc("/", Home) //homepage
	http.HandleFunc("/1", CollectData11)  //bitcoin
	http.HandleFunc("/2", CollectData22)  //bitcoin further detail
	http.HandleFunc("/3", CollectData33)  //five different crypto currencies
	http.HandleFunc("/4", CollectData44) //litecoin detail
	http.HandleFunc("/5", CollectData55) //ripple detail
	http.HandleFunc("/6", CollectData66) //ethereum detail
	http.HandleFunc("/7", CollectData77) // monero detail
	//calling templates here
	http.HandleFunc("/8", TemplateBicoin)
	http.HandleFunc("/9", TemplateLitecoin)
	http.HandleFunc("/10", TemplateRipple)
	http.HandleFunc("/11", TemplateEthereum)
	http.HandleFunc("/12", TemplateMonero)

	//checking error and establish server on 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//Function will setup the homepage using a template file.
func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	items := struct {
		Ole string
	}{
		Ole: "Ali er masterkoder",
	}
	t.Execute(w, items)
}

//function 1 collect api data and present live details of bitcoin crypto currency on the web.
func CollectData11 (w http.ResponseWriter, r *http.Request) {
	url := "https://api.coinmarketcap.com/v1/ticker/bitcoin/?convert=NOK"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() // close the body after loading data.

	var crypto Crypto4 //instance which will hold the data in the shape of structure.

	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil { //here decode the data and check error to return successfully
		return
	}
	// Printing out the information get lively through api after decoding.
	fmt.Fprintf(w,"The price of %v today is '%v' usd \n", crypto[0].Name, crypto[0].PriceUsd)
	fmt.Fprintf(w,"The price in norwegeian krones of %v is '%v' nok \n", crypto[0].Name, crypto[0].PriceNok)
	fmt.Fprintf(w,"The market capital of %v today is '%v' \n", crypto[0].Name, crypto[0].MarketCapNok)
	fmt.Fprintf(w,"The ID of %v is '%v'  \n", crypto[0].Name, crypto[0].ID)
	fmt.Fprintf(w,"The Symbol of  %v is '%v'  \n", crypto[0].Name, crypto[0].Symbol)
	fmt.Fprintf(w,"The avaialable supply of %v today is '%v' \n", crypto[0].Name, crypto[0].AvailableSupply)
	fmt.Fprintf(w,"The percentage changed in 24h of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange24H)
	fmt.Fprintf(w,"The percentage changed in last 7-days of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange7D)
	fmt.Fprintf(w,"The last updated result of %v is '%v' \n", crypto[0].Name, crypto[0].LastUpdated)
	fmt.Fprintf(w,"****************************************************************************\n")
}

//function 1 collect api data and present live details of litecoin crypto currency on the web.
func CollectData44 (w http.ResponseWriter, r *http.Request) {
	url := "https://api.coinmarketcap.com/v1/ticker/litecoin/?convert=NOK"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() // close the body after loading data.

	var crypto Crypto4 //instance which will hold the data in the shape of structure.

	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil { //here decode the data and check error to return successfully
		return
	}
	// Printing out the information get lively through api after decoding.
	fmt.Fprintf(w,"The price of %v today is '%v' usd \n", crypto[0].Name, crypto[0].PriceUsd)
	fmt.Fprintf(w,"The price in norwegeian krones of %v is '%v' nok \n", crypto[0].Name, crypto[0].PriceNok)
	fmt.Fprintf(w,"The market capital of %v today is '%v' \n", crypto[0].Name, crypto[0].MarketCapNok)
	fmt.Fprintf(w,"The ID of %v is '%v'  \n", crypto[0].Name, crypto[0].ID)
	fmt.Fprintf(w,"The Symbol of  %v is '%v'  \n", crypto[0].Name, crypto[0].Symbol)
	fmt.Fprintf(w,"The avaialable supply of %v today is '%v' \n", crypto[0].Name, crypto[0].AvailableSupply)
	fmt.Fprintf(w,"The percentage changed in 24h of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange24H)
	fmt.Fprintf(w,"The percentage changed in last 7-days of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange7D)
	fmt.Fprintf(w,"The last updated result of %v is '%v' \n", crypto[0].Name, crypto[0].LastUpdated)
	fmt.Fprintf(w,"****************************************************************************")
}

//function 1 collect api data and present live details of ripple crypto currency on the web.
func CollectData55 (w http.ResponseWriter, r *http.Request) {
	url := "https://api.coinmarketcap.com/v1/ticker/ripple/?convert=NOK"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() // close the body after loading data.

	var crypto Crypto4 //instance which will hold the data in the shape of structure.

	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil { //here decode the data and check error to return successfully
		return
	}
	// Printing out the information get lively through api after decoding.
	fmt.Fprintf(w,"The price of %v today is '%v' usd \n", crypto[0].Name, crypto[0].PriceUsd)
	fmt.Fprintf(w,"The price in norwegeian krones of %v is '%v' nok \n", crypto[0].Name, crypto[0].PriceNok)
	fmt.Fprintf(w,"The market capital of %v today is '%v' \n", crypto[0].Name, crypto[0].MarketCapNok)
	fmt.Fprintf(w,"The ID of %v is '%v'  \n", crypto[0].Name, crypto[0].ID)
	fmt.Fprintf(w,"The Symbol of  %v is '%v'  \n", crypto[0].Name, crypto[0].Symbol)
	fmt.Fprintf(w,"The avaialable supply of %v today is '%v' \n", crypto[0].Name, crypto[0].AvailableSupply)
	fmt.Fprintf(w,"The percentage changed in 24h of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange24H)
	fmt.Fprintf(w,"The percentage changed in last 7-days of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange7D)
	fmt.Fprintf(w,"The last updated result of %v is '%v' \n", crypto[0].Name, crypto[0].LastUpdated)
	fmt.Fprintf(w,"****************************************************************************")
}

//function 1 collect api data and present live details of ethereum crypto currency on the web.
func CollectData66 (w http.ResponseWriter, r *http.Request) {
	url := "https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=NOK"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() // close the body after loading data.

	var crypto Crypto4 //instance which will hold the data in the shape of structure.

	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil { //here decode the data and check error to return successfully
		return
	}
	// Printing out the information get lively through api after decoding.
	fmt.Fprintf(w,"The price of %v today is '%v' usd \n", crypto[0].Name, crypto[0].PriceUsd)
	fmt.Fprintf(w,"The price in norwegeian krones of %v is '%v' nok \n", crypto[0].Name, crypto[0].PriceNok)
	fmt.Fprintf(w,"The market capital of %v today is '%v' \n", crypto[0].Name, crypto[0].MarketCapNok)
	fmt.Fprintf(w,"The ID of %v is '%v'  \n", crypto[0].Name, crypto[0].ID)
	fmt.Fprintf(w,"The Symbol of  %v is '%v'  \n", crypto[0].Name, crypto[0].Symbol)
	fmt.Fprintf(w,"The avaialable supply of %v today is '%v' \n", crypto[0].Name, crypto[0].AvailableSupply)
	fmt.Fprintf(w,"The percentage changed in 24h of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange24H)
	fmt.Fprintf(w,"The percentage changed in last 7-days of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange7D)
	fmt.Fprintf(w,"The last updated result of %v is '%v' \n", crypto[0].Name, crypto[0].LastUpdated)
	fmt.Fprintf(w,"****************************************************************************")
}

//function 1 collect api data and present live details of monero crypto currency on the web.
func CollectData77 (w http.ResponseWriter, r *http.Request) {
	url := "https://api.coinmarketcap.com/v1/ticker/monero/?convert=NOK"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() // close the body after loading data.

	var crypto Crypto4 //instance which will hold the data in the shape of structure.

	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil { //here decode the data and check error to return successfully
		return
	}
	// Printing out the information get lively through api after decoding.
	fmt.Fprintf(w,"The price of %v today is '%v' usd \n", crypto[0].Name, crypto[0].PriceUsd)
	fmt.Fprintf(w,"The price in norwegeian krones of %v is '%v' nok \n", crypto[0].Name, crypto[0].PriceNok)
	fmt.Fprintf(w,"The market capital of %v today is '%v' \n", crypto[0].Name, crypto[0].MarketCapNok)
	fmt.Fprintf(w,"The ID of %v is '%v'  \n", crypto[0].Name, crypto[0].ID)
	fmt.Fprintf(w,"The Symbol of  %v is '%v'  \n", crypto[0].Name, crypto[0].Symbol)
	fmt.Fprintf(w,"The avaialable supply of %v today is '%v' \n", crypto[0].Name, crypto[0].AvailableSupply)
	fmt.Fprintf(w,"The percentage changed in 24h of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange24H)
	fmt.Fprintf(w,"The percentage changed in last 7-days of %v is '%v' \n", crypto[0].Name, crypto[0].PercentChange7D)
	fmt.Fprintf(w,"The last updated result of %v is '%v' \n", crypto[0].Name, crypto[0].LastUpdated)
	fmt.Fprintf(w,"****************************************************************************")
}

//function 2 collect another api data and presents live details of btc on the web.
func CollectData22 (w http.ResponseWriter, r *http.Request) {
	url := "https://api.cryptonator.com/api/full/btc-usd"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close() //close the body after loading data.

	var crypto2 Crypto5 //store the data into crypto instance
	if err := json.NewDecoder(resp.Body).Decode(&crypto2); err != nil { //here decode the data and check error to return successfully
		return
	}
	// Output BTC detail.
	fmt.Fprintf(w,"The base is '%v' and target we have, is '%v' \n", crypto2.Ticker.Base, crypto2.Ticker.Target)
	fmt.Fprintf(w,"The price of base %v is '%v' \n", crypto2.Ticker.Base, crypto2.Ticker.Price)
	fmt.Fprintf(w,"The volume of its '%v' and we have also a change '%v' in market \n", crypto2.Ticker.Volume, crypto2.Ticker.Change)
}

//funtion 3 collect another api data and present live details of five different crypto-currencies on the web.
func CollectData33(w http.ResponseWriter, r *http.Request)  {
	url := "https://api.cryptonator.com/api/full/btc-usd"
	resp, err:= http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close() //close the body after loading data.

	var data Crypto5 //storing data into the instance
	if err := json.NewDecoder(resp.Body).Decode(&data); err !=nil{ //here decode the data and check error to return successfully
		return
	}
	//Here are the output of five different crypto currencies with details.

	fmt.Fprintf(w,"Below are other five crypto currencies with prices and their volume numbers:" + "\n")
	fmt.Fprintf(w, "1:\tThe market of '%v' with the price of '%v' and the volume is '%v' \n", data.Ticker.Markets[1].Market, data.Ticker.Markets[1].Price, data.Ticker.Markets[1].Volume)
	fmt.Fprintf(w, "\tThe timestamp of '%v' is '%v' and success result is '%v' \n", data.Ticker.Markets[1].Market, data.Timestamp, data.Success)
	fmt.Fprintf(w, "2:\tThe market has also '%v' with the price of '%v' and the volume is '%v' \n", data.Ticker.Markets[2].Market, data.Ticker.Markets[2].Price, data.Ticker.Markets[2].Volume)
	fmt.Fprintf(w, "\tThe timestamp of '%v' is '%v' and success result is '%v' \n", data.Ticker.Markets[2].Market, data.Timestamp, data.Success)
	fmt.Fprintf(w, "3:\tThis is '%v' with the price of '%v' and the volume is '%v' \n", data.Ticker.Markets[3].Market, data.Ticker.Markets[3].Price, data.Ticker.Markets[3].Volume)
	fmt.Fprintf(w, "\tThe timestamp of '%v' is '%v' and success result is '%v' \n", data.Ticker.Markets[3].Market, data.Timestamp, data.Success)
	fmt.Fprintf(w, "4:\tThe market has '%v' with the price of '%v' and the volume is '%v' \n", data.Ticker.Markets[4].Market, data.Ticker.Markets[4].Price, data.Ticker.Markets[4].Volume)
	fmt.Fprintf(w, "\tThe timestamp of '%v' is '%v' and success result '%v' \n", data.Ticker.Markets[4].Market, data.Timestamp, data.Success)
	fmt.Fprintf(w, "5:\tOn this position is '%v' and price is '%v' with the volume of '%v' \n", data.Ticker.Markets[5].Market, data.Ticker.Markets[5].Price, data.Ticker.Markets[5].Volume)
	fmt.Fprintf(w, "\tThe timestamp of '%v' is '%v' and success result is '%v' \n", data.Ticker.Markets[5].Market, data.Timestamp, data.Success)
	fmt.Fprintf(w,"************************************************************************************************ \n")
	}
	
	
	

// Template for bitcoin
func TemplateBicoin(w http.ResponseWriter, r *http.Request) {
	url := "http://api.coinmarketcap.com/v1/ticker/bitcoin"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() //close the body after loading data.

	var crypto Crypto4 	//storing data into the instance
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}

	// Selects the html template file for print data
	t, err := template.ParseFiles("test.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Prints the API data on a temporary website, which are stored in memory
	err = t.Execute(w, crypto[0])
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

// Template for litecoing
func TemplateLitecoin(w http.ResponseWriter, r *http.Request) {
	url := "http://api.coinmarketcap.com/v1/ticker/litecoin"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() //close the body after loading data.

	var crypto Crypto4 //storing data into the instance
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}
	// Selects the html template file for print data
	t, err := template.ParseFiles("test.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Prints the API data on a temporary website, which are stored in memory
	err = t.Execute(w, crypto[0])
	if err != nil {
		log.Print("template executing error: ", err)
	}

}

// Template for Ripple
func TemplateRipple(w http.ResponseWriter, r *http.Request) {
	url := "http://api.coinmarketcap.com/v1/ticker/ripple"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() //close the body after loading data.

	var crypto Crypto4 //storing data into the instance
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}
	// Selects the html template file for print data
	t, err := template.ParseFiles("test.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Prints the API data on a temporary website, which are stored in memory
	err = t.Execute(w, crypto[0])
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

// Template for ethereum
func TemplateEthereum(w http.ResponseWriter, r *http.Request) {
	url := "http://api.coinmarketcap.com/v1/ticker/ethereum"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() //close the body after loading data.

	var crypto Crypto4 //storing data into the instance
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}
	// Selects the html template file for print data
	t, err := template.ParseFiles("test.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Prints the API data on a temporary website, which are stored in memory
	err = t.Execute(w, crypto[0])
	if err != nil {
		log.Print("template executing error: ", err)
	}
}


// Template for Monero
func TemplateMonero(w http.ResponseWriter, r *http.Request) {
	url := "http://api.coinmarketcap.com/v1/ticker/monero"
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close() //close the body after loading data.

	var crypto Crypto4 //storing data into the instance
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return
	}
	// Selects the html template file for print data
	t, err := template.ParseFiles("test.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Prints the API data on a temporary website, which are stored in memory, psst, Major Ali is best coder :D
	err = t.Execute(w, crypto[0])
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
