package gogenerics

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
	A      T
	B      T
	First  bool
	Last   bool
	ValidA bool
	ValidB bool
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
					A:      data[i],
					B:      p,
					First:  IF(i == 0, true, false),
					Last:   false,
					ValidA: true,
					ValidB: true,
				}
			}
			ch <- Pair[T]{
				A:      data[len(data)-1],
				B:      *new(T),
				First:  false,
				Last:   true,
				ValidA: true,
				ValidB: false,
			}

		case len(data) == 1:
			ch <- Pair[T]{
				A:      data[0],
				B:      *new(T),
				First:  true,
				Last:   true,
				ValidA: true,
				ValidB: false,
			}

		default:
			ch <- Pair[T]{
				A:      *new(T),
				B:      *new(T),
				First:  false,
				Last:   false,
				ValidA: false,
				ValidB: false,
			}
		}
	}()
	return ch
}

type Triple[T any] struct {
	A      T
	B      T
	C      T
	First  bool
	Last   bool
	ValidA bool
	ValidB bool
	ValidC bool
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
					A:      data[i],
					ValidA: true,
					B:      data[i+1],
					ValidB: true,
					C:      p,
					ValidC: true,
					First:  IF(i == 0, true, false),
					Last:   false,
				}
			}
			ch <- Triple[T]{
				A:      data[len(data)-2],
				ValidA: true,
				B:      data[len(data)-1],
				ValidB: true,
				C:      *new(T),
				ValidC: false,
				First:  false,
				Last:   false,
			}
			ch <- Triple[T]{
				A:      data[len(data)-1],
				ValidA: true,
				B:      *new(T),
				ValidB: false,
				C:      *new(T),
				ValidC: false,
				First:  false,
				Last:   true,
			}

		case len(data) == 2:
			ch <- Triple[T]{
				A:      data[0],
				ValidA: true,
				B:      data[1],
				ValidB: true,
				C:      *new(T),
				ValidC: false,
				First:  true,
				Last:   false,
			}
			ch <- Triple[T]{
				A:      data[1],
				ValidA: true,
				B:      *new(T),
				ValidB: false,
				C:      *new(T),
				ValidC: false,
				First:  false,
				Last:   true,
			}

		case len(data) == 1:
			ch <- Triple[T]{
				A:      data[0],
				ValidA: true,
				B:      *new(T),
				ValidB: false,
				C:      *new(T),
				ValidC: false,
				First:  true,
				Last:   true,
			}

		default:
			ch <- Triple[T]{
				A:      *new(T),
				ValidA: false,
				B:      *new(T),
				ValidB: false,
				C:      *new(T),
				ValidC: false,
				First:  false,
				Last:   false,
			}
		}
	}()
	return ch
}

type Cache[T any] struct {
	Elem  T
	Cache []T
	First bool
	Last  bool
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
				Elem:  e,
				Cache: cache,
				First: IF(i == 0, true, false),
				Last:  IF(i == len(data)-1, true, false),
			}
		}
	}()
	return ch
}
