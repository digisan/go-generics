package v2

import (
	"reflect"
)

func KindOf(v any) string {
	return reflect.ValueOf(v).Kind().String()
}

// slowest
// func TypeOf1(v any) string {
// 	return fmt.Sprintf("%T", v)
// }

// 2nd fast
// func TypeOf(v any) string {
// 	return reflect.TypeOf(v).String()
// }

// fastest, but cannot cover all, e.g. pointer type
func TypeOf(v any) string {
	switch v.(type) {

	case string:
		return "string"
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"

	case uint:
		return "uint"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case float32:
		return "float32"
	case complex64:
		return "complex64"
	case complex128:
		return "complex128"

	default:
		return reflect.TypeOf(v).String()
	}
}
