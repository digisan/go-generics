package v2

import (
	"fmt"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
	"time"
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
	return In(reflect.TypeOf(v).Kind(), reflect.Slice, reflect.Array)
}

func LenOfMustArrOrSlc(v any) int {
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
func IsEmail(str string) bool {
	_, err := mail.ParseAddress(str)
	return err == nil
}

func IsURL(str string) bool {
	_, err := url.ParseRequestURI(str)
	return err == nil
}

func IsDateTime(str string) bool {
	var layouts = []string{
		// standard
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		// Handy time stamps.
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
	for _, lo := range layouts {
		if _, err := time.Parse(lo, str); err == nil {
			return true
		}
	}
	return false
}

func IsDateUS(str string) bool {
	var layouts = []string{
		"January 2, 2006",
		"Jan 2, 2006",
		"01/02/06",
		"01/02/2006",
		"Jan-02-06",
	}
	for _, lo := range layouts {
		if _, err := time.Parse(lo, str); err == nil {
			return true
		}
	}
	return false
}

func IsDateUK(str string) bool {
	var layouts = []string{
		"2 January, 2006",
		"2 Jan, 2006",
		"02/01/06",
		"02/01/2006",
		"02-Jan-06",
	}
	for _, lo := range layouts {
		if _, err := time.Parse(lo, str); err == nil {
			return true
		}
	}
	return false
}

func IsTime(str string) bool {
	var layouts = []string{
		"15:04:05",
		"3:04:05PM",
		"3:04:05 PM",
		"3:04:05pm",
		"3:04:05 pm",
		"3:04:05 P.M.",
		"3:04:05 p.m.",
	}
	for _, lo := range layouts {
		if _, err := time.Parse(lo, str); err == nil {
			return true
		}
	}
	return false
}
