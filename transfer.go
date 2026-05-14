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

// TransferService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferService] method instead.
type TransferService struct {
	options []option.RequestOption
}

// NewTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransferService(opts ...option.RequestOption) (r TransferService) {
	r = TransferService{}
	r.options = opts
	return
}

// Create an approval-backed bank or instant transfer.
func (r *TransferService) New(ctx context.Context, body TransferNewParams, opts ...option.RequestOption) (res *TransferNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TransferNewResponse struct {
	Data TransferNewResponseData `json:"data" api:"required"`
	Meta TransferNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseData struct {
	ActionID string `json:"action_id" api:"required" format:"uuid"`
	// Any of "trade", "transfer", "data_fix", "automation", "account_create",
	// "unknown".
	ActionKind          string                                `json:"action_kind" api:"required"`
	ActionType          string                                `json:"action_type" api:"required"`
	ApprovalStatusURL   string                                `json:"approval_status_url" api:"required"`
	ApprovalURL         string                                `json:"approval_url" api:"required"`
	CreatedAt           string                                `json:"created_at" api:"required"`
	DecidedAt           string                                `json:"decided_at" api:"required"`
	DecisionNeeds       TransferNewResponseDataDecisionNeeds  `json:"decision_needs" api:"required"`
	Description         string                                `json:"description" api:"required"`
	DetailsV1           TransferNewResponseDataDetailsV1Union `json:"details_v1" api:"required"`
	DisplayErrorMessage string                                `json:"display_error_message" api:"required"`
	ExecutedAt          string                                `json:"executed_at" api:"required"`
	ExpiresAt           string                                `json:"expires_at" api:"required"`
	Status              string                                `json:"status" api:"required"`
	ThreadID            string                                `json:"thread_id" api:"required" format:"uuid"`
	UpdatedAt           string                                `json:"updated_at" api:"required"`
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
func (r TransferNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseDataDecisionNeeds struct {
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
func (r TransferNewResponseDataDecisionNeeds) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseDataDecisionNeeds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransferNewResponseDataDetailsV1Union contains all possible properties and
// values from [TransferNewResponseDataDetailsV1Object],
// [TransferNewResponseDataDetailsV1Object2],
// [TransferNewResponseDataDetailsV1Object3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TransferNewResponseDataDetailsV1Union struct {
	// This field is a union of [TransferNewResponseDataDetailsV1ObjectAccount],
	// [TransferNewResponseDataDetailsV1Object2Account]
	Account                 TransferNewResponseDataDetailsV1UnionAccount `json:"account"`
	AmountLabel             string                                       `json:"amount_label"`
	BottomMarkdown          string                                       `json:"bottom_markdown"`
	CurrentPriceDisplay     string                                       `json:"current_price_display"`
	EstimatedCostDisplay    string                                       `json:"estimated_cost_display"`
	InputMode               string                                       `json:"input_mode"`
	Kind                    string                                       `json:"kind"`
	MoneyMovementDisclosure string                                       `json:"money_movement_disclosure"`
	OrderTypeDisplay        string                                       `json:"order_type_display"`
	PriceChangeDirection    string                                       `json:"price_change_direction"`
	PriceChangeDisplay      string                                       `json:"price_change_display"`
	PrimaryActionLabel      string                                       `json:"primary_action_label"`
	ProposedByLabel         string                                       `json:"proposed_by_label"`
	Quantity                string                                       `json:"quantity"`
	QuantityDisplay         string                                       `json:"quantity_display"`
	QuoteAsOf               time.Time                                    `json:"quote_as_of"`
	QuoteAsOfLabel          string                                       `json:"quote_as_of_label"`
	ReasonMarkdown          string                                       `json:"reason_markdown"`
	SecondaryActionLabel    string                                       `json:"secondary_action_label"`
	SecurityName            string                                       `json:"security_name"`
	Side                    string                                       `json:"side"`
	SideLabel               string                                       `json:"side_label"`
	Symbol                  string                                       `json:"symbol"`
	TimingDisclosure        string                                       `json:"timing_disclosure"`
	TopMarkdown             string                                       `json:"top_markdown"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	Amount string `json:"amount"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	AmountDisplay string `json:"amount_display"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	CurrencyCode string `json:"currency_code"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	Description string `json:"description"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	From TransferNewResponseDataDetailsV1Object3From `json:"from"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	Method string `json:"method"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	MethodDisplay string `json:"method_display"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	RequestID string `json:"request_id"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	TimingLabel string `json:"timing_label"`
	// This field is from variant [TransferNewResponseDataDetailsV1Object3].
	To   TransferNewResponseDataDetailsV1Object3To `json:"to"`
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

func (u TransferNewResponseDataDetailsV1Union) AsTransferNewResponseDataDetailsV1Object() (v TransferNewResponseDataDetailsV1Object) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransferNewResponseDataDetailsV1Union) AsTransferNewResponseDataDetailsV1Object2() (v TransferNewResponseDataDetailsV1Object2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransferNewResponseDataDetailsV1Union) AsTransferNewResponseDataDetailsV1Object3() (v TransferNewResponseDataDetailsV1Object3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransferNewResponseDataDetailsV1Union) RawJSON() string { return u.JSON.raw }

func (r *TransferNewResponseDataDetailsV1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransferNewResponseDataDetailsV1UnionAccount is an implicit subunion of
// [TransferNewResponseDataDetailsV1Union].
// TransferNewResponseDataDetailsV1UnionAccount provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TransferNewResponseDataDetailsV1Union].
type TransferNewResponseDataDetailsV1UnionAccount struct {
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

func (r *TransferNewResponseDataDetailsV1UnionAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseDataDetailsV1Object struct {
	Account              TransferNewResponseDataDetailsV1ObjectAccount `json:"account" api:"required"`
	AmountLabel          string                                        `json:"amount_label" api:"required"`
	BottomMarkdown       string                                        `json:"bottom_markdown" api:"required"`
	CurrentPriceDisplay  string                                        `json:"current_price_display" api:"required"`
	EstimatedCostDisplay string                                        `json:"estimated_cost_display" api:"required"`
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
func (r TransferNewResponseDataDetailsV1Object) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseDataDetailsV1Object) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseDataDetailsV1ObjectAccount struct {
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
func (r TransferNewResponseDataDetailsV1ObjectAccount) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseDataDetailsV1ObjectAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseDataDetailsV1Object2 struct {
	Account              TransferNewResponseDataDetailsV1Object2Account `json:"account" api:"required"`
	AmountLabel          string                                         `json:"amount_label" api:"required"`
	BottomMarkdown       string                                         `json:"bottom_markdown" api:"required"`
	CurrentPriceDisplay  string                                         `json:"current_price_display" api:"required"`
	EstimatedCostDisplay string                                         `json:"estimated_cost_display" api:"required"`
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
func (r TransferNewResponseDataDetailsV1Object2) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseDataDetailsV1Object2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseDataDetailsV1Object2Account struct {
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
func (r TransferNewResponseDataDetailsV1Object2Account) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseDataDetailsV1Object2Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseDataDetailsV1Object3 struct {
	Amount        string                                      `json:"amount" api:"required"`
	AmountDisplay string                                      `json:"amount_display" api:"required"`
	AmountLabel   string                                      `json:"amount_label" api:"required"`
	CurrencyCode  string                                      `json:"currency_code" api:"required"`
	Description   string                                      `json:"description" api:"required"`
	From          TransferNewResponseDataDetailsV1Object3From `json:"from" api:"required"`
	// Any of "transfer".
	Kind string `json:"kind" api:"required"`
	// Any of "bank_transfer", "instant".
	Method                  string                                    `json:"method" api:"required"`
	MethodDisplay           string                                    `json:"method_display" api:"required"`
	MoneyMovementDisclosure string                                    `json:"money_movement_disclosure" api:"required"`
	PrimaryActionLabel      string                                    `json:"primary_action_label" api:"required"`
	RequestID               string                                    `json:"request_id" api:"required"`
	SecondaryActionLabel    string                                    `json:"secondary_action_label" api:"required"`
	TimingDisclosure        string                                    `json:"timing_disclosure" api:"required"`
	TimingLabel             string                                    `json:"timing_label" api:"required"`
	To                      TransferNewResponseDataDetailsV1Object3To `json:"to" api:"required"`
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
func (r TransferNewResponseDataDetailsV1Object3) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseDataDetailsV1Object3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseDataDetailsV1Object3From struct {
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
func (r TransferNewResponseDataDetailsV1Object3From) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseDataDetailsV1Object3From) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseDataDetailsV1Object3To struct {
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
func (r TransferNewResponseDataDetailsV1Object3To) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseDataDetailsV1Object3To) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewResponseMeta struct {
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
func (r TransferNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewParams struct {
	Amount TransferNewParamsAmountUnion `json:"amount,omitzero" api:"required"`
	FromID string                       `json:"from_id" api:"required"`
	// Any of "bank_transfer", "instant".
	Method      TransferNewParamsMethod `json:"method,omitzero" api:"required"`
	ToID        string                  `json:"to_id" api:"required"`
	Description param.Opt[string]       `json:"description,omitzero"`
	RequestID   param.Opt[string]       `json:"request_id,omitzero" format:"uuid"`
	// Any of "USD".
	CurrencyCode TransferNewParamsCurrencyCode `json:"currency_code,omitzero"`
	paramObj
}

func (r TransferNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TransferNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransferNewParamsAmountUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransferNewParamsAmountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransferNewParamsAmountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type TransferNewParamsMethod string

const (
	TransferNewParamsMethodBankTransfer TransferNewParamsMethod = "bank_transfer"
	TransferNewParamsMethodInstant      TransferNewParamsMethod = "instant"
)

type TransferNewParamsCurrencyCode string

const (
	TransferNewParamsCurrencyCodeUsd TransferNewParamsCurrencyCode = "USD"
)
