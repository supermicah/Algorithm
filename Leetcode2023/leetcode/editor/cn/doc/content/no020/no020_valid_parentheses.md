<p>给定一个只包括 <code>'('</code>，<code>')'</code>，<code>'{'</code>，<code>'}'</code>，<code>'['</code>，<code>']'</code>&nbsp;的字符串 <code>s</code> ，判断字符串是否有效。</p>

<p>有效字符串需满足：</p>

<ol> 
 <li>左括号必须用相同类型的右括号闭合。</li> 
 <li>左括号必须以正确的顺序闭合。</li> 
 <li>每个右括号都有一个对应的相同类型的左括号。</li> 
</ol>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "()"
<strong>输出：</strong>true
</pre>

<p><strong>示例&nbsp;2：</strong></p>

<pre>
<strong>输入：</strong>s = "()[]{}"
<strong>输出：</strong>true
</pre>

<p><strong>示例&nbsp;3：</strong></p>

<pre>
<strong>输入：</strong>s = "(]"
<strong>输出：</strong>false
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= s.length &lt;= 10<sup>4</sup></code></li> 
 <li><code>s</code> 仅由括号 <code>'()[]{}'</code> 组成</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>栈 | 字符串</details><br>

<div>👍 3904, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**



<p><strong><a href="https://labuladong.github.io/article/slug.html?slug=valid-parentheses" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

栈是一种先进后出的数据结构，处理括号问题的时候尤其有用。

遇到左括号就入栈，遇到右括号就去栈中寻找最近的左括号，看是否匹配。

**详细题解：[如何解决括号相关的问题](https://labuladong.github.io/article/fname.html?fname=括号插入)**

**标签：括号问题，[栈](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2121993002939219969)**

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
    bool isValid(string str) {
        stack<char> left;
        for (char c : str) {
            if (c == '(' || c == '{' || c == '[')
                left.push(c);
            else // 字符 c 是右括号
                if (!left.empty() && leftOf(c) == left.top())
                    left.pop();
                else
                    // 和最近的左括号不匹配
                    return false;
        }
        // 是否所有的左括号都被匹配了
        return left.empty();
    }

    char leftOf(char c) {
        if (c == '}') return '{';
        if (c == ')') return '(';
        return '[';
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def isValid(self, s: str) -> bool:
        left = []  # 使用栈结构，存储所有待匹配的左括号
        for c in s:
            if c == '(' or c == '{' or c == '[':
                left.append(c)  # 如果字符 c 是左括号，则将其加入左括号栈 left 中
            else:
                if left and self.leftOf(c) == left[-1]:  # 如果字符 c 是右括号，则比较它与最近一次加入栈 left 中的左括号是否匹配
                    left.pop()  # 如果匹配，则将最近的左括号出栈，否则返回 False
                else:
                    return False
        return not left  # 最后判断栈是否为空，如果是则说明所有的左括号都被匹配了，返回 True，否则返回 False

    def leftOf(self, c: str) -> str:
        if c == '}':
            return '{'
        elif c == ')':
            return '('
        else:
            return '['
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public boolean isValid(String str) {
        Stack<Character> left = new Stack<>();
        for (char c : str.toCharArray()) {
            if (c == '(' || c == '{' || c == '[')
                left.push(c);
            else // 字符 c 是右括号
                if (!left.isEmpty() && leftOf(c) == left.peek())
                    left.pop();
                else
                    // 和最近的左括号不匹配
                    return false;
        }
        // 是否所有的左括号都被匹配了
        return left.isEmpty();
    }

    char leftOf(char c) {
        if (c == '}') return '{';
        if (c == ')') return '(';
        return '[';
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

// 包名为 main
// 定义一个字节数组栈类型
type Stack []byte

// 入栈操作
func (s *Stack) push(str byte) {
    *s = append(*s, str)
}

// 出栈操作
func (s *Stack) pop() byte {
    if len(*s) == 0 {
        return 0
    }
    res := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return res
}

// 判断给定字符串是否是合法的括号序列
func isValid(str string) bool {
    // 定义一个栈 left 保存左括号
    var left Stack
    // 遍历字符
    for i := range str {
        c := str[i]
        // 当 c 是左括号时，入栈 left
        if c == '(' || c == '[' || c == '{' {
            left.push(c)
        } else { // 当 c 是右括号时
            // 如果栈 left 非空，且栈顶的左括号和当前右括号匹配，则弹出栈顶元素
            if len(left) != 0 && leftOf(c) == left.pop() {
                continue
            } else { // 当前左括号和最近的左括号不匹配
                return false
            }
        }
    }
    // 是否所有的左括号都被匹配了
    return len(left) == 0
}

// 返回左括号
func leftOf(c byte) byte {
    if c == '}' {
        return '{'
    } else if c == ')' {
        return '('
    } else {
        return '['
    }
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var isValid = function(str) {
    // 建立一个栈
    let left = [];
    // 遍历字符串中的每一个字符
    for (let c of str) {
        // 如果是左括号，则入栈
        if (c == '(' || c == '{' || c == '[')
            left.push(c);
        else { // 字符 c 是右括号
            // 如果栈不为空，并且最近入栈的左括号可以匹配，则出栈
            if (left.length && leftOf(c) == left[left.length-1])
                left.pop();
            else
                // 和最近的左括号不匹配
                return false;
        }
    }
    // 是否所有的左括号都被匹配了
    return !left.length;
}

function leftOf(c) {
    if (c == '}') return '{';
    if (c == ')') return '(';
    return '[';
}
```

</div></div>
</div></div>

**类似题目**：
  - [1541. 平衡括号字符串的最少插入次数 🟠](/problems/minimum-insertions-to-balance-a-parentheses-string)
  - [921. 使括号有效的最少添加 🟠](/problems/minimum-add-to-make-parentheses-valid)

</details>
</div>



