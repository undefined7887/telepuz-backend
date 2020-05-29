package endpoint

import "encoding/json"

type RequestInfo struct {
	MethodName string `json:"method_name"`
}

type Request struct {
	Data interface{} `json:"data"`
}

func (r *Request) String() string {
	bytes, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	return "Request " + string(bytes)
}
