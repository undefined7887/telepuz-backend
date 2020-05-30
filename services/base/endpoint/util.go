package endpoint

import "encoding/json"

func toJSON(value interface{}) string {
	bytes, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	return string(bytes)
}
