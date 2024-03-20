package gogenerics

import (
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {

	// r := NewFailResult(100.0, "sorry failed")
	r := NewResult(1000, nil)

	var (
		err error
	)

	if ok, v := r.SafeValue(&err); ok {
		fmt.Printf("\n%v SafeValue ok\n", v)
	} else {
		fmt.Printf("cannot fetch valid value from result (invalid value is %v)\nerr: %v\n\n", v, err)
	}
	{
		v := r.Expect("oh!!!")
		fmt.Printf("\n%v Expect ok\n", v)
	}
	{
		v := r.Unwrap()
		fmt.Printf("\n%v Unwrap ok\n", v)
	}
}
