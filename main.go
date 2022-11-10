package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/beefsack/go-rate"
	"github.com/grishinsana/goftx"
	"github.com/grishinsana/goftx/models"
	"github.com/kataras/golog"
	"github.com/ncruces/zenity"
)

var limiter *rate.RateLimiter = rate.New(28, time.Second)

func main() {
	done := make(chan struct{})

	key, err := zenity.Entry("API Key", zenity.Title("Paste your API key"))
	if err != nil {
		os.Exit(1)
	}

	secret, err := zenity.Entry("API Secret", zenity.Title("Paste your API secret"))
	if err != nil {
		os.Exit(1)
	}

	golog.Info("Starting download of account data")

	client := goftx.New(goftx.WithAuth(key, secret))

	accList, err := client.GetSubaccounts()

	if err != nil {
		golog.Error(err)
	}

	run := func(subAcc string) {
		label := subAcc
		if subAcc == "" {
			label = "Main"
		}

		var subClient *goftx.Client

		if subAcc == "" {
			subClient = goftx.New(goftx.WithAuth(key, secret))
		} else {
			subClient = goftx.New(goftx.WithAuth(key, secret), goftx.WithSubaccount(subAcc))
		}

		count, err := fetchTransactions(subClient, fmt.Sprintf("%s_transaction_history.csv", label))

		golog.Infof("Downloaded %d transactions for %s", count, label)

		if err != nil {
			golog.Error(err)
		}

		count, err = fetchWithdrawals(subClient, fmt.Sprintf("%s_withdrawal_history.csv", label))

		golog.Infof("Downloaded %d withdrawals for %s", count, label)

		if err != nil {
			golog.Error(err)
		}

		count, err = fetchDeposits(subClient, fmt.Sprintf("%s_deposit_history.csv", label))

		golog.Infof("Downloaded %d deposits for %s", count, label)

		if err != nil {
			golog.Error(err)
		}

		count, err = fetchRefRebates(subClient, fmt.Sprintf("%s_referral_rebates.csv", label))

		golog.Infof("Downloaded %d referral rebates for %s", count, label)

		if err != nil {
			golog.Error(err)
		}

		count, err = fetchFunding(subClient, fmt.Sprintf("%s_futures_funding.csv", label))

		golog.Infof("Downloaded %d funding records for %s", count, label)

		if err != nil {
			golog.Error(err)
		}

		count, err = fetchBorrowHistory(subClient, fmt.Sprintf("%s_borrow_history.csv", label))

		golog.Infof("Downloaded %d borrow history for %s", count, label)

		if err != nil {
			golog.Error(err)
		}

		count, err = fetchLendingHistory(subClient, fmt.Sprintf("%s_lending_history.csv", label))

		golog.Infof("Downloaded %d lending history for %s", count, label)

		if err != nil {
			golog.Error(err)
		}

		fetchAccountDetails(subClient, fmt.Sprintf("%s_account_details.json", label))
	}

	run("")

	for _, sa := range accList {
		run(sa.Nickname)
	}

	fmt.Println("FINISHED!")
	fmt.Println("Press CTRL+C to close this window")
	<-done
}

func fetchAccountDetails(client *goftx.Client, outFile string) error {
	acc, err := client.GetAccountInformation()

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(acc, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(outFile, data, 0777)
}

func fetchFunding(client *goftx.Client, outFile string) (int64, error) {
	since := time.Unix(0, 0)

	file, err := os.Create(outFile)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{
		"Future",
		"ID",
		"Payment",
		"Time",
	})

	start := since.Unix()
	end := time.Now().Unix()
	var count int64 = 0

	for {
		limiter.Wait()
		recs, err := client.GetFundingPayments(start, end)
		if err != nil {
			return 0, err
		}

		if len(recs) == 0 {
			break
		}

		start = recs[0].Time.Unix()
		end = recs[len(recs)-1].Time.Unix()

		for i := range recs {
			count++
			f := recs[i]

			csvWriter.Write([]string{
				f.Future,
				fmt.Sprintf("%d", f.ID),
				f.Payment.String(),
				f.Time.String(),
			})
		}
	}

	csvWriter.Flush()
	return count, nil
}

func fetchBorrowHistory(client *goftx.Client, outFile string) (int64, error) {
	since := time.Unix(0, 0)

	file, err := os.Create(outFile)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{
		"Coin",
		"Cost",
		"Rate",
		"Size",
		"Time",
	})

	start := since.Unix()
	end := time.Now().Unix()
	var count int64 = 0

	for {
		limiter.Wait()
		recs, err := client.SpotMargin.GetBorrowHistory(start, end)

		if err != nil {
			return 0, err
		}

		if len(recs) == 0 {
			break
		}

		start = recs[0].Time.Unix()
		end = recs[len(recs)-1].Time.Unix()

		for i := range recs {
			count++
			f := recs[i]

			csvWriter.Write([]string{
				f.Coin,
				f.Cost.String(),
				f.Rate.String(),
				f.Size.String(),
				f.Time.String(),
			})
		}
	}

	csvWriter.Flush()
	return count, nil
}

