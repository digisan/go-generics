package v2

import (
	"reflect"
	"testing"
)

func TestMaxMin(t *testing.T) {
	{
		type args struct {
			arr []float32
		}
		tests := []struct {
			name string
			args args
			want float32
		}{
			// TODO: Add test cases.
			{
				name: "OK",
				args: args{
					arr: []float32{3, 1, 2, 6, 7, 4, 5, 2, 5},
				},
				want: 7,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Max(tt.args.arr...); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Max() = %v, want %v", got, tt.want)
				}
			})
		}
	}
	{
		type args struct {
			arr []int
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			// TODO: Add test cases.
			{
				name: "OK",
				args: args{
					arr: []int{3, 1, 2, 6, 7, 4, 5, 2, 5},
				},
				want: 1,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Min(tt.args.arr...); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Min() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
