package common

import (
	"encoding/json"
	"reflect"
)

// StructToMap 结构体转map
func StructToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	b, err := json.Marshal(s)
	if err != nil {
		return m
	}
	_ = json.Unmarshal(b, &m)
	return m
}

// RemoveZeroValues 去掉零值 跟 id
func RemoveZeroValues(m map[string]interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})
	for k, v := range m {
		if k == "id" {
			continue
		}
		if !isZero(reflect.ValueOf(v)) {
			filtered[k] = v
		}
	}
	return filtered
}

func isZero(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
