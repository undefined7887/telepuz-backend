package endpoint

import "encoding/json"

type ResponseInfo struct {
	MethodName string `json:"method_name"`
}

type Response struct {
	Result int         `json:"result"`
	Data   interface{} `json:"data,omitempty"`
}

func (r *Response) String() string {
	bytes, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	return "Response " + string(bytes)
}
