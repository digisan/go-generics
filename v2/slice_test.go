package v2

import (
	"fmt"
	"testing"
)

func TestSetNestedSlice(t *testing.T) {

	slc := []any{[]any{2, []any{3, 4}}, 1, []any{"a"}, 3, 8}
	fmt.Println(slc)
	fmt.Println()

	paths := TraverseNestedSlice(slc)
	fmt.Println(paths)

	cap := CapacityForSlice(paths...)
	fmt.Println(cap)

	dest := InitNestedSlice(cap...)
	fmt.Println(dest)

	paths = TraverseNestedSlice(dest)
	fmt.Println(paths)

	ok := SetNestedSlice(dest, "OK", 2, 1, 0)
	fmt.Println(ok)
	fmt.Println(dest)

	ok = SetNestedSlice(slc, "OKey", 0, 1, 1)
	fmt.Println(ok)
	fmt.Println(slc)
}

func TestInitNestedSlice(t *testing.T) {

	slc := InitNestedSlice(5, 2, 2)
	fmt.Println(slc)

	SetNestedSlice(slc, "OK0", 2, 1, 0)
	SetNestedSlice(slc, "OK1", 1, 0, 0)

	fmt.Println(slc)
}

// func TestMergeNestedSlice(t *testing.T) {
// 	slc1 := []any{[]any{300, nil}, 300}
// 	slc2 := []any{[]any{nil, 300}, nil, nil, nil, 300}
// 	fmt.Println(MergeTwoSlices(slc1, slc2))
// }
