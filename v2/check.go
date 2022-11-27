package v2

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func IsArrOrSlc(v any) bool {
	return In(reflect.TypeOf(v).Kind(), reflect.Slice, reflect.Array)
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
func IsContinuous[T Integer](numbers ...T) (ok bool, minIfOk T, maxIfOk T) {
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

// IsXML : Check str is valid XML
func IsXML(str string) bool {
	return xml.Unmarshal([]byte(str), new(any)) == nil
}

// IsJSON : Check str is valid JSON
func IsJSON(str string) bool {
	return json.Unmarshal([]byte(str), new(any)) == nil
}

// IsCSV : Check str is valid CSV
func IsCSV(str string) bool {
	records, err := csv.NewReader(strings.NewReader(str)).ReadAll()
	return err == nil && len(records) > 0
}
