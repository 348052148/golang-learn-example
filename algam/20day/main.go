package main

import (
	"fmt"
)

func main() {
	fmt.Println(titleToNumber("AA"))
}

func titleToNumber(s string) int {
	temp := 0
	for _,v := range s {
		num := int(v - 64)
		temp = temp * 26 + num
	}
	return temp
}