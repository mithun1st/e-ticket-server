package utils

import "fmt"

func NilToStrDB[T any](data *T) string {
	switch any(data).(type) {
	case *string:
		if data == nil {
			return "null"
		} else {
			return fmt.Sprintf("'%v'", *data)
		}
	default:
		if data == nil {
			return "null"
		} else {
			return fmt.Sprintf("%v", *data)
		}
	}

}
