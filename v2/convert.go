package v2

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"unsafe"
)

// StrToConstBytes converts string to byte slice without a memory allocation.
func StrToConstBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

// ConstBytesToStr converts byte slice to string without a memory allocation.
func ConstBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func AnyTryToType[T any](v any) (T, bool) {

	sv := fmt.Sprint(v)

	fv, errF := strconv.ParseFloat(sv, 64)
	iv, errI := strconv.ParseInt(sv, 10, 64)
	uv, errU := strconv.ParseUint(sv, 10, 64)
	bv, errB := strconv.ParseBool(sv)
	cv, errC := strconv.ParseComplex(sv, 128)

	switch fmt.Sprintf("%T", new(T)) {

	case "*float64":
		if errF == nil {
			r := float64(fv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*float32":
		if errF == nil {
			r := float32(fv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "*int":
		if errI == nil {
			r := int(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*int8":
		if errI == nil {
			r := int8(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*int16":
		if errI == nil {
			r := int16(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*int32":
		if errI == nil {
			r := int32(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*int64":
		if errI == nil {
			r := int64(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "*uint":
		if errU == nil {
			r := uint(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*uint8":
		if errU == nil {
			r := uint8(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*uint16":
		if errU == nil {
			r := uint16(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*uint32":
		if errU == nil {
			r := uint32(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "*uint64":
		if errU == nil {
			r := uint64(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "*bool":
		if errB == nil {
			r := bv
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "*complex64":
		if errC == nil {
			r := complex64(cv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "*complex128":
		if errC == nil {
			r := complex128(cv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "*string":
		return *(*T)(unsafe.Pointer(&sv)), true
	}

	return *new(T), false
}

// T is return type
func AnysToTypes[T any](s []any) []T {
	if s == nil {
		return nil
	}
	rt := make([]T, 0, len(s))
	for _, a := range s {
		rt = append(rt, a.(T))
	}
	return rt
}

// s(any) must be []any. T is return type. If return nil, convert failed
func AnysAsAnyToTypes[T any](s any) []T {
	if s == nil {
		return nil
	}
	switch v := s.(type) {
	case []any:
		return AnysToTypes[T](v)
	}
	return nil
}

// s(any) is any actual type of slice or array
func SlcToAnys(s any) (rt []any) {
	if IsArrOrSlc(s) {
		s := reflect.ValueOf(s)
		for i := 0; i < s.Len(); i++ {
			rt = append(rt, s.Index(i).Interface())
		}
	}
	return
}

// s(any) is any actual type of slice or array, T is return type
func SlcToTypes[T any](s any) []T {
	return AnysToTypes[T](SlcToAnys(s))
}

/////////////////////////////////////////////////////////////////

func SyncMapToMap[T1 comparable, T2 any](sm sync.Map) map[T1]T2 {
	rt := make(map[T1]T2)
	sm.Range(func(key, value any) bool {
		rt[key.(T1)] = value.(T2)
		return true
	})
	return rt
}

func MapToSyncMap[T1 comparable, T2 any](m map[T1]T2) sync.Map {
	rt := sync.Map{}
	for k, v := range m {
		rt.Store(k, v)
	}
	return rt
}

func MapAnyToType[T1 comparable, T2 any](m map[any]any) map[T1]T2 {
	if m == nil {
		return nil
	}
	rt := make(map[T1]T2)
	for k, v := range m {
		rt[k.(T1)] = v.(T2)
	}
	return rt
}
