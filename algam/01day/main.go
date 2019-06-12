package _1day

import "fmt"

type ListNode struct {
	value int
	next *ListNode
}

func (n *ListNode)Print()  {
	head := n
	for head != nil {
		fmt.Printf(" %d", head.value)
		head = head.next
	}
	fmt.Println()
}

func generate(nums []int) *ListNode  {
	var root, head *ListNode
	for _,num := range nums {
		n := new(ListNode)
		n.value = num
		if root == nil {
			head = n
			root = head
		}else {
			head.next = n
			head = head.next
		}
	}
	return root
}

func main() {
	l1 := generate([]int{1,2,4})
	l2 := generate([]int{1,3,4})
	l1.Print()
	l2.Print()
	mergeTwoLists(l1, l2).Print()

	numbers := [][]int{
		{ 9,0,0,0,0,0 },
		{ 2,6,0,0,0,0 },
		{ 4,5,6,0,0,0 },
		{ 8,2,5,7,0,0 },
		{ 9,3,6,2,4,0 } ,
		{ 2,5,3,7,6,8 },
	}
	fmt.Println(sum(numbers,0,0))
}
//1.遍历处理排序，然后生成新节点
//2.直接链接节点 让较小的节点作为头节点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil  {
		return l2
	}
	if l2 == nil  {
		return l1
	}
	var head *ListNode
	if l1.value <= l2.value {
		head = l1
		fmt.Println("L1", head.value)
		head.next = mergeTwoLists(l1.next, l2)
	} else {
		head = l2
		fmt.Println("L2", head.value)
		head.next = mergeTwoLists(l1, l2.next)
	}
	return head
}

func sum(nums [][]int, r,c int) int {

	val := nums[r][c]
	fmt.Println(val)
	if r<0||r >= 5 ||c<0 ||c >= 5{
		return val
	}

	tmp := 0
	if c-1 < 0 || nums[r+1][c+1] > nums[r+1][c-1] {
		tmp = sum(nums,r+1,c+1)
	} else {
		tmp = sum(nums,r+1,c-1)
	}
	return tmp + val
}
