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

// CreditDebtService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditDebtService] method instead.
type CreditDebtService struct {
	options []option.RequestOption
}

// NewCreditDebtService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCreditDebtService(opts ...option.RequestOption) (r CreditDebtService) {
	r = CreditDebtService{}
	r.options = opts
	return
}

// Get the latest credit score alongside a breakdown of debt accounts (credit cards
// and loans) with balances, rates, and payments.
func (r *CreditDebtService) Get(ctx context.Context, opts ...option.RequestOption) (res *CreditDebtGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credit-debt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CreditDebtGetResponse struct {
	Data CreditDebtGetResponseData `json:"data" api:"required"`
	Meta CreditDebtGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditDebtGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditDebtGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditDebtGetResponseData struct {
	CreditScore          CreditDebtGetResponseDataCreditScore  `json:"credit_score" api:"required"`
	DebtByType           []CreditDebtGetResponseDataDebtByType `json:"debt_by_type" api:"required"`
	Debts                []CreditDebtGetResponseDataDebt       `json:"debts" api:"required"`
	TotalDebt            float64                               `json:"total_debt" api:"required"`
	TotalMinimumPayments float64                               `json:"total_minimum_payments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditScore          respjson.Field
		DebtByType           respjson.Field
		Debts                respjson.Field
		TotalDebt            respjson.Field
		TotalMinimumPayments respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditDebtGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *CreditDebtGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditDebtGetResponseDataCreditScore struct {
	AsOf   string  `json:"as_of" api:"required"`
	Score  float64 `json:"score" api:"required"`
	Source string  `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsOf        respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditDebtGetResponseDataCreditScore) RawJSON() string { return r.JSON.raw }
func (r *CreditDebtGetResponseDataCreditScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditDebtGetResponseDataDebtByType struct {
	AccountCount int64   `json:"account_count" api:"required"`
	TotalBalance float64 `json:"total_balance" api:"required"`
	Type         string  `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountCount respjson.Field
		TotalBalance respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditDebtGetResponseDataDebtByType) RawJSON() string { return r.JSON.raw }
func (r *CreditDebtGetResponseDataDebtByType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditDebtGetResponseDataDebt struct {
	ID                       string    `json:"id" api:"required"`
	BalanceCurrent           float64   `json:"balance_current" api:"required"`
	BalanceLimit             float64   `json:"balance_limit" api:"required"`
	InstitutionName          string    `json:"institution_name" api:"required"`
	InterestRatePercentage   float64   `json:"interest_rate_percentage" api:"required"`
	InterestRateType         string    `json:"interest_rate_type" api:"required"`
	LastPaymentAmount        float64   `json:"last_payment_amount" api:"required"`
	LastPaymentDate          time.Time `json:"last_payment_date" api:"required" format:"date"`
	Name                     string    `json:"name" api:"required"`
	NextPaymentDueDate       time.Time `json:"next_payment_due_date" api:"required" format:"date"`
	NextPaymentMinimumAmount float64   `json:"next_payment_minimum_amount" api:"required"`
	Subtype                  string    `json:"subtype" api:"required"`
	Type                     string    `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		BalanceCurrent           respjson.Field
		BalanceLimit             respjson.Field
		InstitutionName          respjson.Field
		InterestRatePercentage   respjson.Field
		InterestRateType         respjson.Field
		LastPaymentAmount        respjson.Field
		LastPaymentDate          respjson.Field
		Name                     respjson.Field
		NextPaymentDueDate       respjson.Field
		NextPaymentMinimumAmount respjson.Field
		Subtype                  respjson.Field
		Type                     respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditDebtGetResponseDataDebt) RawJSON() string { return r.JSON.raw }
func (r *CreditDebtGetResponseDataDebt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditDebtGetResponseMeta struct {
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
func (r CreditDebtGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *CreditDebtGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
