package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	re := regexp.MustCompile("<.*?>")
	

	c.OnHTML("body", func(e *colly.HTMLElement) {
		s := re.ReplaceAllString(e.Text, "")
		s = strings.ReplaceAll(s, "\n", "")
		s = strings.TrimSpace(s)
		fmt.Println(":", s)
	})

	c.Visit("https://www.openpr.com/news/3227327/coffee-substitutes-market-2023-industry-analysis-key")
}
