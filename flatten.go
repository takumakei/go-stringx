package stringx

import (
	"fmt"
	"reflect"
)

// Flatten returns a single slice of all elements from all slices or arrays of strings.
func Flatten(value ...interface{}) []string {
	s := make([]string, 0, len(value))
	for _, e := range value {
		switch v := e.(type) {
		case string:
			s = append(s, v)

		case []string:
			s = append(s, v...)

		default:
			s = flatten(s, e)
		}
	}
	return s
}

func flatten(s []string, a interface{}) []string {
	ty := reflect.TypeOf(a)
	if k := ty.Kind(); k == reflect.Slice || k == reflect.Array {
		switch ty.Elem().Kind() {
		case reflect.String:
			v := reflect.ValueOf(a)
			for i := 0; i < v.Len(); i++ {
				s = append(s, v.Index(i).String())
			}
			return s

		case reflect.Slice, reflect.Array:
			v := reflect.ValueOf(a)
			for i := 0; i < v.Len(); i++ {
				s = flatten(s, v.Index(i).Interface())
			}
			return s
		}
	}

	panic(fmt.Sprintf("%v is not a string nor a string slice", a))
}
