package goftx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/grishinsana/goftx/models"
)

const (
	apiGetAccountInformation    = "/account"
	apiGetPositions             = "/positions"
	apiGetBalances              = "/wallet/balances"
	apiAccountValueHistory      = "/wallet/usd_value_snapshots?limit=%d"
	apiPostLeverage             = "/account/leverage"
	apiGetReferralRebateHistory = "/referral_rebate_history"
	apiGetWithdrawalHistory     = "/wallet/withdrawals?start_time=%d&end_time=%d"
	apiGetDespositHistory       = "/wallet/deposits?start_time=%d&end_time=%d"
	apiGetLoginStatus           = "/login_status"
	apiGetFundingPayments       = "/funding_payments?start_time=%d&end_time=%d"
)

type Account struct {
	client *Client
}

func (a *Account) GetLoginStatus() ([]byte, error) {
	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetLoginStatus),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := a.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response, nil
}

func (a *Account) GetAccountInformation() (*models.AccountInformation, error) {
	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetAccountInformation),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := a.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result *models.AccountInformation
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (s *Account) GetBalances() ([]*models.Balance, error) {
	request, err := s.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetBalances),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := s.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Balance
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (s *Account) GetAccountValueHistory(limit uint) (*models.AccountValueHistory, error) {
	request, err := s.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, fmt.Sprintf(apiAccountValueHistory, limit)),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := s.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result *models.AccountValueHistory
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (a *Account) GetPositions() ([]*models.Position, error) {
	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetPositions),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := a.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Position
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (a *Account) ChangeAccountLeverage(leverage decimal.Decimal) error {
	body, err := json.Marshal(struct {
		Leverage decimal.Decimal `json:"leverage"`
	}{Leverage: leverage})
	if err != nil {
		return errors.WithStack(err)
	}

	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodPost,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiPostLeverage),
		Body:   body,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = a.client.do(request)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (a *Account) GetReferralRebateHistory() ([]*models.ReferralRebateHistory, error) {
	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetReferralRebateHistory),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := a.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.ReferralRebateHistory
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (a *Account) GetFundingPayments(start, end int64) ([]*models.FundingPayment, error) {
	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, fmt.Sprintf(apiGetFundingPayments, start, end)),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := a.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.FundingPayment
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (a *Account) GetWithdrawalHistory(start, end int64) ([]*models.WithdrawalHistory, error) {
	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, fmt.Sprintf(apiGetWithdrawalHistory, start, end)),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := a.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.WithdrawalHistory
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (a *Account) GetDepositHistory(start, end int64) ([]*models.DepositHistory, error) {
	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, fmt.Sprintf(apiGetDespositHistory, start, end)),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := a.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.DepositHistory
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
