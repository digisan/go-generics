package v2

import (
	"reflect"
	"sort"
	"unsafe"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64 | ~string
}

func Max[T Ordered](arr ...T) T {
	if len(arr) == 0 {
		panic("Max args at least has one element")
	}
	m := arr[0]
	for _, a := range arr[1:] {
		if a > m {
			m = a
		}
	}
	return m
}

func MaxIdx[T Ordered](arr ...T) (T, int) {
	if len(arr) == 0 {
		panic("Max args at least has one element")
	}
	idx, m := 0, arr[0]
	for i, a := range arr[1:] {
		if a > m {
			m = a
			idx = i + 1
		}
	}
	return m, idx
}

func Min[T Ordered](arr ...T) T {
	if len(arr) == 0 {
		panic("Min args at least has one element")
	}
	m := arr[0]
	for _, a := range arr[1:] {
		if a < m {
			m = a
		}
	}
	return m
}

func MinIdx[T Ordered](arr ...T) (T, int) {
	if len(arr) == 0 {
		panic("Min args at least has one element")
	}
	idx, m := 0, arr[0]
	for i, a := range arr[1:] {
		if a < m {
			m = a
			idx = i + 1
		}
	}
	return m, idx
}

// IdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IdxOf[T comparable](e T, arr ...T) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// LastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func LastIdxOf[T comparable](e T, arr ...T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// In : if arr has element e, return true. otherwise false
func In[T comparable](e T, arr ...T) bool {
	return IdxOf(e, arr...) != -1
}

// NotIn : if arr does NOT have element e, return true. otherwise false
func NotIn[T comparable](e T, arr ...T) bool {
	return !In(e, arr...)
}

func DelEleOrderlyAt[T comparable](arr *[]T, i int) {
	if i >= 0 && i < len(*arr) {
		*arr = append((*arr)[:i], (*arr)[i+1:]...)
	}
}

func DelEleAt[T comparable](arr *[]T, i int) {
	if i >= 0 && i < len(*arr) {
		(*arr)[i] = (*arr)[len(*arr)-1]
		(*reflect.SliceHeader)(unsafe.Pointer(arr)).Len--
	}
}

func DelOneEle[T comparable](arr *[]T, ele T) {
	if idx := IdxOf(ele, (*arr)...); idx >= 0 {
		DelEleAt(arr, idx)
	}
}

func DelOneEleOrderly[T comparable](arr *[]T, ele T) {
	if idx := IdxOf(ele, (*arr)...); idx >= 0 {
		DelEleOrderlyAt(arr, idx)
	}
}

// Settify : remove repeated elements in arr
func Settify[T comparable](arr ...T) (set []T) {
	if len(arr) == 0 {
		return arr
	}
	set = make([]T, 0, len(arr))
	m := make(map[T]struct{})
	for _, ele := range arr {
		if _, ok := m[ele]; !ok {
			set = append(set, ele)
			m[ele] = struct{}{}
		}
	}
	if len(set) == 0 {
		return []T{}
	}
	return
}

func Filter[T any](data *[]T, check func(i int, e T) bool) []T {
	if check == nil {
		return *data
	}

	p := *data
	var k = 0
	for i, v := range p {
		if check(i, v) {
			p[k], p[i] = p[i], p[k]
			k++
		}
	}
	(*reflect.SliceHeader)(unsafe.Pointer(data)).Len = k
	return p[:k]
}

func Map[T1, T2 any](arr []T1, mapper func(i int, e T1) T2) (r []T2) {
	r = make([]T2, 0, len(arr))
	if mapper != nil {
		for i, e := range arr {
			r = append(r, mapper(i, e))
		}
	}
	return
}

// FilterMap : Filter & Modify A slice, return B slice
func FilterMap[T1, T2 any](arr []T1, filter func(i int, e T1) bool, mapper func(i int, e T1) T2) (r []T2) {
	return Map(Filter(&arr, filter), mapper)
}

// for Map2KVs
type kv struct {
	key any
	val any
}

// Map2KVs : map to key slice & value slice
func Map2KVs[T1, T2 comparable](m map[T1]T2, less4k func(i, j T1) bool, less4v func(i, j T2) bool) (keys []T1, values []T2) {

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
		values = append(values, kvEle.val.(T2))
	}
	return
}

// TODO...
// MapSafeMerge
// MapReplaceMerge

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
