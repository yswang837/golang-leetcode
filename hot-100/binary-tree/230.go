package binary_tree

func kthSmallest(root *TreeNode, k int) int {
	var f func(root *TreeNode)
	var ss []int
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)
		ss = append(ss, root.Val)
		f(root.Right)
	}
	f(root)
	return ss[k-1]
}
