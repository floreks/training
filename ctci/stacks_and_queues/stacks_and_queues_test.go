package stacks_and_queues_test

import (
	. "github.com/floreks/training/ctci/stacks_and_queues"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StacksAndQueues", func() {
	It("Test Stack", func() {
		s := NewStack()
		s.Push(5)
		s.Push(2)
		s.Push(1)
		Expect(s.Size()).To(Equal(3))
		Expect(s.Peek()).To(Equal(1))
		Expect(s.Size()).To(Equal(3))
		Expect(s.Pop()).To(Equal(1))
		Expect(s.Size()).To(Equal(2))
	})

	It("Test StackQueue", func() {
		s := NewStackQueue()
		s.Add(5)
		s.Add(2)
		s.Add(1)
		Expect(s.Size()).To(Equal(3))
		Expect(s.Peek()).To(Equal(5))
		Expect(s.Size()).To(Equal(3))
		Expect(s.Remove()).To(Equal(5))
		Expect(s.Size()).To(Equal(2))
		Expect(s.Remove()).To(Equal(2))
		Expect(s.Size()).To(Equal(1))
	})

	It("Test SortStack", func() {
		s := NewStack()
		s.Push(4)
		s.Push(2)
		s.Push(3)
		s.Push(5)
		s.Push(1)

		s = SortStack(s)

		Expect(s.Pop()).To(Equal(1))
		Expect(s.Pop()).To(Equal(2))
		Expect(s.Pop()).To(Equal(3))
		Expect(s.Pop()).To(Equal(4))
		Expect(s.Pop()).To(Equal(5))
	})

	It("Test SortStack2", func() {
		s := NewStack()
		s.Push(1)
		s.Push(5)
		s.Push(5)
		s.Push(2)
		s.Push(4)

		s = SortStack(s)

		Expect(s.Pop()).To(Equal(1))
		Expect(s.Pop()).To(Equal(2))
		Expect(s.Pop()).To(Equal(4))
		Expect(s.Pop()).To(Equal(5))
		Expect(s.Pop()).To(Equal(5))
	})
})
