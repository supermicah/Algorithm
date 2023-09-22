<p>给你一个单链表的头节点 <code>head</code> ，请你判断该链表是否为回文链表。如果是，返回 <code>true</code> ；否则，返回 <code>false</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/03/03/pal1linked-list.jpg" style="width: 422px; height: 62px;" /> 
<pre>
<strong>输入：</strong>head = [1,2,2,1]
<strong>输出：</strong>true
</pre>

<p><strong>示例 2：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/03/03/pal2linked-list.jpg" style="width: 182px; height: 62px;" /> 
<pre>
<strong>输入：</strong>head = [1,2]
<strong>输出：</strong>false
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li>链表中节点数目在范围<code>[1, 10<sup>5</sup>]</code> 内</li> 
 <li><code>0 &lt;= Node.val &lt;= 9</code></li> 
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>你能否用&nbsp;<code>O(n)</code> 时间复杂度和 <code>O(1)</code> 空间复杂度解决此题？</p>

<details><summary><strong>Related Topics</strong></summary>栈 | 递归 | 链表 | 双指针</details><br>

<div>👍 1789, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员，[新版刷题打卡挑战](https://labuladong.gitee.io/algo/challenge/) 上线！**



<p><strong><a href="https://labuladong.github.io/article/slug.html?slug=palindrome-linked-list" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

PS：这道题在[《算法小抄》](https://item.jd.com/12759911.html) 的第 277 页。

这道题的关键在于，单链表无法倒着遍历，无法使用双指针技巧。

那么最简单的办法就是，把原始链表反转存入一条新的链表，然后比较这两条链表是否相同。

更聪明一些的办法是借助双指针算法：

**1、先通过 [双指针技巧](https://labuladong.github.io/article/fname.html?fname=链表技巧) 中的快慢指针来找到链表的中点**：

![](https://labuladong.github.io/pictures/回文链表/1.jpg)

**2、如果 `fast` 指针没有指向 `null`，说明链表长度为奇数，`slow` 还要再前进一步**：

![](https://labuladong.github.io/pictures/回文链表/2.jpg)

**3、从 `slow` 开始反转后面的链表，现在就可以开始比较回文串了**：

![](https://labuladong.github.io/pictures/回文链表/3.jpg)

**详细题解：[如何判断回文链表](https://labuladong.github.io/article/fname.html?fname=判断回文链表)**

**标签：回文问题，[数据结构](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318892385270808576)，[链表](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120596033251475465)，[链表双指针](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120596033251475465)**

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
    bool isPalindrome(ListNode* head) {
        ListNode *slow, *fast;   // Define two pointers
        slow = fast = head;      // Initialize the pointers
        while (fast != nullptr && fast->next != nullptr) {  // Loop until fast pointer reaches the end
            slow = slow->next;   // Move slow pointer one step
            fast = fast->next->next;   // Move fast pointer two step
        }

        if (fast != nullptr)    // If fast pointer is not nullptr
            slow = slow->next;  // Move the slow pointer one step

        ListNode *left = head;  // Initialize left pointer to head
        ListNode *right = reverse(slow);   // Reverse the right half of the list and make the right pointer point to the new head
        while (right != nullptr) {        // Loop until right pointer is nullptr
            if (left->val != right->val)  // If the values of left and right pointers are not equal
                return false;            // Return false
            left = left->next;           // Move the left pointer one step
            right = right->next;         // Move the right pointer one step
        }

        return true;   // Return true if all elements in the list are equal to their corresponding elements in the reversed list
    }

    ListNode* reverse(ListNode* head) {
        ListNode *pre = nullptr, *cur = head;   // Define two pointers
        while (cur != nullptr) {   // Loop until the last node of the list is reached
            ListNode *next = cur->next;    // Store the next pointer temporarily
            cur->next = pre;               // Reverse the pointer
            pre = cur;                     // Move the pre pointer one step
            cur = next;                    // Move the cur pointer one step forward
        }
        return pre;            // Return the new head
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        slow, fast = head, head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        
        if fast:
            slow = slow.next

        left = head
        right = self.reverse(slow)
        while right:
            if left.val != right.val:
                return False
            left = left.next
            right = right.next
            
        return True
    
    def reverse(self, head: ListNode) -> ListNode:
        pre, cur = None, head
        while cur:
            next_node = cur.next
            cur.next = pre
            pre = cur
            cur = next_node
            
        return pre
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public boolean isPalindrome(ListNode head) {
        ListNode slow, fast;
        slow = fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }

        if (fast != null)
            slow = slow.next;

        ListNode left = head;
        ListNode right = reverse(slow);
        while (right != null) {
            if (left.val != right.val)
                return false;
            left = left.next;
            right = right.next;
        }

        return true;
    }

    ListNode reverse(ListNode head) {
        ListNode pre = null, cur = head;
        while (cur != null) {
            ListNode next = cur.next;
            cur.next = pre;
            pre = cur;
            cur = next;
        }
        return pre;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func isPalindrome(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    if fast != nil {
        slow = slow.Next
    }

    left, right := head, reverse(slow)
    for right != nil {
        if left.Val != right.Val {
            return false
        }
        left = left.Next
        right = right.Next
    }

    return true
}

func reverse(head *ListNode) *ListNode {
    var pre, cur *ListNode
    cur = head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var isPalindrome = function(head) {
  let [slow, fast] = [head, head];
  while (fast !== null && fast.next !== null) {
    slow = slow.next;
    fast = fast.next.next;
  }

  if (fast !== null)
    slow = slow.next;

  let left = head;
  let right = reverse(slow);
  while (right !== null) {
    if (left.val !== right.val)
      return false;
    left = left.next;
    right = right.next;
  }

  return true;
};

function reverse(head) {
  let [pre, cur] = [null, head];
  while (cur !== null) {
    let next = cur.next;
    cur.next = pre;
    pre = cur;
    cur = next;
  }
  return pre;
}
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🌟🌟 算法可视化 🌟🌟</strong></summary><div id="data_palindrome-linked-list" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_palindrome-linked-list"></div></div>
</details><hr /><br />

**类似题目**：
  - [剑指 Offer II 027. 回文链表 🟢](/problems/aMhZSa)

</details>
</div>

