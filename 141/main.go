package main

//快满指针,龟兔赛跑
//slow
//fast
type LinkNode struct {
	Val  int
	Next *LinkNode
}

func hasCycle(head *LinkNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
func main() {

}
