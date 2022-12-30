package v2

import (
	"fmt"
	"reflect"
	"strings"
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
			if got := MapReplMerge(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
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
			gotKeys, gotValues := MapToKVs(tt.args.m, tt.args.less4key, tt.args.less4value)
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
			if got := MapAllValAreEmpty(tt.args.m); got != tt.want {
				t.Errorf("MapAllEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapFlattenAndNested(t *testing.T) {

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
										"X": []any{
											"XXX",
											[]any{0.001, "", nil, []any{[]any{}}},
											nil,
											struct{}{},
										},
										"Y": "YYY",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Printf("m1: %+v\n", m1)
	str1 := fmt.Sprint(m1)

	fm := MapNestedToFlat(m1)
	n := 0
	for k, v := range fm {
		fmt.Printf("%2d: %s %v\n", n, k, v)
		n++
	}

	m2 := map[string]any{
		"C1.c2":                         100,
		"C1.C2.C3.arr.1.Z4.3":           false,
		"C1.C2.p3":                      nil,
		"C1.C2.Arr.0":                   0,
		"C1.C2.Arr.2":                   2,
		"C1.p2":                         "abc",
		"C1.C2.C3.m4":                   "final",
		"C1.C2.C3.arr.0.Z3":             "ABCDEFG",
		"C1.C2.C3.arr.1.Z4.1":           12,
		"C1.C2.z3":                      "ok",
		"C1.C2.ArrF.0":                  "A",
		"C1.C2.ArrF.2":                  "C",
		"C1.C2.C3.t4":                   20.0,
		"C1.C2.C3.n4":                   "final",
		"C1.C2.C3.arr.1.z2":             2000.0,
		"C1.C2.C3.arr.1.Z4.2":           nil,
		"C1.C2.C3.arr.1.Z4.4.2.deepest": "real-final3",
		"C1.C2.C3.arr.1.Z4.4.2.Y":       "YYY",
		"C1.C2.C3.arr.1.Z4.4.2.X.0":     "XXX",
		"C1.C2.C3.arr.1.Z4.4.2.X.1.0":   0.001,
		"C1.C2.C3.arr.1.Z4.4.2.X.1.1":   "",
		"C1.C2.C3.arr.1.Z4.4.2.X.1.2":   nil,
		"C1.C2.C3.arr.1.Z4.4.2.X.1.3.0": []any{},
		"C1.C2.C3.arr.1.Z4.4.2.X.2":     nil,
		"C1.C2.C3.arr.1.Z4.4.2.X.3":     struct{}{},
		"C1.C2.Arr.1":                   1,
		"C1.C2.ArrF.1":                  "B",
		"P1":                            "ABC",
		"C1.C2.C3.arr.0.z1":             1000,
		"C1.C2.C3.arr.1.Z4.0":           "ab",
		"C1.C2.C3.arr.1.Z4.4.0":         "real-final1",
		"C1.C2.C3.arr.1.Z4.4.1":         "real-final2",
	}

	fmt.Println("fm == m1 ?", reflect.DeepEqual(fm, m2))

	///////////////////////////////////////////////////

	fmt.Println("MapFlatToNested")

	nm := MapFlatToNested(m2, nil)
	fmt.Printf("nm: %+v\n", nm)
	str2 := fmt.Sprint(nm)

	///////////////////////////////////////////////////

	// fmt.Println("m1 == nm ?", reflect.DeepEqual(m1, nm))
	fmt.Println("m1Str == nmStr ?", str1 == str2)
}

func TestMapTryToSlc(t *testing.T) {

	m := map[string]any{
		"6": map[string]any{"key": 200},
		"5": map[string]any{"key": 100},
		"2": "c",
		"3": "d",
		"1": "b",
		"4": "e",
		"0": "A",
	}

	fmt.Println(m)

	s, err := MapTryToSlc(m)
	fmt.Println(err)
	fmt.Println(s)
}

func TestSetNestedMapIgnoreIdx(t *testing.T) {
	m := make(map[string]any)
	// SetNestedMapIgnoreIdx(m, "ABD", "A")
	SetNestedMapIgnoreIdx(m, "ABD", "A", "0")
	SetNestedMapIgnoreIdx(m, "ABD", "A", "1")
	SetNestedMapIgnoreIdx(m, "ACD", "A", "2", "D")
	SetNestedMapIgnoreIdx(m, "ACE", "B", "0", "E")
	fmt.Println(m)
}

// func TestSetNestedMap(t *testing.T) {
// 	var m any
// 	SetNestedMapWithIdx(m, "100", "0")

// 	fmt.Println(m)
// }

func TestTemp(t *testing.T) {

	// arr := InitNestedSlice(5, 1, 2)
	// SetNestedSlice(arr, 20, 4)
	// fmt.Println(arr)

	m := make(map[string]any)

	err := SetNestedMap(m, "XXX", "A", "B", "5", "D")
	fmt.Println(m)
	fmt.Println(err)

	err = SetNestedMap(m, "XX", "A", "B", "1", "C")
	fmt.Println(m)
	fmt.Println(err)

	err = SetNestedMap(m, "ZZZ", "B", "C", "3")
	fmt.Println(m)
	fmt.Println(err)

	err = SetNestedMap(m, "TTT", "B", "D", "3")
	fmt.Println(m)
	fmt.Println(err)

	err = SetNestedMap(m, "13", "BB", "1", "3")
	fmt.Println(m)
	fmt.Println(err)

	err = SetNestedMap(m, 02, "BB", "1", "2", "1", "E")
	fmt.Println(m)
	fmt.Println(err)

	err = SetNestedMap(m, "XXX", "D", "1", "2", "M", "N")
	fmt.Println(m)
	fmt.Println(err)
}

func TestMapFlatToNested(t *testing.T) {

	m2 := map[string]any{
		"X.1.1": "XXX",
		"X.0.0": 0,
	}

	fmt.Println(MapFlatToNested(m2, nil))

	fmt.Println(MapFlatToNested(m2, func(path string, value any) (p string, v any) {
		if strings.HasSuffix(path, ".1") {
			return "", nil
		}
		if strings.HasSuffix(path, ".0") {
			return path, value.(int) + 1
		}
		return path, value
	}))
}

func TestObjsonToFlatMap(t *testing.T) {

	type myStruct struct {
		A int
		B float64
		C []string
	}

	myData := myStruct{
		A: 100,
		B: 0.123,
		C: []string{"11", "1.2", "43.76"},
	}

	fm, err := ObjsonToFlatMap(myData)
	if err != nil {
		panic("Fatal: ObjsonToFlatMap")
	}
	fmt.Printf("%+v\n", fm)

	fmt.Println(FlatMapValTryToType[float64](fm, "A"))
	fmt.Println(FlatMapValsTryToTypes[float32](fm, "C"))
}

func TestMapValTryToType(t *testing.T) {

	m := map[string]any{
		"C1.0.z1":     "1000",
		"C1.1.Z4.0":   "ab",
		"C1.1.Z4.4.0": "real-final1",
		"C1.1.Z4.4.1": "real-final2",
		"C1.2a.0":     "777",
		"C1.2a.1":     "666",
	}

	v, ok := FlatMapValTryToType[int](m, "C1.0.z1")
	fmt.Println(ok)
	fmt.Println(v)

	v, ok = FlatMapValTryToType[int](m, "C1.2a.0")
	fmt.Println(ok)
	fmt.Println(v)

	vs, ok := FlatMapValsTryToTypes[int](m, "C1.2a")
	fmt.Println(ok)
	fmt.Println(vs)

}
