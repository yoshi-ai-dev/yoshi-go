// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// PaperTradingAccountHoldingService contains methods and other services that help
// with interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperTradingAccountHoldingService] method instead.
type PaperTradingAccountHoldingService struct {
	options []option.RequestOption
}

// NewPaperTradingAccountHoldingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPaperTradingAccountHoldingService(opts ...option.RequestOption) (r PaperTradingAccountHoldingService) {
	r = PaperTradingAccountHoldingService{}
	r.options = opts
	return
}

// Get current holdings/positions for a Test Drive account.
func (r *PaperTradingAccountHoldingService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *PaperTradingAccountHoldingListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("paper-trading/accounts/%s/holdings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PaperTradingAccountHoldingListResponse struct {
	Data PaperTradingAccountHoldingListResponseData `json:"data" api:"required"`
	Meta PaperTradingAccountHoldingListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountHoldingListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountHoldingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountHoldingListResponseData struct {
	Holdings []PaperTradingAccountHoldingListResponseDataHolding `json:"holdings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Holdings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountHoldingListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountHoldingListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountHoldingListResponseDataHolding struct {
	ID               string  `json:"id" api:"required"`
	AsOf             string  `json:"as_of" api:"required"`
	CostBasis        float64 `json:"cost_basis" api:"required"`
	InstitutionPrice float64 `json:"institution_price" api:"required"`
	InstitutionValue float64 `json:"institution_value" api:"required"`
	Name             string  `json:"name" api:"required"`
	Quantity         float64 `json:"quantity" api:"required"`
	SecurityID       string  `json:"security_id" api:"required"`
	SecurityType     string  `json:"security_type" api:"required"`
	Symbol           string  `json:"symbol" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AsOf             respjson.Field
		CostBasis        respjson.Field
		InstitutionPrice respjson.Field
		InstitutionValue respjson.Field
		Name             respjson.Field
		Quantity         respjson.Field
		SecurityID       respjson.Field
		SecurityType     respjson.Field
		Symbol           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountHoldingListResponseDataHolding) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountHoldingListResponseDataHolding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountHoldingListResponseMeta struct {
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
func (r PaperTradingAccountHoldingListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountHoldingListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
