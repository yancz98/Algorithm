package recursion

import (
	"fmt"
	"testing"

	"algorithm/DS"
)

// 递归测试
func TestRecursion(t *testing.T) {
	// 阶乘
	fmt.Println(factorial(2))

	// 反转单链表
	node4 := &DS.SingleNode{Data: 4, Next: nil}
	node3 := &DS.SingleNode{Data: 3, Next: node4}
	node2 := &DS.SingleNode{Data: 2, Next: node3}
	node1 := &DS.SingleNode{Data: 1, Next: node2}

	fmt.Println("Original:")
	for cur := node1; cur != nil; cur = cur.Next {
		fmt.Println(cur.Data)
	}

	reversed := ReverseSingleListByRecursion(node1)
	fmt.Println("reverseSingleListByRecursion:")
	for cur := reversed; cur != nil; cur = cur.Next {
		fmt.Println(cur.Data)
	}
}
