class Solution {
    public boolean isValid(String s) {
		char[] charArr = s.toCharArray();
		Stack<Character> stack = new Stack<>();
		for (int i = 0; i < charArr.length; i++) {
			switch (charArr[i]) {
				case '(':
					stack.push(charArr[i]);
					break;
				case '{':
					stack.push(charArr[i]);
					break;
				case '[':
					stack.push(charArr[i]);
					break;
				case ')':
					if (!stack.isEmpty() && stack.peek() == '(') {
						stack.pop();
					} else {
						return false;
					}
					break;
				case '}':
					if (!stack.isEmpty() && stack.peek() == '{') {
						stack.pop();
					} else {
						return false;
					}
					break;
				case ']':
					if (!stack.isEmpty() && stack.peek() == '[') {
						stack.pop();
					} else {
						return false;
					}
					break;
			}
		}
		return stack.isEmpty();
    }
}
