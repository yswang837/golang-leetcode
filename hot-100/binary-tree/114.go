package binary_tree

// 动态规划，划分子问题，传递一棵树，返回一个链表。只需要考虑当前root节点应该怎么做，递归函数会自动遍历出所有节点。
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	node := root
	for node.Right != nil {
		node = node.Right
	}
	node.Right = right
}
