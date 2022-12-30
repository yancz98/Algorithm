package Algorithm

import (
    "fmt"
    "testing"
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
func Test_reverseSingleList(t *testing.T) {
    node4 := &SingleNode{Value: 1, Next: nil}
    node3 := &SingleNode{Value: 1, Next: node4}
    node2 := &SingleNode{Value: 1, Next: node3}
    node1 := &SingleNode{Value: 1, Next: node2}

    fmt.Println("Original:")
    for cur := node1; cur != nil; cur = cur.Next {
        fmt.Println(cur.Value)
    }

    // 反转
    /*reversed := reverseSingleList(node1)
      fmt.Println("Reversed:")
      for cur := reversed; cur != nil; cur = cur.Next {
          fmt.Println(cur.Value)
      }*/

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
    stack := StackWithList{}

    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    stack.Push(4)

    for !stack.Empty() {
        fmt.Println(stack.Pop())
    }
    //fmt.Println(stack.Pop())

    stack2 := StackWithArray{}

    stack2.Push(1)
    stack2.Push(2)
    stack2.Push(3)
    stack2.Push(4)

    for !stack2.Empty() {
        fmt.Println(stack2.Pop())
    }
    //fmt.Println(stack2.Pop())
}

func TestQueue(t *testing.T) {
    queue := QueueWithList{}

    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)

    for !queue.Empty() {
        fmt.Println(queue.Dequeue())
    }
    //fmt.Println(queue.Dequeue())

    queue2 := MakeQueue(5)

    queue2.Enqueue(5)
    queue2.Enqueue(6)
    queue2.Enqueue(7)
    queue2.Enqueue(8)
    queue2.Enqueue(9)
    //queue2.Enqueue(10)

    for !queue2.Empty() {
        fmt.Println(queue2.Dequeue())
    }
    //fmt.Println(queue2.Dequeue())
}
