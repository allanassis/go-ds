package linkedlist_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/allanassis/go-ds/src/linkedlist"
)

var _ = Describe("Linkedlist", func() {
	var root *linkedlist.Node[int]

	BeforeEach(func() {
		root = linkedlist.NewLinkedList(0)
	})

	Describe("Adding to the linkedList", func() {
		Context("with numbers", func() {
			It("should have the itens added", func() {
				// arrange/acr
				root.Add(1)
				root.Add(2)
				// assert
				Expect(root.ToString()).To(Equal("0 - 1 - 2 - "))
			})
			It("should have the itens removed", func() {
				// arrange
				root.Add(1)
				node := root.Add(2)
				// act
				root.Remove(node.Id)
				// assert
				Expect(root.ToString()).To(Equal("0 - 1 - "))
			})
			It("should have the itens removed", func() {
				// arrange
				root.Add(1)
				node := root.Add(2)
				//act
				root.Update(node.Id, 3)
				// assert
				Expect(root.ToString()).To(Equal("0 - 1 - 3 "))
			})
		})
	})
})
