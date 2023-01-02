package linkedlist_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/allanassis/go-ds/src/linkedlist"
)

var _ = Describe("Linkedlist", func() {
	var list *linkedlist.Linkedlist[int]

	BeforeEach(func() {
		list = linkedlist.NewLinkedList(0)
	})

	Describe("Adding to the linkedList", func() {
		Context("with numbers", func() {
			It("should have the itens added", func() {
				// arrange/acr
				list.Add(1)
				list.Add(2)
				// assert
				Expect(list.ToString()).To(Equal("0 - 1 - 2 - "))
			})
			It("should have the itens removed", func() {
				// arrange
				node1 := list.Add(1)
				list.Add(2)
				node2 := list.Add(3)
				// act
				list.Remove(node1.Id)
				list.Remove(node2.Id)
				// assert
				Expect(list.ToString()).To(Equal("0 - 2 - "))
			})
			It("should have the itens removed", func() {
				// arrange
				list.Add(1)
				node := list.Add(2)
				//act
				list.Update(node.Id, 3)
				// assert
				Expect(list.ToString()).To(Equal("0 - 1 - 3 - "))
			})
		})
	})
})
