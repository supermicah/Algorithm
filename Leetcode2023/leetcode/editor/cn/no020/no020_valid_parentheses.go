package main

import "fmt"

/***
*
* 2023-05-06 10:45:25
*
***/

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
//
// Related Topics栈 | 字符串
//
// 👍 3904, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	n := len(s)
	//if n%2 == 1 {
	//	return false
	//}
	m := make(map[byte]byte)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['
	q := make([]byte, 0)
	for i := 0; i < n; i++ {
		b := s[i]
		if v, ok := m[b]; ok {
			if len(q) > 0 && q[len(q)-1] == v {
				q = q[:len(q)-1]
			} else {
				return false
			}
		} else {
			q = append(q, b)
		}
	}
	if len(q) > 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := isValid("()[]{}")
	fmt.Println(fmt.Sprintf("%+v", value))
}
