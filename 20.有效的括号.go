/*
	20.有效的括号

	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

   	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。

	示例 1:
		输入: "()"
		输出: true
	示例 2:
		输入: "()[]{}"
		输出: true
	示例 3:
		输入: "(]"
		输出: false
	示例 4:
		输入: "([)]"
		输出: false
	示例 5:
		输入: "{[]}"
		输出: true
*/

package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	symbolMap := map[string]string {
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
		s.data = s.data[0:s.Length()-1]
		return lastCharacter
	}
}

func (s *Stack) Length () int {
	return len(s.data)
}
