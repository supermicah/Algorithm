package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func toNode(node []int) *ListNode {
	if node == nil {
		return nil
	}
	n := &ListNode{Val: node[0]}
	if len(node) > 1 {
		n.Next = toNode(node[1:])
	}

	return n
}

func toSlice(node *ListNode) []int {
	s := make([]int, 0)
	for node != nil {
		s = append(s, node.Val)
		node = node.Next
	}
	return s
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
