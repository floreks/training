package stacks_and_queues

func SortStack(stack *Stack) *Stack {
	if stack.Size() == 0 {
		return stack
	}

	sortedStack := NewStack()
	sortedStack.Push(stack.Pop())

	popped := 0
	for !stack.IsEmpty() {
		popped = stack.Pop()
		if popped < sortedStack.Peek() {
			sortedStack.Push(popped)
		} else {
			for !sortedStack.IsEmpty() && sortedStack.Peek() < popped {
				stack.Push(sortedStack.Pop())
			}

			sortedStack.Push(popped)
		}
	}

	return sortedStack
}
