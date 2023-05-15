<p>实现&nbsp;<a href="https://www.cplusplus.com/reference/valarray/pow/">pow(<em>x</em>,&nbsp;<em>n</em>)</a>&nbsp;，即计算 x 的 n 次幂函数（即，x<sup>n</sup>）。不得使用库函数，同时不需要考虑大数问题。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>x = 2.00000, n = 10
<strong>输出：</strong>1024.00000
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>x = 2.10000, n = 3
<strong>输出：</strong>9.26100</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>x = 2.00000, n = -2
<strong>输出：</strong>0.25000
<strong>解释：</strong>2<sup>-2</sup> = 1/2<sup>2</sup> = 1/4 = 0.25</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>-100.0 &lt;&nbsp;x&nbsp;&lt; 100.0</code></li> 
 <li><code>-2<sup>31</sup>&nbsp;&lt;= n &lt;=&nbsp;2<sup>31</sup>-1</code></li> 
 <li><code>-10<sup>4</sup>&nbsp;&lt;= x<sup>n</sup>&nbsp;&lt;= 10<sup>4</sup></code></li> 
</ul>

<p>&nbsp;</p>

<p>注意：本题与主站 50 题相同：<a href="https://leetcode-cn.com/problems/powx-n/">https://leetcode-cn.com/problems/powx-n/</a></p>

<details><summary><strong>Related Topics</strong></summary>递归 | 数学</details><br>

<div>👍 415, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题和 [50. Pow(x, n)](/problems/powx-n) 相同。

幂运算是经典的数学运算技巧了，建议你看下前文 [如何高效进行模幂运算](https://labuladong.github.io/article/fname.html?fname=superPower) 就能很容易理解解法代码里的思想了。这道题唯一有点恶心的就是 `k` 的取值范围特别大，不能直接加符号，否则会造成整型溢出，具体解法看代码吧。

**标签：[数学](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122023604245659649)**

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
public:
    double myPow(double a, int k) {
        if (k == 0) return 1;

        if (k == INT_MIN) {
            // 把 k 是 INT_MIN 的情况单独拿出来处理
            // 避免 -k 整型溢出
            return myPow(1 / a, -(k + 1)) / a;
        }

        if (k < 0) {
            return myPow(1 / a, -k);
        }

        if (k % 2 == 1) {
            // k 是奇数
            return (a * myPow(a, k - 1));
        } else {
            // k 是偶数
            double sub = myPow(a, k / 2);
            return (sub * sub);
        }
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def myPow(self, a: float, k: int) -> float:
        if k == 0:
            return 1

        if k == -2147483648:  # k 是 INT_MIN 的情况单独拿出来处理
            # 避免 -k 整型溢出
            return self.myPow(1 / a, -(k + 1)) / a

        if k < 0:
            return self.myPow(1 / a, -k)

        if k % 2 == 1:  # k 是奇数
            return a * self.myPow(a, k - 1)
        else:  # k 是偶数
            sub = self.myPow(a, k // 2)
            return sub * sub
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public double myPow(double a, int k) {
        if (k == 0) return 1;

        if (k == Integer.MIN_VALUE) {
            // 把 k 是 INT_MIN 的情况单独拿出来处理
            // 避免 -k 整型溢出
            return myPow(1 / a, -(k + 1)) / a;
        }

        if (k < 0) {
            return myPow(1 / a, -k);
        }

        if (k % 2 == 1) {
            // k 是奇数
            return (a * myPow(a, k - 1));
        } else {
            // k 是偶数
            double sub = myPow(a, k / 2);
            return (sub * sub);
        }
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func myPow(a float64, k int) float64 {
    if k == 0 { // 如果 k 等于 0，直接返回 1
        return 1
    }

    if k == -2147483648 { // 把 k 是 INT_MIN 的情况单独拿出来处理
        // 避免 -k 整型溢出
        return myPow(1 / a, -(k + 1)) / a
    }

    if k < 0 { // 如果 k 小于 0，递归转换成 1 / a 的相反数
        return myPow(1 / a, -k)
    }

    if k % 2 == 1 { // 如果 k 是奇数，递归运算 k - 1
        return a * myPow(a, k - 1)
    } else { // 如果 k 是偶数，递归运算 k / 2
        sub := myPow(a, k / 2)
        return sub * sub
    }
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * @param {number} a
 * @param {number} k
 * @return {number}
 */
var myPow = function(a, k) {
    if (k === 0) return 1;

    if (k === -2147483648) {
        // 把 k 是 INT_MIN 的情况单独拿出来处理
        // 避免 -k 整型溢出
        return myPow(1 / a, -(k + 1)) / a;
    }

    if (k < 0) {
        return myPow(1 / a, -k);
    }

    if (k % 2 === 1) {
        // k 是奇数
        return (a * myPow(a, k - 1));
    } else {
        // k 是偶数
        var sub = myPow(a, k / 2);
        return (sub * sub);
    }
};
```

</div></div>
</div></div>

**类似题目**：
  - [剑指 Offer 16. 数值的整数次方 🟠](/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/)

</details>
</div>



