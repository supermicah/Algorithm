package main

import "fmt"

/***
*
* 2023-11-03 20:53:39
*
***/

//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
//
// Related Topics树 | 广度优先搜索 | 二叉树
//
// 👍 827, 👎 0bug 反馈 | 使用指南 | 更多配套插件
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	res = append(res, []int{root.Val})
	s := make([]*TreeNode, 0)
	if root.Left != nil {
		s = append(s, root.Left)
	}
	if root.Right != nil {
		s = append(s, root.Right)
	}
	zig := true
	for len(s) > 0 {
		values := make([]int, 0)
		ts := make([]*TreeNode, 0)
		for _, node := range s {
			values = append(values, node.Val)
			if node.Left != nil {
				ts = append(ts, node.Left)
			}
			if node.Right != nil {
				ts = append(ts, node.Right)
			}
		}
		if zig {
			values = reverse(values)
		}
		s = ts
		res = append(res, values)
		zig = !zig
	}

	return res
}

func reverse(nums []int) []int {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	value := 1
	fmt.Println(fmt.Sprintf("%+v", value))
}
