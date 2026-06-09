// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"iter"
	"net/http"

	"github.com/yoshi-ai-dev/yoshi-go/internal/apijson"
	"github.com/yoshi-ai-dev/yoshi-go/internal/requestconfig"
	"github.com/yoshi-ai-dev/yoshi-go/option"
	"github.com/yoshi-ai-dev/yoshi-go/packages/param"
	"github.com/yoshi-ai-dev/yoshi-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CursorPagePagination struct {
	Count      int64  `json:"count"`
	HasMore    bool   `json:"has_more"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CursorPagePagination) RawJSON() string { return r.JSON.raw }
func (r *CursorPagePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CursorPage[T any] struct {
	Data       []T                  `json:"data"`
	Pagination CursorPagePagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageAutoPager[T]) Index() int {
	return r.run
}

func (r *CursorPageAutoPager[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for r.Next() {
			if !yield(r.Current()) {
				break
			}
		}
	}
}
