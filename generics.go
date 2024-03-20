package gogenerics

import (
	"bytes"
	"encoding/gob"
	"log"
	"unsafe"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64 | ~string
}

func AppendBytes(blocks ...[]byte) []byte {
	return bytes.Join(blocks, []byte{})
}

func AppendIf[T any](ok bool, arr []T, elems ...T) []T {
	if ok {
		return append(arr, elems...)
	}
	return arr
}

func Count[T comparable](arr []T, elem T) (n int) {
	if len(arr) == 0 {
		return 0
	}
	for _, a := range arr {
		if a == elem {
			n++
		}
	}
	return
}

func AllSameEle[T comparable](arr ...T) bool {
	if len(arr) == 0 {
		return false
	}
	for i := 1; i < len(arr); i++ {
		if arr[0] != arr[i] {
			return false
		}
	}
	return true
}

func AllSameEleAs[T comparable](arr []T, e T) bool {
	return AllSameEle(arr...) && arr[0] == e
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

func MaxCeiling[T Ordered](ceiling T, close bool, arr ...T) (m T, ok bool) {
	if len(arr) == 0 {
		panic("Max args at least has one element")
	}
	var idx int
	m = arr[0]
	if close {
		for i, a := range arr {
			if a <= ceiling {
				idx, m = i, a
				ok = true
				break
			}
		}
		for _, a := range arr[idx:] {
			if a > m && a <= ceiling {
				m = a
			}
		}
	} else {
		for i, a := range arr {
			if a < ceiling {
				idx, m = i, a
				ok = true
				break
			}
		}
		for _, a := range arr[idx:] {
			if a > m && a < ceiling {
				m = a
			}
		}
	}
	return
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

func MinFloor[T Ordered](floor T, close bool, arr ...T) (m T, ok bool) {
	if len(arr) == 0 {
		panic("Min args at least has one element")
	}
	var idx int
	m = arr[0]
	if close {
		for i, a := range arr {
			if a >= floor {
				idx, m = i, a
				ok = true
				break
			}
		}
		for _, a := range arr[idx:] {
			if a < m && a >= floor {
				m = a
			}
		}
	} else {
		for i, a := range arr {
			if a > floor {
				idx, m = i, a
				ok = true
				break
			}
		}
		for _, a := range arr[idx:] {
			if a < m && a > floor {
				m = a
			}
		}
	}
	return
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

func InSlices[T comparable](e T, arr ...[]T) bool {
	for _, a := range arr {
		if In(e, a...) {
			return true
		}
	}
	return false
}

func NotInSlices[T comparable](e T, arr ...[]T) bool {
	return !InSlices(e, arr...)
}

func OnlyIn2[T comparable](arr1, arr2 []T) (elems []T) {
	for _, a2 := range arr2 {
		if NotIn(a2, arr1...) {
			elems = append(elems, a2)
		}
	}
	return
}

func OnlyIn1[T comparable](arr1, arr2 []T) (elems []T) {
	for _, a1 := range arr1 {
		if NotIn(a1, arr2...) {
			elems = append(elems, a1)
		}
	}
	return
}

func SameElements[T comparable](arr1, arr2 []T) (same []T) {
	for _, a1 := range arr1 {
		if In(a1, arr2...) {
			same = append(same, a1)
		}
	}
	return
}

func CanCover[T comparable](arr1, arr2 []T) bool {
	return len(OnlyIn2(arr1, arr2)) == 0
}

// ***
func DelEleOrderlyAt[T comparable](arr *[]T, i int) {
	if i >= 0 && i < len(*arr) {
		*arr = append((*arr)[:i], (*arr)[i+1:]...)
	}
}

// ***
func DelEleAt[T comparable](arr *[]T, i int) {
	if length := len(*arr); i >= 0 && i < length {
		(*arr)[i] = (*arr)[len(*arr)-1]
		// (*reflect.SliceHeader)(unsafe.Pointer(arr)).Len--
		*arr = unsafe.Slice(unsafe.SliceData(*arr), length-1)
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
func FilterFast[T any](pData *[]T, filter func(i int, e T) bool) []T {
	if filter == nil {
		return *pData
	}

	p := *pData
	var k = 0
	for i, v := range p {
		if filter(i, v) {
			p[k], p[i] = p[i], p[k]
			k++
		}
	}
	// (*reflect.SliceHeader)(unsafe.Pointer(pData)).Len = k
	*pData = unsafe.Slice(unsafe.SliceData(*pData), k)
	return p
}

// *** input data keeps original, return filtered & ordered copy
func Filter[T any](data []T, filter func(i int, e T) bool) []T {
	if filter == nil {
		return append([]T{}, data...)
	}

	r := make([]T, 0, len(data))
	for i, v := range data {
		if filter(i, v) {
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

// FilterMap : Filter A slice, return A=>B slice. B could be different type from A
func FilterMap[T1, T2 any](arr []T1, filter func(i int, e T1) bool, mapper func(i int, e T1) T2) (r []T2) {
	if mapper == nil {
		panic("mapper cannot be nil")
	}

	r = make([]T2, 0, len(arr))
	for i, e := range arr {
		if (filter != nil && filter(i, e)) || (filter == nil) {
			r = append(r, mapper(i, e))
		}
	}
	return
}

// Filter A slice, return A=>B slice. B is the same type as A
func FilterMap4SglTyp[T any](arr []T, filter func(i int, e T) bool, mapper func(i int, e T) T) (r []T) {
	r = make([]T, 0, len(arr))
	for i, e := range arr {
		if (filter != nil && filter(i, e)) || (filter == nil) {
			if mapper != nil {
				r = append(r, mapper(i, e))
			} else {
				r = append(r, e)
			}
		}
	}
	return
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

// *** SmashArrays : [{2,2}, {3,3,4}, {6,6,8,8}] => {2,2,3,3,4,6,6,8,8}
func SmashArrays[T any](arrays ...[]T) (merged []T) {
	for _, arr := range arrays {
		merged = append(merged, arr...)
	}
	return
}

// *** SmashSets : [{2,2}, {3,3,4}, {6,6,8,8}] => {2,3,4,6,8}
func SmashSets[T comparable](arrays ...[]T) (merged []T) {
	for _, arr := range arrays {
		merged = append(merged, arr...)
	}
	return Settify(merged...)
}

func Clone(dst, src any) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(src); err != nil {
		return err
	}
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}
