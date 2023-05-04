package main

import "fmt"

/***
*
* 2023-05-04 23:32:06
*
***/

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics字典树 | 字符串
//
// 👍 2758, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	first := strs[0]
	al := len(first)
	for i := 1; i < l; i++ {
		current := strs[i]
		ml := min(al, len(current))
		ll := 0
		for j := 0; j < ml; j++ {
			if first[j] != current[j] {
				break
			}
			ll++
		}
		al = min(al, ll)
	}
	if al > 0 {
		return first[:al]
	}
	return ""
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(fmt.Sprintf("%+v", value))
}
