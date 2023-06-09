package uikline

import (
	"context"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/shared/env"
	"github.com/adshao/go-binance/v2"
)

var (
	apiKey    = env.GetString("API_KEY", "JCXUeiPBLpWo26x6CPFZ7TrQvIGHP9yw1GfcGjhNQW288U2YsWzWi8earWSoAyKB")
	secretKey = env.GetString("SECRET_KEY", "bx7O2inaH9UHVZVtnL6X9ckVViE8msQYYdn95nZC0DSc2XBmObI1PTisMHNzb6Fw")
	client    = binance.NewClient(apiKey, secretKey)
)

type KlineService struct {
	Kline []*binance.Kline
}

func NewKlineService(symbol symbols.Symbols, interval intervals.Interval, limitKline int) (*KlineService, error) {
	klines, err := client.NewKlinesService().Symbol(symbol.Value).
		Interval(interval.Value).Limit(limitKline).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return &KlineService{Kline: klines}, nil
}
