package v2

import (
	"fmt"
	"testing"
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
}

type SUB struct {
	c    int
	C    int
	d    complex128
	D    complex128
	M    string
	Ssub SSUB
}

type SSUB struct {
	Z float64
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
			c: 9,
			C: 99,
			d: 4 + 5i,
			D: 5 + 6i,
		},
		sub: SUB{},
	}

	v, err := PathValue(test, "Sub.C")
	fmt.Printf("Sub.C: %v\n", v)
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

	test := &TEST{
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
	}

	if ok, err := SetFieldValue(test, "RUNE", "1233335"); !ok {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", *test)

	if ok, err := SetFieldValue(test, "BYTE", "1725"); !ok {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", *test)
}
