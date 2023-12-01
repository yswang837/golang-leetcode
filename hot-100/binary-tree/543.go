package binary_tree

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	maxDepth(root, &maxDiameter)
	return maxDiameter
}

func maxDepth(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left, maxDiameter)
	rightDepth := maxDepth(root.Right, maxDiameter)
	curDepth := leftDepth + rightDepth
	if curDepth > *maxDiameter {
		*maxDiameter = curDepth
	}
	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
