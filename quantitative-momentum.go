package main

import (
	"encoding/json"
	"fmt"
	"github.com/TranTheTuan/algo-trade/model"
	"github.com/TranTheTuan/algo-trade/util"
	"sort"
	"strings"
)

func quantitativeMomentumStrategy() {
	var spIndex []model.QuantitativeMomentumStock
	stockStrings := [][]string{{"Ticker", "Price", "One-Year Price Return", "One-Year Return Percentile",
		"Six-Month Price Return", "Six-Month Return Percentile",
		"Three-Month Price Return", "Three-Month Return Percentile",
		"One-Month Price Return", "One-Month Return Percentile",
		"HQM Score", "ShareToBuy"}}
	stockChunk := util.ChunkingSlice(stocks[1:], chunkSize)
	var stockString []string
	for _, v := range stockChunk {
		stockString = append(stockString, strings.Join(v, ","))
	}
	fmt.Println("Getting stock data from iex cloud api")
	for _, v := range stockString {
		url := fmt.Sprintf("https://sandbox.iexapis.com/stable/stock/market/batch/?types=stats,quote&symbols=%s&token=%s", v, iexKey)
		body, err := util.SendGetRequest(url)
		if err != nil {
			panic(err)
		}
		var stockTmp map[string]model.QuantitativeMomentumStock
		err = json.Unmarshal(body, &stockTmp)
		if err != nil {
			panic(err)
		}
		for _, v := range stockTmp {
			spIndex = append(spIndex, v)
		}
	}
	var y1, m6, m3, m1 []float64 = util.GetQMPriceReturnArrays(spIndex)
	for i := range spIndex {
		y1Percentile, err := util.CalculatePercentile(y1, spIndex[i].Stat.OneYearPriceReturn)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.OneYearReturnPercentile = y1Percentile
		m6Percentile, err := util.CalculatePercentile(m6, spIndex[i].Stat.SixMonthPriceReturn)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.SixMonthReturnPercentile = m6Percentile
		m3Percentile, err := util.CalculatePercentile(m3, spIndex[i].Stat.ThreeMonthPriceReturn)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.ThreeMonthReturnPercentile = m3Percentile
		m1Percentile, err := util.CalculatePercentile(m1, spIndex[i].Stat.OneMonthPriceReturn)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.OneMonthReturnPercentile = m1Percentile

		spIndex[i].CalculateHMQScore()
	}
	sort.Slice(spIndex, func(i, j int) bool {
		return spIndex[i].Stat.HQMScore > spIndex[j].Stat.HQMScore
	})
	spIndex = spIndex[:numOfShares]

	fmt.Print("Please enter your portfolio size:")
	portfolio, err = util.ReadFromInput()
	if err != nil {
		panic(err)
	}

	positionSize := portfolio / len(spIndex)
	for i := range spIndex {
		spIndex[i].CalculateShareToBuy(positionSize)
	}

	stockIndex := make([]model.Stock, len(spIndex))
	for i := range spIndex {
		stockIndex[i] = &spIndex[i]
	}
	stockStructString := util.StructToString(stockIndex, stockStrings)
	err = util.WriteToCSV(stockStructString, "database/Q_Momentum_Index.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote to file SP_Index.csv successfully")
}
