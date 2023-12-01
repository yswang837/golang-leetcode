package binary_tree

// 思路1：动态规划，充分利用遍历函数invertTree的返回值
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// 思路2：回溯算法，定义无返回值的traverse
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	traverse(root)
	return root
}
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	traverse(root.Left)
	traverse(root.Right)

	return
}
