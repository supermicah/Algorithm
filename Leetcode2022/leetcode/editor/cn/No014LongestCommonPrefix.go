package main

import (
	"fmt"
	"strings"
)

/***
*
* 2022-08-20 11:03:03
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
// Related Topics字符串
//
// 👍 2402, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	val := ""
	for i := 1; i <= 200; i++ {
		if len(strs[0]) < i {
			break
		}
		t := strs[0][:i]
		for _, str := range strs {
			if !strings.HasPrefix(str, t) {
				return val
			}
		}
		val = t
	}
	return val
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := longestCommonPrefix([]string{"a"})
	fmt.Println(fmt.Sprintf("%+v", value))
}
