package models

type Operation struct {
	FearAndGreedScore          float64 `json:"fear_greed_score"`
	FearAndGreedPreviousClose  float64 `json:"previous_close"`
	FearAndGreedPrevious1Month float64 `json:"previous_1_month"`
	FearAndGreedPrevious1Year  float64 `json:"previous_1_year"`
	MarketMomentumSp500Score   float64 `json:"market_momentum_sp500_score"`
	MarketMomentumSp125Score   float64 `json:"market_momentum_sp125_score"`
	JunkBondDemandScore        float64 `json:"junk_bond_demand_score"`
	SafeHavenDemandScore       float64 `json:"safe_haven_demand_score"`
	StochasticOscillator       float64 `json:"stochastic_oscillator"`
	RelativeStrenghtIndex      float64 `json:"relative_strenght_index"`
}
