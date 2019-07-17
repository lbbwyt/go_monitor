package common

import (
	"encoding/json"
)

func Obj2JsonStr(v interface{} ) (string) {
	jData, err:= json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(jData[:])
}
