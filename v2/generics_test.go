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

func TestMap2KVs(t *testing.T) {
	type args struct {
		m          map[int]string
		less4key   func(i, j int) bool
		less4value func(i, j string) bool
	}
	tests := []struct {
		name       string
		args       args
		wantKeys   []int
		wantValues []string
	}{
		// TODO: Add test cases.
		{
			args: args{
				m: map[int]string{
					6: "6",
					1: "1",
					2: "2",
					3: "A",
					4: "A",
					5: "5",
				},
				less4key:   func(i, j int) bool { return i > j },
				less4value: func(i, j string) bool { return i < j },
			},
			wantKeys:   []int{1, 2, 5, 6, 4, 3},
			wantValues: []string{"1", "2", "5", "6", "A", "A"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeys, gotValues := Map2KVs(tt.args.m, tt.args.less4key, tt.args.less4value)
			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("Map2KVs() gotKeys = %v, want %v", gotKeys, tt.wantKeys)
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("Map2KVs() gotValues = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestMapMerge(t *testing.T) {
	type args struct {
		ms []map[int]string
	}
	tests := []struct {
		name string
		args args
		want map[int][]string
	}{
		// TODO: Add test cases.
		{
			args: args{
				ms: []map[int]string{
					{
						6: "66",
						1: "11",
						2: "22",
						3: "A",
						4: "A",
						5: "55",
					},
					{
						6: "66",
						1: "11",
						2: "22",
						3: "B",
						4: "B",
						5: "55",
					},
					{
						6: "6",
						1: "1",
						2: "2",
						3: "A",
						4: "A",
						5: "5",
					},
				},
			},
			want: map[int][]string{
				6: {"66", "6"},
				1: {"11", "1"},
				2: {"22", "2"},
				3: {"A", "B"},
				4: {"A", "B"},
				5: {"55", "5"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapMerge(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
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
		{
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 6, 7},
				filter: func(i, e int) bool { return e > 3 },
				mapper: nil,
			},
			wantR: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := FilterMap(tt.args.arr, tt.args.filter, tt.args.mapper); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("FilterMap() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

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
			args: args{
				data:  &[]int{1, 2, 3, 4, 5},
				check: func(i, e int) bool { return e < 4 },
			},
			want: []int{1, 2, 3},
		},
		{
			args: args{
				data:  &[]int{1, 2, 3, 4, 5},
				check: nil,
			},
			want: []int{1, 2, 3, 4, 5},
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
		{
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				mapper: nil,
			},
			wantR: []string{},
		},
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

func TestMapSafeMerge(t *testing.T) {
	type args struct {
		ms []map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
		{
			args: args{
				ms: []map[string]int{
					{
						"a": 1,
						"b": 2,
						"c": 3,
					},
					{
						"d": 4,
						"b": 5,
						"f": 6,
					},
				},
			},
			want: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
				"f": 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapSafeMerge(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapSafeMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapReplaceMerge(t *testing.T) {
	type args struct {
		ms []map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
		{
			args: args{
				ms: []map[string]int{
					{
						"a": 1,
						"b": 2,
						"c": 3,
					},
					{
						"d": 4,
						"b": 5,
						"f": 6,
					},
				},
			},
			want: map[string]int{
				"a": 1,
				"b": 5,
				"c": 3,
				"d": 4,
				"f": 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapReplaceMerge(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapReplaceMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapCopy(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
		{
			args: args{
				m: map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			want: map[string]int{
				"b": 2,
				"c": 3,
				"a": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapCopy(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZipArray(t *testing.T) {
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
			if gotZipped := ZipArray(tt.args.arrays...); !reflect.DeepEqual(gotZipped, tt.wantZipped) {
				t.Errorf("ZipArray() = %v, want %v", gotZipped, tt.wantZipped)
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
