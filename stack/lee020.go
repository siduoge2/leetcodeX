package stack

import "strings"

func isValid(s string) bool {
	strs := strings.Split(s, "")
	mystack := StringStack{stck: nil}
	for _, str := range strs {
		if str == " " || str == "" {
			continue
		}
		if str == "(" || str == "[" || str == "{" {
			mystack.Push(str)
		} else if mystack.isMatched(str) {
			mystack.Pop()
		} else {
			return false
		}
	}
	if mystack.IsEmpty() {
		return true
	}
	return false
}
