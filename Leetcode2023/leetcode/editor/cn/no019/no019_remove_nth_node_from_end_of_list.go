package main

import "fmt"

/***
*
* 2023-05-05 11:46:33
*
***/

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
// Related Topics链表 | 双指针
//
// 👍 2518, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	dummy := head
	l := 0
	for dummy != nil {
		l++
		dummy = dummy.Next
	}
	if l < n {
		return head
	}

	loop := head
	c := l - n
	for c > 1 {
		c--
		loop = loop.Next
	}
	loop.Next = loop.Next.Next
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := removeNthFromEnd(toNode([]int{1, 2, 3, 4, 5}), 4)
	fmt.Println(fmt.Sprintf("%+v", value))
}

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
