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
			It("should return the root node", func() {
				// arrange/acr
				list.Add(1)
				list.Add(2)
				// assert
				Expect(list.Root().Data).To(Equal(0))
			})
			It("should return the tail node", func() {
				// arrange/acr
				list.Add(1)
				list.Add(2)
				// assert
				Expect(list.Tail().Data).To(Equal(2))
			})
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
			It("should have the itens updated", func() {
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
