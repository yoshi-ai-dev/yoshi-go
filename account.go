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

// Create a real estate property account directly from an address. If
// selected_value is omitted, Yoshi estimates the property value with web search
// before creating the account.
func (r *AccountService) NewRealEstate(ctx context.Context, body AccountNewRealEstateParams, opts ...option.RequestOption) (res *AccountNewRealEstateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "accounts/real-estate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type AccountNewRealEstateResponse struct {
	Data AccountNewRealEstateResponseData `json:"data" api:"required"`
	Meta AccountNewRealEstateResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountNewRealEstateResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountNewRealEstateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewRealEstateResponseData struct {
	Account         AccountNewRealEstateResponseDataAccount         `json:"account" api:"required"`
	RealEstateAsset AccountNewRealEstateResponseDataRealEstateAsset `json:"real_estate_asset" api:"required"`
	Valuation       AccountNewRealEstateResponseDataValuation       `json:"valuation" api:"required"`
	WasCreated      bool                                            `json:"was_created" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account         respjson.Field
		RealEstateAsset respjson.Field
		Valuation       respjson.Field
		WasCreated      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountNewRealEstateResponseData) RawJSON() string { return r.JSON.raw }
func (r *AccountNewRealEstateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewRealEstateResponseDataAccount struct {
	ID             string    `json:"id" api:"required" format:"uuid"`
	AsOf           time.Time `json:"as_of" api:"required" format:"date-time"`
	BalanceCurrent float64   `json:"balance_current" api:"required"`
	// Any of "plaid", "method", "column", "apex", "paper", "custom".
	ExternalSource string `json:"external_source" api:"required"`
	Name           string `json:"name" api:"required"`
	// Any of "checking", "cash management", "cd", "ebt", "hsa", "money market",
	// "paypal", "prepaid", "savings", "529", "401a", "401k", "403b", "457b",
	// "brokerage", "crypto exchange", "education savings account", "fixed annuity",
	// "health reimbursement arrangement", "ira", "keogh", "life insurance", "mutual
	// fund", "non-custodial wallet", "non-taxable brokerage account", "other", "other
	// annuity", "other insurance", "pension", "profit sharing plan", "qshr",
	// "retirement", "roth", "roth 401k", "rrsp", "sarsep", "sep ira", "simple ira",
	// "stock plan", "thrift savings plan", "trust", "ugma", "utma", "variable
	// annuity", "auto", "business", "commercial", "construction", "consumer", "home
	// equity", "line of credit", "loan", "mortgage", "overdraft", "student", "credit
	// card", "real estate".
	Subtype string `json:"subtype" api:"required"`
	// Any of "investment", "credit", "depository", "loan", "brokerage", "real_estate",
	// "other".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AsOf           respjson.Field
		BalanceCurrent respjson.Field
		ExternalSource respjson.Field
		Name           respjson.Field
		Subtype        respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountNewRealEstateResponseDataAccount) RawJSON() string { return r.JSON.raw }
func (r *AccountNewRealEstateResponseDataAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewRealEstateResponseDataRealEstateAsset struct {
	ID             string  `json:"id" api:"required" format:"uuid"`
	AccountID      string  `json:"account_id" api:"required" format:"uuid"`
	AddressDisplay string  `json:"address_display" api:"required"`
	SelectedValue  float64 `json:"selected_value" api:"required"`
	// Any of "estimated_market_value", "last_sale_price", "user_entered".
	ValueSource string `json:"value_source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccountID      respjson.Field
		AddressDisplay respjson.Field
		SelectedValue  respjson.Field
		ValueSource    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountNewRealEstateResponseDataRealEstateAsset) RawJSON() string { return r.JSON.raw }
func (r *AccountNewRealEstateResponseDataRealEstateAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewRealEstateResponseDataValuation struct {
	// Any of "high", "medium", "low".
	Confidence           string    `json:"confidence" api:"required"`
	EstimatedMarketValue float64   `json:"estimated_market_value" api:"required"`
	FetchedAt            time.Time `json:"fetched_at" api:"required" format:"date-time"`
	LastSalePrice        float64   `json:"last_sale_price" api:"required"`
	SourceTitle          string    `json:"source_title" api:"required"`
	SourceURL            string    `json:"source_url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence           respjson.Field
		EstimatedMarketValue respjson.Field
		FetchedAt            respjson.Field
		LastSalePrice        respjson.Field
		SourceTitle          respjson.Field
		SourceURL            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountNewRealEstateResponseDataValuation) RawJSON() string { return r.JSON.raw }
func (r *AccountNewRealEstateResponseDataValuation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewRealEstateResponseMeta struct {
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
func (r AccountNewRealEstateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AccountNewRealEstateResponseMeta) UnmarshalJSON(data []byte) error {
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

type AccountNewRealEstateParams struct {
	AddressLine1   string             `json:"address_line1" api:"required"`
	AccountName    param.Opt[string]  `json:"account_name,omitzero"`
	AddressCity    param.Opt[string]  `json:"address_city,omitzero"`
	AddressState   param.Opt[string]  `json:"address_state,omitzero"`
	AddressZipCode param.Opt[string]  `json:"address_zip_code,omitzero"`
	SelectedValue  param.Opt[float64] `json:"selected_value,omitzero"`
	paramObj
}

func (r AccountNewRealEstateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountNewRealEstateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountNewRealEstateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
