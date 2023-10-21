package stockapi

import "errors"

type StockResponse struct {
	Symbol string `csv:"Symbol"`
	Open   string `csv:"Open"`
	High   string `csv:"High"`
	Low    string `csv:"Low"`
}

type StockApi interface {
	StockGet(code string) (StockResponse, error)
}

var ErrNoDataFound = errors.New("no data found")
