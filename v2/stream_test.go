package v2

import (
	"fmt"
	"testing"
)

func TestStream(t *testing.T) {

	str := "ABCD"
	r := StringToStream(str)
	s := StreamToString(r)
	fmt.Println(s)

	data := []byte(str)
	r = BytesToStream(data)
	data = StreamToBytes(r)
	fmt.Println(string(data))

}
