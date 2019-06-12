package main

import "fmt"

func main() {
	a := []int{0}
	b := []int{2,5,6}
	merge(a,0,b, 1)
	fmt.Println(a)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i,j := m-1,n-1
	for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[i+j+1] = nums2[j]
			j--
		} else {
			nums1[i+j+1] = nums1[i]
			i--
		}
	}
	for j >= 0{
		nums1[j] = nums2[j]
		j--
	}
}