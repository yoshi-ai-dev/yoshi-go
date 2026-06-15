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

func TestGoalNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Goals.New(context.TODO(), yoshi.GoalNewParams{
		GoalType:         yoshi.GoalNewParamsGoalTypeCash,
		Name:             "x",
		TargetAmount:     yoshi.Float(0),
		LinkedAccountIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		Priority:         yoshi.Int(0),
		Status:           yoshi.GoalNewParamsStatusActive,
		TargetDate:       yoshi.String("7321-69-10"),
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGoalUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Goals.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		yoshi.GoalUpdateParams{
			CurrentValue:     yoshi.Float(0),
			GoalType:         yoshi.GoalUpdateParamsGoalTypeCash,
			LinkedAccountIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			Name:             yoshi.String("x"),
			Priority:         yoshi.Int(0),
			Status:           yoshi.GoalUpdateParamsStatusActive,
			TargetAmount:     yoshi.Float(0),
			TargetDate:       yoshi.String("7321-69-10"),
		},
	)
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGoalListWithOptionalParams(t *testing.T) {
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
	_, err := client.Goals.List(context.TODO(), yoshi.GoalListParams{
		Status: yoshi.GoalListParamsStatusActive,
	})
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGoalDelete(t *testing.T) {
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
	_, err := client.Goals.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *yoshi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
