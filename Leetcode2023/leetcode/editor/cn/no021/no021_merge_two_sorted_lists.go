package main

import "fmt"

/***
*
* 2023-05-06 10:56:44
*
***/

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
//
// Related Topics递归 | 链表
//
// 👍 3093, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	list := &ListNode{Val: 0}
	dummy := list
	for list1 != nil && list2 != nil {
		v1 := list1.Val
		v2 := list2.Val
		if v1 < v2 {
			list.Next = &ListNode{Val: v1}
			list1 = list1.Next
		} else {
			list.Next = &ListNode{Val: v2}
			list2 = list2.Next
		}
		list = list.Next
	}
	if list1 != nil {
		list.Next = list1
	}
	if list2 != nil {
		list.Next = list2
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := mergeTwoLists(toNode([]int{1, 2, 4}), toNode([]int{1, 3, 4}))
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
