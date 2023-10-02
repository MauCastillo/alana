package trends

import (
	"context"
	"testing"

	"github.com/groovili/gogtrends"
	"github.com/stretchr/testify/require"
)

func TestGetExplore(t *testing.T) {
	c := require.New(t)

	ctx := context.Background()

	trends, err := GetExplore(ctx)
	c.NoError(err)
	c.True(len(trends) > 0)
}

func TestGetExploreInput(t *testing.T) {
	c := require.New(t)

	ctx := context.Background()
	comparisonItems := []*gogtrends.ComparisonItem{
		{
			Keyword: "Go",
			Geo:     "US",
			Time:    "today 12-m",
		},
		{
			Keyword: "Python",
			Geo:     "US",
			Time:    "today 12-m",
		},
		{
			Keyword: "PHP",
			Geo:     "US",
			Time:    "today 12-m",
		},
	}

	exploreResponse, err := GetExploreInput(ctx, comparisonItems)
	c.NoError(err)
	c.True(exploreResponse.Len() > 10)
}

func TestGetTrendsCategories(t *testing.T) {
	c := require.New(t)

	category := GetTrendsCategories()
	trendsCategories := map[string]string{
		"all": "all",
		"b":   "business",
		"h":   "main news",
		"m":   "health",
		"t":   "science and technics",
		"e":   "entertainment",
		"s":   "sport",
	}
	c.Equal(category, trendsCategories)
}

func TestGetSearch(t *testing.T) {
	c := require.New(t)

	ctx := context.Background()

	exploreResponse, err := GetSearch(ctx)
	c.NoError(err)
	c.Equal(exploreResponse[0].Title, "Bitcoin")
}
