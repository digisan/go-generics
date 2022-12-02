package v2

import (
	"fmt"
	"testing"
)

var (
	fPln = fmt.Println
)

func TestIter(t *testing.T) {

	for i := range Iter(-10) {
		fPln(i)
	}
	fPln(" ------------ ")

	for i := range Iter(1, 5) {
		fPln(i)
	}
	fPln(" ------------ ")

	for i := range Iter(2, 3, 10) {
		fPln(i)
	}
	fPln(" ------------ ")

	func(slc ...int) {
		for _, a := range slc {
			fPln(a)
		}
	}(IterToSlc(11, -3, 2)...)

	// ----------------------- //
	fPln(" ************************************* ")

	for i := range Iter(10, -2, -1) {
		fPln(i)
	}
	fPln(" ------------ ")

	for i := range Iter(10, 14) {
		fPln(i)
	}
}

func TestRangePair(t *testing.T) {
	var params []int = []int{2, 4, 6, 8, 10}
	for pair := range IterPair(params...) {
		fmt.Printf("%+v\n", pair)
	}
}

// func TestRangeTriple(t *testing.T) {
// 	var params []int = []int{1, 3, 5, 7, 9, 11, 13, 15}
// 	for pair := range IterTriple(params...) {
// 		fmt.Printf("%+v\n", pair)
// 	}
// }
