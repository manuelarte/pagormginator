package model

// UnPaged No pagination
var UnPaged = PageRequest{0, 0}

// PageRequest Struct that contains the pagination information
type PageRequest struct {
	Page int
	Size int
}

// Page Struct that contains the result of the query
type Page[T any] struct {
	TotalPages    int
	Page          int
	Size          int
	TotalElements int
	Data          []*T
}
