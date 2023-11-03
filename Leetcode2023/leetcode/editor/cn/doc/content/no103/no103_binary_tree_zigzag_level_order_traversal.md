<p>给你二叉树的根节点 <code>root</code> ，返回其节点值的 <strong>锯齿形层序遍历</strong> 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" style="width: 277px; height: 302px;" /> 
<pre>
<strong>输入：</strong>root = [3,9,20,null,null,15,7]
<strong>输出：</strong>[[3],[20,9],[15,7]]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>root = [1]
<strong>输出：</strong>[[1]]
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>root = []
<strong>输出：</strong>[]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li>树中节点数目在范围 <code>[0, 2000]</code> 内</li> 
 <li><code>-100 &lt;= Node.val &lt;= 100</code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>树 | 广度优先搜索 | 二叉树</details><br>

<div>👍 827, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员，全新纸质书[《labuladong 的算法笔记》](https://labuladong.gitee.io/algo/images/book/book_intro_qrcode.jpg) 出版，签名版限时半价！**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这题和 [102. 二叉树的层序遍历](/problems/binary-tree-level-order-traversal) 几乎是一样的，只要用一个布尔变量 `flag` 控制遍历方向即可。

**标签：[BFS 算法](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122002916411604996)，[二叉树](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2121994699837177859)**

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
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> res;
        if (root == nullptr) {
            return res;
        }

        queue<TreeNode*> q;
        q.push(root);
        // 为 true 时向右，false 时向左
        bool flag = true;

        // while 循环控制从上向下一层层遍历
        while (!q.empty()) {
            int sz = q.size();
            // 记录这一层的节点值
            list<int> level;
            // for 循环控制每一层从左向右遍历
            for (int i = 0; i < sz; i++) {
                TreeNode* cur = q.front();
                q.pop();
                // 实现 z 字形遍历
                if (flag) {
                    level.push_back(cur->val);
                } else {
                    level.push_front(cur->val);
                }
                if (cur->left != nullptr)
                    q.push(cur->left);
                if (cur->right != nullptr)
                    q.push(cur->right);
            }
            // 切换方向
            flag = !flag;
            res.emplace_back(level.begin(), level.end());
        }
        return res;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        res = []
        if not root:
            return res

        q = collections.deque()
        q.append(root)
        # 为 True 时向右，False 时向左
        flag = True

        # while 循环控制从上向下一层层遍历
        while q:
            sz = len(q)
            # 记录这一层的节点值
            level = collections.deque()
            # for 循环控制每一层从左向右遍历
            for i in range(sz):
                cur = q.popleft()
                # 实现 z 字形遍历
                if flag:
                    level.append(cur.val)
                else:
                    level.appendleft(cur.val)
                if cur.left:
                    q.append(cur.left)
                if cur.right:
                    q.append(cur.right)
            # 切换方向
            flag = not flag
            res.append(list(level))
        return res
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> res = new LinkedList<>();
        if (root == null) {
            return res;
        }

        Queue<TreeNode> q = new LinkedList<>();
        q.offer(root);
        // 为 true 时向右，false 时向左
        boolean flag = true;

        // while 循环控制从上向下一层层遍历
        while (!q.isEmpty()) {
            int sz = q.size();
            // 记录这一层的节点值
            LinkedList<Integer> level = new LinkedList<>();
            // for 循环控制每一层从左向右遍历
            for (int i = 0; i < sz; i++) {
                TreeNode cur = q.poll();
                // 实现 z 字形遍历
                if (flag) {
                    level.addLast(cur.val);
                } else {
                    level.addFirst(cur.val);
                }
                if (cur.left != null)
                    q.offer(cur.left);
                if (cur.right != null)
                    q.offer(cur.right);
            }
            // 切换方向
            flag = !flag;
            res.add(level);
        }
        return res;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func zigzagLevelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }

    q := make([]*TreeNode, 0)
    q = append(q, root)
    // 为 true 时向右，false 时向左
    flag := true

    // while 循环控制从上向下一层层遍历
    for len(q) > 0 {
        sz := len(q)
        // 记录这一层的节点值
        level := make([]int, 0, sz)
        // for 循环控制每一层从左向右遍历
        for i := 0; i < sz; i++ {
            cur := q[0]
            q = q[1:]
            // 实现 z 字形遍历
            if flag {
                level = append(level, cur.Val)
            } else {
                level = append([]int{cur.Val}, level...)
            }
            if cur.Left != nil {
                q = append(q, cur.Left)
            }
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        }
        // 切换方向
        flag = !flag
        res = append(res, level)
    }
    return res
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var zigzagLevelOrder = function(root) {
    let res = [];
    if (root == null) {
        return res;
    }

    let q = [];
    q.push(root);
    // 为 true 时向右，false 时向左
    let flag = true;

    // while 循环控制从上向下一层层遍历
    while (q.length > 0) {
        let sz = q.length;
        // 记录这一层的节点值
        let level = [];
        // for 循环控制每一层从左向右遍历
        for (let i = 0; i < sz; i++) {
            let cur = q.shift();
            // 实现 z 字形遍历
            if (flag) {
                level.push(cur.val);
            } else {
                level.unshift(cur.val);
            }
            if (cur.left != null)
                q.push(cur.left);
            if (cur.right != null)
                q.push(cur.right);
        }
        // 切换方向
        flag = !flag;
        res.push(level);
    }
    return res;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🥳🥳 算法可视化 🥳🥳</strong></summary><div id="data_binary-tree-zigzag-level-order-traversal" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_binary-tree-zigzag-level-order-traversal"></div></div>
</details><hr /><br />

**类似题目**：
  - [1609. 奇偶树 🟠](/problems/even-odd-tree)
  - [剑指 Offer 32 - III. 从上到下打印二叉树 III 🟠](/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof)

</details>
</div>

