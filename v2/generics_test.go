package v2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppendBytes(t *testing.T) {
	fmt.Println(string(AppendBytes([]byte{49, 50, 51}, []byte{58, 36, 73, 100})))
}

func TestCount(t *testing.T) {
	arr := []int{1, 2, 3, 4, 1}
	fmt.Println(Count(arr, 1))
	fmt.Println(Count(arr, 3))
	fmt.Println(Count(arr, 0))
}

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

func TestFilterMap(t *testing.T) {

	// arr := []int{1, 2, 3, 4, 5}
	// out := FilterMap(arr, func(i int, e int) bool { return e > 3 }, func(i int, e int) int { return e })
	// fmt.Println("arr:", arr)
	// fmt.Println("out:", out)
	// return

	type args struct {
		arr    []int
		filter func(i int, e int) bool
		mapper func(i int, e int) string
	}
	tests := []struct {
		name  string
		args  args
		wantR []string
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 6, 7},
				filter: func(i, e int) bool { return e > 3 },
				mapper: func(i, e int) string { return fmt.Sprint(e) },
			},
			wantR: []string{"4", "5", "6", "7"},
		},
		{
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 6, 7},
				filter: nil,
				mapper: func(i, e int) string { return fmt.Sprint(e) },
			},
			wantR: []string{"1", "2", "3", "4", "5", "6", "7"},
		},
		// {
		// 	args: args{
		// 		arr:    []int{1, 2, 3, 4, 5, 6, 7},
		// 		filter: func(i, e int) bool { return e > 3 },
		// 		mapper: nil, // panic
		// 	},
		// 	wantR: []string{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := FilterMap(tt.args.arr, tt.args.filter, tt.args.mapper); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("FilterMap() = %v, want %v", gotR, tt.wantR)
			}
		})
		fmt.Println(tt.args.arr)
	}
}

func TestFilterMap4SglTyp(t *testing.T) {
	type args struct {
		arr    []int
		filter func(i int, e int) bool
		mapper func(i int, e int) int
	}
	tests := []struct {
		name  string
		args  args
		wantR []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 6, 7},
				filter: func(i, e int) bool { return e > 3 },
				mapper: func(i, e int) int { return i },
			},
			wantR: []int{3, 4, 5, 6},
		},
		{
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 6, 7},
				filter: func(i, e int) bool { return e > 3 },
				mapper: nil,
			},
			wantR: []int{4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := FilterMap4SglTyp(tt.args.arr, tt.args.filter, tt.args.mapper); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("FilterMap4SglTyp() = %v, want %v", gotR, tt.wantR)
			}
		})
		// fmt.Println(tt.args.arr)
	}
}

func TestFilterFast(t *testing.T) {
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
			args: args{
				data:  &[]int{6, 1, 5, 3, 4, 2},
				check: func(i, e int) bool { return e < 4 },
			},
			want: []int{1, 3, 2, 6, 4, 5},
		},
		{
			args: args{
				data:  &[]int{6, 1, 5, 3, 4, 2},
				check: nil,
			},
			want: []int{6, 1, 5, 3, 4, 2},
		},
		{
			args: args{
				data:  &[]int{6, 1, 5, 3, 4, 2},
				check: func(i, e int) bool { return e < 10 },
			},
			want: []int{6, 1, 5, 3, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterFast(tt.args.data, tt.args.check); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
		fmt.Println("filtered:", *tt.args.data) // original data is changed, let's have a look
		fmt.Println()
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		data  []int
		check func(i int, e int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				data:  []int{1, 5, 3, 4, 2},
				check: func(i, e int) bool { return e < 4 },
			},
			want: []int{1, 3, 2},
		},
		{
			args: args{
				data:  []int{1, 5, 3, 4, 2},
				check: nil,
			},
			want: []int{1, 5, 3, 4, 2},
		},
		{
			args: args{
				data:  []int{1, 5, 3, 4, 2},
				check: func(i, e int) bool { return e < 10 },
			},
			want: []int{1, 5, 3, 4, 2},
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

func TestMap(t *testing.T) {
	type args struct {
		arr    []int
		mapper func(i int, e int) string
	}
	tests := []struct {
		name  string
		args  args
		wantR []string
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				mapper: func(i, e int) string { return fmt.Sprint(e) },
			},
			wantR: []string{"1", "2", "3", "4", "5"},
		},
		// {
		// 	args: args{
		// 		arr:    []int{1, 2, 3, 4, 5},
		// 		mapper: nil, // panic
		// 	},
		// 	wantR: []string{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Map(tt.args.arr, tt.args.mapper); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Map() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func Test_equals(t *testing.T) {
	type args struct {
		setA []int
		setB []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				setA: []int{1, 2, 3, 4},
				setB: []int{3, 1, 2, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equals(tt.args.setA, tt.args.setB); got != tt.want {
				t.Errorf("equals() = %v, want %v", got, tt.want)
			}
		})
		fmt.Println("SetA", tt.args.setA)
		fmt.Println("SetB", tt.args.setB)
	}
}

