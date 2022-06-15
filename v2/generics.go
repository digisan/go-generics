package v2

import (
	"log"
	"reflect"
	"sort"
	"unsafe"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64 | ~string
}

// last one @ idx=1; first one @ idx=len(arr)
func Last[T any](arr []T, idx int) T {
	if len(arr) == 0 {
		panic("arr's length must > 0")
	}
	if idx <= 0 || idx > len(arr) {
		panic("idx must between [1, len(arr)]")
	}
	return arr[len(arr)-idx]
}

// ***
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

// ***
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

// ***
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

// ***
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

// *** IdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IdxOf[T comparable](e T, arr ...T) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// *** LastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func LastIdxOf[T comparable](e T, arr ...T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// *** In : if arr has element e, return true. otherwise false
func In[T comparable](e T, arr ...T) bool {
	return IdxOf(e, arr...) != -1
}

// *** NotIn : if arr does NOT have element e, return true. otherwise false
func NotIn[T comparable](e T, arr ...T) bool {
	return !In(e, arr...)
}

// ***
func DelEleOrderlyAt[T comparable](arr *[]T, i int) {
	if i >= 0 && i < len(*arr) {
		*arr = append((*arr)[:i], (*arr)[i+1:]...)
	}
}

// ***
func DelEleAt[T comparable](arr *[]T, i int) {
	if i >= 0 && i < len(*arr) {
		(*arr)[i] = (*arr)[len(*arr)-1]
		(*reflect.SliceHeader)(unsafe.Pointer(arr)).Len--
	}
}

// ***
func DelOneEle[T comparable](arr *[]T, ele T) {
	if idx := IdxOf(ele, (*arr)...); idx >= 0 {
		DelEleAt(arr, idx)
	}
}

// ***
func DelOneEleOrderly[T comparable](arr *[]T, ele T) {
	if idx := IdxOf(ele, (*arr)...); idx >= 0 {
		DelEleOrderlyAt(arr, idx)
	}
}

// *** Settify : remove repeated elements in arr
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

// *** input data will be changed, filtered data keeps order, BUT
// return RE-ORDERED original set
func FilterFast[T any](pData *[]T, check func(i int, e T) bool) []T {
	if check == nil {
		return *pData
	}

	p := *pData
	var k = 0
	for i, v := range p {
		if check(i, v) {
			p[k], p[i] = p[i], p[k]
			k++
		}
	}
	(*reflect.SliceHeader)(unsafe.Pointer(pData)).Len = k
	return p
}

// *** input data keeps original, return filtered & ordered copy
func Filter[T any](data []T, check func(i int, e T) bool) []T {
	if check == nil {
		return append([]T{}, data...)
	}

	r := make([]T, 0, len(data))
	for i, v := range data {
		if check(i, v) {
			r = append(r, v)
		}
	}
	return r
}

// ***
func Map4SglTyp[T any](arr []T, mapper func(i int, e T) T) (r []T) {
	if mapper == nil {
		return append([]T{}, arr...)
	}
	r = make([]T, 0, len(arr))
	for i, e := range arr {
		r = append(r, mapper(i, e))
	}
	return
}

func Map[T1, T2 any](arr []T1, mapper func(i int, e T1) T2) (r []T2) {
	if mapper == nil {
		log.Fatalln("mapper CANNOT be nil")
	}
	r = make([]T2, 0, len(arr))
	for i, e := range arr {
		r = append(r, mapper(i, e))
	}
	return
}

// *** FilterMap : Filter A slice, return A=>B slice. B could be different type from A
func FilterMap[T1, T2 any](arr []T1, filter func(i int, e T1) bool, mapper func(i int, e T1) T2) (r []T2) {
	// tmp := make([]T1, len(arr))
	// copy(tmp, arr)
	// FilterFast(&tmp, filter)
	// return Map(tmp, mapper)

	return Map(Filter(arr, filter), mapper)
}

