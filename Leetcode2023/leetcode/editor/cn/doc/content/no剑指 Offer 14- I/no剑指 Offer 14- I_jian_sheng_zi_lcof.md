<p>给你一根长度为 <code>n</code> 的绳子，请把绳子剪成整数长度的 <code>m</code> 段（m、n都是整数，n&gt;1并且m&gt;1），每段绳子的长度记为 <code>k[0],k[1]...k[m-1]</code> 。请问 <code>k[0]*k[1]*...*k[m-1]</code> 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入: </strong>2
<strong>输出: </strong>1
<strong>解释: </strong>2 = 1 + 1, 1 × 1 = 1</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入: </strong>10
<strong>输出: </strong>36
<strong>解释: </strong>10 = 3 + 3 + 4, 3 ×&nbsp;3 ×&nbsp;4 = 36</pre>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>2 &lt;= n &lt;= 58</code></li> 
</ul>

<p>注意：本题与主站 343 题相同：<a href="https://leetcode-cn.com/problems/integer-break/">https://leetcode-cn.com/problems/integer-break/</a></p>

<details><summary><strong>Related Topics</strong></summary>数学 | 动态规划</details><br>

<div>👍 590, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题和 [343. 整数拆分](/problems/integer-break) 一样，按照 [动态规划核心套路](https://labuladong.github.io/article/fname.html?fname=动态规划详解进阶) 的流程来就行了。

详细的思路看第 343 题吧，只要改个函数名就能解决这道题。

**标签：[动态规划](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318881141113536512)**

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
    vector<int> memo;

public:
    int cuttingRope(int n) {
        memo = vector<int>(n + 1);
        return dp(n);
    }

    // 定义：拆分 n 获得的最大乘积为 dp(n)
    int dp(int n) {
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return 1;
        }
        if (memo[n] > 0) {
            return memo[n];
        }

        // 状态转移方程
        int res = INT_MIN;
        for (int i = 1; i <= n; i++) {
            res = max(res, i * max(dp(n - i), n - i));
        }
        memo[n] = res;
        return res;
    }

};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def cuttingRope(self, n: int) -> int:
        memo = [0] * (n+1)
        return self.dp(n, memo)

    # 定义：拆分 n 获得的最大乘积为 dp(n)
    def dp(self, n: int, memo: List[int]) -> int:
        if n == 0:
            return 0
        if n == 1:
            return 1
        if memo[n] > 0:
            return memo[n]

        # 状态转移方程
        res = float('-inf')
        for i in range(1, n+1):
            res = max(res, i * max(self.dp(n - i, memo), n - i))

        memo[n] = res
        return res
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    int[] memo;

    public int cuttingRope(int n) {
        memo = new int[n + 1];
        return dp(n);
    }

    // 定义：拆分 n 获得的最大乘积为 dp(n)
    int dp(int n) {
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return 1;
        }
        if (memo[n] > 0) {
            return memo[n];
        }

        // 状态转移方程
        int res = Integer.MIN_VALUE;
        for (int i = 1; i <= n; i++) {
            res = Math.max(res, i * Math.max(dp(n - i), n - i)
            );
        }
        memo[n] = res;
        return res;
    }

}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func cuttingRope(n int) int {
    memo := make([]int, n+1)
    return dp(n, memo)
}

// 定义：拆分 n 获得的最大乘积为 dp(n)
func dp(n int, memo []int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    if memo[n] > 0 {
        return memo[n]
    }

    // 状态转移方程
    res := math.MinInt32
    for i := 1; i <= n; i++ {
        res = max(res, i*max(dp(n-i, memo), n-i))
    }
    memo[n] = res
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var cuttingRope = function(n) {
    var memo = new Array(n + 1).fill(0);
    return dp(n);

    function dp(n) {
        if (n === 0) {
            return 0;
        }
        if (n === 1) {
            return 1;
        }
        if (memo[n] > 0) {
            return memo[n];
        }

        // 状态转移方程
        var res = Number.MIN_SAFE_INTEGER;
        for (var i = 1; i <= n; i++) {
            res = Math.max(res, i * Math.max(dp(n - i), n - i));
        }
        memo[n] = res;
        return res;
    }
};
```

</div></div>
</div></div>

</details>
</div>



