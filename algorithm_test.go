package Algorithm

import (
    "fmt"
    "math/rand"
    "testing"
    "time"
)

// 位运算
func Test_swap(t *testing.T) {
    swap(5, 10)
}

func Test_onlyOneOdd(t *testing.T) {
    onlyOneOdd()
}

func Test_retainRightmost1(t *testing.T) {
    retainRightmost1(250)

    fmt.Println(-74 << 2)
}

func Test_twoOdd(t *testing.T) {
    twoOdd()
}

func Test_binary1Num(t *testing.T) {
    binary1Num()
}

// 比较排序
func TestSelectSort(t *testing.T) {
    arr := []int{9, 2, 5, 8, 3, 7, 6}
    fmt.Println("Unordered:", arr)
    SelectSort(arr)
    fmt.Println("Orderly:", arr)
}

func TestBubbleSort(t *testing.T) {
    arr := []int{9, 2, 5, 8, 3, 7, 6}
    fmt.Println("Unordered:", arr)
    BubbleSort(arr)
    fmt.Println("Orderly:", arr)
}

func TestInsertSort(t *testing.T) {
    arr := []int{9, 2, 5, 8, 3, 7, 6}
    fmt.Println("Unordered:", arr)
    InsertSort(arr)
    fmt.Println("Orderly:", arr)

    a := 10
    fmt.Println(a ^ 0)
}

// 基本数据结构
func TestSingleList(t *testing.T) {
    node4 := &SingleNode{Value: 4, Next: nil}
    node3 := &SingleNode{Value: 3, Next: node4}
    node2 := &SingleNode{Value: 2, Next: node3}
    node1 := &SingleNode{Value: 1, Next: node2}

    fmt.Println("Original:")
    for cur := node1; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }

    // 反转
    reversed := reverseSingleList(node1)
    fmt.Println("Reversed:")
    for cur := reversed; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }

    // 删除
    deleted := deleteGiveValueForSingleList(node1, 1)
    fmt.Println("Deleted:")
    for cur := deleted; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }
}

func Test_reverseDoubleList(t *testing.T) {
    node5 := &DoubleNode{Value: 5, Prev: nil, Next: nil}
    node4 := &DoubleNode{Value: 4, Prev: nil, Next: node5}
    node3 := &DoubleNode{Value: 3, Prev: nil, Next: node4}
    node2 := &DoubleNode{Value: 2, Prev: nil, Next: node3}
    node1 := &DoubleNode{Value: 1, Prev: nil, Next: node2}
    node2.Prev = node1
    node3.Prev = node2
    node4.Prev = node3
    node5.Prev = node4

    fmt.Println("Original:")
    for cur := node1; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }

    // 反转
    reversed := reverseDoubleList(node1)
    fmt.Println("Reversed:")
    for cur := reversed; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }

    // 删除

}

func TestList(t *testing.T) {
    list := List{}

    //list.FrontPush(1)
    //list.BackPush(2)
    //list.BackPush(3)
    //list.FrontPush(4)

    fmt.Println("head=", list.FrontPop())
    //fmt.Println("head=", list.FrontPop().Value)
    //fmt.Println("head=", list.FrontPop().Value)
    //fmt.Println("head=", list.FrontPop().Value)
    fmt.Println("tail=", list.BackPop())

    fmt.Println("从前往后：")
    for cur := list.Head; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }

    fmt.Println("从后往前：")
    for cur := list.Tail; cur != nil; cur = cur.Prev {
        fmt.Println(cur.Value)
    }
}

func TestStack(t *testing.T) {
    fmt.Println("StackWithList...")
    stack := StackWithList{}
    stack.Push(1)
    stack.Push(2)
    fmt.Println(stack.Pop())
    stack.Push(3)
    stack.Push(4)
    fmt.Println(stack.Pop())
    stack.Push(5)

    for !stack.Empty() {
        fmt.Println(stack.Pop())
    }
    //fmt.Println(stack.Pop())

    fmt.Println("StackWithArray...")
    stack2 := StackWithArray{}
    stack2.Push(1)
    stack2.Push(2)
    fmt.Println(stack2.Pop())
    stack2.Push(3)
    stack2.Push(4)
    fmt.Println(stack2.Pop())
    stack2.Push(5)

    for !stack2.Empty() {
        fmt.Println(stack2.Pop())
    }
    //fmt.Println(stack2.Pop())

    fmt.Println("StackWithQueue...")
    stack3 := StackWithQueue{}
    stack3.Push(1)
    stack3.Push(2)
    fmt.Println(stack3.Pop())
    stack3.Push(3)
    stack3.Push(4)
    fmt.Println(stack3.Pop())
    stack3.Push(5)

    for !stack3.Empty() {
        fmt.Println(stack3.Pop())
    }
    //fmt.Println(stack3.Pop())
}

