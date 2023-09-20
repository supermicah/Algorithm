<p>定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。</p>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 1-&gt;2-&gt;3-&gt;4-&gt;5-&gt;NULL
<strong>输出:</strong> 5-&gt;4-&gt;3-&gt;2-&gt;1-&gt;NULL</pre>

<p>&nbsp;</p>

<p><strong>限制：</strong></p>

<p><code>0 &lt;= 节点个数 &lt;= 5000</code></p>

<p>&nbsp;</p>

<p><strong>注意</strong>：本题与主站 206 题相同：<a href="https://leetcode-cn.com/problems/reverse-linked-list/">https://leetcode-cn.com/problems/reverse-linked-list/</a></p>

<details><summary><strong>Related Topics</strong></summary>递归 | 链表</details><br>

<div>👍 629, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员，[新版刷题打卡挑战](https://labuladong.gitee.io/algo/challenge/) 上线！**



<p><strong><a href="https://labuladong.github.io/article/slug.html?slug=fan-zhuan-lian-biao-lcof" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题和 [206. 反转链表](/problems/reverse-linked-list) 相同。

递归实现反转链表常常用来考察递归思想，我这里就用纯递归来翻转链表。

**对于递归算法，最重要的就是明确递归函数的定义**。具体来说，我们的 `reverse` 函数定义是这样的：

**输入一个节点 `head`，将「以 `head` 为起点」的链表反转，并返回反转之后的头结点**。

明白了函数的定义，再来看这个问题。比如说我们想反转这个链表：

![](https://labuladong.github.io/pictures/反转链表/1.jpg)

那么输入 `reverse(head)` 后，会在这里进行递归：

```java
ListNode last = reverse(head.next);
```

不要跳进递归（你的脑袋能压几个栈呀？），而是要根据刚才的函数定义，来弄清楚这段代码会产生什么结果：

![](https://labuladong.github.io/pictures/反转链表/2.jpg)

这个 `reverse(head.next)` 执行完成后，整个链表就成了这样：

![](https://labuladong.github.io/pictures/反转链表/3.jpg)

并且根据函数定义，`reverse` 函数会返回反转之后的头结点，我们用变量 `last` 接收了。

现在再来看下面的代码：

```java
head.next.next = head;
```

![](https://labuladong.github.io/pictures/反转链表/4.jpg)

接下来：

```java
head.next = null;
return last;
```

![](https://labuladong.github.io/pictures/反转链表/5.jpg)

神不神奇，这样整个链表就反转过来了！

**详细题解：[递归魔法：反转单链表](https://labuladong.github.io/article/fname.html?fname=递归反转链表的一部分)**

**标签：单链表**

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
    ListNode* reverseList(ListNode* head) {
        // 如果链表为空或只有一个节点，则直接返回该链表
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        ListNode* last = nullptr;
        // 开始迭代翻转节点
        while (head != nullptr) {
            // 先保存当前节点的后继节点
            ListNode* next = head->next;
            // 将当前节点的指针指向前一个节点
            head->next = last;
            // 更新前一个节点为当前节点
            last = head;
            // 更新当前节点为后继节点
            head = next;
        }
        // 翻转链表完成，返回头节点
        return last;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        # 如果链表为空或链表只有一个元素，则直接返回head
        if head == None or head.next==None:
            return head
        # 递归反转链表后面的部分
        last = self.reverseList(head.next)
        """
        反转之前的链表
        head      --> 3 --> 4 --> 5 --> null
        last      --> 5 --> null
        """
        # 将链表的下一个节点的后继指针指向当前节点
        head.next.next = head
        """
        反转之后的链表
        last      --> 5 --> 4 --> 3 --> null
                          ↑
        head      --> 4 ---┘
        """
        # 将当前节点的后继指针指向null
        head.next = None
        """
        反转之后的链表
        last      --> 5 --> 4 --> 3 --> null
        head      --> null
        """
        # 返回反转后的链表
        return last
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode last = reverseList(head.next);/**<extend up -200>![](https://labuladong.github.io/pictures/反转链表/3.jpg) */
        head.next.next = head;/**<extend up -200>![](https://labuladong.github.io/pictures/反转链表/4.jpg) */
        head.next = null;/**<extend up -200>![](https://labuladong.github.io/pictures/反转链表/5.jpg) */
        return last;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    // 如果链表为空或者链表只有一个节点，直接返回head
    if head == nil || head.Next == nil {
        return head
    }
    // 递归处理下一个节点
    last := reverseList(head.Next)
    // 对于当前节点：
    // 原来指向下一个节点的指针反转指向上一个节点
    head.Next.Next = head
    // 原来指向下一个节点的指针指向nil
    head.Next = nil
    // 返回链表的最后一个节点，即反转后链表的头节点
    return last
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
  if (head === null || head.next === null) {
    return head;
  }
  var last = reverseList(head.next);/**<extend up -200>![](https://labuladong.github.io/pictures/反转链表/3.jpg) */
  head.next.next = head;/**<extend up -200>![](https://labuladong.github.io/pictures/反转链表/4.jpg) */
  head.next = null;/**<extend up -200>![](https://labuladong.github.io/pictures/反转链表/5.jpg) */
  return last;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🎃🎃 算法可视化 🎃🎃</strong></summary><div id="data_reverse-linked-list" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_reverse-linked-list"></div></div>
</details><hr /><br />

**类似题目**：
  - [92. 反转链表 II 🟠](/problems/reverse-linked-list-ii)
  - [剑指 Offer 24. 反转链表 🟢](/problems/fan-zhuan-lian-biao-lcof)
  - [剑指 Offer II 024. 反转链表 🟢](/problems/UHnkqh)

</details>
</div>