// Filter A slice, return A=>B slice. B is the same type as A
func FilterMap4SglTyp[T any](arr []T, filter func(i int, e T) bool, mapper func(i int, e T) T) (r []T) {
	// tmp := make([]T, len(arr))
	// copy(tmp, arr)
	// FilterFast(&tmp, filter)
	// return Map4SglTyp(tmp, mapper)

	return Map4SglTyp(Filter(arr, filter), mapper)
}

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
		values = append(values, kvEle.val.(T2))
	}
	return
}

// *** MapSafeMerge:
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

// *** MapMerge:
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

// *** MapMergeOnValSlc :
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

// *** MapFilter:
func MapFilter[T1 comparable, T2 any](m map[T1]T2, filter func(k T1, v T2) bool) map[T1]T2 {
	rt := make(map[T1]T2)
	for k, v := range m {
		if filter(k, v) {
			rt[k] = v
		}
	}
	return rt
}

// *** MapCopy:
func MapCopy[T1 comparable, T2 any](m map[T1]T2) map[T1]T2 {
	return MapFilter(m, func(k T1, v T2) bool { return true })
}

// *** MapToValAny:
func MapToValAny[T1 comparable, T2 any](m map[T1]T2) map[T1]any {
	ret := make(map[T1]any)
	for k, v := range m {
		ret[k] = v
	}
	return ret
}

// *** MapToArrValAny:
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
func MapAllEmptyFields[T comparable](m map[T]any) bool {
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

/////////////////////////////////////////////////////////////////////////

// IsSuper :
func IsSuper[T comparable](setA, setB []T) bool {
NEXT_B:
	for _, b := range setB {
		for _, a := range setA {
			if a == b {
				continue NEXT_B
			}
		}
		return false
	}
	return len(setA) > len(setB)
}

// IsSub :
func IsSub[T comparable](setA, setB []T) bool {
	return IsSuper(setB, setA)
}

// *** equals :
func equals[T comparable](setA, setB []T) bool {
	if (setA == nil && setB != nil) || (setA != nil && setB == nil) {
		return false
	}
	if len(setA) != len(setB) {
		return false
	}

	tmpA := make([]T, len(setA))
	tmpB := make([]T, len(setB))
	copy(tmpA, setA)
	copy(tmpB, setB)

AGAIN:
	for i, a := range tmpA {
		for j, b := range tmpB {
			if a == b {
				DelEleAt(&tmpA, i)
				DelEleAt(&tmpB, j)
				goto AGAIN
			}
		}
	}
	return len(tmpA) == 0 && len(tmpB) == 0
}

// *** Equals
func Equals[T comparable](sets ...[]T) bool {
	for i := 0; i < len(sets)-1; i++ {
		this := sets[i]
		next := sets[i+1]
		if !equals(this, next) {
			return false
		}
	}
	return true
}

// SupEq :
func SupEq[T comparable](setA, setB []T) bool {
	return IsSuper(setA, setB) || Equals(setA, setB)
}

// SubEq :
func SubEq[T comparable](setA, setB []T) bool {
	return IsSub(setA, setB) || Equals(setA, setB)
}

// union :
func union[T comparable](setA, setB []T) (set []T) {
	if setA == nil && setB == nil {
		return nil
	}
	if setA == nil && setB != nil {
		return setB
	}
	if setA != nil && setB == nil {
		return setA
	}

	m := make(map[T]struct{})
	for _, a := range setA {
		if _, ok := m[a]; !ok {
			set = append(set, a)
			m[a] = struct{}{}
		}
	}
	for _, b := range setB {
		if _, ok := m[b]; !ok {
			set = append(set, b)
			m[b] = struct{}{}
		}
	}
	if set == nil {
		return []T{}
	}
	return
}

// Union :
func Union[T comparable](sets ...[]T) (set []T) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = union(set, s)
	}
	return set
}

