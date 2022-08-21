package main

import "fmt"

/***
*
* 2022-08-21 20:22:18
*
***/

//实现 strStr() 函数。
//
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如
//果不存在，则返回 -1 。
//
// 说明：
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。
//
//
//
// 示例 1：
//
//
//输入：haystack = "hello", needle = "ll"
//输出：2
//
//
// 示例 2：
//
//
//输入：haystack = "aaaaa", needle = "bba"
//输出：-1
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
// Related Topics双指针 | 字符串 | 字符串匹配
//
// 👍 1559, 👎 0
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
	value := strStr("aaa", "a")
	fmt.Println(fmt.Sprintf("%+v", value))
}
