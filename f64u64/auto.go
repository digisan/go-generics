package f64u64

import "sort"

// FilterMap : Filter & Modify []float64 slice, return []uint slice
func FilterMap(arr []float64, filter func(i int, e float64) bool, modifier func(i int, e float64) uint) (r []uint) {
	if modifier == nil {
		panic("modifier cannot be nil")
	}

	switch {
	case filter != nil:
		for i, e := range arr {
			if filter(i, e) {
				r = append(r, modifier(i, e))
			}
		}
	default:
		for i, e := range arr {
			r = append(r, modifier(i, e))
		}
	}
	return
}

var (
	FM = FilterMap
)

// Map2KVs : map to key slice & value slice
func Map2KVs(m map[float64]uint, less4key func(i float64, j float64) bool, less4value func(i uint, j uint) bool) (keys []float64, values []uint) {

	if m == nil {
		return nil, nil
	}
	if len(m) == 0 {
		return []float64{}, []uint{}
	}

	type kv struct {
		key   float64
		value uint
	}

	kvSlc := []kv{}
	for k, v := range m {
		kvSlc = append(kvSlc, kv{key: k, value: v})
	}

	switch {
	case less4key != nil && less4value == nil:
		sort.SliceStable(kvSlc, func(i, j int) bool { return less4key(kvSlc[i].key, kvSlc[j].key) })

	case less4key == nil && less4value != nil:
		sort.SliceStable(kvSlc, func(i, j int) bool { return less4value(kvSlc[i].value, kvSlc[j].value) })

	case less4key != nil && less4value != nil:
		sort.SliceStable(kvSlc, func(i, j int) bool {
			if kvSlc[i].value == kvSlc[j].value {
				return less4key(kvSlc[i].key, kvSlc[j].key)
			}
			return less4value(kvSlc[i].value, kvSlc[j].value)
		})

	default:
		// do not sort
	}

	for _, kvEle := range kvSlc {
		keys = append(keys, kvEle.key)
		values = append(values, kvEle.value)
	}
	return
}

// MapMerge:
func MapMerge(ms ...map[float64]uint) map[float64][]uint {
	res := map[float64][]uint{}
	for _, m := range ms {
	srcMap:
		for k, v := range m {
			// Check if (k,v) was added before:
			for _, v2 := range res[k] {
				if v == v2 {
					continue srcMap
				}
			}
			res[k] = append(res[k], v)
		}
	}
	return res
}

// MapFilter:
func MapFilter(m map[float64]uint, filter func(k float64, v uint) bool) map[float64]uint {
	rm := make(map[float64]uint)
	for k, v := range m {
		if filter(k, v) {
			rm[k] = v
		}
	}
	return rm
}
