package main

import "fmt"

func main() {
	nums := []int{-1}
	//rotate(nums,2)
	//fmt.Println(nums)
	rotate1(nums, 2)
	//fmt.Println(nums)
}

func rotate1(nums []int, k int)  {
	reverse(nums, 0, len(nums)-1)
	fmt.Println(nums)
	reverse(nums,0, k-1)
	fmt.Println(nums)
	reverse(nums, k, len(nums)-1)
	fmt.Println(nums)
}

func reverse(nums []int, start, end int){
	for start < end {
		nums[start],nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int)  {
	l := len(nums)
	k = k % l
	count := 0
	for start := 0 ; count < l ; start++ {
		prev := nums[start]
		current := (k + start) % l
		for count < l && current != start {
			fmt.Println(current)

			t := nums[current]
			nums[current] = prev
			prev = t

			current = (k + current) % l
			count++
		}
		//fmt.Println("", prev, (k + 1) % l, count)
		//break
	}
}