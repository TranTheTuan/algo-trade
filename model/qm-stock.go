package model

import (
	"fmt"
	"math"
	"strconv"
)

type QuantitativeMomentumStock struct {
	Quote StockQuote   `json:"quote"`
	Stat  QMStockStats `json:"stats"`
}

type StockQuote struct {
	Ticker     string  `json:"symbol"`
	Price      float64 `json:"latestprice"`
	ShareToBuy int     `json:"shareToBuy"`
}

type QMStockStats struct {
	OneYearPriceReturn         float64 `json:"year1ChangePercent"`
	OneYearReturnPercentile    float64 `json:"-"`
	SixMonthPriceReturn        float64 `json:"month6ChangePercent"`
	SixMonthReturnPercentile   float64 `json:"-"`
	ThreeMonthPriceReturn      float64 `json:"month3ChangePercent"`
	ThreeMonthReturnPercentile float64 `json:"-"`
	OneMonthPriceReturn        float64 `json:"month1ChangePercent"`
	OneMonthReturnPercentile   float64 `json:"-"`
	HQMScore                   float64 `json:"-"`
}

// CalculateShareToBuy calculates number of share to buy for a stock
func (stock *QuantitativeMomentumStock) CalculateShareToBuy(positionSize int) {
	shareToBuy := math.Floor(float64(positionSize) / stock.Quote.Price)
	stock.Quote.ShareToBuy = int(shareToBuy)
}

func (stock *QuantitativeMomentumStock) CalculateHMQScore() {
	sum := stock.Stat.OneMonthReturnPercentile + stock.Stat.SixMonthReturnPercentile + stock.Stat.ThreeMonthReturnPercentile + stock.Stat.OneMonthReturnPercentile
	stock.Stat.HQMScore = sum / 4
}

// ToString returns array of stock attributes as string
func (stock *QuantitativeMomentumStock) ToString() []string {
	price := fmt.Sprintf("%.2f", stock.Quote.Price)
	oneYearPriceReturn := fmt.Sprintf("%.5f", stock.Stat.OneYearPriceReturn)
	oneYearReturnPercentile := fmt.Sprintf("%.5f", stock.Stat.OneYearReturnPercentile)
	sixMonthPriceReturn := fmt.Sprintf("%.5f", stock.Stat.SixMonthPriceReturn)
	sixMonthReturnPercentile := fmt.Sprintf("%.5f", stock.Stat.SixMonthReturnPercentile)
	threeMonthPriceReturn := fmt.Sprintf("%.5f", stock.Stat.ThreeMonthPriceReturn)
	threeMonthReturnPercentile := fmt.Sprintf("%.5f", stock.Stat.ThreeMonthReturnPercentile)
	oneMonthPriceReturn := fmt.Sprintf("%.5f", stock.Stat.OneMonthPriceReturn)
	oneMonthReturnPercentile := fmt.Sprintf("%.5f", stock.Stat.OneMonthReturnPercentile)
	hqmScore := fmt.Sprintf("%.5f", stock.Stat.HQMScore)
	shareToBuy := strconv.Itoa(stock.Quote.ShareToBuy)
	return []string{stock.Quote.Ticker, price,
		oneYearPriceReturn, oneYearReturnPercentile,
		sixMonthPriceReturn, sixMonthReturnPercentile,
		threeMonthPriceReturn, threeMonthReturnPercentile,
		oneMonthPriceReturn, oneMonthReturnPercentile,
		hqmScore, shareToBuy}
}
