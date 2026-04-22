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
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// WebhookEndpointService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookEndpointService] method instead.
type WebhookEndpointService struct {
	options []option.RequestOption
}

// NewWebhookEndpointService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookEndpointService(opts ...option.RequestOption) (r WebhookEndpointService) {
	r = WebhookEndpointService{}
	r.options = opts
	return
}

// Register a new webhook endpoint. Maximum 5 endpoints per user.
func (r *WebhookEndpointService) New(ctx context.Context, body WebhookEndpointNewParams, opts ...option.RequestOption) (res *WebhookEndpointNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get details for a specific webhook endpoint.
func (r *WebhookEndpointService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookEndpointGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a webhook endpoint (URL, event filter, status, description).
func (r *WebhookEndpointService) Update(ctx context.Context, id string, body WebhookEndpointUpdateParams, opts ...option.RequestOption) (res *WebhookEndpointUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List the user's webhook endpoints.
func (r *WebhookEndpointService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookEndpointListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a webhook endpoint. This stops all future deliveries to this URL.
func (r *WebhookEndpointService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookEndpointDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Rotate the webhook signing secret. Both old and new secrets are valid during
// rotation.
func (r *WebhookEndpointService) Rotate(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookEndpointRotateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/endpoints/%s/rotate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Send a synthetic test event to verify your endpoint is receiving webhooks
// correctly.
func (r *WebhookEndpointService) Test(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookEndpointTestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/endpoints/%s/test", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebhookEndpointNewResponse struct {
	Data WebhookEndpointNewResponseData `json:"data" api:"required"`
	Meta WebhookEndpointNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointNewResponseData struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Description string    `json:"description" api:"required"`
	FilterTypes []string  `json:"filter_types" api:"required"`
	Secret      string    `json:"secret" api:"required"`
	Status      string    `json:"status" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	URL         string    `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		FilterTypes respjson.Field
		Secret      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointNewResponseMeta struct {
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
func (r WebhookEndpointNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointGetResponse struct {
	Data WebhookEndpointGetResponseData `json:"data" api:"required"`
	Meta WebhookEndpointGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointGetResponseData struct {
	Endpoint WebhookEndpointGetResponseDataEndpoint `json:"endpoint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoint    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointGetResponseDataEndpoint struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Description string    `json:"description" api:"required"`
	FilterTypes []string  `json:"filter_types" api:"required"`
	Status      string    `json:"status" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	URL         string    `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		FilterTypes respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointGetResponseDataEndpoint) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointGetResponseDataEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointGetResponseMeta struct {
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
func (r WebhookEndpointGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointUpdateResponse struct {
	Data WebhookEndpointUpdateResponseData `json:"data" api:"required"`
	Meta WebhookEndpointUpdateResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointUpdateResponseData struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Description string    `json:"description" api:"required"`
	FilterTypes []string  `json:"filter_types" api:"required"`
	Status      string    `json:"status" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	URL         string    `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		FilterTypes respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointUpdateResponseMeta struct {
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
func (r WebhookEndpointUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointListResponse struct {
	Data WebhookEndpointListResponseData `json:"data" api:"required"`
	Meta WebhookEndpointListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointListResponseData struct {
	Endpoints []WebhookEndpointListResponseDataEndpoint `json:"endpoints" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoints   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointListResponseDataEndpoint struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Description string    `json:"description" api:"required"`
	FilterTypes []string  `json:"filter_types" api:"required"`
	Status      string    `json:"status" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	URL         string    `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		FilterTypes respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointListResponseDataEndpoint) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointListResponseDataEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointListResponseMeta struct {
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
func (r WebhookEndpointListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointDeleteResponse struct {
	Data WebhookEndpointDeleteResponseData `json:"data" api:"required"`
	Meta WebhookEndpointDeleteResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointDeleteResponseData struct {
	// Any of true.
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointDeleteResponseMeta struct {
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
func (r WebhookEndpointDeleteResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointDeleteResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointRotateResponse struct {
	Data WebhookEndpointRotateResponseData `json:"data" api:"required"`
	Meta WebhookEndpointRotateResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointRotateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointRotateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointRotateResponseData struct {
	Secret string `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointRotateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointRotateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointRotateResponseMeta struct {
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
func (r WebhookEndpointRotateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointRotateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointTestResponse struct {
	Data WebhookEndpointTestResponseData `json:"data" api:"required"`
	Meta WebhookEndpointTestResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointTestResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointTestResponseData struct {
	EventID string `json:"event_id" api:"required"`
	// Any of true.
	Sent bool `json:"sent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID     respjson.Field
		Sent        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEndpointTestResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointTestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointTestResponseMeta struct {
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
func (r WebhookEndpointTestResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookEndpointTestResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointNewParams struct {
	URL         string            `json:"url" api:"required" format:"uri"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Event types to subscribe to. Empty array means all events.
	FilterTypes []string `json:"filter_types,omitzero"`
	paramObj
}

func (r WebhookEndpointNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookEndpointNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookEndpointNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	URL         param.Opt[string] `json:"url,omitzero" format:"uri"`
	FilterTypes []string          `json:"filter_types,omitzero"`
	// Any of "active", "disabled".
	Status WebhookEndpointUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r WebhookEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookEndpointUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookEndpointUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEndpointUpdateParamsStatus string

const (
	WebhookEndpointUpdateParamsStatusActive   WebhookEndpointUpdateParamsStatus = "active"
	WebhookEndpointUpdateParamsStatusDisabled WebhookEndpointUpdateParamsStatus = "disabled"
)
