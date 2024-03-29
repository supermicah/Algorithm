<p>给你链表的头节点 <code>head</code> ，每&nbsp;<code>k</code><em>&nbsp;</em>个节点一组进行翻转，请你返回修改后的链表。</p>

<p><code>k</code> 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是&nbsp;<code>k</code><em>&nbsp;</em>的整数倍，那么请将最后剩余的节点保持原有顺序。</p>

<p>你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg" style="width: 542px; height: 222px;" /> 
<pre>
<strong>输入：</strong>head = [1,2,3,4,5], k = 2
<strong>输出：</strong>[2,1,4,3,5]
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg" style="width: 542px; height: 222px;" /></p>

<pre>
<strong>输入：</strong>head = [1,2,3,4,5], k = 3
<strong>输出：</strong>[3,2,1,4,5]
</pre>

<p>&nbsp;</p> 
<strong>提示：</strong>

<ul> 
 <li>链表中的节点数目为 <code>n</code></li> 
 <li><code>1 &lt;= k &lt;= n &lt;= 5000</code></li> 
 <li><code>0 &lt;= Node.val &lt;= 1000</code></li> 
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>你可以设计一个只用 <code>O(1)</code> 额外内存空间的算法解决此问题吗？</p>

<ul> 
</ul>

<details><summary><strong>Related Topics</strong></summary>递归 | 链表</details><br>

