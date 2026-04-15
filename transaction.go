// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/apiquery"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/pagination"
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// TransactionService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.options = opts
	return
}

// List transactions for the authenticated user with cursor-based pagination.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *pagination.CursorPage[TransactionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "transactions"
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

// List transactions for the authenticated user with cursor-based pagination.
func (r *TransactionService) ListAutoPaging(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[TransactionListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type TransactionListResponse struct {
	ID                  string  `json:"id" api:"required"`
	AccountID           string  `json:"account_id" api:"required"`
	AccountName         string  `json:"account_name" api:"required"`
	Amount              float64 `json:"amount" api:"required"`
	AmountAbsolute      float64 `json:"amount_absolute" api:"required"`
	CashFlowDirection   string  `json:"cash_flow_direction" api:"required"`
	CategoryLabel       string  `json:"category_label" api:"required"`
	CategoryTier1       string  `json:"category_tier1" api:"required"`
	CategoryTier2       string  `json:"category_tier2" api:"required"`
	CounterpartyLogoURL string  `json:"counterparty_logo_url" api:"required"`
	CounterpartyName    string  `json:"counterparty_name" api:"required"`
	DateAuthorized      string  `json:"date_authorized" api:"required"`
	DatePosted          string  `json:"date_posted" api:"required"`
	IsInternalTransfer  bool    `json:"is_internal_transfer" api:"required"`
	OriginalDescription string  `json:"original_description" api:"required"`
	Pending             bool    `json:"pending" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountID           respjson.Field
		AccountName         respjson.Field
		Amount              respjson.Field
		AmountAbsolute      respjson.Field
		CashFlowDirection   respjson.Field
		CategoryLabel       respjson.Field
		CategoryTier1       respjson.Field
		CategoryTier2       respjson.Field
		CounterpartyLogoURL respjson.Field
		CounterpartyName    respjson.Field
		DateAuthorized      respjson.Field
		DatePosted          respjson.Field
		IsInternalTransfer  respjson.Field
		OriginalDescription respjson.Field
		Pending             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListParams struct {
	// Filter by account ID
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
