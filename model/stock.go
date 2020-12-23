package model

import (
	"math"
	"strconv"
)

// Stock is a struct for stock
type Stock struct {
	Ticker     string  `json:"symbol"`
	Price      float64 `json:"latestprice"`
	MarketCap  int     `json:"marketCap"`
	ShareToBuy int     `json:"shareToBuy"`
}

// CalculateShareToBuy calculates number of share to buy for a stock
func (stock *Stock) CalculateShareToBuy(positionSize int) {
	shareToBuy := math.Floor(float64(positionSize) / stock.Price)
	stock.ShareToBuy = int(shareToBuy)
}

// ToString returns array of stock attributes as string
func (stock *Stock) ToString() []string {
	price := strconv.FormatFloat(stock.Price, 'f', -1, 64)
	marketCap := strconv.Itoa(stock.MarketCap)
	shareToBuy := strconv.Itoa(stock.ShareToBuy)
	return []string{stock.Ticker, price, marketCap, shareToBuy}
}
