package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	var l = len(digits)
	var tmpSlice []int
	if l == 0 {
		return digits
	}
	if digits[l-1] == 9 {
		if (l == 1) {
			digits = append([]int{0},digits...)
			l = len(digits)
			}
		tmpSlice = plusOne(digits[:l-1])
		digits[l-1] = 0
	}else {
		digits[l-1]++
		tmpSlice = digits[:l-1]
	}

	return append(tmpSlice,digits[l-1])
}
