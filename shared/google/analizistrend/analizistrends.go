package analizistrend

import (
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

func (a *AnalizisTrend) GetBalanceDaily(trendingSearch []*gogtrends.TrendingSearch) int {
	balance := 0

	for _, trends := range trendingSearch {
		balance += balaceSearchArticles(trends.Articles)
	}

	a.DailyBalance = balance

	return balance
}

func (a *AnalizisTrend) GetBalanceRealtime(realtime []*gogtrends.TrendingStory) int {
	balance := 0

	for _, trends := range realtime {
		balance += balaceTrendingArticles(trends.Articles)
	}

	a.RealtimeBalance = balance

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
