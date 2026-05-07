type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{queue: []int{}}
}

func (q *MyStack) Push(x int) {
	q.queue = append(q.queue, x)
    for i := 0; i < len(q.queue)-1; i++ {
    q.queue = append(q.queue, q.queue[0])
    q.queue = q.queue[1:]
    }
}

func (q *MyStack) Pop() int {
	if q.Empty() {
        fmt.Errorf("queue is empty")
		return -1
	}
	topElem := q.queue[0]
	q.queue[0] = 0
	q.queue = q.queue[1:]
	return topElem
}

func (q *MyStack) Top() int {
    if q.Empty() {
        fmt.Errorf("queue is empty")
        return -1
    }
	return q.queue[0]
}

func (q *MyStack) Empty() bool {
	return len(q.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
