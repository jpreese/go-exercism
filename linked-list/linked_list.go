package linkedlist

import "errors"

var ErrIndexOutOfRange = errors.New("Index out of range")

type LinkedList struct {
	Length  int
	head    *node
	current *node
}

type node struct {
	next  *node
	value int
}

func (list *LinkedList) Add(value int) *node {
	if list.Length == 0 {
		list.Length++
		list.head = &node{nil, value}
		return list.head
	}

	list.current = list.head
	for list.current.next != nil {
		list.current = list.current.next
	}

	list.current.next = &node{nil, value}
	list.Length++

	return list.current
}

func (list *LinkedList) Get(index int) (int, error) {

	if index < 0 || index >= list.Length {
		return 0, ErrIndexOutOfRange
	}

	list.current = list.head
	for i := 0; i < index; i++ {
		list.current = list.current.next
	}

	return list.current.value, nil
}
