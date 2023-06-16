package Algorithm

/**
 * 链表面试题
 *  
 * 常用数据结构和技巧
 *  - 使用容器：哈希表、数组等
 *  - 快慢指针（为了省空间）
 */

////////////////
//  快慢指针  //
////////////////

// 1、输入链表头节点，奇数长度返回中点，偶数长度返回上中点
func ListMidpointOrUpperMidpoint(head *SingleNode) *SingleNode {
	// 零个节点 | 一个节点 | 两个节点
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	
	// 快慢指针
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		// 慢指针走 1 步
		slow = slow.Next
		// 快指针走 2 步
		fast = fast.Next.Next
	}

	// 返回奇数个的中点，偶数个的上中点
	return slow
}

// 2、输入链表头节点，奇数长度返回中点，偶数长度返回下中点
func ListMidpointOrLowerMidpoint(head *SingleNode) *SingleNode {
	// 零个节点 | 一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		// 慢指针走 1 步
		slow = slow.Next
		// 快指针走 2 步
		fast = fast.Next.Next
	}

	// 返回奇数个的中点，偶数个的下中点
	return slow.Next
}

// 3、输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
func ListBeforeMidpointOrUpperMidpoint(head *SingleNode) *SingleNode {
	// 零个节点 | 一个节点 | 两个节点 
	// 中点或上中点的前面没有节点
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	
	// 快慢指针
	slow := head
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		// 慢指针走 1 步
		slow = slow.Next
		// 快指针走 2 步
		fast = fast.Next.Next
	}

	// 返回奇数个的中点前一个，偶数个的上中点前一个
	return slow
}

// 4、输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
func ListBeforeMidpointOrLowerMidpoint(head *SingleNode) *SingleNode {
	// 零个节点 | 一个节点
	// 中点的前面没有节点
	if head == nil || head.Next == nil {
		return nil
	}

	// 快慢指针
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		// 慢指针走 1 步
		slow = slow.Next
		// 快指针走 2 步
		fast = fast.Next.Next
	}

	// 返回奇数个的中点前一个，偶数个的下中点前一个
	return slow
}

////////////////
//  使用数组  //
////////////////

// 笔试时（不考虑空间复杂度时）
// 使用数组，然后直接计算出中点
func ListMidpoint(head *SingleNode, cs int) *SingleNode {
	// 将所有节点保存到数组
	array := make([]*SingleNode, 0)
	for head != nil {
		array = append(array, head)
		head = head.Next
	}

	nums := len(array)
	
	// 计算中点
	switch cs {
	case 1:
		if nums == 0 {
			return nil
		}
		//  1、返回中点或上中点
		return array[(len(array)-1)/2]
	case 2:
		if nums == 0 {
			return nil
		}
		// 2、返回中点或下中点
		return array[len(array)/2]
	case 3:
		if nums < 3 {
			return nil
		}
		// 3、返回中点前一个或上中点前一个
		return array[(len(array)-1)/2 - 1]
	case 4:
		if nums < 3 {
			return nil
		}
		// 4、返回中点前一个或下中点前一个
		return array[len(array)/2 - 1]
	}

	return nil
}
