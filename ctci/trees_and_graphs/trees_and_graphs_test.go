package trees_and_graphs_test

import (
	. "github.com/floreks/training/ctci/trees_and_graphs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TreesAndGraphs", func() {
	It("Test MinHeap", func() {
		arr := []int{87, 7, 90, 50, 55, 4}
		heap := NewHeapFromArray(arr...)

		Expect(heap.Delete()).To(Equal(4))
		Expect(heap.Peek()).To(Equal(7))
		heap.Insert(2)
		Expect(heap.Peek()).To(Equal(2))
		Expect(heap.Delete()).To(Equal(2))
		Expect(heap.Delete()).To(Equal(7))
		Expect(heap.Delete()).To(Equal(50))
		Expect(heap.Delete()).To(Equal(55))
	})
})
