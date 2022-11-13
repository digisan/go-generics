package ptb

import (
	"image"
	"sort"
)

// FilterMap : Filter & Modify []image.Point slice, return []bool slice
func FilterMap(arr []image.Point, filter func(i int, e image.Point) bool, modifier func(i int, e image.Point) bool) (r []bool) {
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
func Map2KVs(m map[image.Point]bool, less4key func(i image.Point, j image.Point) bool, less4value func(i bool, j bool) bool) (keys []image.Point, values []bool) {

	if m == nil {
		return nil, nil
	}
	if len(m) == 0 {
		return []image.Point{}, []bool{}
	}

	type kv struct {
		key   image.Point
		value bool
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
func MapMerge(ms ...map[image.Point]bool) map[image.Point][]bool {
	res := map[image.Point][]bool{}
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
func MapFilter(m map[image.Point]bool, filter func(k image.Point, v bool) bool) map[image.Point]bool {
	rt := make(map[image.Point]bool)
	for k, v := range m {
		if filter(k, v) {
			rt[k] = v
		}
	}
	return rt
}

// MapCopy:
func MapCopy(m map[image.Point]bool) map[image.Point]bool {
	return MapFilter(m, func(k image.Point, v bool) bool { return true })
}
