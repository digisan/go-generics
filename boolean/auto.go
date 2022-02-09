package boolean

import (
	"reflect"
	"unsafe"
)

func DelEleOrderly(arr *[]bool, i int) {
	*arr = append((*arr)[:i], (*arr)[i+1:]...)
}

func DelEle(arr *[]bool, i int) {
	(*arr)[i] = (*arr)[len(*arr)-1]
	(*reflect.SliceHeader)(unsafe.Pointer(arr)).Len--
}

// In : if arr has element e, return true. otherwise false
func In(e bool, arr ...bool) bool {
	return IdxOf(e, arr...) != -1
}

// NotIn : if arr does NOT have element e, return true. otherwise false
func NotIn(e bool, arr ...bool) bool {
	return !In(e, arr...)
}

// IdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IdxOf(e bool, arr ...bool) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// LastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func LastIdxOf(e bool, arr ...bool) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// MkSet : remove repeated elements in arr
func MkSet(arr ...bool) (set []bool) {
	if arr == nil {
		return nil
	}
	m := make(map[bool]struct{})
	for _, ele := range arr {
		if _, ok := m[ele]; !ok {
			set = append(set, ele)
			m[ele] = struct{}{}
		}
	}
	if len(set) == 0 {
		return []bool{}
	}
	return
}

// Superset :
func Superset(setA, setB []bool) bool {
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
func Subset(setA, setB []bool) bool {
	return Superset(setB, setA)
}

// equal :
func equal(setA, setB []bool) bool {
	if (setA == nil && setB != nil) || (setA != nil && setB == nil) {
		return false
	}
	if len(setA) != len(setB) {
		return false
	}

	tmpA := make([]bool, len(setA))
	tmpB := make([]bool, len(setB))
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
func Equal(sets ...[]bool) bool {
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
func SuperEq(setA, setB []bool) bool {
	return Superset(setA, setB) || Equal(setA, setB)
}

// SubEq :
func SubEq(setA, setB []bool) bool {
	return Subset(setA, setB) || Equal(setA, setB)
}

// union :
func union(setA, setB []bool) (set []bool) {
	if setA == nil && setB == nil {
		return nil
	}
	if setA == nil && setB != nil {
		return setB
	}
	if setA != nil && setB == nil {
		return setA
	}

	m := make(map[bool]struct{})
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
		return []bool{}
	}
	return
}

// Union :
func Union(sets ...[]bool) (set []bool) {
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
func intersect(setA, setB []bool) (set []bool) {
	if setA == nil || setB == nil {
		return nil
	}

	copyA, copyB := make([]bool, len(setA)), make([]bool, len(setB))
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
		return []bool{}
	}
	return
}

// Intersect :
func Intersect(sets ...[]bool) (set []bool) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = intersect(set, s)
	}
	return set
}

func minus(setA, setB []bool) (set []bool) {
	if setA == nil {
		return nil
	}
	set = make([]bool, 0)

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

func Minus(setA []bool, setOthers ...[]bool) (set []bool) {
	return minus(setA, Union(setOthers...))
}

// Reorder : any index must less than len(arr)
func Reorder(arr []bool, indices []int) (orders []bool) {
	if arr == nil || indices == nil {
		return nil
	}
	if len(arr) == 0 || len(indices) == 0 {
		return []bool{}
	}
	for _, idx := range indices {
		orders = append(orders, arr[idx])
	}
	return orders
}

// Reverse : [1,2,3] => [3,2,1]
func Reverse(arr []bool) []bool {
	indices := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		indices[i] = len(arr) - 1 - i
	}
	return Reorder(arr, indices)
}

// Reduce :
func Reduce(arr []bool, reduce func(e0, e1 bool) bool) bool {
	switch len(arr) {
	case 0, 1:
		panic("Reduce at least receives 2 parameters")
	default:
		var r bool
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
func ZipArray(arrays ...[]bool) (zipped [][]bool) {

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
		tuple := []bool{}
		for _, arr := range arrays {
			tuple = append(tuple, arr[i])
		}
		zipped = append(zipped, tuple)
	}
	return
}

func Filter(data *[]bool, check func(i int, e bool) bool) []bool {
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

// MapFilter:
func MapFilter(m map[bool]bool, filter func(k bool, v bool) bool) map[bool]bool {
	rt := make(map[bool]bool)
	for k, v := range m {
		if filter(k, v) {
			rt[k] = v
		}
	}
	return rt
}

// MapCopy:
func MapCopy(m map[bool]bool) map[bool]bool {
	return MapFilter(m, func(k bool, v bool) bool { return true })
}
