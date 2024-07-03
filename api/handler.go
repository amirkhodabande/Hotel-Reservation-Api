package api

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"message"`
	Data    any   `json:"data"`
}

func SuccessResponse(data any) *Response {
	return &Response{
		Success: true,
		Msg:     "Operation finished successfully",
		Data:    data,
	}
}
