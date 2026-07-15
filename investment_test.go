// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/yoshi-ai-dev/yoshi-go"
	"github.com/yoshi-ai-dev/yoshi-go/internal/testutil"
	"github.com/yoshi-ai-dev/yoshi-go/option"
)

func TestInvestmentListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := yoshi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Investments.List(context.TODO(), yoshi.InvestmentListParams{
		AccountID: yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Category:  yoshi.String("category"),
		SortBy:    yoshi.InvestmentListParamsSortBySymbol,
		SortDir:   yoshi.InvestmentListParamsSortDirAsc,
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentHoldingHistoryWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := yoshi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Investments.HoldingHistory(context.TODO(), yoshi.InvestmentHoldingHistoryParams{
		AccountID:  yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Cursor:     yoshi.String("cursor"),
		EndDate:    yoshi.String("7321-69-10"),
		Limit:      yoshi.Int(1),
		Origin:     yoshi.InvestmentHoldingHistoryParamsOriginFromSource,
		SecurityID: yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		StartDate:  yoshi.String("7321-69-10"),
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentHoldingsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := yoshi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Investments.Holdings(context.TODO(), yoshi.InvestmentHoldingsParams{
		AccountID: yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Cursor:    yoshi.String("cursor"),
		Limit:     yoshi.Int(1),
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentPerformanceWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := yoshi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Investments.Performance(context.TODO(), yoshi.InvestmentPerformanceParams{
		AccountID: yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Period:    yoshi.InvestmentPerformanceParamsPeriod1w,
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentTaxLotsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := yoshi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Investments.TaxLots(context.TODO(), yoshi.InvestmentTaxLotsParams{
		AccountID:  yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Cursor:     yoshi.String("cursor"),
		Limit:      yoshi.Int(1),
		SecurityID: yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Symbol:     yoshi.String("x"),
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentTransactionsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := yoshi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Investments.Transactions(context.TODO(), yoshi.InvestmentTransactionsParams{
		AccountID:  yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Cursor:     yoshi.String("cursor"),
		EndDate:    yoshi.String("7321-69-10"),
		FeesOnly:   yoshi.Bool(true),
		Limit:      yoshi.Int(1),
		SecurityID: yoshi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		StartDate:  yoshi.String("7321-69-10"),
		Subtype:    yoshi.InvestmentTransactionsParamsSubtypeAccountFee,
		Type:       yoshi.InvestmentTransactionsParamsTypeBuy,
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
