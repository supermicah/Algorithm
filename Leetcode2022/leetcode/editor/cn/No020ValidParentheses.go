package main

import "fmt"

/***
*
* 2022-08-20 14:36:03
*
***/

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
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
// 示例 4：
//
//
//输入：s = "([)]"
//输出：false
//
//
// 示例 5：
//
//
//输入：s = "{[]}"
//输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
// Related Topics栈 | 字符串
//
// 👍 3460, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	m := make(map[byte]byte)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['
	var stack []byte
	for _, val := range s {
		v, ok := m[byte(val)]
		if !ok {
			stack = append(stack, byte(val))
			continue
		}
		if len(stack) == 0 {
			return false
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != v {
			return false
		}
	}
	return len(stack) == 0

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := isValid("()[()]{}")
	fmt.Println(fmt.Sprintf("%+v", value))
}
