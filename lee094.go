package leecode

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var rest []int
	if root != nil {
		fmt.Println(root)
		res = append(inorderTraversal(root.Left), root.Val)
		fmt.Println("res:", res)
		rest = append(res, inorderTraversal(root.Right)...)
		fmt.Println("rest:", rest)
	}
	return rest
}

func main() {
	arr := []int{1, nil, 2, 3}
	result := inorderTraversal(arr)
}
