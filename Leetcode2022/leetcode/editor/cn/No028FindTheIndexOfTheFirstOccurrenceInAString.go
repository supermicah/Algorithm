package main

import "fmt"

/***
*
* 2023-04-01 15:54:43
*
***/

//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
//如果 needle 不是 haystack 的一部分，则返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：haystack = "sadbutsad", needle = "sad"
//输出：0
//解释："sad" 在下标 0 和 6 处匹配。
//第一个匹配项的下标是 0 ，所以返回 0 。
//
//
// 示例 2：
//
//
//输入：haystack = "leetcode", needle = "leeto"
//输出：-1
//解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
//
//
//
//
// 提示：
//
//
// 1 <= haystack.length, needle.length <= 10⁴
// haystack 和 needle 仅由小写英文字符组成
//
//
// Related Topics双指针 | 字符串 | 字符串匹配
//
// 👍 1815, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	idx := -1
	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[0] {
			continue
		}
		hasMatch := true
		for j := 1; j < len(needle); j++ {
			if i+j >= len(haystack) {
				hasMatch = false
				break
			}
			if haystack[i+j] != needle[j] {
				hasMatch = false
				continue
			}
		}
		if hasMatch {
			idx = i
			break
		}
	}
	return idx
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := 1
	fmt.Println(fmt.Sprintf("%+v", value))
}
