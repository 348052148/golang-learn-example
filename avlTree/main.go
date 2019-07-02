package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	tree := CreateTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(7)
	fmt.Println(tree.Root.Right.Height)
	tree.ForEach(func(node *Node) {
		fmt.Printf("%d ,", node.Data)
	})

	var R Node
	fmt.Printf("\n B %v \n",R)
	&R = &Node{
		Height:2,
	}
	fmt.Printf("\n B %v \n",R)
	var L *Node
	fmt.Printf("%v \n", L)
	L = new(Node)
	L.Left = new(Node)
	// l-> nil {指向数据的地址}
	// a-> nil {指向数据的地址}
	a := L.Left
	fmt.Printf("p1 %v \n", unsafe.Pointer(L.Left))
	fmt.Printf("p2 %v \n", unsafe.Pointer(a))

	fmt.Printf("&p2 %v \n", unsafe.Pointer(&a))
	a.Height = 2
	//指向数据的地址被替换
	a = &Node{
		Height:2,
	}
	fmt.Printf("%v \n", unsafe.Pointer(&a))
	fmt.Printf(" %v \n", L.Left)

	ac := 1
	fmt.Printf("%p \n", unsafe.Pointer(&ac))
	ac = 2
	fmt.Printf("%p \n", unsafe.Pointer(&ac))

	var d string
	d = "213"
	fmt.Printf("12 %p \n", unsafe.Pointer(&d))
	d = "321"
	fmt.Printf("12 %p \n", unsafe.Pointer(&d))
}

type DataType int

type Node struct {
	Data DataType
	Left, Right *Node
	Height int
}

type Tree struct {
	Root *Node
	Count int
}

func CreateTree() *Tree {
	return &Tree{
		Count:0,
	}
}

func Height(node *Node) int {
	if node == nil {
		return -1
	}else {
		return node.Height
	}
}

func (tree *Tree)Insert(data DataType)  {
	tree.Root = findInsert( data,tree.Root)
	tree.Count++

}

func findInsert(data DataType, Root *Node) *Node {
	fmt.Println("%v", unsafe.Pointer(Root))
	if Root == nil {
		Root = &Node{
			Data:data,
			Height:0,
		}

		fmt.Println(unsafe.Pointer(Root))
	} else {
		if data < Root.Data {
			Root.Left = findInsert(data, Root.Left)
			if Height(Root.Left) - Height(Root.Right) == 2 {
				//LL
				if Root.Left !=nil && data < Root.Left.Data  {
					Root = R(Root)
				} else //LR
				{
					Root = R(Root)
					Root = L(Root)
				}
			}
		}else {
			Root.Right = findInsert(data, Root.Right)
			if Height(Root.Right) - Height(Root.Left) == 2 {
				//RR
				if Root.Right != nil && data < Root.Right.Data {
					Root = L(Root)
				} else //RL
				{
					Root = L(Root)
					Root = R(Root)
				}
			}
		}
	}
	Root.Height = int(math.Max(float64(Height(Root.Left)),float64(Height(Root.Right))) +1)
	return Root
}

func R(node *Node) *Node {
	L := node.Left
	node.Left = L.Right
	L.Right = node
	return L
}
func L(node *Node) *Node {
	R := node.Right
	node.Right = R.Left
	R.Left = node
	return R
}

func (tree *Tree)ForEach(fn func(node *Node))  {
	For(fn, tree.Root)
}

func For(fn func(node *Node), node *Node)  {
	if node != nil {
		For(fn, node.Left)
		fn(node)
		For(fn, node.Right)
	}
}