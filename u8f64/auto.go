package u8f64

import "sort"

// FilterMap : Filter & Modify []byte slice, return []float64 slice
func FilterMap(arr []byte, filter func(i int, e byte) bool, modifier func(i int, e byte) float64) (r []float64) {
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
func Map2KVs(m map[byte]float64, less4key func(i byte, j byte) bool, less4value func(i float64, j float64) bool) (keys []byte, values []float64) {

	if m == nil {
		return nil, nil
	}
	if len(m) == 0 {
		return []byte{}, []float64{}
	}

	type kv struct {
		key   byte
		value float64
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
func MapMerge(ms ...map[byte]float64) map[byte][]float64 {
	res := map[byte][]float64{}
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
func MapFilter(m map[byte]float64, filter func(k byte, v float64) bool) map[byte]float64 {
	rm := make(map[byte]float64)
	for k, v := range m {
		if filter(k, v) {
			rm[k] = v
		}
	}
	return rm
}
