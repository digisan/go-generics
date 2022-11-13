package v2

import (
	"log"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Nums2Floats[T Number](numbers ...T) []float64 {
	if numbers == nil {
		return nil
	}
	rt := make([]float64, 0, len(numbers))
	for _, a := range numbers {
		rt = append(rt, float64(a))
	}
	return rt
}

func Sum[T Number](arr ...T) T {
	sum := new(T)
	for _, a := range arr {
		*sum += a
	}
	return *sum
}

func InSpan[T Number](n T, nOpen bool, a [2]T) bool {
	if a[0] > a[1] {
		log.Fatalf("a[0]-%v must NOT greater than a[1]-%v", a[0], a[1])
	}
	if nOpen {
		return n > a[0] && n < a[1]
	}
	return n >= a[0] && n < a[1]
}

func Overlapped[T Number](a1, a2 [2]T) bool {
	for i, n1 := range a1 {
		if InSpan(n1, i%2 == 1, a2) {
			return true
		}
	}
	for i, n2 := range a2 {
		if InSpan(n2, i%2 == 1, a1) {
			return true
		}
	}
	return false
}

func IsOverlapped[T Number](s1, s2 []T) bool {
	a1 := (*[2]T)(s1)
	a2 := (*[2]T)(s2)
	return Overlapped(*a1, *a2)
}

func SpanJoin[T Number](s1, s2 []T, ocJoin bool) ([]T, bool) {
	max := Max(s1[0], s1[1], s2[0], s2[1])
	min := Min(s1[0], s1[1], s2[0], s2[1])
	if IsOverlapped(s1, s2) {
		return []T{min, max}, true
	}
	if ocJoin {
		if s1[1] == s2[0] || s2[1] == s1[0] {
			return []T{min, max}, true
		}
	}
	return nil, false
}

func HasOverlapped[T Number](sn ...[]T) bool {
	if len(sn) < 2 {
		return false
	}
	for i, s1 := range sn {
		for j, s2 := range sn {
			if i == j {
				continue
			}
			if IsOverlapped(s1, s2) {
				return true
			}
		}
	}
	return false
}
