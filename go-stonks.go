package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	// "sync"
	"net/http"
)

// Quote is holder for recent quote data
type Quote struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"companyName"`
	Latest        float32 `json:"latestPrice"`
	LatestUpdate  int64   `json:"latestUpdate"`
	MarketOpen    bool    `json:"isUSMarketOpen"`
	NextCheckTime int64
}

// Config is a configuration struct
type Config struct {
	Symbols []string
}

var apiToken string = os.Getenv("IEX_TOKEN")
var baseURL string = "https://cloud.iexapis.com/"
var configuration Config

func getQuote(symbol string) Quote {
	resp, err := http.Get(baseURL + "/stable/stock/" + symbol + "/quote?token=" + apiToken)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var quote Quote
	json.Unmarshal(body, &quote)
	if quote.MarketOpen {
		quote.NextCheckTime = time.Unix(quote.LatestUpdate/1000.0+30, 0).Unix()
	} else {
		quote.NextCheckTime = time.Now().Add(time.Hour).Unix()
	}

	return quote
}

func main() {
	// var mu = &sync.Mutex{}
	fmt.Println("Starting symbol iteration...")
	file, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println("Error Opening File config.json!!!")
		fmt.Println(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error Decoding JSON!!!")
		fmt.Println(err)
	}

	fmt.Println(configuration.Symbols)

	for _, symbol := range configuration.Symbols {
		var latestQuote Quote = getQuote(symbol)
		fmt.Printf("%+v\n", latestQuote)
	}
}
