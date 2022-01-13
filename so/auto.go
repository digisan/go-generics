package so

import "sort"

// FilterMap : Filter & Modify []string slice, return []interface{} slice
func FilterMap(arr []string, filter func(i int, e string) bool, modifier func(i int, e string) interface{}) (r []interface{}) {
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
func Map2KVs(m map[string]interface{}, less4key func(i string, j string) bool, less4value func(i interface{}, j interface{}) bool) (keys []string, values []interface{}) {

	if m == nil {
		return nil, nil
	}
	if len(m) == 0 {
		return []string{}, []interface{}{}
	}

	type kv struct {
		key   string
		value interface{}
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
func MapMerge(ms ...map[string]interface{}) map[string][]interface{} {
	res := map[string][]interface{}{}
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
func MapFilter(m map[string]interface{}, filter func(k string, v interface{}) bool) map[string]interface{} {
	rm := make(map[string]interface{})
	for k, v := range m {
		if filter(k, v) {
			rm[k] = v
		}
	}
	return rm
}
