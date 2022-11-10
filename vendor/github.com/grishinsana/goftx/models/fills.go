package models

import (
	"github.com/shopspring/decimal"
)

type Fill struct {
	Fee           decimal.Decimal `json:"fee"`
	FeeCurrency   string          `json:"feeCurrency"`
	FeeRate       decimal.Decimal `json:"feeRate"`
	Future        string          `json:"future"`
	ID            int64           `json:"id"`
	Liquidity     string          `json:"liquidity"`
	Market        string          `json:"market"`
	BaseCurrency  string          `json:"baseCurrency"`
	QuoteCurrency string          `json:"quoteCurrency"`
	OrderID       int64           `json:"orderId"`
	TradeID       int64           `json:"tradeId"`
	Price         decimal.Decimal `json:"price"`
	Side          string          `json:"side"`
	Size          decimal.Decimal `json:"size"`
	Time          FTXTime         `json:"time"`
	Type          string          `json:"type"`
}

type FillsParams struct {
	Market    *string `json:"market"`
	StartTime *int    `json:"start_time"`
	EndTime   *int    `json:"end_time"`
	Order     *string `json:"order"`
	OrderID   *int64  `json:"orderId"`
}
