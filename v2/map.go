package v2

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// for Map2KVs
type kv struct {
	key any
	val any
}

// *** Map2KVs : map to key slice & value slice
func Map2KVs[T1 comparable, T2 any](m map[T1]T2, less4k func(i, j T1) bool, less4v func(i, j T2) bool) (keys []T1, values []T2) {

	kvSlc := make([]kv, 0, len(m))
	for k, v := range m {
		kvSlc = append(kvSlc, kv{key: k, val: v})
	}

	switch {
	case less4k != nil && less4v == nil:
		sort.SliceStable(kvSlc, func(i, j int) bool { return less4k(kvSlc[i].key.(T1), kvSlc[j].key.(T1)) })

	case less4k == nil && less4v != nil:
		sort.SliceStable(kvSlc, func(i, j int) bool { return less4v(kvSlc[i].val.(T2), kvSlc[j].val.(T2)) })

	case less4k != nil && less4v != nil:
		sort.SliceStable(kvSlc, func(i, j int) bool {
			if kvSlc[i].val == kvSlc[j].val {
				return less4k(kvSlc[i].key.(T1), kvSlc[j].key.(T1))
			}
			return less4v(kvSlc[i].val.(T2), kvSlc[j].val.(T2))
		})

	default:
		// do not sort
	}

	keys = make([]T1, 0, len(m))
	values = make([]T2, 0, len(m))
	for _, kvEle := range kvSlc {
		keys = append(keys, kvEle.key.(T1))
		if kvEle.val != nil {
			values = append(values, kvEle.val.(T2))
		} else {
			values = append(values, *new(T2))
		}
	}
	return
}

// MapSafeMerge:
func MapSafeMerge[T1 comparable, T2 any](ms ...map[T1]T2) map[T1]T2 {
	res := map[T1]T2{}
	for _, m := range ms {
		for k, v := range m {
			if _, ok := res[k]; !ok {
				res[k] = v
			}
		}
	}
	return res
}

// MapReplaceMerge
func MapReplaceMerge[T1 comparable, T2 any](ms ...map[T1]T2) map[T1]T2 {
	res := map[T1]T2{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// MapMerge:
func MapMerge[T1, T2 comparable](ms ...map[T1]T2) map[T1][]T2 {
	res := map[T1][]T2{}
	for _, m := range ms {
	srcMap:
		for k, v := range m {
			// Check if (k,v) was added before:
			for _, v2 := range res[k] {
				if v == v2 {
					continue srcMap
				}
			}
			res[k] = append(res[k], v)
		}
	}
	return res
}

// MapMergeOnValSlc :
func MapMergeOnValSlc[T1, T2 comparable](ms ...map[T1][]T2) map[T1][]T2 {
	res := map[T1][]T2{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = append(res[k], v...)
			res[k] = Settify(res[k]...)
		}
	}
	return res
}

// MapFilter:
func MapFilter[T1 comparable, T2 any](m map[T1]T2, filter func(k T1, v T2) bool) map[T1]T2 {
	rt := make(map[T1]T2)
	for k, v := range m {
		if filter(k, v) {
			rt[k] = v
		}
	}
	return rt
}

// MapCopy:
func MapCopy[T1 comparable, T2 any](m map[T1]T2) map[T1]T2 {
	return MapFilter(m, func(k T1, v T2) bool { return true })
}

// MapToValAny:
func MapToValAny[T1 comparable, T2 any](m map[T1]T2) map[T1]any {
	ret := make(map[T1]any)
	for k, v := range m {
		ret[k] = v
	}
	return ret
}

// MapToArrValAny:
func MapToArrValAny[T1 comparable, T2 any](m map[T1][]T2) map[T1][]any {
	ret := make(map[T1][]any)
	for k, v := range m {
		ret[k] = make([]any, 0, len(v))
		for _, item := range v {
			ret[k] = append(ret[k], item)
		}
	}
	return ret
}

// e.g. [ nil, "", []int{}, XXX ptr(nil) ] are 'empty'
// [ &[]int{}, &XXX{} ] are NOT 'empty'
func MapAllValuesAreEmpty[T comparable](m map[T]any) bool {
	for _, v := range m {
		if sv, ok := v.(string); ok {
			if len(sv) > 0 {
				return false
			}
		} else if v != nil {
			rv := reflect.ValueOf(v)
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
				if rv.Len() > 0 {
					return false
				}
			case reflect.Pointer:
				if !rv.IsNil() {
					return false
				}
			default:
				return false
			}
		}
	}
	return true
}

func dumpMap(pk string, jv any, mflat *map[string]any) {

	switch m := jv.(type) {
	case float64, float32, string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, complex64, complex128, nil:
		(*mflat)[pk] = m

	case []any:
		for i, a := range m {
			idx := fmt.Sprintf("%s.%d", pk, i)
			dumpMap(idx, a, mflat)
		}

	case map[string]any:
		{
			for k, v := range m {
				if pk != "" {
					k = fmt.Sprintf("%s.%s", pk, k)
				}

				switch mv := v.(type) {
				case []any:
					for i, a := range v.([]any) {
						idx := fmt.Sprintf("%s.%d", k, i)
						dumpMap(idx, a, mflat)
					}
				default:
					dumpMap(k, mv, mflat)
				}
			}
		}
	}
}

func MapNestedToFlat(m map[string]any) map[string]any {
	flatMap := make(map[string]any)
	dumpMap("", m, &flatMap)
	return flatMap
}

func SetNestedMap[T comparable](m map[T]any, value any, keySegs ...T) {
	pM := &m
	for i, seg := range keySegs {
		if i < len(keySegs)-1 {
			if subM, ok := (*pM)[seg]; !ok {
				deepM := make(map[T]any)
				(*pM)[seg] = deepM
				pM = &deepM
			} else {
				m := subM.(map[T]any)
				pM = &m
			}
		} else {
			(*pM)[seg] = value
		}
	}
}

func MapFlatToNested(m map[string]any) map[string]any {

	keys, vals := Map2KVs(m, func(i, j string) bool { return strings.Count(i, ".") < strings.Count(j, ".") }, nil)
	// fmt.Println(keys)
	// fmt.Println(vals)

	rt := make(map[string]any)
	for i, key := range keys {
		val := vals[i]
		SetNestedMap(rt, val, strings.Split(key, ".")...)
	}
	return rt
}
