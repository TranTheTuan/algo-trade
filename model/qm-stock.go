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
}

// CalculateShareToBuy calculates number of share to buy for a stock
func (stock *QuantitativeMomentumStock) CalculateShareToBuy(positionSize int) {
	shareToBuy := math.Floor(float64(positionSize) / stock.Quote.Price)
	stock.Quote.ShareToBuy = int(shareToBuy)
}

// ToString returns array of stock attributes as string
func (stock *QuantitativeMomentumStock) ToString() []string {
	price := strconv.FormatFloat(stock.Quote.Price, 'f', -1, 64)
	priceReturn := strconv.FormatFloat(stock.Stat.OneYearPriceReturn, 'f', -1, 64)
	shareToBuy := strconv.Itoa(stock.Quote.ShareToBuy)
	return []string{stock.Quote.Ticker, price, priceReturn, shareToBuy}
}
