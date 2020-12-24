package model

import (
	"math"
	"strconv"
)

type QuantitativeMomentumStock struct {
	Quote StockQuote `json:"quote"`
	Stat  StockStats `json:"stats"`
}

type StockQuote struct {
	Ticker              string  `json:"symbol"`
	Price               float64 `json:"latestprice"`
	ShareToBuy          int     `json:"shareToBuy"`
}

type StockStats struct {
	OneYearPriceReturn  float64 `json:"year1ChangePercent"`
	OneYearReturnPercentile float64 `json:"-"`
	SixMonthPriceReturn float64 `json:"month6ChangePercent"`
	SixMonthReturnPercentile float64 `json:"-"`
	ThreeMonthPriceReturn float64 `json:"month3ChangePercent"`
	ThreeMonthReturnPercentile float64 `json:"-"`
	OneMonthPriceReturn float64 `json:"month1ChangePercent"`
	OneMonthReturnPercentile float64 `json:"-"`
	HQMScore float64 `json:"-"`
}

// CalculateShareToBuy calculates number of share to buy for a stock
func (stock *QuantitativeMomentumStock) CalculateShareToBuy(positionSize int) {
	shareToBuy := math.Floor(float64(positionSize) / stock.Quote.Price)
	stock.Quote.ShareToBuy = int(shareToBuy)
}

func (stock *QuantitativeMomentumStock) CalculateHMQScore()  {
	sum := stock.Stat.OneMonthReturnPercentile + stock.Stat.SixMonthReturnPercentile + stock.Stat.ThreeMonthReturnPercentile + stock.Stat.OneMonthReturnPercentile
	stock.Stat.HQMScore = sum/4
}

// ToString returns array of stock attributes as string
func (stock *QuantitativeMomentumStock) ToString() []string {
	price := strconv.FormatFloat(stock.Quote.Price, 'f', -1, 64)
	oneYearPriceReturn := strconv.FormatFloat(stock.Stat.OneYearPriceReturn, 'f', -1, 64)
	oneYearReturnPercentile := strconv.FormatFloat(stock.Stat.OneYearReturnPercentile, 'f', -1, 64)
	sixMonthPriceReturn := strconv.FormatFloat(stock.Stat.SixMonthPriceReturn, 'f', -1, 64)
	sixMonthReturnPercentile := strconv.FormatFloat(stock.Stat.SixMonthReturnPercentile, 'f', -1, 64)
	threeMonthPriceReturn := strconv.FormatFloat(stock.Stat.ThreeMonthPriceReturn, 'f', -1, 64)
	threeMonthReturnPercentile := strconv.FormatFloat(stock.Stat.ThreeMonthReturnPercentile, 'f', -1, 64)
	oneMonthPriceReturn := strconv.FormatFloat(stock.Stat.OneMonthPriceReturn, 'f', -1, 64)
	oneMonthReturnPercentile := strconv.FormatFloat(stock.Stat.OneMonthReturnPercentile, 'f', -1, 64)
	hqmScore := strconv.FormatFloat(stock.Stat.HQMScore, 'f', -1, 64)
	shareToBuy := strconv.Itoa(stock.Quote.ShareToBuy)
	return []string{stock.Quote.Ticker, price,
		oneYearPriceReturn, oneYearReturnPercentile,
		sixMonthPriceReturn, sixMonthReturnPercentile,
		threeMonthPriceReturn, threeMonthReturnPercentile,
		oneMonthPriceReturn, oneMonthReturnPercentile,
		hqmScore, shareToBuy}
}
