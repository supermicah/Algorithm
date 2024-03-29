<p>给定一个仅包含数字&nbsp;<code>2-9</code>&nbsp;的字符串，返回所有它能表示的字母组合。答案可以按 <strong>任意顺序</strong> 返回。</p>

<p>给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/11/09/200px-telephone-keypad2svg.png" style="width: 200px;" /></p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>digits = "23"
<strong>输出：</strong>["ad","ae","af","bd","be","bf","cd","ce","cf"]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>digits = ""
<strong>输出：</strong>[]
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>digits = "2"
<strong>输出：</strong>["a","b","c"]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>0 &lt;= digits.length &lt;= 4</code></li> 
 <li><code>digits[i]</code> 是范围 <code>['2', '9']</code> 的一个数字。</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>哈希表 | 字符串 | 回溯</details><br>

<div>👍 2445, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

你需要先看前文 [回溯算法详解](https://labuladong.github.io/article/fname.html?fname=回溯算法详解修订版) 和 [回溯算法团灭子集、排列、组合问题](https://labuladong.github.io/article/fname.html?fname=子集排列组合)，然后看这道题就很简单了，无非是回溯算法的运用而已。

组合问题本质上就是遍历一棵回溯树，套用回溯算法代码框架即可。

**标签：[回溯算法](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122002916411604996)，[数学](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122023604245659649)**

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
    // 每个数字到字母的映射
    vector<string> mapping = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    vector<string> res;

public:
    vector<string> letterCombinations(string digits) {
        if (digits.empty()) {
            return res;
        }
        // 从 digits[0] 开始进行回溯
        backtrack(digits, 0, "");
        return res;
    }

    // 回溯算法主函数
    void backtrack(const string& digits, int start, string cur) {
        if (cur.size() == digits.size()) {
            // 到达回溯树底部
            res.push_back(cur);
            return;
        }
        // 回溯算法框架
        for (int i = start; i < digits.size(); i++) {
            int digit = digits[i] - '0';
            for (char c : mapping[digit]) {
                // 做选择
                cur.push_back(c);
                // 递归下一层回溯树
                backtrack(digits, i + 1, cur);
                // 撤销选择
                cur.pop_back();
            }
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
    # 每个数字到字母的映射
    mapping = ["", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]

    def __init__(self):
        self.res = []

    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return self.res
        # 从 digits[0] 开始进行回溯
        self.backtrack(digits, 0, [])
        return self.res

    # 回溯算法主函数
    def backtrack(self, digits: str, start: int, path: List[str]):
        if len(path) == len(digits):
            # 到达回溯树底部
            self.res.append(''.join(path))
            return
        # 回溯算法框架
        for i in range(start, len(digits)):
            digit = int(digits[i])
            for c in self.mapping[digit]:
                # 做选择
                path.append(c)
                # 递归下一层回溯树
                self.backtrack(digits, i + 1, path)
                # 撤销选择
                path.pop()
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    // 每个数字到字母的映射
    String[] mapping = new String[] {
            "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"
    };

    List<String> res = new LinkedList<>();

    public List<String> letterCombinations(String digits) {
        if (digits.isEmpty()) {
            return res;
        }
        // 从 digits[0] 开始进行回溯
        backtrack(digits, 0, new StringBuilder());
        return res;
    }

    // 回溯算法主函数
    void backtrack(String digits, int start, StringBuilder sb) {
        if (sb.length() == digits.length()) {
            // 到达回溯树底部
            res.add(sb.toString());
            return;
        }
        // 回溯算法框架
        for (int i = start; i < digits.length(); i++) {
            int digit = digits.charAt(i) - '0';
            for (char c : mapping[digit].toCharArray()) {
                // 做选择
                sb.append(c);
                // 递归下一层回溯树
                backtrack(digits, i + 1, sb);
                // 撤销选择
                sb.deleteCharAt(sb.length() - 1);
            }
        }
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

func letterCombinations(digits string) []string {
    mapping := []string{
        "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
    }
    var res []string
    if len(digits) == 0 {
        return res
    }
    // 从 digits[0] 开始进行回溯
    backtrack(digits, 0, &strings.Builder{}, mapping, &res)
    return res
}

// 回溯算法主函数
func backtrack(digits string, start int, sb *strings.Builder, mapping []string, res *[]string) {
    if sb.Len() == len(digits) {
        // 到达回溯树底部
        *res = append(*res, sb.String())
        return
    }
    // 回溯算法框架
    for i := start; i < len(digits); i++ {
        digit := digits[i] - '0'
        for _, c := range mapping[digit] {
            // 做选择
            sb.WriteRune(c)
            // 递归下一层回溯树
            backtrack(digits, i+1, sb, mapping, res)
            // 撤销选择
            sb.Truncate(sb.Len() - 1)
        }
    }
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var letterCombinations = function(digits) {
    // 每个数字到字母的映射
    const mapping = ["", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"];

    const res = [];

    function backtrack(start, sb) {
        if (sb.length === digits.length) {
            // 到达回溯树底部
            res.push(sb.join(''));
            return;
        }
        // 回溯算法框架
        for (let i = start; i < digits.length; i++) {
            const digit = digits.charAt(i) - '0';
            for (const c of mapping[digit]) {
                // 做选择
                sb.push(c);
                // 递归下一层回溯树
                backtrack(i + 1, sb);
                // 撤销选择
                sb.pop();
            }
        }
    }

    if (digits.length === 0) {
        return res;
    }
    // 从 digits[0] 开始进行回溯
    backtrack(0, []);
    return res;
};
```

</div></div>
</div></div>

</details>
</div>





