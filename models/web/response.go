package web

type Response struct {
	StatusCode responseCode `json:"status_code"`
	Message    string       `json:"message"`
	Data       interface{}  `json:"data"`
}

func NewResponse() *Response {
	return &Response{
		StatusCode: RcSuccess,
		Message:    "success",
		Data:       nil,
	}
}
