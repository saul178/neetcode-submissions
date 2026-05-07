func isValid(s string) bool {
   	if len(s) <= 1 {
		return false
	}

	stack := []byte{}
	for i := range s {
		switch s[i] {
		case '(':
			stack = append(stack, s[i])
		case '{':
			stack = append(stack, s[i])
		case '[':
			stack = append(stack, s[i])
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
 
}
