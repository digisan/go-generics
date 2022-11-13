package i32

import (
	"fmt"
	"testing"
)

func TestDelOneEleOrderly(t *testing.T) {
	type args struct {
		arr *[]rune
		i   rune
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				arr: &[]rune{1, 2, 0, 3, 0, 4},
				i:   1,
			},
		},
		{
			name: "OK",
			args: args{
				arr: &[]rune{1},
				i:   1,
			},
		},
		{
			name: "OK",
			args: args{
				arr: &[]rune{},
				i:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelOneEle(tt.args.arr, tt.args.i)
			fmt.Println(*tt.args.arr)
		})
	}
}

func TestDelEleOrderlyAt(t *testing.T) {
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
			DelEleOrderlyAt(tt.args.arr, tt.args.i)
			fmt.Println(*tt.args.arr)
		})
	}
}

func TestDelEleAt(t *testing.T) {
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
			DelEleAt(tt.args.arr, tt.args.i)
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
