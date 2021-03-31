package stack

//20 有效的括号
func isValid(s string) bool {
	stack := NewArrayStack(0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			//stack.Push(v)
		} else {
			topChar := stack.Pop()
			if topChar == '(' && v != ')' {
				return false
			} else if topChar == '[' && v != ']' {
				return false
			} else if topChar == '{' && v != '}' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
