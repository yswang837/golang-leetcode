package binary_tree

// 思路1
// 递归划分子问题的思想，采用了后续遍历，这是动态规划的思想，未借助额外的遍历函数，充分利用函数maxDepth本身返回的值。
// 定义：给我一棵树，我就能返回这棵树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret, depth := 0, 0
	var f func(root1 *TreeNode)
	f = func(root1 *TreeNode) {
		if root1 == nil {
			return
		}
		depth++
		if root1.Left == nil && root1.Right == nil {
			// 到叶节点了，更新最大深度
			ret = maxInt(ret, depth)
		}
		f(root1.Left)
		f(root1.Right)
		depth--
	}
	f(root)
	return ret
}

func maxInt(i, j int) int {
	if i >= j {
		return i
	}
	return j

}
