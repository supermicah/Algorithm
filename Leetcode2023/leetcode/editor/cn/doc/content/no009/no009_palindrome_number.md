<p>给你一个整数 <code>x</code> ，如果 <code>x</code> 是一个回文整数，返回 <code>true</code> ；否则，返回 <code>false</code> 。</p>

<p>回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。</p>

<ul> 
 <li>例如，<code>121</code> 是回文，而 <code>123</code> 不是。</li> 
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>x = 121
<strong>输出：</strong>true
</pre>

<p><strong>示例&nbsp;2：</strong></p>

<pre>
<strong>输入：</strong>x = -121
<strong>输出：</strong>false
<strong>解释：</strong>从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>x = 10
<strong>输出：</strong>false
<strong>解释：</strong>从右向左读, 为 01 。因此它不是一个回文数。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>-2<sup>31</sup>&nbsp;&lt;= x &lt;= 2<sup>31</sup>&nbsp;- 1</code></li> 
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>你能不将整数转为字符串来解决这个问题吗？</p>

<details><summary><strong>Related Topics</strong></summary>数学</details><br>

<div>👍 2515, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

如果让你判断回文串应该很简单，我在 [数组双指针技巧汇总](https://labuladong.github.io/article/fname.html?fname=双指针技巧) 中讲过。

操作数字没办法像操作字符串那么简单粗暴，但只要你要知道我在 [Rabin Karp 算法详解](https://labuladong.github.io/article/fname.html?fname=rabinkarp) 中讲到的从最高位开始生成数字的技巧，就能轻松解决这个问题：

```java
String s = "8264";
int number = 0;
for (int i = 0; i < s.size(); i++) {
    // 将字符转化成数字
    number = 10 * number + (s[i] - '0');
    print(number);
}
// 打印输出：
// 8
// 82
// 826
// 8264
```

你**从后往前**把 `x` 的每一位拿出来，用这个技巧生成一个数字 `y`，如果 `y` 和 `x` 相等，则说明 `x` 是回文数字。

如何**从后往前**拿出一个数字的每一位？和 10 求余数就行了呗。看代码吧。

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
    bool isPalindrome(int x) {
        if (x < 0) {
            return false;
        }
        int temp = x;
        // y 是 x 翻转后的数字
        int y = 0;
        while (temp > 0) {
            int last_num = temp % 10;
            temp = temp / 10;
            // 从最高位生成数字的技巧
            y = y * 10 + last_num;
        }
        return y == x;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        temp = x
        # y 是 x 翻转后的数字
        y = 0
        while temp > 0:
            last_num = temp % 10
            temp = temp // 10
            # 从最高位生成数字的技巧
            y = y * 10 + last_num
        return y == x
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) {
            return false;
        }
        int temp = x;
        // y 是 x 翻转后的数字
        int y = 0;
        while (temp > 0) {
            int last_num = temp % 10;
            temp = temp / 10;
            // 从最高位生成数字的技巧
            y = y * 10 + last_num;
        }
        return y == x;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

// IsPalindrome 是判断一个整数是否为回文数的函数
func IsPalindrome(x int) bool {
	// 如果 x 是负数，那么它不可能为回文数，直接返回 false
	if x < 0 {
		return false
	}

	// temp 是 x 的副本
	temp := x
	// y 是 x 翻转后的数字
	y := 0

	// 将 temp 逆序生成 y
	for temp > 0 {
		lastNum := temp % 10
		temp = temp / 10
		y = y*10 + lastNum
	}

	// 如果 x 和 y 相等，那么 x 就是回文数
	return y == x
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    if (x < 0) {
        return false;
    }
    var temp = x;
    // y 是 x 翻转后的数字
    var y = 0;
    while (temp > 0) {
        var last_num = temp % 10;
        temp = Math.floor(temp / 10);
        // 从最高位生成数字的技巧
        y = y * 10 + last_num;
    }
    return y == x;
};
```

</div></div>
</div></div>

</details>
</div>



