package main

import "fmt"

func main() {
	fmt.Println(
		bucketSort([]int{1,9,7,4,20,22,29,28,23,25,30,12,32,31,11,37,33,35,34,36}, 100))

	fmt.Println(removeDuplicates1([]int{1,2,2,3,4,4,4,4,4,5,5}))

	fmt.Println(removeElement([]int{2,2,1,2,13,1,4,5},2))

	fmt.Println(strStr("a","a"))
}
// 1000 / 10 = 100
// 99 / 100 = 0

func bucketSort(nums []int, bucketNum int) []int {
	var buckets [][]int
	for i:=0;i < bucketNum ; i++ {
		buckets = append(buckets,[]int{})
	}
	for _,num := range nums {
		index := num / bucketNum
		buckets[index] = append(buckets[index], num)
	}
	meg := []int{}
	for _,bucket := range buckets {
		fmt.Println("bucket :",bucket)
		meg = merge(meg, bucket)
	}
	return meg
}

func merge(bucket1, bucket2 []int)  (result []int) {
	l, r := 0, 0
	//合并。如果l 完了。仍然会去和最后的进行对比
	for l < len(bucket1) && r < len(bucket2)  {
		if bucket1[l] > bucket2[r] {
			result = append(result, bucket2[r])
			r++
		}else {
			result = append(result, bucket1[l])
			l++
		}
	}
	result = append(result, bucket1[l:]...)
	result = append(result, bucket2[r:]...)
	fmt.Println(result)
	return
}
// 1,3,3,4,5
func removeDuplicates(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i],nums[i+1:]...)
			i--
		}
	}
	return nums
}

func removeDuplicates1(nums []int) []int {
	i,j := 0,1
	for j < len(nums) {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return nums[:i]
}

func removeElement(nums []int, val int) []int {
	i,j := 0, 0
	for j < len(nums){
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return nums[:i]
}

func strStr(haystack string, needle string) int {
	nlen := len(needle)
	if nlen == 0 {
		return 0
	}
	for i:=0; i < len(haystack); i++ {
		if len(haystack) >= i+nlen && string(haystack[i:i+nlen]) == needle {
			return i
		}
	}
	return -1
}