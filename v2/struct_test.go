package v2

import (
	"fmt"
	"testing"
	"time"
)

type TEST struct {
	a    int     `json:"email0" validate:"required,email,email-db"`
	A    float64 `json:"email1" validate:"email1,email-db"`
	B    uint
	s    string `json:"email2" validate:"email2,email-db"`
	S    string `json:"email3" validate:"email3,email-db"`
	Sub  SUB
	sub  SUB
	RUNE rune
	BYTE byte
	Arr  []int
	Tm   time.Time
}

type SUB struct {
	c      int
	C      int
	d      complex128
	D      complex128
	M      string
	Ssub   SSUB
	SubArr []SSUB
}

type SSUB struct {
	Z float64
}

type TA struct {
	A int
}

type TB struct {
	B float64
}

type TC struct {
	C string
}

type ABC struct {
	TA
	TB
	TC
}

func TestPartialAsMap(t *testing.T) {

	test := &TEST{
		a: 1,
		A: 12,
		s: "s",
		S: "SS",
		Sub: SUB{
			c: 9,
			C: 99,
			d: 4 + 5i,
			D: 5 + 6i,
			M: "sub",
			Ssub: SSUB{
				Z: 1.11,
			},
		},
		sub: SUB{},
	}

	fmt.Println(PartialAsMap(test, "A", "S", "B", "Sub.C", "Sub.M", "Sub.Ssub.Z"))
}

func TestFieldValueOnMap(t *testing.T) {
	m := map[string]any{
		"a": 1,
		"A": 11,
		"b": 2,
	}

	v, err := FieldValue(m, "A")
	fmt.Printf("FieldValue applies to map: %v\n", v)
	fmt.Printf("Err: %v\n", err)
}

func TestFields(t *testing.T) {

	test := &TEST{
		a: 1,
		A: 12,
		s: "s",
		S: "SS",
		Sub: SUB{
			c:      9,
			C:      99,
			d:      4 + 5i,
			D:      5 + 6i,
			SubArr: []SSUB{{Z: 1.111}, {Z: 2.222}, {Z: 3.333}, {Z: 4.444}},
		},
		sub: SUB{},
		Arr: []int{11, 22, 33},
	}

	v, err := PathValue(test, "Sub.C")
	fmt.Printf("Sub.C: %v\n", v)
	fmt.Printf("Err: %v\n", err)

	fmt.Println()

	v, err = PathValue(test, "Arr.2")
	fmt.Printf("Arr.2: %v\n", v)
	fmt.Printf("Err: %v\n", err)

	fmt.Println()

	v, err = PathValue(test, "Sub.SubArr.2.Z")
	fmt.Printf("Sub.SubArr.2.Z: %v\n", v)
	fmt.Printf("Err: %v\n", err)

	fmt.Println()

	sub, _ := FieldValue(test, "Sub")
	fmt.Printf("%v\n", sub)

	fmt.Println(FieldValue(sub, "D"))

	fmt.Println()

	fmt.Println(FieldValue(*test, "Z"))

	fmt.Println()

	for _, f := range Fields(test) {
		fmt.Println(FieldValue(test, f))
	}

	fmt.Println()

	for _, f := range Fields(*test) {
		fmt.Println(FieldValue(*test, f))
	}

	fmt.Println()

	fmt.Println(ValidatorTags(test, "validate", "email"))
	fmt.Println(Tags(*test, "validate", "required", "email-db"))

	fmt.Println()

	fmt.Println(JsonTags(test))
	fmt.Println(Tags(*test, "json", "email3"))
}

func TestSetFieldValue(t *testing.T) {

	test := TEST{
		a: 1,
		A: 12.1,
		s: "s",
		S: "SS",
		Sub: SUB{
			c: 9,
			C: 99,
			d: 4 + 5i,
			D: 5 + 6i,
		},
		sub: SUB{},
		Arr: []int{11, 22, 33},
	}

	if err := SetFieldValue(&test, "Tm", "1988-03-01 13:02:11"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", test)

	if err := SetFieldValue(&test, "RUNE", "1233335"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", test)

	testArr := []int{9, 8, 7}
	if err := SetFieldValue(testArr, "1", 888); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", testArr)

	abc := ABC{}
	if err := SetFieldValue(&abc, "A", 1000); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", abc)

	// if err := SetFieldValue(&test, "BYTE", "1725.6"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%+v\n", test)

	if err := SetFieldValue(&test, "NOFIELD", "NONONO"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", test)
}
