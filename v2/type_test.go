package v2

import (
	"fmt"
	"io"
	"os"
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

func TestIsNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{"123"},
			want: true,
		},
		{
			name: "OK",
			args: args{".123"},
			want: true,
		},
		{
			name: "OK",
			args: args{"a123"},
			want: false,
		},
		{
			name: "OK",
			args: args{"000123"},
			want: true,
		},
		{
			name: "OK",
			args: args{"000123.0231"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.s); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContinuous(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name    string
		args    args
		wantOk  bool
		wantMin int
		wantMax int
	}{
		// TODO: Add test cases.
		{
			name:    "OK",
			args:    args{[]int{1, 2, 3, 4, 5}},
			wantOk:  true,
			wantMin: 1,
			wantMax: 5,
		},
		{
			name:    "OK",
			args:    args{[]int{1, 2, 3, 4, 6}},
			wantOk:  false,
			wantMin: 1,
			wantMax: 6,
		},
		{
			name:    "OK",
			args:    args{[]int{5, 4, 3, 2, 1}},
			wantOk:  true,
			wantMin: 1,
			wantMax: 5,
		},
		{
			name:    "OK",
			args:    args{[]int{6, 4, 3, 2, 1}},
			wantOk:  false,
			wantMin: 1,
			wantMax: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotMin, gotMax := IsContinuous(tt.args.ints...)
			if gotOk != tt.wantOk {
				t.Errorf("IsContinuous() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotMin != tt.wantMin {
				t.Errorf("IsContinuous() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("IsContinuous() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func checkInterface(i io.Reader) {
	fmt.Println("[i == nil]", i == nil)
	fmt.Println("[nil any]", IsNil(i))
}

func TestIsInterfaceNil(t *testing.T) {

	fmt.Println(IsNil(nil))
	fmt.Println("-----------------")

	file, err := os.Open("../go.mod")
	if err != nil {
		panic(err)
	}
	checkInterface(file)

	fmt.Println("-----------------")
	file = nil
	fmt.Println("[file == nil]", file == nil)
	checkInterface(file)

	fmt.Println("-----------------")
	var file1 *os.File = nil
	fmt.Println("[file1 == nil]", file1 == nil)
	checkInterface(file1)
}

func TestIsNil(t *testing.T) {
	var p1 *int
	var p2 *float64
	var p any

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p)

	fmt.Println(fmt.Sprint(p1) == "<nil>")
	fmt.Println(fmt.Sprint(p2) == "<nil>")
	fmt.Println(fmt.Sprint(p) == "<nil>")

	f := func(a1, a2 any) bool {
		return a1 == a2
	}

	fIsNil := func(a any) bool {
		return a == nil
	}

	fmt.Println("p1 == nil", p1 == nil)
	fmt.Println("p2 == nil", p2 == nil)

	// fmt.Println("p1 == p2", p1 == p2)
	fmt.Println("p1 == p2 in func(any, any)", f(p1, p2))

	fmt.Println("p1 == nil in func(any)", fIsNil(p1))
	fmt.Println("p2 == nil in func(any)", fIsNil(p2))

	fmt.Println("IsNil p1", IsNil(p1))
	fmt.Println("IsNil p2", IsNil(p2))
}

func TestIsEmail(t *testing.T) {
	fmt.Println(IsEmailFormat("abdf@asd.com"))
}
