给你单链表的头指针 <code>head</code> 和两个整数 <code>left</code> 和 <code>right</code> ，其中 <code>left <= right</code> 。请你反转从位置 <code>left</code> 到位置 <code>right</code> 的链表节点，返回 <strong>反转后的链表</strong> 。

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/rev2ex2.jpg" style="width: 542px; height: 222px;" /> 
<pre>
<strong>输入：</strong>head = [1,2,3,4,5], left = 2, right = 4
<strong>输出：</strong>[1,4,3,2,5]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>head = [5], left = 1, right = 1
<strong>输出：</strong>[5]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li>链表中节点数目为 <code>n</code></li> 
 <li><code>1 &lt;= n &lt;= 500</code></li> 
 <li><code>-500 &lt;= Node.val &lt;= 500</code></li> 
 <li><code>1 &lt;= left &lt;= right &lt;= n</code></li> 
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong> 你可以使用一趟扫描完成反转吗？</p>

<details><summary><strong>Related Topics</strong></summary>链表</details><br>

<div>👍 1651, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员，[新版刷题打卡挑战](https://labuladong.gitee.io/algo/challenge/) 上线！**



<p><strong><a href="https://labuladong.github.io/article/slug.html?slug=reverse-linked-list-ii" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

PS：这道题在[《算法小抄》](https://item.jd.com/12759911.html) 的第 283 页。

迭代解法很简单，用一个 for 循环即可，但这道题经常用来考察递归思维，让你用纯递归的形式来解决，这里就给出递归解法的思路。

要想真正理解递归操作链表的代码思路，需要从递归翻转整条链表的算法开始，推导出递归翻转前 `N` 个节点的算法，最后改写出递归翻转区间内的节点的解法代码。

关键点还是要明确递归函数的定义，由于内容和图比较多，这里就不写了，请看详细题解。

**详细题解：[递归魔法：反转单链表](https://labuladong.github.io/article/fname.html?fname=递归反转链表的一部分)**

**标签：[数据结构](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318892385270808576)，递归，[链表](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120596033251475465)，[链表双指针](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120596033251475465)**

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
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        // base case
        if (m == 1) {
            return reverseN(head, n);
        }
        // 前进到反转的起点触发 base case
        head->next = reverseBetween(head->next, m - 1, n - 1);
        return head;
    }

private:
    ListNode* successor = nullptr; // 后驱节点
    // 反转以 head 为起点的 n 个节点，返回新的头结点
    ListNode* reverseN(ListNode* head, int n) {
        if (n == 1) {
            // 记录第 n + 1 个节点
            successor = head->next;
            return head;
        }
        // 以 head->next 为起点，需要反转前 n - 1 个节点
        ListNode* last = reverseN(head->next, n - 1);

        head->next->next = head;
        // 让反转之后的 head 节点和后面的节点连起来
        head->next = successor;
        return last;/**<extend up -90>![](https://labuladong.github.io/pictures/反转链表/7.jpg) */
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        # base case
        if m == 1:
            return self.reverseN(head, n)
        # 前进到反转的起点触发 base case
        head.next = self.reverseBetween(head.next, m - 1, n - 1)
        return head

    successor = None # 后驱节点
    # 反转以 head 为起点的 n 个节点，返回新的头结点
    def reverseN(self, head: ListNode, n: int) -> ListNode:
        if n == 1:
            # 记录第 n + 1 个节点
            self.successor = head.next
            return head
        # 以 head.next 为起点，需要反转前 n - 1 个节点
        last = self.reverseN(head.next, n - 1)

        head.next.next = head
        # 让反转之后的 head 节点和后面的节点连起来
        head.next = self.successor
        return last # <extend up -90>![](https://labuladong.github.io/pictures/反转链表/7.jpg) #
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public ListNode reverseBetween(ListNode head, int m, int n) {
        // base case
        if (m == 1) {
            return reverseN(head, n);
        }
        // 前进到反转的起点触发 base case
        head.next = reverseBetween(head.next, m - 1, n - 1);
        return head;
    }

    ListNode successor = null; // 后驱节点
    // 反转以 head 为起点的 n 个节点，返回新的头结点
    ListNode reverseN(ListNode head, int n) {
        if (n == 1) {
            // 记录第 n + 1 个节点
            successor = head.next;
            return head;
        }
        // 以 head.next 为起点，需要反转前 n - 1 个节点
        ListNode last = reverseN(head.next, n - 1);

        head.next.next = head;
        // 让反转之后的 head 节点和后面的节点连起来
        head.next = successor;
        return last;/**<extend up -90>![](https://labuladong.github.io/pictures/反转链表/7.jpg) */
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
 
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    var successor *ListNode // 后驱节点

    // 反转以 head 为起点的 n 个节点，返回新的头结点
    var reverseN func(head *ListNode, n int) (*ListNode)
    reverseN = func(head *ListNode, n int) (*ListNode) {
        if n == 1 {
            // 记录第 n + 1 个节点
            successor = head.Next
            return head
        }
        // 以 head.Next 为起点，需要反转前 n - 1 个节点
        last := reverseN(head.Next, n - 1)

        head.Next.Next = head
        // 让反转之后的 head 节点和后面的节点连起来
        head.Next = successor
        return last
    }

    // base case
    if m == 1 {
        return reverseN(head, n)
    }
    // 前进到反转的起点触发 base case
    head.Next = reverseBetween(head.Next, m - 1, n - 1)
    return head
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var reverseBetween = function(head, m, n) {
    let successor = null;
    // 反转以 head 为起点的 n 个节点，返回新的头结点
    const reverseN = function(head, n) {
        if (n == 1) {
            // 记录第 n + 1 个节点
            successor = head.next;
            return head;
        }
        const last = reverseN(head.next, n - 1);
        head.next.next = head;
        // 让反转之后的 head 节点和后面的节点连起来
        head.next = successor;
        return last;/**<extend up -90>![](https://labuladong.github.io/pictures/反转链表/7.jpg) */
    };
    // base case
    if (m == 1) {
        return reverseN(head, n);
    }
    // 前进到反转的起点触发 base case
    head.next = reverseBetween(head.next, m - 1, n - 1);
    return head;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🎃🎃 算法可视化 🎃🎃</strong></summary><div id="data_reverse-linked-list-ii" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_reverse-linked-list-ii"></div></div>
</details><hr /><br />

**类似题目**：
  - [206. 反转链表 🟢](/problems/reverse-linked-list)
  - [剑指 Offer 24. 反转链表 🟢](/problems/fan-zhuan-lian-biao-lcof)
  - [剑指 Offer II 024. 反转链表 🟢](/problems/UHnkqh)

</details>
</div>

