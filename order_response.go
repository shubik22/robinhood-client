package robinhood

import (
	"time"
)

type OrderResponse struct {
	UpdatedAt          time.Time     `json:"updated_at"`
	RefID              interface{}   `json:"ref_id"`
	TimeInForce        string        `json:"time_in_force"`
	Fees               string        `json:"fees"`
	Cancel             interface{}   `json:"cancel"`
	ID                 string        `json:"id"`
	CumulativeQuantity string        `json:"cumulative_quantity"`
	StopPrice          interface{}   `json:"stop_price"`
	RejectReason       interface{}   `json:"reject_reason"`
	Instrument         string        `json:"instrument"`
	State              string        `json:"state"`
	Trigger            string        `json:"trigger"`
	Type               string        `json:"type"`
	LastTransactionAt  time.Time     `json:"last_transaction_at"`
	Price              string        `json:"price"`
	Executions         []interface{} `json:"executions"`
	ExtendedHours      bool          `json:"extended_hours"`
	Account            string        `json:"account"`
	URL                string        `json:"url"`
	CreatedAt          time.Time     `json:"created_at"`
	Side               string        `json:"side"`
	Position           string        `json:"position"`
	AveragePrice       interface{}   `json:"average_price"`
	Quantity           string        `json:"quantity"`
}
