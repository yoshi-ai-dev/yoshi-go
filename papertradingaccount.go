// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// PaperTradingAccountService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperTradingAccountService] method instead.
type PaperTradingAccountService struct {
	options  []option.RequestOption
	Holdings PaperTradingAccountHoldingService
	Trades   PaperTradingAccountTradeService
}

// NewPaperTradingAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaperTradingAccountService(opts ...option.RequestOption) (r PaperTradingAccountService) {
	r = PaperTradingAccountService{}
	r.options = opts
	r.Holdings = NewPaperTradingAccountHoldingService(opts...)
	r.Trades = NewPaperTradingAccountTradeService(opts...)
	return
}

// Create a Test Drive account. Requires user approval in the Yoshi web app before
// the account is created.
func (r *PaperTradingAccountService) New(ctx context.Context, body PaperTradingAccountNewParams, opts ...option.RequestOption) (res *PaperTradingAccountNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "paper-trading/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List Test Drive accounts with current balances for the user.
func (r *PaperTradingAccountService) List(ctx context.Context, opts ...option.RequestOption) (res *PaperTradingAccountListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "paper-trading/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PaperTradingAccountNewResponse struct {
	Data PaperTradingAccountNewResponseData `json:"data" api:"required"`
	Meta PaperTradingAccountNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseData struct {
	ActionID          string `json:"action_id" api:"required" format:"uuid"`
	ApprovalStatusURL string `json:"approval_status_url" api:"required"`
	ApprovalURL       string `json:"approval_url" api:"required"`
	Description       string `json:"description" api:"required"`
	Status            string `json:"status" api:"required"`
	ThreadID          string `json:"thread_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID          respjson.Field
		ApprovalStatusURL respjson.Field
		ApprovalURL       respjson.Field
		Description       respjson.Field
		Status            respjson.Field
		ThreadID          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseMeta struct {
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
func (r PaperTradingAccountNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountListResponse struct {
	Data PaperTradingAccountListResponseData `json:"data" api:"required"`
	Meta PaperTradingAccountListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountListResponseData struct {
	Accounts []PaperTradingAccountListResponseDataAccount `json:"accounts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountListResponseDataAccount struct {
	ID               string  `json:"id" api:"required"`
	AsOf             string  `json:"as_of" api:"required"`
	BalanceAvailable float64 `json:"balance_available" api:"required"`
	BalanceCurrent   float64 `json:"balance_current" api:"required"`
	CreatedAt        string  `json:"created_at" api:"required"`
	Name             string  `json:"name" api:"required"`
	Status           string  `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AsOf             respjson.Field
		BalanceAvailable respjson.Field
		BalanceCurrent   respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountListResponseDataAccount) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountListResponseDataAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountListResponseMeta struct {
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
func (r PaperTradingAccountListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// Skip the approval flow and execute immediately. Not supported yet — reserved for
	// future use.
	SkipApproval        param.Opt[bool]   `json:"skip_approval,omitzero"`
	StartingCashBalance param.Opt[string] `json:"starting_cash_balance,omitzero"`
	paramObj
}

func (r PaperTradingAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperTradingAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperTradingAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
