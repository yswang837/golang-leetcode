package binary_tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	var f func(node *TreeNode, depth int)
	f = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > len(ret) {
			ret = append(ret, node.Val)
		}
		f(node.Right, depth+1)
		f(node.Left, depth+1)
	}
	f(root, 1)
	return ret
}
