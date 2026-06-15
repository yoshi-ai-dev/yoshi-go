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
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// SpendingService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendingService] method instead.
type SpendingService struct {
	options []option.RequestOption
}

// NewSpendingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpendingService(opts ...option.RequestOption) (r SpendingService) {
	r = SpendingService{}
	r.options = opts
	return
}

// Get the spending breakdown by category for a period. Spending counts outflow
// transactions excluding internal transfers.
func (r *SpendingService) Get(ctx context.Context, query SpendingGetParams, opts ...option.RequestOption) (res *SpendingGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "spending"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SpendingGetResponse struct {
	Data SpendingGetResponseData `json:"data" api:"required"`
	Meta SpendingGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SpendingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendingGetResponseData struct {
	Categories  []SpendingGetResponseDataCategory `json:"categories" api:"required"`
	GrandTotal  float64                           `json:"grand_total" api:"required"`
	PeriodEnd   time.Time                         `json:"period_end" api:"required" format:"date"`
	PeriodStart time.Time                         `json:"period_start" api:"required" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		GrandTotal  respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendingGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *SpendingGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendingGetResponseDataCategory struct {
	GroupLabel       string  `json:"group_label" api:"required"`
	Percentage       float64 `json:"percentage" api:"required"`
	TotalAmount      float64 `json:"total_amount" api:"required"`
	TransactionCount float64 `json:"transaction_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupLabel       respjson.Field
		Percentage       respjson.Field
		TotalAmount      respjson.Field
		TransactionCount respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendingGetResponseDataCategory) RawJSON() string { return r.JSON.raw }
func (r *SpendingGetResponseDataCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendingGetResponseMeta struct {
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
func (r SpendingGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SpendingGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendingGetParams struct {
	// Filter by account ID
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Limit to the top N groups
	TopN param.Opt[int64] `query:"top_n,omitzero" json:"-"`
	// Category granularity to group by
	//
	// Any of "tier1", "tier2", "category".
	GroupBy SpendingGetParamsGroupBy `query:"group_by,omitzero" json:"-"`
	// Trailing window for the breakdown
	//
	// Any of "1w", "1m", "3m", "ytd", "all".
	Period SpendingGetParamsPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpendingGetParams]'s query parameters as `url.Values`.
func (r SpendingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Category granularity to group by
type SpendingGetParamsGroupBy string

const (
	SpendingGetParamsGroupByTier1    SpendingGetParamsGroupBy = "tier1"
	SpendingGetParamsGroupByTier2    SpendingGetParamsGroupBy = "tier2"
	SpendingGetParamsGroupByCategory SpendingGetParamsGroupBy = "category"
)

// Trailing window for the breakdown
type SpendingGetParamsPeriod string

const (
	SpendingGetParamsPeriod1w  SpendingGetParamsPeriod = "1w"
	SpendingGetParamsPeriod1m  SpendingGetParamsPeriod = "1m"
	SpendingGetParamsPeriod3m  SpendingGetParamsPeriod = "3m"
	SpendingGetParamsPeriodYtd SpendingGetParamsPeriod = "ytd"
	SpendingGetParamsPeriodAll SpendingGetParamsPeriod = "all"
)
