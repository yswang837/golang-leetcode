package link_list

func isPalindrome(head *ListNode) bool {
	var ret []int
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	left, right := 0, len(ret)-1
	for left < right {
		if ret[left] != ret[right] {
			return false
		}
		left++
		right--
	}
	return true
}
