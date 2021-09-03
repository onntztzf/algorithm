package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

//20. 有效的括号
//link: https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	symbolMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	stack := NewStack()
	for _, val := range s {
		character := string(val)
		if rightSymbol := symbolMap[character]; rightSymbol != "" {
			stack.Push(character)
		} else if symbolMap[stack.Pop()] != character {
			return false
		}
	}
	return 0 == stack.Length()
}

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() string {
	if s.Length() == 0 {
		return ""
	} else {
		lastCharacter := s.data[s.Length()-1]
		s.data = s.data[0 : s.Length()-1]
		return lastCharacter
	}
}

func (s *Stack) Length() int {
	return len(s.data)
}
