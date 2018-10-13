package linkedlist

import "fmt"

type LinkedList struct {
	Length int
}

type Node struct {
	Next  *Node
	Value int
}

var head *Node
var current *Node

func (list *LinkedList) Add(value int) *Node {

	if list.Length == 0 {
		list.Length++
		head = &Node{nil, value}
		return head
	}

	current = head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &Node{nil, value}
	list.Length++

	return current
}

func (list *LinkedList) Get(index int) (int, error) {

	if index < 0 || index >= list.Length {
		return 0, fmt.Errorf("Index %d out of range", index)
	}

	current = head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, nil
}
