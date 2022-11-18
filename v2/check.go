package v2

import "reflect"

func IsArrayOrSlice(v any) bool {
	return In(reflect.TypeOf(v).Kind(), reflect.Slice, reflect.Array)
}
