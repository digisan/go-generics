package v2

import (
	"reflect"
	"sync"
)

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