package main

import (
	"fmt"
	"strconv"
)

//给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
//
// 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//
//
// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
//
//
// 给定 n 和 k，返回第 k 个排列。
//
//
//
// 示例 1：
//
//
//输入：n = 3, k = 3
//输出："213"
//
//
// 示例 2：
//
//
//输入：n = 4, k = 9
//输出："2314"
//
//
// 示例 3：
//
//
//输入：n = 3, k = 1
//输出："123"
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
// 1 <= k <= n!
//
// Related Topics 递归 数学 👍 681 👎 0

func main() {
	no60Print("%+v", getPermutation(4, 9))
}

func no60Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--

	ans := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k/factorial[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n-i]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