func TestEquals(t *testing.T) {
	type args struct {
		sets [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				sets: [][]int{
					{1, 2, 3, 4},
					{3, 1, 2, 4},
					{3, 1, 2, 4, 5},
				},
			},
			want: false,
		},
		{
			args: args{
				sets: [][]int{
					{1, 2, 3, 4},
					{3, 1, 2, 4},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.sets...); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
		fmt.Println("sets", tt.args.sets)
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		arr    []int
		reduce func(e0, e1 int) int
	}
	tests := []struct {
		name  string
		args  args
		wantR int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				reduce: func(e0, e1 int) int { return e0 + e1 },
			},
			wantR: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Reduce(tt.args.arr, tt.args.reduce); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Reduce() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestIn(t *testing.T) {
	type args struct {
		e   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				e:   10,
				arr: []int{3, 5, 2, 10, 2, 5, 0, 7},
			},
			want: true,
		},
		{
			args: args{
				e:   10,
				arr: []int{3, 5, 2, 11, 2, 5, 0, 7},
			},
			want: false,
		},
		{
			args: args{
				e:   10,
				arr: []int{},
			},
			want: false,
		},
		{
			args: args{
				e:   10,
				arr: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.args.e, tt.args.arr...); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotIn(t *testing.T) {
	type args struct {
		e   string
		arr []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				e:   "10",
				arr: []string{"3", "5", "2", "10", "2", "5", "0", "7"},
			},
			want: false,
		},
		{
			args: args{
				e:   "10",
				arr: []string{"3", "5", "2", "11", "2", "5", "0", "7"},
			},
			want: true,
		},
		{
			args: args{
				e:   "",
				arr: []string{},
			},
			want: true,
		},
		{
			args: args{
				e:   "",
				arr: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotIn(tt.args.e, tt.args.arr...); got != tt.want {
				t.Errorf("NotIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZipSlice(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantZipped [][]int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arrays: [][]int{
					{1, 2},
					{3, 4, 5},
					{6, 7, 8, 9},
				},
			},
			wantZipped: [][]int{
				{1, 3, 6}, {2, 4, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZipped := ZipSlice(tt.args.arrays...); !reflect.DeepEqual(gotZipped, tt.wantZipped) {
				t.Errorf("ZipSlice() = %v, want %v", gotZipped, tt.wantZipped)
			}
		})
	}
}

func TestZipDim2(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantZipped [][2]int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arrays: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantZipped: [][2]int{
				{1, 4}, {2, 5}, {3, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZipped := ZipDim2(tt.args.arrays...); !reflect.DeepEqual(gotZipped, tt.wantZipped) {
				t.Errorf("ZipDim2() = %v, want %v", gotZipped, tt.wantZipped)
			}
		})
	}
}

func TestZipDim3(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantZipped [][3]int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arrays: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantZipped: [][3]int{
				{1, 4, 7}, {2, 5, 8}, {3, 6, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZipped := ZipDim3(tt.args.arrays...); !reflect.DeepEqual(gotZipped, tt.wantZipped) {
				t.Errorf("ZipDim3() = %v, want %v", gotZipped, tt.wantZipped)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr: []int{},
			},
			want: []int{},
		},
		{
			args: args{
				arr: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReorder(t *testing.T) {
	type args struct {
		arr     []int
		indices []int
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr:     []int{4, 2, 3, 1},
				indices: []int{2, 1, 3, 0},
			},
			wantOrders: []int{3, 2, 1, 4},
		},
		{
			args: args{
				arr:     []int{4, 2, 3, 1},
				indices: []int{2, 1},
			},
			wantOrders: []int{3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrders := Reorder(tt.args.arr, tt.args.indices); !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("Reorder() = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args struct {
		idx int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				idx: 1,
				arr: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			args: args{
				idx: 2,
				arr: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Last(tt.args.arr, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeArray(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantMerged []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arrays: [][]int{{2, 2}, {3, 3, 4}, {6, 6, 8, 8}},
			},
			wantMerged: []int{2, 2, 3, 3, 4, 6, 6, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMerged := SmashArrays(tt.args.arrays...); !reflect.DeepEqual(gotMerged, tt.wantMerged) {
				t.Errorf("SmashArrays() = %v, want %v", gotMerged, tt.wantMerged)
			}
		})
	}
}

func TestMergeSet(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantMerged []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				arrays: [][]int{{2, 2}, {3, 3, 4}, {6, 6, 8, 8}},
			},
			wantMerged: []int{2, 3, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMerged := SmashSets(tt.args.arrays...); !reflect.DeepEqual(gotMerged, tt.wantMerged) {
				t.Errorf("SmashSets() = %v, want %v", gotMerged, tt.wantMerged)
			}
		})
	}
}

func TestAppendIf(t *testing.T) {
	type args struct {
		ok    bool
		arr   []int
		elems []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				ok:    true,
				arr:   []int{1, 2, 3},
				elems: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			args: args{
				ok:    false,
				arr:   []int{1, 2, 3},
				elems: []int{4, 5, 6},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendIf(tt.args.ok, tt.args.arr, tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllSameEle(t *testing.T) {
	arrBool := []bool{true, true}
	fmt.Println(AllSameEle(arrBool...))
	fmt.Println(AllSameEleAs(arrBool, false))
	fmt.Println(AllSameEleAs(arrBool, true))

	arrBool = []bool{}
	fmt.Println(AllSameEle(arrBool...))
	fmt.Println(AllSameEleAs(arrBool, false))
	fmt.Println(AllSameEleAs(arrBool, true))

	arrBool = nil
	fmt.Println(AllSameEle(arrBool...))
	fmt.Println(AllSameEleAs(arrBool, false))
	fmt.Println(AllSameEleAs(arrBool, true))

	arrFloat := []float64{0, 0, 0, 0}
	fmt.Println(AllSameEle(arrFloat...))
	fmt.Println(AllSameEleAs(arrFloat, 0))
	fmt.Println(AllSameEleAs(arrFloat, 0.0))
}

func TestMaxCeiling(t *testing.T) {
	type args struct {
		ceiling int
		arr     []int
	}
	tests := []struct {
		name   string
		args   args
		wantM  int
		wantOk bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				ceiling: 10,
				arr:     []int{3, 5, 33, 2, 1, 12, 66},
			},
			wantM:  5,
			wantOk: true,
		},
		{
			args: args{
				ceiling: -1,
				arr:     []int{3, 5, 33, 2, 1, 12, 66},
			},
			wantM:  3,
			wantOk: false,
		},
		{
			args: args{
				ceiling: 60,
				arr:     []int{60, 3, 5, 33, 2, 1, 12, 66, 59},
			},
			wantM:  59,
			wantOk: true,
		},
		{
			args: args{
				ceiling: 50,
				arr:     []int{60},
			},
			wantM:  60,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, gotOk := MaxCeiling(tt.args.ceiling, false, tt.args.arr...)
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("MaxCeiling() gotM = %v, want %v", gotM, tt.wantM)
			}
			if gotOk != tt.wantOk {
				t.Errorf("MaxCeiling() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMinFloor(t *testing.T) {
	type args struct {
		floor int
		arr   []int
	}
	tests := []struct {
		name   string
		args   args
		wantM  int
		wantOk bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				floor: 10,
				arr:   []int{3, 5, 33, 2, 1, 12, 66, 10},
			},
			wantM:  10,
			wantOk: true,
		},
		{
			args: args{
				floor: 60,
				arr:   []int{3, 5, 33, 2, 1, 12, 66},
			},
			wantM:  66,
			wantOk: true,
		},
		{
			args: args{
				floor: 90,
				arr:   []int{3, 5, 33, 2, 1, 12, 66},
			},
			wantM:  3,
			wantOk: false,
		},
		{
			args: args{
				floor: 4,
				arr:   []int{3},
			},
			wantM:  3,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, gotOk := MinFloor(tt.args.floor, true, tt.args.arr...)
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("MinFloor() gotM = %v, want %v", gotM, tt.wantM)
			}
			if gotOk != tt.wantOk {
				t.Errorf("MinFloor() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
