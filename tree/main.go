package main

import "fmt"

func main() {
	tree := CreateEmptyTree()
	tree.InsertNode(&Node{"A",-1})
	tree.InsertNode(&Node{"B",0})
	tree.InsertNode(&Node{"C",0})
	tree.InsertNode(&Node{"D",0})//3
	tree.InsertNode(&Node{"E",1})
	tree.InsertNode(&Node{"F",2})
	tree.InsertNode(&Node{"G",2})

	tree.InsertNode(&Node{"H",3})
	tree.InsertNode(&Node{"I",3})

	tree.foreach(func(node *Node) {
		fmt.Printf("%s , %d |", node.Data, node.Parent)
	})

	tree.RemoveNode(3)
	fmt.Println()
	tree.foreach(func(node *Node) {
		fmt.Printf("%s , %d |", node.Data, node.Parent)
	})
	fmt.Println()
	for _,n := range tree.FindChild(2) {
		fmt.Printf("%s ", n.Data)
	}

	fmt.Println()
	//two
	ttree := CreateTwoTree()
	ttree.Insert(5)
	ttree.Insert(3)
	ttree.Insert(7)
	ttree.Insert(1)
	ttree.Insert(2)
	ttree.Insert(4)
	ttree.Insert(6)
	ttree.Insert(8)

	//fmt.Println(ttree.Root.Left)
	//fmt.Println(ttree.Root.Right)
	ttree.Foreach(func(node *TwoNode) {
		fmt.Printf("%d ",node.Data)
	})

}

type Node struct {
	Data string
	Parent int
}

type Tree struct {
	PNodes []*Node
	Count int
}

func (tree *Tree)InitTree()  {

}
func CreateEmptyTree() *Tree {
	return &Tree{
		Count:0,
	}
}

func (tree *Tree)InsertNode(node *Node)  {
	if node.Parent == -1 || tree.PNodes[node.Parent] != nil {
		tree.PNodes = append(tree.PNodes, node)
		tree.Count++
	}
}

func (tree *Tree)foreach(fn func(node *Node))  {
	for _,node := range tree.PNodes {
		if node != nil {
			fn(node)
		}
	}
}

//删除后节点移入到上一级中
func (tree *Tree)RemoveNode(i int)  {
	//1 判断 i 是否是别人的双亲 不是则直接删除
	if i <= tree.Count {
		pnode := tree.PNodes[i]
		tree.foreach(func(node *Node) {
			if node.Parent == i {
				node.Parent = pnode.Parent
			}
		})
		tree.PNodes[i] = nil
	}
	//2 是则将节点移入
}

func (tree *Tree)FindChild(i int) []*Node {
	var child []*Node
	tree.foreach(func(node *Node) {
		if node.Parent == i {
			child = append(child, node)
		}
	})
	return child
}

//         深度， 父节点， 孩子节点，添加， 删除
//双亲      O(n)  O(1)   O(n)     O(1)  O(1)
//孩子      O(n)  O(n)   O(1)     O(1)  O(1)
//双亲+孩子  O(n)  O(1)   O(1)     O(1)  O(1)
//最左兄弟+双亲
//最左兄弟+双亲+孩子

type TwoNode struct {
	Data int
	Left, Right *TwoNode
}

type TwoTree struct {
	Root *TwoNode
	Count int
}

func CreateTwoTree() *TwoTree {
	return &TwoTree{
		Count:0,
	}
}

func (tree *TwoTree)Foreach(Fn func(node *TwoNode))  {
	Find(tree.Root, Fn)
}

func Find(head *TwoNode, Fn func(node *TwoNode))  {
	if head != nil {
		//Fn(head)
		Find(head.Left, Fn)
		//Fn(head)
		Find(head.Right, Fn)
		Fn(head)
	}
}

func (tree *TwoTree)Insert(v int)  {
	if tree.Root == nil {
		tree.Root = &TwoNode{
			Data:v,
			Left:nil,
			Right:nil,
		}
	} else {
		var pre *TwoNode
		head := tree.Root
		for head != nil {
			pre = head
			if v < head.Data {
				head = head.Left
			}else {
				head = head.Right
			}
		}
		if pre.Data > v {
			pre.Left = &TwoNode{
				Data:v,
				Left:nil,
				Right:nil,
			}
		}else {
			pre.Right = &TwoNode{
				Data:v,
				Left:nil,
				Right:nil,
			}
		}
	}
	tree.Count++
}

