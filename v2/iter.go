package v2

// Iter : for i := range Iter(end) / (start, end) / (start, step, end)
func Iter[T Integer](params ...T) <-chan T {

	start, end, step := *new(T), *new(T), *new(T)

	switch len(params) {

	case 1:
		end = params[0]
		step = IF[T](end >= 0, 1, -1)

	case 2:
		start, end = params[0], params[1]
		step = IF[T](end >= start, 1, -1)

	case 3:
		start, step, end = params[0], params[1], params[2]
		if start > end {
			if step > 0 {
				panic("step error, must be NEGATIVE for bigger start")
			}
		}
		if start < end {
			if step < 0 {
				panic("step error, must be POSITIVE for smaller start")
			}
		}

	default:
		panic("params' count only can be 1, 2 or 3")
	}

	ch := make(chan T)
	if start > end {
		go func() {
			defer close(ch)
			for i := start; i > end; i += step {
				ch <- i
			}
		}()
	} else {
		go func() {
			defer close(ch)
			for i := start; i < end; i += step {
				ch <- i
			}
		}()
	}
	return ch
}

// IterToSlc : for i := range Iter(end) / (start, end) / (start, step, end)
func IterToSlc[T Integer](params ...T) (slc []T) {
	// if len(params) == 1 {
	// 	for i := range make([]struct{}, params[0]) {
	// 		slc = append(slc, i)
	// 	}
	// 	return
	// }
	for i := range Iter(params...) {
		slc = append(slc, i)
	}
	return
}
