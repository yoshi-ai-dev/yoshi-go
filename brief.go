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

// BriefService contains methods and other services that help with interacting with
// the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBriefService] method instead.
type BriefService struct {
	options []option.RequestOption
}

// NewBriefService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBriefService(opts ...option.RequestOption) (r BriefService) {
	r = BriefService{}
	r.options = opts
	return
}

// Get a single brief thread by ID, including the full opening message.
func (r *BriefService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BriefGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("briefs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List brief threads (system-generated insights, alerts, and action requests),
// newest first.
func (r *BriefService) List(ctx context.Context, query BriefListParams, opts ...option.RequestOption) (res *pagination.CursorPage[BriefListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "briefs"
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

// List brief threads (system-generated insights, alerts, and action requests),
// newest first.
func (r *BriefService) ListAutoPaging(ctx context.Context, query BriefListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[BriefListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type BriefGetResponse struct {
	Data BriefGetResponseData `json:"data" api:"required"`
	Meta BriefGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseData struct {
	ID             string `json:"id" api:"required"`
	BriefType      string `json:"brief_type" api:"required"`
	CreatedAt      string `json:"created_at" api:"required"`
	Decision       string `json:"decision" api:"required"`
	Description    string `json:"description" api:"required"`
	InitialMessage string `json:"initial_message" api:"required"`
	Read           bool   `json:"read" api:"required"`
	Status         string `json:"status" api:"required"`
	Title          string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BriefType      respjson.Field
		CreatedAt      respjson.Field
		Decision       respjson.Field
		Description    respjson.Field
		InitialMessage respjson.Field
		Read           respjson.Field
		Status         respjson.Field
		Title          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseMeta struct {
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
func (r BriefGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefListResponse struct {
	ID          string `json:"id" api:"required"`
	BriefType   string `json:"brief_type" api:"required"`
	CreatedAt   string `json:"created_at" api:"required"`
	Decision    string `json:"decision" api:"required"`
	Description string `json:"description" api:"required"`
	Read        bool   `json:"read" api:"required"`
	Status      string `json:"status" api:"required"`
	Title       string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BriefType   respjson.Field
		CreatedAt   respjson.Field
		Decision    respjson.Field
		Description respjson.Field
		Read        respjson.Field
		Status      respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefListResponse) RawJSON() string { return r.JSON.raw }
func (r *BriefListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefListParams struct {
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by brief classification
	//
	// Any of "action_requested", "insight", "alert", "completed".
	BriefType BriefListParamsBriefType `query:"brief_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BriefListParams]'s query parameters as `url.Values`.
func (r BriefListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by brief classification
type BriefListParamsBriefType string

const (
	BriefListParamsBriefTypeActionRequested BriefListParamsBriefType = "action_requested"
	BriefListParamsBriefTypeInsight         BriefListParamsBriefType = "insight"
	BriefListParamsBriefTypeAlert           BriefListParamsBriefType = "alert"
	BriefListParamsBriefTypeCompleted       BriefListParamsBriefType = "completed"
)