func TestQueue(t *testing.T) {
    fmt.Println("QueueWithList...")
    queue := QueueWithList{}
    queue.Enqueue(1)
    queue.Enqueue(2)
    fmt.Println(queue.Dequeue())
    queue.Enqueue(3)
    for !queue.Empty() {
        fmt.Println(queue.Dequeue())
    }
    //fmt.Println(queue.Dequeue())

    fmt.Println("QueueWithArray...")
    queue2 := MakeQueue(5)
    queue2.Enqueue(5)
    queue2.Enqueue(6)
    queue2.Enqueue(7)
    fmt.Println(queue2.Dequeue())
    queue2.Enqueue(8)
    queue2.Enqueue(9)
    //queue2.Enqueue(10) // Full...
    for !queue2.Empty() {
        fmt.Println(queue2.Dequeue())
    }
    //fmt.Println(queue2.Dequeue()) // Empty...

    fmt.Println("QueueWithStack...")
    queue3 := QueueWithStack{}
    queue3.Enqueue(5)
    queue3.Enqueue(6)
    queue3.Enqueue(7)
    fmt.Println(queue3.Dequeue())
    queue3.Enqueue(8)
    queue3.Enqueue(9)
    for !queue3.Empty() {
        fmt.Println(queue3.Dequeue())
    }
    //fmt.Println(queue3.Dequeue()) // Empty...
}

// 递归测试
func TestRecursion(t *testing.T) {
    // 阶乘
    fmt.Println(factorial(2))

    // 反转单链表
    node4 := &SingleNode{Value: 4, Next: nil}
    node3 := &SingleNode{Value: 3, Next: node4}
    node2 := &SingleNode{Value: 2, Next: node3}
    node1 := &SingleNode{Value: 1, Next: node2}

    fmt.Println("Original:")
    for cur := node1; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }

    reversed := reverseSingleListByRecursion(node1)
    fmt.Println("reverseSingleListByRecursion:")
    for cur := reversed; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }
}

// 归并排序
func TestMergeSort(t *testing.T) {
    arr := []int{3, 1, 2, 4, 6, 8, 5, 7}
    MergeSort(arr)
    fmt.Println(arr)

    arr2 := []int{3, 1, 2, 6, 4, 8, 5, 7, 9}
    MergeSortByIterate(arr2)
    fmt.Println(arr2)
}

// 计算数组的小和
func TestSmallSum(t *testing.T) {
    fmt.Println(SmallSum([]int{1, 3, 4, 2, 5}))    // 16
    fmt.Println(SmallSum([]int{1, 3, 5, 2, 4, 6})) // 27
}

// 数组中的逆序对
func TestReversePair(t *testing.T) {
    fmt.Println(ReversePair([]int{7, 5, 6, 4}))             // 5
    fmt.Println(ReversePair([]int{1, 2, 3, 4, 5, 6, 7, 0})) // 7
    fmt.Println(ReversePair([]int{1, 2, 3, 4, 5, 6, 7}))    // 0
}

// 快速排序
func TestQuickSort(t *testing.T) {
    rand.Seed(time.Now().UnixNano())

    arr := rand.Perm(6)
    fmt.Println("arr:", arr)
    fmt.Println(partition(arr, 0, len(arr)-1))
    fmt.Println("after:", arr)

    arr = rand.Perm(7)
    fmt.Println("arr:", arr)
    fmt.Println(dutchFlag(arr, 0, len(arr)-1))
    fmt.Println("after:", arr)

    arr1 := rand.Perm(8)
    fmt.Println("arr1:", arr1)
    QuickSort1(arr1)
    fmt.Println("sort1:", arr1)

    arr2 := rand.Perm(9)
    fmt.Println("arr2:", arr2)
    QuickSort2(arr2)
    fmt.Println("sort2:", arr2)

    arr3 := rand.Perm(10)
    fmt.Println("arr3:", arr3)
    QuickSort3(arr3)
    fmt.Println("sort3:", arr3)
}

// 堆结构
func TestHeapStructure(t *testing.T) {
    // heapInsert
    var arr []int 
    arr = heapInsert(arr, 1)
    arr = heapInsert(arr, 9)
    arr = heapInsert(arr, 6)
    arr = heapInsert(arr, 8)
    arr = heapInsert(arr, 2)
    arr = heapInsert(arr, 3)
    arr = heapInsert(arr, 5)
    fmt.Println("heapInsert:", arr)

    // heapify
    arr = []int{1, 4, 7, 3, 0, 2, 5}
    fmt.Println("heapify before:", arr)
    for i := 0; i < len(arr); i++ {
        heapify(arr, i, len(arr))
    }
    fmt.Println("heapify after:", arr)

    // Push
    heap := Heap{}
    heap.Push(2)
    heap.Push(5)
    heap.Push(8)
    heap.Push(3)
    heap.Push(6)
    heap.Push(9)
    fmt.Println("Push End...", heap.heap)

    // Pop
    for !heap.Empty() {
        fmt.Println(heap.Pop())
    }

    // fmt.Println(heap.Pop())
     
    // 将数组初始化为堆
    rand.Seed(time.Now().UnixNano())
    arr = rand.Perm(8)
    fmt.Println("arr:", arr)
    h := Init(arr)
    fmt.Println("heap:", h.heap)
}

func TestHeapSort(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    arr := rand.Perm(8)
    fmt.Println("arr:", arr)
    HeapSort(arr)
    fmt.Println("heapSort:", arr)
}
