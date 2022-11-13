package i64

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		data  *[]int
		check func(i int, e int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				data: &[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				check: func(i, e int) bool {
					return e%2 == 0
				},
			},
			want: []int{0, 2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.data, tt.args.check); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
	type args struct {
		arr      []int
		filter   func(i int, e int) bool
		modifier func(i int, e int) int
	}
	tests := []struct {
		name  string
		args  args
		wantR []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				filter: func(i, e int) bool {
					return e%2 == 0
				},
				modifier: nil,
			},
			wantR: []int{0, 2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := FilterMap(tt.args.arr, tt.args.filter, tt.args.modifier); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("FilterMap() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func buildOrginalData() []int {
	s := make([]int, 1024*1024)
	for i := range s {
		s[i] = i
	}
	return s
}

func Benchmark_Filter(b *testing.B) {
	data := buildOrginalData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(&data, func(i, e int) bool { return e%2 == 0 })
	}
}

func Benchmark_FilterMap(b *testing.B) {
	data := buildOrginalData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterMap(data, func(i, e int) bool { return e%2 == 0 }, nil)
	}
}
