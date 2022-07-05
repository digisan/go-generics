package v2

import (
	"fmt"
	"testing"
)

func TestMapCvt(t *testing.T) {
	m := map[any]any{
		1: "11",
		2: "22",
	}
	m1 := MapCvt[int, string](m)
	// m1 := m.(map[int]string)
	fmt.Println(m1[1] + m1[2])
}

func TestSlcCvt(t *testing.T) {
	s := []any{1, 2, 3, 4, 5}
	s1 := SlcCvt[int](s)
	// s1 := s.([]int)
	fmt.Println(Sum(s1...))
}
