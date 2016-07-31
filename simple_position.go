package robinhood

import (
	"time"
)

type SimplePosition struct {
	PurchaseTime    time.Time `json:"purchase_time"`
	Quantity        int       `json:"quantity"`
	Symbol          string    `json:"symbol"`
	AverageBuyPrice float64   `json:"average_buy_price"`
	LastTradePrice  float64   `json:"last_trade_price"`
}
