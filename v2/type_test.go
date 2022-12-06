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
}
