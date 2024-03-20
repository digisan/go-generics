package gogenerics

import (
	"fmt"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Sum[T Number](arr ...T) T {
	sum := new(T)
	for _, a := range arr {
		*sum += a
	}
	return *sum
}

type RangeType int

const (
	CloseOpen RangeType = iota
	CloseClose
	OpenOpen
	OpenClose
)

func InRange[T Number](n, left, right T, rt RangeType) (bool, error) {
	if left >= right {
		return false, fmt.Errorf("[left]-%v must less than [right]-%v", left, right)
	}
	switch rt {
	case CloseOpen:
		return n >= left && n < right, nil
	case CloseClose:
		return n >= left && n <= right, nil
	case OpenOpen:
		return n > left && n < right, nil
	case OpenClose:
		return n > left && n <= right, nil
	default:
		return n >= left && n < right, nil
	}
}

func NotInRange[T Number](n, left, right T, rt RangeType) (bool, error) {
	ok, err := InRange(n, left, right, rt)
	return !ok, err
}

func InRanges[T Number](n T, rt RangeType, ranges ...[2]T) (bool, error) {
	for _, rng := range ranges {
		ok, err := InRange(n, rng[0], rng[1], rt)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}

func NotInRanges[T Number](n T, rt RangeType, ranges ...[2]T) (bool, error) {
	ok, err := InRanges(n, rt, ranges...)
	return !ok, err
}

func HasOverlap[T Number](a1 [2]T, a2 [2]T, rt RangeType) bool {
	for _, e := range a1 {
		if e > a2[0] && e < a2[1] {
			return true
		}
	}
	for _, e := range a2 {
		if e > a1[0] && e < a1[1] {
			return true
		}
	}

	s1, e1 := a1[0], a1[1]
	s2, e2 := a2[0], a2[1]
	switch rt {
	case CloseClose:
		if s1 == s2 || e1 == s2 || s1 == e2 || e1 == e2 {
			return true
		}
	case CloseOpen:
		if s1 == s2 {
			return true
		}
	case OpenClose:
		if e1 == e2 {
			return true
		}
	case OpenOpen:
		if s1 == s2 && e1 == e2 {
			return true
		}
	default:
		return false
	}
	return false
}

func HasOverlapSlc[T Number](s1 []T, s2 []T, rt RangeType) bool {
	a1 := (*[2]T)(s1)
	a2 := (*[2]T)(s2)
	return HasOverlap(*a1, *a2, rt)
}

func ShareOverlap[T Number](rt RangeType, sn ...[]T) bool {
	if len(sn) < 2 {
		return false
	}
	for i, s1 := range sn {
		for j, s2 := range sn {
			if i == j {
				continue
			}
			if HasOverlapSlc(s1, s2, rt) {
				return true
			}
		}
	}
	return false
}

func RangeUnion[T Number](s1, s2 []T, rt RangeType) ([]T, bool) {
	min := Min(s1[0], s2[0])
	max := Max(s1[1], s2[1])
	if HasOverlapSlc(s1, s2, rt) {
		return []T{min, max}, true
	}
	return nil, false
}
