package main

import (
	"fmt"
	"math"
)

/***
*
* 2023-11-03 22:02:48
*
***/

//给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。
//
// 整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。
//
// 返回被除数 dividend 除以除数 divisor 得到的 商 。
//
// 注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−2³¹, 231 − 1] 。本题中，如果商 严格大于 231 − 1 ，则返回 2
//31 − 1 ；如果商 严格小于 -2³¹ ，则返回 -2³¹ 。
//
//
//
// 示例 1:
//
//
//输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
//
// 示例 2:
//
//
//输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。
//
//
//
// 提示：
//
//
// -2³¹ <= dividend, divisor <= 2³¹ - 1
// divisor != 0
//
//
// Related Topics位运算 | 数学
//
// 👍 1192, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func divide(dividend int, divisor int) int {
	isMin := false
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
		isMin = true
	}

	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if divisor == 0 {
		return 0
	}
	negative := false
	if dividend < 0 {
		negative = !negative
		dividend = -dividend
	}
	if divisor < 0 {
		negative = !negative
		divisor = -divisor
	}
	candidates := []int{divisor}
	t := divisor
	for t < dividend-t {
		t += t
		candidates = append(candidates, t)
	}

	ans := 0
	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] <= dividend {
			ans |= 1 << i
			dividend -= candidates[i]
		}
	}
	if isMin && divisor == dividend {
		ans += 1
	}
	if negative {
		ans = -ans
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := divide(-2147483648, -2)
	fmt.Println(fmt.Sprintf("%+v", value))
}
