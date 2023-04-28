package main

import "fmt"

/***
*
* 2023-04-28 15:36:51
*
***/

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//
// 例如，121 是回文，而 123 不是。
//
//
//
//
// 示例 1：
//
//
//输入：x = 121
//输出：true
//
//
// 示例 2：
//
//
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3：
//
//
//输入：x = 10
//输出：false
//解释：从右向左读, 为 01 。因此它不是一个回文数。
//
//
//
//
// 提示：
//
//
// -2³¹ <= x <= 2³¹ - 1
//
//
//
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？
//
// Related Topics数学
//
// 👍 2515, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	t := 0
	old := x
	for t < old {
		t = t*10 + x%10
		x = x / 10
	}
	if t == old || old == t/10 {
		return true
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := isPalindrome(121)
	fmt.Println(fmt.Sprintf("%+v", value))
}
