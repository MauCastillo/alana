package trends

import (
	"context"

	// "github.com/MauCastillo/alana/shared/google/trends"
	"github.com/groovili/gogtrends"
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
			Category:        31, // Programming category
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

func GetTrendsRealTime(ctx context.Context, lenguage, localitation, category string) ([]*gogtrends.TrendingStory, error) {
	realtime, err := gogtrends.Realtime(ctx, lenguage, localitation, category)
	if err != nil {
		return nil, err
	}

	return realtime, nil
}
