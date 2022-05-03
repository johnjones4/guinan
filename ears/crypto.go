package ears

import (
	"encoding/json"
	"io"
	"main/core"
	"net/http"
	"net/url"
	"strconv"
)

type alphavantageResponse struct {
	Body struct {
		ExchangeRate string `json:"5. Exchange Rate"`
	} `json:"Realtime Currency Exchange Rate"`
}

type CryptoEar struct {
	APIKey string
}

func (e *CryptoEar) fetch(c string) (float64, error) {
	params := make(url.Values)
	params.Set("function", "CURRENCY_EXCHANGE_RATE")
	params.Set("from_currency", c)
	params.Set("to_currency", "USD")
	params.Set("apikey", e.APIKey)
	u := url.URL{
		Scheme:   "https",
		Host:     "www.alphavantage.co",
		Path:     "/query",
		RawQuery: params.Encode(),
	}
	res, err := http.Get(u.String())
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var resp alphavantageResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseFloat(resp.Body.ExchangeRate, 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func (e *CryptoEar) FetchAndPopulate(r *core.Record) error {
	btc, err := e.fetch("BTC")
	if err != nil {
		return err
	}

	eth, err := e.fetch("ETH")
	if err != nil {
		return err
	}

	r.Info.BTCPrice = btc
	r.Info.ETHPrice = eth

	return nil
}
