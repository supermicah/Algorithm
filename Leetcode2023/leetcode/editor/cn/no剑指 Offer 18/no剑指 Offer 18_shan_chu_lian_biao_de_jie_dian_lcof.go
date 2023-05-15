package main

import "fmt"

/***
*
* 2023-05-15 10:02:58
*
***/

//给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
// 返回删除后的链表的头节点。
//
// 注意：此题对比原题有改动
//
// 示例 1:
//
// 输入: head = [4,5,1,9], val = 5
//输出: [4,1,9]
//解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//
//
// 示例 2:
//
// 输入: head = [4,5,1,9], val = 1
//输出: [4,5,9]
//解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
//
//
//
//
// 说明：
//
//
// 题目保证链表中节点的值互不相同
// 若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
//
//
// Related Topics链表
//
// 👍 310, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	dummy := head
	for head.Next != nil {
		if head.Next.Val != val {
			head = head.Next
			continue
		}

		if head.Next.Next == nil {
			head.Next = nil
			break
		}
		head.Next.Val = head.Next.Next.Val
		head.Next = head.Next.Next
	}
	return dummy
}

// leetcode submit region end(Prohibit modification and deletion)

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

func main() {
	value := toSlice(deleteNode(toNode([]int{4, 5, 1, 9}), 4))
	fmt.Println(fmt.Sprintf("%+v", value))
}
