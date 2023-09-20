package main

import "fmt"

/***
*
* 2023-09-20 08:01:39
*
***/

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//
// 示例 2：
//
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//
//
// 提示：
//
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
//
// Related Topics链表
//
// 👍 1651, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//if head == nil || head.Next == nil {
	//	return head
	//}

	//dummy := &ListNode{Next: head}
	//p0 := dummy
	//for i := 0; i < left-1; i++ {
	//	p0 = p0.Next
	//}
	//
	//var pre *ListNode
	//current := p0.Next
	//
	//for i := 0; i < right-left+1; i++ {
	//	next := current.Next
	//	current.Next = pre
	//	pre = current
	//	current = next
	//}
	//p0.Next.Next = current
	//p0.Next = pre
	//return dummy.Next

	//if head == nil || head.Next == nil {
	//	return head
	//}
	//var successor *ListNode
	//
	//var reverseN func(head *ListNode, n int) *ListNode
	//reverseN = func(head *ListNode, n int) *ListNode {
	//	if n == 1 {
	//		successor = head.Next
	//		return head
	//	}
	//	last := reverseN(head.Next, n-1)
	//	head.Next.Next = head
	//	head.Next = successor
	//	return last
	//}
	//if left == 1 {
	//	return reverseN(head, right)
	//}
	//head.Next = reverseBetween(head.Next, left-1, right-1)
	//return head

	if head == nil || head.Next == nil {
		return head
	}
	var ln *ListNode
	var reverseN func(head *ListNode, n int) *ListNode
	reverseN = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			ln = head.Next
			return head
		}
		last := reverseN(head.Next, n-1)
		head.Next.Next = head
		head.Next = ln
		return last
	}
	if left == 1 {
		return reverseN(head, left)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := reverseBetween(toNode([]int{1, 2, 3, 4, 5}), 2, 4)
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
