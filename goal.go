// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/apiquery"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// GoalService contains methods and other services that help with interacting with
// the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGoalService] method instead.
type GoalService struct {
	options []option.RequestOption
}

// NewGoalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGoalService(opts ...option.RequestOption) (r GoalService) {
	r = GoalService{}
	r.options = opts
	return
}

// List financial goals with optional status filter.
func (r *GoalService) List(ctx context.Context, query GoalListParams, opts ...option.RequestOption) (res *GoalListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "goals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GoalListResponse struct {
	Data GoalListResponseData `json:"data" api:"required"`
	Meta GoalListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalListResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalListResponseData struct {
	Goals []GoalListResponseDataGoal `json:"goals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Goals       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalListResponseData) RawJSON() string { return r.JSON.raw }
func (r *GoalListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalListResponseDataGoal struct {
	ID           string  `json:"id" api:"required"`
	CurrentValue float64 `json:"current_value" api:"required"`
	GoalType     string  `json:"goal_type" api:"required"`
	Name         string  `json:"name" api:"required"`
	Status       string  `json:"status" api:"required"`
	TargetAmount float64 `json:"target_amount" api:"required"`
	TargetDate   string  `json:"target_date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CurrentValue respjson.Field
		GoalType     respjson.Field
		Name         respjson.Field
		Status       respjson.Field
		TargetAmount respjson.Field
		TargetDate   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalListResponseDataGoal) RawJSON() string { return r.JSON.raw }
func (r *GoalListResponseDataGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalListResponseMeta struct {
	RequestID string    `json:"request_id" api:"required"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *GoalListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalListParams struct {
	// Any of "active", "paused", "completed", "cancelled".
	Status GoalListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GoalListParams]'s query parameters as `url.Values`.
func (r GoalListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GoalListParamsStatus string

const (
	GoalListParamsStatusActive    GoalListParamsStatus = "active"
	GoalListParamsStatusPaused    GoalListParamsStatus = "paused"
	GoalListParamsStatusCompleted GoalListParamsStatus = "completed"
	GoalListParamsStatusCancelled GoalListParamsStatus = "cancelled"
)
