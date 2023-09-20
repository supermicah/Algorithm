package main

import "fmt"

/***
*
* 2023-09-20 07:54:18
*
***/

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
// Related Topics递归 | 链表
//
// 👍 3352, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := 1
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
