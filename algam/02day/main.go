package main

import "fmt"

type TreeNode struct {
	value int
	left, right *TreeNode
}
/*
1
2     1
1 2  2  1
 */

func main() {
	root := &TreeNode{1, nil, nil}
	c1_l := &TreeNode{2, nil, nil}
	c1_r := &TreeNode{1, nil, nil}

	c2_l := &TreeNode{1, nil, nil}
	c2_r := &TreeNode{2, nil, nil}

	c3_l := &TreeNode{2, nil, nil}
	c3_r := &TreeNode{1, nil, nil}
	root.left = c1_l
	root.right = c1_r

	c1_l.left = c2_l
	c1_l.right = c2_r

	c1_r.left = c3_l
	c1_r.right = c3_r
	ForEach(root, 4, []int{})

	fmt.Println(revertString("abcdefg"))

	fmt.Println(revertString01("ABCDEFG"))

	fmt.Println(mergeSort([]int{2,3,1,4,7,5,6}))
}

func ForEach(node *TreeNode, target int, v[]int)  {
	v = append(v, node.value)
	if (target - node.value) == 0 {
		fmt.Println(v)
	}
	target = target - node.value

	if node.left != nil {
		ForEach(node.left, target, v)
	}
	//可以理解为左和右是分别执行的
	if node.right != nil {
		ForEach(node.right, target, v)
	}
}

func revertString(str string) string {
	l := len(str) / 2
	str_arr := []rune(str)
	for i:=0; i< l ; i++  {
		str_arr[i],str_arr[len(str_arr)-i-1]  = str_arr[len(str)-i-1],str_arr[i]
	}
	return string(str_arr)
}

func revertString01(str string) string {
	if (len(str) == 1) {
		return str
	}
	s := str[len(str)-1:]
	return s + revertString01(string(str[:len(str)-1]))
}

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])
	return merge(left,right)
}
func merge(left, right []int) (result []int)  {
	fmt.Println(left, right)
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if (left[l] > right[r]) {
			result = append(result, right[r])
			r++
		}else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result,left[l:]...)
	result = append(result,right[r:]...)
	return
}

