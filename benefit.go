// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"github.com/yoshi-ai-dev/yoshi-go/option"
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
