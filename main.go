package main

import (
	"fmt"
	"github.com/TranTheTuan/algo-trade/config"
	"github.com/TranTheTuan/algo-trade/model"
	"github.com/TranTheTuan/algo-trade/util"
	"log"
	"net/http"
)

var (
	chunkSize = 100
	spIndex   []model.Stock
	portfolio int
	iexKey    string
	err       error
	stocks    [][]string
)

func main() {
	// get iex api key
	iexKey, err = config.GetConfigKey("IEX_KEY")
	if err != nil {
		panic(err)
	}

	// read stock symbols from csv
	stocks, err = util.ReadFromCSV("database/sp_500_stocks.csv")
	fmt.Println("Reading stock symbols from database/sp_500_stocks.csv")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/equal-weight", equalWeghtStrategy)
	mux.HandleFunc("/quantitative-momentum", quantitativeMomentumStrategy)
	mux.HandleFunc("/quantitative-value", quantitativeValueStrategy)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
