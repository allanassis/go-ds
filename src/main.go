package main

import "github.com/allanassis/go-ds/src/linkedlist"

func main() {
	l := linkedlist.NewLinkedList("bla")
	node := l.Add("aaa")
	l.Add("Some")
	l.Update(node.Id, "BBB")
	// l.Remove(node.Id)
	l.Print()
}
