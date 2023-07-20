package main

import "fmt"

func main() {
	//    //      0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	nums := []int{8, 9, 9, 8, 7, 6, 4, 1, 2, 0}
	buildHeep(nums, len(nums))
	fmt.Println(nums)
}

func buildHeep(nums []int, len int) {
	// 找到最后一个节点的父节点
	parent := len/2 - 1
	for parent >= 0 {
		heapify(nums, parent, len)
		parent--
	}
}

func heapify(nums []int, parent, len int) {
	// 判断两个子节点是否比父节点大，如果是的话替换
	max := parent
	lson := parent*2 + 1
	rson := parent*2 + 2
	if lson < len && nums[lson] > nums[max] {
		// 左节点是否大于父节点
		max = lson
	}
	if rson < len && nums[rson] > nums[max] {
		// 右节点是否大于父节点
		max = rson
	}
	if parent != max {
		swap(&nums[max], &nums[parent])
		heapify(nums, max, len)
	}
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
