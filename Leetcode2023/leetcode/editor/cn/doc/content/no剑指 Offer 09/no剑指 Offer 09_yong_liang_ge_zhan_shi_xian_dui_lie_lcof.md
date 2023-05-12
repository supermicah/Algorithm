<p>用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 <code>appendTail</code> 和 <code>deleteHead</code> ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，<code>deleteHead</code>&nbsp;操作返回 -1 )</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>
["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
[[],[3],[],[],[]]
<strong>输出：</strong>[null,null,3,-1,-1]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
<strong>输出：</strong>[null,-1,null,null,5,2]
</pre>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= values &lt;= 10000</code></li> 
 <li>最多会对<code>&nbsp;appendTail、deleteHead </code>进行<code>&nbsp;10000</code>&nbsp;次调用</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>栈 | 设计 | 队列</details><br>

<div>👍 739, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**



<p><strong><a href="https://labuladong.github.io/article/slug.html?slug=yong-liang-ge-zhan-shi-xian-dui-lie-lcof" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题和 [232. 用栈实现队列](/problems/implement-queue-using-stacks) 相同。

我们使用两个栈 `s1, s2` 就能实现一个队列的功能。

当调用 `push` 让元素入队时，只要把元素压入 `s1` 即可：

![](https://labuladong.github.io/pictures/栈队列/3.jpg)

使用 `peek` 或 `pop` 操作队头的元素时，若 `s2` 为空，可以把 `s1` 的所有元素取出再添加进 `s2`，**这时候 `s2` 中元素就是先进先出顺序了**：

![](https://labuladong.github.io/pictures/栈队列/4.jpg)

**详细题解：[队列实现栈以及栈实现队列](https://labuladong.github.io/article/fname.html?fname=队列实现栈栈实现队列)**

**标签：[数据结构](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318892385270808576)，[栈](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2121993002939219969)，[队列](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2121993002939219969)**

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
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

class MyQueue {
private:
    stack<int> s1, s2;
public:
    MyQueue() {
        s1 = stack<int>();
        s2 = stack<int>();
    }

    /**
     * 添加元素到队尾
     */
    void push(int x) {
        s1.push(x);
    }

    /** 
     * 删除队头的元素并返回
     */
    int pop() {
        // 先调用 peek 保证 s2 非空
        peek();
        int val = s2.top();
        s2.pop();
        return val;
    }

    /** 
     * 返回队头元素
     */
    int peek() {
        if (s2.empty())
            // 把 s1 元素压入 s2
            while (!s1.empty()) {
                int val = s1.top();
                s1.pop();
                s2.push(val);
            }
        return s2.top();
    }

    /**
     * 判断队列是否为空
     */
    bool empty() {
        return s1.empty() && s2.empty();
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

class MyQueue:
    def __init__(self):
        """
        初始化一个队列，使用两个堆栈 s1 和 s2
        """
        self.s1 = []
        self.s2 = []
    
    def push(self, x: int) -> None:
        """
        添加元素到队尾
        """
        self.s1.append(x)
    
    def pop(self) -> int:
        """
        删除队头的元素并返回
        """
        # 先调用 peek 保证 s2 非空
        self.peek()
        return self.s2.pop()
    
    def peek(self) -> int:
        """
        返回队头元素
        """
        if not self.s2:
            # 把 s1 元素压入 s2
            while self.s1:
                self.s2.append(self.s1.pop())
        return self.s2[-1]
    
    def empty(self) -> bool:
        """
        判断队列是否为空
        """
        return not self.s1 and not self.s2
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class MyQueue {
    private Stack<Integer> s1, s2;

    public MyQueue() {
        s1 = new Stack<>();
        s2 = new Stack<>();
    }

    /**
     * 添加元素到队尾
     */
    public void push(int x) {
        s1.push(x);
    }

    /**
     * 删除队头的元素并返回
     */
    public int pop() {
        // 先调用 peek 保证 s2 非空
        peek();
        return s2.pop();
    }

    /**
     * 返回队头元素
     */
    public int peek() {
        if (s2.isEmpty())
            // 把 s1 元素压入 s2
            while (!s1.isEmpty())
                s2.push(s1.pop());
        return s2.peek();
    }

    /**
     * 判断队列是否为空
     */
    public boolean empty() {
        return s1.isEmpty() && s2.isEmpty();
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

type MyQueue struct {
    s1, s2 *stack.Stack
}

func Constructor() MyQueue {
    return MyQueue{
        s1: stack.New(),
        s2: stack.New(),
    }
}

/**
 * 添加元素到队尾
 */
func (this *MyQueue) Push(x int) {
    this.s1.Push(x)
}

/**
 * 删除队头的元素并返回
 */
func (this *MyQueue) Pop() int {
    // 先调用 Peek 保证 s2 非空
    this.Peek()
    return this.s2.Pop().(int)
}

/**
 * 返回队头元素
 */
func (this *MyQueue) Peek() int {
    if this.s2.Empty() {
        // 把 s1 元素压入 s2
        for !this.s1.Empty() {
            this.s2.Push(this.s1.Pop())
        }
    }
    return this.s2.Peek().(int)
}

/**
 * 判断队列是否为空
 */
func (this *MyQueue) Empty() bool {
    return this.s1.Empty() && this.s2.Empty()
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

/**
 * Initialize your data structure here.
 */
var MyQueue = function() {
    this.s1 = [];
    this.s2 = [];
};

/**
 * Push element x to the back of queue. 
 * @param {number} x
 * @return {void}
 */
MyQueue.prototype.push = function(x) {
    this.s1.push(x);
};

/**
 * Removes the element from in front of queue and returns that element.
 * @return {number}
 */
MyQueue.prototype.pop = function() {
    // 先调用 peek 保证 s2 非空
    this.peek();
    return this.s2.pop();
};

/**
 * Get the front element.
 * @return {number}
 */
MyQueue.prototype.peek = function() {
    if (this.s2.length === 0) {
        // 把 s1 元素压入 s2
        while (this.s1.length !== 0) {
            this.s2.push(this.s1.pop());
        }
    }
    return this.s2[this.s2.length - 1];
};

/**
 * Returns whether the queue is empty.
 * @return {boolean}
 */
MyQueue.prototype.empty = function() {
    return this.s1.length === 0 && this.s2.length === 0;
};
```

</div></div>
</div></div>

**类似题目**：
  - [225. 用队列实现栈 🟢](/problems/implement-stack-using-queues)
  - [剑指 Offer 09. 用两个栈实现队列 🟢](/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof)

</details>
</div>



