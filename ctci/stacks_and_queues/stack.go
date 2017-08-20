package stacks_and_queues

type Stack struct {
	s []int
}

func (self *Stack) Size() int {
	return len(self.s)
}

func (self *Stack) Pop() int {
	if self.Size() > 0 {
		ret := self.s[self.Size()-1]
		self.s = self.s[:self.Size()-1]
		return ret
	}

	panic("Can not pop from an empty stack")
}

func (self *Stack) Push(n int) {
	self.s = append(self.s, n)
}

func (self *Stack) Peek() int {
	if self.Size() > 0 {
		ret := self.s[self.Size()-1]
		return ret
	}

	panic("Can not peek on an empty stack")
}

func NewStack() *Stack {
	return &Stack{s: make([]int, 0)}
}
