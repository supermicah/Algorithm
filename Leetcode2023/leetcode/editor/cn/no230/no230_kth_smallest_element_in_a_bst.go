package main

import "fmt"

/***
*
* 2023-10-09 15:22:04
*
***/

//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,1,4,null,2], k = 1
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3
//
//
//
//
//
//
// 提示：
//
//
// 树中的节点数为 n 。
// 1 <= k <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
//
//
//
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
//
// Related Topics树 | 深度优先搜索 | 二叉搜索树 | 二叉树
//
// 👍 782, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	_, res := traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) (idx int, res int) {
	if root == nil {
		return k, 0
	}

	lk, lVal := traverse(root.Left, k)
	if 0 == lk {
		return 0, lVal
	}
	lk--
	if lk == 0 {
		return 0, root.Val
	}
	rk, rVal := traverse(root.Right, lk)
	if rk == 0 {
		return 0, rVal
	}
	return rk, 0
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}
	value := kthSmallest(node, 3)
	fmt.Println(fmt.Sprintf("%+v", value))
}
