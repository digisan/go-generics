package v2

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	s := Stack[string](dataStr)
	n := s.Push("e", "f", "g")
	fmt.Println(n)
	fmt.Println(s, s.Len())
}

func TestPop(t *testing.T) {
	s := Stack[string](dataStr)
	s.Push("e", "f", "g")
	fmt.Println(s, s.Len())
	n := s.Len()
	for i := 0; i <= n; i++ {
		ele, ok := s.Pop()
		fmt.Println(ele, ok, "left:", s.Len())
	}
}

func TestStackPeek(t *testing.T) {
	s := Stack[int](dataInt)
	n := s.Len()
	for i := 0; i <= n; i++ {
		ele, ok := s.Peek()
		fmt.Println(ele, ok, "left:", s.Len())
	}
}

func TestStackClear(t *testing.T) {
	s := Stack[int](dataInt)
	fmt.Println("s:", s, s.Len())
	o := s.Clear()
	fmt.Println("ori:", o, o.Len())
	fmt.Println("after clear, s:", s, s.Len())
}

func TestStackCopy(t *testing.T) {
	s := Stack[int](dataInt)
	fmt.Println("s:", s, s.Len())
	cp := s.Copy()
	s.Clear()
	fmt.Println("after clear, s:", s, s.Len())
	fmt.Println("cp", cp, cp.Len())
}

func TestStackSink(t *testing.T) {
	s := Stack[string](dataStr)
	fmt.Println("s:", s, s.Len())
	for _, e := range s.Sink() {
		fmt.Println(e)
	}
}
