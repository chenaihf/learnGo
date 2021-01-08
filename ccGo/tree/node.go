package tree

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func (node *TreeNode) Print()  {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(v int){
	node.Value = v
}

func (node *TreeNode) Traverse()  {
	//if node == nil{
	//	return
	//}

	fmt.Println(node.Value)
	if node.Left != nil{
		node.Left.Traverse()
	}
	if node.Right != nil{
		node.Right.Traverse()
	}

}

//func main() {
//	root := treeNode{value: 1}
//	root.print()
//	root.left = &treeNode{value: 2}
//	root.right = &treeNode{value: 3}
//	root.left.left = &treeNode{}
//	root.left.left.setValue(4)
//	root.traverse()
//
//}


