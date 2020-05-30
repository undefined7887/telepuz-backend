package endpoint

type Method interface {
	NewRequest() *Request
	Call(req *Request) *Response
}
