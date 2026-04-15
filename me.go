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

// MeService contains methods and other services that help with interacting with
// the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeService] method instead.
type MeService struct {
	options []option.RequestOption
}

// NewMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeService(opts ...option.RequestOption) (r MeService) {
	r = MeService{}
	r.options = opts
	return
}

// Get the authenticated user profile (non-sensitive fields only).
func (r *MeService) Get(ctx context.Context, opts ...option.RequestOption) (res *MeGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a comprehensive financial summary including accounts, scores, and goals.
func (r *MeService) Summary(ctx context.Context, query MeSummaryParams, opts ...option.RequestOption) (res *MeSummaryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "me/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MeGetResponse struct {
	Data MeGetResponseData `json:"data" api:"required"`
	Meta MeGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeGetResponseData struct {
	ID            string `json:"id" api:"required"`
	CreatedAt     string `json:"created_at" api:"required"`
	Email         string `json:"email" api:"required"`
	Image         string `json:"image" api:"required"`
	PreferredName string `json:"preferred_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Email         respjson.Field
		Image         respjson.Field
		PreferredName respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeGetResponseMeta struct {
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
func (r MeGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeSummaryResponse struct {
	Data MeSummaryResponseData `json:"data" api:"required"`
	Meta MeSummaryResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *MeSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeSummaryResponseData struct {
	Accounts MeSummaryResponseDataAccounts `json:"accounts" api:"required"`
	Goals    []MeSummaryResponseDataGoal   `json:"goals" api:"required"`
	Scores   MeSummaryResponseDataScores   `json:"scores" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		Goals       respjson.Field
		Scores      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeSummaryResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeSummaryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeSummaryResponseDataAccounts struct {
	Items []MeSummaryResponseDataAccountsItem `json:"items" api:"required"`
	Total int64                               `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeSummaryResponseDataAccounts) RawJSON() string { return r.JSON.raw }
func (r *MeSummaryResponseDataAccounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeSummaryResponseDataAccountsItem struct {
	ID               string  `json:"id" api:"required"`
	BalanceAvailable float64 `json:"balance_available" api:"required"`
	BalanceCurrent   float64 `json:"balance_current" api:"required"`
	BalanceLimit     float64 `json:"balance_limit" api:"required"`
	Hidden           bool    `json:"hidden" api:"required"`
	InstitutionName  string  `json:"institution_name" api:"required"`
	Name             string  `json:"name" api:"required"`
	Status           string  `json:"status" api:"required"`
	Subtype          string  `json:"subtype" api:"required"`
	Type             string  `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		BalanceAvailable respjson.Field
		BalanceCurrent   respjson.Field
		BalanceLimit     respjson.Field
		Hidden           respjson.Field
		InstitutionName  respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		Subtype          respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeSummaryResponseDataAccountsItem) RawJSON() string { return r.JSON.raw }
func (r *MeSummaryResponseDataAccountsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeSummaryResponseDataGoal struct {
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
func (r MeSummaryResponseDataGoal) RawJSON() string { return r.JSON.raw }
func (r *MeSummaryResponseDataGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeSummaryResponseDataScores struct {
	Baseline        float64 `json:"baseline" api:"required"`
	Capacity        float64 `json:"capacity" api:"required"`
	ObservationDate string  `json:"observation_date" api:"required"`
	Recovery        float64 `json:"recovery" api:"required"`
	Yoshi           float64 `json:"yoshi" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Baseline        respjson.Field
		Capacity        respjson.Field
		ObservationDate respjson.Field
		Recovery        respjson.Field
		Yoshi           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeSummaryResponseDataScores) RawJSON() string { return r.JSON.raw }
func (r *MeSummaryResponseDataScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeSummaryResponseMeta struct {
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
func (r MeSummaryResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MeSummaryResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeSummaryParams struct {
	// Any of "true", "false".
	Hidden MeSummaryParamsHidden `query:"hidden,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeSummaryParams]'s query parameters as `url.Values`.
func (r MeSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeSummaryParamsHidden string

const (
	MeSummaryParamsHiddenTrue  MeSummaryParamsHidden = "true"
	MeSummaryParamsHiddenFalse MeSummaryParamsHidden = "false"
)
