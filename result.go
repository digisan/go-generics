package gogenerics

import (
	"errors"
	"fmt"
)

type Result[T any] struct {
	value T
	err   error
}

func (r *Result[T]) Unwrap() T {
	if r.err == nil {
		return r.value
	}
	panic(r.err.Error())
}

func (r *Result[T]) Expect(failMsg string) T {
	if r.err == nil {
		return r.value
	}
	err := fmt.Errorf("%s: %w", failMsg, r.err)
	if len(failMsg) != 0 {
		panic(err.Error())
	}
	panic(errors.Unwrap(err).Error())
}

func (r *Result[T]) SafeValue(err *error) (bool, T) {
	defer func() {
		if err == nil {
			panic("'err' cannot be nil for holding error")
		}
		*err = r.err
	}()
	return r.err == nil, r.value
}

func NewResult[T any](value T, err error) Result[T] {
	return Result[T]{
		value: value,
		err:   err,
	}
}

func NewFailResult[T any](value T, failMsg string) Result[T] {
	return NewResult[T](value, errors.New(failMsg))
}
