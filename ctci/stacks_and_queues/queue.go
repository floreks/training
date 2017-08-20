package stacks_and_queues

type StackQueue struct {
	new, old *Stack
}

func (self *StackQueue) Size() int {
	return self.new.Size() + self.old.Size()
}

func (self *StackQueue) Add(val int) {
	self.new.Push(val)
}

func (self *StackQueue) Peek() int {
	self.rollStacks()
	return self.old.Peek()
}

func (self *StackQueue) Remove() int {
	self.rollStacks()
	return self.old.Pop()
}

func (self *StackQueue) rollStacks() {
	if self.old.Size() == 0 {
		for self.new.Size() > 0 {
			self.old.Push(self.new.Pop())
		}
	}
}

func NewStackQueue() *StackQueue {
	return &StackQueue{
		old: NewStack(),
		new: NewStack(),
	}
}
