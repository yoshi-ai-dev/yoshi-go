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

// AccountBalanceSeriesService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBalanceSeriesService] method instead.
type AccountBalanceSeriesService struct {
	options []option.RequestOption
}

// NewAccountBalanceSeriesService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBalanceSeriesService(opts ...option.RequestOption) (r AccountBalanceSeriesService) {
	r = AccountBalanceSeriesService{}
	r.options = opts
	return
}

// Get historical daily balance series for an account.
func (r *AccountBalanceSeriesService) List(ctx context.Context, id string, query AccountBalanceSeriesListParams, opts ...option.RequestOption) (res *AccountBalanceSeriesListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/balances", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AccountBalanceSeriesListResponse struct {
	Data AccountBalanceSeriesListResponseData `json:"data" api:"required"`
	Meta AccountBalanceSeriesListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBalanceSeriesListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountBalanceSeriesListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBalanceSeriesListResponseData struct {
	AccountID       string                                      `json:"account_id" api:"required"`
	AccountName     string                                      `json:"account_name" api:"required"`
	Days            int64                                       `json:"days" api:"required"`
	ISOCurrencyCode string                                      `json:"iso_currency_code" api:"required"`
	Points          []AccountBalanceSeriesListResponseDataPoint `json:"points" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID       respjson.Field
		AccountName     respjson.Field
		Days            respjson.Field
		ISOCurrencyCode respjson.Field
		Points          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBalanceSeriesListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AccountBalanceSeriesListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBalanceSeriesListResponseDataPoint struct {
	Balance float64 `json:"balance" api:"required"`
	Date    string  `json:"date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance     respjson.Field
		Date        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBalanceSeriesListResponseDataPoint) RawJSON() string { return r.JSON.raw }
func (r *AccountBalanceSeriesListResponseDataPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBalanceSeriesListResponseMeta struct {
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
func (r AccountBalanceSeriesListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AccountBalanceSeriesListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBalanceSeriesListParams struct {
	Days param.Opt[int64] `query:"days,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountBalanceSeriesListParams]'s query parameters as
// `url.Values`.
func (r AccountBalanceSeriesListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
