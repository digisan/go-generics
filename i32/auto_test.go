package i32

import (
	"fmt"
	"testing"
)

func TestDelEleOrderly(t *testing.T) {
	type args struct {
		arr *[]rune
		i   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				arr: &[]rune{1, 2, 3, 4},
				i:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelEleOrderly(tt.args.arr, tt.args.i)
			fmt.Println(*tt.args.arr)
		})
	}
}

func TestDelEle(t *testing.T) {
	type args struct {
		arr *[]rune
		i   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				arr: &[]rune{1, 2, 3, 4},
				i:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelEle(tt.args.arr, tt.args.i)
			fmt.Println(*tt.args.arr)
		})
	}
}

func TestMapCopy(t *testing.T) {
	m := map[rune]rune{
		'A': 'a',
		'B': 'b',
	}
	fmt.Println("m original:", m)

	fmt.Println("copying... to cp")
	cp := MapCopy(m)

	fmt.Println("modifying m...")
	m['B'] = 'B'
	fmt.Println("m modified:", m)

	fmt.Println("cp", cp)
}
