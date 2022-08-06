package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func detectCycle(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	third := head
	for third != slow {
		third = third.Next
		slow = slow.Next
	}
	return slow
}

func main() {

}
