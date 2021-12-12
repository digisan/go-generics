package i32

import (
	"fmt"
	"testing"
)

func TestDelEleOrderly(t *testing.T) {
	type args struct {
		arr *[]int
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
				arr: &[]int{1, 2, 3, 4},
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
		arr *[]int
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
				arr: &[]int{1, 2, 3, 4},
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
