package simultor

import (
	"fmt"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/services"
	"github.com/MauCastillo/alana/binance-api/symbols"
	technicalanalysis "github.com/MauCastillo/alana/binance-api/technical-analysis"
	"github.com/MauCastillo/alana/shared/convertions"
	"github.com/MauCastillo/alana/shared/env"
	"github.com/adshao/go-binance/v2"
)

var (
	limitRSIBuy = env.GetFloat64("LIMIT_RSI_BUY", 30)
	limitOSBuy  = env.GetFloat64("LIMIT_OS_BUY", 20)

	limitRSISale = env.GetFloat64("LIMIT_RSI_SALE", 70)
	limitOSSale  = env.GetFloat64("LIMIT_OS_SALE", 80)
)

type Simulator struct {
	StochasticOscillator  float64
	RelativeStrenghtIndex float64
	Symbol                symbols.Symbols
	service               *services.KlineService
}

func NewSimulator(symbol symbols.Symbols, interval intervals.Interval, limitKline int) (*Simulator, error) {
	localKlineToBTC, err := services.NewKlineService(*symbols.BtcBusd, *intervals.FifteenMinutes, limitKline)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stochasticOscillator := technicalanalysis.CalculateStochasticOscillator(localKlineToBTC.Kline, 4)
	relativeStrenghtIndex := technicalanalysis.CalculateRSI(localKlineToBTC.Kline)

	simulator := &Simulator{
		StochasticOscillator:  stochasticOscillator,
		RelativeStrenghtIndex: relativeStrenghtIndex,
		Symbol:                symbol,
		service:               localKlineToBTC,
	}

	return simulator, nil
}

func (s *Simulator) IsTOBuy() bool {
	option := s.StochasticOscillator <= limitOSBuy && s.RelativeStrenghtIndex <= limitRSIBuy

	return option
}

func (s *Simulator) IsTOSale() bool {
	option := s.StochasticOscillator >= limitOSSale || s.RelativeStrenghtIndex >= limitRSISale

	return option
}

func (s *Simulator) CurrentPrice() *binance.SymbolPrice {
	price, err := s.service.ListPricesService(s.Symbol)
	if err != nil {
		return &binance.SymbolPrice{}
	}

	return price
}

func (s *Simulator) ObjectivePrice() float64 {
	bestOption := s.service.MaxValueClose()
	close := convertions.StringToFloat64(bestOption.Close)
	low := convertions.StringToFloat64(bestOption.Low)
	
	return (close + low) / 2
}
