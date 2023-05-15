package main

import "fmt"

/***
*
* 2023-05-15 09:34:02
*
***/

//输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
//
// 示例 1:
//
// 输入: n = 1
//输出: [1,2,3,4,5,6,7,8,9]
//
//
//
//
// 说明：
//
//
// 用返回一个整数列表来代替打印
// n 为正整数
//
//
// Related Topics数组 | 数学
//
// 👍 294, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func printNumbers(n int) []int {
	max := 1
	for i := 1; i <= n; i++ {
		max *= 10
	}

	ans := make([]int, 0, max-1)
	for i := 1; i < max; i++ {
		ans = append(ans, i)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := printNumbers(1)
	fmt.Println(fmt.Sprintf("%+v", value))
}
