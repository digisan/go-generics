package v2

import (
	"fmt"
	"testing"
)

func TestIF(t *testing.T) {
	a := IF(3 > 2, 4, 6)
	fmt.Println(a)
	b := IF[any](3 > 5, 4, "6")
	fmt.Println(b)
}

func TestMATCH(t *testing.T) {
	a := MATCH("d", "a", "b", "c", "d", "A", "B", "C", "D", "default")
	fmt.Println(a)
	b := MATCH[any]("e", "a", "b", "c", "d", 1, 2, 3, 4, 100)
	fmt.Println(b)
}
