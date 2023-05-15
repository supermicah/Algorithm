<p>请实现一个函数用来匹配包含<code>'. '</code>和<code>'*'</code>的正则表达式。模式中的字符<code>'.'</code>表示任意一个字符，而<code>'*'</code>表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串<code>"aaa"</code>与模式<code>"a.a"</code>和<code>"ab*ac*a"</code>匹配，但与<code>"aa.a"</code>和<code>"ab*a"</code>均不匹配。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong>
s = "aa"
p = "a"
<strong>输出:</strong> false
<strong>解释:</strong> "a" 无法匹配 "aa" 整个字符串。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong>
s = "aa"
p = "a*"
<strong>输出:</strong> true
<strong>解释:</strong>&nbsp;因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong>
s = "ab"
p = ".*"
<strong>输出:</strong> true
<strong>解释:</strong>&nbsp;".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
</pre>

<p><strong>示例 4:</strong></p>

<pre><strong>输入:</strong>
s = "aab"
p = "c*a*b"
<strong>输出:</strong> true
<strong>解释:</strong>&nbsp;因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
</pre>

<p><strong>示例 5:</strong></p>

<pre><strong>输入:</strong>
s = "mississippi"
p = "mis*is*p*."
<strong>输出:</strong> false</pre>

<ul> 
 <li><code>s</code>&nbsp;可能为空，且只包含从&nbsp;<code>a-z</code>&nbsp;的小写字母。</li> 
 <li><code>p</code>&nbsp;可能为空，且只包含从&nbsp;<code>a-z</code>&nbsp;的小写字母以及字符&nbsp;<code>.</code>&nbsp;和&nbsp;<code>*</code>，无连续的 <code>'*'</code>。</li> 
</ul>

<p>注意：本题与主站 10&nbsp;题相同：<a href="https://leetcode-cn.com/problems/regular-expression-matching/">https://leetcode-cn.com/problems/regular-expression-matching/</a></p>

<details><summary><strong>Related Topics</strong></summary>递归 | 字符串 | 动态规划</details><br>

<div>👍 536, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**



<p><strong><a href="https://labuladong.gitee.io/article/slug.html?slug=zheng-ze-biao-da-shi-pi-pei-lcof" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题和 [10. 正则表达式匹配](/problems/regular-expression-matching) 相同。

