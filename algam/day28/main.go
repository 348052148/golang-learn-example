package main

import (
	"fmt"
	"container/list"
)

func main() {
	incx := inc(0)
	fmt.Println(incx(1))
	fmt.Println(incx(2))
	fmt.Println(incx(3))
	//
	fmt.Println("-----")
	root := &Node{Val:1}

	left := &Node{Val:2}
	right := &Node{Val:3}

	root.Left = left
	root.Right = right

	lleft := &Node{Val:4}
	lright := &Node{Val:5}

	left.Left = lleft
	left.Right = lright

	rleft := &Node{Val:6}
	rright := &Node{Val:7}

	right.Left = rleft
	right.Right = rright

	DFS(root)
	fmt.Println(BFS01(root))
	fmt.Println(BFS02(root))

	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJumpDp([]int{2,3,1,1,4}))
}
func canJumpDp(nums []int) bool {
	// 2,3,1,1,4
	//2,3
	//f(n)表示n位置是否可达，f(2) = max(f(n-1) + postion+p[n])
	//dp(0) & p[0] + 0 >= i 可达
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j:=0; j < i ; j++ {
			//
			if dp[i-1] && nums[j] + j >= i {
				dp[i] = true
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

func canJump(nums []int) bool {
	return canjumpFinish(0, nums)
}

func canjumpFinish(position int, nums []int) bool {
	if position == len(nums) - 1 {
		return true
	}
	jumpLen := 0
	if position + nums[position] < len(nums) - 1 {
		jumpLen = position + nums[position]
	} else {
		jumpLen = len(nums) - 1
	}
	//【2,3,1,1,4】
	//利用递归回溯
	//1步1步走可走2步
	for i := position + 1; i <= jumpLen; i++ {
		if canjumpFinish(i, nums) {
			return true
		}
	}
	return false
}
func DFS(root *Node)  {
	if root != nil {
		DFS(root.Left)
		fmt.Println(root.Val)
		DFS(root.Right)
	}
}
/*
	1
 2 		3
4  5   6  7

 */
func BFS02(root *Node) []int {
	q := list.List{}
	q.PushFront(root)
	var list []int
	revert := 0
	for q.Len() > 0 {
		node := q.Back()
		q.Remove(node)

		list = append(list, node.Value.(*Node).Val)
		if (revert%2 == 0) {
			if node.Value.(*Node).Left != nil {
				q.PushFront(node.Value.(*Node).Right)
			}
			if node.Value.(*Node).Right != nil {
				q.PushFront(node.Value.(*Node).Left)
			}
		}else {
			if node.Value.(*Node).Right != nil {
				q.PushFront(node.Value.(*Node).Left)
			}
			if node.Value.(*Node).Left != nil {
				q.PushFront(node.Value.(*Node).Right)
			}
		}

		revert++
	}
	return list
}
func BFS01(root *Node) []int {
	stack1 := list.List{}
	stack2 := list.List{}
	stack1.PushFront(root)
	var list []int
	for stack1.Len() > 0 || stack2.Len() > 0 {
		for stack1.Len() > 0 {
			node := stack1.Back()
			stack1.Remove(node)
			if node.Value.(*Node).Left != nil {
				stack2.PushBack(node.Value.(*Node).Left)
			}
			if node.Value.(*Node).Right != nil {
				stack2.PushBack(node.Value.(*Node).Right)
			}

			list = append(list, node.Value.(*Node).Val)
		}

		for stack2.Len() > 0 {
			node := stack2.Back()
			stack2.Remove(node)
			if node.Value.(*Node).Right != nil {
				stack1.PushBack(node.Value.(*Node).Right)
			}
			if node.Value.(*Node).Left != nil {
				stack1.PushBack(node.Value.(*Node).Left)
			}

			list = append(list, node.Value.(*Node).Val)
		}
	}
	return list
}
func BFS(root *Node) []int {
	q := list.List{}
	q.PushFront(root)
	var list []int
	for q.Len() > 0 {
		node := q.Front()
		q.Remove(node)

		list = append(list, node.Value.(*Node).Val)
		if node.Value.(*Node).Left != nil {
			q.PushFront(node.Value.(*Node).Left)
		}
		if node.Value.(*Node).Right != nil {
			q.PushFront(node.Value.(*Node).Right)
		}
	}
	return list
}


func inc(x int) func(int) int {
	return func(y int) int {
		x = x+y
		return x
	}
}
type Node struct {
	Val int
	Left, Right *Node
}

