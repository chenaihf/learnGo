package main

import (
	"ccGo/mytree"
	"ccGo/queue"
	"ccGo/tree"
	"fmt"
)

func main() {
	root := tree.TreeNode{Value: 1}
	root.Print()
	root.Left = &tree.TreeNode{Value: 2}
	root.Right = &tree.TreeNode{Value: 3}
	root.Left.Left = &tree.TreeNode{}
	root.Left.Left.SetValue(4)
	root.Traverse()

	my_root := mytree.MyTree{root}
	my_root.Traverse()
	my_root.Print()

	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	if !(q.Empty()){
		fmt.Println(q.Pop())
	}
	fmt.Println(q.Empty())
}