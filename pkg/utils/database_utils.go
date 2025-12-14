package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func _cast(data any) string {

	var nullStr string = "null"

	switch value := any(data).(type) {

	case nil:
		return nullStr
	// String
	case *string:
		if value == nil {
			return nullStr
		}
		return fmt.Sprintf("'%s'", *value)
	case string:
		return fmt.Sprintf("'%s'", value)

	// Numeric & Boolean
	case
		*uint, *uint8, *uint16, *uint32, *uint64,
		*int, *int8, *int16, *int32, *int64,
		*float32, *float64,
		*bool:

		rv := reflect.ValueOf(value)
		if rv.IsNil() {
			return nullStr
		}
		return fmt.Sprintf("%v", rv.Elem().Interface())

	default:
		return fmt.Sprintf("%v", value)
	}

}

func DbNames(name ...string) string {
	return strings.Join(name, ",")
}

func DbValues(values ...any) string {
	var arr []string = make([]string, 0)
	for _, e := range values {
		arr = append(arr, _cast(e))
	}
	return DbNames(arr...)
}
