package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var depth int

func maxDepth(root *TreeNode) int {
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		res = maxA(res, depth)
	}
	traverse(root.Left)
	traverse(root.Right)
	depth--
}

func maxA(i, j int) int {
	if i > j {
		return i
	}
	return j
}
