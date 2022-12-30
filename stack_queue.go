package Algorithm

/**
 * 栈 & 队列
 *  - 用双向链表实现栈
 *  - 用数组实现栈
 *  - 用双向链表实现队列
 *  - 用数组实现队列
 *  - 用两个栈实现队列
 *  - 用两个队列实现栈
 */

////////////////////////////////////////////////////
// 栈
////////////////////////////////////////////////////

// 1、用双向链表实现栈
type StackWithList struct {
    stack List
}

// 入栈
func (this *StackWithList) Push(v int) {
    this.stack.FrontPush(v)
}

// 出栈
func (this *StackWithList) Pop() int {
    if this.Empty() {
        panic("Stack Empty...")
    }

    return this.stack.FrontPop().Value
}

// 判断是否为空
func (this *StackWithList) Empty() bool {
    if this.stack.Head == nil || this.stack.Tail == nil {
        return true
    }

    return false
}

// 2、用数组实现栈
type StackWithArray struct {
    stack []int
}

// 入栈
func (this *StackWithArray) Push(v int) {
    this.stack = append(this.stack, v)
}

// 出栈
func (this *StackWithArray) Pop() int {
    if this.Empty() {
        panic("Stack Empty...")
    }

    v := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]

    return v
}

// 判断是否为空
func (this *StackWithArray) Empty() bool {
    if len(this.stack) == 0 {
        return true
    }

    return false
}

////////////////////////////////////////////////////
// 队列
////////////////////////////////////////////////////

// 3、用双向链表实现队列
type QueueWithList struct {
    queue List
}

// 入队
func (this *QueueWithList) Enqueue(v int) {
    this.queue.FrontPush(v)
}

// 出队
func (this *QueueWithList) Dequeue() int {
    if this.Empty() {
        panic("Queue Empty...")
    }

    return this.queue.BackPop().Value
}

// 判断是否为空
func (this *QueueWithList) Empty() bool {
    if this.queue.Head == nil || this.queue.Tail == nil {
        return true
    }

    return false
}

// 4、用数组实现队列
type QueueWithArray struct {
    queue   []int // 环形数组
    qSize   uint  // 队列大小
    bufSize uint  // 缓冲区大小
    enIdx   uint  // 入队位置，初始为 0
    deIdx   uint  // 出队位置，初始为 0
}

// 新建队列
func MakeQueue(size uint) QueueWithArray {
    return QueueWithArray{
        queue:   make([]int, size),
        qSize:   0,
        bufSize: size,
        enIdx:   0,
        deIdx:   0,
    }
}

// 入队
func (this *QueueWithArray) Enqueue(v int) {
    if this.qSize == this.bufSize {
        panic("Queue Full...")
    }

    // 存 入队位置
    this.queue[this.enIdx] = v
    // 加库存，下一个入队位置
    this.qSize++
    this.enIdx++
    // 环形数组实现
    if this.enIdx == this.bufSize {
        this.enIdx = 0
    }
}

// 出队
func (this *QueueWithArray) Dequeue() int {
    if this.Empty() {
        panic("Queue Empty...")
    }

    // 取 出队位置
    v := this.queue[this.deIdx]
    // 减库存，下一个出队位置
    this.qSize--
    this.deIdx++
    // 环形数组实现
    if this.deIdx == this.bufSize {
        this.deIdx = 0
    }

    return v
}

// 判断是否为空
func (this *QueueWithArray) Empty() bool {
    if this.qSize == 0 {
        return true
    }

    return false
}

////////////////////////////////////////////////////

// 5、用两个队列实现栈
type StackWithQueue struct {
    q1 QueueWithList // 主队
    q2 QueueWithList // 辅队
}

// 入栈
func (this *StackWithQueue) Push(v int) {
    // 入栈时，将数据写入非空的队列
    if !this.q1.Empty() {
        this.q1.Enqueue(v)
    } else {
        this.q2.Enqueue(v)
    }
}

// 出栈
func (this *StackWithQueue) Pop() int {
    if this.Empty() {
        panic("Stack Empty...")
    }

    // 出栈时，将非空队列的前 N-1 个元素导入到另一个队列
    // 弹出最后一个元素
    var last int
    if !this.q1.Empty() {
        for !this.q1.Empty() {
            // 先出队
            v := this.q1.Dequeue()
            // 如果出队后，队列为空，说明这是最后一个元素
            if this.q1.Empty() {
                last = v
                break
            }

            // 否则，丢到另一个队列
            this.q2.Enqueue(v)
        }
    } else {
        for !this.q2.Empty() {
            // 先出队
            v := this.q2.Dequeue()
            // 如果出队后，队列为空，说明这是最后一个元素
            if this.q2.Empty() {
                last = v
                break
            }

            // 否则，丢到另一个队列
            this.q1.Enqueue(v)
        }
    }

    return last
}

// 判断是否为空
func (this *StackWithQueue) Empty() bool {
    if this.q1.Empty() && this.q2.Empty() {
        return true
    }

    return false
}

// 6、用两个栈实现队列
type QueueWithStack struct {
    ins  StackWithArray // 入栈
    outs StackWithArray // 出栈
}

// 入队
func (this *QueueWithStack) Enqueue(v int) {
    // 入队只往 in 栈中插
    this.ins.Push(v)
}

// 出队
func (this *QueueWithStack) Dequeue() int {
    if this.Empty() {
        panic("Queue Empty...")
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

// 判断是否为空
func (this *QueueWithStack) Empty() bool {
    if this.ins.Empty() && this.outs.Empty() {
        return true
    }

    return false
}
