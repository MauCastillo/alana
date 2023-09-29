package analizistrend

import (
	"testing"

	"github.com/groovili/gogtrends"
	"github.com/stretchr/testify/require"
)

func TestGetBalanceRealtime(t *testing.T) {
	c := require.New(t)

	trendingStory := []*gogtrends.TrendingStory{{
		Articles: []*gogtrends.TrendingArticle{
			{Title: "Mets owner Steve Cohen apologizes to Marlins, who were ..."},
			{Title: "Trump Org. tries to figure out the future of its business after fraud ruling"}},
	}}
	
	analizis := NewAnalizisTrend()

	balance := analizis.GetBalanceRealtime(trendingStory)
	c.Equal(0, balance)
}

func TestGetBalanceDaily(t *testing.T) {
	c := require.New(t)

	trendingSearch := []*gogtrends.TrendingSearch{{

		Articles: []*gogtrends.SearchArticle{
			{Title: "Mets owner Steve Cohen apologizes to Marlins, who were ..."},
			{Title: "Trump Org. tries to figure out the future of its business after fraud ruling"}},
	}}
	
	analizis := NewAnalizisTrend()

	balance := analizis.GetBalanceDaily(trendingSearch)
	c.Equal(0, balance)
}