<div>👍 2168, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员，[新版刷题打卡挑战](https://labuladong.gitee.io/algo/challenge/) 上线！**



<p><strong><a href="https://labuladong.github.io/article/slug.html?slug=reverse-nodes-in-k-group" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

PS：这道题在[《算法小抄》](https://item.jd.com/12759911.html) 的第 289 页。

输入 `head`，`reverseKGroup` 函数能够把以 `head` 为头的这条链表进行翻转。

我们要充分利用这个递归函数的定义，把原问题分解成规模更小的子问题进行求解。

**1、先反转以 `head` 开头的 `k` 个元素**。

![](https://labuladong.github.io/pictures/kgroup/3.jpg)

**2、将第 `k + 1` 个元素作为 `head` 递归调用 `reverseKGroup` 函数**。

![](https://labuladong.github.io/pictures/kgroup/4.jpg)

**3、将上述两个过程的结果连接起来**。

![](https://labuladong.github.io/pictures/kgroup/5.jpg)

最后函数递归完成之后就是这个结果，完全符合题意：

![](https://labuladong.github.io/pictures/kgroup/7.jpg)

**详细题解：[如何 K 个一组反转链表](https://labuladong.github.io/article/fname.html?fname=k个一组反转链表)**

**标签：[数据结构](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318892385270808576)，[链表](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120596033251475465)，[链表双指针](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120596033251475465)**

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
    ListNode* reverseKGroup(ListNode* head, int k) {
         if(head == NULL) return NULL;
    	//区间[a,b)包含k个待反转元素
    	ListNode *a, *b;
    	a = b = head;
    	for (int i = 0; i < k; i++) {
    		//不足k个，直接返回
    		if (b == NULL) return head;
    		b = b->next;
    	}
    	//反转前k个元素
    	ListNode *newHead = reverse(a, b);
    	//递归反转后续链表并连接起来
    	a->next = reverseKGroup(b, k);
    	return newHead;
    }
    ListNode *reverse(ListNode *a, ListNode *b) {
    	ListNode *pre, *cur, *nxt;
    	cur = a; pre = NULL;
    	while (cur != b) {
    		nxt = cur->next;
    		cur->next = pre;
    		pre = cur;
    		cur = nxt;
    	}
    	return pre;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        if not head:
            return None
        # 区间 [a, b) 包含 k 个待反转元素
        a = b = head
        for i in range(k):
            # 不足 k 个，不需要反转，base case
            if not b:
                return head
            b = b.next
        # 反转前 k 个元素
        newHead = self.reverse(a, b)
        # 递归反转后续链表并连接起来
        a.next = self.reverseKGroup(b, k) # <extend up -90>![](https://labuladong.github.io/pictures/kgroup/6.jpg) #
        return newHead

    """ 反转区间 [a, b) 的元素，注意是左闭右开 """
    def reverse(self, a: ListNode, b: ListNode) -> ListNode: # <extend up -300>![](https://labuladong.github.io/pictures/kgroup/8.gif) #
        pre, cur, nxt = None, a, a
        # while 终止的条件改一下就行了
        while cur != b:
            nxt = cur.next
            cur.next = pre
            pre = cur
            cur = nxt
        # 返回反转后的头结点
        return pre
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        if (head == null) return null;
        // 区间 [a, b) 包含 k 个待反转元素
        ListNode a, b;
        a = b = head;
        for (int i = 0; i < k; i++) {
            // 不足 k 个，不需要反转，base case
            if (b == null) return head;
            b = b.next;
        }
        // 反转前 k 个元素
        ListNode newHead = reverse(a, b);
        // 递归反转后续链表并连接起来
        a.next = reverseKGroup(b, k);/**<extend up -90>![](https://labuladong.github.io/pictures/kgroup/6.jpg) */
        return newHead;
    }

    /* 反转区间 [a, b) 的元素，注意是左闭右开 */
    ListNode reverse(ListNode a, ListNode b) {/**<extend up -300>![](https://labuladong.github.io/pictures/kgroup/8.gif) */
        ListNode pre, cur, nxt;
        pre = null;
        cur = a;
        nxt = a;
        // while 终止的条件改一下就行了
        while (cur != b) {
            nxt = cur.next;
            cur.next = pre;
            pre = cur;
            cur = nxt;
        }
        // 返回反转后的头结点
        return pre;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

// 给出一个链表和一个数k，比如1->2->3->4->5->6，k=2，那么你需要返回的结果是2->1->4->3->6->5。
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    
    a, b := head, head
    // 找到需要翻转的区间 [a, b)
    for i := 0; i < k; i++ {
        if b == nil {
            return head
        }
        b = b.Next
    }

    // 反转区间内的链表
    newHead := reverse(a, b)
    // 递归将后续链表的区间也翻转，然后再将它链接到新的区间内
    a.Next = reverseKGroup(b, k)/**<extend up -90>![](https://labuladong.github.io/pictures/kgroup/6.jpg) */

    return newHead
}

// 翻转区间 [a, b)，包头不包尾
func reverse(a, b *ListNode) *ListNode {
    // 初始化三个指针
    pre, cur, nxt := (*ListNode)(nil), a, a

    // 循环将当前节点指向前一个节点，然后将前一个节点和当前节点往后移动
    for cur != b {
        nxt = cur.Next
        cur.Next = pre
        pre, cur = cur, nxt
    }

    // 返回新的头结点
    return pre
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var reverseKGroup = function(head, k) {
    if (!head) return null;
    // 区间 [a, b) 包含 k 个待反转元素
    let a = head, b = head;
    for (let i = 0; i < k; i++) {
        // 不足 k 个，不需要反转，base case
        if (!b) return head;
        b = b.next;
    }
    // 反转前 k 个元素
    let newHead = reverse(a, b);
    // 递归反转后续链表并连接起来
    a.next = reverseKGroup(b, k);/**<extend up -90>![](https://labuladong.github.io/pictures/kgroup/6.jpg) */
    return newHead;
};

/* 反转区间 [a, b) 的元素，注意是左闭右开 */
var reverse = function(a, b) {/**<extend up -300>![](https://labuladong.github.io/pictures/kgroup/8.gif) */
    let pre = null, cur = a, nxt = a;
    // while 终止的条件改一下就行了
    while (cur !== b) {
        nxt = cur.next;
        cur.next = pre;
        pre = cur;
        cur = nxt;
    }
    // 返回反转后的头结点
    return pre;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🥳🥳 算法可视化 🥳🥳</strong></summary><div id="data_reverse-nodes-in-k-group" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_reverse-nodes-in-k-group"></div></div>
</details><hr /><br />

**类似题目**：
  - [24. 两两交换链表中的节点 🟠](/problems/swap-nodes-in-pairs)

</details>
</div>

