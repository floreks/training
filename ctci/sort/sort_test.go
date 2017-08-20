package sort_test

import (
	. "github.com/floreks/training/ctci/sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sort", func() {
	It("Test Quicksort", func() {
		arr := []int{5,2,1,6,3,2,6}
		Quicksort(arr)
		Expect(arr).To(Equal([]int{1,2,2,3,5,6,6}))
	})
})
