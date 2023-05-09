package main

import "fmt"

/***
*
* 2023-05-07 16:48:21
*
***/

//一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
// 示例 1：
//
// 输入：n = 2
//输出：2
//
//
// 示例 2：
//
// 输入：n = 7
//输出：21
//
//
// 示例 3：
//
// 输入：n = 0
//输出：1
//
// 提示：
//
//
// 0 <= n <= 100
//
//
// 注意：本题与主站 70 题相同：https://leetcode-cn.com/problems/climbing-stairs/
//
//
//
// Related Topics记忆化搜索 | 数学 | 动态规划
//
// 👍 389, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
var m = make(map[int]int)

func numWays(n int) int {
	if n <= 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	}

	way := (numWays(n-1) + numWays(n-2)) % 1000000007
	m[n] = way
	return way
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := numWays(7)
	fmt.Println(fmt.Sprintf("%+v", value))
}
