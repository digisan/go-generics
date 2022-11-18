package v2

import (
	"fmt"
	"reflect"
	"strconv"
)

func IsArrayOrSlice(v any) bool {
	return In(reflect.TypeOf(v).Kind(), reflect.Slice, reflect.Array)
}

// IsUint : Check str is valid numeric style
func IsUint(v any) bool {
	_, err := strconv.ParseUint(fmt.Sprint(v), 10, 64)
	return err == nil
}
