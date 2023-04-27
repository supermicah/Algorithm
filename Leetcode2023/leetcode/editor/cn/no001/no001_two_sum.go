package main

import "fmt"

/***
*
* 2023-04-27 14:36:48
*
***/

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//
//
// 示例 3：
//
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
// 只会存在一个有效答案
//
//
//
//
// 进阶：你可以想出一个时间复杂度小于 O(n²) 的算法吗？
//
// Related Topics数组 | 哈希表
//
// 👍 16900, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	m := make(map[int]int)
	for i, num := range nums {
		if idx, ok := m[num]; ok {
			if idx != i {
				return []int{i, idx}
			}
		}
		m[target-num] = i
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(fmt.Sprintf("%+v", value))
}
