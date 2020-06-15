package utils

import "encoding/json"

func PrettyStruct(prefix string, value interface{}) string {
	bytes, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	return prefix + " " + string(bytes)
}
