package main

import (
	"fmt"
)

/***
*
* 2023-05-11 15:07:40
*
***/

//地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一
//格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但
//它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
//
//
//
// 示例 1：
//
// 输入：m = 2, n = 3, k = 1
//输出：3
//
//
// 示例 2：
//
// 输入：m = 3, n = 1, k = 0
//输出：1
//
//
// 提示：
//
//
// 1 <= n,m <= 100
// 0 <= k <= 20
//
//
// Related Topics深度优先搜索 | 广度优先搜索 | 动态规划
//
// 👍 633, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
type M struct {
	x int
	y int
}

func movingCount(m int, n int, k int) int {
	queue := make([]M, 0)
	queue = append(queue, M{x: 0, y: 0})
	count := 1
	d := []M{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visit := make(map[int]map[int]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make(map[int]bool, n)
	}
	visit[0][0] = true
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		for _, dd := range d {
			nx := current.x + dd.x
			ny := current.y + dd.y
			if nx >= 0 && nx < m && ny >= 0 && ny < n && cal(nx, ny) <= k {
				if _, ok := visit[nx][ny]; ok {
					continue
				}
				queue = append(queue, M{nx, ny})
				count++
				visit[nx][ny] = true
			}
		}

	}
	return count
}
func cal(i, j int) int {
	total := 0
	for i > 0 {
		total += i % 10
		i = i / 10
	}
	for j > 0 {
		total += j % 10
		j = j / 10
	}
	return total
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := movingCount(16, 8, 4) //15
	//value := movingCount(14, 14, 5) //21
	fmt.Println(fmt.Sprintf("%+v", value))
}
