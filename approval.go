// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// ApprovalService contains methods and other services that help with interacting
// with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApprovalService] method instead.
type ApprovalService struct {
	options []option.RequestOption
}

// NewApprovalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApprovalService(opts ...option.RequestOption) (r ApprovalService) {
	r = ApprovalService{}
	r.options = opts
	return
}

// Check the approval status of a pending action (paper trading account creation,
// trade execution, etc.).
func (r *ApprovalService) Get(ctx context.Context, threadID string, opts ...option.RequestOption) (res *ApprovalGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("approvals/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ApprovalGetResponse struct {
	Data ApprovalGetResponseData `json:"data" api:"required"`
	Meta ApprovalGetResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ApprovalGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalGetResponseData struct {
	ActionID    string `json:"action_id" api:"required" format:"uuid"`
	CreatedAt   string `json:"created_at" api:"required"`
	Description string `json:"description" api:"required"`
	Status      string `json:"status" api:"required"`
	ThreadID    string `json:"thread_id" api:"required" format:"uuid"`
	UpdatedAt   string `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Status      respjson.Field
		ThreadID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ApprovalGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalGetResponseMeta struct {
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
func (r ApprovalGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ApprovalGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
