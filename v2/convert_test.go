package v2

import (
	"fmt"
	"sync"
	"testing"
)

func TestMapCvt(t *testing.T) {
	m := map[any]any{
		1: "11",
		2: "22",
	}
	m1 := MapAnyToType[int, string](m)
	// m1 := m.(map[int]string)
	fmt.Println(m1[1] + m1[2])
}

func TestSyncMap2Map(t *testing.T) {
	sm := sync.Map{}
	sm.Store(1, "11")
	sm.Store(2, "22")
	m := SyncMapToMap[int, string](sm)
	fmt.Println(m[1] + m[2])
}

func TestMap2SyncMap(t *testing.T) {
	m := map[int]string{
		1: "11",
		2: "22",
	}
	sm := MapToSyncMap(m)
	sm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

func TestSlcCvt(t *testing.T) {
	s := []any{1, 2, 3, 4, 5}
	s1 := AnysToTypes[int](s)
	// s1 := s.([]int)
	fmt.Println(Sum(s1...))
}

func TestSlc2Types(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	s1 := SlcToTypes[int](s)
	fmt.Println(Sum(s1...))

	s2 := Nums2Floats(s1...)
	fmt.Println(Sum(s2...))

	// s2 := s.([]int)
	// fmt.Println(Sum(s2...))
}

func TestSlc2Anys(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	gs1 := SlcToAnys(s1)
	fmt.Println(gs1)

	s2 := []string{"a1", "b2", "c3", "d4", "e5", "f6", "g7", "h8"}
	gs2 := SlcToAnys(s2)
	fmt.Println(gs2)
}
