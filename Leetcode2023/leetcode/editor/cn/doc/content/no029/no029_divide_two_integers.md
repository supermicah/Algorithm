<p>给你两个整数，被除数&nbsp;<code>dividend</code>&nbsp;和除数&nbsp;<code>divisor</code>。将两数相除，要求 <strong>不使用</strong> 乘法、除法和取余运算。</p>

<p>整数除法应该向零截断，也就是截去（<code>truncate</code>）其小数部分。例如，<code>8.345</code> 将被截断为 <code>8</code> ，<code>-2.7335</code> 将被截断至 <code>-2</code> 。</p>

<p>返回被除数&nbsp;<code>dividend</code>&nbsp;除以除数&nbsp;<code>divisor</code>&nbsp;得到的 <strong>商</strong> 。</p>

<p><strong>注意：</strong>假设我们的环境只能存储 <strong>32 位</strong> 有符号整数，其数值范围是 <code>[−2<sup>31</sup>,&nbsp; 2<sup>31&nbsp;</sup>− 1]</code> 。本题中，如果商 <strong>严格大于</strong> <code>2<sup>31&nbsp;</sup>− 1</code> ，则返回 <code>2<sup>31&nbsp;</sup>− 1</code> ；如果商 <strong>严格小于</strong> <code>-2<sup>31</sup></code> ，则返回 <code>-2<sup>31</sup></code><sup> </sup>。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre>
<strong>输入:</strong> dividend = 10, divisor = 3
<strong>输出:</strong> 3
<strong>解释: </strong>10/3 = 3.33333.. ，向零截断后得到 3 。</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre>
<strong>输入:</strong> dividend = 7, divisor = -3
<strong>输出:</strong> -2
<strong>解释:</strong> 7/-3 = -2.33333.. ，向零截断后得到 -2 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>-2<sup>31</sup> &lt;= dividend, divisor &lt;= 2<sup>31</sup> - 1</code></li> 
 <li><code>divisor != 0</code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>位运算 | 数学</details><br>

<div>👍 1192, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员，全新纸质书[《labuladong 的算法笔记》](https://labuladong.gitee.io/algo/images/book/book_intro_qrcode.jpg) 出版，签名版限时半价！**

</div>

