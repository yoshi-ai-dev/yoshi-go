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

// AccountService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	options       []option.RequestOption
	BalanceSeries AccountBalanceSeriesService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.options = opts
	r.BalanceSeries = NewAccountBalanceSeriesService(opts...)
	return
}

// List linked financial accounts with display metadata and public lifecycle
// status.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AccountListResponse struct {
	Data AccountListResponseData `json:"data" api:"required"`
	Meta AccountListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseData struct {
	Accounts []AccountListResponseDataAccount `json:"accounts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseDataAccount struct {
	ID                     string  `json:"id" api:"required"`
	ApyPercentDisplay      string  `json:"apy_percent_display" api:"required"`
	AsOf                   string  `json:"as_of" api:"required"`
	BalanceAvailable       float64 `json:"balance_available" api:"required"`
	BalanceCurrent         float64 `json:"balance_current" api:"required"`
	BalanceLimit           float64 `json:"balance_limit" api:"required"`
	ConnectionStatus       string  `json:"connection_status" api:"required"`
	ExternalSource         string  `json:"external_source" api:"required"`
	Hidden                 bool    `json:"hidden" api:"required"`
	InstitutionLogo        string  `json:"institution_logo" api:"required"`
	InstitutionName        string  `json:"institution_name" api:"required"`
	InterestRatePercentage float64 `json:"interest_rate_percentage" api:"required"`
	Mask                   string  `json:"mask" api:"required"`
	Name                   string  `json:"name" api:"required"`
	NextPaymentDueDate     string  `json:"next_payment_due_date" api:"required"`
	// Any of "open", "closed".
	Status    string `json:"status" api:"required"`
	Subtype   string `json:"subtype" api:"required"`
	Type      string `json:"type" api:"required"`
	UpdatedAt string `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		ApyPercentDisplay      respjson.Field
		AsOf                   respjson.Field
		BalanceAvailable       respjson.Field
		BalanceCurrent         respjson.Field
		BalanceLimit           respjson.Field
		ConnectionStatus       respjson.Field
		ExternalSource         respjson.Field
		Hidden                 respjson.Field
		InstitutionLogo        respjson.Field
		InstitutionName        respjson.Field
		InterestRatePercentage respjson.Field
		Mask                   respjson.Field
		Name                   respjson.Field
		NextPaymentDueDate     respjson.Field
		Status                 respjson.Field
		Subtype                respjson.Field
		Type                   respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListResponseDataAccount) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponseDataAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseMeta struct {
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
func (r AccountListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListParams struct {
	// Any of "true", "false".
	Hidden AccountListParamsHidden `query:"hidden,omitzero" json:"-"`
	// Any of "open", "closed", "all".
	Status AccountListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountListParamsHidden string

const (
	AccountListParamsHiddenTrue  AccountListParamsHidden = "true"
	AccountListParamsHiddenFalse AccountListParamsHidden = "false"
)

type AccountListParamsStatus string

const (
	AccountListParamsStatusOpen   AccountListParamsStatus = "open"
	AccountListParamsStatusClosed AccountListParamsStatus = "closed"
	AccountListParamsStatusAll    AccountListParamsStatus = "all"
)
