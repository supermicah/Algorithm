package main

import "fmt"

/***
*
* 2023-09-18 15:09:20
*
***/

//统计一个数字在排序数组中出现的次数。
//
//
//
// 示例 1:
//
//
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: 2
//
// 示例 2:
//
//
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: 0
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums 是一个非递减数组
// -10⁹ <= target <= 10⁹
//
//
//
//
// 注意：本题与主站 34 题相同（仅返回值不同）：https://leetcode-cn.com/problems/find-first-and-last-
//position-of-element-in-sorted-array/
//
// Related Topics数组 | 二分查找
//
// 👍 446, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans := 1
			for i := mid + 1; i <= right; i++ {
				if nums[i] == target {
					ans++
				} else {
					break
				}
			}
			for i := mid - 1; i >= left; i-- {
				if nums[i] == target {
					ans++
				} else {
					break
				}
			}
			return ans
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := search([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Println(fmt.Sprintf("%+v", value))
}
