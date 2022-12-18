package lunch

import (
	"context"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type WijkandersFetcher struct {
}

func (f *WijkandersFetcher) FetchByDate(ctx context.Context, date time.Time, lang string) (*LunchFetchResult, error) {
	return nil, nil
}

func (f *WijkandersFetcher) FetchByWeek(ctx context.Context, date time.Time, lang string) ([]LunchFetchResult, error) {
	var returnDays []LunchFetchResult
	c := colly.NewCollector(
		colly.AllowedDomains("https://wijkanders.se"),
	)
	c.OnHTML(".post-content .entry-content", func(e *colly.HTMLElement) {
		e.ForEach("p", func(i int, p *colly.HTMLElement) {
			if 13 >= i && i >= 9 {
				returnDays = append(returnDays)
			}
		})
	})

	c.Visit("https://wijkanders.se/restaurangen/")

	return returnDays, nil
}
