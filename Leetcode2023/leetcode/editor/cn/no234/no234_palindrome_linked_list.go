package main

import (
	"fmt"
)

/***
*
* 2023-09-21 16:17:44
*
***/

//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,2,1]
//输出：true
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics栈 | 递归 | 链表 | 双指针
//
// 👍 1789, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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
func isPalindrome(head *ListNode) bool {
	// 第一种：如果用slice来处理，耗时和内存占用都比较大
	// 	执行耗时:124 ms,击败了50.65% 的Go用户
	//	内存消耗:8.3 MB,击败了73.55% 的Go用户
	//s := make([]int, 0, 10000)
	//for head != nil {
	//	s = append(s, head.Val)
	//	head = head.Next
	//}
	//left, right := 0, len(s)-1
	//for left < right {
	//	if s[left] != s[right] {
	//		return false
	//	}
	//	left++
	//	right--
	//}
	//return true

	// 第二种：思路使用快慢指针找出中间点。这里的中间点有两种情况：
	// 1. 偶数：有两个中间点
	// 2. 奇数：有一个中间点
	// 找到后半段和头这边相同步数，对比
	// 执行耗时:132 ms,击败了35.35% 的Go用户
	// 内存消耗:8.3 MB,击败了69.68% 的Go用户
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next // 重要
	}
	left, right := head, reverse(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	var pre, current *ListNode = nil, node
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := isPalindrome(toNode([]int{1, 2, 2, 1}))
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

func toSlice(node *ListNode) []int {
	s := make([]int, 0)
	for node != nil {
		s = append(s, node.Val)
		node = node.Next
	}
	return s
}
