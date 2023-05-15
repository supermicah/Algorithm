<p>编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为&nbsp;<a href="http://en.wikipedia.org/wiki/Hamming_weight" target="_blank">汉明重量</a>).）。</p>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li>请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。</li> 
 <li>在 Java 中，编译器使用 <a href="https://baike.baidu.com/item/二进制补码/5295284">二进制补码</a> 记法来表示有符号整数。因此，在上面的&nbsp;<strong>示例 3&nbsp;</strong>中，输入表示有符号整数 <code>-3</code>。</li> 
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>n = 11 (控制台输入 00000000000000000000000000001011)
<strong>输出：</strong>3
<strong>解释：</strong>输入的二进制串 <span><code><strong>00000000000000000000000000001011</strong>&nbsp;中，共有三位为 '1'。</code></span>
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>n = 128 (控制台输入 00000000000000000000000010000000)
<strong>输出：</strong>1
<strong>解释：</strong>输入的二进制串 <strong>00000000000000000000000010000000</strong>&nbsp;中，共有一位为 '1'。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>n = 4294967293 (控制台输入 11111111111111111111111111111101，部分语言中 n = -3）
<strong>输出：</strong>31
<strong>解释：</strong>输入的二进制串 <strong>11111111111111111111111111111101</strong> 中，共有 31 位为 '1'。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li>输入必须是长度为 <code>32</code> 的 <strong>二进制串</strong> 。</li> 
</ul>

<p>&nbsp;</p>

<p>注意：本题与主站 191 题相同：<a href="https://leetcode-cn.com/problems/number-of-1-bits/">https://leetcode-cn.com/problems/number-of-1-bits/</a></p>

<details><summary><strong>Related Topics</strong></summary>位运算</details><br>

<div>👍 317, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**



<p><strong><a href="https://labuladong.gitee.io/article/slug.html?slug=er-jin-zhi-zhong-1de-ge-shu-lcof" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题和 [191. 位 1 的个数](/problems/number-of-1-bits) 相同。

`n & (n-1)` 这个操作是算法中常见的，作用是消除数字 `n` 的二进制表示中的最后一个 1：

![](https://labuladong.github.io/pictures/位操作/1.png)

不断消除数字 `n` 中的 1，直到 `n` 变为 0。

**详细题解：[常用的位操作](https://labuladong.github.io/article/fname.html?fname=常用的位操作)**

**标签：[位运算](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122023604245659649)，[数学](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122023604245659649)**

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
    int hammingWeight(uint32_t n) {
        int res = 0;
        while (n != 0) {
            n = n & (n - 1);
            res++;
        }
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
    def hammingWeight(self, n: int) -> int:
        """
        :type n: int
        :rtype: int
        """
        res = 0
        while n != 0:
            n &= n - 1
            res += 1
        return res
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
public class Solution {
    // you need to treat n as an unsigned value
    public int hammingWeight(int n) {
        int res = 0;
        while (n != 0) {
            n = n & (n - 1);
            res++;
        }
        return res;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func hammingWeight(num uint32) int {
    res := 0
    for num != 0 {
        num &= (num - 1)
        res++
    }
    return res
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * @param {number} n - a positive integer
 * @return {number} - the number of bits set to 1 in the binary representation of n
 */
var hammingWeight = function(n) {
    let res = 0;
    while (n !== 0) {
        n = n & (n - 1);
        res++;
    }
    return res;
};
```

</div></div>
</div></div>

**类似题目**：
  - [136. 只出现一次的数字 🟢](/problems/single-number)
  - [231. 2 的幂 🟢](/problems/power-of-two)
  - [268. 丢失的数字 🟢](/problems/missing-number)
  - [剑指 Offer 15. 二进制中1的个数 🟢](/problems/er-jin-zhi-zhong-1de-ge-shu-lcof)

</details>
</div>



