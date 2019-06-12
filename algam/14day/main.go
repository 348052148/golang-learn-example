package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}
func removeOuterParentheses(S string) string {

	new := make([]rune,0)
	num :=0
	last := 0
	for i,v := range []rune(S) {
		if v == '(' {
			num++
		}else if v == ')' {
			num--
		}
		if (num == 0) {
			new = append(new, []rune(S)[last+1:i]...)
			last=i+1
		}
	}
	return string(new)
}