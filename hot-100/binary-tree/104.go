package binary_tree

// 思路1
// 递归划分子问题的思想，采用了后续遍历，这是动态规划的思想，未借助额外的遍历函数，充分利用函数maxDepth本身返回的值。
// 定义：给我一棵树，我就能返回这棵树的最大深度
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth1(root.Left)
	right := maxDepth1(root.Right)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}

// 思路2
// 递归地遍历一遍树(并未划分子问题)，得到问题答案，这是回溯算法的思想，借助了traverse函数，记录了额外的遍历值，如res和depth，不需要有返回值，给该函数传递指针就行
func maxDepth2(root *TreeNode) int {
	res, depth := 0, 0
	traverse(root, &res, &depth)
	return res
}

func traverse(root *TreeNode, res, depth *int) {
	if root == nil {
		return
	}
	*depth++
	if root.Left == nil && root.Right == nil {
		*res = maxInt(*res, *depth)
	}
	traverse(root.Left, res, depth)
	traverse(root.Right, res, depth)
	*depth--
}

func maxInt(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
