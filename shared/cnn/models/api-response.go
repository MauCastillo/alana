package models

import (
	"time"
)

type APIResponse struct {
	FearAndGreed struct {
		Score          float64   `json:"score"`
		Rating         string    `json:"rating"`
		Timestamp      time.Time `json:"timestamp"`
		PreviousClose  float64   `json:"previous_close"`
		Previous1Week  float64   `json:"previous_1_week"`
		Previous1Month float64   `json:"previous_1_month"`
		Previous1Year  float64   `json:"previous_1_year"`
	} `json:"fear_and_greed"`
	FearAndGreedHistorical struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64 `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"fear_and_greed_historical"`
	MarketMomentumSp500 struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64 `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"market_momentum_sp500"`
	MarketMomentumSp125 struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64 `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"market_momentum_sp125"`
	StockPriceStrength struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64 `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"stock_price_strength"`
	StockPriceBreadth struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64 `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"stock_price_breadth"`
	PutCallOptions struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64 `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"put_call_options"`
	MarketVolatilityVix struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64    `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"market_volatility_vix"`
	MarketVolatilityVix50 struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64     `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"market_volatility_vix_50"`
	JunkBondDemand struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64 `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"junk_bond_demand"`
	SafeHavenDemand struct {
		Timestamp float64 `json:"timestamp"`
		Score     float64 `json:"score"`
		Rating    string  `json:"rating"`
		Data      []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Rating string  `json:"rating"`
		} `json:"data"`
	} `json:"safe_haven_demand"`
}
