package main

import "fmt"

/***
*
* 2023-05-04 16:36:35
*
***/

//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
//
//
//
// 示例 1：
//
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
//
//
//输入：height = [1,1]
//输出：1
//
//
//
//
// 提示：
//
//
// n == height.length
// 2 <= n <= 10⁵
// 0 <= height[i] <= 10⁴
//
//
// Related Topics贪心 | 数组 | 双指针
//
// 👍 4282, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0
	left, right := 0, len(height)-1
	for left < right {
		t := m(height[left], height[right]) * (right - left)
		if max < t {
			max = t
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func m(first, last int) int {
	if first > last {
		return last
	}
	return first
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(fmt.Sprintf("%+v", value))
}
