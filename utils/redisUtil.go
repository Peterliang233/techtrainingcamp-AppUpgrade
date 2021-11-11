package utils

import "encoding/json"

// MapToString map转string字符串
func MapToString(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)

	return string(dataType)
}
