package v2

import (
	"fmt"
	"testing"
)

func BenchmarkTypeOf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		number := 3987
		TypeOf(number)
	}
}

func TestTypeOf(t *testing.T) {
	number := 123
	p := &number
	fmt.Println(TypeOf(number))
	fmt.Println(TypeOf(&p))

	obj := struct{ Z int }{Z: 100}
	fmt.Println(TypeOf(obj))  // struct { Z int }
	fmt.Println(KindOf(obj))  // struct
	fmt.Println(KindOf(&obj)) // ptr

	slc := []int{}
	fmt.Println(TypeOf(slc))
	fmt.Println(KindOf(slc))

	arr := [3]int{}
	fmt.Println(TypeOf(arr))
	fmt.Println(KindOf(arr))

	m := map[string]any{}
	fmt.Println(TypeOf(m))
	fmt.Println(KindOf(m))
}
