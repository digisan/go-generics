package gogenerics

import (
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
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

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func IsArrOrSlc(v any) bool {
	if v == nil {
		return false
	}
	return In(reflect.TypeOf(v).Kind(), reflect.Slice, reflect.Array)
}

func LenOfMustArrOrSlc(v any) int {
	if v == nil {
		log.Fatalf("[%v] is NOT Slice or Array Type", v)
	}
	return reflect.ValueOf(v).Len()
}

// IsInt : Check v is valid int numeric style
func IsInt(v any) bool {
	_, err := strconv.ParseInt(fmt.Sprint(v), 10, 64)
	return err == nil
}

// IsUint : Check v is valid uint numeric style
func IsUint(v any) bool {
	_, err := strconv.ParseUint(fmt.Sprint(v), 10, 64)
	return err == nil
}

// IsNumeric : Check v is valid numeric style
func IsNumeric(v any) bool {
	_, err := strconv.ParseFloat(fmt.Sprint(v), 64)
	return err == nil
}

// IsContinuous : check numbers is continuous int slice
func IsContinuous[T Integer](numbers ...T) (ok bool, rtMin T, rtMax T) {
	if len(numbers) == 0 {
		return false, *new(T), *new(T)
	}
	if len(numbers) == 1 {
		return true, numbers[0], numbers[0]
	}
	s, e := numbers[0], numbers[len(numbers)-1]
	if s < e {
		return reflect.DeepEqual(IterToSlc(s, e+1), numbers), s, e
	}
	return reflect.DeepEqual(IterToSlc(s, e-1), numbers), e, s
}

// nil-pointer could be non-nil any
func IsNil(i any) bool {
	return i == nil || reflect.ValueOf(i).Kind() == 0 || reflect.ValueOf(i).IsNil() || fmt.Sprint(i) == "<nil>"
}

// check string format is email
func IsEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

// check string is valid IP
func IsIP(s string) bool {
	return net.ParseIP(s) != nil
}

func IsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func IsDateTime(s string) bool {
	_, ok := TryToDateTime(s)
	return ok
}

func IsDateUS(s string) bool {
	_, ok := TryToDateUS(s)
	return ok
}

func IsDateUK(s string) bool {
	_, ok := TryToDateUK(s)
	return ok
}

func IsTime(s string) bool {
	_, ok := TryToTime(s)
	return ok
}
