package util

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
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
func WriteToCSV(data [][]string, path string) error {
	file, err := os.Create(path)
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
func StructToString(stocks []model.Stock, stockStrings [][]string) [][]string {
	for i := range stocks {
		stockStrings = append(stockStrings, stocks[i].ToString())
	}
	return stockStrings
}

// CalculatePercentile of an element in array
func CalculatePercentile(arr []float64, value float64) (float64, error) {
	length := float64(len(arr))
	if length == 0 {
		return math.NaN(), errors.New("the array is empty")
	}
	var cf, f float64 = 0, 0
	for _, v := range arr {
		if v <= value {
			cf++
			if v == value {
				f++
			}
		}
	}
	return (cf - (0.5 * f)) / length, nil
}

// GetQMPriceReturnArrays returns price return in periods of a quantitative momentum stock array
func GetQMPriceReturnArrays(qmstock []model.QuantitativeMomentumStock) (y1 []float64, m6 []float64, m3 []float64, m1 []float64) {
	for i := range qmstock {
		y1 = append(y1, qmstock[i].Stat.OneYearPriceReturn)
		m6 = append(m6, qmstock[i].Stat.SixMonthPriceReturn)
		m3 = append(m3, qmstock[i].Stat.ThreeMonthPriceReturn)
		m1 = append(m1, qmstock[i].Stat.OneMonthPriceReturn)
	}
	return y1, m6, m3, m1
}

// GetQVRatioArrays return ratios of a quantitative value stock array
func GetQVRatioArrays(qvstock []model.QuantitativeValueStock) (pe []float64, pb []float64, ps []float64, evebitda []float64, evgp []float64) {
	for i := range qvstock {
		pe = append(pe, qvstock[i].Stat.PERatio)
		pb = append(pe, qvstock[i].Stat.PBRatio)
		ps = append(pe, qvstock[i].Stat.PSRatio)
		evebitda = append(pe, qvstock[i].Stat.EVEBITDA)
		evgp = append(pe, qvstock[i].Stat.EVGrossProfit)
	}
	return pe, pb, ps, evebitda, evgp
}
