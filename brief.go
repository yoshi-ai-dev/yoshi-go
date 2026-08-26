// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"encoding/json"
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

// Get a single Brief thread and its canonical ordered content sections.
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
	ID string `json:"id" api:"required"`
	// Any of "unread", "needs_action", "completed".
	AttentionStatus string                                    `json:"attention_status" api:"required"`
	BriefID         string                                    `json:"brief_id" api:"required" format:"uuid"`
	BriefType       string                                    `json:"brief_type" api:"required"`
	ContentSections []BriefGetResponseDataContentSectionUnion `json:"content_sections" api:"required"`
	CreatedAt       string                                    `json:"created_at" api:"required"`
	Decision        string                                    `json:"decision" api:"required"`
	Description     string                                    `json:"description" api:"required"`
	// Any of "active", "archived", "deleted".
	LifecycleStatus string `json:"lifecycle_status" api:"required"`
	Read            bool   `json:"read" api:"required"`
	Status          string `json:"status" api:"required"`
	ThreadID        string `json:"thread_id" api:"required"`
	Title           string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttentionStatus respjson.Field
		BriefID         respjson.Field
		BriefType       respjson.Field
		ContentSections respjson.Field
		CreatedAt       respjson.Field
		Decision        respjson.Field
		Description     respjson.Field
		LifecycleStatus respjson.Field
		Read            respjson.Field
		Status          respjson.Field
		ThreadID        respjson.Field
		Title           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BriefGetResponseDataContentSectionUnion contains all possible properties and
// values from [BriefGetResponseDataContentSectionObject],
// [BriefGetResponseDataContentSectionObject2],
// [BriefGetResponseDataContentSectionObject3],
// [BriefGetResponseDataContentSectionObject4],
// [BriefGetResponseDataContentSectionObject5].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BriefGetResponseDataContentSectionUnion struct {
	Key string `json:"key"`
	// This field is from variant [BriefGetResponseDataContentSectionObject].
	Markdown string `json:"markdown"`
	Type     string `json:"type"`
	// This field is a union of [[]BriefGetResponseDataContentSectionObject2Row],
	// [[]BriefGetResponseDataContentSectionObject5Row]
	Rows BriefGetResponseDataContentSectionUnionRows `json:"rows"`
	// This field is from variant [BriefGetResponseDataContentSectionObject2].
	Presentation string `json:"presentation"`
	Title        string `json:"title"`
	// This field is from variant [BriefGetResponseDataContentSectionObject3].
	Body string `json:"body"`
	// This field is from variant [BriefGetResponseDataContentSectionObject3].
	Tone string `json:"tone"`
	// This field is from variant [BriefGetResponseDataContentSectionObject4].
	Evidence BriefGetResponseDataContentSectionObject4Evidence `json:"evidence"`
	// This field is from variant [BriefGetResponseDataContentSectionObject5].
	Caption string `json:"caption"`
	// This field is from variant [BriefGetResponseDataContentSectionObject5].
	Columns []BriefGetResponseDataContentSectionObject5Column `json:"columns"`
	JSON    struct {
		Key          respjson.Field
		Markdown     respjson.Field
		Type         respjson.Field
		Rows         respjson.Field
		Presentation respjson.Field
		Title        respjson.Field
		Body         respjson.Field
		Tone         respjson.Field
		Evidence     respjson.Field
		Caption      respjson.Field
		Columns      respjson.Field
		raw          string
	} `json:"-"`
}

func (u BriefGetResponseDataContentSectionUnion) AsBriefGetResponseDataContentSectionObject() (v BriefGetResponseDataContentSectionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BriefGetResponseDataContentSectionUnion) AsBriefGetResponseDataContentSectionObject2() (v BriefGetResponseDataContentSectionObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BriefGetResponseDataContentSectionUnion) AsBriefGetResponseDataContentSectionObject3() (v BriefGetResponseDataContentSectionObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BriefGetResponseDataContentSectionUnion) AsBriefGetResponseDataContentSectionObject4() (v BriefGetResponseDataContentSectionObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BriefGetResponseDataContentSectionUnion) AsBriefGetResponseDataContentSectionObject5() (v BriefGetResponseDataContentSectionObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BriefGetResponseDataContentSectionUnion) RawJSON() string { return u.JSON.raw }

func (r *BriefGetResponseDataContentSectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BriefGetResponseDataContentSectionUnionRows is an implicit subunion of
// [BriefGetResponseDataContentSectionUnion].
// BriefGetResponseDataContentSectionUnionRows provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BriefGetResponseDataContentSectionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBriefGetResponseDataContentSectionObject2Rows
// OfBriefGetResponseDataContentSectionObject5Rows]
type BriefGetResponseDataContentSectionUnionRows struct {
	// This field will be present if the value is a
	// [[]BriefGetResponseDataContentSectionObject2Row] instead of an object.
	OfBriefGetResponseDataContentSectionObject2Rows []BriefGetResponseDataContentSectionObject2Row `json:",inline"`
	// This field will be present if the value is a
	// [[]BriefGetResponseDataContentSectionObject5Row] instead of an object.
	OfBriefGetResponseDataContentSectionObject5Rows []BriefGetResponseDataContentSectionObject5Row `json:",inline"`
	JSON                                            struct {
		OfBriefGetResponseDataContentSectionObject2Rows respjson.Field
		OfBriefGetResponseDataContentSectionObject5Rows respjson.Field
		raw                                             string
	} `json:"-"`
}

func (r *BriefGetResponseDataContentSectionUnionRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject struct {
	Key      string `json:"key" api:"required"`
	Markdown string `json:"markdown" api:"required"`
	// Any of "markdown".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Markdown    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject2 struct {
	Key  string                                         `json:"key" api:"required"`
	Rows []BriefGetResponseDataContentSectionObject2Row `json:"rows" api:"required"`
	// Any of "row_group".
	Type string `json:"type" api:"required"`
	// Any of "compact".
	Presentation string `json:"presentation"`
	Title        string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key          respjson.Field
		Rows         respjson.Field
		Type         respjson.Field
		Presentation respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject2) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject2Row struct {
	Key            string `json:"key" api:"required"`
	Label          string `json:"label" api:"required"`
	SupportingText string `json:"supporting_text" api:"required"`
	Value          string `json:"value" api:"required"`
	// Any of "neutral", "info", "warning", "success".
	Tone string `json:"tone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key            respjson.Field
		Label          respjson.Field
		SupportingText respjson.Field
		Value          respjson.Field
		Tone           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject2Row) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject2Row) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject3 struct {
	Body  string `json:"body" api:"required"`
	Key   string `json:"key" api:"required"`
	Title string `json:"title" api:"required"`
	// Any of "neutral", "info", "warning", "success".
	Tone string `json:"tone" api:"required"`
	// Any of "callout".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Key         respjson.Field
		Title       respjson.Field
		Tone        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject3) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject4 struct {
	Evidence BriefGetResponseDataContentSectionObject4Evidence `json:"evidence" api:"required"`
	Key      string                                            `json:"key" api:"required"`
	// Any of "evidence".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Evidence    respjson.Field
		Key         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject4) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject4Evidence struct {
	Metrics []BriefGetResponseDataContentSectionObject4EvidenceMetric `json:"metrics" api:"required"`
	Visual  BriefGetResponseDataContentSectionObject4EvidenceVisual   `json:"visual" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics     respjson.Field
		Visual      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject4Evidence) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject4Evidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject4EvidenceMetric struct {
	Label           string `json:"label" api:"required"`
	SupportingLabel string `json:"supporting_label" api:"required"`
	// Any of "positive", "negative", "neutral".
	Tone       string `json:"tone" api:"required"`
	ValueLabel string `json:"value_label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label           respjson.Field
		SupportingLabel respjson.Field
		Tone            respjson.Field
		ValueLabel      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject4EvidenceMetric) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject4EvidenceMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject4EvidenceVisual struct {
	Bars []BriefGetResponseDataContentSectionObject4EvidenceVisualBar `json:"bars" api:"required"`
	// Any of "spark", "bars", "grid".
	Kind   string    `json:"kind" api:"required"`
	Points []float64 `json:"points" api:"required"`
	Title  string    `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bars        respjson.Field
		Kind        respjson.Field
		Points      respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject4EvidenceVisual) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject4EvidenceVisual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject4EvidenceVisualBar struct {
	Label      string  `json:"label" api:"required"`
	Value      float64 `json:"value" api:"required"`
	ValueLabel string  `json:"value_label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Value       respjson.Field
		ValueLabel  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject4EvidenceVisualBar) RawJSON() string {
	return r.JSON.raw
}
func (r *BriefGetResponseDataContentSectionObject4EvidenceVisualBar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject5 struct {
	Caption string                                            `json:"caption" api:"required"`
	Columns []BriefGetResponseDataContentSectionObject5Column `json:"columns" api:"required"`
	Key     string                                            `json:"key" api:"required"`
	Rows    []BriefGetResponseDataContentSectionObject5Row    `json:"rows" api:"required"`
	// Any of "table".
	Type  string `json:"type" api:"required"`
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Columns     respjson.Field
		Key         respjson.Field
		Rows        respjson.Field
		Type        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject5) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject5Column struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "label", "numeric", "text".
	Presentation string `json:"presentation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key          respjson.Field
		Label        respjson.Field
		Presentation respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject5Column) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject5Column) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefGetResponseDataContentSectionObject5Row struct {
	Cells map[string]string `json:"cells" api:"required"`
	Key   string            `json:"key" api:"required"`
	// Any of "proposal", "net_effect", "total", "warning", "informational".
	Role string `json:"role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cells       respjson.Field
		Key         respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefGetResponseDataContentSectionObject5Row) RawJSON() string { return r.JSON.raw }
func (r *BriefGetResponseDataContentSectionObject5Row) UnmarshalJSON(data []byte) error {
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
	ID string `json:"id" api:"required"`
	// Any of "unread", "needs_action", "completed".
	AttentionStatus BriefListResponseAttentionStatus `json:"attention_status" api:"required"`
	BriefID         string                           `json:"brief_id" api:"required" format:"uuid"`
	BriefType       string                           `json:"brief_type" api:"required"`
	CreatedAt       string                           `json:"created_at" api:"required"`
	Decision        string                           `json:"decision" api:"required"`
	Description     string                           `json:"description" api:"required"`
	// Any of "active", "archived", "deleted".
	LifecycleStatus BriefListResponseLifecycleStatus `json:"lifecycle_status" api:"required"`
	Read            bool                             `json:"read" api:"required"`
	Status          string                           `json:"status" api:"required"`
	ThreadID        string                           `json:"thread_id" api:"required"`
	Title           string                           `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttentionStatus respjson.Field
		BriefID         respjson.Field
		BriefType       respjson.Field
		CreatedAt       respjson.Field
		Decision        respjson.Field
		Description     respjson.Field
		LifecycleStatus respjson.Field
		Read            respjson.Field
		Status          respjson.Field
		ThreadID        respjson.Field
		Title           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BriefListResponse) RawJSON() string { return r.JSON.raw }
func (r *BriefListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BriefListResponseAttentionStatus string

const (
	BriefListResponseAttentionStatusUnread      BriefListResponseAttentionStatus = "unread"
	BriefListResponseAttentionStatusNeedsAction BriefListResponseAttentionStatus = "needs_action"
	BriefListResponseAttentionStatusCompleted   BriefListResponseAttentionStatus = "completed"
)

type BriefListResponseLifecycleStatus string

const (
	BriefListResponseLifecycleStatusActive   BriefListResponseLifecycleStatus = "active"
	BriefListResponseLifecycleStatusArchived BriefListResponseLifecycleStatus = "archived"
	BriefListResponseLifecycleStatusDeleted  BriefListResponseLifecycleStatus = "deleted"
)

type BriefListParams struct {
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by brief classification
	//
	// Any of "needs_you", "insight", "alert".
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
	BriefListParamsBriefTypeNeedsYou BriefListParamsBriefType = "needs_you"
	BriefListParamsBriefTypeInsight  BriefListParamsBriefType = "insight"
	BriefListParamsBriefTypeAlert    BriefListParamsBriefType = "alert"
)
