package binary_tree

func pathSum(root *TreeNode, targetSum int) int {
	presumCntMap := map[int]int{0: 1}
	var f func(n *TreeNode, curSum int) (ans int)
	f = func(n *TreeNode, curSum int) (ans int) {
		if n == nil {
			return
		}
		curSum += n.Val
		if cnt, ok := presumCntMap[curSum-targetSum]; ok {
			ans += cnt
		}
		presumCntMap[curSum]++
		ans = ans + f(n.Left, curSum) + f(n.Right, curSum)
		presumCntMap[curSum]--
		return
	}
	return f(root, 0)
}
