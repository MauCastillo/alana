package scraping

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

var (
	htmlTags = regexp.MustCompile("<.*?>")
)

type Report struct {
	URL  string `json:"url" bson:"url"`
	Body string `json:"body" bson:"body"`
}

func NewReport(url string) (*Report, error) {
	report := Report{URL: url}
	err := report.scraper()
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *Report) scraper() error {
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		r.Body = clearBodyString(e.Text)
	})

	err := c.Visit("https://www.openpr.com/news/3227327/coffee-substitutes-market-2023-industry-analysis-key")
	if err != nil {
		return err
	}

	return nil
}

func clearBodyString(body string) string {
	if body == "" {
		return ""
	}

	bodyRaw := htmlTags.ReplaceAllString(body, "")
	bodyRaw = strings.ReplaceAll(bodyRaw, "\n", "")
	bodyRaw = strings.TrimSpace(bodyRaw)

	return bodyRaw
}
