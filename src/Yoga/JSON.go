package Yoga_JSON

import (
	"encoding/json"
	"fmt"
	"gopurs/output/Foreign"
	"gopurs/output/gopurs_runtime"
)

var _Undefined = gopurs_runtime.Value{Type: 0}

func _ParseJSON(payload string) interface{} {
	var v interface{}
	err := json.Unmarshal([]byte(payload), &v)
	if err != nil {
		panic(fmt.Sprintf("JSON parsing error: %v", err))
	}
	return v
}

func UnsafeStringify(data interface{}) interface{} {
	b, err := json.Marshal(Foreign.UnboxForJSON(data))
	if err != nil {
		panic(fmt.Sprintf("JSON stringify error: %v", err))
	}
	return string(b)
}

func _UnsafePrettyStringify(spaces int, data interface{}) interface{} {
	indent := ""
	for i := 0; i < spaces; i++ {
		indent += " "
	}
	b, err := json.MarshalIndent(Foreign.UnboxForJSON(data), "", indent)
	if err != nil {
		panic(fmt.Sprintf("JSON pretty stringify error: %v", err))
	}
	return string(b)
}
