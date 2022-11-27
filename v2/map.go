package v2

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
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

func dumpMap(pk string, jv any, mflat *map[string]any) {

	switch m := jv.(type) {

	case float64, float32, string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, complex64, complex128, nil:
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
		if IsArrOrSlc(m) {
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
		if ik, err := strconv.ParseInt(fmt.Sprint(k), 10, 64); err != nil || int64(i) != ik {
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

// This can only set single value in nested slice, useless for further assignment
//
// func MakeNestedSlice(value any, idxSegs ...any) (slc []any, err error) {
// 	if len(idxSegs) == 0 {
// 		return []any{value}, nil
// 	}
// 	for _, seg := range idxSegs {
// 		if !IsUint(seg) {
// 			return nil, fmt.Errorf("index path (idxSegs) must can be converted to unsigned int")
// 		}
// 	}

// 	for i, seg := range Reverse(idxSegs) {
// 		idx, _ := AnyTryToType[int](seg)

// 		if i == 0 { // value seg

// 			nNil := idx - len(slc)
// 			for i := 0; i < nNil; i++ {
// 				slc = append(slc, nil)
// 			}
// 			slc = append(slc, value)

// 		} else { // path seg

// 			slc = []any{slc}

// 			nNil := idx - len(slc)
// 			for i := 0; i <= nNil; i++ {
// 				slc = append(slc, []any{})
// 			}
// 			slc = Reverse(slc)
// 		}
// 	}
// 	return
// }

func traverseNestedSlice(slc []any, path string, sb *strings.Builder) {
	for i, e := range slc {
		if IsArrOrSlc(e) {
			traverseNestedSlice(e.([]any), fmt.Sprintf("%v.%d", path, i), sb)
		} else {

			// **
			// with value
			// **
			// path := fmt.Sprintf("%v @ %v.%d", e, path, i)
			// path = strings.Replace(path, "@ .", "@ ", 1)

			// **
			// without value
			// **
			path := fmt.Sprintf("%v.%d", path, i)
			path = strings.TrimPrefix(path, ".")

			sb.WriteString(path + "\n")

			// debug
			// if e != nil {
			// 	fmt.Printf("%v @%p %s\n", slc[i], &slc[i], path)
			// } else {
			// 	fmt.Printf("%v %s\n", slc[i], path)
			// }
		}
	}
}

func TraverseNestedSlice(slc []any) []string {
	var (
		path = ""
		sb   = &strings.Builder{}
	)
	traverseNestedSlice(slc, path, sb)
	return strings.Split(strings.TrimSpace(sb.String()), "\n")
}

// path is like '1.2.3'
func CapacityForSlice(paths ...string) []int {

	mLvlIndices := make(map[int][]int)
	for _, path := range paths {
		// fmt.Println(i, path)
		for lvl, seg := range strings.Split(path, ".") {
			if n, ok := AnyTryToType[int](seg); ok {
				mLvlIndices[lvl] = append(mLvlIndices[lvl], n)
			}
		}
	}

	mLvlCap := make(map[int]int)
	for lvl, indices := range mLvlIndices {
		mLvlCap[lvl] = Max(indices...) + 1
	}

	_, values := MapToKVs(mLvlCap, func(ki int, kj int) bool { return ki < kj }, nil)
	return values
}

func initNestedSlice(slc *any, caps ...int) {
	for lvl, c := range caps {
		*slc = make([]any, c)
		if lvl < len(caps)-1 {
			for i := 0; i < c; i++ {
				initNestedSlice(&((*slc).([]any))[i], caps[1:]...)
			}
		} else if lvl == len(caps)-1 {
			for i := 0; i < c; i++ {
				((*slc).([]any))[i] = struct{}{}
			}
		}
		break
	}
}

func InitNestedSlice(caps ...int) []any {
	var slc any = []any{}
	initNestedSlice(&slc, caps...)
	return slc.([]any)
}

// slc must have enough capacity for all element
func SetNestedSlice(slc []any, value any, idxSegs ...any) (ok bool) {
	if len(idxSegs) == 0 {
		return false
	}
	for _, seg := range idxSegs {
		if !IsUint(seg) {
			return false
		}
	}

	for i, seg := range idxSegs {
		if idx, _ := AnyTryToType[int](seg); idx < len(slc) {

			if i == len(idxSegs)-1 {
				slc[idx] = value
				return true
			} else {
				var ok bool
				if slc, ok = slc[idx].([]any); !ok {
					return false
				}
			}

		} else {
			return false
		}
	}
	return true
}

// func SetNestedMapWithIdx[T comparable](m any, value any, keySegs ...T) any {
// 	if len(keySegs) == 0 {
// 		panic("key path (keySegs) must be provided")
// 	}
// 	for _, seg := range keySegs {
// 		if IsNumeric(seg) {

// 			idx, _ := strconv.ParseInt(fmt.Sprint(seg), 10, 64)
// 			if idx == 0 {

// 				rt = []any{value}

// 			} else {

// 				if int(idx) != len(rt.([]any)) {
// 					panic("error index order")
// 				}
// 				pM = append(pM.([]any), value)
// 			}

// 		} else {

// 		}
// 	}

// return rt
// }

func MapFlatToNested(m map[string]any) map[string]any {

	keys, vals := MapToKVs(m, func(i, j string) bool { return strings.Count(i, ".") < strings.Count(j, ".") }, nil)
	// fmt.Println(keys)
	// fmt.Println(vals)

	rt := make(map[string]any)
	for i, key := range keys {
		val := vals[i]
		SetNestedMapIgnoreIdx(rt, val, strings.Split(key, ".")...)
	}
	return rt
}
