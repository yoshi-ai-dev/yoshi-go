// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/apiquery"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// SecurityPriceHistoryService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityPriceHistoryService] method instead.
type SecurityPriceHistoryService struct {
	options []option.RequestOption
}

// NewSecurityPriceHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecurityPriceHistoryService(opts ...option.RequestOption) (r SecurityPriceHistoryService) {
	r = SecurityPriceHistoryService{}
	r.options = opts
	return
}

// List daily OHLCV/NAV price history for a supported security symbol.
func (r *SecurityPriceHistoryService) List(ctx context.Context, symbol string, query SecurityPriceHistoryListParams, opts ...option.RequestOption) (res *SecurityPriceHistoryListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if symbol == "" {
		err = errors.New("missing required symbol parameter")
		return nil, err
	}
	path := fmt.Sprintf("securities/%s/price-history", url.PathEscape(symbol))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SecurityPriceHistoryListResponse struct {
	Data       SecurityPriceHistoryListResponseData       `json:"data" api:"required"`
	Meta       SecurityPriceHistoryListResponseMeta       `json:"meta" api:"required"`
	Pagination SecurityPriceHistoryListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityPriceHistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *SecurityPriceHistoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityPriceHistoryListResponseData struct {
	Prices   []SecurityPriceHistoryListResponseDataPrice  `json:"prices" api:"required"`
	Security SecurityPriceHistoryListResponseDataSecurity `json:"security" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prices      respjson.Field
		Security    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityPriceHistoryListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SecurityPriceHistoryListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityPriceHistoryListResponseDataPrice struct {
	AdjustedClose float64 `json:"adjusted_close" api:"required"`
	ClosePrice    float64 `json:"close_price" api:"required"`
	Date          string  `json:"date" api:"required"`
	HighPrice     float64 `json:"high_price" api:"required"`
	LowPrice      float64 `json:"low_price" api:"required"`
	OpenPrice     float64 `json:"open_price" api:"required"`
	Volume        float64 `json:"volume" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdjustedClose respjson.Field
		ClosePrice    respjson.Field
		Date          respjson.Field
		HighPrice     respjson.Field
		LowPrice      respjson.Field
		OpenPrice     respjson.Field
		Volume        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityPriceHistoryListResponseDataPrice) RawJSON() string { return r.JSON.raw }
func (r *SecurityPriceHistoryListResponseDataPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityPriceHistoryListResponseDataSecurity struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "us_equity", "international_equity", "fixed_income", "cash_equivalent",
	// "alternative", "real_estate", "commodity".
	AssetClass string `json:"asset_class" api:"required"`
	Cusip      string `json:"cusip" api:"required"`
	// Any of "quodd", "plaid", "manual".
	DataProvider string  `json:"data_provider" api:"required"`
	ExpenseRatio float64 `json:"expense_ratio" api:"required"`
	Isin         string  `json:"isin" api:"required"`
	Name         string  `json:"name" api:"required"`
	Sedol        string  `json:"sedol" api:"required"`
	Symbol       string  `json:"symbol" api:"required"`
	// Any of "cash", "cryptocurrency", "derivative", "equity", "etf", "fixed income",
	// "loan", "mutual fund", "other".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AssetClass   respjson.Field
		Cusip        respjson.Field
		DataProvider respjson.Field
		ExpenseRatio respjson.Field
		Isin         respjson.Field
		Name         respjson.Field
		Sedol        respjson.Field
		Symbol       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityPriceHistoryListResponseDataSecurity) RawJSON() string { return r.JSON.raw }
func (r *SecurityPriceHistoryListResponseDataSecurity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityPriceHistoryListResponseMeta struct {
	RequestID string `json:"request_id" api:"required"`
	Timestamp string `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityPriceHistoryListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SecurityPriceHistoryListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityPriceHistoryListResponsePagination struct {
	// Number of items in the current page (equals data.length).
	Count   int64 `json:"count" api:"required"`
	HasMore bool  `json:"has_more" api:"required"`
	// Opaque cursor for the next page; pass back as `cursor`. null when there are no
	// more pages.
	NextCursor string `json:"next_cursor" api:"required"`
	// Page size requested for this page. Present only when the endpoint echoes its
	// limit.
	Limit int64 `json:"limit"`
	// Total items across all pages. Present only when the source can compute it.
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		Limit       respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityPriceHistoryListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SecurityPriceHistoryListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityPriceHistoryListParams struct {
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Inclusive price-date upper bound
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Inclusive price-date lower bound
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecurityPriceHistoryListParams]'s query parameters as
// `url.Values`.
func (r SecurityPriceHistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
