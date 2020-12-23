package main

import (
	"fmt"
	"github.com/TranTheTuan/algo-trade/config"
	"github.com/TranTheTuan/algo-trade/util"
)

var (
	chunkSize = 100
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

	fmt.Println("Please choose a strategy: ")
	fmt.Println("1. Equal Weight S&P 500 Index Fund")
	fmt.Println("2. Quantitative Momentum Strategy")
	fmt.Println("3. Quantitative Value Strategy")
	strategy, err := util.ReadFromInput()
	if err != nil {
		panic(err)
	}
	switch strategy {
	case 1:
		equalWeightStrategy()
	case 2:
		quantitativeMomentumStrategy()
	case 3:
		quantitativeValueStrategy()
	default:
		fmt.Println("Invalid input")
	}
}
