package Yoga_JSON

import (
	"encoding/json"
	"fmt"
	"gopurs/output/gopurs_runtime"
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

func _UnsafeStringify(data interface{}) string {
	b, err := json.Marshal(deepUnbox(data))
	if err != nil {
		panic(fmt.Sprintf("JSON stringify error: %v", err))
	}
	return string(b)
}

func _UnsafePrettyStringify(spaces int, data interface{}) string {
	indent := ""
	for i := 0; i < spaces; i++ {
		indent += " "
	}
	b, err := json.MarshalIndent(deepUnbox(data), "", indent)
	if err != nil {
		panic(fmt.Sprintf("JSON pretty stringify error: %v", err))
	}
	return string(b)
}

func deepUnbox(v interface{}) interface{} {
	if val, ok := v.(gopurs_runtime.Value); ok {
		switch val.Type {
		case gopurs_runtime.TypeInt:
			return val.IntVal
		case gopurs_runtime.TypeFloat:
			return *(*float64)(val.UnsafePtr)
		case gopurs_runtime.TypeString:
			return *(*string)(val.UnsafePtr)
		case gopurs_runtime.TypeBool:
			return val.IntVal != 0
		case gopurs_runtime.TypeArray:
			arr := *(*[]gopurs_runtime.Value)(val.UnsafePtr)
			res := make([]interface{}, len(arr))
			for i, x := range arr {
				res[i] = deepUnbox(x)
			}
			return res
		case gopurs_runtime.TypeRecord:
			rec := *(*map[string]gopurs_runtime.Value)(val.UnsafePtr)
			res := make(map[string]interface{})
			for k, x := range rec {
				res[k] = deepUnbox(x)
			}
			return res
		case gopurs_runtime.TypeRecord0:
			return make(map[string]interface{})
		case gopurs_runtime.TypeRecord1:
			rec := *(*gopurs_runtime.RecordData1)(val.UnsafePtr)
			return map[string]interface{}{rec.K0: deepUnbox(rec.V0)}
		case gopurs_runtime.TypeRecord2:
			rec := *(*gopurs_runtime.RecordData2)(val.UnsafePtr)
			return map[string]interface{}{rec.K0: deepUnbox(rec.V0), rec.K1: deepUnbox(rec.V1)}
		case gopurs_runtime.TypeRecord3:
			rec := *(*gopurs_runtime.RecordData3)(val.UnsafePtr)
			return map[string]interface{}{rec.K0: deepUnbox(rec.V0), rec.K1: deepUnbox(rec.V1), rec.K2: deepUnbox(rec.V2)}
		case gopurs_runtime.TypeRecord4:
			rec := *(*gopurs_runtime.RecordData4)(val.UnsafePtr)
			return map[string]interface{}{rec.K0: deepUnbox(rec.V0), rec.K1: deepUnbox(rec.V1), rec.K2: deepUnbox(rec.V2), rec.K3: deepUnbox(rec.V3)}
		case gopurs_runtime.TypeRecord5:
			rec := *(*gopurs_runtime.RecordData5)(val.UnsafePtr)
			return map[string]interface{}{rec.K0: deepUnbox(rec.V0), rec.K1: deepUnbox(rec.V1), rec.K2: deepUnbox(rec.V2), rec.K3: deepUnbox(rec.V3), rec.K4: deepUnbox(rec.V4)}
		case gopurs_runtime.TypeRecordData:
			rec := *(*gopurs_runtime.RecordData)(val.UnsafePtr)
			res := make(map[string]interface{})
			for i, k := range rec.Keys {
				res[k] = deepUnbox(rec.Vals[i])
			}
			return res
		case gopurs_runtime.TypeAny:
			ptr := *(*any)(val.UnsafePtr)
			return deepUnbox(ptr)
		default:
			return nil
		}
	}
	if valSlice, ok := v.([]gopurs_runtime.Value); ok {
		res := make([]interface{}, len(valSlice))
		for i, x := range valSlice {
			res[i] = deepUnbox(x)
		}
		return res
	}
	return v
}
