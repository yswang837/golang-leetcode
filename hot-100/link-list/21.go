package link_list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := &ListNode{-1, nil}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

// dummy,cur
// l1: 1->2->4
// l2: 1->3->4
//             cur
// l1: dummy -> 1 -> 2 -> 4
//cur
// l2: dummy -> 1 -> 1 -> 3 -> 4
//cur
// l1: dummy -> 1 -> 1 -> 2 -> 4
//cur
// l2: dummy -> 1 -> 1 -> 2 -> 3 -> 4
//cur
// l1: dummy -> 1 -> 1 -> 2 -> 3 -> 4 // 循环终止
//cur
// l2: dummy -> 1 -> 1 -> 2 -> 3 -> 4 -> 4
