package endpoint

type RequestInfo struct {
	MethodName string `json:"method_name"`
}

func (r *RequestInfo) String() string {
	return "RequestInfo " + toJSON(r)
}

type Request struct {
	Data interface{} `json:"data"`
}

func (r *Request) String() string {
	return "Request " + toJSON(r)
}
