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

// 写法2
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := f(root.Left)
		right := f(root.Right)
		if left >= right {
			ret = left + 1
		} else {
			ret = right + 1
		}
		return ret
	}
	return f(root)
}

// 思路2
// 递归地遍历一遍树(并未划分子问题)，得到问题答案，这是回溯算法的思想，借助了traverse函数，记录了额外的遍历值，如res和depth，不需要有返回值，给该函数传递指针就行
func maxDepth(root *TreeNode) int {
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

// 思路2-写法2
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
			ret = maxInt(ret, depth)
		}
		f(root1.Left)
		f(root1.Right)
		depth--
	}
	f(root)
	return ret
}
