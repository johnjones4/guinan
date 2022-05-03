package ears

import (
	"main/core"

	"github.com/mmcdole/gofeed"
)

func fetchRss(url string) ([]core.Headline, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}
	headlines := make([]core.Headline, len(feed.Items))
	for i, item := range feed.Items {
		headlines[i] = core.Headline{
			Title: item.Title,
			URL:   item.Link,
		}
	}
	return headlines, nil
}

type NYTimesEar struct {
}

func (e *NYTimesEar) FetchAndPopulate(r *core.Record) error {
	headlines, err := fetchRss("https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml")
	if err != nil {
		return err
	}
	r.Info.NYTHeadlines = headlines
	return nil
}
