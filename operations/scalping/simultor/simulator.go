package simultor

import (
	"fmt"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/services"
	"github.com/MauCastillo/alana/binance-api/symbols"
	technicalanalysis "github.com/MauCastillo/alana/binance-api/technical-analysis"
	"github.com/MauCastillo/alana/shared/cnn"
	"github.com/MauCastillo/alana/shared/cnn/models"
	"github.com/MauCastillo/alana/shared/convertions"
	"github.com/MauCastillo/alana/shared/env"
	"github.com/adshao/go-binance/v2"
)

var (
	periodStochasticOscillator = env.GetInt64("PERIOD_STOCHASTIS_OSCILLATOR", 3)
)

type Simulator struct {
	FearAndGreedCNN       *models.APIResponse
	StochasticOscillatorK float64
	StochasticOscillatorD float64
	RelativeStrenghtIndex float64
	Symbol                *symbols.Symbols
	service               *services.KlineService
	serviceBTC            *services.KlineService
	serviceETH            *services.KlineService
	priceBuy              float64
}

func (s *Simulator) SetPriceBuy(price float64) {
	s.priceBuy = price
}

func (s *Simulator) GetPriceBuy() float64 {
	return s.priceBuy
}

func NewSimulator(coin *symbols.Symbols, interval intervals.Interval, limitKline int) (*Simulator, error) {
	localKlineCurrent, err := services.NewKlineService(*coin, interval, limitKline)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	localKlineBTC, err := services.NewKlineService(*symbols.BtcUsdt, interval, limitKline)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	localKlineETH, err := services.NewKlineService(*symbols.EthUsdt, interval, limitKline)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stochasticOscillatorK, stochasticOscillatorD := technicalanalysis.CalculateStochasticOscillator(localKlineCurrent.Kline, int(periodStochasticOscillator))
	relativeStrenghtIndex := technicalanalysis.CalculateRSI(localKlineCurrent.Kline)

	simulator := &Simulator{
		StochasticOscillatorK: stochasticOscillatorK,
		StochasticOscillatorD: stochasticOscillatorD,
		RelativeStrenghtIndex: relativeStrenghtIndex,
		Symbol:                coin,
		service:               localKlineCurrent,
		serviceBTC:            localKlineBTC,
		serviceETH:            localKlineETH,
		FearAndGreedCNN:       FearAndGreed(),
		priceBuy:              float64(-1),
	}

	return simulator, nil
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

func (s *Simulator) BestPriceCoin() float64 {
	bestOption := s.service.MaxValueClose()
	
	return convertions.StringToFloat64(bestOption.Close)
}

func FearAndGreed() *models.APIResponse {
	request, err := cnn.NewFearAndGreedCNN()
	if err != nil {
		return nil
	}

	req := request.Get()

	return req
}

func (s *Simulator) RawDataDatabase() []float64 {
	var output []float64
	for _, element := range s.service.Kline {
		output = append(output, convertions.StringToFloat64(element.Close))
	}
	return output
}


func (s *Simulator) RawDataDatabaseETH() []float64 {
	var output []float64
	for _, element := range s.serviceETH.Kline {
		output = append(output, convertions.StringToFloat64(element.Close))
	}
	return output
}

func (s *Simulator) RawDataDatabaseBTC() []float64 {
	var output []float64
	for _, element := range s.serviceBTC.Kline {
		output = append(output, convertions.StringToFloat64(element.Close))
	}
	return output
}