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

type Query struct {
	Page          int
	PageSize      int
	Term          string
	SortColumn    string
	SortDirection SortDirection
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

func ParseQuery(r *http.Request) Query {
	return Query{
		Page:     page(r),
		PageSize: pageSize(r),
		Term:     term(r),
	}
}
