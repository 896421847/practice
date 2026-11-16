package main

import "fmt"

func isvalid(s string) bool {

	stack := []rune{}
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false

			}

			top := stack[len(stack)-1]
			if ch == ')' && top != '(' {
				return false

			}
			if ch == ']' && top != '[' {
				return false

			}
			if ch == '}' && top != '{' {
				return false

			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0
}
func main() {
	s := "({})[]"
	fmt.Println(isvalid(s))
}
