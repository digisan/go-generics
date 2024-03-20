package gogenerics

import (
	"fmt"
	"strings"
)

type Queue[T any] []T

// *** Enqueue :
func (q *Queue[T]) Enqueue(items ...T) int {
	*q = append(*q, items...)
	return len(items)
}

// *** Len :
func (q *Queue[T]) Len() int {
	return len(*q)
}

// *** Dequeue :
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Len() > 0 {
		first := (*q)[0]
		*q = (*q)[1:]
		return first, true
	}
	return *new(T), false
}

// *** Peek :
func (q *Queue[T]) Peek() (T, bool) {
	if q.Len() > 0 {
		return (*q)[0], true
	}
	return *new(T), false
}

// *** Clear :
func (q *Queue[T]) Clear() Queue[T] {
	cp := q.Copy()
	*q = Queue[T]{}
	return cp
}

// *** Copy :
func (q *Queue[T]) Copy() Queue[T] {
	tmp := make([]T, q.Len())
	copy(tmp, *q)
	return Queue[T](tmp)
}

// *** Sink :
func (q *Queue[T]) Sink() []T {
	n := q.Len()
	arr := make([]T, 0, n)
	for {
		if ele, ok := q.Dequeue(); ok {
			arr = append(arr, ele)
		} else {
			break
		}
	}
	return arr
}

// *** String :
func (q Queue[T]) String() string {
	sep := ","
	sb := strings.Builder{}
	for _, ele := range q {
		sb.WriteString(fmt.Sprintf("%v", ele))
		sb.WriteString(sep)
	}
	return strings.TrimRight(sb.String(), sep)
}
