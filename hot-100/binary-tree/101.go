package binary_tree

// 思路1
func isSymmetric(root *TreeNode) bool {
	return traverse(root, root)
}

func traverse(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && traverse(root1.Left, root2.Right) && traverse(root1.Right, root2.Left)
}

// 思路2
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{} // 借助一个队列
	p, q := root, root     // 初始化两个节点，并将其放入队列
	queue = append(append(queue, p), q)
	for len(queue) > 0 {
		p, q = queue[0], queue[1]
		queue = queue[2:] // 出队两个元素
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(append(queue, p.Left), q.Right)
		queue = append(append(queue, p.Right), q.Left)
	}
	return true
}
