package main

import (
	"container/heap"
	"fmt"
)

/***
*
* 2023-09-18 19:54:49
*
***/

//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
// Related Topics链表 | 分治 | 堆（优先队列） | 归并排序
//
// 👍 2640, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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

type Queue []*ListNode

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].Val < q[j].Val
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(*ListNode))
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{Val: -1}
	p := dummy

	pq := make(Queue, 0, len(lists))
	for _, head := range lists {
		if head != nil {
			pq = append(pq, head)
		}
	}
	heap.Init(&pq)

	for pq.Len() != 0 {
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
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
