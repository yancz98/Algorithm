package DS

import (
	"fmt"
)

// 数组 vs 单链表：
//
// | 操作 | 数组  | 单链表 |
// | --- | ---- | ----- |
// | 插入 | O(n) | O(1)  |
// | 删除 | O(n) | O(1)  |
// | 查找 | O(1) | O(n)  |

/**
 * 链表
 *  - 单向链表
 *     - 反转链表
 *     - 把给定值都删除
 *  - 双向链表
 *     - 实现头插、尾插、头出、尾出
 */

// SingleNode 单链表节点
type SingleNode struct {
	Data any
	Next *SingleNode
}

// NewSingleLinked 创建带头结点的单链表
//
// 注意：这里的 head 指向的是头结点（空的），而不是数据结点
//  在实际应用中，为了简化对链表状态的判定和处理，特别引入一个不存储数据元素的结点，称为头结点。
//
// Example:
//  head := NewSingleLinked()
//  head.Push(1)
//  head.Push(2)
//  head.Push(3)
//
//  head(0, 0x...) ---> (1, 0x...) ---> (2, 0x...) ---> (3, 0x...)
func NewSingleLinked() *SingleNode {
	return &SingleNode{
		Data: nil,
		Next: nil,
	}
}

// Find 查找元素
func (head *SingleNode) Find(idx uint) *SingleNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 从数据节点开始查找
	var i uint
	for cur := head; cur != nil; cur = cur.Next {
		if i == idx {
			return cur
		}

		if i < idx {
			i++
		}
	}

	return nil
}

// Insert 插入元素
//
// 在 idx 位置上插入 val
func (head *SingleNode) Insert(idx uint, val any) error {
	if head == nil {
		return fmt.Errorf("head is nil")
	}

	if idx == 0 {
		// 插入到第一个位置
		first := &SingleNode{
			Data: val,
			Next: head.Next,
		}

		head.Next = first
		return nil
	}

	// 找到要插入的前一个
	exist := head.Find(idx - 1)
	if exist == nil {
		return fmt.Errorf("idx %d not exist", idx-1)
	}

	exist.Next = &SingleNode{
		Data: val,
		Next: exist.Next,
	}

	return nil
}

// Delete 删除元素
func (head *SingleNode) Delete(idx uint) error {
	if head == nil || head.Next == nil {
		return fmt.Errorf("head is nil")
	}

	if idx == 0 {
		// 删除第一个元素
		head.Next = head.Next.Next
		return nil
	}

	// 找到要删除的前一个
	prev := head.Find(idx - 1)
	if prev == nil {
		return fmt.Errorf("idx %d not exist", idx-1)
	}

	prev.Next = prev.Next.Next
	return nil
}

// Push 尾插
func (head *SingleNode) Push(val any) error {
	if head == nil {
		return fmt.Errorf("head is nil")
	}

	// 空链表
	if head.Next == nil {
		head.Next = &SingleNode{
			Data: val,
			Next: nil,
		}

		return nil
	}

	// 找到尾结点
	cur := head.Next
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = &SingleNode{
		Data: val,
		Next: nil,
	}

	return nil
}

func (head *SingleNode) Print() {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Println(cur.Data)
	}
}

// ReverseSingleLinkedByRecursive 反转单链表（递归法）
//
// 注：递归反转法仅适用于反转不带头节点的链表
func ReverseSingleLinkedByRecursive(head *SingleNode) *SingleNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 最后一个节点，被返回到这里
	last := ReverseSingleLinkedByRecursive(head.Next)
	// 反转结点（从倒数第 2 个开始处理）
	head.Next.Next = head
	// 断开本层的最后一个节点
	head.Next = nil

	return last
}

// ReverseSingleLinkedByIterative 反转单链表 - 迭代法（双指针）
func ReverseSingleLinkedByIterative(head *SingleNode) *SingleNode {
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

// DeleteGiveValueForSingleList 把单链表中给定值都删除
func DeleteGiveValueForSingleList(head *SingleNode, v any) *SingleNode {
	// 先确定新 head 的位置
	// 若前几个元素都是需要删除的，则跳到第一个不是 v 的位置
	for head != nil {
		if head.Data != v {
			break
		}

		head = head.Next
	}

	// prev 记录前一个非 v 节点
	prev := head
	// 上一轮中已确定 head 不是 v，所以直接从 head 的下一个开始判断
	cur := head

	for cur != nil {
		// 命中
		if cur.Data == v {
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
