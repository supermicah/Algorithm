package main

import "fmt"

/***
*
* 2023-05-05 10:20:29
*
***/

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
//
// Related Topics哈希表 | 字符串 | 回溯
//
// 👍 2445, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	keyMap := make(map[byte][]byte)
	keyMap['2'] = []byte{'a', 'b', 'c'}
	keyMap['3'] = []byte{'d', 'e', 'f'}
	keyMap['4'] = []byte{'g', 'h', 'i'}
	keyMap['5'] = []byte{'j', 'k', 'l'}
	keyMap['6'] = []byte{'m', 'n', 'o'}
	keyMap['7'] = []byte{'p', 'q', 'r', 's'}
	keyMap['8'] = []byte{'t', 'u', 'v'}
	keyMap['9'] = []byte{'w', 'x', 'y', 'z'}

	ans := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		temp := make([]string, 0)
		b := digits[i]
		for _, b2 := range keyMap[b] {
			if len(ans) == 0 {
				temp = append(temp, string(b2))
				continue
			}
			for _, an := range ans {
				temp = append(temp, an+string(b2))
			}
		}
		ans = temp
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := letterCombinations("23")
	fmt.Println(fmt.Sprintf("%+v", value))
}
