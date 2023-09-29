package trends

import (
	"context"

	"github.com/groovili/gogtrends"
	"github.com/MauCastillo/alana/shared/google/analizistrend"
)

func GetExplore(ctx context.Context) ([]*gogtrends.TrendingSearch, error) {
	dailySearches, err := gogtrends.Daily(ctx, "EN", "US")
	if err != nil {
		return nil, err
	}
	return dailySearches, nil
}

func GetSearch(ctx context.Context) ([]*gogtrends.KeywordTopic, error) {
	explore, err := gogtrends.Search(ctx,"bitcoin","EN")
	if err != nil {
		return nil, err
	}
	
	return explore, nil
}

func GetExploreInput(ctx context.Context, inputs []*gogtrends.ComparisonItem) (*gogtrends.ExploreResponse, error) {
	compare, err := gogtrends.Explore(ctx,
		&gogtrends.ExploreRequest{
			ComparisonItems: inputs,
			Category:        31,
			Property:        "",
		}, "EN")

	if err != nil {
		return nil, err
	}

	return &compare, nil

}

func GetTrendsCategories() map[string]string {
	cats := gogtrends.TrendsCategories()
	return cats
}

func GetBalanceTrendsRealTime(ctx context.Context, lenguage, localitation, category string) (int, error) {
	gogtrends.Debug(false)
	realtime, err := gogtrends.Realtime(ctx, lenguage, localitation, category)
	if err != nil {
		return 0, err
	}

	analizis := analizistrend.NewAnalizisTrend()
	balance := analizis.GetBalanceRealtime(realtime)

	return balance, nil
}


func GetBalanceDaily(ctx context.Context, lenguage, localitation string) (int, error) {
	gogtrends.Debug(false)

	daily, err := gogtrends.Daily(ctx, lenguage, localitation)
	if err != nil {
		return 0, err
	}

	analizis := analizistrend.NewAnalizisTrend()
	balance := analizis.GetBalanceDaily(daily)




	return int(balance), nil
}
