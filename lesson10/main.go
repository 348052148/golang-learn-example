package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6,7,8}
	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) (int, int)  {
	maps := make(map[int]int)
	for i,v := range nums {
		v1 := target - v
		if i1,ok := maps[v1]; ok {
			return i,i1
		}
		maps[v] = i
	}
	return 0,0
}