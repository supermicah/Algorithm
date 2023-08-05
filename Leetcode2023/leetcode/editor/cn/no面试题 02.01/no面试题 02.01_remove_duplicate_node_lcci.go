package main

import "fmt"

/***
*
* 2023-08-05 15:49:57
*
***/

//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
// 示例1:
//
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//
//
// 示例2:
//
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//
//
// 提示：
//
//
// 链表长度在[0, 20000]范围内。
// 链表元素在[0, 20000]范围内。
//
//
// 进阶：
//
// 如果不得使用临时缓冲区，该怎么解决？
//
// Related Topics哈希表 | 链表 | 双指针
//
// 👍 189, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := head
	m := make(map[int]struct{})
	m[dummy.Val] = struct{}{}
	for dummy.Next != nil {
		if _, ok := m[dummy.Next.Val]; ok {
			dummy.Next = dummy.Next.Next
		} else {
			m[dummy.Next.Val] = struct{}{}
			dummy = dummy.Next
		}

	}
	return head

	//if head == nil {
	//	return head
	//}
	//occurred := map[int]bool{head.Val: true}
	//pos := head
	//for pos.Next != nil {
	//	cur := pos.Next
	//	if !occurred[cur.Val] {
	//		occurred[cur.Val] = true
	//		pos = pos.Next
	//	} else {
	//		pos.Next = pos.Next.Next
	//	}
	//}
	////pos.Next = nil
	//return head
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
