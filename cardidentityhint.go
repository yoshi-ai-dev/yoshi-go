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

// CardIdentityHintService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardIdentityHintService] method instead.
type CardIdentityHintService struct {
	options []option.RequestOption
}

// NewCardIdentityHintService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCardIdentityHintService(opts ...option.RequestOption) (r CardIdentityHintService) {
	r = CardIdentityHintService{}
	r.options = opts
	return
}

// List source-data hints that help identify card products. Yoshi does not
// normalize these into rewards products.
func (r *CardIdentityHintService) List(ctx context.Context, query CardIdentityHintListParams, opts ...option.RequestOption) (res *CardIdentityHintListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "card-identity-hints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type CardIdentityHintListResponse struct {
	Data CardIdentityHintListResponseData `json:"data" api:"required"`
	Meta CardIdentityHintListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardIdentityHintListResponse) RawJSON() string { return r.JSON.raw }
func (r *CardIdentityHintListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardIdentityHintListResponseData struct {
	CardIdentityHints []CardIdentityHintListResponseDataCardIdentityHint `json:"card_identity_hints" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardIdentityHints respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardIdentityHintListResponseData) RawJSON() string { return r.JSON.raw }
func (r *CardIdentityHintListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardIdentityHintListResponseDataCardIdentityHint struct {
	AccountID        string                                                          `json:"account_id" api:"required"`
	AccountName      string                                                          `json:"account_name" api:"required"`
	ConnectionStatus string                                                          `json:"connection_status" api:"required"`
	InstitutionName  string                                                          `json:"institution_name" api:"required"`
	Mask             string                                                          `json:"mask" api:"required"`
	MethodCardBrand  CardIdentityHintListResponseDataCardIdentityHintMethodCardBrand `json:"method_card_brand" api:"required"`
	PlaidAccount     CardIdentityHintListResponseDataCardIdentityHintPlaidAccount    `json:"plaid_account" api:"required"`
	// Any of "open", "closed".
	Status string `json:"status" api:"required"`
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
	Type      string `json:"type" api:"required"`
	UpdatedAt string `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID        respjson.Field
		AccountName      respjson.Field
		ConnectionStatus respjson.Field
		InstitutionName  respjson.Field
		Mask             respjson.Field
		MethodCardBrand  respjson.Field
		PlaidAccount     respjson.Field
		Status           respjson.Field
		Subtype          respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardIdentityHintListResponseDataCardIdentityHint) RawJSON() string { return r.JSON.raw }
func (r *CardIdentityHintListResponseDataCardIdentityHint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardIdentityHintListResponseDataCardIdentityHintMethodCardBrand struct {
	Brands    []CardIdentityHintListResponseDataCardIdentityHintMethodCardBrandBrand `json:"brands" api:"required"`
	Source    string                                                                 `json:"source" api:"required"`
	Status    string                                                                 `json:"status" api:"required"`
	UpdatedAt string                                                                 `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brands      respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardIdentityHintListResponseDataCardIdentityHintMethodCardBrand) RawJSON() string {
	return r.JSON.raw
}
func (r *CardIdentityHintListResponseDataCardIdentityHintMethodCardBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardIdentityHintListResponseDataCardIdentityHintMethodCardBrandBrand struct {
	ID            string `json:"id" api:"required"`
	CardProductID string `json:"card_product_id" api:"required"`
	Description   string `json:"description" api:"required"`
	Name          string `json:"name" api:"required"`
	Network       string `json:"network" api:"required"`
	NetworkTier   string `json:"network_tier" api:"required"`
	Type          string `json:"type" api:"required"`
	URL           string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CardProductID respjson.Field
		Description   respjson.Field
		Name          respjson.Field
		Network       respjson.Field
		NetworkTier   respjson.Field
		Type          respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardIdentityHintListResponseDataCardIdentityHintMethodCardBrandBrand) RawJSON() string {
	return r.JSON.raw
}
func (r *CardIdentityHintListResponseDataCardIdentityHintMethodCardBrandBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardIdentityHintListResponseDataCardIdentityHintPlaidAccount struct {
	Name         string `json:"name" api:"required"`
	OfficialName string `json:"official_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		OfficialName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardIdentityHintListResponseDataCardIdentityHintPlaidAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *CardIdentityHintListResponseDataCardIdentityHintPlaidAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardIdentityHintListResponseMeta struct {
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
func (r CardIdentityHintListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *CardIdentityHintListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardIdentityHintListParams struct {
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Any of "open", "closed", "all".
	Status CardIdentityHintListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CardIdentityHintListParams]'s query parameters as
// `url.Values`.
func (r CardIdentityHintListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CardIdentityHintListParamsStatus string

const (
	CardIdentityHintListParamsStatusOpen   CardIdentityHintListParamsStatus = "open"
	CardIdentityHintListParamsStatusClosed CardIdentityHintListParamsStatus = "closed"
	CardIdentityHintListParamsStatusAll    CardIdentityHintListParamsStatus = "all"
)
