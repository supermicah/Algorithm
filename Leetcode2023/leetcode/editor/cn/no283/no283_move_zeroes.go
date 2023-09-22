package main

import "fmt"

/***
*
* 2023-09-22 08:02:20
*
***/

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
// 示例 1:
//
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
//
// 示例 2:
//
//
//输入: nums = [0]
//输出: [0]
//
//
//
// 提示:
//
//
//
// 1 <= nums.length <= 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能尽量减少完成的操作次数吗？
//
// Related Topics数组 | 双指针
//
// 👍 2174, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
		right++
	}

	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

	fmt.Println(fmt.Sprintf("%+v", nums))
}
