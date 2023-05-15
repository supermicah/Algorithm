package main

import "fmt"

/***
*
* 2023-05-14 21:28:27
*
***/

//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1]
// 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘
//积是18。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
//
//
// 示例 1：
//
// 输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1
//
// 示例 2:
//
// 输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
//
//
//
// 提示：
//
//
// 2 <= n <= 1000
//
//
// 注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/
//
// Related Topics数学 | 动态规划
//
// 👍 254, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	// n = 3a + b
	ret := 1
	// b 等于0，1，2情况最大乘4,所以n > 4
	for n > 4 {
		ret = ret * 3 % 1000000007
		n -= 3
	}
	// 最后结果只会剩下2,3,4 所以直接乘以n再取余1000000007
	return ret * n % 1000000007
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := cuttingRope(120) // 953271190
	fmt.Println(fmt.Sprintf("%+v", value))
}
