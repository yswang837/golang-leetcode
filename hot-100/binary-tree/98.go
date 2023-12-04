package binary_tree

import "math"

func dfs(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}
	x := node.Val
	return left < x && x < right && dfs(node.Left, left, x) && dfs(node.Right, x, right)
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}
