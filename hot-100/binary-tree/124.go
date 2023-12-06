package binary_tree

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := maxInt(maxGain(node.Left), 0)
		rightGain := maxInt(maxGain(node.Right), 0)
		currentSum := node.Val + leftGain + rightGain
		maxSum = maxInt(maxSum, currentSum)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func maxInt(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
