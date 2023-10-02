package scraping

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScraping(t *testing.T) {
	c := require.New(t)

	url := "https://www.openpr.com/news/3227327/coffee-substitutes-market-2023-industry-analysis-key"

	report, err := NewReport("")
	c.EqualError(err, ErrorEmptyURL.Error())
	c.Nil(report)

	report, err = NewReport(url)
	c.NoError(err)

	c.Equal(url, report.URL)
	c.Contains(report.Body, "Coffee Substitutes Market 2023")

}
