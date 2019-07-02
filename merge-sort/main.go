package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{3,5,1,4,2,7,6,8,9,3,4}))
}


func mergeSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	left := mergeSort(nums[:l/2])
	right := mergeSort(nums[l/2:])
	a := merge(left, right)
	fmt.Println(a)
	return a
}

func merge(left, right []int) []int {
	var tmp []int
	l,r :=0, 0
	for l < len(left) && r < len(right)  {
		if left[l] < right[r] {
			tmp = append(tmp,left[l])
			l++
		}else{
			tmp = append(tmp,right[r])
			r++
		}
	}
	if l < len(left) {
		tmp = append(tmp, left[l:]...)
	}else {
		tmp = append(tmp, right[r:]...)
	}

	return tmp
}