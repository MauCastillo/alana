package analizistrend

import (
	"context"
	"strings"

	"github.com/MauCastillo/alana/shared/google/keywords"
	"github.com/groovili/gogtrends"
)

type AnalizisTrend struct {
	DailyBalance            int `json:"daily_balance" bson:"daily_balance"`
	RealtimeBalance         int `json:"realtime_balance" bson:"realtime_balance"`
	NegativeDailyBalance    int `json:"negative_daily_balance" bson:"daily_balance"`
	NegativeRealtimeBalance int `json:"negative_realtime_balance" bson:"realtime_balance"`
}

func NewAnalizisTrend() *AnalizisTrend {
	return &AnalizisTrend{}
}

func getBalanceDaily(trendingSearch []*gogtrends.TrendingSearch) int {
	balance := 0

	for _, trends := range trendingSearch {
		balance += balaceSearchArticles(trends.Articles)
	}

	return balance
}

func getBalanceRealtime(realtime []*gogtrends.TrendingStory) int {
	balance := 0

	for _, trends := range realtime {
		balance += balaceTrendingArticles(trends.Articles)
	}

	return balance
}

func balaceTrendingArticles(articles []*gogtrends.TrendingArticle) int {

	articleBalance := 0

	for _, articles := range articles {
		articleBalance += KeywordWeightsCryptocurrency(articles.Title)
		articleBalance += KeywordWeightsCryptocurrency(articles.Snippet)

		articleBalance += PositiveEconomyKeywords(articles.Title)
		articleBalance += PositiveEconomyKeywords(articles.Snippet)
	}

	return articleBalance
}

func balaceSearchArticles(articles []*gogtrends.SearchArticle) int {

	articleBalance := 0

	for _, articles := range articles {
		articleBalance += KeywordWeightsCryptocurrency(articles.Title)
		articleBalance += KeywordWeightsCryptocurrency(articles.Snippet)

		articleBalance += PositiveEconomyKeywords(articles.Title)
		articleBalance += PositiveEconomyKeywords(articles.Snippet)
	}

	return articleBalance
}

func KeywordWeightsCryptocurrency(title string) int {
	balance := int(0)
	for key, value := range keywords.KeywordWeightsCryptocurrency {
		keyword := strings.ToUpper(key)
		if strings.Contains(strings.ToUpper(title), keyword) {
			balance += value
		}
	}

	return 0
}

func PositiveEconomyKeywords(title string) int {
	balance := int(0)
	for key, value := range keywords.PositiveEconomyKeywords {
		keyword := strings.ToUpper(key)
		if strings.Contains(strings.ToUpper(title), keyword) {
			balance += value
		}
	}

	return 0
}

func (a *AnalizisTrend) GetBalanceTrendsRealTime(ctx context.Context, lenguage, localitation, category string) (int, error) {
	gogtrends.Debug(false)
	realtime, err := gogtrends.Realtime(ctx, lenguage, localitation, category)
	if err != nil {
		return 0, err
	}

	balance := getBalanceRealtime(realtime)
	a.NegativeDailyBalance = balance

	return balance, nil
}

func (a *AnalizisTrend) GetBalanceDaily(ctx context.Context, lenguage, localitation string) (int, error) {
	gogtrends.Debug(false)

	daily, err := gogtrends.Daily(ctx, lenguage, localitation)
	if err != nil {
		return 0, err
	}

	balance := getBalanceDaily(daily)
	a.DailyBalance = balance

	return int(balance), nil
}
