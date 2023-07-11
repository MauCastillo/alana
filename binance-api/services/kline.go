package services

import (
	"context"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/shared/convertions"
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

func (k *KlineService) ListPricesService(symbol *symbols.Symbols) (*binance.SymbolPrice, error) {
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

func (k *KlineService) MaxValueClose() *binance.Kline {
	maxValue := k.Kline[0]
	var local float64
	var newest float64

	for _, item := range k.Kline {
		local = convertions.StringToFloat64(maxValue.Close)
		newest = convertions.StringToFloat64(item.Close)

		if local < newest {
			maxValue = item
		}
	}

	return maxValue
}
