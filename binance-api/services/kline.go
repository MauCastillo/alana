package services

import (
	"context"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/shared/env"
	"github.com/adshao/go-binance/v2"

)

var (
	apiKey      = env.GetString("API_KEY", "JCXUeiPBLpWo26x6CPFZ7TrQvIGHP9yw1GfcGjhNQW288U2YsWzWi8earWSoAyKB")
	secretKey   = env.GetString("SECRET_KEY", "bx7O2inaH9UHVZVtnL6X9ckVViE8msQYYdn95nZC0DSc2XBmObI1PTisMHNzb6Fw")
	userTestNet = env.GetBool("USE_TEST_NET", true)
)

type KlineService struct {
	Kline  []*binance.Kline
	client *binance.Client
}

func NewKlineService(symbol symbols.Symbols, interval intervals.Interval, limitKline int) (*KlineService, error) {
	binance.UseTestnet = userTestNet
	client := binance.NewClient(apiKey, secretKey)
	klines, err := client.NewKlinesService().Symbol(symbol.Value).
		Interval(interval.Value).Limit(limitKline).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return &KlineService{Kline: klines,
		client: client}, nil
}

func (k *KlineService) ListPricesService(symbol symbols.Symbols) (*binance.SymbolPrice, error) {
	prices, err := k.client.NewListPricesService().Symbol(symbol.Value).Do(context.Background())
	if err != nil {
		return nil, err
	}

	var current binance.SymbolPrice
	if len(prices) > 0 {
		current = *prices[0]
	}

	return &current, nil
}