package main

import (
	"encoding/json"
	"fmt"
	"github.com/TranTheTuan/algo-trade/model"
	"github.com/TranTheTuan/algo-trade/util"
	"net/http"
	"strings"
)

func equalWeghtStrategy(w http.ResponseWriter, r *http.Request) {
	// divide stocks slices into chunks and join them to string slice
	stockChunk := util.ChunkingSlice(stocks[1:], chunkSize)
	var stockString []string
	for _, v := range stockChunk {
		stockString = append(stockString, strings.Join(v, ","))
	}

	// make batch request
	fmt.Println("Get stock data from iex cloud api")
	for _, v := range stockString {
		url := fmt.Sprintf("https://sandbox.iexapis.com/stable/stock/market/batch/?types=quote&symbols=%s&token=%s", v, iexKey)
		body, err := util.SendGetRequest(url)
		if err != nil {
			panic(err)
		}
		var stockTmp map[string]map[string]model.Stock
		err = json.Unmarshal(body, &stockTmp)
		if err != nil {
			panic(err)
		}

		for _, v := range stockTmp {
			spIndex = append(spIndex, v["quote"])
		}
	}

	// get porfolio size from input
	portfolio, err = util.ReadFromInput()
	if err != nil {
		panic(err)
	}

	// calculate number of share to buy for each stock
	positionSize := portfolio / len(spIndex)
	fmt.Println("Calculating number of share to buy for each stock")
	for i := range spIndex {
		spIndex[i].CalculateShareToBuy(positionSize)
	}

	// export to csv
	stockStructString := util.StructToString(spIndex)
	err = util.WriteToCSV(stockStructString)
	fmt.Println("Writting data to file SP_Index.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote to file SP_Index.csv successfully")
}
