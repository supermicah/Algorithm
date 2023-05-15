package main

import "fmt"

/***
*
* 2023-05-15 09:33:59
*
***/

//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xⁿ）。不得使用库函数，同时不需要考虑大数问题。
//
//
//
// 示例 1：
//
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//
//
// 示例 2：
//
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//
// 示例 3：
//
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2⁻² = 1/2² = 1/4 = 0.25
//
//
//
// 提示：
//
//
// -100.0 < x < 100.0
// -2³¹ <= n <= 2³¹-1
// -10⁴ <= xⁿ <= 10⁴
//
//
//
//
// 注意：本题与主站 50 题相同：https://leetcode-cn.com/problems/powx-n/
//
// Related Topics递归 | 数学
//
// 👍 415, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	if n%2 == 1 {
		return myPow(x, n-1) * x
	}
	sub := myPow(x, n/2)
	return sub * sub
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := myPow(0.00001, 2147483647)
	fmt.Println(fmt.Sprintf("%+v", value))
}
