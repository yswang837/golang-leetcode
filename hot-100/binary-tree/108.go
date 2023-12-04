package binary_tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return traverse(nums, 0, len(nums)-1)
}

func traverse(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	middle := (left + right) / 2
	root := &TreeNode{Val: nums[middle]}
	root.Left = traverse(nums, left, middle-1)
	root.Right = traverse(nums, middle+1, right)
	return root
}
