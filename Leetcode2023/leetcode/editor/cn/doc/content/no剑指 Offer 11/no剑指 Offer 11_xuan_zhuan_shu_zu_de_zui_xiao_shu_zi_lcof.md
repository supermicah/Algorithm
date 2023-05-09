<p>把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。</p>

<p>给你一个可能存在&nbsp;<strong>重复</strong>&nbsp;元素值的数组&nbsp;<code>numbers</code>&nbsp;，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的<strong>最小元素</strong>。例如，数组&nbsp;<code>[3,4,5,1,2]</code> 为 <code>[1,2,3,4,5]</code> 的一次旋转，该数组的最小值为 1。&nbsp;&nbsp;</p>

<p>注意，数组 <code>[a[0], a[1], a[2], ..., a[n-1]]</code> 旋转一次 的结果为数组 <code>[a[n-1], a[0], a[1], a[2], ..., a[n-2]]</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong><span><code>numbers = </code></span>[3,4,5,1,2]
<strong>输出：</strong>1
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong><span><code>numbers = </code></span>[2,2,2,0,1]
<strong>输出：</strong>0
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>n == numbers.length</code></li> 
 <li><code>1 &lt;= n &lt;= 5000</code></li> 
 <li><code>-5000 &lt;= numbers[i] &lt;= 5000</code></li> 
 <li><code>numbers</code> 原来是一个升序排序的数组，并进行了 <code>1</code> 至 <code>n</code> 次旋转</li> 
</ul>

<p>注意：本题与主站 154 题相同：<a href="https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/">https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/</a></p>

<details><summary><strong>Related Topics</strong></summary>数组 | 二分查找</details><br>

<div>👍 811, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

</div>



