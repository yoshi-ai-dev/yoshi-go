// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/apiquery"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
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

// Create a new financial goal (cash savings, investment, or debt payoff).
func (r *GoalService) New(ctx context.Context, body GoalNewParams, opts ...option.RequestOption) (res *GoalNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "goals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an existing goal. Only provided fields change; `target_date: null` clears
// the date.
func (r *GoalService) Update(ctx context.Context, id string, body GoalUpdateParams, opts ...option.RequestOption) (res *GoalUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("goals/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List financial goals with optional status filter.
func (r *GoalService) List(ctx context.Context, query GoalListParams, opts ...option.RequestOption) (res *GoalListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "goals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete an existing goal. System-managed goals cannot be deleted.
func (r *GoalService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GoalDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("goals/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type GoalNewResponse struct {
	Data GoalNewResponseData `json:"data" api:"required"`
	Meta GoalNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalNewResponseData struct {
	Goal GoalNewResponseDataGoal `json:"goal" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Goal        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *GoalNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalNewResponseDataGoal struct {
	ID           string    `json:"id" api:"required"`
	CurrentValue float64   `json:"current_value" api:"required"`
	GoalType     string    `json:"goal_type" api:"required"`
	Name         string    `json:"name" api:"required"`
	Status       string    `json:"status" api:"required"`
	TargetAmount float64   `json:"target_amount" api:"required"`
	TargetDate   time.Time `json:"target_date" api:"required" format:"date"`
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
func (r GoalNewResponseDataGoal) RawJSON() string { return r.JSON.raw }
func (r *GoalNewResponseDataGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalNewResponseMeta struct {
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
func (r GoalNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *GoalNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalUpdateResponse struct {
	Data GoalUpdateResponseData `json:"data" api:"required"`
	Meta GoalUpdateResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalUpdateResponseData struct {
	Goal GoalUpdateResponseDataGoal `json:"goal" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Goal        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *GoalUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalUpdateResponseDataGoal struct {
	ID           string    `json:"id" api:"required"`
	CurrentValue float64   `json:"current_value" api:"required"`
	GoalType     string    `json:"goal_type" api:"required"`
	Name         string    `json:"name" api:"required"`
	Status       string    `json:"status" api:"required"`
	TargetAmount float64   `json:"target_amount" api:"required"`
	TargetDate   time.Time `json:"target_date" api:"required" format:"date"`
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
func (r GoalUpdateResponseDataGoal) RawJSON() string { return r.JSON.raw }
func (r *GoalUpdateResponseDataGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalUpdateResponseMeta struct {
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
func (r GoalUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *GoalUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	ID           string    `json:"id" api:"required"`
	CurrentValue float64   `json:"current_value" api:"required"`
	GoalType     string    `json:"goal_type" api:"required"`
	Name         string    `json:"name" api:"required"`
	Status       string    `json:"status" api:"required"`
	TargetAmount float64   `json:"target_amount" api:"required"`
	TargetDate   time.Time `json:"target_date" api:"required" format:"date"`
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

type GoalDeleteResponse struct {
	Data GoalDeleteResponseData `json:"data" api:"required"`
	Meta GoalDeleteResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalDeleteResponseData struct {
	ID      string `json:"id" api:"required"`
	Deleted bool   `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *GoalDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalDeleteResponseMeta struct {
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
func (r GoalDeleteResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *GoalDeleteResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalNewParams struct {
	TargetAmount param.Opt[float64] `json:"target_amount,omitzero" api:"required"`
	// Any of "cash", "investment", "debt_payoff".
	GoalType         GoalNewParamsGoalType `json:"goal_type,omitzero" api:"required"`
	Name             string                `json:"name" api:"required"`
	Priority         param.Opt[int64]      `json:"priority,omitzero"`
	TargetDate       param.Opt[string]     `json:"target_date,omitzero"`
	LinkedAccountIDs []string              `json:"linked_account_ids,omitzero" format:"uuid"`
	// Any of "active", "paused", "completed", "cancelled".
	Status GoalNewParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r GoalNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GoalNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalNewParamsGoalType string

const (
	GoalNewParamsGoalTypeCash       GoalNewParamsGoalType = "cash"
	GoalNewParamsGoalTypeInvestment GoalNewParamsGoalType = "investment"
	GoalNewParamsGoalTypeDebtPayoff GoalNewParamsGoalType = "debt_payoff"
)

type GoalNewParamsStatus string

const (
	GoalNewParamsStatusActive    GoalNewParamsStatus = "active"
	GoalNewParamsStatusPaused    GoalNewParamsStatus = "paused"
	GoalNewParamsStatusCompleted GoalNewParamsStatus = "completed"
	GoalNewParamsStatusCancelled GoalNewParamsStatus = "cancelled"
)

type GoalUpdateParams struct {
	CurrentValue param.Opt[float64] `json:"current_value,omitzero"`
	Priority     param.Opt[int64]   `json:"priority,omitzero"`
	TargetAmount param.Opt[float64] `json:"target_amount,omitzero"`
	TargetDate   param.Opt[string]  `json:"target_date,omitzero"`
	Name         param.Opt[string]  `json:"name,omitzero"`
	// Any of "cash", "investment", "debt_payoff".
	GoalType         GoalUpdateParamsGoalType `json:"goal_type,omitzero"`
	LinkedAccountIDs []string                 `json:"linked_account_ids,omitzero" format:"uuid"`
	// Any of "active", "paused", "completed", "cancelled".
	Status GoalUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r GoalUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GoalUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalUpdateParamsGoalType string

const (
	GoalUpdateParamsGoalTypeCash       GoalUpdateParamsGoalType = "cash"
	GoalUpdateParamsGoalTypeInvestment GoalUpdateParamsGoalType = "investment"
	GoalUpdateParamsGoalTypeDebtPayoff GoalUpdateParamsGoalType = "debt_payoff"
)

type GoalUpdateParamsStatus string

const (
	GoalUpdateParamsStatusActive    GoalUpdateParamsStatus = "active"
	GoalUpdateParamsStatusPaused    GoalUpdateParamsStatus = "paused"
	GoalUpdateParamsStatusCompleted GoalUpdateParamsStatus = "completed"
	GoalUpdateParamsStatusCancelled GoalUpdateParamsStatus = "cancelled"
)

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
