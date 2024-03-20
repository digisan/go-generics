package gogenerics

// IsSuper :
func IsSuper[T comparable](setA, setB []T) bool {
NEXT_B:
	for _, b := range setB {
		for _, a := range setA {
			if a == b {
				continue NEXT_B
			}
		}
		return false
	}
	return len(setA) > len(setB)
}

// IsSub :
func IsSub[T comparable](setA, setB []T) bool {
	return IsSuper(setB, setA)
}

// equals :
func equals[T comparable](setA, setB []T) bool {
	if (setA == nil && setB != nil) || (setA != nil && setB == nil) {
		return false
	}
	if len(setA) != len(setB) {
		return false
	}

	tmpA := make([]T, len(setA))
	tmpB := make([]T, len(setB))
	copy(tmpA, setA)
	copy(tmpB, setB)

AGAIN:
	for i, a := range tmpA {
		for j, b := range tmpB {
			if a == b {
				DelEleAt(&tmpA, i)
				DelEleAt(&tmpB, j)
				goto AGAIN
			}
		}
	}
	return len(tmpA) == 0 && len(tmpB) == 0
}

// Equals
func Equals[T comparable](sets ...[]T) bool {
	for i := 0; i < len(sets)-1; i++ {
		this := sets[i]
		next := sets[i+1]
		if !equals(this, next) {
			return false
		}
	}
	return true
}

// SupEq :
func SupEq[T comparable](setA, setB []T) bool {
	return IsSuper(setA, setB) || Equals(setA, setB)
}

// SubEq :
func SubEq[T comparable](setA, setB []T) bool {
	return IsSub(setA, setB) || Equals(setA, setB)
}

// union :
func union[T comparable](setA, setB []T) (set []T) {
	if setA == nil && setB == nil {
		return nil
	}
	if setA == nil && setB != nil {
		return setB
	}
	if setA != nil && setB == nil {
		return setA
	}

	m := make(map[T]struct{})
	for _, a := range setA {
		if _, ok := m[a]; !ok {
			set = append(set, a)
			m[a] = struct{}{}
		}
	}
	for _, b := range setB {
		if _, ok := m[b]; !ok {
			set = append(set, b)
			m[b] = struct{}{}
		}
	}
	if set == nil {
		return []T{}
	}
	return
}

// Union :
func Union[T comparable](sets ...[]T) (set []T) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = union(set, s)
	}
	return set
}

// intersect :
func intersect[T comparable](setA, setB []T) (set []T) {
	if setA == nil || setB == nil {
		return nil
	}

	copyA, copyB := make([]T, len(setA)), make([]T, len(setB))
	copy(copyA, setA)
	copy(copyB, setB)

AGAIN:
	for i, a := range copyA {
		for j, b := range copyB {
			if a == b {
				set = append(set, a)
				DelEleAt(&copyA, i)
				DelEleAt(&copyB, j)
				goto AGAIN
			}
		}
	}
	if set == nil {
		return []T{}
	}
	return
}

// Intersect :
func Intersect[T comparable](sets ...[]T) (set []T) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = intersect(set, s)
	}
	return set
}

func minus[T comparable](setA, setB []T) (set []T) {
	if setA == nil {
		return nil
	}
	set = make([]T, 0)

NEXT_A:
	for _, a := range setA {
		for _, b := range setB {
			if a == b {
				continue NEXT_A
			}
		}
		set = append(set, a)
	}
	return
}

func Minus[T comparable](setA []T, setOthers ...[]T) (set []T) {
	return minus(setA, Union(setOthers...))
}
