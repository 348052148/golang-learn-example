package main

func main() {
	bt := &BNode{
		Count:2,
		Parent:nil,
		Keys:[M+1]int{0,21,38},
		Branch:[M+1]*BNode{},
	}
}

const M  = 4
const Min int= M/2 -1

type BNode struct {
	Count int
	Parent *BNode
	Keys [M+1]int
	Branch [M+1]*BNode
}

type BTree struct {
	Root *BNode
}

func (t *BNode) Search(value int) (bool,int,*BNode) {
	node := &BNode{}
	var i int
	if t == nil {
		return false, 0, nil
	}
	for t != nil {
		for i = t.Count; i > 0 && value <= t.Keys[i]; i-- {
			if value == t.Keys[i] {
				return true, i, t
			}
		}
		if t.Branch[i] == nil {
			node = t
		}
		t = t.Branch[i]
	}
	return false, i, node
}

func (node *BNode)Insert(value int)  {
	var i int
	ok,_,node := node.Search(value)
	if !ok {
		node.Keys[0] = value
		for i = node.Count;i>0 && value<node.Keys[i];i-- {
			node.Keys[i+1] = node.Keys[i]
		}
		node.Keys[i+1] = value
		node.Count++

		if node.Count < M {
			return
		} else {
			parent := node.Split()
			for parent.Parent != nil {
				parent = parent.Parent
			}
			return

		}

	}
}

func (node *BNode)Split() *BNode {
	newNode := &BNode{}
	parent := node.Parent
	//根节点情况
	if parent == nil {
		parent = &BNode{}
	}
	mid := node.Count/2+1
	newNode.Count = M - mid
	node.Count = mid -1
	j := 1
	k:=mid+1
	for ;k<=M;k++ {  //新生成的右节点
		newNode.Keys[j] = node.Keys[k]
		newNode.Branch[j-1] = node.Branch[k-1]
		j = j+1
	}
	newNode.Branch[j-1] = node.Branch[k-1]
	//新节点和原始节点都指向新的父节点
	newNode.Parent = parent
	node.Parent = parent

	//将该节点中间节点插入到父节点
	k=parent.Count
	for ;node.Keys[mid]<parent.Keys[k];k-- {
		parent.Keys[k+1] = parent.Keys[k]
		parent.Branch[k+1] = parent.Branch[k]
	}
	parent.Keys[k+1] = node.Keys[mid]
	parent.Branch[k] = node
	parent.Branch[k+1] = newNode
	parent.Count++
	if parent.Count >=M {
		return parent.Split()
	}
	return parent
}
