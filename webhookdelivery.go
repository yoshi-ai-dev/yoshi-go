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

// WebhookDeliveryService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookDeliveryService] method instead.
type WebhookDeliveryService struct {
	options []option.RequestOption
}

// NewWebhookDeliveryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookDeliveryService(opts ...option.RequestOption) (r WebhookDeliveryService) {
	r = WebhookDeliveryService{}
	r.options = opts
	return
}

// List recent webhook delivery attempts from the event log.
func (r *WebhookDeliveryService) List(ctx context.Context, query WebhookDeliveryListParams, opts ...option.RequestOption) (res *WebhookDeliveryListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/deliveries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WebhookDeliveryListResponse struct {
	Data WebhookDeliveryListResponseData `json:"data" api:"required"`
	Meta WebhookDeliveryListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListResponseData struct {
	Deliveries []WebhookDeliveryListResponseDataDelivery `json:"deliveries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deliveries  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListResponseDataDelivery struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	EventID   string    `json:"event_id" api:"required"`
	EventType string    `json:"event_type" api:"required"`
	Status    string    `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventID     respjson.Field
		EventType   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListResponseMeta struct {
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
func (r WebhookDeliveryListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListParams struct {
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
