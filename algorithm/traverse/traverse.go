package traverse

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历：1、迭代遍历；2、递归遍历

// 数组

// 1、迭代
func traverse1(arr []int) {
	for i := 0; i < len(arr); i++ {

	}
}

// 2、递归
func traverse2(arr []int, i int) {
	if i == len(arr) {
		return
	}
	// 前序位置
	traverse2(arr, i+1)
	// 后序位置
}

// 链表
// 1、迭代
func traverse3(head *ListNode) {
	for p := head; p != nil; p = p.Next {

	}
}

// 2、递归
func traverse4(head *ListNode) {
	if head == nil {
		return
	}
	// 前序位置
	traverse4(head.Next)
	// 后序位置
}

// 树
// 1、迭代，二叉树迭代遍历稍难，一般二叉树都是用递归遍历。
func traverse5() {

}

// 2、递归
func traverse6(root *TreeNode) {
	if root == nil {
		return
	}
	// 前序位置
	fmt.Println(root.Value)
	traverse6(root.Left)
	// 中序位置
	traverse6(root.Right)
	// 后序位置
}
