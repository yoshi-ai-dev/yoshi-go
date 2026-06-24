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

func TestBenchmarkReplayWithOptionalParams(t *testing.T) {
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
	_, err := client.Benchmarks.Replay(context.TODO(), yoshi.BenchmarkReplayParams{
		Allocations: []yoshi.BenchmarkReplayParamsAllocation{{
			Ticker: "ticker",
			Weight: 1,
		}},
		CashFlows: []yoshi.BenchmarkReplayParamsCashFlow{{
			Amount: 0,
			Date:   "7321-69-10",
		}},
		ActualEndValue: yoshi.Float(0),
		EndDate:        yoshi.String("7321-69-10"),
		StartDate:      yoshi.String("7321-69-10"),
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
