package v2

import (
	"fmt"
	"reflect"
	"testing"
)

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

func TestMapMerge4ValSlc(t *testing.T) {
	type args struct {
		ms []map[string][]string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		// TODO: Add test cases.
		{
			args: args{
				ms: []map[string][]string{
					{
						"a": {"a", "A"},
						"b": {"b", "B"},
						"c": {"c", "C"},
					},
					{
						"a": {"Aa", "A"},
						"b": {"Bb", "B"},
						"c": {"Cc", "C"},
					},
				},
			},
			want: map[string][]string{
				"a": {"a", "A", "Aa"},
				"b": {"b", "B", "Bb"},
				"c": {"c", "C", "Cc"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapMergeOnValSlc(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapMerge4ValSlc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToValAny(t *testing.T) {
	type args struct {
		m map[int]int
	}
	tests := []struct {
		name string
		args args
		want map[int]any
	}{
		// TODO: Add test cases.
		{
			args: args{
				m: map[int]int{
					1: 1,
					2: 2,
					3: 3,
				},
			},
			want: map[int]any{
				1: 1,
				2: 2,
				3: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToValAny(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToValAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToArrValAny(t *testing.T) {
	type args struct {
		m map[int][]int
	}
	tests := []struct {
		name string
		args args
		want map[int][]any
	}{
		// TODO: Add test cases.
		{
			args: args{
				m: map[int][]int{
					1: {1, 2, 3},
					2: {4, 5, 6},
				},
			},
			want: map[int][]any{
				1: {1, 2, 3},
				2: {4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToArrValAny(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToArrValAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllValuesAreEmpty(t *testing.T) {

	type P struct {
		name string
	}
	var p *P = nil

	type args struct {
		m map[string]any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				m: map[string]any{
					"a": "",
					"b": []int{},
					"p": p,
					"n": nil,
				},
			},
			want: true,
		},
		{
			args: args{
				m: map[string]any{
					"a": "",
					"b": []int{},
					"c": &[]int{},
					"p": p,
					"n": nil,
				},
			},
			want: false,
		},
		{
			args: args{
				m: map[string]any{
					"a": "",
					"b": []int{},
					"c": &P{},
					"p": p,
					"n": nil,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapAllValuesAreEmpty(tt.args.m); got != tt.want {
				t.Errorf("MapAllEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetNestedMap(t *testing.T) {
	m := make(map[string]any)
	SetNestedMap(m, "ABD", "A", "B", "D")
	SetNestedMap(m, "ACD", "A", "C", "D")
	fmt.Println(m)
}

func TestMapFlatten(t *testing.T) {

	m := map[string]any{
		"P1": "ABC",
		"C1": map[string]any{
			"p2": "abc",
			"c2": 100,
			"C2": map[string]any{
				"p3":   nil,
				"z3":   "ok",
				"Arr":  []int{0, 1, 2},
				"ArrF": []string{"A", "B", "C"},
				"C3": map[string]any{
					"t4": 20.0,
					"m4": "final",
					"n4": "final",
					"arr": []map[string]any{
						{
							"z1": 1000,
							"Z3": "ABCDEFG",
						},
						{
							"z2": 2000.0,
							"Z4": []any{
								"ab",
								12,
								nil,
								false,
								[]any{
									"real-final1",
									"real-final2",
									map[string]any{
										"deepest": "real-final3",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(m)

	fm := MapNestedToFlat(m)
	n := 0
	for k, v := range fm {
		fmt.Printf("%2d: %s %v\n", n, k, v)
		n++
	}
}

func TestMapFlatToNested(t *testing.T) {

	m := map[string]any{
		"C1.c2":       100,
		"C1.C2.p3":    nil,
		"C1.p2":       "abc",
		"P1":          "ABC",
		"C1.C2.z3":    "ok",
		"C1.C2.C3.t4": 20.0,
		"C1.C2.C3.m4": "final",
	}

	nm := MapFlatToNested(m)
	fmt.Println(nm)
}

func TestMapCompare(t *testing.T) {

	m1 := map[string]any{
		"P1": "ABC",
		"C1": map[string]any{
			"p2": "abc",
			"c2": 100,
			"C2": map[string]any{
				"p3":   nil,
				"z3":   "ok",
				"Arr":  []int{0, 1, 2},
				"ArrF": []string{"A", "B", "C"},
				"C3": map[string]any{
					"t4": 20.0,
					"m4": "final",
					"n4": "final",
					"arr": []map[string]any{
						{
							"z1": 1000,
						},
						{
							"z2": 2000.0,
						},
					},
				},
			},
		},
	}

	m2 := map[string]any{
		"C1.c2":       100,
		"C1.C2.p3":    nil,
		"C1.p2":       "abc",
		"P1":          "ABC",
		"C1.C2.z3":    "ok",
		"C1.C2.Arr.0": 0,
		"C1.C2.Arr.1": 1,
		"C1.C2.Arr.2": 2,
		"C1.C2.C3.t4": 20.0,
		"C1.C2.C3.m4": "final",
		"C1.C2.C3.n4": "final",
	}

	fm := MapNestedToFlat(m1)
	fmt.Println(fm)
	fmt.Println(reflect.DeepEqual(fm, m2))

	nm := MapFlatToNested(m2)
	fmt.Println(nm)
	fmt.Println(reflect.DeepEqual(m1, nm))
}
