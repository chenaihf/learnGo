package mytree

import (
	"ccGo/tree"
	"fmt"
)

type MyTree struct {
	tree.TreeNode
}

func (node *MyTree) Print()  {
	fmt.Println("MyTree Print")
}
