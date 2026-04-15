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

// InvestmentService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestmentService] method instead.
type InvestmentService struct {
	options []option.RequestOption
}

// NewInvestmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestmentService(opts ...option.RequestOption) (r InvestmentService) {
	r = InvestmentService{}
	r.options = opts
	return
}

// Get investment holdings grouped by asset class.
func (r *InvestmentService) List(ctx context.Context, query InvestmentListParams, opts ...option.RequestOption) (res *InvestmentListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "investments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type InvestmentListResponse struct {
	Data InvestmentListResponseData `json:"data" api:"required"`
	Meta InvestmentListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *InvestmentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentListResponseData struct {
	AccountTotal InvestmentListResponseDataAccountTotal `json:"account_total" api:"required"`
	Groups       []InvestmentListResponseDataGroup      `json:"groups" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountTotal respjson.Field
		Groups       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentListResponseData) RawJSON() string { return r.JSON.raw }
func (r *InvestmentListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentListResponseDataAccountTotal struct {
	CostBasisTotal       float64 `json:"cost_basis_total" api:"required"`
	CurrentValue         float64 `json:"current_value" api:"required"`
	DayChange            float64 `json:"day_change" api:"required"`
	DayChangePercent     float64 `json:"day_change_percent" api:"required"`
	TotalGainLoss        float64 `json:"total_gain_loss" api:"required"`
	TotalGainLossPercent float64 `json:"total_gain_loss_percent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CostBasisTotal       respjson.Field
		CurrentValue         respjson.Field
		DayChange            respjson.Field
		DayChangePercent     respjson.Field
		TotalGainLoss        respjson.Field
		TotalGainLossPercent respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentListResponseDataAccountTotal) RawJSON() string { return r.JSON.raw }
func (r *InvestmentListResponseDataAccountTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentListResponseDataGroup struct {
	AssetClass string                                   `json:"asset_class" api:"required"`
	Holdings   []InvestmentListResponseDataGroupHolding `json:"holdings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetClass  respjson.Field
		Holdings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentListResponseDataGroup) RawJSON() string { return r.JSON.raw }
func (r *InvestmentListResponseDataGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentListResponseDataGroupHolding struct {
	AssetClass           string  `json:"asset_class" api:"required"`
	AverageCostBasis     float64 `json:"average_cost_basis" api:"required"`
	Category             string  `json:"category" api:"required"`
	CostBasisTotal       float64 `json:"cost_basis_total" api:"required"`
	CurrentValue         float64 `json:"current_value" api:"required"`
	DayChange            float64 `json:"day_change" api:"required"`
	DayChangePercent     float64 `json:"day_change_percent" api:"required"`
	LastPrice            float64 `json:"last_price" api:"required"`
	LastPriceAt          string  `json:"last_price_at" api:"required"`
	Name                 string  `json:"name" api:"required"`
	PercentOfAccount     float64 `json:"percent_of_account" api:"required"`
	Quantity             float64 `json:"quantity" api:"required"`
	SecurityID           string  `json:"security_id" api:"required"`
	SecurityType         string  `json:"security_type" api:"required"`
	Symbol               string  `json:"symbol" api:"required"`
	TotalGainLoss        float64 `json:"total_gain_loss" api:"required"`
	TotalGainLossPercent float64 `json:"total_gain_loss_percent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetClass           respjson.Field
		AverageCostBasis     respjson.Field
		Category             respjson.Field
		CostBasisTotal       respjson.Field
		CurrentValue         respjson.Field
		DayChange            respjson.Field
		DayChangePercent     respjson.Field
		LastPrice            respjson.Field
		LastPriceAt          respjson.Field
		Name                 respjson.Field
		PercentOfAccount     respjson.Field
		Quantity             respjson.Field
		SecurityID           respjson.Field
		SecurityType         respjson.Field
		Symbol               respjson.Field
		TotalGainLoss        respjson.Field
		TotalGainLossPercent respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentListResponseDataGroupHolding) RawJSON() string { return r.JSON.raw }
func (r *InvestmentListResponseDataGroupHolding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentListResponseMeta struct {
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
func (r InvestmentListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *InvestmentListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentListParams struct {
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Filter by asset class
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Any of "symbol", "last_price", "day_change", "total_gain_loss", "current_value",
	// "quantity", "cost_basis_total".
	SortBy InvestmentListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortDir InvestmentListParamsSortDir `query:"sort_dir,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvestmentListParams]'s query parameters as `url.Values`.
func (r InvestmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentListParamsSortBy string

const (
	InvestmentListParamsSortBySymbol         InvestmentListParamsSortBy = "symbol"
	InvestmentListParamsSortByLastPrice      InvestmentListParamsSortBy = "last_price"
	InvestmentListParamsSortByDayChange      InvestmentListParamsSortBy = "day_change"
	InvestmentListParamsSortByTotalGainLoss  InvestmentListParamsSortBy = "total_gain_loss"
	InvestmentListParamsSortByCurrentValue   InvestmentListParamsSortBy = "current_value"
	InvestmentListParamsSortByQuantity       InvestmentListParamsSortBy = "quantity"
	InvestmentListParamsSortByCostBasisTotal InvestmentListParamsSortBy = "cost_basis_total"
)

type InvestmentListParamsSortDir string

const (
	InvestmentListParamsSortDirAsc  InvestmentListParamsSortDir = "asc"
	InvestmentListParamsSortDirDesc InvestmentListParamsSortDir = "desc"
)
