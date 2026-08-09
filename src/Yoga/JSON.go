package Yoga_JSON

import (
	"encoding/json"
	"fmt"
	"gopurs/output/Foreign"
	"gopurs/output/gopurs_runtime"
	"math/big"
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

func transformForStringify(v interface{}) interface{} {
	if b, ok := v.(*big.Int); ok {
		return b.String()
	}
	if m, ok := v.(map[string]interface{}); ok {
		res := make(map[string]interface{})
		for k, val := range m {
			if val != Foreign.UndefinedForJSON {
				res[k] = transformForStringify(val)
			}
		}
		return res
	}
	if a, ok := v.([]interface{}); ok {
		res := make([]interface{}, len(a))
		for i, val := range a {
			res[i] = transformForStringify(val)
		}
		return res
	}
	return v
}

func UnsafeStringify(data interface{}) interface{} {
	b, err := json.Marshal(transformForStringify(Foreign.UnboxForJSON(data)))
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
	b, err := json.MarshalIndent(transformForStringify(Foreign.UnboxForJSON(data)), "", indent)
	if err != nil {
		panic(fmt.Sprintf("JSON pretty stringify error: %v", err))
	}
	return string(b)
}
