package v2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxMin(t *testing.T) {
	{
		type f32 float32

		type args struct {
			arr []f32
		}
		tests := []struct {
			name string
			args args
			want f32
		}{
			// TODO: Add test cases.
			{
				name: "OK",
				args: args{
					arr: []f32{3, 1, 2, 6, 7, 4, 5, 2, 5},
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

func TestMaxMinIdx(t *testing.T) {
	{
		type args struct {
			arr []int
		}
		tests := []struct {
			name  string
			args  args
			want  int
			want1 int
		}{
			// TODO: Add test cases.
			{
				name: "OK",
				args: args{
					arr: []int{3},
				},
				want:  3,
				want1: 0,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, got1 := MaxIdx(tt.args.arr...)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("MaxIdx() got = %v, want %v", got, tt.want)
				}
				if got1 != tt.want1 {
					t.Errorf("MaxIdx() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	}
	{
		type args struct {
			arr []float64
		}
		tests := []struct {
			name  string
			args  args
			want  float64
			want1 int
		}{
			// TODO: Add test cases.
			{
				name: "OK",
				args: args{
					arr: []float64{3, 1, 2, 6, 7, 4, 5, 2, 5},
				},
				want:  1,
				want1: 1,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, got1 := MinIdx(tt.args.arr...)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("MinIdx() got = %v, want %v", got, tt.want)
				}
				if got1 != tt.want1 {
					t.Errorf("MinIdx() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	}
}

func TestIdxOf(t *testing.T) {
	{
		type args struct {
			e   int
			arr []int
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			// TODO: Add test cases.
			{
				name: "",
				args: args{
					e:   3,
					arr: []int{2, 6, 1, 4, 3, 2, 3, 7, 9},
				},
				want: 4,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := IdxOf(tt.args.e, tt.args.arr...); got != tt.want {
					t.Errorf("IdxOf() = %v, want %v", got, tt.want)
				}
			})
		}
	}
	{
		type args struct {
			e   int
			arr []int
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			// TODO: Add test cases.
			{
				name: "",
				args: args{
					e:   3,
					arr: []int{2, 6, 1, 4, 3, 2, 3, 7, 9},
				},
				want: 6,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := LastIdxOf(tt.args.e, tt.args.arr...); got != tt.want {
					t.Errorf("LastIdxOf() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestDelEleOrderlyAt(t *testing.T) {
	type args struct {
		arr *[]string
		i   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arr: &[]string{"a", "b", "c", "d", "e"},
				i:   4,
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
		arr *[]string
		i   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arr: &[]string{"a", "b", "c", "d", "e"},
				i:   4,
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

func TestDelOneEle(t *testing.T) {
	type args struct {
		arr *[]string
		ele string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arr: &[]string{"a", "b", "c", "d", "e"},
				ele: "b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelOneEle(tt.args.arr, tt.args.ele)
			fmt.Println(*tt.args.arr)
		})
	}
}

func TestDelOneEleOrderly(t *testing.T) {
	type args struct {
		arr *[]string
		ele string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arr: &[]string{"a", "b", "c", "d", "e"},
				ele: "b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelOneEleOrderly(tt.args.arr, tt.args.ele)
			fmt.Println(*tt.args.arr)
		})
	}
}

func TestSettify(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name    string
		args    args
		wantSet []string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arr: []string{"a", "b", "a", "c", "d", "c", "e", "b"},
			},
			wantSet: []string{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := Settify(tt.args.arr...); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("Settify() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func TestMapFilter(t *testing.T) {
	type args struct {
		m      map[int]string
		filter func(k int, v string) bool
	}
	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		// TODO: Add test cases.
		{
			args: args{
				m: map[int]string{
					0: "0",
					1: "1",
					2: "2",
					3: "3",
					4: "4",
				},
				filter: func(k int, v string) bool {
					return k >= 3
				},
			},
			want: map[int]string{
				3: "3",
				4: "4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapFilter(tt.args.m, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
