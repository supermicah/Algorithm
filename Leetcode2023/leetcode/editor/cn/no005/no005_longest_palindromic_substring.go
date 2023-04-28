package main

import "fmt"

/***
*
* 2023-04-28 14:19:03
*
***/

//给你一个字符串 s，找到 s 中最长的回文子串。
//
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics字符串 | 动态规划
//
// 👍 6446, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	start, end := 0, 0

	left1, left2, right1, right2 := 0, 0, 0, 0
	for i := 0; i < l; i++ {
		// 1: 中间是单数
		left1, right1 = loop(s, i, i)
		// 2: 中间是双数
		left2, right2 = loop(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}

	}
	return s[start : end+1]
}

func loop(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	return left + 1, right - 1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := longestPalindrome("cbbd")
	fmt.Println(fmt.Sprintf("%+v", value))
}
