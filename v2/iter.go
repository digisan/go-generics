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

////////////////////////////////////////////////////////////////////////

type Pair[T any] struct {
	a      T
	b      T
	first  bool
	last   bool
	validA bool
	validB bool
}

// 1,2,3,4 => (1,2), (2,3), (3,4), (4, junk)
func IterPair[T any](data []T) <-chan Pair[T] {
	ch := make(chan Pair[T])
	go func() {
		defer close(ch)
		switch {
		case len(data) >= 2:
			for i, p := range data[1:] {
				ch <- Pair[T]{
					a:      data[i],
					b:      p,
					first:  IF(i == 0, true, false),
					last:   false,
					validA: true,
					validB: true,
				}
			}
			ch <- Pair[T]{
				a:      data[len(data)-1],
				b:      *new(T),
				first:  false,
				last:   true,
				validA: true,
				validB: false,
			}

		case len(data) == 1:
			ch <- Pair[T]{
				a:      data[0],
				b:      *new(T),
				first:  true,
				last:   true,
				validA: true,
				validB: false,
			}

		default:
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

type Triple[T any] struct {
	a      T
	b      T
	c      T
	first  bool
	last   bool
	validA bool
	validB bool
	validC bool
}

// 1,2,3,4 => (1,2,3), (2,3,4), (3,4,junk), (4,junk,junk)
func IterTriple[T any](data []T) <-chan Triple[T] {
	ch := make(chan Triple[T])
	go func() {
		defer close(ch)
		switch {
		case len(data) >= 3:
			for i, p := range data[2:] {
				ch <- Triple[T]{
					a:      data[i],
					validA: true,
					b:      data[i+1],
					validB: true,
					c:      p,
					validC: true,
					first:  IF(i == 0, true, false),
					last:   false,
				}
			}
			ch <- Triple[T]{
				a:      data[len(data)-2],
				validA: true,
				b:      data[len(data)-1],
				validB: true,
				c:      *new(T),
				validC: false,
				first:  false,
				last:   false,
			}
			ch <- Triple[T]{
				a:      data[len(data)-1],
				validA: true,
				b:      *new(T),
				validB: false,
				c:      *new(T),
				validC: false,
				first:  false,
				last:   true,
			}

		case len(data) == 2:
			ch <- Triple[T]{
				a:      data[0],
				validA: true,
				b:      data[1],
				validB: true,
				c:      *new(T),
				validC: false,
				first:  true,
				last:   false,
			}
			ch <- Triple[T]{
				a:      data[1],
				validA: true,
				b:      *new(T),
				validB: false,
				c:      *new(T),
				validC: false,
				first:  false,
				last:   true,
			}

		case len(data) == 1:
			ch <- Triple[T]{
				a:      data[0],
				validA: true,
				b:      *new(T),
				validB: false,
				c:      *new(T),
				validC: false,
				first:  true,
				last:   true,
			}

		default:
			ch <- Triple[T]{
				a:      *new(T),
				validA: false,
				b:      *new(T),
				validB: false,
				c:      *new(T),
				validC: false,
				first:  false,
				last:   false,
			}
		}
	}()
	return ch
}

type Cache[T any] struct {
	elem  T
	cache []T
	first bool
	last  bool
}

func IterCache[T any](data []T, nPrev, nNext int, junk T) <-chan Cache[T] {
	ch := make(chan Cache[T])
	go func() {
		defer close(ch)
		for i, e := range data {
			var (
				cache     []T
				idxS      int
				idxE      int
				nHeadJunk = 0
				nTailJunk = 0
				headJunk  []T
				tailJunk  []T
			)
			if i-nPrev < 0 {
				nHeadJunk = nPrev - i
				for i := 0; i < nHeadJunk; i++ {
					headJunk = append(headJunk, junk)
				}
			}
			if i+nNext >= len(data) {
				nTailJunk = i + nNext - len(data) + 1
				for i := 0; i < nTailJunk; i++ {
					tailJunk = append(tailJunk, junk)
				}
			}
			idxS = IF(i-nPrev >= 0, i-nPrev, 0)
			idxE = IF(i+nNext < len(data), i+nNext+1, len(data))
			cache = append(cache, headJunk...)
			cache = append(cache, data[idxS:idxE]...)
			cache = append(cache, tailJunk...)
			ch <- Cache[T]{
				elem:  e,
				cache: cache,
				first: IF(i == 0, true, false),
				last:  IF(i == len(data)-1, true, false),
			}
		}
	}()
	return ch
}
