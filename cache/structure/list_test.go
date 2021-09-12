package structure

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_MyDoubleEndList(t *testing.T) {
	ll := NewList()
	node1 := ll.AddLast(5)
	_ = ll.InsertBefore(9, node1)
	ll.AddFirst(12)
	ll.RemoveToFirst(node1)

	fmt.Printf("base info, length = %v, first value = %v and last value = %v\n",
		ll.Len(), ll.GetFirst().pair, ll.GetLast().pair)
	fmt.Println("------concrete info:")
	for i := ll.Len(); i > 0; i-- {
		fmt.Println(ll.Remove(ll.GetFirst()))
	}
}

func Test_LibraryList(t *testing.T) {
	ll := list.New()
	node1 := ll.PushBack(5)
	node2 := ll.InsertAfter(6, node1)

	fmt.Printf("length: %d\n", ll.Len())
	fmt.Printf("back: %v\n", ll.Back().Value)
	fmt.Printf("front: %v\n", ll.Front().Value)

	fmt.Printf("node1: %v\n", node1.Value)
	fmt.Printf("node2: %v\n", node2.Value)
}
