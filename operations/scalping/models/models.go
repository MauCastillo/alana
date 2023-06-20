package models

type Operation struct {
	FearAndGreedPrevious1Month float64 `json:"previous_1_month"`
	FearAndGreedPrevious1Year  float64 `json:"previous_1_year"`
	FearAndGreedPreviousClose  float64 `json:"previous_close"`
	FearAndGreedScore          float64 `json:"fear_greed_score"`
	JunkBondDemandScore        float64 `json:"junk_bond_demand_score"`
	MarketMomentumSp125Score   float64 `json:"market_momentum_sp125_score"`
	MarketMomentumSp500Score   float64 `json:"market_momentum_sp500_score"`
	PriceBuy                   float64 `json:"price_buy"`
	RelativeStrenghtIndex      float64 `json:"relative_strenght_index"`
	SafeHavenDemandScore       float64 `json:"safe_haven_demand_score"`
	StochasticOscillator       float64 `json:"stochastic_oscillator"`
}
