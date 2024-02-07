package v2

import (
	"fmt"
	"strings"
)

type Stack[T any] []T

// *** Push
func (stk *Stack[T]) Push(items ...T) int {
	*stk = append(*stk, items...)
	return len(items)
}

// *** Len :
func (stk *Stack[T]) Len() int {
	return len(*stk)
}

// *** Pop :
func (stk *Stack[T]) Pop() (T, bool) {
	if stk.Len() > 0 {
		last := (*stk)[stk.Len()-1]
		*stk = (*stk)[:stk.Len()-1]
		return last, true
	}
	return *new(T), false
}

// *** Peek :
func (stk *Stack[T]) Peek() (T, bool) {
	if stk.Len() > 0 {
		return (*stk)[stk.Len()-1], true
	}
	return *new(T), false
}

// *** Clear :
func (stk *Stack[T]) Clear() Stack[T] {
	cp := stk.Copy()
	*stk = Stack[T]{}
	return cp
}

// *** Copy :
func (stk *Stack[T]) Copy() Stack[T] {
	tmp := make([]T, stk.Len())
	copy(tmp, *stk)
	return Stack[T](tmp)
}

// *** Sink :
func (stk *Stack[T]) Sink() []T {
	n := stk.Len()
	arr := make([]T, 0, n)
	for {
		if ele, ok := stk.Pop(); ok {
			arr = append(arr, ele)
		} else {
			break
		}
	}
	return arr
}

// String :
func (stk Stack[T]) String() string {
	sep := ","
	sb := strings.Builder{}
	for _, ele := range stk {
		sb.WriteString(fmt.Sprintf("%v", ele))
		sb.WriteString(sep)
	}
	return strings.TrimRight(sb.String(), sep)
}
