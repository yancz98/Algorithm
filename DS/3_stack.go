package DS

/**
 * 栈
 *  - 用数组实现栈（顺序存储）
 *  - 用链表实现栈（链式存储）
 *  - 用两个队列实现栈
 */

// StackWithArray - 用数组实现栈（顺序存储）
type StackWithArray []any

// Empty 判断是否为空
func (this *StackWithArray) Empty() bool {
	return len(*this) == 0
}

// Push 入栈
func (this *StackWithArray) Push(v any) {
	*this = append(*this, v)
}

// Pop 出栈
func (this *StackWithArray) Pop() any {
	if this.Empty() {
		return nil
	}

	v := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]

	return v
}

// Top 读栈顶元素
func (this *StackWithArray) Top() any {
	if this.Empty() {
		return nil
	}

	return (*this)[len(*this)-1]
}

// StackWithDLinked - 用链表实现栈（顺序存储）
type StackWithDLinked struct {
	DLinked
}

// Empty 判断是否为空
func (this *StackWithDLinked) Empty() bool {
	return this.Head == nil || this.Tail == nil
}

// Push 入栈
func (this *StackWithDLinked) Push(v int) {
	this.Unshift(v)
}

// Pop 出栈
func (this *StackWithDLinked) Pop() any {
	if this.Empty() {
		return nil
	}

	return this.Shift().Value
}

func (this *StackWithDLinked) Top() any {
	if this.Empty() {
		return nil
	}

	return this.Head.Value
}

// StackWithTwoQueues - 用两个队列实现栈
//
//  - 入栈时，将数据写入非空的队列 (如果两个队列都为空，则随便写入一个)
//  - 出栈时，将非空队列的前 N-1 个元素导入到另一个队列，返回最后一个元素
type StackWithTwoQueues struct {
	q1 QueueWithDLinked // 主队
	q2 QueueWithDLinked // 辅队
}

// Empty 判断是否为空
func (this *StackWithTwoQueues) Empty() bool {
	return this.q1.Empty() && this.q2.Empty()
}

// Push 入栈
func (this *StackWithTwoQueues) Push(v int) {
	// 入栈时，将数据写入非空的队列
	if !this.q1.Empty() {
		this.q1.Enqueue(v)
	} else {
		this.q2.Enqueue(v)
	}
}

// Pop 出栈
func (this *StackWithTwoQueues) Pop() any {
	if this.Empty() {
		return nil
	}

	for !this.q1.Empty() {
		// 先出队
		v := this.q1.Dequeue()

		// 如果出队后，队列为空，说明这是最后一个元素
		if this.q1.Empty() {
			// 弹出最后一个元素
			return v
		}

		// 否则，丢到另一个队列
		this.q2.Enqueue(v)
	}

	for !this.q2.Empty() {
		// 先出队
		v := this.q2.Dequeue()

		// 如果出队后，队列为空，说明这是最后一个元素
		if this.q2.Empty() {
			// 弹出最后一个元素
			return v
		}

		// 否则，丢到另一个队列
		this.q1.Enqueue(v)
	}

	return nil
}
