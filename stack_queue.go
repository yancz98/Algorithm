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

func (this *StackWithArray) Push(v int) {
    this.stack = append(this.stack, v)
}

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
    mainQueue      QueueWithList // 主队
    secondaryQueue QueueWithList // 辅队
}

func (this *StackWithQueue) Push(v int) {
    // 入栈时，将数据写入 inQueue
    //this.inQueue.Enqueue(v)
}

func (this *StackWithQueue) Pop() int {
    // 出栈时，如果 outStack

    return 0
}

// 6、用两个栈实现队列
