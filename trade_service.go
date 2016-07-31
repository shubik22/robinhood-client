package robinhood

import (
	"net/http"
	"net/url"
	"strconv"
)

type TradeParams struct {
	AccountUrl    string
	InstrumentUrl string
	Symbol        string
	Quantity      int
	OrderType     string
}

type TradeService service

func (s *TradeService) PlaceTrade(symbol, orderType string, quantity int) (*http.Response, error) {
	c := s.client
	ar, _, err := c.Accounts.ListAccounts()
	if err != nil {
		return nil, err
	}
	accountUrl := ar.Results[0].URL

	i, _, err := c.Instruments.GetInstrumentFromSymbol(symbol)
	if err != nil {
		return nil, err
	}
	instrumentUrl := i.URL
	tp := &TradeParams{
		AccountUrl:    accountUrl,
		InstrumentUrl: instrumentUrl,
		Symbol:        symbol,
		Quantity:      quantity,
		OrderType:     orderType,
	}
	return s.placeTrade(tp)
}

func (s *TradeService) placeTrade(tp *TradeParams) (*http.Response, error) {
	params := url.Values{}

	params.Add("account", tp.AccountUrl)
	params.Add("instrument", tp.InstrumentUrl)
	// params.add("price", bid_price)
	// params.add("stop_price", stop_price)
	params.Add("quantity", strconv.Itoa(tp.Quantity))
	params.Add("side", tp.OrderType)
	params.Add("symbol", tp.Symbol)
	params.Add("time_in_force", "gfd")
	params.Add("trigger", "immediate")
	params.Add("type", "market")

	var r interface{}
	resp, err := s.client.PostForm("orders/", params, &r)

	return resp, err
}
