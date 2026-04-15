// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package yoshi

import (
	"github.com/yoshi-ai-dev/yoshi-go/option"
)

// PaperTradingService contains methods and other services that help with
// interacting with the yoshi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperTradingService] method instead.
type PaperTradingService struct {
	options  []option.RequestOption
	Accounts PaperTradingAccountService
}

// NewPaperTradingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperTradingService(opts ...option.RequestOption) (r PaperTradingService) {
	r = PaperTradingService{}
	r.options = opts
	r.Accounts = NewPaperTradingAccountService(opts...)
	return
}
