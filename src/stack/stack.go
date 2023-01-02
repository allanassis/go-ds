package stack

import (
	"fmt"

	"github.com/allanassis/go-ds/src/linkedlist"
)

type Stack[T any] struct {
	list *linkedlist.Linkedlist[T]
}

func (stack *Stack[T]) Push(value T) error {
	stack.list.Add(value)
	return nil
}

func (stack *Stack[T]) Pop() *linkedlist.Node[T] {
	tail := stack.list.Tail()
	removed := stack.list.Remove(tail.Id)
	return removed
}

func (stack *Stack[T]) Show() *linkedlist.Node[T] {
	tail := stack.list.Tail()
	fmt.Printf("Top: %v", tail)
}

func (stack *Stack[T]) Top() *linkedlist.Node[T] {
	tail := stack.list.Tail()
	return tail
}
