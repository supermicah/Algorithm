<p>给你一个由&nbsp;<code>'1'</code>（陆地）和 <code>'0'</code>（水）组成的的二维网格，请你计算网格中岛屿的数量。</p>

<p>岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。</p>

<p>此外，你可以假设该网格的四条边均被水包围。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
<strong>输出：</strong>1
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
<strong>输出：</strong>3
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>m == grid.length</code></li> 
 <li><code>n == grid[i].length</code></li> 
 <li><code>1 &lt;= m, n &lt;= 300</code></li> 
 <li><code>grid[i][j]</code> 的值为 <code>'0'</code> 或 <code>'1'</code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>深度优先搜索 | 广度优先搜索 | 并查集 | 数组 | 矩阵</details><br>

<div>👍 2236, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员！[第 21 期打卡挑战](https://opedk.xet.tech/s/4ptSo2) 最后一天报名！**



<p><strong><a href="https://labuladong.github.io/article/slug.html?slug=number-of-islands" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

岛屿系列问题可以用 DFS/BFS 算法或者 [Union-Find 并查集算法](https://labuladong.github.io/article/fname.html?fname=UnionFind算法详解) 来解决。

用 DFS 算法解决岛屿题目是最常见的，每次遇到一个岛屿中的陆地，就用 DFS 算法吧这个岛屿「淹掉」。

如何使用 DFS 算法遍历二维数组？你把二维数组中的每个格子看做「图」中的一个节点，这个节点和周围的四个节点连通，这样二维矩阵就被抽象成了一幅网状的「图」。

为什么每次遇到岛屿，都要用 DFS 算法把岛屿「淹了」呢？主要是为了省事，避免维护 `visited` 数组。

[图算法遍历基础](https://labuladong.github.io/article/fname.html?fname=图) 说了，遍历图是需要 `visited` 数组记录遍历过的节点防止走回头路。

因为 `dfs` 函数遍历到值为 `0` 的位置会直接返回，所以只要把经过的位置都设置为 `0`，就可以起到不走回头路的作用。

**详细题解：[一文秒杀所有岛屿题目](https://labuladong.github.io/article/fname.html?fname=岛屿题目)**

**标签：[DFS 算法](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122002916411604996)，二维矩阵**

## 解法代码

提示：🟢 标记的是我写的解法代码，🤖 标记的是 chatGPT 翻译的多语言解法代码。如有错误，可以 [点这里](https://github.com/labuladong/fucking-algorithm/issues/1113) 反馈和修正。

<div class="tab-panel"><div class="tab-nav">
<button data-tab-item="cpp" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">cpp🤖</button>

<button data-tab-item="python" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">python🤖</button>

<button data-tab-item="java" class="tab-nav-button btn active" data-tab-group="default" onclick="switchTab(this)">java🟢</button>

<button data-tab-item="go" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">go🤖</button>

<button data-tab-item="javascript" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">javascript🤖</button>
</div><div class="tab-content">
<div data-tab-item="cpp" class="tab-item " data-tab-group="default"><div class="highlight">

```cpp
// 注意：cpp 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution {
    // 主函数，计算岛屿数量
public:
    int numIslands(vector<vector<char>>& grid) {
        int res = 0;
        int m = grid.size(), n = grid[0].size();
        // 遍历 grid
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    // 每发现一个岛屿，岛屿数量加一
                    res++;
                    // 然后使用 DFS 将岛屿淹了
                    dfs(grid, i, j);
                }
            }
        }
        return res;
    }

private:
    // 从 (i, j) 开始，将与之相邻的陆地都变成海水
    void dfs(vector<vector<char>>& grid, int i, int j) {
        int m = grid.size(), n = grid[0].size();
        if (i < 0 || j < 0 || i >= m || j >= n) {
            // 超出索引边界
            return;
        }
        if (grid[i][j] == '0') {
            // 已经是海水了
            return;
        }
        // 将 (i, j) 变成海水
        grid[i][j] = '0';
        // 淹没上下左右的陆地
        dfs(grid, i + 1, j);
        dfs(grid, i, j + 1);
        dfs(grid, i - 1, j);
        dfs(grid, i, j - 1);
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        res = 0
        m = len(grid)
        n = len(grid[0])
        # 遍历 grid
        for i in range(m):
            for j in range(n):
                if grid[i][j] == '1':
                    # 每发现一个岛屿，岛屿数量加一
                    res += 1
                    # 然后使用 DFS 将岛屿淹了
                    self.dfs(grid, i, j)
        return res

    # 从 (i, j) 开始，将与之相邻的陆地都变成海水
    def dfs(self, grid: List[List[str]], i: int, j: int) -> None:
        m = len(grid)
        n = len(grid[0])
        if i < 0 or j < 0 or i >= m or j >= n:
            # 超出索引边界
            return
        if grid[i][j] == '0':
            # 已经是海水了
            return
        # 将 (i, j) 变成海水
        grid[i][j] = '0'
        # 淹没上下左右的陆地
        self.dfs(grid, i + 1, j)
        self.dfs(grid, i, j + 1)
        self.dfs(grid, i - 1, j)
        self.dfs(grid, i, j - 1)
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    // 主函数，计算岛屿数量
    public int numIslands(char[][] grid) {
        int res = 0;
        int m = grid.length, n = grid[0].length;
        // 遍历 grid
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    // 每发现一个岛屿，岛屿数量加一
                    res++;
                    // 然后使用 DFS 将岛屿淹了
                    dfs(grid, i, j);
                }
            }
        }
        return res;
    }

    // 从 (i, j) 开始，将与之相邻的陆地都变成海水
    void dfs(char[][] grid, int i, int j) {
        int m = grid.length, n = grid[0].length;
        if (i < 0 || j < 0 || i >= m || j >= n) {
            // 超出索引边界
            return;
        }
        if (grid[i][j] == '0') {
            // 已经是海水了
            return;
        }
        // 将 (i, j) 变成海水
        grid[i][j] = '0';
        // 淹没上下左右的陆地
        dfs(grid, i + 1, j);
        dfs(grid, i, j + 1);
        dfs(grid, i - 1, j);
        dfs(grid, i, j - 1);
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func numIslands(grid [][]byte) int {
    res := 0
    m, n := len(grid), len(grid[0])
    // 遍历 grid
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                // 每发现一个岛屿，岛屿数量加一
                res++
                // 然后使用 DFS 将岛屿淹了
                dfs(grid, i, j)
            }
        }
    }
    return res
}

// 从 (i, j) 开始，将与之相邻的陆地都变成海水
func dfs(grid [][]byte, i, j int) {
    m, n := len(grid), len(grid[0])
    if i < 0 || j < 0 || i >= m || j >= n {
        // 超出索引边界
        return
    }
    if grid[i][j] == '0' {
        // 已经是海水了
        return
    }
    // 将 (i, j) 变成海水
    grid[i][j] = '0'
    // 淹没上下左右的陆地
    dfs(grid, i+1, j)
    dfs(grid, i, j+1)
    dfs(grid, i-1, j)
    dfs(grid, i, j-1)
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var numIslands = function(grid) {
    var res = 0;
    var m = grid.length, n = grid[0].length;
    // 遍历 grid
    for (var i = 0; i < m; i++) {
        for (var j = 0; j < n; j++) {
            if (grid[i][j] == '1') {
                // 每发现一个岛屿，岛屿数量加一
                res++;
                // 然后使用 DFS 将岛屿淹了
                dfs(grid, i, j);
            }
        }
    }
    return res;
};

// 从 (i, j) 开始，将与之相邻的陆地都变成海水
function dfs(grid, i, j) {
    var m = grid.length, n = grid[0].length;
    if (i < 0 || j < 0 || i >= m || j >= n) {
        // 超出索引边界
        return;
    }
    if (grid[i][j] == '0') {
        // 已经是海水了
        return;
    }
    // 将 (i, j) 变成海水
    grid[i][j] = '0';
    // 淹没上下左右的陆地
    dfs(grid, i + 1, j);
    dfs(grid, i, j + 1);
    dfs(grid, i - 1, j);
    dfs(grid, i, j - 1);
}
```

</div></div>
</div></div>

<details open><summary><strong>👾👾 算法可视化 👾👾</strong></summary><div id="data_number-of-islands" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_number-of-islands"></div></div>
</details><hr /><br />

**类似题目**：
  - [1020. 飞地的数量 🟠](/problems/number-of-enclaves)
  - [1254. 统计封闭岛屿的数目 🟠](/problems/number-of-closed-islands)
  - [1905. 统计子岛屿 🟠](/problems/count-sub-islands)
  - [694. 不同岛屿的数量 🟠](/problems/number-of-distinct-islands)
  - [695. 岛屿的最大面积 🟠](/problems/max-area-of-island)
  - [剑指 Offer II 105. 岛屿的最大面积 🟠](/problems/ZL6zAn)

</details>
</div>

