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
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// AutomationService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationService] method instead.
type AutomationService struct {
	options []option.RequestOption
}

// NewAutomationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutomationService(opts ...option.RequestOption) (r AutomationService) {
	r = AutomationService{}
	r.options = opts
	return
}

// List the user automations (active and paused), excluding deleted ones.
func (r *AutomationService) List(ctx context.Context, query AutomationListParams, opts ...option.RequestOption) (res *AutomationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "automations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AutomationListResponse struct {
	Data AutomationListResponseData `json:"data" api:"required"`
	Meta AutomationListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationListResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationListResponseData struct {
	Automations []AutomationListResponseDataAutomation `json:"automations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Automations respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AutomationListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationListResponseDataAutomation struct {
	ID              string `json:"id" api:"required"`
	ApprovalMode    string `json:"approval_mode" api:"required"`
	CreatedAt       string `json:"created_at" api:"required"`
	Description     string `json:"description" api:"required"`
	NextScheduledAt string `json:"next_scheduled_at" api:"required"`
	Origin          string `json:"origin" api:"required"`
	// Any of "active", "paused".
	Status       string `json:"status" api:"required"`
	Title        string `json:"title" api:"required"`
	TriggerEvent string `json:"trigger_event" api:"required"`
	TriggerMode  string `json:"trigger_mode" api:"required"`
	Type         string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ApprovalMode    respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		NextScheduledAt respjson.Field
		Origin          respjson.Field
		Status          respjson.Field
		Title           respjson.Field
		TriggerEvent    respjson.Field
		TriggerMode     respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationListResponseDataAutomation) RawJSON() string { return r.JSON.raw }
func (r *AutomationListResponseDataAutomation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationListResponseMeta struct {
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
func (r AutomationListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AutomationListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationListParams struct {
	// Filter by lifecycle status
	//
	// Any of "active", "paused".
	Status AutomationListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AutomationListParams]'s query parameters as `url.Values`.
func (r AutomationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by lifecycle status
type AutomationListParamsStatus string

const (
	AutomationListParamsStatusActive AutomationListParamsStatus = "active"
	AutomationListParamsStatusPaused AutomationListParamsStatus = "paused"
)
