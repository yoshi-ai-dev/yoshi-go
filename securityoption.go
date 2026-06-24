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

// SecurityOptionService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityOptionService] method instead.
type SecurityOptionService struct {
	options []option.RequestOption
}

// NewSecurityOptionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityOptionService(opts ...option.RequestOption) (r SecurityOptionService) {
	r = SecurityOptionService{}
	r.options = opts
	return
}

// Get a bounded delayed options chain slice for an equity symbol.
func (r *SecurityOptionService) Chain(ctx context.Context, symbol string, query SecurityOptionChainParams, opts ...option.RequestOption) (res *SecurityOptionChainResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if symbol == "" {
		err = errors.New("missing required symbol parameter")
		return nil, err
	}
	path := fmt.Sprintf("securities/%s/options/chain", url.PathEscape(symbol))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SecurityOptionChainResponse struct {
	Data SecurityOptionChainResponseData `json:"data" api:"required"`
	Meta SecurityOptionChainResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityOptionChainResponse) RawJSON() string { return r.JSON.raw }
func (r *SecurityOptionChainResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityOptionChainResponseData struct {
	Contracts   []SecurityOptionChainResponseDataContract   `json:"contracts" api:"required"`
	Expirations []SecurityOptionChainResponseDataExpiration `json:"expirations" api:"required"`
	Filters     SecurityOptionChainResponseDataFilters      `json:"filters" api:"required"`
	// Any of "quodd", "unsupported".
	Source string `json:"source" api:"required"`
	// Any of "available", "unavailable", "error".
	Status     string                                    `json:"status" api:"required"`
	Summary    SecurityOptionChainResponseDataSummary    `json:"summary" api:"required"`
	Underlying SecurityOptionChainResponseDataUnderlying `json:"underlying" api:"required"`
	Message    string                                    `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contracts   respjson.Field
		Expirations respjson.Field
		Filters     respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		Summary     respjson.Field
		Underlying  respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityOptionChainResponseData) RawJSON() string { return r.JSON.raw }
func (r *SecurityOptionChainResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityOptionChainResponseDataContract struct {
	Ask            float64 `json:"ask" api:"required"`
	Bid            float64 `json:"bid" api:"required"`
	ContractSymbol string  `json:"contract_symbol" api:"required"`
	Currency       string  `json:"currency" api:"required"`
	Delta          float64 `json:"delta" api:"required"`
	ExpirationDate string  `json:"expiration_date" api:"required"`
	ExtrinsicValue float64 `json:"extrinsic_value" api:"required"`
	Gamma          float64 `json:"gamma" api:"required"`
	InTheMoney     bool    `json:"in_the_money" api:"required"`
	IntrinsicValue float64 `json:"intrinsic_value" api:"required"`
	Last           float64 `json:"last" api:"required"`
	Mark           float64 `json:"mark" api:"required"`
	// Any of "itm", "atm", "otm", "unknown".
	Moneyness    string  `json:"moneyness" api:"required"`
	OpenInterest float64 `json:"open_interest" api:"required"`
	// Any of "call", "put", "unknown".
	OptionType    string  `json:"option_type" api:"required"`
	PercentChange float64 `json:"percent_change" api:"required"`
	Rho           float64 `json:"rho" api:"required"`
	Strike        float64 `json:"strike" api:"required"`
	Theta         float64 `json:"theta" api:"required"`
	Vega          float64 `json:"vega" api:"required"`
	Volume        float64 `json:"volume" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ask            respjson.Field
		Bid            respjson.Field
		ContractSymbol respjson.Field
		Currency       respjson.Field
		Delta          respjson.Field
		ExpirationDate respjson.Field
		ExtrinsicValue respjson.Field
		Gamma          respjson.Field
		InTheMoney     respjson.Field
		IntrinsicValue respjson.Field
		Last           respjson.Field
		Mark           respjson.Field
		Moneyness      respjson.Field
		OpenInterest   respjson.Field
		OptionType     respjson.Field
		PercentChange  respjson.Field
		Rho            respjson.Field
		Strike         respjson.Field
		Theta          respjson.Field
		Vega           respjson.Field
		Volume         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityOptionChainResponseDataContract) RawJSON() string { return r.JSON.raw }
func (r *SecurityOptionChainResponseDataContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityOptionChainResponseDataExpiration struct {
	Calls            int64  `json:"calls" api:"required"`
	DaysToExpiration int64  `json:"days_to_expiration" api:"required"`
	ExpirationDate   string `json:"expiration_date" api:"required"`
	Puts             int64  `json:"puts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Calls            respjson.Field
		DaysToExpiration respjson.Field
		ExpirationDate   respjson.Field
		Puts             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityOptionChainResponseDataExpiration) RawJSON() string { return r.JSON.raw }
func (r *SecurityOptionChainResponseDataExpiration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityOptionChainResponseDataFilters struct {
	ExpirationDate string `json:"expiration_date" api:"required"`
	IncludeGreeks  bool   `json:"include_greeks" api:"required"`
	MaxContracts   int64  `json:"max_contracts" api:"required"`
	// Any of "all", "itm", "atm", "otm".
	Moneyness string `json:"moneyness" api:"required"`
	Month     int64  `json:"month" api:"required"`
	// Any of "both", "calls", "puts".
	OptionType          string  `json:"option_type" api:"required"`
	StrikeWindowPercent float64 `json:"strike_window_percent" api:"required"`
	Symbol              string  `json:"symbol" api:"required"`
	Year                int64   `json:"year" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpirationDate      respjson.Field
		IncludeGreeks       respjson.Field
		MaxContracts        respjson.Field
		Moneyness           respjson.Field
		Month               respjson.Field
		OptionType          respjson.Field
		StrikeWindowPercent respjson.Field
		Symbol              respjson.Field
		Year                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityOptionChainResponseDataFilters) RawJSON() string { return r.JSON.raw }
func (r *SecurityOptionChainResponseDataFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityOptionChainResponseDataSummary struct {
	ReturnedContracts       int64 `json:"returned_contracts" api:"required"`
	TotalAvailableContracts int64 `json:"total_available_contracts" api:"required"`
	Truncated               bool  `json:"truncated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReturnedContracts       respjson.Field
		TotalAvailableContracts respjson.Field
		Truncated               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityOptionChainResponseDataSummary) RawJSON() string { return r.JSON.raw }
func (r *SecurityOptionChainResponseDataSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityOptionChainResponseDataUnderlying struct {
	AsOf      string  `json:"as_of" api:"required"`
	Currency  string  `json:"currency" api:"required"`
	LastPrice float64 `json:"last_price" api:"required"`
	Name      string  `json:"name" api:"required"`
	Symbol    string  `json:"symbol" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsOf        respjson.Field
		Currency    respjson.Field
		LastPrice   respjson.Field
		Name        respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityOptionChainResponseDataUnderlying) RawJSON() string { return r.JSON.raw }
func (r *SecurityOptionChainResponseDataUnderlying) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityOptionChainResponseMeta struct {
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
func (r SecurityOptionChainResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SecurityOptionChainResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityOptionChainParams struct {
	StrikeWindowPercent param.Opt[float64] `query:"strike_window_percent,omitzero" json:"-"`
	// Exact option expiration date as YYYY-MM-DD.
	ExpirationDate param.Opt[string] `query:"expiration_date,omitzero" json:"-"`
	IncludeGreeks  param.Opt[bool]   `query:"include_greeks,omitzero" json:"-"`
	// Maximum returned contracts. Defaults to 20, capped at 50.
	MaxContracts param.Opt[int64] `query:"max_contracts,omitzero" json:"-"`
	// Expiration month.
	Month param.Opt[int64] `query:"month,omitzero" json:"-"`
	// Expiration year.
	Year param.Opt[int64] `query:"year,omitzero" json:"-"`
	// Any of "all", "itm", "atm", "otm".
	Moneyness SecurityOptionChainParamsMoneyness `query:"moneyness,omitzero" json:"-"`
	// Any of "both", "calls", "puts".
	OptionType SecurityOptionChainParamsOptionType `query:"option_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecurityOptionChainParams]'s query parameters as
// `url.Values`.
func (r SecurityOptionChainParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SecurityOptionChainParamsMoneyness string

const (
	SecurityOptionChainParamsMoneynessAll SecurityOptionChainParamsMoneyness = "all"
	SecurityOptionChainParamsMoneynessItm SecurityOptionChainParamsMoneyness = "itm"
	SecurityOptionChainParamsMoneynessAtm SecurityOptionChainParamsMoneyness = "atm"
	SecurityOptionChainParamsMoneynessOtm SecurityOptionChainParamsMoneyness = "otm"
)

type SecurityOptionChainParamsOptionType string

const (
	SecurityOptionChainParamsOptionTypeBoth  SecurityOptionChainParamsOptionType = "both"
	SecurityOptionChainParamsOptionTypeCalls SecurityOptionChainParamsOptionType = "calls"
	SecurityOptionChainParamsOptionTypePuts  SecurityOptionChainParamsOptionType = "puts"
)
