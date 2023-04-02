<p>给你一个由小写字母组成的字符串&nbsp;<code>s</code>，和一个整数&nbsp;<code>k</code>。</p>

<p>请你按下面的要求分割字符串：</p>

<ul> 
 <li>首先，你可以将&nbsp;<code>s</code>&nbsp;中的部分字符修改为其他的小写英文字母。</li> 
 <li>接着，你需要把&nbsp;<code>s</code>&nbsp;分割成&nbsp;<code>k</code>&nbsp;个非空且不相交的子串，并且每个子串都是回文串。</li> 
</ul>

<p>请返回以这种方式分割字符串所需修改的最少字符数。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>s = "abc", k = 2
<strong>输出：</strong>1
<strong>解释：</strong>你可以把字符串分割成 "ab" 和 "c"，并修改 "ab" 中的 1 个字符，将它变成回文串。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>s = "aabbc", k = 3
<strong>输出：</strong>0
<strong>解释：</strong>你可以把字符串分割成 "aa"、"bb" 和 "c"，它们都是回文串。</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>s = "leetcode", k = 8
<strong>输出：</strong>0
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= k &lt;= s.length &lt;= 100</code></li> 
 <li><code>s</code>&nbsp;中只含有小写英文字母。</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>字符串 | 动态规划</details><br>

<div>👍 112, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://mp.weixin.qq.com/s/NF8mmVyXVfC1ehdMOsO7Cw' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。过年前最后一期打卡挑战 [开始报名](https://aep.xet.tech/s/1a9ByX)。**

</div>



