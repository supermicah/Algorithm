package main

import "fmt"

/***
*
* 2022-08-20 14:43:11
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
// Related Topics递归 | 链表
//
// 👍 2617, 👎 0
//
//
//
//

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dump := &ListNode{}
	head := dump

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return dump.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := 1
	fmt.Println(fmt.Sprintf("%+v", value))
}
