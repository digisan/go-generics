package u8

import (
	"reflect"
	"unsafe"
)

func DelEleOrderlyAt(arr *[]byte, i int) {
	*arr = append((*arr)[:i], (*arr)[i+1:]...)
}

func DelEleAt(arr *[]byte, i int) {
	(*arr)[i] = (*arr)[len(*arr)-1]
	(*reflect.SliceHeader)(unsafe.Pointer(arr)).Len--
}

func DelOneEle(arr *[]byte, ele byte) {
	if idx := IdxOf(ele, (*arr)...); idx >= 0 {
		DelEleAt(arr, idx)
	}
}

func DelOneEleOrderly(arr *[]byte, ele byte) {
	if idx := IdxOf(ele, (*arr)...); idx >= 0 {
		DelEleOrderlyAt(arr, idx)
	}
}

func Max(arr ...byte) byte {
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

func MaxIdx(arr ...byte) (byte, int) {
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

func Min(arr ...byte) byte {
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

func MinIdx(arr ...byte) (byte, int) {
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
func In(e byte, arr ...byte) bool {
	return IdxOf(e, arr...) != -1
}

// NotIn : if arr does NOT have element e, return true. otherwise false
func NotIn(e byte, arr ...byte) bool {
	return !In(e, arr...)
}

// IdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IdxOf(e byte, arr ...byte) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// LastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func LastIdxOf(e byte, arr ...byte) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// MkSet : remove repeated elements in arr
func MkSet(arr ...byte) (set []byte) {
	if arr == nil {
		return nil
	}
	m := make(map[byte]struct{})
	for _, ele := range arr {
		if _, ok := m[ele]; !ok {
			set = append(set, ele)
			m[ele] = struct{}{}
		}
	}
	if len(set) == 0 {
		return []byte{}
	}
	return
}

// Superset :
func Superset(setA, setB []byte) bool {
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
func Subset(setA, setB []byte) bool {
	return Superset(setB, setA)
}

// equal :
func equal(setA, setB []byte) bool {
	if (setA == nil && setB != nil) || (setA != nil && setB == nil) {
		return false
	}
	if len(setA) != len(setB) {
		return false
	}

	tmpA := make([]byte, len(setA))
	tmpB := make([]byte, len(setB))
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
func Equal(sets ...[]byte) bool {
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
func SuperEq(setA, setB []byte) bool {
	return Superset(setA, setB) || Equal(setA, setB)
}

// SubEq :
func SubEq(setA, setB []byte) bool {
	return Subset(setA, setB) || Equal(setA, setB)
}

// union :
func union(setA, setB []byte) (set []byte) {
	if setA == nil && setB == nil {
		return nil
	}
	if setA == nil && setB != nil {
		return setB
	}
	if setA != nil && setB == nil {
		return setA
	}

	m := make(map[byte]struct{})
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
		return []byte{}
	}
	return
}

// Union :
func Union(sets ...[]byte) (set []byte) {
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
func intersect(setA, setB []byte) (set []byte) {
	if setA == nil || setB == nil {
		return nil
	}

	copyA, copyB := make([]byte, len(setA)), make([]byte, len(setB))
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
		return []byte{}
	}
	return
}

// Intersect :
func Intersect(sets ...[]byte) (set []byte) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = intersect(set, s)
	}
	return set
}

func minus(setA, setB []byte) (set []byte) {
	if setA == nil {
		return nil
	}
	set = make([]byte, 0)

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

func Minus(setA []byte, setOthers ...[]byte) (set []byte) {
	return minus(setA, Union(setOthers...))
}

// Reorder : any index must less than len(arr)
func Reorder(arr []byte, indices []int) (orders []byte) {
	if arr == nil || indices == nil {
		return nil
	}
	if len(arr) == 0 || len(indices) == 0 {
		return []byte{}
	}
	for _, idx := range indices {
		orders = append(orders, arr[idx])
	}
	return orders
}

// Reverse : [1,2,3] => [3,2,1]
func Reverse(arr []byte) []byte {
	indices := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		indices[i] = len(arr) - 1 - i
	}
	return Reorder(arr, indices)
}

// Reduce :
func Reduce(arr []byte, reduce func(e0, e1 byte) byte) byte {
	switch len(arr) {
	case 0, 1:
		panic("Reduce at least receives 2 parameters")
	default:
		var r byte
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
func ZipArray(arrays ...[]byte) (zipped [][]byte) {

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
		tuple := []byte{}
		for _, arr := range arrays {
			tuple = append(tuple, arr[i])
		}
		zipped = append(zipped, tuple)
	}
	return
}

func Filter(data *[]byte, check func(i int, e byte) bool) []byte {
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
func MapFilter(m map[byte]byte, filter func(k byte, v byte) bool) map[byte]byte {
	rt := make(map[byte]byte)
	for k, v := range m {
		if filter(k, v) {
			rt[k] = v
		}
	}
	return rt
}

// MapCopy:
func MapCopy(m map[byte]byte) map[byte]byte {
	return MapFilter(m, func(k byte, v byte) bool { return true })
}
