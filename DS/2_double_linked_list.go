package DS

import (
	"fmt"
)

// DLinkedNode 双向链表节点
type DLinkedNode struct {
	Value any
	Prev  *DLinkedNode
	Next  *DLinkedNode
}

// DLinked 双向链表
type DLinked struct {
	Head *DLinkedNode
	Tail *DLinkedNode
}

// Unshift 头插
func (d *DLinked) Unshift(v any) {
	node := &DLinkedNode{Value: v}

	if d.Head == nil {
		d.Head = node
		d.Tail = node
	} else {
		// 向链表头插入，处理前、后指针
		d.Head.Prev = node
		node.Next = d.Head
		// 头指针前移
		d.Head = node
	}
}

// Shift 头出
func (d *DLinked) Shift() *DLinkedNode {
	if d.Head == nil {
		return nil
	}

	// 保存头结点
	node := d.Head
	if d.Head.Next == nil {
		// 该节点为最后一个节点
		// 将 Head、Tail 置 nil
		d.Head = nil
		d.Tail = nil
	} else {
		// 头结点后移
		d.Head = d.Head.Next
		// 新头结点的 Prev 置 nil，原头结点丢失
		d.Head.Prev = nil
	}

	return node
}

// Push 尾插
func (d *DLinked) Push(v any) {
	node := &DLinkedNode{Value: v}

	if d.Tail == nil {
		d.Head = node
		d.Tail = node
	} else {
		// 向链表尾插入，处理前、后指针
		d.Tail.Next = node
		node.Prev = d.Tail
		// 尾指针后移
		d.Tail = node
	}
}

// Pop 尾出
func (d *DLinked) Pop() *DLinkedNode {
	if d.Tail == nil {
		return nil
	}

	// 保存尾节点
	node := d.Tail
	if d.Tail.Prev == nil {
		// 该节点为最后一个节点
		// 将 Head、Tail 置 nil
		d.Head = nil
		d.Tail = nil
	} else {
		// 尾结点前移
		d.Tail = d.Tail.Prev
		// 新尾结点的 Next 置 nil，原尾结点丢失
		d.Tail.Next = nil
	}

	return node
}

// Remove 删除节点
func (d *DLinked) Remove(node *DLinkedNode) {
	if node == nil {
		return
	}

	if node.Prev == nil {
		// 删除头节点
		d.Head = node.Next
	} else {
		node.Prev.Next = node.Next
	}

	if node.Next == nil {
		// 删除尾节点
		d.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}
}

func (d *DLinked) Print() {
	for node := d.Head; node != nil; node = node.Next {
		fmt.Println(node.Value)
	}
}

// 把给定值都删除 - 双链表
func deleteGiveValueForDoubleList(head *DLinkedNode, k int) *DLinkedNode {
	return nil
}

// ReverseDLinked 反转双链表
func ReverseDLinked(head *DLinkedNode) *DLinkedNode {
	var prev, next *DLinkedNode

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
