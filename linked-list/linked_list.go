package linkedlist

import "fmt"

type LinkedList struct {
	Length  int
	head    *Node
	current *Node
}

type Node struct {
	Next  *Node
	Value int
}

func (list *LinkedList) Add(value int) *Node {

	if list.Length == 0 {
		list.Length++
		list.head = &Node{nil, value}
		return list.head
	}

	list.current = list.head
	for list.current.Next != nil {
		list.current = list.current.Next
	}

	list.current.Next = &Node{nil, value}
	list.Length++

	return list.current
}

func (list *LinkedList) Get(index int) (int, error) {

	if index < 0 || index >= list.Length {
		return 0, fmt.Errorf("Index %d out of range", index)
	}

	list.current = list.head
	for i := 0; i < index; i++ {
		list.current = list.current.Next
	}

	return list.current.Value, nil
}
