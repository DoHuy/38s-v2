package dtos

// BaseResponse represents struct of base response
type BaseResponse struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}

// Meta represents struct of meta
type Meta struct {
	Code int `json:"code,omitempty"`
}

// PagingMeta paging meta for listing API
type PagingMeta struct {
	Meta
	Limit      int64 `json:"limit"`
	Offset     int64 `json:"offset"`
	TotalItems int64 `json:"total"`
}

// ErrorItem represents struct of error item
type ErrorItem struct {
	Key     string `json:"key,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// ErrorData represent struct of error data
type ErrorData struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Counter int         `json:"counter"`
	Errors  []ErrorItem `json:"errors,omitempty"`
}