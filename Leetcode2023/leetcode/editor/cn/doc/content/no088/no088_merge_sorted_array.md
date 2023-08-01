<p>给你两个按 <strong>非递减顺序</strong> 排列的整数数组&nbsp;<code>nums1</code><em> </em>和 <code>nums2</code>，另有两个整数 <code>m</code> 和 <code>n</code> ，分别表示 <code>nums1</code> 和 <code>nums2</code> 中的元素数目。</p>

<p>请你 <strong>合并</strong> <code>nums2</code><em> </em>到 <code>nums1</code> 中，使合并后的数组同样按 <strong>非递减顺序</strong> 排列。</p>

<p><strong>注意：</strong>最终，合并后数组不应由函数返回，而是存储在数组 <code>nums1</code> 中。为了应对这种情况，<code>nums1</code> 的初始长度为 <code>m + n</code>，其中前 <code>m</code> 个元素表示应合并的元素，后 <code>n</code> 个元素为 <code>0</code> ，应忽略。<code>nums2</code> 的长度为 <code>n</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
<strong>输出：</strong>[1,2,2,3,5,6]
<strong>解释：</strong>需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [<em><strong>1</strong></em>,<em><strong>2</strong></em>,2,<em><strong>3</strong></em>,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [1], m = 1, nums2 = [], n = 0
<strong>输出：</strong>[1]
<strong>解释：</strong>需要合并 [1] 和 [] 。
合并结果是 [1] 。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [0], m = 0, nums2 = [1], n = 1
<strong>输出：</strong>[1]
<strong>解释：</strong>需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>nums1.length == m + n</code></li> 
 <li><code>nums2.length == n</code></li> 
 <li><code>0 &lt;= m, n &lt;= 200</code></li> 
 <li><code>1 &lt;= m + n &lt;= 200</code></li> 
 <li><code>-10<sup>9</sup> &lt;= nums1[i], nums2[j] &lt;= 10<sup>9</sup></code></li> 
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>你可以设计实现一个时间复杂度为 <code>O(m + n)</code> 的算法解决此问题吗？</p>

<details><summary><strong>Related Topics</strong></summary>数组 | 双指针 | 排序</details><br>

<div>👍 1947, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员！[第 21 期打卡挑战](https://opedk.xet.tech/s/4ptSo2) 最后一天报名！**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题很像前文 [链表的双指针技巧汇总](https://labuladong.github.io/article/fname.html?fname=链表技巧) 中讲过的 [21. 合并两个有序链表](/problems/merge-two-sorted-lists)，这里让你合并两个有序数组。

对于单链表来说，我们直接用双指针从头开始合并即可，但对于数组来说会出问题。因为题目让我直接把结果存到 `nums1` 中，而 `nums1` 的开头有元素，如果我们无脑复制单链表的逻辑，会覆盖掉 `nums1` 的原始元素，导致错误。

但 `nums1` 后面是空的呀，所以这道题需要我们稍微变通一下：**将双指针初始化在数组的尾部，然后从后向前进行合并**，这样即便覆盖了 `nums1` 中的元素，这些元素也必然早就被用过了，不会影响答案的正确性。

**标签：[数据结构](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318892385270808576)，[数组双指针](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)**

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
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        // 两个指针分别初始化在两个数组的最后一个元素（类似拉链两端的锯齿）
        int i = m - 1, j = n - 1;
        // 生成排序的结果（类似拉链的拉锁）
        int p = nums1.size() - 1;
        // 从后向前生成结果数组，类似合并两个有序链表的逻辑
        while (i >= 0 && j >= 0) {
            if (nums1[i] > nums2[j]) {
                nums1[p] = nums1[i];
                i--;
            } else {
                nums1[p] = nums2[j];
                j--;
            }
            p--;
        }
        // 可能其中一个数组的指针走到尽头了，而另一个还没走完
        // 因为我们本身就是在往 nums1 中放元素，所以只需考虑 nums2 是否剩元素即可
        while (j >= 0) {
            nums1[p] = nums2[j];
            j--;
            p--;
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
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        # 两个指针分别初始化在两个数组的最后一个元素（类似拉链两端的锯齿）
        i, j = m - 1, n - 1
        # 生成排序的结果（类似拉链的拉锁）
        p = len(nums1) - 1
        # 从后向前生成结果数组，类似合并两个有序链表的逻辑
        while i >= 0 and j >= 0:
            if nums1[i] > nums2[j]:
                nums1[p] = nums1[i]
                i -= 1
            else:
                nums1[p] = nums2[j]
                j -= 1
            p -= 1
        # 可能其中一个数组的指针走到尽头了，而另一个还没走完
        # 因为我们本身就是在往 nums1 中放元素，所以只需考虑 nums2 是否剩元素即可
        while j >= 0:
            nums1[p] = nums2[j]
            j -= 1
            p -= 1
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        // 两个指针分别初始化在两个数组的最后一个元素（类似拉链两端的锯齿）
        int i = m - 1, j = n - 1;
        // 生成排序的结果（类似拉链的拉锁）
        int p = nums1.length - 1;
        // 从后向前生成结果数组，类似合并两个有序链表的逻辑
        while (i >= 0 && j >= 0) {
            if (nums1[i] > nums2[j]) {
                nums1[p] = nums1[i];
                i--;
            } else {
                nums1[p] = nums2[j];
                j--;
            }
            p--;
        }
        // 可能其中一个数组的指针走到尽头了，而另一个还没走完
        // 因为我们本身就是在往 nums1 中放元素，所以只需考虑 nums2 是否剩元素即可
        while (j >= 0) {
            nums1[p] = nums2[j];
            j--;
            p--;
        }
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j, p := m-1, n-1, len(nums1)-1 // 初始化指针
    for i >= 0 && j >= 0 { // 两个数组都未遍历完时进行比较
        if nums1[i] > nums2[j] { // 挑选大的元素放入 nums1 的末位
            nums1[p] = nums1[i]
            i--
        } else {
            nums1[p] = nums2[j]
            j--
        }
        p-- // 从后往前生成结果
    }
    for j >= 0 { // nums2 剩余元素放入 nums1
        nums1[p] = nums2[j]
        j--
        p--
    }
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var merge = function(nums1, m, nums2, n) {
    // 两个指针分别初始化在两个数组的最后一个元素（类似拉链两端的锯齿）
    var i = m - 1, j = n - 1;
    // 生成排序的结果（类似拉链的拉锁）
    var p = nums1.length - 1;
    // 从后向前生成结果数组，类似合并两个有序链表的逻辑
    while (i >= 0 && j >= 0) {
        if (nums1[i] > nums2[j]) {
            nums1[p] = nums1[i];
            i--;
        } else {
            nums1[p] = nums2[j];
            j--;
        }
        p--;
    }
    // 可能其中一个数组的指针走到尽头了，而另一个还没走完
    // 因为我们本身就是在往 nums1 中放元素，所以只需考虑 nums2 是否剩元素即可
    while (j >= 0) {
        nums1[p] = nums2[j];
        j--;
        p--;
    }
};
```

</div></div>
</div></div>

<details open><summary><strong>🌟🌟 算法可视化 🌟🌟</strong></summary><div id="data_merge-sorted-array" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_merge-sorted-array"></div></div>
</details><hr /><br />

**类似题目**：
  - [977. 有序数组的平方 🟢](/problems/squares-of-a-sorted-array)

</details>
</div>

