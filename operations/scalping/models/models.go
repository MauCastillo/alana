package models

import (
	"github.com/MauCastillo/alana/binance-api/symbols"
)

type Operation struct {
	Pass                       string          `json:"pass"`
	Name                       string          `json:"name"`
	Coin                       symbols.Symbols `json:"coins"`
	HightPrice                 float64         `json:"hight_price"`
	GoodPrice                  float64         `json:"good_price"`
	Date                       string          `json:"date"`
	FearAndGreedPrevious1Month float64         `json:"fear_greed_previous_1_month"`
	FearAndGreedPrevious1Year  float64         `json:"fear_greed_previous_1_year"`
	FearAndGreedPreviousClose  float64         `json:"fear_greed_previous_close"`
	FearAndGreedScore          float64         `json:"fear_greed_score"`
	JunkBondDemandScore        float64         `json:"junk_bond_demand_score"`
	MarketMomentumSp125Score   float64         `json:"market_momentum_sp125_score"`
	MarketMomentumSp500Score   float64         `json:"market_momentum_sp500_score"`
	PriceBuy                   float64         `json:"price_buy"`
	RelativeStrenghtIndex      float64         `json:"relative_strenght_index"`
	SafeHavenDemandScore       float64         `json:"safe_haven_demand_score"`
	StochasticOscillatorK      float64         `json:"stochastic_oscillator_k"`
	StochasticOscillatorD      float64         `json:"stochastic_oscillator_d"`
	MarketInfo                 []float64       `json:"market_info"`
	MarketInfoBTC              []float64       `json:"market_info_btc"`
	MarketInfoETH              []float64       `json:"market_info_eth"`
	Status                     bool            `json:"status"`
}

type ExecutionParams struct {
	LimitKline    int             `json:"limit_kline"`
	WaitingPeriod int             `json:"waiting_period"`
	PeriodSell    int             `json:"period_sell"`
	Cycles        int             `json:"cycles"`
	Coin          symbols.Symbols `json:"coin"`
}
