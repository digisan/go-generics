package v2

func MapCvt[T1 comparable, T2 any](m map[any]any) map[T1]T2 {
	rt := make(map[T1]T2)
	for k, v := range m {
		rt[k.(T1)] = v.(T2)
	}
	return rt
}

func SlcCvt[T any](s []any) []T {
	rt := make([]T, 0, len(s))
	for _, a := range s {
		rt = append(rt, a.(T))
	}
	return rt
}
