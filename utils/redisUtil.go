package utils

import "encoding/json"

// MapToString map转string字符串
func MapToString(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)

	return string(dataType)
}

// StringToMap 将String转Map
func StringToMap(s string) map[string]interface{} {
	var temp map[string]interface{}

	_ = json.Unmarshal([]byte(s), &temp)

	return temp
}
