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

// ScoreService contains methods and other services that help with interacting with
// the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScoreService] method instead.
type ScoreService struct {
	options []option.RequestOption
}

// NewScoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScoreService(opts ...option.RequestOption) (r ScoreService) {
	r = ScoreService{}
	r.options = opts
	return
}

// Get the latest financial health scores.
func (r *ScoreService) List(ctx context.Context, opts ...option.RequestOption) (res *ScoreListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "scores"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ScoreListResponse struct {
	Data ScoreListResponseData `json:"data" api:"required"`
	Meta ScoreListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoreListResponse) RawJSON() string { return r.JSON.raw }
func (r *ScoreListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoreListResponseData struct {
	Baseline        float64 `json:"baseline" api:"required"`
	Capacity        float64 `json:"capacity" api:"required"`
	ObservationDate string  `json:"observation_date" api:"required"`
	Recovery        float64 `json:"recovery" api:"required"`
	Yoshi           float64 `json:"yoshi" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Baseline        respjson.Field
		Capacity        respjson.Field
		ObservationDate respjson.Field
		Recovery        respjson.Field
		Yoshi           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoreListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ScoreListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoreListResponseMeta struct {
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
func (r ScoreListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ScoreListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
