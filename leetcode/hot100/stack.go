package hot100

import "fmt"

type stack struct {
	arr string
}

func newStack() *stack {
	return &stack{
		arr: "",
	}
}

func (s *stack) push(e byte) {
	s.arr = s.arr + string(e)
}

func (s *stack) pop() byte {
	if len(s.arr) == 0 {
		return ' '
	}
	b := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return b
}

//20 有效的括号
func isValid(s string) bool {
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}
	sta := newStack()
	for i, _ := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			sta.push(s[i])
			continue
		}
		char := sta.pop()
		if s[i] == ')' && char != '(' {
			return false
		} else if s[i] == ']' && char != '[' {
			return false
		} else if s[i] == '}' && char != '{' {
			return false
		}
	}
	if len(sta.arr) != 0 {
		return false
	}
	return true
}

func TestisValid() {
	s:="()"
	fmt.Println(isValid(s))
}
