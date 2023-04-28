package main

import "fmt"

/***
*
* 2023-04-27 19:16:50
*
***/

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
//
// Related Topics数组 | 二分查找 | 分治
//
// 👍 6490, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	mid := l / 2
	s := false
	if mid*2 < l {
		mid++
		s = true
	}

	last := 0

	i := 0
	for len(nums1) > 0 || len(nums2) > 0 {
		if len(nums1) > 0 && len(nums2) > 0 {
			if nums1[0] < nums2[0] {
				last = nums1[0]
				nums1 = nums1[1:]
			} else {
				last = nums2[0]
				nums2 = nums2[1:]
			}
		} else {
			if len(nums1) > 0 {
				last = nums1[0]
				nums1 = nums1[1:]
			} else {
				last = nums2[0]
				nums2 = nums2[1:]
			}
		}
		i++
		if i == mid {
			if s {
				return float64(last)
			} else {
				return (float64(last) + float64(min(nums1, nums2))) / 2
			}
		}
	}
	return 0.0
}

func min(num1, num2 []int) int {
	if len(num1) == 0 {
		return num2[0]
	}
	if len(num2) == 0 {
		return num1[0]
	}
	if num1[0] < num2[0] {
		return num1[0]
	}
	return num2[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := findMedianSortedArrays([]int{1, 2}, []int{3})
	fmt.Println(fmt.Sprintf("%+v", value))
}
