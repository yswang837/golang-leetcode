package link_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{-1, head}
	slow, fast := dummy, dummy
	for i := 0; i <= n; i++ { // 因为加了个虚拟节点dummy，所以需要等于，为什么要加dummy的原因在于：为了删除第一个节点的case
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 执行到这里fast已经是nil了
	slow.Next = slow.Next.Next
	return dummy.Next
}

//              f
//        s
// [a,1,2,3,4,5], n = 2

//      f
//  s
// [a,1], n = 1

//        f
//    s
// [a,1,2], n = 1
