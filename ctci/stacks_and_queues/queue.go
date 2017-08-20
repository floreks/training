package stacks_and_queues

type StackQueue struct {
	newStack, oldStack *Stack
}

func (self *StackQueue) Size() int {
	return self.newStack.Size() + self.oldStack.Size()
}

func (self *StackQueue) Add(val int) {
	self.newStack.Push(val)
}

func (self *StackQueue) Peek() int {
	self.rollStacks()
	return self.oldStack.Peek()
}

func (self *StackQueue) Remove() int {
	self.rollStacks()
	return self.oldStack.Pop()
}

func (self *StackQueue) rollStacks() {
	if self.oldStack.Size() == 0 {
		for self.newStack.Size() > 0 {
			self.oldStack.Push(self.newStack.Pop())
		}
	}
}

func NewStackQueue() *StackQueue {
	return &StackQueue{
		oldStack: NewStack(),
		newStack: NewStack(),
	}
}
