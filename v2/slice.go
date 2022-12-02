package v2

import (
	"fmt"
	"strings"
)

func traverseNestedSlice(slc []any, path string, sb *strings.Builder) {
	for i, e := range slc {
		if IsArrOrSlc(e) {
			traverseNestedSlice(e.([]any), fmt.Sprintf("%v.%d", path, i), sb)
		} else {

			// **
			// with value
			// **
			// path := fmt.Sprintf("%v @ %v.%d", e, path, i)
			// path = strings.Replace(path, "@ .", "@ ", 1)

			// **
			// without value
			// **
			path := fmt.Sprintf("%v.%d", path, i)
			path = strings.TrimPrefix(path, ".")

			sb.WriteString(path + "\n")

			// debug
			// if e != nil {
			// 	fmt.Printf("%v @%p %s\n", slc[i], &slc[i], path)
			// } else {
			// 	fmt.Printf("%v %s\n", slc[i], path)
			// }
		}
	}
}

func TraverseNestedSlice(slc []any) []string {
	var (
		path = ""
		sb   = &strings.Builder{}
	)
	traverseNestedSlice(slc, path, sb)
	return strings.Split(strings.TrimSpace(sb.String()), "\n")
}

// paths are like '1.2.3', '4.6.2'
func CapacityForSlice(paths ...string) []int {

	mLvlIndices := make(map[int][]int)
	for _, path := range paths {
		// fmt.Println(i, path)
		for lvl, seg := range strings.Split(path, ".") {
			if n, ok := AnyTryToType[int](seg); ok {
				mLvlIndices[lvl] = append(mLvlIndices[lvl], n)
			}
		}
	}

	mLvlCap := make(map[int]int)
	for lvl, indices := range mLvlIndices {
		mLvlCap[lvl] = Max(indices...) + 1
	}

	_, values := MapToKVs(mLvlCap, func(ki int, kj int) bool { return ki < kj }, nil)
	return values
}

func initNestedSlice(slc *any, caps ...int) {
	for lvl, c := range caps {
		*slc = make([]any, c)
		if lvl < len(caps)-1 {
			for i := 0; i < c; i++ {
				initNestedSlice(&((*slc).([]any))[i], caps[1:]...)
			}
		} else if lvl == len(caps)-1 {
			for i := 0; i < c; i++ {
				((*slc).([]any))[i] = struct{}{}
			}
		}
		break // keep 'break' here
	}
}

func InitNestedSlice(caps ...int) []any {
	var slc any = []any{}
	initNestedSlice(&slc, caps...)
	return slc.([]any)
}

// slc must have enough capacity for all element
func SetNestedSlice(slc []any, value any, idxSegs ...any) (ok bool) {
	if len(idxSegs) == 0 {
		return false
	}
	for _, seg := range idxSegs {
		if !IsUint(seg) {
			return false
		}
	}

	for i, seg := range idxSegs {
		if idx, _ := AnyTryToType[int](seg); idx < len(slc) {

			if i == len(idxSegs)-1 {
				slc[idx] = value
				return true
			} else {
				var ok bool
				if slc, ok = slc[idx].([]any); !ok {
					return false
				}
			}

		} else {
			return false
		}
	}
	return true
}

// This can only set single value in nested slice, useless for further assignment
//
// func MakeNestedSlice(value any, idxSegs ...any) (slc []any, err error) {
// 	if len(idxSegs) == 0 {
// 		return []any{value}, nil
// 	}
// 	for _, seg := range idxSegs {
// 		if !IsUint(seg) {
// 			return nil, fmt.Errorf("index path (idxSegs) must can be converted to unsigned int")
// 		}
// 	}

// 	for i, seg := range Reverse(idxSegs) {
// 		idx, _ := AnyTryToType[int](seg)

// 		if i == 0 { // value seg

// 			nNil := idx - len(slc)
// 			for i := 0; i < nNil; i++ {
// 				slc = append(slc, nil)
// 			}
// 			slc = append(slc, value)

// 		} else { // path seg

// 			slc = []any{slc}

// 			nNil := idx - len(slc)
// 			for i := 0; i <= nNil; i++ {
// 				slc = append(slc, []any{})
// 			}
// 			slc = Reverse(slc)
// 		}
// 	}
// 	return
// }
