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

// *** MapToKVs : map to key slice & value slice
func MapToKVs[T1 comparable, T2 any](m map[T1]T2, less4k func(ki, kj T1) bool, less4v func(vi, vj T2) bool) (keys []T1, values []T2) {

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

// MapReplMerge
func MapReplMerge[T1 comparable, T2 any](ms ...map[T1]T2) map[T1]T2 {
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
func MapValTypeToAny[T1 comparable, T2 any](m map[T1]T2) map[T1]any {
	rt := make(map[T1]any)
	for k, v := range m {
		rt[k] = v
	}
	return rt
}

func MapValAnyToType[T1 comparable, T2 any](m map[T1]any) map[T1]T2 {
	rt := make(map[T1]T2)
	for k, v := range m {
		rt[k] = v.(T2)
	}
	return rt
}

func MapValTypesToAnys[T1 comparable, T2 any](m map[T1][]T2) map[T1][]any {
	rt := make(map[T1][]any)
	for k, v := range m {
		rt[k] = SlcToAnys(v)
	}
	return rt
}

func MapValAnysToTypes[T1 comparable, T2 any](m map[T1][]any) map[T1][]T2 {
	rt := make(map[T1][]T2)
	for k, v := range m {
		rt[k] = AnysToTypes[T2](v)
	}
	return rt
}

// e.g. [ nil, "", []int{}, XXX ptr(nil) ] are 'empty'
// [ &[]int{}, &XXX{} ] are NOT 'empty'
func MapAllValAreEmpty[T comparable](m map[T]any) bool {
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

func dumpMap(pk string, v any, mflat *map[string]any) {

	switch m := v.(type) {

	case float64, float32, string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, complex64, complex128, nil, struct{}:
		(*mflat)[pk] = m

	case map[string]any:
		{
			for k, v := range m {
				if pk != "" {
					k = fmt.Sprintf("%s.%s", pk, k)
				}
				dumpMap(k, v, mflat)
			}
		}

	default:

		// otherwise, jump into next...
		if IsArrOrSlc(m) {

			// empty slice or array is leaf
			if LenOfMustArrOrSlc(m) == 0 {
				(*mflat)[pk] = m
			}

			for i, a := range SlcToAnys(m) {
				idx := fmt.Sprintf("%s.%d", pk, i)
				dumpMap(idx, a, mflat)
			}
		}
	}
}

func MapNestedToFlat(m map[string]any) map[string]any {
	flatMap := make(map[string]any)
	dumpMap("", m, &flatMap)
	return flatMap
}

///////////////////////////////////////////////////////

func MapTryToSlc[T1 comparable, T2 any](m map[T1]T2) ([]T2, bool) {
	for k := range m {
		if !IsNumeric(k) {
			return nil, false
		}
	}
	keys, values := MapToKVs(m, func(ki T1, kj T1) bool {
		i, _ := AnyTryToType[uint](ki)
		j, _ := AnyTryToType[uint](kj)
		return i < j
	}, nil)
	for i, k := range keys {
		if ik, ok := AnyTryToType[int](k); !ok || ik != i {
			return nil, false
		}
	}
	return values, true
}

func SetNestedMapIgnoreIdx[T comparable](m map[T]any, value any, keySegs ...T) error {
	if len(keySegs) == 0 {
		return fmt.Errorf("key path (keySegs) must be provided")
	}

	pM := m
	for i, seg := range keySegs {
		if i < len(keySegs)-1 {
			if subM, ok := pM[seg]; !ok {
				pM[seg] = make(map[T]any)
				pM = pM[seg].(map[T]any)
			} else {
				pM = subM.(map[T]any)
			}
		} else {
			pM[seg] = value
		}
	}
	return nil
}

// m MUST have enough indexed space for elements in nested array
func SetNestedMap[T comparable](m map[T]any, value any, kiSegs ...T) error {
	if len(kiSegs) == 0 {
		return fmt.Errorf("key path (kiSegs) must be provided")
	}

	var pM any = m
	for pair := range IterPair(kiSegs...) {

		if ki := pair.a; pair.validA {

			if pair.first && IsNumeric(ki) {
				return fmt.Errorf("cannot set nested map on 1st level key as number, (want to set a slice?)")
			}

			if pair.last && !pair.validB {

				if !IsUint(ki) {
					pM.(map[T]any)[ki] = value
				} else {
					idx, _ := AnyTryToType[int](ki)
					pM.([]any)[idx] = value
				}

			} else {

				switch {
				case !IsUint(ki) && !IsUint(pair.b): // e.g. "A", "B"
					if _, ok := pM.(map[T]any)[ki]; !ok {
						pM.(map[T]any)[ki] = make(map[T]any)
					}
					pM = pM.(map[T]any)[ki]

				case !IsUint(ki) && IsUint(pair.b): // e.g. "B", "0"
					idx, _ := AnyTryToType[int](pair.b)
					if _, ok := pM.(map[T]any)[ki]; !ok {
						pM.(map[T]any)[ki] = InitSlice(idx + 1) // here once only, allocate enough space
					}
					pM = pM.(map[T]any)[ki]

				case IsUint(ki) && !IsUint(pair.b): // e.g. "0", "C"
					idx, _ := AnyTryToType[int](ki)
					if pM.([]any)[idx] == struct{}{} {
						pM.([]any)[idx] = make(map[T]any)
					}
					pM = pM.([]any)[idx]

				case IsUint(ki) && IsUint(pair.b): // e.g. "0", "1"
					idx1, _ := AnyTryToType[int](ki)
					idx2, _ := AnyTryToType[int](pair.b)
					if pM.([]any)[idx1] == struct{}{} {
						pM.([]any)[idx1] = InitSlice(idx2 + 1) // here once only, allocate enough space
					}
					pM = pM.([]any)[idx1]
				}
			}
		}
	}

	return nil
}

// if fm return's 'p' is "", then ignore this (path-value) being into Nested.
func MapFlatToNested(m map[string]any, fm func(path string, value any) (p string, v any)) map[string]any {

	// *** ERROR if put less segment path at top. if so, following short segment path may have bigger index
	//

	// keys, vals := MapToKVs(m,
	// 	func(pathi, pathj string) bool {
	// 		ni, nj := strings.Count(pathi, "."), strings.Count(pathj, ".")
	// 		if ni == nj {
	// 			ssi, ssj := strings.Split(pathi, "."), strings.Split(pathj, ".")
	// 		NEXT:
	// 			for i := 0; i < ni+1; i++ {
	// 				si, sj := ssi[i], ssj[i]
	// 				if si == sj {
	// 					continue NEXT
	// 				}
	// 				if IsUint(si) && IsUint(sj) {
	// 					idxI, _ := AnyTryToType[uint](si)
	// 					idxJ, _ := AnyTryToType[uint](sj)
	// 					if idxI == idxJ {
	// 						continue NEXT
	// 					}
	// 					return idxI > idxJ
	// 				}
	// 				return si < sj // ascii ASC, uppercase first
	// 			}
	// 		}
	// 		return ni > nj
	// 	},
	// 	nil,
	// )

	// ***
	// re-order path-keys for aim nested map has enough space to be filled
	// bigger array index path to be top area
	//

	keys, vals := MapToKVs(m,
		func(pathi, pathj string) bool {
			ni := strings.Count(pathi, ".")
			ssi, ssj := strings.Split(pathi, "."), strings.Split(pathj, ".")
		NEXT:
			for i := 0; i < ni+1; i++ {
				si, sj := ssi[i], ssj[i]
				if si == sj {
					continue NEXT
				}
				if IsUint(si) && IsUint(sj) {
					idxI, _ := AnyTryToType[uint](si)
					idxJ, _ := AnyTryToType[uint](sj)
					if idxI == idxJ {
						continue NEXT
					}
					return idxI > idxJ
				}
				return si < sj // ascii ASC, uppercase first
			}
			return true // keep original order
		},
		nil,
	)

	// fmt.Println(keys)
	// fmt.Println(vals)
	// fmt.Println()

	rt := make(map[string]any)
	for i, key := range keys {
		val := vals[i]
		if fm != nil {
			if p, v := fm(key, val); len(p) > 0 {
				SetNestedMap(rt, v, strings.Split(p, ".")...)
			}
		} else {
			SetNestedMap(rt, val, strings.Split(key, ".")...)
		}
	}
	return rt
}
