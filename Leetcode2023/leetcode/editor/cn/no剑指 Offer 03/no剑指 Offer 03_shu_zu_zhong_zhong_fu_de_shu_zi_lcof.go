package main

import (
	"fmt"
)

/***
*
* 2023-05-06 11:06:30
*
***/

//找出数组中重复的数字。
//
// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
//请找出数组中任意一个重复的数字。
//
// 示例 1：
//
// 输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
//
//
//
//
// 限制：
//
// 2 <= n <= 100000
//
// Related Topics数组 | 哈希表 | 排序
//
// 👍 1163, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatNumber(nums []int) int {
	if nums == nil {
		return -1
	}
	// 1. 用last计数
	//sort.Ints(nums)
	//last := -1
	//for _, num := range nums {
	//	if num == last {
	//		return num
	//	}
	//	last = num
	//
	//}
	//return -1

	// 2. 用map计数
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = true
	}
	return -1
	//
	//for _, num := range nums {
	//	if nums[abs(num)] < 0 {
	//		return abs(num)
	//	} else {
	//		nums[abs(num)] *= -1
	//	}
	//}
	//return 0
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := findRepeatNumber([]int{1, 1, 1})
	fmt.Println(fmt.Sprintf("%+v", value))
}
