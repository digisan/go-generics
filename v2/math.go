package v2

import "log"

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

func InRange[T Number](n, left, right T, leftClose, rightClose bool) bool {
	if left > right {
		log.Fatalf("[left]-%v must NOT greater than [right]-%v", left, right)
	}
	switch {
	case leftClose && rightClose:
		return n >= left && n <= right
	case !leftClose && rightClose:
		return n > left && n <= right
	case leftClose && !rightClose:
		return n >= left && n < right
	default:
		return n > left && n < right
	}
}

func NotInRange[T Number](n, left, right T, leftClose, rightClose bool) bool {
	return !InRange(n, left, right, leftClose, rightClose)
}

func InCloseRange[T Number](n, left, right T) bool {
	return InRange(n, left, right, true, true)
}

func NotInCloseRange[T Number](n, left, right T) bool {
	return !InCloseRange(n, left, right)
}

func InOpenRange[T Number](n, left, right T) bool {
	return InRange(n, left, right, false, false)
}

func NotInOpenRange[T Number](n, left, right T) bool {
	return !InOpenRange(n, left, right)
}

func InCloseOpenRange[T Number](n, left, right T) bool {
	return InRange(n, left, right, true, false)
}

func NotInCloseOpenRange[T Number](n, left, right T) bool {
	return !InCloseOpenRange(n, left, right)
}

// a1, a2 are both close-open range
func OverlappedA[T Number](a1, a2 [2]T) bool {
	switch {
	case InCloseOpenRange(a1[0], a2[0], a2[1]):
		return true
	case a2[0] < a1[1] && a1[1] < a2[1]:
		return true
	case InCloseOpenRange(a2[0], a1[0], a1[1]):
		return true
	case a1[0] < a2[1] && a2[1] < a1[1]:
		return true
	default:
		return false
	}
}

// s1, s2 are both close-open range
func OverlappedS[T Number](s1, s2 []T) bool {
	a1 := (*[2]T)(s1)
	a2 := (*[2]T)(s2)
	return OverlappedA(*a1, *a2)
}

func SpanJoin[T Number](s1, s2 []T, ocJoin bool) ([]T, bool) {
	max := Max(s1[0], s1[1], s2[0], s2[1])
	min := Min(s1[0], s1[1], s2[0], s2[1])
	if OverlappedS(s1, s2) {
		return []T{min, max}, true
	}
	if ocJoin {
		if s1[1] == s2[0] || s2[1] == s1[0] {
			return []T{min, max}, true
		}
	}
	return nil, false
}

// sn are all close-open range
func HasOverlapped[T Number](sn ...[]T) bool {
	if len(sn) < 2 {
		return false
	}
	for i, s1 := range sn {
		for j, s2 := range sn {
			if i == j {
				continue
			}
			if OverlappedS(s1, s2) {
				return true
			}
		}
	}
	return false
}