// intersect :
func intersect[T comparable](setA, setB []T) (set []T) {
	if setA == nil || setB == nil {
		return nil
	}

	copyA, copyB := make([]T, len(setA)), make([]T, len(setB))
	copy(copyA, setA)
	copy(copyB, setB)

AGAIN:
	for i, a := range copyA {
		for j, b := range copyB {
			if a == b {
				set = append(set, a)
				DelEleAt(&copyA, i)
				DelEleAt(&copyB, j)
				goto AGAIN
			}
		}
	}
	if set == nil {
		return []T{}
	}
	return
}

// Intersect :
func Intersect[T comparable](sets ...[]T) (set []T) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = intersect(set, s)
	}
	return set
}

func minus[T comparable](setA, setB []T) (set []T) {
	if setA == nil {
		return nil
	}
	set = make([]T, 0)

NEXT_A:
	for _, a := range setA {
		for _, b := range setB {
			if a == b {
				continue NEXT_A
			}
		}
		set = append(set, a)
	}
	return
}

func Minus[T comparable](setA []T, setOthers ...[]T) (set []T) {
	return minus(setA, Union(setOthers...))
}

// *** Reorder : any index must less than len(arr); ([4,2,3,1],[2,1,3,0]) => [3,2,1,4]
func Reorder[T any](arr []T, indices []int) (orders []T) {
	if arr == nil || indices == nil {
		return nil
	}
	if len(arr) == 0 || len(indices) == 0 {
		return []T{}
	}
	for _, idx := range indices {
		orders = append(orders, arr[idx])
	}
	return orders
}

// *** Reverse : [1,2,3] => [3,2,1]
func Reverse[T any](arr []T) []T {
	indices := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		indices[i] = len(arr) - 1 - i
	}
	return Reorder(arr, indices)
}

// *** Reduce :
func Reduce[T any](arr []T, reduce func(e0, e1 T) T) (r T) {
	if len(arr) < 2 {
		panic("Reduce at least receives 2 parameters")
	}

	for i := 0; i < len(arr)-1; i++ {
		j := i + 1
		e0, e1 := arr[i], arr[j]
		if i > 0 {
			e0 = r
		}
		r = reduce(e0, e1)
	}
	return r
}

// *** ZipSlice : [{1,2}, {3,4,5}, {6,7,8,9}] =>  [{1,3,6}, {2,4,7}]
func ZipSlice[T any](slices ...[]T) (zipped [][]T) {
	Min := func(data ...int) int {
		min := data[0]
		for i := 1; i < len(data); i++ {
			if data[i] < min {
				min = data[i]
			}
		}
		return min
	}

	lens := []int{}
	for _, slc := range slices {
		lens = append(lens, len(slc))
	}
	min := Min(lens...)
	for i := 0; i < min; i++ {
		tuple := []T{}
		for _, slc := range slices {
			tuple = append(tuple, slc[i])
		}
		zipped = append(zipped, tuple)
	}
	return
}

func ZipDim2[T any](slices ...[]T) (za [][2]T) {
	for _, s := range ZipSlice(slices...) {
		za = append(za, *(*[2]T)(s))
	}
	return
}

func ZipDim3[T any](slices ...[]T) (za [][3]T) {
	for _, s := range ZipSlice(slices...) {
		za = append(za, *(*[3]T)(s))
	}
	return
}

func ZipDim4[T any](slices ...[]T) (za [][4]T) {
	for _, s := range ZipSlice(slices...) {
		za = append(za, *(*[4]T)(s))
	}
	return
}

// *** MergeArray : [{2,2}, {3,3,4}, {6,6,8,8}] => {2,2,3,3,4,6,6,8,8}
func MergeArray[T any](arrays ...[]T) (merged []T) {
	for _, arr := range arrays {
		merged = append(merged, arr...)
	}
	return
}

// *** MergeSet : [{2,2}, {3,3,4}, {6,6,8,8}] => {2,3,4,6,8}
func MergeSet[T comparable](arrays ...[]T) (merged []T) {
	for _, arr := range arrays {
		merged = append(merged, arr...)
	}
	return Settify(merged...)
}
