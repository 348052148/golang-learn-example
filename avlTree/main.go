package main

import (
	"fmt"
)

func main() {
	root := Insert(nil, 10)
	Insert(root, 5)
	Insert(root, 20)

	Insert(root, 16)
	Insert(root, 13)
	Insert(root, 19)
	//Insert(root, 17)
	//Insert(root, 18)

	Delete(root, 16)

	ForEach(func(node *Node) {
		fmt.Printf("%d, parent = %v\n", node.Data, node.Parent)
	}, root)
}

type DataType int

type Node struct {
	Data DataType
	Parent *Node
	Left, Right *Node
	Height int
}

func Insert(node *Node, data DataType) *Node {
	if node == nil {
		node = &Node{Data:data}
	}else if node.Data > data {
		//node.data < data insertiong left
		node.Left = Insert(node.Left, data)
		//set Left node parentNode
		node.Left.Parent = node
	}else {
		//node.data > data insertiong right
		node.Right = Insert(node.Right, data)
		//set Right node parentNode
		node.Right.Parent = node
	}
	return node
}

func Delete(node *Node, data DataType)  {
	//search Node.data equals data
	snode := Search(node, data)
	if snode == nil {
		return
	}
	//if node is leaf node
	if snode.Left == nil && snode.Right == nil {
		if snode.Parent.Left != nil && snode.Parent.Left.Data == data {
			snode.Parent.Left = nil
		}else if  snode.Parent.Right != nil {
			snode.Parent.Right = nil
		}
	} else if snode.Right == nil {
		//if del node is parentnode.left
		snode.Left.Parent = snode.Parent
		if snode.Parent.Left.Data == data {
			snode.Parent.Left = snode.Left
		}else {
			snode.Parent.Right = snode.Left
		}
	}else if snode.Left == nil {
		snode.Right.Parent = snode.Parent
		if snode.Parent.Left.Data == data {
			snode.Parent.Left = snode.Right
		}else {
			snode.Parent.Right = snode.Right
		}

	} else {
		//if node have left.node and right.node
		pnode := postNode(snode.Right)

		if snode.Parent.Left.Data == data {
			snode.Parent.Left = pnode
		}else {
			snode.Parent.Right = pnode
		}

		//因为会抑制left fund 下去 只会有右节点有的情况
		if pnode.Right != nil {
			pnode.Parent.Left = pnode.Right
			pnode.Right.Parent = pnode.Parent
		}
		//设置新节点的父节点到老节点的父节点
		pnode.Parent = snode.Parent

		//设置原节点的左和右到新节点
		pnode.Left = snode.Left
		//设置原节点的左和右的父节点
		pnode.Left.Parent = pnode

		//如果替换节点是删除的right节点
		if pnode != snode.Right {
			pnode.Right = snode.Right
			pnode.Right.Parent = pnode
		}
	}

}

//search sub tree just big than current node
func postNode(node *Node) *Node {
	//if node.Right.Left == nil
	if node.Left == nil {
		return node
	}else {
		return postNode(node.Left)
	}
}

//search node.data = data
func Search(node *Node, data DataType) *Node {
	if node == nil {
		return nil
	}
	if node.Data == data {
		return node
	}
	l := Search(node.Left, data)
	r := Search(node.Right, data)
	if l != nil {
		return l
	}else if r != nil {
		return r
	}
	return nil
}

func ForEach(fn func(node *Node), node *Node)  {
	For(fn, node)
}

func For(fn func(node *Node), node *Node)  {
	if node != nil {
		fn(node)
		For(fn, node.Left)
		For(fn, node.Right)
	}
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