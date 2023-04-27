package main

import "fmt"

/***
*
* 2023-04-27 14:44:54
*
***/

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
//
// Related Topics递归 | 链表 | 数学
//
// 👍 9534, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	node := &ListNode{}
	head := node
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		node.Val = sum % 10

		if carry == 0 && l1 == nil && l2 == nil {
			return head
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	return head
}

// 性能不好，需要重新写
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	return addTwoNumbers1(0, l1, l2)
//}
//
//func addTwoNumbers1(parent int, l1, l2 *ListNode) *ListNode {
//	if l1 == nil && l2 == nil {
//		if parent == 0 {
//			return nil
//		}
//		return &ListNode{Val: parent}
//	}
//	if l1 == nil {
//		l1 = &ListNode{}
//	}
//	if l2 == nil {
//		l2 = &ListNode{}
//	}
//	s := l1.Val + l2.Val + parent
//	return &ListNode{Val: s % 10, Next: addTwoNumbers1(s/10, l1.Next, l2.Next)}
//}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := addTwoNumbers(toNode([]int{9, 9, 9, 9, 9, 9, 9}), toNode([]int{9, 9, 9, 9}))
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
