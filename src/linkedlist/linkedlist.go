package linkedlist

import (
	"fmt"

	"github.com/google/uuid"
)

type Linkedlist[T any] struct {
	root *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	Id   string
	Data T
	next *Node[T]
}

func (list *Linkedlist[T]) Root() *Node[T] {
	return list.root
}

func (list *Linkedlist[T]) Tail() *Node[T] {
	return list.tail
}

func (list *Linkedlist[T]) Add(data T) *Node[T] {
	n := &Node[T]{
		Id:   uuid.NewString(),
		Data: data,
		next: nil,
	}
	list.tail = n

	current := list.root
	for current.next != nil {
		current = current.next
	}
	current.next = n

	return n
}

func (list *Linkedlist[T]) Remove(id string) *Node[T] {
	if list.root.Id == id {
		list.root = list.root.next
	}

	previews := list.root
	current := list.root
	for current != nil {
		if current.Id == id {
			if current.Id == list.tail.Id {
				list.tail = previews
			}
			previews.next = current.next
			return current
		}
		previews = current
		current = current.next
	}
	return nil
}

func (list *Linkedlist[T]) Update(id string, data T) *Node[T] {
	current := list.root

	for current != nil {
		if current.Id == id {
			current.Data = data
			return current
		}
		current = current.next
	}

	return nil
}

func (list *Linkedlist[T]) Print() {
	current := list.root
	for current != nil {
		fmt.Printf("Node: id: %s | value: %v\n", current.Id, current.Data)
		current = current.next
	}
}

func (list *Linkedlist[T]) ToString() string {
	current := list.root
	raw := ""
	for current != nil {
		raw = raw + fmt.Sprintf("%v - ", current.Data)
		current = current.next
	}
	return raw
}

func NewLinkedList[T any](data T) *Linkedlist[T] {
	node := &Node[T]{
		Id:   uuid.NewString(),
		Data: data,
		next: nil,
	}

	return &Linkedlist[T]{
		root: node,
		tail: node,
	}
}
