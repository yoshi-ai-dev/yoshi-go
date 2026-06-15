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
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// PaperTradingAccountTradeService contains methods and other services that help
// with interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperTradingAccountTradeService] method instead.
type PaperTradingAccountTradeService struct {
	options []option.RequestOption
}

// NewPaperTradingAccountTradeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPaperTradingAccountTradeService(opts ...option.RequestOption) (r PaperTradingAccountTradeService) {
	r = PaperTradingAccountTradeService{}
	r.options = opts
	return
}

// Place a buy or sell trade. Requires user approval in the Yoshi web app before
// execution.
func (r *PaperTradingAccountTradeService) New(ctx context.Context, accountID string, body PaperTradingAccountTradeNewParams, opts ...option.RequestOption) (res *PaperTradingAccountTradeNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("paper-trading/accounts/%s/trades", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List trade history for a Test Drive account with cursor-based pagination.
func (r *PaperTradingAccountTradeService) List(ctx context.Context, accountID string, query PaperTradingAccountTradeListParams, opts ...option.RequestOption) (res *PaperTradingAccountTradeListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("paper-trading/accounts/%s/trades", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type PaperTradingAccountTradeNewResponse struct {
	Data PaperTradingAccountTradeNewResponseData `json:"data" api:"required"`
	Meta PaperTradingAccountTradeNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountTradeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseData struct {
	ActionID string `json:"action_id" api:"required" format:"uuid"`
	// Any of "trade", "transfer", "data_fix", "automation", "account_create",
	// "unknown".
	ActionKind          string                                                `json:"action_kind" api:"required"`
	ActionType          string                                                `json:"action_type" api:"required"`
	ApprovalStatusURL   string                                                `json:"approval_status_url" api:"required"`
	ApprovalURL         string                                                `json:"approval_url" api:"required"`
	CreatedAt           string                                                `json:"created_at" api:"required"`
	DecidedAt           string                                                `json:"decided_at" api:"required"`
	DecisionNeeds       PaperTradingAccountTradeNewResponseDataDecisionNeeds  `json:"decision_needs" api:"required"`
	Description         string                                                `json:"description" api:"required"`
	DetailsV1           PaperTradingAccountTradeNewResponseDataDetailsV1Union `json:"details_v1" api:"required"`
	DisplayErrorMessage string                                                `json:"display_error_message" api:"required"`
	ExecutedAt          string                                                `json:"executed_at" api:"required"`
	ExpiresAt           string                                                `json:"expires_at" api:"required"`
	Status              string                                                `json:"status" api:"required"`
	ThreadID            string                                                `json:"thread_id" api:"required" format:"uuid"`
	UpdatedAt           string                                                `json:"updated_at" api:"required"`
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
func (r PaperTradingAccountTradeNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseDataDecisionNeeds struct {
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
func (r PaperTradingAccountTradeNewResponseDataDecisionNeeds) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponseDataDecisionNeeds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperTradingAccountTradeNewResponseDataDetailsV1Union contains all possible
// properties and values from
// [PaperTradingAccountTradeNewResponseDataDetailsV1Object],
// [PaperTradingAccountTradeNewResponseDataDetailsV1Object2],
// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PaperTradingAccountTradeNewResponseDataDetailsV1Union struct {
	// This field is a union of
	// [PaperTradingAccountTradeNewResponseDataDetailsV1ObjectAccount],
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object2Account]
	Account                 PaperTradingAccountTradeNewResponseDataDetailsV1UnionAccount `json:"account"`
	AmountLabel             string                                                       `json:"amount_label"`
	BottomMarkdown          string                                                       `json:"bottom_markdown"`
	CurrentPriceDisplay     string                                                       `json:"current_price_display"`
	EstimatedCostDisplay    string                                                       `json:"estimated_cost_display"`
	InputMode               string                                                       `json:"input_mode"`
	Kind                    string                                                       `json:"kind"`
	MoneyMovementDisclosure string                                                       `json:"money_movement_disclosure"`
	OrderTypeDisplay        string                                                       `json:"order_type_display"`
	PriceChangeDirection    string                                                       `json:"price_change_direction"`
	PriceChangeDisplay      string                                                       `json:"price_change_display"`
	PrimaryActionLabel      string                                                       `json:"primary_action_label"`
	ProposedByLabel         string                                                       `json:"proposed_by_label"`
	Quantity                string                                                       `json:"quantity"`
	QuantityDisplay         string                                                       `json:"quantity_display"`
	QuoteAsOf               time.Time                                                    `json:"quote_as_of"`
	QuoteAsOfLabel          string                                                       `json:"quote_as_of_label"`
	ReasonMarkdown          string                                                       `json:"reason_markdown"`
	SecondaryActionLabel    string                                                       `json:"secondary_action_label"`
	SecurityName            string                                                       `json:"security_name"`
	Side                    string                                                       `json:"side"`
	SideLabel               string                                                       `json:"side_label"`
	Symbol                  string                                                       `json:"symbol"`
	TimingDisclosure        string                                                       `json:"timing_disclosure"`
	TopMarkdown             string                                                       `json:"top_markdown"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	Amount string `json:"amount"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	AmountDisplay string `json:"amount_display"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	CurrencyCode string `json:"currency_code"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	Description string `json:"description"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	From PaperTradingAccountTradeNewResponseDataDetailsV1Object3From `json:"from"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	Method string `json:"method"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	MethodDisplay string `json:"method_display"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	RequestID string `json:"request_id"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	TimingLabel string `json:"timing_label"`
	// This field is from variant
	// [PaperTradingAccountTradeNewResponseDataDetailsV1Object3].
	To   PaperTradingAccountTradeNewResponseDataDetailsV1Object3To `json:"to"`
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

func (u PaperTradingAccountTradeNewResponseDataDetailsV1Union) AsPaperTradingAccountTradeNewResponseDataDetailsV1Object() (v PaperTradingAccountTradeNewResponseDataDetailsV1Object) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperTradingAccountTradeNewResponseDataDetailsV1Union) AsPaperTradingAccountTradeNewResponseDataDetailsV1Object2() (v PaperTradingAccountTradeNewResponseDataDetailsV1Object2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaperTradingAccountTradeNewResponseDataDetailsV1Union) AsPaperTradingAccountTradeNewResponseDataDetailsV1Object3() (v PaperTradingAccountTradeNewResponseDataDetailsV1Object3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaperTradingAccountTradeNewResponseDataDetailsV1Union) RawJSON() string { return u.JSON.raw }

func (r *PaperTradingAccountTradeNewResponseDataDetailsV1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaperTradingAccountTradeNewResponseDataDetailsV1UnionAccount is an implicit
// subunion of [PaperTradingAccountTradeNewResponseDataDetailsV1Union].
// PaperTradingAccountTradeNewResponseDataDetailsV1UnionAccount provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PaperTradingAccountTradeNewResponseDataDetailsV1Union].
type PaperTradingAccountTradeNewResponseDataDetailsV1UnionAccount struct {
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

func (r *PaperTradingAccountTradeNewResponseDataDetailsV1UnionAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseDataDetailsV1Object struct {
	Account              PaperTradingAccountTradeNewResponseDataDetailsV1ObjectAccount `json:"account" api:"required"`
	AmountLabel          string                                                        `json:"amount_label" api:"required"`
	BottomMarkdown       string                                                        `json:"bottom_markdown" api:"required"`
	CurrentPriceDisplay  string                                                        `json:"current_price_display" api:"required"`
	EstimatedCostDisplay string                                                        `json:"estimated_cost_display" api:"required"`
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
func (r PaperTradingAccountTradeNewResponseDataDetailsV1Object) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponseDataDetailsV1Object) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseDataDetailsV1ObjectAccount struct {
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
func (r PaperTradingAccountTradeNewResponseDataDetailsV1ObjectAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperTradingAccountTradeNewResponseDataDetailsV1ObjectAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseDataDetailsV1Object2 struct {
	Account              PaperTradingAccountTradeNewResponseDataDetailsV1Object2Account `json:"account" api:"required"`
	AmountLabel          string                                                         `json:"amount_label" api:"required"`
	BottomMarkdown       string                                                         `json:"bottom_markdown" api:"required"`
	CurrentPriceDisplay  string                                                         `json:"current_price_display" api:"required"`
	EstimatedCostDisplay string                                                         `json:"estimated_cost_display" api:"required"`
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
func (r PaperTradingAccountTradeNewResponseDataDetailsV1Object2) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponseDataDetailsV1Object2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseDataDetailsV1Object2Account struct {
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
func (r PaperTradingAccountTradeNewResponseDataDetailsV1Object2Account) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperTradingAccountTradeNewResponseDataDetailsV1Object2Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseDataDetailsV1Object3 struct {
	Amount        string                                                      `json:"amount" api:"required"`
	AmountDisplay string                                                      `json:"amount_display" api:"required"`
	AmountLabel   string                                                      `json:"amount_label" api:"required"`
	CurrencyCode  string                                                      `json:"currency_code" api:"required"`
	Description   string                                                      `json:"description" api:"required"`
	From          PaperTradingAccountTradeNewResponseDataDetailsV1Object3From `json:"from" api:"required"`
	// Any of "transfer".
	Kind string `json:"kind" api:"required"`
	// Any of "bank_transfer", "instant", "wire".
	Method                  string                                                    `json:"method" api:"required"`
	MethodDisplay           string                                                    `json:"method_display" api:"required"`
	MoneyMovementDisclosure string                                                    `json:"money_movement_disclosure" api:"required"`
	PrimaryActionLabel      string                                                    `json:"primary_action_label" api:"required"`
	RequestID               string                                                    `json:"request_id" api:"required"`
	SecondaryActionLabel    string                                                    `json:"secondary_action_label" api:"required"`
	TimingDisclosure        string                                                    `json:"timing_disclosure" api:"required"`
	TimingLabel             string                                                    `json:"timing_label" api:"required"`
	To                      PaperTradingAccountTradeNewResponseDataDetailsV1Object3To `json:"to" api:"required"`
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
func (r PaperTradingAccountTradeNewResponseDataDetailsV1Object3) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponseDataDetailsV1Object3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseDataDetailsV1Object3From struct {
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
func (r PaperTradingAccountTradeNewResponseDataDetailsV1Object3From) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperTradingAccountTradeNewResponseDataDetailsV1Object3From) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseDataDetailsV1Object3To struct {
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
func (r PaperTradingAccountTradeNewResponseDataDetailsV1Object3To) RawJSON() string {
	return r.JSON.raw
}
func (r *PaperTradingAccountTradeNewResponseDataDetailsV1Object3To) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewResponseMeta struct {
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
func (r PaperTradingAccountTradeNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeListResponse struct {
	// Number of items in the current page
	Count int64 `json:"count" api:"required"`
	// Opaque cursor for the next page, null if no more pages
	Cursor  string                                     `json:"cursor" api:"required"`
	Data    []PaperTradingAccountTradeListResponseData `json:"data" api:"required"`
	HasMore bool                                       `json:"has_more" api:"required"`
	Meta    PaperTradingAccountTradeListResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Cursor      respjson.Field
		Data        respjson.Field
		HasMore     respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountTradeListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeListResponseData struct {
	ID        string  `json:"id" api:"required"`
	Amount    float64 `json:"amount" api:"required"`
	CreatedAt string  `json:"created_at" api:"required"`
	Date      string  `json:"date" api:"required"`
	Name      string  `json:"name" api:"required"`
	Price     float64 `json:"price" api:"required"`
	Quantity  float64 `json:"quantity" api:"required"`
	Subtype   string  `json:"subtype" api:"required"`
	Symbol    string  `json:"symbol" api:"required"`
	Type      string  `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Date        respjson.Field
		Name        respjson.Field
		Price       respjson.Field
		Quantity    respjson.Field
		Subtype     respjson.Field
		Symbol      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaperTradingAccountTradeListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeListResponseMeta struct {
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
func (r PaperTradingAccountTradeListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PaperTradingAccountTradeListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewParams struct {
	// Any of "buy", "sell".
	Side     PaperTradingAccountTradeNewParamsSide `json:"side,omitzero" api:"required"`
	Symbol   string                                `json:"symbol" api:"required"`
	Notional param.Opt[float64]                    `json:"notional,omitzero"`
	Quantity param.Opt[float64]                    `json:"quantity,omitzero"`
	// Skip the approval flow and execute immediately. Not supported yet — reserved for
	// future use.
	SkipApproval param.Opt[bool] `json:"skip_approval,omitzero"`
	paramObj
}

func (r PaperTradingAccountTradeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaperTradingAccountTradeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaperTradingAccountTradeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaperTradingAccountTradeNewParamsSide string

const (
	PaperTradingAccountTradeNewParamsSideBuy  PaperTradingAccountTradeNewParamsSide = "buy"
	PaperTradingAccountTradeNewParamsSideSell PaperTradingAccountTradeNewParamsSide = "sell"
)

type PaperTradingAccountTradeListParams struct {
	// Opaque cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Items per page (1-100, default 50)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaperTradingAccountTradeListParams]'s query parameters as
// `url.Values`.
func (r PaperTradingAccountTradeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
