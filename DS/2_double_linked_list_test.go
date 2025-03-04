package DS

import (
	"fmt"
	"testing"
)

func Test_DoubleLinked(t *testing.T) {
	dl := new(DLinked)
	dl.Unshift(3)
	dl.Unshift(2)
	dl.Unshift(1)
	dl.Push(4)
	dl.Push(5)
	dl.Push(6)

	t.Log("head=", dl.Shift().Value)
	t.Log("tail=", dl.Pop().Value)
	t.Log("head=", dl.Shift().Value)
	t.Log("tail=", dl.Pop().Value)

	dl.Unshift(2)
	dl.Push(5)

	t.Log("从前往后：")
	for cur := dl.Head; cur != nil; cur = cur.Next {
		t.Log(cur.Value)
	}

	t.Log("从后往前：")
	for cur := dl.Tail; cur != nil; cur = cur.Prev {
		t.Log(cur.Value)
	}
}

func Test_ReverseDLinked(t *testing.T) {
	node5 := &DLinkedNode{Value: 5, Prev: nil, Next: nil}
	node4 := &DLinkedNode{Value: 4, Prev: nil, Next: node5}
	node3 := &DLinkedNode{Value: 3, Prev: nil, Next: node4}
	node2 := &DLinkedNode{Value: 2, Prev: nil, Next: node3}
	node1 := &DLinkedNode{Value: 1, Prev: nil, Next: node2}
	node2.Prev = node1
	node3.Prev = node2
	node4.Prev = node3
	node5.Prev = node4

	fmt.Println("Original:")
	for cur := node1; cur != nil; cur = cur.Next {
		fmt.Println(cur.Value)
	}

	// 反转
	reversed := ReverseDLinked(node1)
	fmt.Println("Reversed:")
	for cur := reversed; cur != nil; cur = cur.Next {
		fmt.Println(cur.Value)
	}
}
