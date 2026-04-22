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

// WebhookPortalService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookPortalService] method instead.
type WebhookPortalService struct {
	options []option.RequestOption
}

// NewWebhookPortalService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookPortalService(opts ...option.RequestOption) (r WebhookPortalService) {
	r = WebhookPortalService{}
	r.options = opts
	return
}

// Get a magic link to the Svix App Portal for managing endpoints and viewing
// delivery logs.
func (r *WebhookPortalService) Get(ctx context.Context, opts ...option.RequestOption) (res *WebhookPortalGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/portal"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WebhookPortalGetResponse struct {
	Data WebhookPortalGetResponseData `json:"data" api:"required"`
	Meta WebhookPortalGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookPortalGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookPortalGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookPortalGetResponseData struct {
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookPortalGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookPortalGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookPortalGetResponseMeta struct {
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
func (r WebhookPortalGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookPortalGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
