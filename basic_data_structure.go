package Algorithm

/**
 * 基本数据结构
 *  - 单向链表
 *  - 双向链表
 *     - 反转链表
 *     - 把给定值都删除
 *  - 栈 & 队列
 *     - 用双向链表实现栈
 *     - 用数组实现栈
 *     - 用双向链表实现队列
 *     - 用数组实现队列
 *     - 用两个栈实现队列
 *     - 用两个队列实现栈
 */

////////////////////////////////////////////////////
// 单向链表
////////////////////////////////////////////////////

// 节点
type SingleNode struct {
    Value int
    Next  *SingleNode
}

// 反转单链表
func reverseSingleList(head *SingleNode) *SingleNode {
    // 当前节点的前一个、后一个节点
    var prev, next *SingleNode

    for head != nil {
        // 记录下当前节点的 next
        next = head.Next
        // 指向前一个
        head.Next = prev
        // 准备进行下一轮
        prev = head
        head = next
    }

    // 迭代结束，此时 head、next 已全部指向 nil
    // prev 指向头
    return prev
}

// 把给定值都删除 - 单链表
func deleteGiveValueForSingleList(head *SingleNode, k int) *SingleNode {
    // 先确定新 head 的位置
    // 若前几个元素都是需要删除的，则跳到第一个不是 k 的位置
    for head != nil {
        if head.Value != k {
            break
        }

        head = head.Next
    }

    // prev 记录前一个非 k 节点
    prev := head
    // 上一轮中已确定 head 不是 k，所以直接从 head 的下一个开始判断
    cur := head

    for cur != nil {
        // 命中
        if cur.Value == k {
            // 跳过当前节点
            prev.Next = cur.Next
        } else {
            // 已命中时，当前节点即将删除，不能作为 prev，故不用更新
            // 未命中时，才将 prev 更新为当前节点，准备进行下一轮
            prev = cur
        }

        // 后移，进入下一轮
        cur = cur.Next
    }

    return head
}

////////////////////////////////////////////////////
// 双向链表
////////////////////////////////////////////////////

// 节点
type DoubleNode struct {
    Value int
    Prev  *DoubleNode
    Next  *DoubleNode
}

// 链表
type List struct {
    Head *DoubleNode
    Tail *DoubleNode
}

// 头插
func (this *List) FrontPush(v int) {
    node := &DoubleNode{Value: v}

    if this.Head == nil {
        this.Head = node
        this.Tail = node
    } else {
        // 向链表头插入，处理前、后指针
        this.Head.Prev = node
        node.Next = this.Head
        // 头指针前移
        this.Head = node
    }
}

// 尾插
func (this *List) BackPush(v int) {
    node := &DoubleNode{Value: v}

    if this.Tail == nil {
        this.Head = node
        this.Tail = node
    } else {
        // 向链表尾插入，处理前、后指针
        this.Tail.Next = node
        node.Prev = this.Tail
        // 尾指针后移
        this.Tail = node
    }
}

// 头出
func (this *List) FrontPop() *DoubleNode {
    if this.Head == nil {
        return nil
    }

    // 保存头结点
    node := this.Head
    if this.Head.Next == nil {
        // 该节点为最后一个节点
        // 将 Head、Tail 置 nil
        this.Head = nil
        this.Tail = nil
    } else {
        // 头结点后移
        this.Head = this.Head.Next
        // 新头结点的 Prev 置 nil，原头结点丢失
        this.Head.Prev = nil
    }

    return node
}

// 尾出
func (this *List) BackPop() *DoubleNode {
    if this.Tail == nil {
        return nil
    }

    // 保存尾节点
    node := this.Tail
    if this.Tail.Prev == nil {
        // 该节点为最后一个节点
        // 将 Head、Tail 置 nil
        this.Head = nil
        this.Tail = nil
    } else {
        // 尾结点前移
        this.Tail = this.Tail.Prev
        // 新尾结点的 Next 置 nil，原尾结点丢失
        this.Tail.Next = nil
    }

    return node
}

// 反转双链表
func reverseDoubleList(head *DoubleNode) *DoubleNode {
    var prev, next *DoubleNode

    for head != nil {
        // 保留下一个节点
        next = head.Next
        // 反转前、后指针
        head.Prev = next
        head.Next = prev
        // 准备下一轮
        prev = head
        head = next
    }

    return prev
}

// 把给定值都删除 - 双链表
func deleteGiveValueForDoubleList(head *DoubleNode, k int) *DoubleNode {
    return nil
}

////////////////////////////////////////////////////
// 栈 & 队列
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
    this.inQueue.Enqueue(v)
}

func (this *StackWithQueue) Pop() int {
    // 出栈时，如果 outStack
}

// 6、用两个栈实现队列
