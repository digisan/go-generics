package gogenerics

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"unsafe"
)

// StrToConstBytes converts string to byte slice without a memory allocation.
func StrToConstBytes(s string) (b []byte) {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// ConstBytesToStr converts byte slice to string without a memory allocation.
func ConstBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// if return type 'T' is []byte, then return value is v's printed string as []byte, NOT binary encoding
//
// optBase is 'parse (u)int base (2, 8, 10, 16)'
func AnyTryToType[T any](v any, optBase ...any) (T, bool) {

	sv := ""
	if TypeOf(v) == "[]uint8" {
		sv = ConstBytesToStr(v.([]byte))
	} else {
		sv = fmt.Sprint(v)
	}

	base := 0
	if len(optBase) > 0 {
		if b, ok := AnyTryToType[int](optBase[0]); ok {
			base = b
		}
	}

	fv, errF := strconv.ParseFloat(sv, 64)
	iv, errI := strconv.ParseInt(sv, base, 64)
	uv, errU := strconv.ParseUint(sv, base, 64)
	bv, errB := strconv.ParseBool(sv)
	cv, errC := strconv.ParseComplex(sv, 128)
	tm, okT := TryToDateTime(sv)

	rType := fmt.Sprintf("%T", *new(T))
	switch rType {

	case "float64":
		if errF == nil {
			r := float64(fv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "float32":
		if errF == nil {
			r := float32(fv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "int":
		if errI == nil {
			r := int(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "int8":
		if errI == nil {
			r := int8(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "int16":
		if errI == nil {
			r := int16(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "int32":
		if errI == nil {
			r := int32(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "int64":
		if errI == nil {
			r := int64(iv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "uint":
		if errU == nil {
			r := uint(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "uint8":
		if errU == nil {
			r := uint8(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "uint16":
		if errU == nil {
			r := uint16(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "uint32":
		if errU == nil {
			r := uint32(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "uint64":
		if errU == nil {
			r := uint64(uv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "bool":
		if errB == nil {
			r := bv
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "complex64":
		if errC == nil {
			r := complex64(cv)
			return *(*T)(unsafe.Pointer(&r)), true
		}
	case "complex128":
		if errC == nil {
			r := complex128(cv)
			return *(*T)(unsafe.Pointer(&r)), true
		}

	case "string", "[]uint8":
		return *(*T)(unsafe.Pointer(&sv)), true

	case "time.Time":
		if okT {
			return *(*T)(unsafe.Pointer(&tm)), true
		}
	}

	return *new(T), false
}

// T is return type
func AnyToType[T any](v any) (T, error) {
	typ := reflect.TypeOf(*new(T))
	if t, ok := v.(T); ok {
		return t, nil
	} else {
		if reflect.TypeOf(v).ConvertibleTo(typ) {
			v := reflect.ValueOf(v).Convert(typ).Interface()
			return v.(T), nil
		}
	}
	return *new(T), fmt.Errorf("%v cannot be converted to Type [%T]", v, typ)
}

// T is return type
func AnysToTypes[T any](s []any) []T {
	if s == nil {
		return nil
	}
	rt := make([]T, 0, len(s))
	for _, a := range s {
		if v, err := AnyToType[T](a); err == nil {
			rt = append(rt, v)
		}
	}
	return rt
}

func AnysTryToTypes[T any](s []any) ([]T, bool) {
	if s == nil {
		return nil, false
	}
	rt, flag := make([]T, 0, len(s)), true
	for _, a := range s {
		if r, ok := AnyTryToType[T](a); ok {
			rt = append(rt, r)
		} else {
			rt = append(rt, *new(T))
			flag = false
		}
	}
	return rt, flag
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

func AnysAsAnyTryToTypes[T any](s any) ([]T, bool) {
	if s == nil {
		return nil, false
	}
	switch v := s.(type) {
	case []any:
		return AnysTryToTypes[T](v)
	}
	return nil, false
}

// s(any) is any actual type of slice or array
func TypesAsAnyToAnys(s any) (rt []any) {
	if IsArrOrSlc(s) {
		s := reflect.ValueOf(s)
		for i := 0; i < s.Len(); i++ {
			rt = append(rt, s.Index(i).Interface())
		}
	}
	return
}

// s(any) is any actual type of slice or array, T is return type
func TypesAsAnyToTypes[T any](s any) []T {
	return AnysToTypes[T](TypesAsAnyToAnys(s))
}

func TypesAsAnyTryToTypes[T any](s any) ([]T, bool) {
	return AnysTryToTypes[T](TypesAsAnyToAnys(s))
}

func SlcToPtrSlc[T any](s ...T) (rt []*T) {
	if s == nil {
		return nil
	}
	for i := 0; i < len(s); i++ {
		rt = append(rt, &s[i])
	}
	return rt
}

func PtrSlcToSlc[T any](s ...*T) (rt []T) {
	if s == nil {
		return nil
	}
	for i := 0; i < len(s); i++ {
		rt = append(rt, *s[i])
	}
	return rt
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

func MapCvtKVTypeToAny[T1 comparable, T2 any](m map[T1]T2) map[any]any {
	if m == nil {
		return nil
	}
	rt := make(map[any]any)
	for k, v := range m {
		rt[k] = v
	}
	return rt
}

func MapCvtKVAnyToType[T1 comparable, T2 any](m map[any]any) map[T1]T2 {
	if m == nil {
		return nil
	}
	rt := make(map[T1]T2)
	for k, v := range m {
		if kt, err := AnyToType[T1](k); err == nil {
			if vt, err := AnyToType[T2](v); err == nil {
				rt[kt] = vt
			}
		}
	}
	return rt
}

// TK1: input map key type; TV1: input map value type;
// TK2: output map key type; TK2: output map value type;
func MapCvtKVTypeToType[TK1, TK2 comparable, TV1, TV2 any](m map[TK1]TV1) map[TK2]TV2 {
	return MapCvtKVAnyToType[TK2, TV2](MapCvtKVTypeToAny(m))
}

func MapCvtKVTypesToAnys[T1 comparable, T2 any](m map[T1][]T2) map[any][]any {
	rt := make(map[any][]any)
	for k, v := range m {
		var a any = k
		rt[a] = TypesAsAnyToAnys(v)
	}
	return rt
}

func MapCvtKVAnysToTypes[T1 comparable, T2 any](m map[any][]any) map[T1][]T2 {
	rt := make(map[T1][]T2)
	for k, v := range m {
		if kt, err := AnyToType[T1](k); err == nil {
			rt[kt] = AnysToTypes[T2](v)
		}
	}
	return rt
}

// TK1: input map key type; TV1: input map value type;
// TK2: output map key type; TK2: output map value type;
func MapCvtKVTypesToTypes[TK1, TK2 comparable, TV1, TV2 any](m map[TK1][]TV1) map[TK2][]TV2 {
	return MapCvtKVAnysToTypes[TK2, TV2](MapCvtKVTypesToAnys(m))
}
