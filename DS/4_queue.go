package DS

import (
	"errors"
)

/**
 * 队列
 *  - 用数组实现队列（顺序存储）
 *  - 用链表实现队列（链式存储）
 *  - 用两个栈实现队列
 */

// QueueWithArray - 用数组实现队列（顺序存储）
//
// 顺序队列无法重复利用空间，故设计为环形队列
type QueueWithArray struct {
	queue   []any // 环形数组
	bufSize uint  // 缓冲区大小
	qSize   uint  // 队列大小（实际元素）
	front   uint  // 队头（出队位置），初始为 0
	rear    uint  // 队尾（入队位置），初始为 0
}

// InitQueue 初始化队列
func InitQueue(size uint) *QueueWithArray {
	return &QueueWithArray{
		queue:   make([]any, size),
		bufSize: size,
		qSize:   0,
		front:   0,
		rear:    0,
	}
}

// Empty 判断是否为空
func (this *QueueWithArray) Empty() bool {
	return this.qSize == 0
}

// Enqueue 入队
func (this *QueueWithArray) Enqueue(v any) error {
	if this.qSize == this.bufSize {
		return errors.New("queue is full")
	}

	// 队尾入
	this.queue[this.rear] = v
	// 加库存，下一个入队位置
	this.qSize++
	this.rear++

	// 环形数组实现
	if this.rear == this.bufSize {
		this.rear = 0
	}

	return nil
}

// Dequeue 出队
func (this *QueueWithArray) Dequeue() any {
	if this.Empty() {
		return nil
	}

	// 队头出
	v := this.queue[this.front]
	// 减库存，下一个出队位置
	this.qSize--
	this.front++

	// 环形数组实现
	if this.front == this.bufSize {
		this.front = 0
	}

	return v
}

// Front 读队头元素（不移除）
func (this *QueueWithArray) Front() any {
	if this.Empty() {
		return nil
	}

	return this.queue[this.front]
}

// QueueWithDLinked - 用链表实现队列（链式存储）
type QueueWithDLinked struct {
	DLinked
}

// Empty 判断是否为空
func (this *QueueWithDLinked) Empty() bool {
	return this.Head == nil || this.Tail == nil
}

// Enqueue 入队
func (this *QueueWithDLinked) Enqueue(v any) {
	this.Unshift(v)
}

// Dequeue 出队
func (this *QueueWithDLinked) Dequeue() any {
	if this.Empty() {
		return nil
	}

	return this.Pop().Value
}

// Front 读队头元素（不移除）
func (this *QueueWithDLinked) Front() any {
	if this.Empty() {
		return nil
	}

	return this.Head.Value
}

// QueueWithTwoStacks - 用两个栈实现队列
//
//  - 入队时：只往 in 栈中插
//  - 出队时：从 out 栈中取，若 out 为空，则把 in 中的数据全部导入到 out 后再取
type QueueWithTwoStacks struct {
	ins  StackWithArray // 入栈
	outs StackWithArray // 出栈
}

// Empty 判断是否为空
func (this *QueueWithTwoStacks) Empty() bool {
	return this.ins.Empty() && this.outs.Empty()
}

// Enqueue 入队
func (this *QueueWithTwoStacks) Enqueue(v any) {
	// 入队只往 in 栈中插
	this.ins.Push(v)
}

// Dequeue 出队
func (this *QueueWithTwoStacks) Dequeue() any {
	if this.Empty() {
		return nil
	}

	// 出栈时，先把 out 栈中的元素拿空
	// 再把 in 栈中的数据全部导入到 out 栈
	if this.outs.Empty() {
		for !this.ins.Empty() {
			this.outs.Push(this.ins.Pop())
		}
	}

	// 从 out 栈中弹出
	return this.outs.Pop()
}
