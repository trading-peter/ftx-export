package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AccountInformation struct {
	BackstopProvider             bool            `json:"backstopProvider"`
	Collateral                   decimal.Decimal `json:"collateral"`
	FreeCollateral               decimal.Decimal `json:"freeCollateral"`
	InitialMarginRequirement     decimal.Decimal `json:"initialMarginRequirement"`
	Liquidating                  bool            `json:"liquidating"`
	MaintenanceMarginRequirement decimal.Decimal `json:"maintenanceMarginRequirement"`
	MakerFee                     decimal.Decimal `json:"makerFee"`
	MarginFraction               decimal.Decimal `json:"marginFraction"`
	OpenMarginFraction           decimal.Decimal `json:"openMarginFraction"`
	TakerFee                     decimal.Decimal `json:"takerFee"`
	TotalAccountValue            decimal.Decimal `json:"totalAccountValue"`
	TotalPositionSize            decimal.Decimal `json:"totalPositionSize"`
	Username                     string          `json:"username"`
	Leverage                     decimal.Decimal `json:"leverage"`
	Positions                    []Position      `json:"positions"`
}

type Position struct {
	Cost                         decimal.Decimal `json:"cost"`
	EntryPrice                   decimal.Decimal `json:"entryPrice"`
	EstimatedLiquidationPrice    decimal.Decimal `json:"estimatedLiquidationPrice"`
	Future                       string          `json:"future"`
	InitialMarginRequirement     decimal.Decimal `json:"initialMarginRequirement"`
	LongOrderSize                decimal.Decimal `json:"longOrderSize"`
	MaintenanceMarginRequirement decimal.Decimal `json:"maintenanceMarginRequirement"`
	NetSize                      decimal.Decimal `json:"netSize"`
	OpenSize                     decimal.Decimal `json:"openSize"`
	RecentPnl                    decimal.Decimal `json:"recentPnl"`
	RealizedPnl                  decimal.Decimal `json:"realizedPnl"`
	ShortOrderSize               decimal.Decimal `json:"shortOrderSize"`
	Side                         string          `json:"side"`
	Size                         decimal.Decimal `json:"size"`
	UnrealizedPnl                decimal.Decimal `json:"unrealizedPnl"`
	CollateralUsed               decimal.Decimal `json:"collateralUsed"`
}

type Balance struct {
	Coin                   string          `json:"coin"`
	Free                   decimal.Decimal `json:"free"`
	Total                  decimal.Decimal `json:"total"`
	UsdValue               decimal.Decimal `json:"usdValue"`
	SpotBorrow             decimal.Decimal `json:"spotBorrow"`
	AvailableWithoutBorrow decimal.Decimal `json:"availableWithoutBorrow"`
}

type AccountValueHistory struct {
	Now     time.Time       `json:"now"`
	Value   decimal.Decimal `json:"value"`
	Records []AccountValue  `json:"records"`
}

type AccountValue struct {
	Time     time.Time       `json:"time"`
	UsdValue decimal.Decimal `json:"usdValue"`
}

type ReferralRebateHistory struct {
	Subaccount string          `json:"subaccount"`
	Size       decimal.Decimal `json:"size"`
	Day        time.Time       `json:"day"`
}

type FundingPayment struct {
	Future  string          `json:"future"`
	ID      int64           `json:"id"`
	Payment decimal.Decimal `json:"payment"`
	Time    time.Time       `json:"time"`
}

type BorrowHistory struct {
	Coin string          `json:"coin"`
	Cost decimal.Decimal `json:"cost"`
	Rate decimal.Decimal `json:"rate"`
	Size decimal.Decimal `json:"size"`
	Time time.Time       `json:"time"`
}

type LendingHistory struct {
	Coin     string          `json:"coin"`
	Proceeds decimal.Decimal `json:"proceeds"`
	Rate     decimal.Decimal `json:"rate"`
	Size     decimal.Decimal `json:"size"`
	Time     time.Time       `json:"time"`
}

type WithdrawalHistory struct {
	Coin    string          `json:"coin"`
	Address string          `json:"address"`
	Tag     string          `json:"tag"`
	Fee     decimal.Decimal `json:"fee"`
	ID      int64           `json:"id"`
	Size    decimal.Decimal `json:"size"`
	Status  string          `json:"status"`
	Time    time.Time       `json:"time"`
	Method  string          `json:"method"`
	Txid    string          `json:"txid"`
	Notes   string          `json:"notes"`
}

type DepositHistory struct {
	Coin          string          `json:"coin"`
	Confirmations int64           `json:"confirmations"`
	ConfirmedTime time.Time       `json:"confirmedTime"`
	Fee           decimal.Decimal `json:"fee"`
	ID            int64           `json:"id"`
	SentTime      time.Time       `json:"sentTime"`
	Size          decimal.Decimal `json:"size"`
	Status        string          `json:"status"`
	Time          time.Time       `json:"time"`
	Txid          string          `json:"txid"`
	Notes         string          `json:"notes"`
}
