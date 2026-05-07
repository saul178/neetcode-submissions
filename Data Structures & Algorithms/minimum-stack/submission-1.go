type MinStack struct {
	elements []int
}

func Constructor() *MinStack {
	stack := []int{}
	return &MinStack{elements: stack}
}

func (ms *MinStack) Push(val int) {
	ms.elements = append(ms.elements, val)
}

func (ms *MinStack) Pop() int {
	if len(ms.elements) > 0 {
		val := ms.elements[len(ms.elements)-1]
		ms.elements = ms.elements[:len(ms.elements)-1]
		return val
	} else {
		fmt.Println("stack is empty")
		return -1
	}
}

func (ms *MinStack) Top() int {
	if len(ms.elements) > 0 {
		return ms.elements[len(ms.elements)-1]
	} else {
		fmt.Println("stack is empty")
		return -1
	}
}

func (ms *MinStack) GetMin() int {
	tmpStack := Constructor()
	minimum := ms.Top()

	for len(ms.elements) > 0 {
		val := ms.Pop()
		minimum = min(minimum, val)
		tmpStack.Push(val)
	}

	for len(tmpStack.elements) > 0 {
		val := tmpStack.Pop()
		ms.Push(val)
	}
	return minimum
}
