package models

import (
	"github.com/shopspring/decimal"
)

type PlaceRequestParams struct {
	Volume string `json:"volume"` // quote volume (base currency)
	Price  string `json:"price"`  // quote price (quote currency)
	Pair   string `json:"pair"`   // BTC/USDT, ETH/USDT ...
}

type PlaceReturn struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
	Data struct {
		ID           string          `json:"id"`
		Pair         string          `json:"pair"`
		Email        string          `json:"email"` // User who own the order.
		QuotePrice   decimal.Decimal `json:"price"`
		QuoteVolume  decimal.Decimal `json:"volume"`       // base currency
		QuoteAmount  decimal.Decimal `json:"quote_amount"` // quote currency
		DealVolume   decimal.Decimal `json:"deal_volume"`
		RemainVolume decimal.Decimal `json:"remain_volume"`
		DealAmount   decimal.Decimal `json:"deal_amount"`
		DealFee      decimal.Decimal `json:"deal_fee"`
		CreatedAt    int64           `json:"created_at"`
		UsedUpAt     int64           `json:"used_up_at"`
		IsCanceled   bool            `json:"is_canceled"`
		CanceledAt   int64           `json:"canceled_at"`
		UpdatedAt    int64           `json:"updated_at"`
		Kind         int             `json:"kind"`
	} `json:"data"`
}
