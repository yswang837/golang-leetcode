package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	isLeftToRight := true

	var res [][]int
	if root == nil {
		return res
	}
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		var level []int
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		var finalLevel []int
		if isLeftToRight {
			finalLevel = level
		} else {
			for i := len(level) - 1; i >= 0; i-- {
				finalLevel = append(finalLevel, level[i])
			}
		}
		if len(finalLevel) > 0 {
			res = append(res, finalLevel)
		}
		isLeftToRight = !isLeftToRight
	}
	return res
}

func main() {

}
