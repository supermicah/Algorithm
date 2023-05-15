<p>给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。</p>

<p>返回删除后的链表的头节点。</p>

<p><strong>注意：</strong>此题对比原题有改动</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> head = [4,5,1,9], val = 5
<strong>输出:</strong> [4,1,9]
<strong>解释: </strong>给定你链表中值为&nbsp;5&nbsp;的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -&gt; 1 -&gt; 9.
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> head = [4,5,1,9], val = 1
<strong>输出:</strong> [4,5,9]
<strong>解释: </strong>给定你链表中值为&nbsp;1&nbsp;的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -&gt; 5 -&gt; 9.
</pre>

<p>&nbsp;</p>

<p><strong>说明：</strong></p>

<ul> 
 <li>题目保证链表中节点的值互不相同</li> 
 <li>若使用 C 或 C++ 语言，你不需要 <code>free</code> 或 <code>delete</code> 被删除的节点</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>链表</details><br>

<div>👍 310, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题常规的思路是通过操作指针来删除值为 `val` 的节点，不过处理起来可能稍微有点麻烦，我们可以使用 [链表双指针技巧汇总](https://labuladong.github.io/article/fname.html?fname=链表技巧) 中讲到的分解链表的思路来解决这道题（你可以先去做下 [86. 分隔链表](/problems/partition-list)）。

你可以认为这道题是把原链表分解成「值为 `val`」和「值不为 `val`」的两条链表，就可以复用 86 题的思路了。

**标签：[链表双指针](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120596033251475465)**

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
    ListNode* deleteNode(ListNode* head, int val) {
        // 存放删除 val 的链表
        ListNode* dummy = new ListNode(-1);
        // q 指针负责生成结果链表
        ListNode* q = dummy;
        // p 负责遍历原链表
        ListNode* p = head;
        while (p != nullptr) {
            if (p->val != val) {
                // 把值不为 val 的节点接到结果链表上
                q->next = p;
                q = q->next;
            }
            // 断开原链表中的每个节点的 next 指针
            ListNode* temp = p->next;
            p->next = nullptr;
            p = temp;
        }

        return dummy->next;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def deleteNode(self, head: ListNode, val: int) -> ListNode:
        # 存放删除 val 的链表
        dummy = ListNode(-1)
        # q 指针负责生成结果链表
        q = dummy
        # p 负责遍历原链表
        p = head
        while p:
            if p.val != val:
                # 把值不为 val 的节点接到结果链表上
                q.next = p
                q = q.next
            # 断开原链表中的每个节点的 next 指针
            temp = p.next
            p.next = None
            p = temp

        return dummy.next
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public ListNode deleteNode(ListNode head, int val) {
        // 存放删除 val 的链表
        ListNode dummy = new ListNode(-1);
        // q 指针负责生成结果链表
        ListNode q = dummy;
        // p 负责遍历原链表
        ListNode p = head;
        while (p != null) {
            if (p.val != val) {
                // 把值不为 val 的节点接到结果链表上
                q.next = p;
                q = q.next;
            }
            // 断开原链表中的每个节点的 next 指针
            ListNode temp = p.next;
            p.next = null;
            p = temp;
        }

        return dummy.next;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func deleteNode(head *ListNode, val int) *ListNode {
    // 存放删除 val 的链表
    dummy := &ListNode{
        Val: -1,
        Next: nil,
    }
    // q 指针负责生成结果链表
    q := dummy
    // p 负责遍历原链表
    p := head
    for p != nil {
        if p.Val != val {
            // 把值不为 val 的节点接到结果链表上
            q.Next = p
            q = q.Next
        }
        // 断开原链表中的每个节点的 next 指针
        temp := p.Next
        p.Next = nil
        p = temp
    }

    return dummy.Next
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

/**
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 */
var deleteNode = function(head, val) {
    // 存放删除 val 的链表
    let dummy = new ListNode(-1);
    // q 指针负责生成结果链表
    let q = dummy;
    // p 负责遍历原链表
    let p = head;
    while (p != null) {
        if (p.val != val) {
            // 把值不为 val 的节点接到结果链表上
            q.next = p;
            q = q.next;
        }
        // 断开原链表中的每个节点的 next 指针
        let temp = p.next;
        p.next = null;
        p = temp;
    }

    return dummy.next;
};
```

</div></div>
</div></div>

</details>
</div>



