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
	var data []int = []int{1, 2, 3, 4, 5}
	for pair := range IterPair(data) {
		fmt.Printf("%+v\n", pair)
	}
}

func TestRangeTriple(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5}
	for triple := range IterTriple(data) {
		fmt.Printf("%+v\n", triple)
	}
}

func TestRangeCache(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5}
	for cache := range IterCache(data, 3, 3, 999) {
		fmt.Printf("%+v\n", cache)
	}
}
