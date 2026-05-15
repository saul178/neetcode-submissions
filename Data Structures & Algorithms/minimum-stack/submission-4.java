class MinStack {
	private Stack<Integer> stack;
	private Stack<Integer> minStack;

	public MinStack() {
		this.stack = new Stack<>();
		this.minStack = new Stack<>();
	}

	public void push(int val) {
		stack.push(val);
		if (minStack.isEmpty() || val <= minStack.peek()) {
			minStack.push(val);
		}
	}

	public void pop() {
		if (!stack.isEmpty()) {
			int top = stack.pop();
			if (top == minStack.peek()) {
				minStack.pop();
			}
		} else {
			System.out.println("Stack is empty!");
			return;
		}
	}

	public int top() {
		if (!stack.isEmpty()) {
			return stack.peek();
		} else {
			System.out.println("Stack is empty!");
			return -1;
		}
	}

	public int getMin() {
		return minStack.peek();
	}
}
