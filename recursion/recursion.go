package recursion

import (
	"algorithm/DS"
)

/**
 * 递归
 */

// 阶乘函数
//
// 递归出口：
//  - Factorial(0) = 1
//  - Factorial(1) = 1
// 递归体：
//  - Factorial(n) = n * Factorial(n - 1)
func factorial(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// ReverseSingleListByRecursion 反转单链表（递归法）
//
// 递归出口：
//  - 链表为空或只有一个节点，直接返回，不用反转
// 递归体：
//  - 递归处理下一个节点 head.Next 的反转
//  - 然后将下一个节点指向当前节点，当前节点指向 nil
//
// 时间复杂度 = O(N)
func ReverseSingleListByRecursion(head *DS.SingleNode) *DS.SingleNode {
	// 递归出口
	if head == nil || head.Next == nil {
		// 此时会找到链表中的最后一个元素，将作为新的头节点
		return head
	}

	// 递归处理下一个节点
	// 返回的新头结点保持不动
	last := ReverseSingleListByRecursion(head.Next)
	// 反转，后一个节点指向前一个
	head.Next.Next = head
	head.Next = nil

	return last
}
