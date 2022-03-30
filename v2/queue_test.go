package v2

import (
	"fmt"
	"testing"
)

var (
	dataStr = []string{"a", "b", "c", "d"}
	dataInt = []int{1, 2, 3, 4, 5}
)

func TestEnqueue(t *testing.T) {
	q := Queue[string](dataStr)
	n := q.Enqueue("e", "f", "g")
	fmt.Println(n)
	fmt.Println(q, q.Len())
}

func TestDequeue(t *testing.T) {
	q := Queue[string](dataStr)
	q.Enqueue("e", "f", "g")
	fmt.Println(q, q.Len())
	n := q.Len()
	for i := 0; i <= n; i++ {
		ele, ok := q.Dequeue()
		fmt.Println(ele, ok, "left:", q.Len())
	}
}

func TestQueuePeek(t *testing.T) {
	q := Queue[int](dataInt)
	n := q.Len()
	for i := 0; i <= n; i++ {
		ele, ok := q.Peek()
		fmt.Println(ele, ok, "left:", q.Len())
	}
}

func TestQueueClear(t *testing.T) {
	q := Queue[int](dataInt)
	fmt.Println("q:", q, q.Len())
	o := q.Clear()
	fmt.Println("ori:", o, o.Len())
	fmt.Println("after clear, q:", q, q.Len())
}

func TestQueueCopy(t *testing.T) {
	q := Queue[int](dataInt)
	fmt.Println("q:", q, q.Len())
	cp := q.Copy()
	q.Clear()
	fmt.Println("after clear, q:", q, q.Len())
	fmt.Println("cp:", cp, cp.Len())
}
