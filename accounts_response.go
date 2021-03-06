package robinhood

import (
	"time"
)

type Account struct {
	Deactivated    bool      `json:"deactivated"`
	UpdatedAt      time.Time `json:"updated_at"`
	MarginBalances struct {
		DayTradeBuyingPower               string      `json:"day_trade_buying_power"`
		CreatedAt                         time.Time   `json:"created_at"`
		OvernightBuyingPowerHeldForOrders string      `json:"overnight_buying_power_held_for_orders"`
		CashHeldForOrders                 string      `json:"cash_held_for_orders"`
		DayTradeBuyingPowerHeldForOrders  string      `json:"day_trade_buying_power_held_for_orders"`
		MarkedPatternDayTraderDate        interface{} `json:"marked_pattern_day_trader_date"`
		Cash                              string      `json:"cash"`
		UnallocatedMarginCash             string      `json:"unallocated_margin_cash"`
		UpdatedAt                         time.Time   `json:"updated_at"`
		CashAvailableForWithdrawal        string      `json:"cash_available_for_withdrawal"`
		MarginLimit                       string      `json:"margin_limit"`
		OvernightBuyingPower              string      `json:"overnight_buying_power"`
		UnclearedDeposits                 string      `json:"uncleared_deposits"`
		UnsettledFunds                    string      `json:"unsettled_funds"`
		DayTradeRatio                     string      `json:"day_trade_ratio"`
		OvernightRatio                    string      `json:"overnight_ratio"`
	} `json:"margin_balances"`
	Portfolio                  string      `json:"portfolio"`
	CashBalances               interface{} `json:"cash_balances"`
	WithdrawalHalted           bool        `json:"withdrawal_halted"`
	CashAvailableForWithdrawal string      `json:"cash_available_for_withdrawal"`
	Type                       string      `json:"type"`
	Sma                        string      `json:"sma"`
	SweepEnabled               bool        `json:"sweep_enabled"`
	DepositHalted              bool        `json:"deposit_halted"`
	BuyingPower                string      `json:"buying_power"`
	User                       string      `json:"user"`
	MaxAchEarlyAccessAmount    string      `json:"max_ach_early_access_amount"`
	CashHeldForOrders          string      `json:"cash_held_for_orders"`
	OnlyPositionClosingTrades  bool        `json:"only_position_closing_trades"`
	URL                        string      `json:"url"`
	Positions                  string      `json:"positions"`
	CreatedAt                  time.Time   `json:"created_at"`
	Cash                       string      `json:"cash"`
	SmaHeldForOrders           string      `json:"sma_held_for_orders"`
	AccountNumber              string      `json:"account_number"`
	UnclearedDeposits          string      `json:"uncleared_deposits"`
	UnsettledFunds             string      `json:"unsettled_funds"`
}

type AccountsResponse struct {
	Previous interface{} `json:"previous"`
	Results  []Account   `json:"results"`
	Next     interface{} `json:"next"`
}
