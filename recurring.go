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
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// RecurringService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecurringService] method instead.
type RecurringService struct {
	options []option.RequestOption
}

// NewRecurringService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecurringService(opts ...option.RequestOption) (r RecurringService) {
	r = RecurringService{}
	r.options = opts
	return
}

// List recurring transaction streams (subscriptions, bills, income).
func (r *RecurringService) List(ctx context.Context, opts ...option.RequestOption) (res *RecurringListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type RecurringListResponse struct {
	Data RecurringListResponseData `json:"data" api:"required"`
	Meta RecurringListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurringListResponse) RawJSON() string { return r.JSON.raw }
func (r *RecurringListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurringListResponseData struct {
	Streams []RecurringListResponseDataStream `json:"streams" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Streams     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurringListResponseData) RawJSON() string { return r.JSON.raw }
func (r *RecurringListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurringListResponseDataStream struct {
	ID                string  `json:"id" api:"required"`
	AccountID         string  `json:"account_id" api:"required"`
	AccountName       string  `json:"account_name" api:"required"`
	AverageAmount     float64 `json:"average_amount" api:"required"`
	CurrencyCode      string  `json:"currency_code" api:"required"`
	Description       string  `json:"description" api:"required"`
	Direction         string  `json:"direction" api:"required"`
	FirstDate         string  `json:"first_date" api:"required"`
	Frequency         string  `json:"frequency" api:"required"`
	IsActive          bool    `json:"is_active" api:"required"`
	LastAmount        float64 `json:"last_amount" api:"required"`
	LastDate          string  `json:"last_date" api:"required"`
	MerchantName      string  `json:"merchant_name" api:"required"`
	PredictedNextDate string  `json:"predicted_next_date" api:"required"`
	Status            string  `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AccountID         respjson.Field
		AccountName       respjson.Field
		AverageAmount     respjson.Field
		CurrencyCode      respjson.Field
		Description       respjson.Field
		Direction         respjson.Field
		FirstDate         respjson.Field
		Frequency         respjson.Field
		IsActive          respjson.Field
		LastAmount        respjson.Field
		LastDate          respjson.Field
		MerchantName      respjson.Field
		PredictedNextDate respjson.Field
		Status            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurringListResponseDataStream) RawJSON() string { return r.JSON.raw }
func (r *RecurringListResponseDataStream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurringListResponseMeta struct {
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
func (r RecurringListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *RecurringListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
