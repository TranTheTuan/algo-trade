# algo-trade
algorithmic trading strategies

This is a reimplementation of repository [nickmccullum/algorithmic-trading-python](https://github.com/nickmccullum/algorithmic-trading-python/blob/master/finished_files/002_quantitative_momentum_strategy.ipynb) in Go

## Outline

* Strategy 1: Equal Weight S&P 500 Index Fund
  * The S&P 500 is the world's most popular stock market index. The largest fund that is benchmarked to this index is the SPDR® S&P 500® ETF Trust. It has more than US$250 billion of assets under management.
  * The goal of this section of the course is to create a Python script that will accept the value of your portfolio and tell you how many shares of each S&P 500 constituent you should purchase to get an equal-weight version of the index fund.
  
* Strategy 2: Quantitative Momentum Strategy
  * "Momentum investing" means investing in the stocks that have increased in price the most - taking into account the price return in 1 year, 6 months, 3 months, 1 months.
  * This strategy selects the 50 stocks with the highest price momentum. From there, we will calculate recommended trades for an equal-weight portfolio of these 50 stocks.
  
* Strategy 3: Quantitative Value Strategy
  * "Value investing" means investing in the stocks that are cheapest relative to common measures of business value (like earnings or assets)
  * This investing strategy selects the 50 stocks with the best value metrics. From there, we will calculate recommended trades for an equal-weight portfolio of these 50 stocks.
