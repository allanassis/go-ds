package linkedlist

import (
	"fmt"

	"github.com/google/uuid"
)

type Node[T any] struct {
	Id   string
	data T
	next *Node[T]
}

func (node *Node[T]) Add(data T) *Node[T] {
	n := &Node[T]{
		Id:   uuid.NewString(),
		data: data,
		next: nil,
	}
	current := node
	for current.next != nil {
		current = current.next
	}
	current.next = n

	return n
}

func (node *Node[T]) Remove(id string) *Node[T] {
	if node.Id == id {
		node = node.next
	}

	previews := node
	current := node
	for current.next != nil {
		if current.Id == id {
			previews.next = current.next
			return current
		}
		previews = current
		current = current.next
	}
	return nil
}

func (node *Node[T]) Update(id string, data T) *Node[T] {
	current := node

	for current != nil {
		if current.Id == id {
			current.data = data
			return current
		}
		current = current.next
	}

	return nil
}

func (node *Node[T]) Print() {
	current := node
	for current != nil {
		fmt.Printf("Node: id: %s | value: %v\n", current.Id, current.data)
		current = current.next
	}
}

func (node *Node[T]) ToString() string {
	current := node
	raw := ""
	for current != nil {
		raw = raw + fmt.Sprintf("%v - ", current.data)
		current = current.next
	}
	return raw
}

func NewLinkedList[T any](data T) *Node[T] {
	return &Node[T]{
		Id:   uuid.NewString(),
		data: data,
		next: nil,
	}
}
