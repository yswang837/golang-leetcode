package binary_tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var level []int
		length := len(queue) // queue的长度会变，需要用变量单独记录，不然level的结果不对
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if len(level) > 0 {
			ret = append(ret, level)
		}
	}
	return ret
}
