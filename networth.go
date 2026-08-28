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

// NetWorthService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetWorthService] method instead.
type NetWorthService struct {
	options []option.RequestOption
}

// NewNetWorthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetWorthService(opts ...option.RequestOption) (r NetWorthService) {
	r = NetWorthService{}
	r.options = opts
	return
}

// Get the net worth trend over time, aggregated daily across all accounts.
// Liability accounts (credit, loan) contribute negatively.
func (r *NetWorthService) History(ctx context.Context, query NetWorthHistoryParams, opts ...option.RequestOption) (res *NetWorthHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "net-worth/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type NetWorthHistoryResponse struct {
	Data NetWorthHistoryResponseData `json:"data" api:"required"`
	Meta NetWorthHistoryResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetWorthHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *NetWorthHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetWorthHistoryResponseData struct {
	CurrentNetWorth float64                                 `json:"current_net_worth" api:"required"`
	DataPoints      int64                                   `json:"data_points" api:"required"`
	History         []NetWorthHistoryResponseDataHistory    `json:"history" api:"required"`
	Period          string                                  `json:"period" api:"required"`
	PeriodChange    NetWorthHistoryResponseDataPeriodChange `json:"period_change" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentNetWorth respjson.Field
		DataPoints      respjson.Field
		History         respjson.Field
		Period          respjson.Field
		PeriodChange    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetWorthHistoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetWorthHistoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetWorthHistoryResponseDataHistory struct {
	Date     string  `json:"date" api:"required"`
	NetWorth float64 `json:"net_worth" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		NetWorth    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetWorthHistoryResponseDataHistory) RawJSON() string { return r.JSON.raw }
func (r *NetWorthHistoryResponseDataHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetWorthHistoryResponseDataPeriodChange struct {
	Amount     float64 `json:"amount" api:"required"`
	EndValue   float64 `json:"end_value" api:"required"`
	Percent    float64 `json:"percent" api:"required"`
	StartValue float64 `json:"start_value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		EndValue    respjson.Field
		Percent     respjson.Field
		StartValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetWorthHistoryResponseDataPeriodChange) RawJSON() string { return r.JSON.raw }
func (r *NetWorthHistoryResponseDataPeriodChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetWorthHistoryResponseMeta struct {
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
func (r NetWorthHistoryResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *NetWorthHistoryResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetWorthHistoryParams struct {
	// Trailing window for the trend (all is capped at 1 year)
	//
	// Any of "3m", "6m", "1y", "all".
	Period NetWorthHistoryParamsPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetWorthHistoryParams]'s query parameters as `url.Values`.
func (r NetWorthHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Trailing window for the trend (all is capped at 1 year)
type NetWorthHistoryParamsPeriod string

const (
	NetWorthHistoryParamsPeriod3m  NetWorthHistoryParamsPeriod = "3m"
	NetWorthHistoryParamsPeriod6m  NetWorthHistoryParamsPeriod = "6m"
	NetWorthHistoryParamsPeriod1y  NetWorthHistoryParamsPeriod = "1y"
	NetWorthHistoryParamsPeriodAll NetWorthHistoryParamsPeriod = "all"
)
