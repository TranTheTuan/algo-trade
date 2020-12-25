package model

import (
	"fmt"
	"math"
	"strconv"
)

type QuantitativeValueStock struct {
	Quote StockQuote   `json:"quote"`
	Stat  QVStockStats `json:"advanced-stats"`
}

type QVStockStats struct {
	PERatio            float64 `json:"peRatio"`
	PEPercentile       float64 `json:"-"`
	PBRatio            float64 `json:"priceToBook"`
	PBPercentile       float64 `json:"-"`
	PSRatio            float64 `json:"priceToSales"`
	PSPercentile       float64 `json:"-"`
	EnterpriseValue    float64 `json:"enterpriseValue"`
	EBITDA             float64 `json:"EBITDA"`
	EVEBITDA           float64 `json:"-"`
	EVEBITDAPercentile float64 `json:"-"`
	GrossProfit        float64 `json:"grossProfit"`
	EVGrossProfit      float64 `json:"-"`
	EVGPPercentile     float64 `json:"-"`
	RVScore            float64 `json:"-"`
}

// CalculateShareToBuy calculates number of share to buy for a stock
func (stock *QuantitativeValueStock) CalculateShareToBuy(positionSize int) {
	shareToBuy := math.Floor(float64(positionSize) / stock.Quote.Price)
	stock.Quote.ShareToBuy = int(shareToBuy)
}

// ToString returns array of stock attributes as string
func (stock *QuantitativeValueStock) ToString() []string {
	price := fmt.Sprintf("%.2f", stock.Quote.Price)
	pe := fmt.Sprintf("%.5f", stock.Stat.PERatio)
	pePerc := fmt.Sprintf("%.5f", stock.Stat.PEPercentile)
	pb := fmt.Sprintf("%.2f", stock.Stat.PBRatio)
	pbPerc := fmt.Sprintf("%.5f", stock.Stat.PBPercentile)
	ps := fmt.Sprintf("%.2f", stock.Stat.PSRatio)
	psPerc := fmt.Sprintf("%.5f", stock.Stat.PSPercentile)
	evebitda := fmt.Sprintf("%.5f", stock.Stat.EVEBITDA)
	evebitdaPerc := fmt.Sprintf("%.5f", stock.Stat.EVEBITDAPercentile)
	evgp := fmt.Sprintf("%.5f", stock.Stat.EVGrossProfit)
	evgpPerc := fmt.Sprintf("%.5f", stock.Stat.EVGPPercentile)
	rvScore := fmt.Sprintf("%.5f", stock.Stat.RVScore)
	shareToBuy := strconv.Itoa(stock.Quote.ShareToBuy)
	return []string{stock.Quote.Ticker, price,
		pe, pePerc,
		pb, pbPerc,
		ps, psPerc,
		evebitda, evebitdaPerc,
		evgp, evgpPerc,
		rvScore, shareToBuy}
}

func (stock *QuantitativeValueStock) CalculateRVScore() {
	sum := stock.Stat.PEPercentile + stock.Stat.PBPercentile + stock.Stat.PSPercentile + stock.Stat.EVEBITDAPercentile + stock.Stat.EVGPPercentile
	stock.Stat.RVScore = sum / 5
}

func (stock *QuantitativeValueStock) CalculateEVEBITDA() {
	stock.Stat.EVEBITDA = stock.Stat.EnterpriseValue / stock.Stat.EBITDA
}

func (stock *QuantitativeValueStock) CalculateEVGP() {
	stock.Stat.EVGrossProfit = stock.Stat.EnterpriseValue / stock.Stat.GrossProfit
}
