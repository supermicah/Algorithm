package main

import "fmt"

/***
*
* 2023-05-07 11:00:54
*
***/

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的
//功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
//
//
// 示例 1：
//
//
//输入：
//["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
//[[],[3],[],[],[]]
//输出：[null,null,3,-1,-1]
//
//
// 示例 2：
//
//
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//
//
// 提示：
//
//
// 1 <= values <= 10000
// 最多会对 appendTail、deleteHead 进行 10000 次调用
//
//
// Related Topics栈 | 设计 | 队列
//
// 👍 739, 👎 0bug 反馈 | 使用指南 | 更多配套插件
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
type CQueue struct {
	in, out []int
}

func Constructor() CQueue {
	return CQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		}
		this.move()
	}
	v := this.out[len(this.out)-1]
	this.out = this.out[0 : len(this.out)-1]
	return v
}
func (this *CQueue) move() {
	for len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	obj.DeleteHead()
	obj.DeleteHead()
	obj.DeleteHead()
	obj.AppendTail(2)
	obj.AppendTail(3)
	obj.DeleteHead()
	obj.DeleteHead()
	obj.AppendTail(4)
	obj.DeleteHead()
	obj.DeleteHead()
	value := obj.DeleteHead()
	fmt.Println(fmt.Sprintf("%+v", value))
}
