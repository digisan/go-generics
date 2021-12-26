package f64

import (
	"reflect"
	"sort"
	"unsafe"
)

func DelEleOrderly(arr *[]float64, i int) {
	*arr = append((*arr)[:i], (*arr)[i+1:]...)
}

func DelEle(arr *[]float64, i int) {
	(*arr)[i] = (*arr)[len(*arr)-1]
	(*reflect.SliceHeader)(unsafe.Pointer(arr)).Len--
}

func Max(arr ...float64) float64 {
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

func MaxIdx(arr ...float64) (float64, int) {
	if len(arr) == 0 {
		panic("Max args at least has one element")
	}
	m := arr[0]
	idx := 0
	for i, a := range arr[1:] {
		if a > m {
			m = a
			idx = i + 1
		}
	}
	return m, idx
}

func Min(arr ...float64) float64 {
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

func MinIdx(arr ...float64) (float64, int) {
	if len(arr) == 0 {
		panic("Min args at least has one element")
	}
	m := arr[0]
	idx := 0
	for i, a := range arr[1:] {
		if a < m {
			m = a
			idx = i + 1
		}
	}
	return m, idx
}

// In : if arr has element e, return true. otherwise false
func In(e float64, arr ...float64) bool {
	return IdxOf(e, arr...) != -1
}

// NotIn : if arr does NOT have element e, return true. otherwise false
func NotIn(e float64, arr ...float64) bool {
	return !In(e, arr...)
}

// IdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IdxOf(e float64, arr ...float64) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// LastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func LastIdxOf(e float64, arr ...float64) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// MkSet : remove repeated elements in arr
func MkSet(arr ...float64) (set []float64) {
	if arr == nil {
		return nil
	}
	m := make(map[float64]struct{})
	for _, ele := range arr {
		if _, ok := m[ele]; !ok {
			set = append(set, ele)
			m[ele] = struct{}{}
		}
	}
	if len(set) == 0 {
		return []float64{}
	}
	return
}

// Superset :
func Superset(setA, setB []float64) bool {
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

// Subset :
func Subset(setA, setB []float64) bool {
	return Superset(setB, setA)
}

// equal :
func equal(setA, setB []float64) bool {
	if (setA == nil && setB != nil) || (setA != nil && setB == nil) {
		return false
	}
	if len(setA) != len(setB) {
		return false
	}

	tmpA := make([]float64, len(setA))
	tmpB := make([]float64, len(setB))
	copy(tmpA, setA)
	copy(tmpB, setB)

AGAIN:
	for i, a := range tmpA {
		for j, b := range tmpB {
			if a == b {
				tmpA = append(tmpA[:i], tmpA[i+1:]...)
				tmpB = append(tmpB[:j], tmpB[j+1:]...)
				goto AGAIN
			}
		}
	}
	return len(tmpA) == 0 && len(tmpB) == 0
}

// Equal
func Equal(sets ...[]float64) bool {
	for i := 0; i < len(sets)-1; i++ {
		this := sets[i]
		next := sets[i+1]
		if !equal(this, next) {
			return false
		}
	}
	return true
}

// SuperEq :
func SuperEq(setA, setB []float64) bool {
	return Superset(setA, setB) || Equal(setA, setB)
}

// SubEq :
func SubEq(setA, setB []float64) bool {
	return Subset(setA, setB) || Equal(setA, setB)
}

// union :
func union(setA, setB []float64) (set []float64) {
	if setA == nil && setB == nil {
		return nil
	}
	if setA == nil && setB != nil {
		return setB
	}
	if setA != nil && setB == nil {
		return setA
	}

	m := make(map[float64]struct{})
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
		return []float64{}
	}
	return
}

// Union :
func Union(sets ...[]float64) (set []float64) {
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
func intersect(setA, setB []float64) (set []float64) {
	if setA == nil || setB == nil {
		return nil
	}

	copyA, copyB := make([]float64, len(setA)), make([]float64, len(setB))
	copy(copyA, setA)
	copy(copyB, setB)

AGAIN:
	for i, a := range copyA {
		for j, b := range copyB {
			if a == b {
				set = append(set, a)
				copyA = append(copyA[:i], copyA[i+1:]...)
				copyB = append(copyB[:j], copyB[i+j:]...)
				goto AGAIN
			}
		}
	}
	if set == nil {
		return []float64{}
	}
	return
}

// Intersect :
func Intersect(sets ...[]float64) (set []float64) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = intersect(set, s)
	}
	return set
}

func minus(setA, setB []float64) (set []float64) {
	if setA == nil {
		return nil
	}
	set = make([]float64, 0)

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

func Minus(setA []float64, setOthers ...[]float64) (set []float64) {
	return minus(setA, Union(setOthers...))
}

// Reorder : any index must less than len(arr)
func Reorder(arr []float64, indices []int) (orders []float64) {
	if arr == nil || indices == nil {
		return nil
	}
	if len(arr) == 0 || len(indices) == 0 {
		return []float64{}
	}
	for _, idx := range indices {
		orders = append(orders, arr[idx])
	}
	return orders
}

// Reverse : [1,2,3] => [3,2,1]
func Reverse(arr []float64) []float64 {
	indices := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		indices[i] = len(arr) - 1 - i
	}
	return Reorder(arr, indices)
}

// Reduce :
func Reduce(arr []float64, reduce func(e0, e1 float64) float64) float64 {
	switch len(arr) {
	case 0, 1:
		panic("Reduce at least receives 2 parameters")
	default:
		var r float64
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
}

// ZipArray :
func ZipArray(arrays ...[]float64) (zipped [][]float64) {

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
	for _, arr := range arrays {
		lens = append(lens, len(arr))
	}
	min := Min(lens...)
	for i := 0; i < min; i++ {
		tuple := []float64{}
		for _, arr := range arrays {
			tuple = append(tuple, arr[i])
		}
		zipped = append(zipped, tuple)
	}
	return
}

func Filter(data *[]float64, check func(i int, e float64) bool) []float64 {
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

// FilterMap : Filter & Modify []float64 slice, return []float64 slice
func FilterMap(arr []float64, filter func(i int, e float64) bool, modifier func(i int, e float64) float64) (r []float64) {
	switch {
	case filter != nil && modifier != nil:
		for i, e := range arr {
			if filter(i, e) {
				r = append(r, modifier(i, e))
			}
		}
	case filter != nil && modifier == nil:
		for i, e := range arr {
			if filter(i, e) {
				r = append(r, e)
			}
		}
	case filter == nil && modifier != nil:
		for i, e := range arr {
			r = append(r, modifier(i, e))
		}
	default:
		return arr
	}
	return
}

var (
	FM = FilterMap
)

// Map2KVs : map to key slice & value slice
func Map2KVs(m map[float64]float64, less4key func(i float64, j float64) bool, less4value func(i float64, j float64) bool) (keys []float64, values []float64) {

	if m == nil {
		return nil, nil
	}
	if len(m) == 0 {
		return []float64{}, []float64{}
	}

	type kv struct {
		key   float64
		value float64
	}

	kvSlc := []kv{}
	for k, v := range m {
		kvSlc = append(kvSlc, kv{key: k, value: v})
	}

	switch {
	case less4key != nil && less4value == nil:
		sort.SliceStable(kvSlc, func(i, j int) bool { return less4key(kvSlc[i].key, kvSlc[j].key) })

	case less4key == nil && less4value != nil:
		sort.SliceStable(kvSlc, func(i, j int) bool { return less4value(kvSlc[i].value, kvSlc[j].value) })

	case less4key != nil && less4value != nil:
		sort.SliceStable(kvSlc, func(i, j int) bool {
			if kvSlc[i].value == kvSlc[j].value {
				return less4key(kvSlc[i].key, kvSlc[j].key)
			}
			return less4value(kvSlc[i].value, kvSlc[j].value)
		})

	default:
		// do not sort
	}

	for _, kvEle := range kvSlc {
		keys = append(keys, kvEle.key)
		values = append(values, kvEle.value)
	}
	return
}

// MapMerge:
func MapMerge(ms ...map[float64]float64) map[float64][]float64 {
	res := map[float64][]float64{}
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
