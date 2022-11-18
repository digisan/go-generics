package v2

import (
	"reflect"
	"sync"
)

func MapCvt[T1 comparable, T2 any](m map[any]any) map[T1]T2 {
	if m == nil {
		return nil
	}
	rt := make(map[T1]T2)
	for k, v := range m {
		rt[k.(T1)] = v.(T2)
	}
	return rt
}

// T is return type
func SlcCvt[T any](s []any) []T {
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
func AnySlc2Slc[T any](s any) []T {
	if s == nil {
		return nil
	}
	switch v := s.(type) {
	case []any:
		return SlcCvt[T](v)
	}
	return nil
}

// s(any) is any actual type of slice or array
func Any2AnySlc(s any) (rt []any) {
	if IsArrayOrSlice(s) {
		s := reflect.ValueOf(s)
		for i := 0; i < s.Len(); i++ {
			rt = append(rt, s.Index(i).Interface())
		}
	}
	return
}

// s(any) is any actual type of slice or array, T is return type
func Any2Slc[T any](s any) []T {
	return AnySlc2Slc[T](Any2AnySlc(s))
}

func SyncMap2Map[T1 comparable, T2 any](sm sync.Map) map[T1]T2 {
	rt := make(map[T1]T2)
	sm.Range(func(key, value any) bool {
		rt[key.(T1)] = value.(T2)
		return true
	})
	return rt
}

func Map2SyncMap[T1 comparable, T2 any](m map[T1]T2) sync.Map {
	rt := sync.Map{}
	for k, v := range m {
		rt.Store(k, v)
	}
	return rt
}
