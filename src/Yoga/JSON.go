package Yoga_JSON

import (
	"encoding/json"
	"fmt"
)

func _Undefined() interface{} {
	return nil
}

func _ParseJSON(payload string) interface{} {
	var v interface{}
	err := json.Unmarshal([]byte(payload), &v)
	if err != nil {
		panic(fmt.Sprintf("JSON parsing error: %v", err))
	}
	return v
}

func _UnsafeStringify(unboxFn func(interface{}) interface{}) interface{} {
	return func(data interface{}) interface{} {
		b, err := json.Marshal(unboxFn(data))
		if err != nil {
			panic(fmt.Sprintf("JSON stringify error: %v", err))
		}
		return string(b)
	}
}

func _UnsafePrettyStringify(unboxFn func(interface{}) interface{}) interface{} {
	return func(spaces int, data interface{}) interface{} {
		indent := ""
		for i := 0; i < spaces; i++ {
			indent += " "
		}
		b, err := json.MarshalIndent(unboxFn(data), "", indent)
		if err != nil {
			panic(fmt.Sprintf("JSON pretty stringify error: %v", err))
		}
		return string(b)
	}
}
