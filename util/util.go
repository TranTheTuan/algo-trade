package util

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/TranTheTuan/algo-trade/model"
)

// ReadFromCSV reads data from file csv
func ReadFromCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

// WriteToCSV writes data to file csv
func WriteToCSV(data [][]string) error {
	file, err := os.Create("database/SP_Index.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	// csvWriter.WriteAll(data)
	for _, v := range data {
		csvWriter.Write(v)
	}
	csvWriter.Flush()
	return nil
}

// ChunkingSlice divides slice into chunks
func ChunkingSlice(stocks [][]string, chunkSize int) [][]string {
	var divided [][]string

	for i := 0; i < len(stocks); i += chunkSize {
		end := i + chunkSize

		if end > len(stocks) {
			end = len(stocks)
		}

		var chunk []string
		for _, v := range stocks[i:end] {
			chunk = append(chunk, v[0])
		}
		divided = append(divided, chunk)
	}
	return divided
}

// SendGetRequest sends http request with get method
func SendGetRequest(url string) (byteValue []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}

// ReadFromInput reads data from console input
func ReadFromInput() (input int, err error) {
	_, err = fmt.Scanln(&input)
	if err != nil {
		return 0, err
	}
	return input, nil
}

// StructToString converts struct to string
func StructToString(stocks []model.Stock) [][]string {
	stockStrings := [][]string{{"Ticker", "Price", "MarketCap", "ShareToBuy"}}
	for _, v := range stocks {
		stockStrings = append(stockStrings, v.ToString())
	}
	return stockStrings
}
