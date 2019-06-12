package main

import (
	"fmt"
	"strconv"
)
type A struct {
	name string
	sub Asub
}
type Asub struct {
	c int
} 
func main() {
	fmt.Println(searchInsert1([]int{1,3,5,6}, 3))
	a := A{"123", Asub{10}}
	b := a
	a.sub.c = 20
	fmt.Println(b.sub)

	fmt.Println(countAndSay(3))
}

func searchInsert(nums []int, target int) int {
	start, end, index := 0, len(nums),0
	for {
		if start >= end {
			break
		}
		mid := (end + start) / 2
		if nums[mid] > target {
			end = mid
			index = mid
		} else if nums[mid] < target{
			start = mid + 1
			index = mid + 1;
		} else {
			index = mid
			return index;
		}
	}
	return index
}

func searchInsert1(nums []int, target int) int {
	if len(nums) == 1 {
		return 0
	}
	start, end := 0, len(nums)
	mid := (start + end)/2

	fmt.Println(mid)
	if nums[mid] == target {
		return mid
	}
	if start == mid {
		if nums[mid] > target {
			return mid + 1
		}else {
			return mid
		}
	}
	if nums[mid] > target {
		return searchInsert1(nums[:mid], target)
	} else {
		return mid + searchInsert1(nums[mid:], target)
	}
}


func countAndSay(n int) string {
	si := "1"
	var cur rune
	for {
		if n < 0 {
			break
		}
		cur = []rune(si)[0]
		count := 1
		curent := ""
		for i, c := range []rune(si) {
			if i == 0 {
				continue
			}
			if c != cur {
				curent = curent + strconv.Itoa(count) + string(cur)
				cur = c
				count = 1
			}else {
				count++
			}
		}
		curent = curent + strconv.Itoa(count) + string(cur)
		si = curent
		n--
	}
	return si
}