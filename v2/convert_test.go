package v2

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func TestConstBytesToStr(t *testing.T) {
	bytes := []byte{65, 66, 67}
	str := ConstBytesToStr(bytes)
	fmt.Println(str)
	bPtr := unsafe.SliceData(bytes)
	sPtr := unsafe.StringData(str)
	fmt.Printf("%p %p\n", bPtr, sPtr)
	if bPtr != sPtr {
		panic("error")
	}
}

func TestStrToConstBytes(t *testing.T) {
	str := "ABC"
	bytes := StrToConstBytes(str)
	fmt.Println(bytes)
	bPtr := unsafe.SliceData(bytes)
	sPtr := unsafe.StringData(str)
	fmt.Printf("%p %p\n", bPtr, sPtr)
	if bPtr != sPtr {
		panic("error")
	}
}

func TestMapCvt(t *testing.T) {
	m := map[any]any{
		1: "11",
		2: "22",
	}
	m1 := MapCvtKVAnyToType[int, string](m)
	// m1 := m.(map[int]string)
	fmt.Println(m1[1] + m1[2])
}

func TestSyncMap2Map(t *testing.T) {
	sm := sync.Map{}
	sm.Store(1, "11")
	sm.Store(2, "22")
	m := SyncMapToMap[int, string](sm)
	fmt.Println(m[1] + m[2])
}

func TestMap2SyncMap(t *testing.T) {
	m := map[int]string{
		1: "11",
		2: "22",
	}
	sm := MapToSyncMap(m)
	sm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

func TestSlcCvt(t *testing.T) {
	s := []any{1, 2, 3, 4, 5}
	s1 := AnysToTypes[int](s)
	// s1 := s.([]int)
	fmt.Println(Sum(s1...))

	S := []any{"1", 2, "3", 4, "5"}
	S1, ok := AnysTryToTypes[int](S)
	fmt.Println(Sum(S1...), ok)
}

func TestSlc2Types(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	s1, ok1 := TypesAsAnyTryToTypes[uint16](s)
	if ok1 {
		fmt.Println(Sum(s1...))
	}

	s3, ok3 := TypesAsAnyTryToTypes[float32](s1)
	if ok3 {
		fmt.Println(Sum(s3...))
	}

	// s2 := s.([]int)
	// fmt.Println(Sum(s2...))
}

func TestSlc2Anys(t *testing.T) {

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	gs1 := TypesAsAnyToAnys(s1)
	fmt.Println(gs1)

	s2 := []string{"a1", "b2", "c3", "d4", "e5", "f6", "g7", "h8"}
	gs2 := TypesAsAnyToAnys(s2)
	fmt.Println(gs2)
}

func TestPtrSlcToSlc(t *testing.T) {

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := []string{"a1", "b2", "c3", "d4", "e5", "f6", "g7", "h8"}

	addressesInt := SlcToPtrSlc(s1...)
	fmt.Println(addressesInt)

	valsInt := PtrSlcToSlc(addressesInt...)
	fmt.Println(valsInt)

	///

	addressesStr := SlcToPtrSlc(s2...)
	fmt.Println(addressesStr)

	valsStr := PtrSlcToSlc(addressesStr...)
	fmt.Println(valsStr)
}

func TestAnyTryToType(t *testing.T) {

	primitive, ok := AnyTryToType[float32]([]byte("123.0"))
	if ok {
		fmt.Println(primitive + 222)
	} else {
		panic("")
	}

	b, ok := AnyTryToType[[]byte](123)
	if ok {
		fmt.Println(b, "string as", string(b))
	} else {
		panic("")
	}

	f, ok := AnyTryToType[float32]("1")
	if ok {
		fmt.Println(f + 100.0)
	} else {
		panic("")
	}

	u, ok := AnyTryToType[uint]("123")
	if ok {
		fmt.Println(u + 100.0)
	} else {
		panic("")
	}

	u64, ok := AnyTryToType[uint64]("123")
	if ok {
		fmt.Println(u64 + 100.0)
	} else {
		panic("")
	}

	c64, ok := AnyTryToType[complex64]("(2+3i)")
	if ok {
		fmt.Println(c64 + 100.0)
	} else {
		panic("")
	}

	c128, ok := AnyTryToType[complex128](2 + 3i)
	if ok {
		fmt.Println(c128 + 200.0)
	} else {
		panic("")
	}

	str, ok := AnyTryToType[string](2 + 3i)
	if ok {
		fmt.Println("hello " + str)
	} else {
		panic("")
	}

	boolean, ok := AnyTryToType[bool]("TRUE")
	if ok {
		fmt.Print("bool ")
		fmt.Println(boolean)
	} else {
		panic("")
	}

	tm, ok := AnyTryToType[time.Time]("2022-12-16 09:11:51")
	if ok {
		fmt.Print("TM OK: ")
		fmt.Println(tm)
	} else {
		panic("")
	}
}

func TestTryToTime(t *testing.T) {
	fmt.Println(time.Now())
	if tm, ok := TryToDateTime("2022-12-16 09:11:51"); ok {
		fmt.Println(tm)
	}
	if tm, ok := TryToDateTime("2022-12-16 09:11:51.0"); ok {
		fmt.Println(tm)
	}
	if tm, ok := TryToDateTime("2022-12-16 09:11:51.123"); ok {
		fmt.Println(tm)
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
			if got := MapCvtVTypeToAny(tt.args.m); !reflect.DeepEqual(got, tt.want) {
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
			if got := MapCvtVTypesToAnys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToArrValAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