PS：这道题在[《算法小抄》](https://item.jd.com/12759911.html) 的第 155 页。

`s` 和 `p` 相互匹配的过程大致是，两个指针 `i, j` 分别在 `s` 和 `p` 上移动，如果最后两个指针都能移动到字符串的末尾，那么就匹配成功，反之则匹配失败。

正则表达算法问题只需要把住一个基本点：看 `s[i]` 和 `p[j]` 两个字符是否匹配，一切逻辑围绕匹配/不匹配两种情况展开即可。

动态规划算法的核心就是「状态」和「选择」，**「状态」无非就是 `i` 和 `j` 两个指针的位置，「选择」就是模式串的 `p[j]` 选择匹配几个字符**。

`dp` 函数的定义如下：

**若 `dp(s, i, p, j) = true`，则表示 `s[i..]` 可以匹配 `p[j..]`；若 `dp(s, i, p, j) = false`，则表示 `s[i..]` 无法匹配 `p[j..]`**。

**详细题解：[经典动态规划：正则表达式](https://labuladong.github.io/article/fname.html?fname=动态规划之正则表达)**

**标签：[二维动态规划](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122017695998050308)，[动态规划](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318881141113536512)，字符串**

## 解法代码

提示：🟢 标记的是我写的解法代码，🤖 标记的是 chatGPT 翻译的多语言解法代码。如有错误，可以 [点这里](https://github.com/labuladong/fucking-algorithm/issues/1113) 反馈和修正。

<div class="tab-panel"><div class="tab-nav">
<button data-tab-item="cpp" class="tab-nav-button btn active" data-tab-group="default" onclick="switchTab(this)">cpp🟢</button>

<button data-tab-item="python" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">python🤖</button>

<button data-tab-item="java" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">java🤖</button>

<button data-tab-item="go" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">go🤖</button>

<button data-tab-item="javascript" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">javascript🤖</button>
</div><div class="tab-content">
<div data-tab-item="cpp" class="tab-item active" data-tab-group="default"><div class="highlight">

```cpp
class Solution {
    public:
    // 备忘录
    vector<vector<int>> memo;

    bool isMatch(string s, string p) {
        int m = s.size(), n = p.size();
        memo = vector<vector<int>>(m, vector<int>(n, -1));
        // 指针 i，j 从索引 0 开始移动
        return dp(s, 0, p, 0);
    }

    /* 计算 p[j..] 是否匹配 s[i..] */
    bool dp(string& s, int i, string& p, int j) {
        int m = s.size(), n = p.size();
        // base case
        if (j == n) {
            return i == m;
        }
        if (i == m) {
            if ((n - j) % 2 == 1) {
                return false;
            }
            for (; j + 1 < n; j += 2) {
                if (p[j + 1] != '*') {
                    return false;
                }
            }
            return true;
        }

        // 查备忘录，防止重复计算
        if (memo[i][j] != -1) {
            return memo[i][j];
        }

        bool res = false;

        if (s[i] == p[j] || p[j] == '.') {
            if (j < n - 1 && p[j + 1] == '*') {
                res = dp(s, i, p, j + 2)
                        || dp(s, i + 1, p, j);
            } else {
                res = dp(s, i + 1, p, j + 1);
            }
        } else {
            if (j < n - 1 && p[j + 1] == '*') {
                res = dp(s, i, p, j + 2);
            } else {
                res = false;
            }
        }
        // 将当前结果记入备忘录
        memo[i][j] = res;
        return res;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        # 备忘录
        memo = []
        # 创建大小为m*n的备忘录，初始值为-1
        m = len(s)
        n = len(p)
        for i in range(m):
            memo.append([-1] * n)
        # 索引指针 i,j 从 0 开始移动
        return self.dp(s, 0, p, 0, memo)

    # 计算 p[j..] 是否匹配 s[i..]
    def dp(self, s, i, p, j, memo):
        m = len(s)
        n = len(p)
        # 如果 j 达到结尾，判断 i 是否也到结尾
        if j == n:
            return i == m
        # 如果 i 达到结尾，判断 j 接下来的字符能否匹配空串
        if i == m:
            if (n - j) % 2 == 1:
                return False
            for k in range(j + 1, n, 2):
                if p[k] != '*':
                    return False
            return True

        # 查备忘录，防止重复计算
        if memo[i][j] != -1:
            return memo[i][j]

        res = False
        if s[i] == p[j] or p[j] == '.':
            if j < n - 1 and p[j + 1] == '*':
                res = self.dp(s, i, p, j + 2, memo) or self.dp(s, i + 1, p, j, memo)
            else:
                res = self.dp(s, i + 1, p, j + 1, memo)
        else:
            if j < n - 1 and p[j + 1] == '*':
                res = self.dp(s, i, p, j + 2, memo)
            else:
                res = False

        # 记入当前结果到备忘录
        memo[i][j] = res
        return res
```

</div></div>

<div data-tab-item="java" class="tab-item " data-tab-group="default"><div class="highlight">

```java
// 注意：java 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution {
    // 备忘录
    private int[][] memo;

    public boolean isMatch(String s, String p) {
        int m = s.length(), n = p.length();
        memo = new int[m][n];
        for (int i = 0; i < m; i++) {
            Arrays.fill(memo[i], -1);
        }
        // 指针 i，j 从索引 0 开始移动
        return dp(s, 0, p, 0);
    }

    /* 计算 p[j..] 是否匹配 s[i..] */
    private boolean dp(String s, int i, String p, int j) {
        int m = s.length(), n = p.length();
        // base case
        if (j == n) {
            return i == m;
        }
        if (i == m) {
            if ((n - j) % 2 == 1) {
                return false;
            }
            for (; j + 1 < n; j += 2) {
                if (p.charAt(j + 1) != '*') {
                    return false;
                }
            }
            return true;
        }

        // 查备忘录，防止重复计算
        if (memo[i][j] != -1) {
            return memo[i][j] == 1;
        }

        boolean res = false;

        if (s.charAt(i) == p.charAt(j) || p.charAt(j) == '.') {
            if (j < n - 1 && p.charAt(j + 1) == '*') {
                res = dp(s, i, p, j + 2)
                        || dp(s, i + 1, p, j);
            } else {
                res = dp(s, i + 1, p, j + 1);
            }
        } else {
            if (j < n - 1 && p.charAt(j + 1) == '*') {
                res = dp(s, i, p, j + 2);
            } else {
                res = false;
            }
        }
        // 将当前结果记入备忘录
        memo[i][j] = res ? 1 : 0;
        return res;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func isMatch(s string, p string) bool {
    // 备忘录
    memo := make([][]int, len(s))
    for i := range memo {
        memo[i] = make([]int, len(p))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    // 指针 i，j 从索引 0 开始移动
    return dp(s, 0, p, 0, memo)
}

/* 计算 p[j..] 是否匹配 s[i..] */
func dp(s string, i int, p string, j int, memo [][]int) bool {
    m, n := len(s), len(p)
    // base case
    if j == n {
        return i == m
    }
    if i == m {
        if (n-j)%2 == 1 {
            return false
        }
        for ; j+1 < n; j += 2 {
            if p[j+1] != '*' {
                return false
            }
        }
        return true
    }
    // 查备忘录，防止重复计算
    if memo[i][j] != -1 {
        return memo[i][j] == 1
    }

    var res bool

    if s[i] == p[j] || p[j] == '.' {
        if j < n-1 && p[j+1] == '*' {
            res = dp(s, i, p, j+2, memo) || dp(s, i+1, p, j, memo)
        } else {
            res = dp(s, i+1, p, j+1, memo)
        }
    } else {
        if j < n-1 && p[j+1] == '*' {
            res = dp(s, i, p, j+2, memo)
        }
    }
    // 将当前结果记入备忘录
    if res {
        memo[i][j] = 1
    } else {
        memo[i][j] = 0
    }
    return res
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 cpp 代码对比查看。

var isMatch = function(s, p) {
    var m = s.length,
        n = p.length;
    var memo = new Array(m + 1); // 先构造一个数组
    for (var i = 0; i <= m; i++) {
        memo[i] = new Array(n + 1).fill(false);
    }
    memo[0][0] = true; // 两个字符串为空串时，互相匹配

    for (i = 0; i <= m; i++) {
        for (var j = 1; j <= n; j++) {
            if (p.charAt(j - 1) === '*') {
                // * 号能匹配 0 次或更多次
                memo[i][j] = memo[i][j - 2] // 匹配 0 次
                || (i > 0 && (s.charAt(i - 1) === p.charAt(j - 2) || p.charAt(j - 2) === '.') && memo[i - 1][j]);
            } else {
                if (i > 0 && (s.charAt(i - 1) === p.charAt(j - 1) || p.charAt(j - 1) === '.')) {
                    memo[i][j] = memo[i - 1][j - 1];
                }
            }
        }
    }

    return memo[m][n]; // 返回从两个字符串的尾部开始，是否相互匹配
};
```

</div></div>
</div></div>

**类似题目**：
  - [44. 通配符匹配 🔴](/problems/wildcard-matching)
  - [剑指 Offer 19. 正则表达式匹配 🔴](/problems/zheng-ze-biao-da-shi-pi-pei-lcof)

</details>
</div>



