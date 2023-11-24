package pagination

import (
	"net/http"
	"strconv"
)

const (
	DefaultPageSize int = 10
)

type (
	SortDirection string
)

type QueryPagination struct {
	Page          int
	PageSize      int
	Term          string
	SortColumn    string
	SortDirection SortDirection
	Links         *Links
}

func (q QueryPagination) OffsetFromPage() int {
	offset := (q.Page - 1) * q.PageSize
	return offset
}

func page(r *http.Request) int {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page <= 0 {
		page = 1
	}
	return page
}

func pageSize(r *http.Request) int {
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pagesize"))
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	return pageSize
}

func term(r *http.Request) string {
	term := r.URL.Query().Get("term")
	return term
}

func ParseQuery(r *http.Request) QueryPagination {
	qp := QueryPagination{
		Page:     page(r),
		PageSize: pageSize(r),
		Term:     term(r),
	}

	currentQuery := r.URL.Query()

	currentQuery.Set("page", strconv.Itoa(qp.Page-1))
	qp.Links.PrevPage = r.URL.Path + currentQuery.Encode()

	currentQuery.Set("page", strconv.Itoa(qp.Page+1))
	qp.Links.NextPage = r.URL.Path + currentQuery.Encode()
	return qp
}

type Links struct {
	NextPage string
	PrevPage string
	Pages    []string
}
