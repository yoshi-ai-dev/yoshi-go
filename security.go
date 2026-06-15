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

// SecurityService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityService] method instead.
type SecurityService struct {
	options []option.RequestOption
}

// NewSecurityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityService(opts ...option.RequestOption) (r SecurityService) {
	r = SecurityService{}
	r.options = opts
	return
}

// Get a basic quote and details for a security by ticker symbol.
func (r *SecurityService) Get(ctx context.Context, symbol string, opts ...option.RequestOption) (res *SecurityGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if symbol == "" {
		err = errors.New("missing required symbol parameter")
		return nil, err
	}
	path := fmt.Sprintf("securities/%s", url.PathEscape(symbol))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fuzzy search for securities by ticker or company name.
func (r *SecurityService) Search(ctx context.Context, query SecuritySearchParams, opts ...option.RequestOption) (res *SecuritySearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "securities/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SecurityGetResponse struct {
	Data SecurityGetResponseData `json:"data" api:"required"`
	Meta SecurityGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SecurityGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGetResponseData struct {
	ID              string  `json:"id" api:"required"`
	AssetClass      string  `json:"asset_class" api:"required"`
	ClosePrice      float64 `json:"close_price" api:"required"`
	ClosePriceAsOf  string  `json:"close_price_as_of" api:"required"`
	DividendYield   float64 `json:"dividend_yield" api:"required"`
	ExpenseRatio    float64 `json:"expense_ratio" api:"required"`
	Industry        string  `json:"industry" api:"required"`
	ISOCurrencyCode string  `json:"iso_currency_code" api:"required"`
	LastPrice       float64 `json:"last_price" api:"required"`
	LastPriceAt     string  `json:"last_price_at" api:"required"`
	MarketCap       float64 `json:"market_cap" api:"required"`
	Name            string  `json:"name" api:"required"`
	PeRatio         float64 `json:"pe_ratio" api:"required"`
	Sector          string  `json:"sector" api:"required"`
	Symbol          string  `json:"symbol" api:"required"`
	Type            string  `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AssetClass      respjson.Field
		ClosePrice      respjson.Field
		ClosePriceAsOf  respjson.Field
		DividendYield   respjson.Field
		ExpenseRatio    respjson.Field
		Industry        respjson.Field
		ISOCurrencyCode respjson.Field
		LastPrice       respjson.Field
		LastPriceAt     respjson.Field
		MarketCap       respjson.Field
		Name            respjson.Field
		PeRatio         respjson.Field
		Sector          respjson.Field
		Symbol          respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *SecurityGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGetResponseMeta struct {
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
func (r SecurityGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SecurityGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecuritySearchResponse struct {
	Data SecuritySearchResponseData `json:"data" api:"required"`
	Meta SecuritySearchResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecuritySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SecuritySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecuritySearchResponseData struct {
	Securities []SecuritySearchResponseDataSecurity `json:"securities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Securities  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecuritySearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *SecuritySearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecuritySearchResponseDataSecurity struct {
	ID          string  `json:"id" api:"required"`
	AssetClass  string  `json:"asset_class" api:"required"`
	ClosePrice  float64 `json:"close_price" api:"required"`
	LastPrice   float64 `json:"last_price" api:"required"`
	LastPriceAt string  `json:"last_price_at" api:"required"`
	Name        string  `json:"name" api:"required"`
	Symbol      string  `json:"symbol" api:"required"`
	Type        string  `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetClass  respjson.Field
		ClosePrice  respjson.Field
		LastPrice   respjson.Field
		LastPriceAt respjson.Field
		Name        respjson.Field
		Symbol      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecuritySearchResponseDataSecurity) RawJSON() string { return r.JSON.raw }
func (r *SecuritySearchResponseDataSecurity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecuritySearchResponseMeta struct {
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
func (r SecuritySearchResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SecuritySearchResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecuritySearchParams struct {
	// Search query (ticker or company name)
	Q string `query:"q" api:"required" json:"-"`
	// Max results (default 20, max 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecuritySearchParams]'s query parameters as `url.Values`.
func (r SecuritySearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
