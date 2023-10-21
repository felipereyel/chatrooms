package stockapi

import (
	"chatrooms/gosrc/utils"
	"errors"
	"fmt"
)

type stockApi struct {
	client utils.HTTPClient
}

func NewStockApi() *stockApi {
	client := utils.HTTPClient{
		BaseUrl: "https://stooq.com",
	}

	return &stockApi{client}
}

func (s *stockApi) StockGet(code string) (StockResponse, error) {
	resp, err := s.client.Request("GET", "/q/l/?s="+code+"&f=sd2t2ohlcv&h&e=csv", nil)
	if err != nil {
		fmt.Println("failed to perform request: " + err.Error())
		return StockResponse{}, err
	}

	var responses []StockResponse
	if err := utils.CSVParse(resp.Body, &responses); err != nil {
		return StockResponse{}, err
	}

	if len(responses) == 0 {
		return StockResponse{}, errors.New("no data found")
	}

	return responses[0], nil
}
