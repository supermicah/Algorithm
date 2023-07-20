<p>请实现一个函数用来判断字符串是否表示<strong>数值</strong>（包括整数和小数）。</p>

<p><strong>数值</strong>（按顺序）可以分成以下几个部分：</p>

<ol> 
 <li>若干空格</li> 
 <li>一个&nbsp;<strong>小数</strong>&nbsp;或者&nbsp;<strong>整数</strong></li> 
 <li>（可选）一个&nbsp;<code>'e'</code>&nbsp;或&nbsp;<code>'E'</code>&nbsp;，后面跟着一个&nbsp;<strong>整数</strong></li> 
 <li>若干空格</li> 
</ol>

<p><strong>小数</strong>（按顺序）可以分成以下几个部分：</p>

<ol> 
 <li>（可选）一个符号字符（<code>'+'</code> 或 <code>'-'</code>）</li> 
 <li>下述格式之一： 
  <ol> 
   <li>至少一位数字，后面跟着一个点 <code>'.'</code></li> 
   <li>至少一位数字，后面跟着一个点 <code>'.'</code> ，后面再跟着至少一位数字</li> 
   <li>一个点 <code>'.'</code> ，后面跟着至少一位数字</li> 
  </ol> </li> 
</ol>

<p><strong>整数</strong>（按顺序）可以分成以下几个部分：</p>

<ol> 
 <li>（可选）一个符号字符（<code>'+'</code> 或 <code>'-'</code>）</li> 
 <li>至少一位数字</li> 
</ol>

<p>部分<strong>数值</strong>列举如下：</p>

<ul> 
 <li><code>["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]</code></li> 
</ul>

<p>部分<strong>非数值</strong>列举如下：</p>

<ul> 
 <li><code>["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]</code></li> 
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "0"
<strong>输出：</strong>true
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "e"
<strong>输出：</strong>false
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "."
<strong>输出：</strong>false</pre>

<p><strong>示例 4：</strong></p>

<pre>
<strong>输入：</strong>s = "&nbsp;&nbsp;&nbsp;&nbsp;.1&nbsp;&nbsp;"
<strong>输出：</strong>true
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= s.length &lt;= 20</code></li> 
 <li><code>s</code> 仅含英文字母（大写和小写），数字（<code>0-9</code>），加号 <code>'+'</code> ，减号 <code>'-'</code> ，空格 <code>' '</code> 或者点 <code>'.'</code> 。</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>字符串</details><br>

<div>👍 520, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员！[第 21 期打卡挑战](https://opedk.xet.tech/s/4ptSo2) 开始报名！**

</div>



