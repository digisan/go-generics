package gogenerics

import "reflect"

// IF : Ternary Operator LIKE < ? : >, BUT NO S/C
// src1 and src2 MUST all valid. e.g. type assert, nil pointer, out of index
func IF[T any](condition bool, src1, src2 T) T {
	if condition {
		return src1
	}
	return src2
}

// MATCH : NO ShortCut, MUST all valid, e.g. type assert, nil pointer, out of index
// MATCH(check, case1, case2, ..., value1, value2, ..., default)
func MATCH[T any](chkCasesValues ...T) T {
	l := len(chkCasesValues)
	if l < 4 || l%2 == 1 {
		panic("PARAM_INVALID, Must be 'MATCH(check, case1, case2, ..., value1, value2, ..., default)'")
	}
	_, l1, l2 := 1, (l-1)/2, (l-1)/2
	check := chkCasesValues[0]
	cases := chkCasesValues[1 : 1+l1]
	values := chkCasesValues[1+l1 : 1+l1+l2]
	for i, c := range cases {
		if reflect.DeepEqual(check, c) {
			return values[i]
		}
	}
	return chkCasesValues[l-1]
}