func fetchLendingHistory(client *goftx.Client, outFile string) (int64, error) {
	since := time.Unix(0, 0)

	file, err := os.Create(outFile)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{
		"Coin",
		"Proceeds",
		"Rate",
		"Size",
		"Time",
	})

	start := since.Unix()
	end := time.Now().Unix()
	var count int64 = 0

	for {
		limiter.Wait()
		recs, err := client.SpotMargin.GetLendingHistory(start, end)

		if err != nil {
			return 0, err
		}

		if len(recs) == 0 {
			break
		}

		start = recs[0].Time.Unix()
		end = recs[len(recs)-1].Time.Unix()

		for i := range recs {
			count++
			f := recs[i]

			csvWriter.Write([]string{
				f.Coin,
				f.Proceeds.String(),
				f.Rate.String(),
				f.Size.String(),
				f.Time.String(),
			})
		}
	}

	csvWriter.Flush()
	return count, nil
}

func fetchWithdrawals(client *goftx.Client, outFile string) (int64, error) {
	since := time.Unix(0, 0)

	file, err := os.Create(outFile)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{
		"Coin",
		"Address",
		"Tag",
		"Fee",
		"ID",
		"Size",
		"Status",
		"Time",
		"Method",
		"Txid",
		"Notes",
	})

	start := since.Unix()
	end := time.Now().Unix()
	var count int64 = 0

	for {
		limiter.Wait()
		recs, err := client.GetWithdrawalHistory(start, end)

		if err != nil {
			return 0, err
		}

		if len(recs) == 0 {
			break
		}

		end = recs[len(recs)-1].Time.Unix()

		for i := range recs {
			count++
			f := recs[i]

			csvWriter.Write([]string{
				f.Coin,
				f.Address,
				f.Tag,
				f.Fee.String(),
				fmt.Sprintf("%d", f.ID),
				f.Size.String(),
				f.Status,
				f.Time.String(),
				f.Method,
				f.Txid,
				f.Notes,
			})
		}
	}

	csvWriter.Flush()
	return count, nil
}

func fetchDeposits(client *goftx.Client, outFile string) (int64, error) {
	since := time.Unix(0, 0)

	file, err := os.Create(outFile)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{
		"Coin",
		"Confirmations",
		"ConfirmedTime",
		"Fee",
		"ID",
		"SentTime",
		"Size",
		"Status",
		"Time",
		"Txid",
		"Notes",
	})

	start := since.Unix()
	end := time.Now().Unix()
	var count int64 = 0

	for {
		limiter.Wait()
		recs, err := client.GetDepositHistory(start, end)

		if err != nil {
			return 0, err
		}

		if len(recs) == 0 {
			break
		}

		end = recs[len(recs)-1].Time.Unix()

		for i := range recs {
			count++
			f := recs[i]

			csvWriter.Write([]string{
				f.Coin,
				fmt.Sprintf("%d", f.Confirmations),
				f.ConfirmedTime.String(),
				f.Fee.String(),
				fmt.Sprintf("%d", f.ID),
				f.SentTime.String(),
				f.Size.String(),
				f.Status,
				f.Time.String(),
				f.Txid,
				f.Notes,
			})
		}
	}

	csvWriter.Flush()
	return count, nil
}

func fetchTransactions(client *goftx.Client, outFile string) (int64, error) {
	since := time.Unix(0, 0)

	file, err := os.Create(outFile)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{
		"ID",
		"BaseCurrency",
		"Fee",
		"FeeCurrency",
		"FeeRate",
		"Future",
		"Liquidity",
		"Market",
		"OrderID",
		"Price",
		"QuoteCurrency",
		"Side",
		"Size",
		"Time",
		"TradeID",
		"Type",
	})

	start := int(since.Unix())
	end := int(time.Now().Unix())
	var count int64 = 0

	for {
		limiter.Wait()
		fills, err := client.Fills.Fills(&models.FillsParams{
			StartTime: &start,
			EndTime:   &end,
		})

		if err != nil {
			return 0, err
		}

		if len(fills) == 0 {
			break
		}

		end = int(fills[len(fills)-1].Time.Time.Unix())

		for i := range fills {
			count++
			f := fills[i]

			csvWriter.Write([]string{
				fmt.Sprintf("%d", f.ID),
				f.BaseCurrency,
				f.Fee.String(),
				f.FeeCurrency,
				f.FeeRate.String(),
				f.Future,
				string(f.Liquidity),
				f.Market,
				fmt.Sprintf("%d", f.OrderID),
				f.Price.String(),
				f.QuoteCurrency,
				string(f.Side),
				f.Size.String(),
				f.Time.Time.String(),
				fmt.Sprintf("%d", f.TradeID),
				f.Type,
			})
		}
	}

	csvWriter.Flush()
	return count, nil
}

func fetchRefRebates(client *goftx.Client, outFile string) (int64, error) {
	file, err := os.Create(outFile)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{
		"Subaccount",
		"Size",
		"Day",
	})

	var count int64 = 0

	limiter.Wait()
	recs, err := client.GetReferralRebateHistory()

	if err != nil {
		return 0, err
	}

	for i := range recs {
		count++
		f := recs[i]

		csvWriter.Write([]string{
			f.Subaccount,
			f.Size.String(),
			f.Day.String(),
		})
	}

	csvWriter.Flush()
	return count, nil
}
