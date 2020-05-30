package endpoint

type ResponseInfo struct {
	MethodName string `json:"method_name"`
}

func (r *ResponseInfo) String() string {
	return "ResponseInfo " + toJSON(r)
}

type Response struct {
	Result int         `json:"result"`
	Data   interface{} `json:"data,omitempty"`
}

func (r *Response) String() string {
	return "Response " + toJSON(r)
}
