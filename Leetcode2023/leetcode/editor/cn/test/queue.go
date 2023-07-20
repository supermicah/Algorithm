package main

import (
	"container/heap"
	"fmt"
	"time"
)

// DelayedTask 定义延迟任务的结构体
type DelayedTask struct {
	message    string    // 任务消息
	expiration time.Time // 任务到期时间
}

// 实现堆接口所需的方法

// GetExpiration 获取任务到期时间
func (t *DelayedTask) GetExpiration() time.Time {
	return t.expiration
}

// Less 比较任务的优先级
func (t *DelayedTask) Less(otherTask DelayedTask) bool {
	return t.expiration.Before(otherTask.expiration)
}

// GetMessage 获取任务消息
func (t *DelayedTask) GetMessage() string {
	return t.message
}

func (t *DelayedTask) String() {
	fmt.Println(t.message)
}

// DelayedQueue 定义延迟队列类型
type DelayedQueue []DelayedTask

// 实现堆接口所需的方法

// Len 获取队列长度
func (q DelayedQueue) Len() int {
	return len(q)
}

// Less 判断队列中的两个任务的优先级
func (q DelayedQueue) Less(i, j int) bool {
	return q[i].expiration.Before(q[j].expiration)
}

// Swap 交换队列中的两个任务位置
func (q DelayedQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// Push 添加任务到队列
func (q *DelayedQueue) Push(x interface{}) {
	*q = append(*q, x.(DelayedTask))
}

// Pop 移除队列中最早到期的任务
func (q *DelayedQueue) Pop() interface{} {
	old := *q
	n := len(old)
	task := old[n-1]
	*q = old[0 : n-1]
	return task
}
func (q *DelayedQueue) String() {
	old := *q
	for i := 0; i < len(old); i++ {
		fmt.Println(old[i])
	}
}

// 启动延迟队列
func startDelayedQueue() {
	// 创建一个空的延迟队列
	queue := &DelayedQueue{}

	// 添加一些延迟任务到队列
	addDelayedTask(queue, "Task 1", 5*time.Second)
	fmt.Println(fmt.Sprintf("%+v", queue))
	addDelayedTask(queue, "Task 2", 1*time.Second)
	fmt.Println(fmt.Sprintf("%+v", queue))
	addDelayedTask(queue, "Task 3", 8*time.Second)
	fmt.Println(fmt.Sprintf("%+v", queue))
	addDelayedTask(queue, "Task 4", 1*time.Second)
	fmt.Println(fmt.Sprintf("%+v", queue))

	// 处理延迟队列中的任务
	processDelayedQueue(queue)
}

// 添加延迟任务到队列
func addDelayedTask(queue *DelayedQueue, message string, delay time.Duration) {
	task := DelayedTask{
		message:    message,
		expiration: time.Now().Add(delay),
	}
	heap.Push(queue, task)
}

// 处理延迟队列中的任务
func processDelayedQueue(queue *DelayedQueue) {
	for queue.Len() > 0 {
		// 获取最早到期的任务
		task := heap.Pop(queue).(DelayedTask)
		now := time.Now()

		// 判断任务是否已到期
		if now.Before(task.GetExpiration()) {
			// 如果任务未到期，则等待到期时间后重新放入队列
			sleepDuration := task.GetExpiration().Sub(now)
			time.Sleep(sleepDuration)
			heap.Push(queue, task)
		} else {
			// 处理到期的任务
			fmt.Println("Task:", task.GetMessage(), "is executed at", now)
		}
	}
}

func main() {
	startDelayedQueue()
}
