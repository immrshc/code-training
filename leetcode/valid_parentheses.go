package leetcode

type Stack []string

func (s *Stack) Empty() bool {
	size := len(*s)
	return size == 0
}

func (s *Stack) Pop() string {
	if s.Empty() {
		return ""
	}
	size := len(*s)
	var str string
	str = (*s)[size-1]
	*s = (*s)[:size-1]
	return str
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

var closeToOpen = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func ValidParentheses(s string) bool {
	stack := &Stack{}
	for _, b := range s {
		bracket := string(b)
		if open, ok := closeToOpen[bracket]; ok {
			if stack.Pop() != open {
				return false
			}
		} else {
			stack.Push(bracket)
		}
	}
	return stack.Empty()
}
