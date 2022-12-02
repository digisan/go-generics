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

type Pair[T any] struct {
	a      T
	b      T
	first  bool
	last   bool
	validA bool
	validB bool
}

// 1,2,3,4... => (1,2), (2,3), (3,4)...
func IterPair[T any](params ...T) <-chan Pair[T] {
	ch := make(chan Pair[T])
	go func() {
		defer close(ch)
		if len(params) > 1 {
			for i, p := range params[1:] {
				ch <- Pair[T]{
					a:      params[i],
					b:      p,
					first:  IF(i == 0, true, false),
					last:   IF(i == len(params)-2, true, false),
					validA: true,
					validB: true,
				}
			}
			ch <- Pair[T]{
				a:      params[len(params)-1],
				b:      *new(T),
				first:  false,
				last:   true,
				validA: true,
				validB: false,
			}
		} else if len(params) == 1 {
			ch <- Pair[T]{
				a:      params[0],
				b:      *new(T),
				first:  true,
				last:   true,
				validA: true,
				validB: false,
			}
		} else {
			ch <- Pair[T]{
				a:      *new(T),
				b:      *new(T),
				first:  false,
				last:   false,
				validA: false,
				validB: false,
			}
		}
	}()
	return ch
}

// type Triple[T any] struct {
// 	a     T
// 	b     T
// 	c     T
// 	first bool
// 	last  bool
// 	ok    bool
// }

// // 1,2,3,4,5,6,7... => (1,2,3), (2,3,4), (3,4,5)...
// func IterTriple[T any](params ...T) <-chan Triple[T] {
// 	ch := make(chan Triple[T])
// 	go func() {
// 		defer close(ch)
// 		if len(params) > 2 {
// 			for i, p := range params[2:] {
// 				ch <- Triple[T]{
// 					a:     params[i],
// 					b:     params[i+1],
// 					c:     p,
// 					first: IF(i == 0, true, false),
// 					last:  IF(i == len(params)-3, true, false),
// 					ok:    true,
// 				}
// 			}
// 		} else if len(params) == 2 {
// 			ch <- Triple[T]{
// 				a:     *new(T),
// 				b:     params[0],
// 				c:     params[1],
// 				first: true,
// 				last:  true,
// 				ok:    false,
// 			}
// 		} else if len(params) == 1 {
// 			ch <- Triple[T]{
// 				a:     *new(T),
// 				b:     *new(T),
// 				c:     params[0],
// 				first: true,
// 				last:  true,
// 				ok:    false,
// 			}
// 		} else {
// 			ch <- Triple[T]{
// 				a:     *new(T),
// 				b:     *new(T),
// 				c:     *new(T),
// 				first: false,
// 				last:  false,
// 				ok:    false,
// 			}
// 		}
// 	}()
// 	return ch
// }
