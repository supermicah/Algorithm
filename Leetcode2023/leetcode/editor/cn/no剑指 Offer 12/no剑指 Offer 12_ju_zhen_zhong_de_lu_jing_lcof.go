package main

import "fmt"

/***
*
* 2023-05-09 14:22:02
*
***/

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
//
//
//
//
//
// 示例 1：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCCED"
//输出：true
//
//
// 示例 2：
//
//
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
//
//
//
//
// 提示：
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
//
// 注意：本题与主站 79 题相同：https://leetcode-cn.com/problems/word-search/
//
// Related Topics数组 | 回溯 | 矩阵
//
// 👍 794, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
type pair struct{ x, y int }

var s = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	h := len(board)
	w := len(board[0])
	visit := make([][]bool, h)
	for i := range board {
		visit[i] = make([]bool, w)
	}
	var ef func(i, j, k int) bool
	ef = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visit[i][j] = true
		defer func() { visit[i][j] = false }()
		for _, dir := range s {
			newI, newJ := i+dir.x, j+dir.y
			if 0 <= newI && newI < h && 0 <= newJ && newJ < w && !visit[newI][newJ] {
				if ef(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, bytes := range board {
		for j := range bytes {
			if ef(i, j, 0) {
				return true
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	value := exist([][]byte{{'a', 'b'}, {'c', 'd'}}, "abdc")
	fmt.Println(fmt.Sprintf("%+v", value))
}
