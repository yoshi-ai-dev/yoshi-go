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
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// BenchmarkService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBenchmarkService] method instead.
type BenchmarkService struct {
	options []option.RequestOption
}

// NewBenchmarkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBenchmarkService(opts ...option.RequestOption) (r BenchmarkService) {
	r = BenchmarkService{}
	r.options = opts
	return
}

// Replay arbitrary supported benchmark tickers against a cash-flow chronology and
// return benchmark value series, return metrics, alpha inputs, and data-quality
// flags.
func (r *BenchmarkService) Replay(ctx context.Context, body BenchmarkReplayParams, opts ...option.RequestOption) (res *BenchmarkReplayResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "benchmarks/replay"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BenchmarkReplayResponse struct {
	Data BenchmarkReplayResponseData `json:"data" api:"required"`
	Meta BenchmarkReplayResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenchmarkReplayResponse) RawJSON() string { return r.JSON.raw }
func (r *BenchmarkReplayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenchmarkReplayResponseData struct {
	Allocations            []BenchmarkReplayResponseDataAllocation `json:"allocations" api:"required"`
	Alpha                  BenchmarkReplayResponseDataAlpha        `json:"alpha" api:"required"`
	BenchmarkEndValue      float64                                 `json:"benchmark_end_value" api:"required"`
	BenchmarkReturnDollars float64                                 `json:"benchmark_return_dollars" api:"required"`
	BenchmarkReturnPercent float64                                 `json:"benchmark_return_percent" api:"required"`
	DataQuality            BenchmarkReplayResponseDataDataQuality  `json:"data_quality" api:"required"`
	EndDate                string                                  `json:"end_date" api:"required"`
	GrossContributions     float64                                 `json:"gross_contributions" api:"required"`
	NetCashFlows           float64                                 `json:"net_cash_flows" api:"required"`
	Series                 []BenchmarkReplayResponseDataSeries     `json:"series" api:"required"`
	StartDate              string                                  `json:"start_date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allocations            respjson.Field
		Alpha                  respjson.Field
		BenchmarkEndValue      respjson.Field
		BenchmarkReturnDollars respjson.Field
		BenchmarkReturnPercent respjson.Field
		DataQuality            respjson.Field
		EndDate                respjson.Field
		GrossContributions     respjson.Field
		NetCashFlows           respjson.Field
		Series                 respjson.Field
		StartDate              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenchmarkReplayResponseData) RawJSON() string { return r.JSON.raw }
func (r *BenchmarkReplayResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenchmarkReplayResponseDataAllocation struct {
	Name       string  `json:"name" api:"required"`
	SecurityID string  `json:"security_id" api:"required"`
	Ticker     string  `json:"ticker" api:"required"`
	Weight     float64 `json:"weight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		SecurityID  respjson.Field
		Ticker      respjson.Field
		Weight      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenchmarkReplayResponseDataAllocation) RawJSON() string { return r.JSON.raw }
func (r *BenchmarkReplayResponseDataAllocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenchmarkReplayResponseDataAlpha struct {
	ActualEndValue float64 `json:"actual_end_value" api:"required"`
	DollarAlpha    float64 `json:"dollar_alpha" api:"required"`
	PercentAlpha   float64 `json:"percent_alpha" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualEndValue respjson.Field
		DollarAlpha    respjson.Field
		PercentAlpha   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenchmarkReplayResponseDataAlpha) RawJSON() string { return r.JSON.raw }
func (r *BenchmarkReplayResponseDataAlpha) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenchmarkReplayResponseDataDataQuality struct {
	Flags    []string `json:"flags" api:"required"`
	Messages []string `json:"messages" api:"required"`
	// Any of "ok", "partial", "error".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flags       respjson.Field
		Messages    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenchmarkReplayResponseDataDataQuality) RawJSON() string { return r.JSON.raw }
func (r *BenchmarkReplayResponseDataDataQuality) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenchmarkReplayResponseDataSeries struct {
	Date  string  `json:"date" api:"required"`
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BenchmarkReplayResponseDataSeries) RawJSON() string { return r.JSON.raw }
func (r *BenchmarkReplayResponseDataSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenchmarkReplayResponseMeta struct {
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
func (r BenchmarkReplayResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *BenchmarkReplayResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BenchmarkReplayParams struct {
	Allocations []BenchmarkReplayParamsAllocation `json:"allocations,omitzero" api:"required"`
	CashFlows   []BenchmarkReplayParamsCashFlow   `json:"cash_flows,omitzero" api:"required"`
	// Optional actual ending portfolio value used to calculate dollar and percent
	// alpha.
	ActualEndValue param.Opt[float64] `json:"actual_end_value,omitzero"`
	EndDate        param.Opt[string]  `json:"end_date,omitzero"`
	StartDate      param.Opt[string]  `json:"start_date,omitzero"`
	paramObj
}

func (r BenchmarkReplayParams) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkReplayParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkReplayParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Ticker, Weight are required.
type BenchmarkReplayParamsAllocation struct {
	Ticker string  `json:"ticker" api:"required"`
	Weight float64 `json:"weight" api:"required"`
	paramObj
}

func (r BenchmarkReplayParamsAllocation) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkReplayParamsAllocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkReplayParamsAllocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Date are required.
type BenchmarkReplayParamsCashFlow struct {
	// External cash flow amount. Contributions are positive; withdrawals are negative.
	Amount float64 `json:"amount" api:"required"`
	Date   string  `json:"date" api:"required"`
	paramObj
}

func (r BenchmarkReplayParamsCashFlow) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkReplayParamsCashFlow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkReplayParamsCashFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
