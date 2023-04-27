package main

import "fmt"

/***
*
* 2023-04-27 15:20:40
*
***/

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
// Related Topics哈希表 | 字符串 | 滑动窗口
//
// 👍 9092, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[byte]int)
	ans := 0

	left := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		m[b]++
		for m[b] > 1 {
			m[s[left]]--
			left++
		}
		ans = max(ans, i-left+1)
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := lengthOfLongestSubstring("tmmzuxt")
	fmt.Println(fmt.Sprintf("value: %+v", value))
}
