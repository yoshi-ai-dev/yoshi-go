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
	"github.com/yoshi-ai-dev/yoshi-go/packages/pagination"
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// PaperTradingAccountTradeService contains methods and other services that help
// with interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperTradingAccountTradeService] method instead.
type PaperTradingAccountTradeService struct {
	options []option.RequestOption
}

// NewPaperTradingAccountTradeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPaperTradingAccountTradeService(opts ...option.RequestOption) (r PaperTradingAccountTradeService) {
	r = PaperTradingAccountTradeService{}
	r.options = opts
	return
}

// Place a buy or sell trade. Requires user approval in the Yoshi web app before
// execution.
func (r *PaperTradingAccountTradeService) New(ctx context.Context, accountID string, body PaperTradingAccountTradeNewParams, opts ...option.RequestOption) (res *PaperTradingAccountTradeNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("paper-trading/accounts/%s/trades", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List trade history for a Test Drive account with cursor-based pagination.
func (r *PaperTradingAccountTradeService) List(ctx context.Context, accountID string, query PaperTradingAccountTradeListParams, opts ...option.RequestOption) (res *pagination.CursorPage[PaperTradingAccountTradeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("paper-trading/accounts/%s/trades", accountID)
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

// List trade history for a Test Drive account with cursor-based pagination.
func (r *PaperTradingAccountTradeService) ListAutoPaging(ctx context.Context, accountID string, query PaperTradingAccountTradeListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[PaperTradingAccountTradeListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, accountID, query, opts...))
}

type PaperTradingAccountTradeNewResponse struct {
	Data PaperTradingAccountTradeNewResponseData `json:"data" api:"required"`
	Meta PaperTradingAccountTradeNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountTradeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseData struct {
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
func (r PaperTradingAccountTradeNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseMeta struct {
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
func (r PaperTradingAccountTradeNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeListResponse struct {
	ID        string  `json:"id" api:"required"`
	Amount    float64 `json:"amount" api:"required"`
	CreatedAt string  `json:"created_at" api:"required"`
	Date      string  `json:"date" api:"required"`
	Name      string  `json:"name" api:"required"`
	Price     float64 `json:"price" api:"required"`
	Quantity  float64 `json:"quantity" api:"required"`
	Subtype   string  `json:"subtype" api:"required"`
	Symbol    string  `json:"symbol" api:"required"`
	Type      string  `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Date        respjson.Field
		Name        respjson.Field
		Price       respjson.Field
		Quantity    respjson.Field
		Subtype     respjson.Field
		Symbol      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountTradeListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewParams struct {
	// Any of "buy", "sell".
	Side     PaperTradingAccountTradeNewParamsSide `json:"side,omitzero" api:"required"`
	Symbol   string                                `json:"symbol" api:"required"`
	Notional param.Opt[float64]                    `json:"notional,omitzero"`
	Quantity param.Opt[float64]                    `json:"quantity,omitzero"`
	// Skip the approval flow and execute immediately. Not supported yet — reserved for
	// future use.
	SkipApproval param.Opt[bool] `json:"skip_approval,omitzero"`
	paramObj
}

func (r PaperTradingAccountTradeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperTradingAccountTradeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperTradingAccountTradeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewParamsSide string

const (
	PaperTradingAccountTradeNewParamsSideBuy  PaperTradingAccountTradeNewParamsSide = "buy"
	PaperTradingAccountTradeNewParamsSideSell PaperTradingAccountTradeNewParamsSide = "sell"
)

type PaperTradingAccountTradeListParams struct {
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperTradingAccountTradeListParams]'s query parameters as
// `url.Values`.
func (r PaperTradingAccountTradeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
