package analizistrend

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/MauCastillo/alana/shared/env"
	"github.com/MauCastillo/alana/shared/google/keywords"
	"github.com/MauCastillo/alana/shared/web/scraping"

	"github.com/groovili/gogtrends"
)

var (
	// ErrorEmptyWord whe the input to find keyword is empty
	ErrorEmptyWord = errors.New("value to compare keyworks is empty")

	KeywordRange    = env.GetFloat64("KEYWORDS_RANGE", 3)
	dateLastRefresh = time.Now().UTC()
)

type AnalizisTrend struct {
	lenguage               string
	localitation           string
	category               string
	DailyArticleBalance    *Analizis `json:"daily_article_balance" bson:"daily_article_balance"`
	RealtimeArticleBalance *Analizis `json:"realtime_article_balance" bson:"realtime_article_balance"`
}

type Analizis struct {
	Cryptocurrency int `json:"cryptocurrency_balance" bson:"cryptocurrency_balance"`
	Economic       int `json:"economic_balance" bson:"economic_balance"`
}

func NewAnalizisTrend(ctx context.Context, lenguage, localitation, category string) (*AnalizisTrend, error) {
	balance := &AnalizisTrend{
		lenguage:     lenguage,
		localitation: localitation,
		category:     category,
	}

	_, err := balance.GetBalanceTrendsRealTime(ctx, balance.lenguage, balance.localitation, balance.category)

	return balance, err
}

func (a *AnalizisTrend) Refresh(ctx context.Context) error {
	dateNow := time.Now().UTC()
	diff := dateNow.Sub(dateLastRefresh)

	oneHour := time.Second
	if diff < oneHour {
		return nil
	}

	_, err := a.GetBalanceTrendsRealTime(ctx, a.lenguage, a.localitation, a.category)
	if err != nil {
		return err
	}

	return nil
}

func getBalanceDaily(trendingSearch []*gogtrends.TrendingSearch) *Analizis {
	output := &Analizis{Economic: 0, Cryptocurrency: 0}

	var balance Analizis

	for _, trends := range trendingSearch {
		balance = balanceSearchArticles(trends.Articles)

		output.Cryptocurrency += balance.Cryptocurrency
		output.Economic += balance.Economic
	}

	return output
}

func getBalanceRealtime(realtime []*gogtrends.TrendingStory) *Analizis {
	output := &Analizis{Economic: 0, Cryptocurrency: 0}

	var balance Analizis

	for _, trends := range realtime {
		balance = balanceTrendingArticles(trends.Articles)

		output.Cryptocurrency += balance.Cryptocurrency
		output.Economic += balance.Economic
	}

	return output
}

func balanceTrendingArticles(articles []*gogtrends.TrendingArticle) Analizis {

	var err error
	report := &scraping.Report{}

	analizisBalaceCryptocurrency := 0
	analizisBalaceEconomic := 0

	for _, articles := range articles {
		report, err = scraping.NewReport(articles.URL)
		if err != nil {
			continue
		}

		analizisBalaceCryptocurrency += KeywordWeightsCryptocurrency(report.Body)
		analizisBalaceEconomic += PositiveEconomyKeywords(report.Body)
	}

	return Analizis{Economic: analizisBalaceEconomic, Cryptocurrency: analizisBalaceCryptocurrency}
}

func balanceSearchArticles(articles []*gogtrends.SearchArticle) Analizis {

	analizisBalaceCryptocurrency := 0
	analizisBalaceEconomic := 0

	var err error
	report := &scraping.Report{}

	for _, articles := range articles {
		report, err = scraping.NewReport(articles.URL)
		if err != nil {
			continue
		}

		analizisBalaceCryptocurrency += KeywordWeightsCryptocurrency(report.Body)
		analizisBalaceEconomic += PositiveEconomyKeywords(report.Body)
	}

	return Analizis{Economic: analizisBalaceEconomic, Cryptocurrency: analizisBalaceCryptocurrency}
}

func KeywordWeightsCryptocurrency(title string) int {
	if title == "" {
		return 0
	}

	positive := 0
	positivePoints := 0

	for key, value := range keywords.KeywordWeightsCryptocurrency {
		keyword := strings.ToUpper(key)
		if strings.Contains(strings.ToUpper(title), keyword) {
			positive++
			positivePoints += value
		}
	}

	negative := 0
	negativePoints := 0
	for key, value := range keywords.NegativeKeywordsCryptocurrency {
		keyword := strings.ToUpper(key)
		if strings.Contains(strings.ToUpper(title), keyword) {
			negative++
			negativePoints -= value
		}
	}

	rangePositive := positive / int(KeywordRange)
	if rangePositive < 1 {
		positivePoints = 0
	}

	rangeNegative := negative / int(KeywordRange)
	if rangeNegative < 1 {
		negativePoints = 0
	}

	return positivePoints + negativePoints
}

func PositiveEconomyKeywords(title string) int {
	if title == "" {
		return 0
	}

	positive := 0
	positivePoints := 0

	for key, value := range keywords.PositiveEconomyKeywords {
		keyword := strings.ToUpper(key)
		if strings.Contains(strings.ToUpper(title), keyword) {
			positive++
			positivePoints += value
		}
	}

	negative := 0
	negativePoints := 0
	for key, value := range keywords.NegativeEconomyKeywords {
		keyword := strings.ToUpper(key)
		if strings.Contains(strings.ToUpper(title), keyword) {
			negative++
			negativePoints -= value
		}
	}

	rangePositive := float64(positive) / KeywordRange
	if rangePositive < 1 {
		positivePoints = 0
	}

	rangeNegative := float64(negative) / KeywordRange
	if rangeNegative < 1 {
		negativePoints = 0
	}

	return positivePoints + negativePoints
}

func (a *AnalizisTrend) GetBalanceTrendsRealTime(ctx context.Context, lenguage, localitation, category string) (*Analizis, error) {
	gogtrends.Debug(false)
	realtime, err := gogtrends.Realtime(ctx, lenguage, localitation, category)
	if err != nil {
		return &Analizis{Economic: 0, Cryptocurrency: 0}, err
	}

	balance := getBalanceRealtime(realtime)
	a.RealtimeArticleBalance = balance

	return balance, nil
}

func (a *AnalizisTrend) GetBalanceDaily(ctx context.Context, lenguage, localitation string) (*Analizis, error) {
	gogtrends.Debug(false)

	daily, err := gogtrends.Daily(ctx, lenguage, localitation)
	if err != nil {
		return nil, err
	}

	balance := getBalanceDaily(daily)
	a.DailyArticleBalance = balance

	dateLastRefresh = time.Now().UTC()
	return balance, nil
}
