package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type record struct {
	res   int
	depth int
}

func maxDepth(root *TreeNode) int {
	r := &record{}
	traverse(root, r)
	return r.res
}

func traverse(root *TreeNode, r *record) {
	if root == nil {
		return
	}
	r.depth++
	if root.Left == nil && root.Right == nil {
		r.res = maxA(r.res, r.depth)
	}
	traverse(root.Left, r)
	traverse(root.Right, r)
	r.depth--
}

func maxA(i, j int) int {
	if i > j {
		return i
	}
	return j
}
