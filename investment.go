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
	"github.com/yoshi-ai-dev/yoshi-go/packages/pagination"
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

// List historical account holding snapshots for investment accounts.
func (r *InvestmentService) HoldingHistory(ctx context.Context, query InvestmentHoldingHistoryParams, opts ...option.RequestOption) (res *pagination.CursorPage[InvestmentHoldingHistoryResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "investments/holdings/history"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List historical account holding snapshots for investment accounts.
func (r *InvestmentService) HoldingHistoryAutoPaging(ctx context.Context, query InvestmentHoldingHistoryParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[InvestmentHoldingHistoryResponse] {
	return pagination.NewCursorPageAutoPager(r.HoldingHistory(ctx, query, opts...))
}

// List current investment holdings per account and security with identifiers and
// fee metadata.
func (r *InvestmentService) Holdings(ctx context.Context, query InvestmentHoldingsParams, opts ...option.RequestOption) (res *pagination.CursorPage[InvestmentHoldingsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "investments/holdings"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List current investment holdings per account and security with identifiers and
// fee metadata.
func (r *InvestmentService) HoldingsAutoPaging(ctx context.Context, query InvestmentHoldingsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[InvestmentHoldingsResponse] {
	return pagination.NewCursorPageAutoPager(r.Holdings(ctx, query, opts...))
}

// Get investment portfolio performance, TWR series, realized gains, unrealized
// gains, and income.
func (r *InvestmentService) Performance(ctx context.Context, query InvestmentPerformanceParams, opts ...option.RequestOption) (res *InvestmentPerformanceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "investments/performance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List normalized current/open tax lots with cost basis, holding period,
// unrealized gain/loss, and coverage metadata.
func (r *InvestmentService) TaxLots(ctx context.Context, query InvestmentTaxLotsParams, opts ...option.RequestOption) (res *pagination.CursorPage[InvestmentTaxLotsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "investments/tax-lots"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List normalized current/open tax lots with cost basis, holding period,
// unrealized gain/loss, and coverage metadata.
func (r *InvestmentService) TaxLotsAutoPaging(ctx context.Context, query InvestmentTaxLotsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[InvestmentTaxLotsResponse] {
	return pagination.NewCursorPageAutoPager(r.TaxLots(ctx, query, opts...))
}

// List investment transactions with security identifiers and explicit fee fields.
func (r *InvestmentService) Transactions(ctx context.Context, query InvestmentTransactionsParams, opts ...option.RequestOption) (res *pagination.CursorPage[InvestmentTransactionsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "investments/transactions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List investment transactions with security identifiers and explicit fee fields.
func (r *InvestmentService) TransactionsAutoPaging(ctx context.Context, query InvestmentTransactionsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[InvestmentTransactionsResponse] {
	return pagination.NewCursorPageAutoPager(r.Transactions(ctx, query, opts...))
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

type InvestmentHoldingHistoryResponse struct {
	AccountID   string `json:"account_id" api:"required" format:"uuid"`
	AccountName string `json:"account_name" api:"required"`
	AsOf        string `json:"as_of" api:"required"`
	// Any of "us_equity", "international_equity", "fixed_income", "cash_equivalent",
	// "alternative", "real_estate", "commodity".
	AssetClass     InvestmentHoldingHistoryResponseAssetClass `json:"asset_class" api:"required"`
	Category       string                                     `json:"category" api:"required"`
	ClosePrice     float64                                    `json:"close_price" api:"required"`
	ClosePriceAsOf string                                     `json:"close_price_as_of" api:"required"`
	CostBasis      float64                                    `json:"cost_basis" api:"required"`
	Cusip          string                                     `json:"cusip" api:"required"`
	// Any of "quodd", "plaid", "manual".
	DataProvider          InvestmentHoldingHistoryResponseDataProvider `json:"data_provider" api:"required"`
	ExpenseRatio          float64                                      `json:"expense_ratio" api:"required"`
	HoldingID             string                                       `json:"holding_id" api:"required" format:"uuid"`
	InstitutionPrice      float64                                      `json:"institution_price" api:"required"`
	InstitutionPriceAsOf  string                                       `json:"institution_price_as_of" api:"required"`
	InstitutionSecurityID string                                       `json:"institution_security_id" api:"required"`
	InstitutionValue      float64                                      `json:"institution_value" api:"required"`
	Isin                  string                                       `json:"isin" api:"required"`
	ISOCurrencyCode       string                                       `json:"iso_currency_code" api:"required"`
	LastPrice             float64                                      `json:"last_price" api:"required"`
	LastPriceAt           string                                       `json:"last_price_at" api:"required"`
	Name                  string                                       `json:"name" api:"required"`
	// Any of "from_source", "reconstructed", "trade".
	Origin     InvestmentHoldingHistoryResponseOrigin `json:"origin" api:"required"`
	Quantity   float64                                `json:"quantity" api:"required"`
	SecurityID string                                 `json:"security_id" api:"required" format:"uuid"`
	// Any of "cash", "cryptocurrency", "derivative", "equity", "etf", "fixed income",
	// "loan", "mutual fund", "other".
	SecurityType           InvestmentHoldingHistoryResponseSecurityType `json:"security_type" api:"required"`
	Sedol                  string                                       `json:"sedol" api:"required"`
	Subtype                string                                       `json:"subtype" api:"required"`
	Symbol                 string                                       `json:"symbol" api:"required"`
	UnofficialCurrencyCode string                                       `json:"unofficial_currency_code" api:"required"`
	VestedQuantity         float64                                      `json:"vested_quantity" api:"required"`
	VestedValue            float64                                      `json:"vested_value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID              respjson.Field
		AccountName            respjson.Field
		AsOf                   respjson.Field
		AssetClass             respjson.Field
		Category               respjson.Field
		ClosePrice             respjson.Field
		ClosePriceAsOf         respjson.Field
		CostBasis              respjson.Field
		Cusip                  respjson.Field
		DataProvider           respjson.Field
		ExpenseRatio           respjson.Field
		HoldingID              respjson.Field
		InstitutionPrice       respjson.Field
		InstitutionPriceAsOf   respjson.Field
		InstitutionSecurityID  respjson.Field
		InstitutionValue       respjson.Field
		Isin                   respjson.Field
		ISOCurrencyCode        respjson.Field
		LastPrice              respjson.Field
		LastPriceAt            respjson.Field
		Name                   respjson.Field
		Origin                 respjson.Field
		Quantity               respjson.Field
		SecurityID             respjson.Field
		SecurityType           respjson.Field
		Sedol                  respjson.Field
		Subtype                respjson.Field
		Symbol                 respjson.Field
		UnofficialCurrencyCode respjson.Field
		VestedQuantity         respjson.Field
		VestedValue            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentHoldingHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *InvestmentHoldingHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentHoldingHistoryResponseAssetClass string

const (
	InvestmentHoldingHistoryResponseAssetClassUsEquity            InvestmentHoldingHistoryResponseAssetClass = "us_equity"
	InvestmentHoldingHistoryResponseAssetClassInternationalEquity InvestmentHoldingHistoryResponseAssetClass = "international_equity"
	InvestmentHoldingHistoryResponseAssetClassFixedIncome         InvestmentHoldingHistoryResponseAssetClass = "fixed_income"
	InvestmentHoldingHistoryResponseAssetClassCashEquivalent      InvestmentHoldingHistoryResponseAssetClass = "cash_equivalent"
	InvestmentHoldingHistoryResponseAssetClassAlternative         InvestmentHoldingHistoryResponseAssetClass = "alternative"
	InvestmentHoldingHistoryResponseAssetClassRealEstate          InvestmentHoldingHistoryResponseAssetClass = "real_estate"
	InvestmentHoldingHistoryResponseAssetClassCommodity           InvestmentHoldingHistoryResponseAssetClass = "commodity"
)

type InvestmentHoldingHistoryResponseDataProvider string

const (
	InvestmentHoldingHistoryResponseDataProviderQuodd  InvestmentHoldingHistoryResponseDataProvider = "quodd"
	InvestmentHoldingHistoryResponseDataProviderPlaid  InvestmentHoldingHistoryResponseDataProvider = "plaid"
	InvestmentHoldingHistoryResponseDataProviderManual InvestmentHoldingHistoryResponseDataProvider = "manual"
)

type InvestmentHoldingHistoryResponseOrigin string

const (
	InvestmentHoldingHistoryResponseOriginFromSource    InvestmentHoldingHistoryResponseOrigin = "from_source"
	InvestmentHoldingHistoryResponseOriginReconstructed InvestmentHoldingHistoryResponseOrigin = "reconstructed"
	InvestmentHoldingHistoryResponseOriginTrade         InvestmentHoldingHistoryResponseOrigin = "trade"
)

type InvestmentHoldingHistoryResponseSecurityType string

const (
	InvestmentHoldingHistoryResponseSecurityTypeCash           InvestmentHoldingHistoryResponseSecurityType = "cash"
	InvestmentHoldingHistoryResponseSecurityTypeCryptocurrency InvestmentHoldingHistoryResponseSecurityType = "cryptocurrency"
	InvestmentHoldingHistoryResponseSecurityTypeDerivative     InvestmentHoldingHistoryResponseSecurityType = "derivative"
	InvestmentHoldingHistoryResponseSecurityTypeEquity         InvestmentHoldingHistoryResponseSecurityType = "equity"
	InvestmentHoldingHistoryResponseSecurityTypeEtf            InvestmentHoldingHistoryResponseSecurityType = "etf"
	InvestmentHoldingHistoryResponseSecurityTypeFixedIncome    InvestmentHoldingHistoryResponseSecurityType = "fixed income"
	InvestmentHoldingHistoryResponseSecurityTypeLoan           InvestmentHoldingHistoryResponseSecurityType = "loan"
	InvestmentHoldingHistoryResponseSecurityTypeMutualFund     InvestmentHoldingHistoryResponseSecurityType = "mutual fund"
	InvestmentHoldingHistoryResponseSecurityTypeOther          InvestmentHoldingHistoryResponseSecurityType = "other"
)

type InvestmentHoldingsResponse struct {
	AccountID   string `json:"account_id" api:"required" format:"uuid"`
	AccountName string `json:"account_name" api:"required"`
	AsOf        string `json:"as_of" api:"required"`
	// Any of "us_equity", "international_equity", "fixed_income", "cash_equivalent",
	// "alternative", "real_estate", "commodity".
	AssetClass     InvestmentHoldingsResponseAssetClass `json:"asset_class" api:"required"`
	Category       string                               `json:"category" api:"required"`
	ClosePrice     float64                              `json:"close_price" api:"required"`
	ClosePriceAsOf string                               `json:"close_price_as_of" api:"required"`
	CostBasis      float64                              `json:"cost_basis" api:"required"`
	Cusip          string                               `json:"cusip" api:"required"`
	// Any of "quodd", "plaid", "manual".
	DataProvider          InvestmentHoldingsResponseDataProvider `json:"data_provider" api:"required"`
	ExpenseRatio          float64                                `json:"expense_ratio" api:"required"`
	HoldingID             string                                 `json:"holding_id" api:"required" format:"uuid"`
	InstitutionPrice      float64                                `json:"institution_price" api:"required"`
	InstitutionPriceAsOf  string                                 `json:"institution_price_as_of" api:"required"`
	InstitutionSecurityID string                                 `json:"institution_security_id" api:"required"`
	InstitutionValue      float64                                `json:"institution_value" api:"required"`
	Isin                  string                                 `json:"isin" api:"required"`
	ISOCurrencyCode       string                                 `json:"iso_currency_code" api:"required"`
	LastPrice             float64                                `json:"last_price" api:"required"`
	LastPriceAt           string                                 `json:"last_price_at" api:"required"`
	Name                  string                                 `json:"name" api:"required"`
	// Any of "from_source", "reconstructed", "trade".
	Origin     InvestmentHoldingsResponseOrigin `json:"origin" api:"required"`
	Quantity   float64                          `json:"quantity" api:"required"`
	SecurityID string                           `json:"security_id" api:"required" format:"uuid"`
	// Any of "cash", "cryptocurrency", "derivative", "equity", "etf", "fixed income",
	// "loan", "mutual fund", "other".
	SecurityType           InvestmentHoldingsResponseSecurityType `json:"security_type" api:"required"`
	Sedol                  string                                 `json:"sedol" api:"required"`
	Subtype                string                                 `json:"subtype" api:"required"`
	Symbol                 string                                 `json:"symbol" api:"required"`
	UnofficialCurrencyCode string                                 `json:"unofficial_currency_code" api:"required"`
	VestedQuantity         float64                                `json:"vested_quantity" api:"required"`
	VestedValue            float64                                `json:"vested_value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID              respjson.Field
		AccountName            respjson.Field
		AsOf                   respjson.Field
		AssetClass             respjson.Field
		Category               respjson.Field
		ClosePrice             respjson.Field
		ClosePriceAsOf         respjson.Field
		CostBasis              respjson.Field
		Cusip                  respjson.Field
		DataProvider           respjson.Field
		ExpenseRatio           respjson.Field
		HoldingID              respjson.Field
		InstitutionPrice       respjson.Field
		InstitutionPriceAsOf   respjson.Field
		InstitutionSecurityID  respjson.Field
		InstitutionValue       respjson.Field
		Isin                   respjson.Field
		ISOCurrencyCode        respjson.Field
		LastPrice              respjson.Field
		LastPriceAt            respjson.Field
		Name                   respjson.Field
		Origin                 respjson.Field
		Quantity               respjson.Field
		SecurityID             respjson.Field
		SecurityType           respjson.Field
		Sedol                  respjson.Field
		Subtype                respjson.Field
		Symbol                 respjson.Field
		UnofficialCurrencyCode respjson.Field
		VestedQuantity         respjson.Field
		VestedValue            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentHoldingsResponse) RawJSON() string { return r.JSON.raw }
func (r *InvestmentHoldingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentHoldingsResponseAssetClass string

const (
	InvestmentHoldingsResponseAssetClassUsEquity            InvestmentHoldingsResponseAssetClass = "us_equity"
	InvestmentHoldingsResponseAssetClassInternationalEquity InvestmentHoldingsResponseAssetClass = "international_equity"
	InvestmentHoldingsResponseAssetClassFixedIncome         InvestmentHoldingsResponseAssetClass = "fixed_income"
	InvestmentHoldingsResponseAssetClassCashEquivalent      InvestmentHoldingsResponseAssetClass = "cash_equivalent"
	InvestmentHoldingsResponseAssetClassAlternative         InvestmentHoldingsResponseAssetClass = "alternative"
	InvestmentHoldingsResponseAssetClassRealEstate          InvestmentHoldingsResponseAssetClass = "real_estate"
	InvestmentHoldingsResponseAssetClassCommodity           InvestmentHoldingsResponseAssetClass = "commodity"
)

type InvestmentHoldingsResponseDataProvider string

const (
	InvestmentHoldingsResponseDataProviderQuodd  InvestmentHoldingsResponseDataProvider = "quodd"
	InvestmentHoldingsResponseDataProviderPlaid  InvestmentHoldingsResponseDataProvider = "plaid"
	InvestmentHoldingsResponseDataProviderManual InvestmentHoldingsResponseDataProvider = "manual"
)

type InvestmentHoldingsResponseOrigin string

const (
	InvestmentHoldingsResponseOriginFromSource    InvestmentHoldingsResponseOrigin = "from_source"
	InvestmentHoldingsResponseOriginReconstructed InvestmentHoldingsResponseOrigin = "reconstructed"
	InvestmentHoldingsResponseOriginTrade         InvestmentHoldingsResponseOrigin = "trade"
)

type InvestmentHoldingsResponseSecurityType string

const (
	InvestmentHoldingsResponseSecurityTypeCash           InvestmentHoldingsResponseSecurityType = "cash"
	InvestmentHoldingsResponseSecurityTypeCryptocurrency InvestmentHoldingsResponseSecurityType = "cryptocurrency"
	InvestmentHoldingsResponseSecurityTypeDerivative     InvestmentHoldingsResponseSecurityType = "derivative"
	InvestmentHoldingsResponseSecurityTypeEquity         InvestmentHoldingsResponseSecurityType = "equity"
	InvestmentHoldingsResponseSecurityTypeEtf            InvestmentHoldingsResponseSecurityType = "etf"
	InvestmentHoldingsResponseSecurityTypeFixedIncome    InvestmentHoldingsResponseSecurityType = "fixed income"
	InvestmentHoldingsResponseSecurityTypeLoan           InvestmentHoldingsResponseSecurityType = "loan"
	InvestmentHoldingsResponseSecurityTypeMutualFund     InvestmentHoldingsResponseSecurityType = "mutual fund"
	InvestmentHoldingsResponseSecurityTypeOther          InvestmentHoldingsResponseSecurityType = "other"
)

type InvestmentPerformanceResponse struct {
	Data InvestmentPerformanceResponseData `json:"data" api:"required"`
	Meta InvestmentPerformanceResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentPerformanceResponse) RawJSON() string { return r.JSON.raw }
func (r *InvestmentPerformanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentPerformanceResponseData struct {
	DataQuality             InvestmentPerformanceResponseDataDataQuality    `json:"data_quality" api:"required"`
	EndDate                 string                                          `json:"end_date" api:"required"`
	EndValue                float64                                         `json:"end_value" api:"required"`
	InvestmentIncomePercent float64                                         `json:"investment_income_percent" api:"required"`
	InvestmentIncomeTotal   float64                                         `json:"investment_income_total" api:"required"`
	NetContributions        float64                                         `json:"net_contributions" api:"required"`
	Period                  string                                          `json:"period" api:"required"`
	RealizedGainsTotal      float64                                         `json:"realized_gains_total" api:"required"`
	ReturnSeries            []InvestmentPerformanceResponseDataReturnSeries `json:"return_series" api:"required"`
	Series                  []InvestmentPerformanceResponseDataSeries       `json:"series" api:"required"`
	StartDate               string                                          `json:"start_date" api:"required"`
	StartValue              float64                                         `json:"start_value" api:"required"`
	TotalReturnDollars      float64                                         `json:"total_return_dollars" api:"required"`
	TotalReturnPercent      float64                                         `json:"total_return_percent" api:"required"`
	UnrealizedGainsTotal    float64                                         `json:"unrealized_gains_total" api:"required"`
	CurrentHoldingsCount    float64                                         `json:"current_holdings_count"`
	// Any of "ok", "empty".
	DataStatus string `json:"data_status"`
	// Any of "no_investment_accounts", "no_daily_balances",
	// "insufficient_balance_coverage".
	EmptyReason          string  `json:"empty_reason"`
	IncludedAccountCount float64 `json:"included_account_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataQuality             respjson.Field
		EndDate                 respjson.Field
		EndValue                respjson.Field
		InvestmentIncomePercent respjson.Field
		InvestmentIncomeTotal   respjson.Field
		NetContributions        respjson.Field
		Period                  respjson.Field
		RealizedGainsTotal      respjson.Field
		ReturnSeries            respjson.Field
		Series                  respjson.Field
		StartDate               respjson.Field
		StartValue              respjson.Field
		TotalReturnDollars      respjson.Field
		TotalReturnPercent      respjson.Field
		UnrealizedGainsTotal    respjson.Field
		CurrentHoldingsCount    respjson.Field
		DataStatus              respjson.Field
		EmptyReason             respjson.Field
		IncludedAccountCount    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentPerformanceResponseData) RawJSON() string { return r.JSON.raw }
func (r *InvestmentPerformanceResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentPerformanceResponseDataDataQuality struct {
	CoverageEndDate      string                                                        `json:"coverage_end_date" api:"required"`
	CoverageStartDate    string                                                        `json:"coverage_start_date" api:"required"`
	ExcludedAccountCount int64                                                         `json:"excluded_account_count" api:"required"`
	ExcludedAccounts     []InvestmentPerformanceResponseDataDataQualityExcludedAccount `json:"excluded_accounts" api:"required"`
	ExcludedValue        float64                                                       `json:"excluded_value" api:"required"`
	Flags                []string                                                      `json:"flags" api:"required"`
	IncludedAccountCount int64                                                         `json:"included_account_count" api:"required"`
	IncludedValue        float64                                                       `json:"included_value" api:"required"`
	IncludedValuePercent float64                                                       `json:"included_value_percent" api:"required"`
	Messages             []string                                                      `json:"messages" api:"required"`
	RequestedEndDate     string                                                        `json:"requested_end_date" api:"required"`
	RequestedStartDate   string                                                        `json:"requested_start_date" api:"required"`
	// Any of "full", "partial", "insufficient".
	Status            string `json:"status" api:"required"`
	TotalAccountCount int64  `json:"total_account_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoverageEndDate      respjson.Field
		CoverageStartDate    respjson.Field
		ExcludedAccountCount respjson.Field
		ExcludedAccounts     respjson.Field
		ExcludedValue        respjson.Field
		Flags                respjson.Field
		IncludedAccountCount respjson.Field
		IncludedValue        respjson.Field
		IncludedValuePercent respjson.Field
		Messages             respjson.Field
		RequestedEndDate     respjson.Field
		RequestedStartDate   respjson.Field
		Status               respjson.Field
		TotalAccountCount    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentPerformanceResponseDataDataQuality) RawJSON() string { return r.JSON.raw }
func (r *InvestmentPerformanceResponseDataDataQuality) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentPerformanceResponseDataDataQualityExcludedAccount struct {
	AccountID         string  `json:"account_id" api:"required"`
	AccountName       string  `json:"account_name" api:"required"`
	BalancePointCount int64   `json:"balance_point_count" api:"required"`
	FirstBalanceDate  string  `json:"first_balance_date" api:"required"`
	LastBalanceDate   string  `json:"last_balance_date" api:"required"`
	LatestBalance     float64 `json:"latest_balance" api:"required"`
	// Any of "no_balance_history", "insufficient_history", "outside_common_window".
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID         respjson.Field
		AccountName       respjson.Field
		BalancePointCount respjson.Field
		FirstBalanceDate  respjson.Field
		LastBalanceDate   respjson.Field
		LatestBalance     respjson.Field
		Reason            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentPerformanceResponseDataDataQualityExcludedAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestmentPerformanceResponseDataDataQualityExcludedAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentPerformanceResponseDataReturnSeries struct {
	Date          string  `json:"date" api:"required"`
	ReturnCents   float64 `json:"return_cents" api:"required"`
	ReturnPercent float64 `json:"return_percent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date          respjson.Field
		ReturnCents   respjson.Field
		ReturnPercent respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentPerformanceResponseDataReturnSeries) RawJSON() string { return r.JSON.raw }
func (r *InvestmentPerformanceResponseDataReturnSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentPerformanceResponseDataSeries struct {
	Date  string  `json:"date" api:"required"`
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentPerformanceResponseDataSeries) RawJSON() string { return r.JSON.raw }
func (r *InvestmentPerformanceResponseDataSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentPerformanceResponseMeta struct {
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
func (r InvestmentPerformanceResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *InvestmentPerformanceResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentTaxLotsResponse struct {
	AccountID        string `json:"account_id" api:"required" format:"uuid"`
	AccountName      string `json:"account_name" api:"required"`
	AcquiredAt       string `json:"acquired_at" api:"required"`
	AsOf             string `json:"as_of" api:"required"`
	CostBasisPerUnit string `json:"cost_basis_per_unit" api:"required"`
	CostBasisTotal   string `json:"cost_basis_total" api:"required"`
	CurrentPrice     string `json:"current_price" api:"required"`
	CurrentPriceAt   string `json:"current_price_at" api:"required"`
	CurrentValue     string `json:"current_value" api:"required"`
	// Any of "missing_acquired_at", "missing_cost_basis", "missing_current_price",
	// "missing_quantity", "unknown_position_type", "unsupported_short_financials".
	DataQuality []string `json:"data_quality" api:"required"`
	// Any of "short_term", "long_term", "unknown".
	HoldingPeriod     InvestmentTaxLotsResponseHoldingPeriod `json:"holding_period" api:"required"`
	HoldingPeriodAsOf string                                 `json:"holding_period_as_of" api:"required"`
	ISOCurrencyCode   string                                 `json:"iso_currency_code" api:"required"`
	LotID             string                                 `json:"lot_id" api:"required"`
	OriginalQuantity  string                                 `json:"original_quantity" api:"required"`
	// Any of "long", "short", "unknown".
	PositionType      InvestmentTaxLotsResponsePositionType `json:"position_type" api:"required"`
	RemainingQuantity string                                `json:"remaining_quantity" api:"required"`
	SecurityID        string                                `json:"security_id" api:"required" format:"uuid"`
	SecurityName      string                                `json:"security_name" api:"required"`
	// Any of "institution", "yoshi_trade".
	Source                    InvestmentTaxLotsResponseSource `json:"source" api:"required"`
	Symbol                    string                          `json:"symbol" api:"required"`
	UnofficialCurrencyCode    string                          `json:"unofficial_currency_code" api:"required"`
	UnrealizedGainLoss        string                          `json:"unrealized_gain_loss" api:"required"`
	UnrealizedGainLossPercent float64                         `json:"unrealized_gain_loss_percent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID                 respjson.Field
		AccountName               respjson.Field
		AcquiredAt                respjson.Field
		AsOf                      respjson.Field
		CostBasisPerUnit          respjson.Field
		CostBasisTotal            respjson.Field
		CurrentPrice              respjson.Field
		CurrentPriceAt            respjson.Field
		CurrentValue              respjson.Field
		DataQuality               respjson.Field
		HoldingPeriod             respjson.Field
		HoldingPeriodAsOf         respjson.Field
		ISOCurrencyCode           respjson.Field
		LotID                     respjson.Field
		OriginalQuantity          respjson.Field
		PositionType              respjson.Field
		RemainingQuantity         respjson.Field
		SecurityID                respjson.Field
		SecurityName              respjson.Field
		Source                    respjson.Field
		Symbol                    respjson.Field
		UnofficialCurrencyCode    respjson.Field
		UnrealizedGainLoss        respjson.Field
		UnrealizedGainLossPercent respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentTaxLotsResponse) RawJSON() string { return r.JSON.raw }
func (r *InvestmentTaxLotsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentTaxLotsResponseHoldingPeriod string

const (
	InvestmentTaxLotsResponseHoldingPeriodShortTerm InvestmentTaxLotsResponseHoldingPeriod = "short_term"
	InvestmentTaxLotsResponseHoldingPeriodLongTerm  InvestmentTaxLotsResponseHoldingPeriod = "long_term"
	InvestmentTaxLotsResponseHoldingPeriodUnknown   InvestmentTaxLotsResponseHoldingPeriod = "unknown"
)

type InvestmentTaxLotsResponsePositionType string

const (
	InvestmentTaxLotsResponsePositionTypeLong    InvestmentTaxLotsResponsePositionType = "long"
	InvestmentTaxLotsResponsePositionTypeShort   InvestmentTaxLotsResponsePositionType = "short"
	InvestmentTaxLotsResponsePositionTypeUnknown InvestmentTaxLotsResponsePositionType = "unknown"
)

type InvestmentTaxLotsResponseSource string

const (
	InvestmentTaxLotsResponseSourceInstitution InvestmentTaxLotsResponseSource = "institution"
	InvestmentTaxLotsResponseSourceYoshiTrade  InvestmentTaxLotsResponseSource = "yoshi_trade"
)

type InvestmentTransactionsResponse struct {
	ID                              string  `json:"id" api:"required" format:"uuid"`
	AccountID                       string  `json:"account_id" api:"required" format:"uuid"`
	AccountName                     string  `json:"account_name" api:"required"`
	Amount                          float64 `json:"amount" api:"required"`
	Cusip                           string  `json:"cusip" api:"required"`
	Date                            string  `json:"date" api:"required"`
	ExternalInvestmentTransactionID string  `json:"external_investment_transaction_id" api:"required"`
	ExternalOrderID                 string  `json:"external_order_id" api:"required"`
	Fees                            float64 `json:"fees" api:"required"`
	IsInternalTransfer              bool    `json:"is_internal_transfer" api:"required"`
	Isin                            string  `json:"isin" api:"required"`
	ISOCurrencyCode                 string  `json:"iso_currency_code" api:"required"`
	Name                            string  `json:"name" api:"required"`
	PostedAt                        string  `json:"posted_at" api:"required"`
	Price                           float64 `json:"price" api:"required"`
	Quantity                        float64 `json:"quantity" api:"required"`
	SecurityID                      string  `json:"security_id" api:"required" format:"uuid"`
	SecurityName                    string  `json:"security_name" api:"required"`
	Sedol                           string  `json:"sedol" api:"required"`
	// Any of "pending", "posted", "cancelled".
	Status InvestmentTransactionsResponseStatus `json:"status" api:"required"`
	// Any of "account fee", "adjustment", "assignment", "buy", "buy to cover",
	// "contribution", "deposit", "distribution", "dividend", "dividend reinvestment",
	// "exercise", "expire", "fund fee", "interest", "interest receivable", "interest
	// reinvestment", "legal fee", "loan payment", "long-term capital gain", "long-term
	// capital gain reinvestment", "management fee", "margin expense", "merger",
	// "miscellaneous fee", "non-qualified dividend", "non-resident tax", "pending
	// credit", "pending debit", "qualified dividend", "rebalance", "return of
	// principal", "request", "sell", "sell short", "send", "short-term capital gain",
	// "short-term capital gain reinvestment", "spin off", "split", "stock
	// distribution", "tax", "tax withheld", "trade", "transfer", "transfer fee",
	// "trust fee", "unqualified gain", "withdrawal".
	Subtype InvestmentTransactionsResponseSubtype `json:"subtype" api:"required"`
	Symbol  string                                `json:"symbol" api:"required"`
	// Any of "buy", "sell", "cancel", "cash", "fee", "transfer".
	Type                   InvestmentTransactionsResponseType `json:"type" api:"required"`
	UnofficialCurrencyCode string                             `json:"unofficial_currency_code" api:"required"`
	UpdatedAt              string                             `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		AccountID                       respjson.Field
		AccountName                     respjson.Field
		Amount                          respjson.Field
		Cusip                           respjson.Field
		Date                            respjson.Field
		ExternalInvestmentTransactionID respjson.Field
		ExternalOrderID                 respjson.Field
		Fees                            respjson.Field
		IsInternalTransfer              respjson.Field
		Isin                            respjson.Field
		ISOCurrencyCode                 respjson.Field
		Name                            respjson.Field
		PostedAt                        respjson.Field
		Price                           respjson.Field
		Quantity                        respjson.Field
		SecurityID                      respjson.Field
		SecurityName                    respjson.Field
		Sedol                           respjson.Field
		Status                          respjson.Field
		Subtype                         respjson.Field
		Symbol                          respjson.Field
		Type                            respjson.Field
		UnofficialCurrencyCode          respjson.Field
		UpdatedAt                       respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestmentTransactionsResponse) RawJSON() string { return r.JSON.raw }
func (r *InvestmentTransactionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestmentTransactionsResponseStatus string

const (
	InvestmentTransactionsResponseStatusPending   InvestmentTransactionsResponseStatus = "pending"
	InvestmentTransactionsResponseStatusPosted    InvestmentTransactionsResponseStatus = "posted"
	InvestmentTransactionsResponseStatusCancelled InvestmentTransactionsResponseStatus = "cancelled"
)

type InvestmentTransactionsResponseSubtype string

const (
	InvestmentTransactionsResponseSubtypeAccountFee                       InvestmentTransactionsResponseSubtype = "account fee"
	InvestmentTransactionsResponseSubtypeAdjustment                       InvestmentTransactionsResponseSubtype = "adjustment"
	InvestmentTransactionsResponseSubtypeAssignment                       InvestmentTransactionsResponseSubtype = "assignment"
	InvestmentTransactionsResponseSubtypeBuy                              InvestmentTransactionsResponseSubtype = "buy"
	InvestmentTransactionsResponseSubtypeBuyToCover                       InvestmentTransactionsResponseSubtype = "buy to cover"
	InvestmentTransactionsResponseSubtypeContribution                     InvestmentTransactionsResponseSubtype = "contribution"
	InvestmentTransactionsResponseSubtypeDeposit                          InvestmentTransactionsResponseSubtype = "deposit"
	InvestmentTransactionsResponseSubtypeDistribution                     InvestmentTransactionsResponseSubtype = "distribution"
	InvestmentTransactionsResponseSubtypeDividend                         InvestmentTransactionsResponseSubtype = "dividend"
	InvestmentTransactionsResponseSubtypeDividendReinvestment             InvestmentTransactionsResponseSubtype = "dividend reinvestment"
	InvestmentTransactionsResponseSubtypeExercise                         InvestmentTransactionsResponseSubtype = "exercise"
	InvestmentTransactionsResponseSubtypeExpire                           InvestmentTransactionsResponseSubtype = "expire"
	InvestmentTransactionsResponseSubtypeFundFee                          InvestmentTransactionsResponseSubtype = "fund fee"
	InvestmentTransactionsResponseSubtypeInterest                         InvestmentTransactionsResponseSubtype = "interest"
	InvestmentTransactionsResponseSubtypeInterestReceivable               InvestmentTransactionsResponseSubtype = "interest receivable"
	InvestmentTransactionsResponseSubtypeInterestReinvestment             InvestmentTransactionsResponseSubtype = "interest reinvestment"
	InvestmentTransactionsResponseSubtypeLegalFee                         InvestmentTransactionsResponseSubtype = "legal fee"
	InvestmentTransactionsResponseSubtypeLoanPayment                      InvestmentTransactionsResponseSubtype = "loan payment"
	InvestmentTransactionsResponseSubtypeLongTermCapitalGain              InvestmentTransactionsResponseSubtype = "long-term capital gain"
	InvestmentTransactionsResponseSubtypeLongTermCapitalGainReinvestment  InvestmentTransactionsResponseSubtype = "long-term capital gain reinvestment"
	InvestmentTransactionsResponseSubtypeManagementFee                    InvestmentTransactionsResponseSubtype = "management fee"
	InvestmentTransactionsResponseSubtypeMarginExpense                    InvestmentTransactionsResponseSubtype = "margin expense"
	InvestmentTransactionsResponseSubtypeMerger                           InvestmentTransactionsResponseSubtype = "merger"
	InvestmentTransactionsResponseSubtypeMiscellaneousFee                 InvestmentTransactionsResponseSubtype = "miscellaneous fee"
	InvestmentTransactionsResponseSubtypeNonQualifiedDividend             InvestmentTransactionsResponseSubtype = "non-qualified dividend"
	InvestmentTransactionsResponseSubtypeNonResidentTax                   InvestmentTransactionsResponseSubtype = "non-resident tax"
	InvestmentTransactionsResponseSubtypePendingCredit                    InvestmentTransactionsResponseSubtype = "pending credit"
	InvestmentTransactionsResponseSubtypePendingDebit                     InvestmentTransactionsResponseSubtype = "pending debit"
	InvestmentTransactionsResponseSubtypeQualifiedDividend                InvestmentTransactionsResponseSubtype = "qualified dividend"
	InvestmentTransactionsResponseSubtypeRebalance                        InvestmentTransactionsResponseSubtype = "rebalance"
	InvestmentTransactionsResponseSubtypeReturnOfPrincipal                InvestmentTransactionsResponseSubtype = "return of principal"
	InvestmentTransactionsResponseSubtypeRequest                          InvestmentTransactionsResponseSubtype = "request"
	InvestmentTransactionsResponseSubtypeSell                             InvestmentTransactionsResponseSubtype = "sell"
	InvestmentTransactionsResponseSubtypeSellShort                        InvestmentTransactionsResponseSubtype = "sell short"
	InvestmentTransactionsResponseSubtypeSend                             InvestmentTransactionsResponseSubtype = "send"
	InvestmentTransactionsResponseSubtypeShortTermCapitalGain             InvestmentTransactionsResponseSubtype = "short-term capital gain"
	InvestmentTransactionsResponseSubtypeShortTermCapitalGainReinvestment InvestmentTransactionsResponseSubtype = "short-term capital gain reinvestment"
	InvestmentTransactionsResponseSubtypeSpinOff                          InvestmentTransactionsResponseSubtype = "spin off"
	InvestmentTransactionsResponseSubtypeSplit                            InvestmentTransactionsResponseSubtype = "split"
	InvestmentTransactionsResponseSubtypeStockDistribution                InvestmentTransactionsResponseSubtype = "stock distribution"
	InvestmentTransactionsResponseSubtypeTax                              InvestmentTransactionsResponseSubtype = "tax"
	InvestmentTransactionsResponseSubtypeTaxWithheld                      InvestmentTransactionsResponseSubtype = "tax withheld"
	InvestmentTransactionsResponseSubtypeTrade                            InvestmentTransactionsResponseSubtype = "trade"
	InvestmentTransactionsResponseSubtypeTransfer                         InvestmentTransactionsResponseSubtype = "transfer"
	InvestmentTransactionsResponseSubtypeTransferFee                      InvestmentTransactionsResponseSubtype = "transfer fee"
	InvestmentTransactionsResponseSubtypeTrustFee                         InvestmentTransactionsResponseSubtype = "trust fee"
	InvestmentTransactionsResponseSubtypeUnqualifiedGain                  InvestmentTransactionsResponseSubtype = "unqualified gain"
	InvestmentTransactionsResponseSubtypeWithdrawal                       InvestmentTransactionsResponseSubtype = "withdrawal"
)

type InvestmentTransactionsResponseType string

const (
	InvestmentTransactionsResponseTypeBuy      InvestmentTransactionsResponseType = "buy"
	InvestmentTransactionsResponseTypeSell     InvestmentTransactionsResponseType = "sell"
	InvestmentTransactionsResponseTypeCancel   InvestmentTransactionsResponseType = "cancel"
	InvestmentTransactionsResponseTypeCash     InvestmentTransactionsResponseType = "cash"
	InvestmentTransactionsResponseTypeFee      InvestmentTransactionsResponseType = "fee"
	InvestmentTransactionsResponseTypeTransfer InvestmentTransactionsResponseType = "transfer"
)

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

type InvestmentHoldingHistoryParams struct {
	// Filter by account ID
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Inclusive as-of date upper bound
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by security ID
	SecurityID param.Opt[string] `query:"security_id,omitzero" format:"uuid" json:"-"`
	// Inclusive as-of date lower bound
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	// Any of "from_source", "reconstructed", "trade".
	Origin InvestmentHoldingHistoryParamsOrigin `query:"origin,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvestmentHoldingHistoryParams]'s query parameters as
// `url.Values`.
func (r InvestmentHoldingHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentHoldingHistoryParamsOrigin string

const (
	InvestmentHoldingHistoryParamsOriginFromSource    InvestmentHoldingHistoryParamsOrigin = "from_source"
	InvestmentHoldingHistoryParamsOriginReconstructed InvestmentHoldingHistoryParamsOrigin = "reconstructed"
	InvestmentHoldingHistoryParamsOriginTrade         InvestmentHoldingHistoryParamsOrigin = "trade"
)

type InvestmentHoldingsParams struct {
	// Filter by account ID
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvestmentHoldingsParams]'s query parameters as
// `url.Values`.
func (r InvestmentHoldingsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentPerformanceParams struct {
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Any of "1w", "1m", "3m", "ytd", "1y", "all".
	Period InvestmentPerformanceParamsPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvestmentPerformanceParams]'s query parameters as
// `url.Values`.
func (r InvestmentPerformanceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentPerformanceParamsPeriod string

const (
	InvestmentPerformanceParamsPeriod1w  InvestmentPerformanceParamsPeriod = "1w"
	InvestmentPerformanceParamsPeriod1m  InvestmentPerformanceParamsPeriod = "1m"
	InvestmentPerformanceParamsPeriod3m  InvestmentPerformanceParamsPeriod = "3m"
	InvestmentPerformanceParamsPeriodYtd InvestmentPerformanceParamsPeriod = "ytd"
	InvestmentPerformanceParamsPeriod1y  InvestmentPerformanceParamsPeriod = "1y"
	InvestmentPerformanceParamsPeriodAll InvestmentPerformanceParamsPeriod = "all"
)

type InvestmentTaxLotsParams struct {
	// Filter by account ID
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by security ID
	SecurityID param.Opt[string] `query:"security_id,omitzero" format:"uuid" json:"-"`
	// Filter by ticker symbol
	Symbol param.Opt[string] `query:"symbol,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvestmentTaxLotsParams]'s query parameters as
// `url.Values`.
func (r InvestmentTaxLotsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentTransactionsParams struct {
	// Filter by account ID
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Inclusive transaction date upper bound
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// Return only explicit fee transactions or rows with fees
	FeesOnly param.Opt[bool] `query:"fees_only,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by security ID
	SecurityID param.Opt[string] `query:"security_id,omitzero" format:"uuid" json:"-"`
	// Inclusive transaction date lower bound
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	// Any of "account fee", "adjustment", "assignment", "buy", "buy to cover",
	// "contribution", "deposit", "distribution", "dividend", "dividend reinvestment",
	// "exercise", "expire", "fund fee", "interest", "interest receivable", "interest
	// reinvestment", "legal fee", "loan payment", "long-term capital gain", "long-term
	// capital gain reinvestment", "management fee", "margin expense", "merger",
	// "miscellaneous fee", "non-qualified dividend", "non-resident tax", "pending
	// credit", "pending debit", "qualified dividend", "rebalance", "return of
	// principal", "request", "sell", "sell short", "send", "short-term capital gain",
	// "short-term capital gain reinvestment", "spin off", "split", "stock
	// distribution", "tax", "tax withheld", "trade", "transfer", "transfer fee",
	// "trust fee", "unqualified gain", "withdrawal".
	Subtype InvestmentTransactionsParamsSubtype `query:"subtype,omitzero" json:"-"`
	// Any of "buy", "sell", "cancel", "cash", "fee", "transfer".
	Type InvestmentTransactionsParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvestmentTransactionsParams]'s query parameters as
// `url.Values`.
func (r InvestmentTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentTransactionsParamsSubtype string

const (
	InvestmentTransactionsParamsSubtypeAccountFee                       InvestmentTransactionsParamsSubtype = "account fee"
	InvestmentTransactionsParamsSubtypeAdjustment                       InvestmentTransactionsParamsSubtype = "adjustment"
	InvestmentTransactionsParamsSubtypeAssignment                       InvestmentTransactionsParamsSubtype = "assignment"
	InvestmentTransactionsParamsSubtypeBuy                              InvestmentTransactionsParamsSubtype = "buy"
	InvestmentTransactionsParamsSubtypeBuyToCover                       InvestmentTransactionsParamsSubtype = "buy to cover"
	InvestmentTransactionsParamsSubtypeContribution                     InvestmentTransactionsParamsSubtype = "contribution"
	InvestmentTransactionsParamsSubtypeDeposit                          InvestmentTransactionsParamsSubtype = "deposit"
	InvestmentTransactionsParamsSubtypeDistribution                     InvestmentTransactionsParamsSubtype = "distribution"
	InvestmentTransactionsParamsSubtypeDividend                         InvestmentTransactionsParamsSubtype = "dividend"
	InvestmentTransactionsParamsSubtypeDividendReinvestment             InvestmentTransactionsParamsSubtype = "dividend reinvestment"
	InvestmentTransactionsParamsSubtypeExercise                         InvestmentTransactionsParamsSubtype = "exercise"
	InvestmentTransactionsParamsSubtypeExpire                           InvestmentTransactionsParamsSubtype = "expire"
	InvestmentTransactionsParamsSubtypeFundFee                          InvestmentTransactionsParamsSubtype = "fund fee"
	InvestmentTransactionsParamsSubtypeInterest                         InvestmentTransactionsParamsSubtype = "interest"
	InvestmentTransactionsParamsSubtypeInterestReceivable               InvestmentTransactionsParamsSubtype = "interest receivable"
	InvestmentTransactionsParamsSubtypeInterestReinvestment             InvestmentTransactionsParamsSubtype = "interest reinvestment"
	InvestmentTransactionsParamsSubtypeLegalFee                         InvestmentTransactionsParamsSubtype = "legal fee"
	InvestmentTransactionsParamsSubtypeLoanPayment                      InvestmentTransactionsParamsSubtype = "loan payment"
	InvestmentTransactionsParamsSubtypeLongTermCapitalGain              InvestmentTransactionsParamsSubtype = "long-term capital gain"
	InvestmentTransactionsParamsSubtypeLongTermCapitalGainReinvestment  InvestmentTransactionsParamsSubtype = "long-term capital gain reinvestment"
	InvestmentTransactionsParamsSubtypeManagementFee                    InvestmentTransactionsParamsSubtype = "management fee"
	InvestmentTransactionsParamsSubtypeMarginExpense                    InvestmentTransactionsParamsSubtype = "margin expense"
	InvestmentTransactionsParamsSubtypeMerger                           InvestmentTransactionsParamsSubtype = "merger"
	InvestmentTransactionsParamsSubtypeMiscellaneousFee                 InvestmentTransactionsParamsSubtype = "miscellaneous fee"
	InvestmentTransactionsParamsSubtypeNonQualifiedDividend             InvestmentTransactionsParamsSubtype = "non-qualified dividend"
	InvestmentTransactionsParamsSubtypeNonResidentTax                   InvestmentTransactionsParamsSubtype = "non-resident tax"
	InvestmentTransactionsParamsSubtypePendingCredit                    InvestmentTransactionsParamsSubtype = "pending credit"
	InvestmentTransactionsParamsSubtypePendingDebit                     InvestmentTransactionsParamsSubtype = "pending debit"
	InvestmentTransactionsParamsSubtypeQualifiedDividend                InvestmentTransactionsParamsSubtype = "qualified dividend"
	InvestmentTransactionsParamsSubtypeRebalance                        InvestmentTransactionsParamsSubtype = "rebalance"
	InvestmentTransactionsParamsSubtypeReturnOfPrincipal                InvestmentTransactionsParamsSubtype = "return of principal"
	InvestmentTransactionsParamsSubtypeRequest                          InvestmentTransactionsParamsSubtype = "request"
	InvestmentTransactionsParamsSubtypeSell                             InvestmentTransactionsParamsSubtype = "sell"
	InvestmentTransactionsParamsSubtypeSellShort                        InvestmentTransactionsParamsSubtype = "sell short"
	InvestmentTransactionsParamsSubtypeSend                             InvestmentTransactionsParamsSubtype = "send"
	InvestmentTransactionsParamsSubtypeShortTermCapitalGain             InvestmentTransactionsParamsSubtype = "short-term capital gain"
	InvestmentTransactionsParamsSubtypeShortTermCapitalGainReinvestment InvestmentTransactionsParamsSubtype = "short-term capital gain reinvestment"
	InvestmentTransactionsParamsSubtypeSpinOff                          InvestmentTransactionsParamsSubtype = "spin off"
	InvestmentTransactionsParamsSubtypeSplit                            InvestmentTransactionsParamsSubtype = "split"
	InvestmentTransactionsParamsSubtypeStockDistribution                InvestmentTransactionsParamsSubtype = "stock distribution"
	InvestmentTransactionsParamsSubtypeTax                              InvestmentTransactionsParamsSubtype = "tax"
	InvestmentTransactionsParamsSubtypeTaxWithheld                      InvestmentTransactionsParamsSubtype = "tax withheld"
	InvestmentTransactionsParamsSubtypeTrade                            InvestmentTransactionsParamsSubtype = "trade"
	InvestmentTransactionsParamsSubtypeTransfer                         InvestmentTransactionsParamsSubtype = "transfer"
	InvestmentTransactionsParamsSubtypeTransferFee                      InvestmentTransactionsParamsSubtype = "transfer fee"
	InvestmentTransactionsParamsSubtypeTrustFee                         InvestmentTransactionsParamsSubtype = "trust fee"
	InvestmentTransactionsParamsSubtypeUnqualifiedGain                  InvestmentTransactionsParamsSubtype = "unqualified gain"
	InvestmentTransactionsParamsSubtypeWithdrawal                       InvestmentTransactionsParamsSubtype = "withdrawal"
)

type InvestmentTransactionsParamsType string

const (
	InvestmentTransactionsParamsTypeBuy      InvestmentTransactionsParamsType = "buy"
	InvestmentTransactionsParamsTypeSell     InvestmentTransactionsParamsType = "sell"
	InvestmentTransactionsParamsTypeCancel   InvestmentTransactionsParamsType = "cancel"
	InvestmentTransactionsParamsTypeCash     InvestmentTransactionsParamsType = "cash"
	InvestmentTransactionsParamsTypeFee      InvestmentTransactionsParamsType = "fee"
	InvestmentTransactionsParamsTypeTransfer InvestmentTransactionsParamsType = "transfer"
)
