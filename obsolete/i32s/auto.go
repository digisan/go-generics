package i32s

import "sort"

// FilterMap : Filter & Modify []rune slice, return []string slice
func FilterMap(arr []rune, filter func(i int, e rune) bool, modifier func(i int, e rune) string) (r []string) {
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
func Map2KVs(m map[rune]string, less4key func(i rune, j rune) bool, less4value func(i string, j string) bool) (keys []rune, values []string) {

	if m == nil {
		return nil, nil
	}
	if len(m) == 0 {
		return []rune{}, []string{}
	}

	type kv struct {
		key   rune
		value string
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
func MapMerge(ms ...map[rune]string) map[rune][]string {
	res := map[rune][]string{}
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
func MapFilter(m map[rune]string, filter func(k rune, v string) bool) map[rune]string {
	rt := make(map[rune]string)
	for k, v := range m {
		if filter(k, v) {
			rt[k] = v
		}
	}
	return rt
}

// MapCopy:
func MapCopy(m map[rune]string) map[rune]string {
	return MapFilter(m, func(k rune, v string) bool { return true })
}
