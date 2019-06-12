package main

import "fmt"

func main() {
	fmt.Println([]int{-2,1,-3,4,-1,2,1,-5,4})
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

//如果结果是负数，直接滑动窗口吧。
func maxSubArray(nums []int) int {
	var res int = nums[0]
	var sum int = 0
	for _,num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if sum > res {
			res = sum
		}
	}
	return  res
}
