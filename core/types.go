package core

import "time"

type Headline struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Record struct {
	Date time.Time `json:"time"`
	Info struct {
		NYTHeadlines []Headline `json:"nytHeadlines"`
		ETHPrice     float64    `json:"ethPrice"`
		BTCPrice     float64    `json:"btcPrice"`
	} `json:"info"`
}

type Ear interface {
	FetchAndPopulate(r *Record) error
}
