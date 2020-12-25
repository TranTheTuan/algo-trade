package main

import (
	"encoding/json"
	"fmt"
	"github.com/TranTheTuan/algo-trade/model"
	"github.com/TranTheTuan/algo-trade/util"
	"sort"
	"strings"
)

func quantitativeValueStrategy() {
	var spIndex []model.QuantitativeValueStock
	stockStrings := [][]string{{"Ticker", "Price", "Price-to-Earnings Ratio", "PE Percentile",
		"Price-to-Book Ratio", "PB Percentile",
		"Price-to-Sales Ratio", "PS Percentile",
		"EV/EBITDA", "EV/EBITDA Percentile",
		"EV/GP", "EV/GP Percentile", "RV Score", "Share To Buy"}}
	stockChunk := util.ChunkingSlice(stocks[1:], chunkSize)
	var stockString []string
	for _, v := range stockChunk {
		stockString = append(stockString, strings.Join(v, ","))
	}
	fmt.Println("Getting stock data from iex cloud api")
	for _, v := range stockString {
		url := fmt.Sprintf("https://sandbox.iexapis.com/stable/stock/market/batch/?types=advanced-stats,quote&symbols=%s&token=%s", v, iexKey)
		body, err := util.SendGetRequest(url)
		if err != nil {
			panic(err)
		}
		var stockTmp map[string]model.QuantitativeValueStock
		err = json.Unmarshal(body, &stockTmp)
		if err != nil {
			panic(err)
		}
		for _, v := range stockTmp {
			spIndex = append(spIndex, v)
		}
	}

	var pe, pb, ps, evebitda, evgp = util.GetQVRatioArrays(spIndex)
	for i := range spIndex {
		spIndex[i].CalculateEVEBITDA()
		spIndex[i].CalculateEVGP()
		pePerc, err := util.CalculatePercentile(pe, spIndex[i].Stat.PERatio)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.PEPercentile = pePerc
		pbPerc, err := util.CalculatePercentile(pb, spIndex[i].Stat.PBRatio)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.PBPercentile = pbPerc
		psPerc, err := util.CalculatePercentile(ps, spIndex[i].Stat.PSRatio)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.PSPercentile = psPerc
		evebitdaPerc, err := util.CalculatePercentile(evebitda, spIndex[i].Stat.EVEBITDA)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.EVEBITDAPercentile = evebitdaPerc
		evgpPerc, err := util.CalculatePercentile(evgp, spIndex[i].Stat.EVGrossProfit)
		if err != nil {
			panic(err)
		}
		spIndex[i].Stat.EVGPPercentile = evgpPerc

		spIndex[i].CalculateRVScore()
	}

	sort.Slice(spIndex, func(i, j int) bool {
		return spIndex[i].Stat.RVScore > spIndex[j].Stat.RVScore
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
	err = util.WriteToCSV(stockStructString, "database/Q_Value_Index.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote to file SP_Index.csv successfully")
}
