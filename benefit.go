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

// BenefitService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBenefitService] method instead.
type BenefitService struct {
	options []option.RequestOption
}

// NewBenefitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBenefitService(opts ...option.RequestOption) (r BenefitService) {
	r = BenefitService{}
	r.options = opts
	return
}

// Get benefits expiring soon that have remaining value.
//
// Send the end user's IANA timezone in the `x-user-timezone` header (for example
// `America/New_York`). Calendar-day results are resolved in that zone; without it
// they resolve in UTC, which differs from the user's own day for part of every
// day.
func (r *BenefitService) Expiring(ctx context.Context, query BenefitExpiringParams, opts ...option.RequestOption) (res *BenefitExpiringResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "benefits/expiring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get a summary of the user's card benefit periods with usage and remaining value.
// A period that has ended on the caller's calendar day is excluded, so the
// reported totals depend on the timezone.
//
// Send the end user's IANA timezone in the `x-user-timezone` header (for example
// `America/New_York`). Calendar-day results are resolved in that zone; without it
// they resolve in UTC, which differs from the user's own day for part of every
// day.
func (r *BenefitService) Summary(ctx context.Context, opts ...option.RequestOption) (res *BenefitSummaryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "benefits/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type BenefitExpiringResponse struct {
	Data BenefitExpiringResponseData `json:"data" api:"required"`
	Meta BenefitExpiringResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenefitExpiringResponse) RawJSON() string { return r.JSON.raw }
func (r *BenefitExpiringResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitExpiringResponseData struct {
	Expiring []BenefitExpiringResponseDataExpiring `json:"expiring" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Expiring    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenefitExpiringResponseData) RawJSON() string { return r.JSON.raw }
func (r *BenefitExpiringResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitExpiringResponseDataExpiring struct {
	BenefitName         string  `json:"benefit_name" api:"required"`
	BenefitType         string  `json:"benefit_type" api:"required"`
	CardIssuer          string  `json:"card_issuer" api:"required"`
	CardProductName     string  `json:"card_product_name" api:"required"`
	DaysRemaining       float64 `json:"days_remaining" api:"required"`
	PeriodEnd           string  `json:"period_end" api:"required"`
	PeriodID            string  `json:"period_id" api:"required" format:"uuid"`
	Status              string  `json:"status" api:"required"`
	ValueLimitCents     float64 `json:"value_limit_cents" api:"required"`
	ValueRemainingCents float64 `json:"value_remaining_cents" api:"required"`
	ValueUsedCents      float64 `json:"value_used_cents" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BenefitName         respjson.Field
		BenefitType         respjson.Field
		CardIssuer          respjson.Field
		CardProductName     respjson.Field
		DaysRemaining       respjson.Field
		PeriodEnd           respjson.Field
		PeriodID            respjson.Field
		Status              respjson.Field
		ValueLimitCents     respjson.Field
		ValueRemainingCents respjson.Field
		ValueUsedCents      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenefitExpiringResponseDataExpiring) RawJSON() string { return r.JSON.raw }
func (r *BenefitExpiringResponseDataExpiring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitExpiringResponseMeta struct {
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
func (r BenefitExpiringResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *BenefitExpiringResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitSummaryResponse struct {
	Data BenefitSummaryResponseData `json:"data" api:"required"`
	Meta BenefitSummaryResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenefitSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *BenefitSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitSummaryResponseData struct {
	Periods             []BenefitSummaryResponseDataPeriod `json:"periods" api:"required"`
	TotalRemainingCents float64                            `json:"total_remaining_cents" api:"required"`
	TotalUsedCents      float64                            `json:"total_used_cents" api:"required"`
	TotalValueCents     float64                            `json:"total_value_cents" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Periods             respjson.Field
		TotalRemainingCents respjson.Field
		TotalUsedCents      respjson.Field
		TotalValueCents     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenefitSummaryResponseData) RawJSON() string { return r.JSON.raw }
func (r *BenefitSummaryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitSummaryResponseDataPeriod struct {
	BenefitDescription  string  `json:"benefit_description" api:"required"`
	BenefitName         string  `json:"benefit_name" api:"required"`
	BenefitType         string  `json:"benefit_type" api:"required"`
	CardIssuer          string  `json:"card_issuer" api:"required"`
	CardProductName     string  `json:"card_product_name" api:"required"`
	CardProductSlug     string  `json:"card_product_slug" api:"required"`
	DaysRemaining       float64 `json:"days_remaining" api:"required"`
	MaxUsesPerPeriod    float64 `json:"max_uses_per_period" api:"required"`
	PeriodEnd           string  `json:"period_end" api:"required"`
	PeriodID            string  `json:"period_id" api:"required" format:"uuid"`
	PeriodStart         string  `json:"period_start" api:"required"`
	ResetCadence        string  `json:"reset_cadence" api:"required"`
	Status              string  `json:"status" api:"required"`
	UsagePct            float64 `json:"usage_pct" api:"required"`
	UsesCount           float64 `json:"uses_count" api:"required"`
	ValueLimitCents     float64 `json:"value_limit_cents" api:"required"`
	ValueRemainingCents float64 `json:"value_remaining_cents" api:"required"`
	ValueUsedCents      float64 `json:"value_used_cents" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BenefitDescription  respjson.Field
		BenefitName         respjson.Field
		BenefitType         respjson.Field
		CardIssuer          respjson.Field
		CardProductName     respjson.Field
		CardProductSlug     respjson.Field
		DaysRemaining       respjson.Field
		MaxUsesPerPeriod    respjson.Field
		PeriodEnd           respjson.Field
		PeriodID            respjson.Field
		PeriodStart         respjson.Field
		ResetCadence        respjson.Field
		Status              respjson.Field
		UsagePct            respjson.Field
		UsesCount           respjson.Field
		ValueLimitCents     respjson.Field
		ValueRemainingCents respjson.Field
		ValueUsedCents      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenefitSummaryResponseDataPeriod) RawJSON() string { return r.JSON.raw }
func (r *BenefitSummaryResponseDataPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitSummaryResponseMeta struct {
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
func (r BenefitSummaryResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *BenefitSummaryResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitExpiringParams struct {
	Days param.Opt[int64] `query:"days,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BenefitExpiringParams]'s query parameters as `url.Values`.
func (r BenefitExpiringParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
