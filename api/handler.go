package api

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"message"`
	Data    any    `json:"data"`
	Count   int64  `json:"count,omitempty"`
	Page    int64  `json:"page,omitempty"`
}

func SuccessResponse(data any) *Response {
	return &Response{
		Success: true,
		Msg:     "Operation finished successfully",
		Data:    data,
	}
}

func (r *Response) WithPagination(count, page int64) *Response {
	r.Count = count
	r.Page = page
	return r
}
