package main

import "fmt"

/***
*
* 2023-05-11 23:01:57
*
***/

//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
//请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18
//。
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
// 提示：
//
//
// 2 <= n <= 58
//
//
// 注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/
//
// Related Topics数学 | 动态规划
//
// 👍 590, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		current := 0
		for j := 1; j < i; j++ {
			m := max(j*(i-j), dp[j]*(i-j))
			current = max(m, current)
		}
		dp[i] = current
	}
	return dp[n]

	//dp := make([]int, n+1)
	//for i := 2; i <= n; i++ {
	//	current := 0
	//	for j := 1; j <= i; j++ {
	//		m := max((i-j)*j, dp[i-j]*(j))
	//		current = max(current, m)
	//	}
	//	dp[i] = current
	//}
	//return dp[n]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := cuttingRope(10)
	fmt.Println(fmt.Sprintf("%+v", value))
}
