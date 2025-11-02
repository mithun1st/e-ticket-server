package utils

import (
	"fmt"
	"strings"
)

func JoinArray[T any](arr []T) string {
	var strArr []string = make([]string, len(arr))
	for index, element := range arr {
		switch any(element).(type) {
		case string:
			strArr[index] = fmt.Sprintf("'%v'", element)
		default:
			strArr[index] = fmt.Sprint(element)
		}
	}
	return strings.Join(strArr, ",")
}
