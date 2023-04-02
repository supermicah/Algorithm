package main

import "fmt"

/***
*
* 2022-08-23 00:28:46
*
***/

//一个 n x n 的二维网络 board 仅由 0 和 1 组成 。每次移动，你能任意交换两列或是两行的位置。
//
// 返回 将这个矩阵变为 “棋盘” 所需的最小移动次数 。如果不存在可行的变换，输出 -1。
//
// “棋盘” 是指任意一格的上下左右四个方向的值均与本身不同的矩阵。
//
//
//
// 示例 1:
//
//
//
//
//输入: board = [[0,1,1,0],[0,1,1,0],[1,0,0,1],[1,0,0,1]]
//输出: 2
//解释:一种可行的变换方式如下，从左到右：
//第一次移动交换了第一列和第二列。
//第二次移动交换了第二行和第三行。
//
//
// 示例 2:
//
//
//
//
//输入: board = [[0, 1], [1, 0]]
//输出: 0
//解释: 注意左上角的格值为0时也是合法的棋盘，也是合法的棋盘.
//
//
// 示例 3:
//
//
//
//
//输入: board = [[1, 0], [1, 0]]
//输出: -1
//解释: 任意的变换都不能使这个输入变为合法的棋盘。
//
//
//
//
// 提示：
//
//
// n == board.length
// n == board[i].length
// 2 <= n <= 30
// board[i][j] 将只包含 0或 1
//
// Related Topics位运算 | 数组 | 数学 | 矩阵
//
// 👍 56, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func movesToChessboard(board [][]int) int {
	w := len(board)
	for i := 0; i < w; i++ {
		if i == 0 {
			continue
		}
		match := true
		if board[i-1][0] == board[i][0] {
			for j := i + 1; j < w; j++ {
				if board[i-1][0] != board[j][0] {
					board[i], board[j] = board[j], board[i]
					match = true
					break
				}
			}
		} else {
			match = true
		}
		if !match {
			return -1
		}
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := movesToChessboard([][]int{{1, 0, 1}, {1, 0, 1}, {0, 1, 0}})
	fmt.Println(fmt.Sprintf("%+v", value))
}
