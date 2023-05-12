package main

import (
	"fmt"
	"strings"
)

/***
*
* 2023-05-06 12:25:03
*
***/

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
// 示例 1：
//
// 输入：s = "We are happy."
//输出："We%20are%20happy."
//
//
//
// 限制：
//
// 0 <= s 的长度 <= 10000
//
// Related Topics字符串
//
// 👍 474, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	var ans strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans.WriteString("%20")
		} else {
			ans.WriteByte(s[i])
		}
	}
	return ans.String()
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := replaceSpace("We are happy.")
	fmt.Println(fmt.Sprintf("%+v", value))
}
