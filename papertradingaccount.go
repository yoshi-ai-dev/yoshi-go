// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// PaperTradingAccountService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperTradingAccountService] method instead.
type PaperTradingAccountService struct {
	options  []option.RequestOption
	Holdings PaperTradingAccountHoldingService
	Trades   PaperTradingAccountTradeService
}

// NewPaperTradingAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaperTradingAccountService(opts ...option.RequestOption) (r PaperTradingAccountService) {
	r = PaperTradingAccountService{}
	r.options = opts
	r.Holdings = NewPaperTradingAccountHoldingService(opts...)
	r.Trades = NewPaperTradingAccountTradeService(opts...)
	return
}

// Create a Test Drive account. Requires user approval in the Yoshi web app before
// the account is created.
func (r *PaperTradingAccountService) New(ctx context.Context, body PaperTradingAccountNewParams, opts ...option.RequestOption) (res *PaperTradingAccountNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "paper-trading/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List Test Drive accounts with current balances for the user.
func (r *PaperTradingAccountService) List(ctx context.Context, opts ...option.RequestOption) (res *PaperTradingAccountListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "paper-trading/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PaperTradingAccountNewResponse struct {
	Data PaperTradingAccountNewResponseData `json:"data" api:"required"`
	Meta PaperTradingAccountNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseData struct {
	ActionID string `json:"action_id" api:"required" format:"uuid"`
	// Any of "trade", "transfer", "data_fix", "automation", "account_create",
	// "unknown".
	ActionKind          string                                           `json:"action_kind" api:"required"`
	ActionType          string                                           `json:"action_type" api:"required"`
	ApprovalStatusURL   string                                           `json:"approval_status_url" api:"required"`
	ApprovalURL         string                                           `json:"approval_url" api:"required"`
	CreatedAt           string                                           `json:"created_at" api:"required"`
	DecidedAt           string                                           `json:"decided_at" api:"required"`
	DecisionNeeds       PaperTradingAccountNewResponseDataDecisionNeeds  `json:"decision_needs" api:"required"`
	Description         string                                           `json:"description" api:"required"`
	DetailsV1           PaperTradingAccountNewResponseDataDetailsV1Union `json:"details_v1" api:"required"`
	DisplayErrorMessage string                                           `json:"display_error_message" api:"required"`
	ExecutedAt          string                                           `json:"executed_at" api:"required"`
	ExpiresAt           string                                           `json:"expires_at" api:"required"`
	Status              string                                           `json:"status" api:"required"`
	ThreadID            string                                           `json:"thread_id" api:"required" format:"uuid"`
	UpdatedAt           string                                           `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID            respjson.Field
		ActionKind          respjson.Field
		ActionType          respjson.Field
		ApprovalStatusURL   respjson.Field
		ApprovalURL         respjson.Field
		CreatedAt           respjson.Field
		DecidedAt           respjson.Field
		DecisionNeeds       respjson.Field
		Description         respjson.Field
		DetailsV1           respjson.Field
		DisplayErrorMessage respjson.Field
		ExecutedAt          respjson.Field
		ExpiresAt           respjson.Field
		Status              respjson.Field
		ThreadID            respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseDataDecisionNeeds struct {
	Approval             bool `json:"approval" api:"required"`
	ClientExecution      bool `json:"client_execution" api:"required"`
	IdentityVerification bool `json:"identity_verification" api:"required"`
	Passkey              bool `json:"passkey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Approval             respjson.Field
		ClientExecution      respjson.Field
		IdentityVerification respjson.Field
		Passkey              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseDataDecisionNeeds) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseDataDecisionNeeds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperTradingAccountNewResponseDataDetailsV1Union contains all possible
// properties and values from [PaperTradingAccountNewResponseDataDetailsV1Object],
// [PaperTradingAccountNewResponseDataDetailsV1Object2],
// [PaperTradingAccountNewResponseDataDetailsV1Object3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperTradingAccountNewResponseDataDetailsV1Union struct {
	// This field is a union of
	// [PaperTradingAccountNewResponseDataDetailsV1ObjectAccount],
	// [PaperTradingAccountNewResponseDataDetailsV1Object2Account]
	Account                 PaperTradingAccountNewResponseDataDetailsV1UnionAccount `json:"account"`
	AmountLabel             string                                                  `json:"amount_label"`
	BottomMarkdown          string                                                  `json:"bottom_markdown"`
	CurrentPriceDisplay     string                                                  `json:"current_price_display"`
	EstimatedCostDisplay    string                                                  `json:"estimated_cost_display"`
	InputMode               string                                                  `json:"input_mode"`
	Kind                    string                                                  `json:"kind"`
	MoneyMovementDisclosure string                                                  `json:"money_movement_disclosure"`
	OrderTypeDisplay        string                                                  `json:"order_type_display"`
	PriceChangeDirection    string                                                  `json:"price_change_direction"`
	PriceChangeDisplay      string                                                  `json:"price_change_display"`
	PrimaryActionLabel      string                                                  `json:"primary_action_label"`
	ProposedByLabel         string                                                  `json:"proposed_by_label"`
	Quantity                string                                                  `json:"quantity"`
	QuantityDisplay         string                                                  `json:"quantity_display"`
	QuoteAsOf               time.Time                                               `json:"quote_as_of"`
	QuoteAsOfLabel          string                                                  `json:"quote_as_of_label"`
	ReasonMarkdown          string                                                  `json:"reason_markdown"`
	SecondaryActionLabel    string                                                  `json:"secondary_action_label"`
	SecurityName            string                                                  `json:"security_name"`
	Side                    string                                                  `json:"side"`
	SideLabel               string                                                  `json:"side_label"`
	Symbol                  string                                                  `json:"symbol"`
	TimingDisclosure        string                                                  `json:"timing_disclosure"`
	TopMarkdown             string                                                  `json:"top_markdown"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	Amount string `json:"amount"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	AmountDisplay string `json:"amount_display"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	CurrencyCode string `json:"currency_code"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	Description string `json:"description"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	From PaperTradingAccountNewResponseDataDetailsV1Object3From `json:"from"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	Method string `json:"method"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	MethodDisplay string `json:"method_display"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	RequestID string `json:"request_id"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	TimingLabel string `json:"timing_label"`
	// This field is from variant [PaperTradingAccountNewResponseDataDetailsV1Object3].
	To   PaperTradingAccountNewResponseDataDetailsV1Object3To `json:"to"`
	JSON struct {
		Account                 respjson.Field
		AmountLabel             respjson.Field
		BottomMarkdown          respjson.Field
		CurrentPriceDisplay     respjson.Field
		EstimatedCostDisplay    respjson.Field
		InputMode               respjson.Field
		Kind                    respjson.Field
		MoneyMovementDisclosure respjson.Field
		OrderTypeDisplay        respjson.Field
		PriceChangeDirection    respjson.Field
		PriceChangeDisplay      respjson.Field
		PrimaryActionLabel      respjson.Field
		ProposedByLabel         respjson.Field
		Quantity                respjson.Field
		QuantityDisplay         respjson.Field
		QuoteAsOf               respjson.Field
		QuoteAsOfLabel          respjson.Field
		ReasonMarkdown          respjson.Field
		SecondaryActionLabel    respjson.Field
		SecurityName            respjson.Field
		Side                    respjson.Field
		SideLabel               respjson.Field
		Symbol                  respjson.Field
		TimingDisclosure        respjson.Field
		TopMarkdown             respjson.Field
		Amount                  respjson.Field
		AmountDisplay           respjson.Field
		CurrencyCode            respjson.Field
		Description             respjson.Field
		From                    respjson.Field
		Method                  respjson.Field
		MethodDisplay           respjson.Field
		RequestID               respjson.Field
		TimingLabel             respjson.Field
		To                      respjson.Field
		raw                     string
	} `json:"-"`
}

func (u PaperTradingAccountNewResponseDataDetailsV1Union) AsPaperTradingAccountNewResponseDataDetailsV1Object() (v PaperTradingAccountNewResponseDataDetailsV1Object) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperTradingAccountNewResponseDataDetailsV1Union) AsPaperTradingAccountNewResponseDataDetailsV1Object2() (v PaperTradingAccountNewResponseDataDetailsV1Object2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperTradingAccountNewResponseDataDetailsV1Union) AsPaperTradingAccountNewResponseDataDetailsV1Object3() (v PaperTradingAccountNewResponseDataDetailsV1Object3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperTradingAccountNewResponseDataDetailsV1Union) RawJSON() string { return u.JSON.raw }

func (r *PaperTradingAccountNewResponseDataDetailsV1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperTradingAccountNewResponseDataDetailsV1UnionAccount is an implicit subunion
// of [PaperTradingAccountNewResponseDataDetailsV1Union].
// PaperTradingAccountNewResponseDataDetailsV1UnionAccount provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperTradingAccountNewResponseDataDetailsV1Union].
type PaperTradingAccountNewResponseDataDetailsV1UnionAccount struct {
	ID                     string `json:"id"`
	AvailableAmountDisplay string `json:"available_amount_display"`
	AvailableAmountLabel   string `json:"available_amount_label"`
	AvailableAmountState   string `json:"available_amount_state"`
	DisplayName            string `json:"display_name"`
	IconURL                string `json:"icon_url"`
	InstitutionName        string `json:"institution_name"`
	Source                 string `json:"source"`
	SourceLabel            string `json:"source_label"`
	Subtitle               string `json:"subtitle"`
	JSON                   struct {
		ID                     respjson.Field
		AvailableAmountDisplay respjson.Field
		AvailableAmountLabel   respjson.Field
		AvailableAmountState   respjson.Field
		DisplayName            respjson.Field
		IconURL                respjson.Field
		InstitutionName        respjson.Field
		Source                 respjson.Field
		SourceLabel            respjson.Field
		Subtitle               respjson.Field
		raw                    string
	} `json:"-"`
}

func (r *PaperTradingAccountNewResponseDataDetailsV1UnionAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseDataDetailsV1Object struct {
	Account              PaperTradingAccountNewResponseDataDetailsV1ObjectAccount `json:"account" api:"required"`
	AmountLabel          string                                                   `json:"amount_label" api:"required"`
	BottomMarkdown       string                                                   `json:"bottom_markdown" api:"required"`
	CurrentPriceDisplay  string                                                   `json:"current_price_display" api:"required"`
	EstimatedCostDisplay string                                                   `json:"estimated_cost_display" api:"required"`
	// Any of "shares", "dollars".
	InputMode string `json:"input_mode" api:"required"`
	// Any of "paper_trade".
	Kind                    string `json:"kind" api:"required"`
	MoneyMovementDisclosure string `json:"money_movement_disclosure" api:"required"`
	OrderTypeDisplay        string `json:"order_type_display" api:"required"`
	// Any of "positive", "negative", "neutral".
	PriceChangeDirection string    `json:"price_change_direction" api:"required"`
	PriceChangeDisplay   string    `json:"price_change_display" api:"required"`
	PrimaryActionLabel   string    `json:"primary_action_label" api:"required"`
	ProposedByLabel      string    `json:"proposed_by_label" api:"required"`
	Quantity             string    `json:"quantity" api:"required"`
	QuantityDisplay      string    `json:"quantity_display" api:"required"`
	QuoteAsOf            time.Time `json:"quote_as_of" api:"required" format:"date-time"`
	QuoteAsOfLabel       string    `json:"quote_as_of_label" api:"required"`
	ReasonMarkdown       string    `json:"reason_markdown" api:"required"`
	SecondaryActionLabel string    `json:"secondary_action_label" api:"required"`
	SecurityName         string    `json:"security_name" api:"required"`
	// Any of "buy", "sell".
	Side             string `json:"side" api:"required"`
	SideLabel        string `json:"side_label" api:"required"`
	Symbol           string `json:"symbol" api:"required"`
	TimingDisclosure string `json:"timing_disclosure" api:"required"`
	TopMarkdown      string `json:"top_markdown" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account                 respjson.Field
		AmountLabel             respjson.Field
		BottomMarkdown          respjson.Field
		CurrentPriceDisplay     respjson.Field
		EstimatedCostDisplay    respjson.Field
		InputMode               respjson.Field
		Kind                    respjson.Field
		MoneyMovementDisclosure respjson.Field
		OrderTypeDisplay        respjson.Field
		PriceChangeDirection    respjson.Field
		PriceChangeDisplay      respjson.Field
		PrimaryActionLabel      respjson.Field
		ProposedByLabel         respjson.Field
		Quantity                respjson.Field
		QuantityDisplay         respjson.Field
		QuoteAsOf               respjson.Field
		QuoteAsOfLabel          respjson.Field
		ReasonMarkdown          respjson.Field
		SecondaryActionLabel    respjson.Field
		SecurityName            respjson.Field
		Side                    respjson.Field
		SideLabel               respjson.Field
		Symbol                  respjson.Field
		TimingDisclosure        respjson.Field
		TopMarkdown             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseDataDetailsV1Object) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseDataDetailsV1Object) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseDataDetailsV1ObjectAccount struct {
	ID                     string `json:"id" api:"required"`
	AvailableAmountDisplay string `json:"available_amount_display" api:"required"`
	AvailableAmountLabel   string `json:"available_amount_label" api:"required"`
	// Any of "available", "unavailable".
	AvailableAmountState string `json:"available_amount_state" api:"required"`
	DisplayName          string `json:"display_name" api:"required"`
	IconURL              string `json:"icon_url" api:"required"`
	InstitutionName      string `json:"institution_name" api:"required"`
	// Any of "paper", "apex".
	Source      string `json:"source" api:"required"`
	SourceLabel string `json:"source_label" api:"required"`
	Subtitle    string `json:"subtitle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AvailableAmountDisplay respjson.Field
		AvailableAmountLabel   respjson.Field
		AvailableAmountState   respjson.Field
		DisplayName            respjson.Field
		IconURL                respjson.Field
		InstitutionName        respjson.Field
		Source                 respjson.Field
		SourceLabel            respjson.Field
		Subtitle               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseDataDetailsV1ObjectAccount) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseDataDetailsV1ObjectAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseDataDetailsV1Object2 struct {
	Account              PaperTradingAccountNewResponseDataDetailsV1Object2Account `json:"account" api:"required"`
	AmountLabel          string                                                    `json:"amount_label" api:"required"`
	BottomMarkdown       string                                                    `json:"bottom_markdown" api:"required"`
	CurrentPriceDisplay  string                                                    `json:"current_price_display" api:"required"`
	EstimatedCostDisplay string                                                    `json:"estimated_cost_display" api:"required"`
	// Any of "shares", "dollars".
	InputMode string `json:"input_mode" api:"required"`
	// Any of "trade".
	Kind                    string `json:"kind" api:"required"`
	MoneyMovementDisclosure string `json:"money_movement_disclosure" api:"required"`
	OrderTypeDisplay        string `json:"order_type_display" api:"required"`
	// Any of "positive", "negative", "neutral".
	PriceChangeDirection string    `json:"price_change_direction" api:"required"`
	PriceChangeDisplay   string    `json:"price_change_display" api:"required"`
	PrimaryActionLabel   string    `json:"primary_action_label" api:"required"`
	ProposedByLabel      string    `json:"proposed_by_label" api:"required"`
	Quantity             string    `json:"quantity" api:"required"`
	QuantityDisplay      string    `json:"quantity_display" api:"required"`
	QuoteAsOf            time.Time `json:"quote_as_of" api:"required" format:"date-time"`
	QuoteAsOfLabel       string    `json:"quote_as_of_label" api:"required"`
	ReasonMarkdown       string    `json:"reason_markdown" api:"required"`
	SecondaryActionLabel string    `json:"secondary_action_label" api:"required"`
	SecurityName         string    `json:"security_name" api:"required"`
	// Any of "buy", "sell".
	Side             string `json:"side" api:"required"`
	SideLabel        string `json:"side_label" api:"required"`
	Symbol           string `json:"symbol" api:"required"`
	TimingDisclosure string `json:"timing_disclosure" api:"required"`
	TopMarkdown      string `json:"top_markdown" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account                 respjson.Field
		AmountLabel             respjson.Field
		BottomMarkdown          respjson.Field
		CurrentPriceDisplay     respjson.Field
		EstimatedCostDisplay    respjson.Field
		InputMode               respjson.Field
		Kind                    respjson.Field
		MoneyMovementDisclosure respjson.Field
		OrderTypeDisplay        respjson.Field
		PriceChangeDirection    respjson.Field
		PriceChangeDisplay      respjson.Field
		PrimaryActionLabel      respjson.Field
		ProposedByLabel         respjson.Field
		Quantity                respjson.Field
		QuantityDisplay         respjson.Field
		QuoteAsOf               respjson.Field
		QuoteAsOfLabel          respjson.Field
		ReasonMarkdown          respjson.Field
		SecondaryActionLabel    respjson.Field
		SecurityName            respjson.Field
		Side                    respjson.Field
		SideLabel               respjson.Field
		Symbol                  respjson.Field
		TimingDisclosure        respjson.Field
		TopMarkdown             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseDataDetailsV1Object2) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseDataDetailsV1Object2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseDataDetailsV1Object2Account struct {
	ID                     string `json:"id" api:"required"`
	AvailableAmountDisplay string `json:"available_amount_display" api:"required"`
	AvailableAmountLabel   string `json:"available_amount_label" api:"required"`
	// Any of "available", "unavailable".
	AvailableAmountState string `json:"available_amount_state" api:"required"`
	DisplayName          string `json:"display_name" api:"required"`
	IconURL              string `json:"icon_url" api:"required"`
	InstitutionName      string `json:"institution_name" api:"required"`
	// Any of "paper", "apex".
	Source      string `json:"source" api:"required"`
	SourceLabel string `json:"source_label" api:"required"`
	Subtitle    string `json:"subtitle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AvailableAmountDisplay respjson.Field
		AvailableAmountLabel   respjson.Field
		AvailableAmountState   respjson.Field
		DisplayName            respjson.Field
		IconURL                respjson.Field
		InstitutionName        respjson.Field
		Source                 respjson.Field
		SourceLabel            respjson.Field
		Subtitle               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseDataDetailsV1Object2Account) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperTradingAccountNewResponseDataDetailsV1Object2Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseDataDetailsV1Object3 struct {
	Amount        string                                                 `json:"amount" api:"required"`
	AmountDisplay string                                                 `json:"amount_display" api:"required"`
	AmountLabel   string                                                 `json:"amount_label" api:"required"`
	CurrencyCode  string                                                 `json:"currency_code" api:"required"`
	Description   string                                                 `json:"description" api:"required"`
	From          PaperTradingAccountNewResponseDataDetailsV1Object3From `json:"from" api:"required"`
	// Any of "transfer".
	Kind string `json:"kind" api:"required"`
	// Any of "bank_transfer", "instant", "wire".
	Method                  string                                               `json:"method" api:"required"`
	MethodDisplay           string                                               `json:"method_display" api:"required"`
	MoneyMovementDisclosure string                                               `json:"money_movement_disclosure" api:"required"`
	PrimaryActionLabel      string                                               `json:"primary_action_label" api:"required"`
	RequestID               string                                               `json:"request_id" api:"required"`
	SecondaryActionLabel    string                                               `json:"secondary_action_label" api:"required"`
	TimingDisclosure        string                                               `json:"timing_disclosure" api:"required"`
	TimingLabel             string                                               `json:"timing_label" api:"required"`
	To                      PaperTradingAccountNewResponseDataDetailsV1Object3To `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                  respjson.Field
		AmountDisplay           respjson.Field
		AmountLabel             respjson.Field
		CurrencyCode            respjson.Field
		Description             respjson.Field
		From                    respjson.Field
		Kind                    respjson.Field
		Method                  respjson.Field
		MethodDisplay           respjson.Field
		MoneyMovementDisclosure respjson.Field
		PrimaryActionLabel      respjson.Field
		RequestID               respjson.Field
		SecondaryActionLabel    respjson.Field
		TimingDisclosure        respjson.Field
		TimingLabel             respjson.Field
		To                      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseDataDetailsV1Object3) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseDataDetailsV1Object3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseDataDetailsV1Object3From struct {
	ID                      string    `json:"id" api:"required"`
	BalanceAsOf             time.Time `json:"balance_as_of" api:"required" format:"date-time"`
	BalanceAvailableDisplay string    `json:"balance_available_display" api:"required"`
	BalanceCurrentDisplay   string    `json:"balance_current_display" api:"required"`
	DisplayName             string    `json:"display_name" api:"required"`
	IconURL                 string    `json:"icon_url" api:"required"`
	InstitutionName         string    `json:"institution_name" api:"required"`
	// Any of "yoshi_account", "plaid_account", "external_counterparty".
	Kind     string `json:"kind" api:"required"`
	Subtitle string `json:"subtitle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		BalanceAsOf             respjson.Field
		BalanceAvailableDisplay respjson.Field
		BalanceCurrentDisplay   respjson.Field
		DisplayName             respjson.Field
		IconURL                 respjson.Field
		InstitutionName         respjson.Field
		Kind                    respjson.Field
		Subtitle                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseDataDetailsV1Object3From) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseDataDetailsV1Object3From) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseDataDetailsV1Object3To struct {
	ID                      string    `json:"id" api:"required"`
	BalanceAsOf             time.Time `json:"balance_as_of" api:"required" format:"date-time"`
	BalanceAvailableDisplay string    `json:"balance_available_display" api:"required"`
	BalanceCurrentDisplay   string    `json:"balance_current_display" api:"required"`
	DisplayName             string    `json:"display_name" api:"required"`
	IconURL                 string    `json:"icon_url" api:"required"`
	InstitutionName         string    `json:"institution_name" api:"required"`
	// Any of "yoshi_account", "plaid_account", "external_counterparty".
	Kind     string `json:"kind" api:"required"`
	Subtitle string `json:"subtitle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		BalanceAsOf             respjson.Field
		BalanceAvailableDisplay respjson.Field
		BalanceCurrentDisplay   respjson.Field
		DisplayName             respjson.Field
		IconURL                 respjson.Field
		InstitutionName         respjson.Field
		Kind                    respjson.Field
		Subtitle                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountNewResponseDataDetailsV1Object3To) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseDataDetailsV1Object3To) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewResponseMeta struct {
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
func (r PaperTradingAccountNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountListResponse struct {
	Data PaperTradingAccountListResponseData `json:"data" api:"required"`
	Meta PaperTradingAccountListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountListResponseData struct {
	Accounts []PaperTradingAccountListResponseDataAccount `json:"accounts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountListResponseDataAccount struct {
	ID               string  `json:"id" api:"required"`
	AsOf             string  `json:"as_of" api:"required"`
	BalanceAvailable float64 `json:"balance_available" api:"required"`
	BalanceCurrent   float64 `json:"balance_current" api:"required"`
	CreatedAt        string  `json:"created_at" api:"required"`
	Name             string  `json:"name" api:"required"`
	Status           string  `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AsOf             respjson.Field
		BalanceAvailable respjson.Field
		BalanceCurrent   respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountListResponseDataAccount) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountListResponseDataAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountListResponseMeta struct {
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
func (r PaperTradingAccountListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountNewParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// Skip the approval flow and execute immediately. Not supported yet — reserved for
	// future use.
	SkipApproval        param.Opt[bool]   `json:"skip_approval,omitzero"`
	StartingCashBalance param.Opt[string] `json:"starting_cash_balance,omitzero"`
	paramObj
}

func (r PaperTradingAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperTradingAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperTradingAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
