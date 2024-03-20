package gogenerics

import (
	"fmt"
	"testing"
)

func TestSetNestedSlice(t *testing.T) {

	slc := []any{[]any{2, []any{3, 4}}, 1, []any{"a"}, 3, 8}
	fmt.Println(slc)
	fmt.Println()

	paths := TraverseSlice(slc)
	fmt.Println(paths)

	cap := CapacityForSlice(paths...)
	fmt.Println(cap)

	dest := InitSlice(cap...)
	fmt.Println(dest)

	paths = TraverseSlice(dest)
	fmt.Println(paths)

	ok := SetSlice(dest, "OK", 2, 1, 0)
	fmt.Println(ok)
	fmt.Println(dest)

	ok = SetSlice(slc, "OKey", 0, 1, 1)
	fmt.Println(ok)
	fmt.Println(slc)
}

func TestInitNestedSlice(t *testing.T) {

	slc := InitSlice(5, 2, 2, 1)
	fmt.Printf("%+v\n", slc)

	if !SetSlice(slc, "OK0", 2, 1, 0, 0) {
		panic("Error SetSlice OK0")
	}
	if !SetSlice(slc, "OK1", 1, 0, 1, 0) {
		panic("Error SetSlice OK1")
	}

	fmt.Printf("%+v\n", slc)
}
