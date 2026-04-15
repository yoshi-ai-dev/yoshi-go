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

// IncomeService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIncomeService] method instead.
type IncomeService struct {
	options []option.RequestOption
}

// NewIncomeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIncomeService(opts ...option.RequestOption) (r IncomeService) {
	r = IncomeService{}
	r.options = opts
	return
}

// Get income analysis and distribution.
func (r *IncomeService) Get(ctx context.Context, opts ...option.RequestOption) (res *IncomeGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "income"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type IncomeGetResponse struct {
	Data IncomeGetResponseData `json:"data" api:"required"`
	Meta IncomeGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncomeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncomeGetResponseData struct {
	Debt          IncomeGetResponseDataDebt          `json:"debt" api:"required"`
	EmergencyFund IncomeGetResponseDataEmergencyFund `json:"emergency_fund" api:"required"`
	Expenses      IncomeGetResponseDataExpenses      `json:"expenses" api:"required"`
	GeneratedAt   string                             `json:"generated_at" api:"required"`
	Investments   IncomeGetResponseDataInvestments   `json:"investments" api:"required"`
	MonthlyIncome float64                            `json:"monthly_income" api:"required"`
	Variability   IncomeGetResponseDataVariability   `json:"variability" api:"required"`
	Window        IncomeGetResponseDataWindow        `json:"window" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Debt          respjson.Field
		EmergencyFund respjson.Field
		Expenses      respjson.Field
		GeneratedAt   respjson.Field
		Investments   respjson.Field
		MonthlyIncome respjson.Field
		Variability   respjson.Field
		Window        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncomeGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncomeGetResponseDataDebt struct {
	Amount    float64            `json:"amount" api:"required"`
	Percent   float64            `json:"percent" api:"required"`
	Breakdown map[string]float64 `json:"breakdown"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Percent     respjson.Field
		Breakdown   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncomeGetResponseDataDebt) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponseDataDebt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncomeGetResponseDataEmergencyFund struct {
	Amount    float64            `json:"amount" api:"required"`
	Percent   float64            `json:"percent" api:"required"`
	Breakdown map[string]float64 `json:"breakdown"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Percent     respjson.Field
		Breakdown   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncomeGetResponseDataEmergencyFund) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponseDataEmergencyFund) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncomeGetResponseDataExpenses struct {
	Amount    float64            `json:"amount" api:"required"`
	Percent   float64            `json:"percent" api:"required"`
	Breakdown map[string]float64 `json:"breakdown"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Percent     respjson.Field
		Breakdown   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncomeGetResponseDataExpenses) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponseDataExpenses) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncomeGetResponseDataInvestments struct {
	Amount    float64            `json:"amount" api:"required"`
	Percent   float64            `json:"percent" api:"required"`
	Breakdown map[string]float64 `json:"breakdown"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Percent     respjson.Field
		Breakdown   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncomeGetResponseDataInvestments) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponseDataInvestments) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncomeGetResponseDataVariability struct {
	Max   float64 `json:"max" api:"required"`
	Min   float64 `json:"min" api:"required"`
	Range float64 `json:"range" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		Range       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncomeGetResponseDataVariability) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponseDataVariability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncomeGetResponseDataWindow struct {
	End    string `json:"end" api:"required"`
	Months int64  `json:"months" api:"required"`
	Start  string `json:"start" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Months      respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncomeGetResponseDataWindow) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponseDataWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncomeGetResponseMeta struct {
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
func (r IncomeGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *IncomeGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
