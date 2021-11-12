package utils

import (
	"encoding/json"
	"strconv"
	"strings"
)

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

// compareVersion 比较两个版本的高低,version1>version2，返回1,version1<version2,返回-1,否则返回0
func CompareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) || i < len(v2); i++ {
		x, y := 0, 0
		if i < len(v1) {
			x, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			y, _ = strconv.Atoi(v2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
