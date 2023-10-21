package stockapi

type fakeStockApi struct{}

func FakeStockApi() StockApi {
	return &fakeStockApi{}
}

func (s *fakeStockApi) StockGet(code string) (StockResponse, error) {
	if code == "golang" {
		return StockResponse{
			Symbol: "GOLANG",
			Open:   "100",
			High:   "200",
			Low:    "50",
		}, nil
	}

	return StockResponse{}, ErrNoDataFound
}
