package trees_and_graphs

type MinHeap struct {
	values []int
}

func (self *MinHeap) Insert(val int) {
	if len(self.values) == 0 {
		self.values = append(self.values, val)
		return
	}

	// Add element at the end of min-heap (bottom rightmost element)
	self.values = append(self.values, val)

	valIdx := len(self.values) - 1
	parentIdx := valIdx / 2
	for parentIdx != valIdx && self.values[parentIdx] > val {
		self.values[parentIdx], self.values[valIdx] = self.values[valIdx], self.values[parentIdx]
		valIdx = parentIdx
		parentIdx = valIdx / 2
	}
}

func (self *MinHeap) Peek() int {
	return self.values[0]
}

func (self *MinHeap) Delete() int {
	if len(self.values) == 0 {
		panic("Can not delete from empty heap")
	}

	min := self.values[0]

	// Replace root with bottom rightmost element
	self.values[0] = self.values[len(self.values)-1]
	// Remove last element
	self.values = self.values[:len(self.values)-1]
	// rebalance the heap
	lowestIdx := 0
	parentIdx := 0
	leftChildIdx := 2*lowestIdx + 1
	rightChildIdx := 2*lowestIdx + 2
	for true {
		if leftChildIdx < len(self.values) && self.values[lowestIdx] > self.values[leftChildIdx] {
			lowestIdx = leftChildIdx
		}

		if rightChildIdx < len(self.values) && self.values[lowestIdx] > self.values[rightChildIdx] {
			lowestIdx = rightChildIdx
		}

		if lowestIdx != parentIdx {
			self.values[parentIdx], self.values[lowestIdx] = self.values[lowestIdx], self.values[parentIdx]
			parentIdx = lowestIdx
			leftChildIdx = 2*lowestIdx + 1
			rightChildIdx = 2*lowestIdx + 2
		} else {
			break
		}
	}

	return min
}

func (self *MinHeap) heapify(arr []int, i int) {
	lowestIdx := i
	leftChildIdx := 2*i + 1
	rightChildIdx := 2*i + 2

	if leftChildIdx < len(arr) && arr[leftChildIdx] < arr[lowestIdx] {
		lowestIdx = leftChildIdx
	}

	if rightChildIdx < len(arr) && arr[rightChildIdx] < arr[lowestIdx] {
		lowestIdx = rightChildIdx
	}

	if lowestIdx != i {
		arr[i], arr[lowestIdx] = arr[lowestIdx], arr[i]
		self.heapify(self.values, lowestIdx)
	}
}

func NewHeap() *MinHeap {
	return &MinHeap{values: make([]int, 0)}
}

func NewHeapFromArray(arr ...int) *MinHeap {
	heap := &MinHeap{values: arr}
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heap.heapify(arr, i)
	}

	return heap
}
