package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("hello World1 "))
}

func lengthOfLastWord(s string) int {
	tmpLen := 0
	for _,str := range []rune(strings.TrimRight(s, " ")) {
		if str == ' ' {
			tmpLen = 0
		} else {
			tmpLen++
		}
	}
	return tmpLen
}
