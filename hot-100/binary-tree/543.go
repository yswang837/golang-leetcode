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

// 写法2
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		curDepth := leftDepth + rightDepth
		if curDepth > maxDiameter {
			maxDiameter = curDepth
		}
		if leftDepth >= rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
	maxDepth(root)
	return maxDiameter
}
