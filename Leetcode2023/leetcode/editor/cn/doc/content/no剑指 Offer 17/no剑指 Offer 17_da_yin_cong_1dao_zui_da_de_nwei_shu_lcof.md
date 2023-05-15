<p>输入数字 <code>n</code>，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> n = 1
<strong>输出:</strong> [1,2,3,4,5,6,7,8,9]
</pre>

<p>&nbsp;</p>

<p>说明：</p>

<ul> 
 <li>用返回一个整数列表来代替打印</li> 
 <li>n 为正整数</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>数组 | 数学</details><br>

<div>👍 294, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这题很简单，估计就是想考察整数的生成吧，你只要知道这个字符串转数字的技巧就行了：

```java
String s = "8264";
int number = 0;
for (int i = 0; i < s.size(); i++) {
    // 将字符转化成数字
    number = 10 * number + (s[i] - '0');
}
```

**标签：字符串**

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
    vector<int> printNumbers(int n) {
        // 先生成 n 位 9
        int max = 0;
        for (int i = 0; i < n; i++) {
            max = 10 * max + 9;
        }
        // 填入结果数组
        vector<int> res(max);
        for (int i = 1; i <= max; i++) {
            res[i - 1] = i;
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
    def printNumbers(self, n: int) -> List[int]:
        # 先生成 n 位 9
        max_num = 0
        for i in range(n):
            max_num = max_num * 10 + 9
            
        # 填入结果数组
        res = []
        for i in range(1, max_num + 1):
            res.append(i)
        
        return res
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public int[] printNumbers(int n) {
        // 先生成 n 位 9
        int max = 0;
        for (int i = 0; i < n; i++) {
            max = 10 * max + 9;
        }
        // 填入结果数组
        int[] res = new int[max];
        for (int i = 1; i <= max; i++) {
            res[i - 1] = i;
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

func printNumbers(n int) []int {
    // 先生成 n 位 9
    max := 0
    for i := 0; i < n; i++ {
        max = 10 * max + 9
    }
    // 填入结果数组
    res := make([]int, max)
    for i := 1; i <= max; i++ {
        res[i-1] = i
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
 * @param {number} n
 * @return {number[]}
 */
var printNumbers = function(n) {
    // 先生成 n 位 9
    let max = 0;
    for (let i = 0; i < n; i++) {
        max = 10 * max + 9;
    }
    // 填入结果数组
    let res = new Array(max);
    for (let i = 1; i <= max; i++) {
        res[i - 1] = i;
    }
    return res;
};
```

</div></div>
</div></div>

</details>
</div>



