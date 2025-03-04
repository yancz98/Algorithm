package Algorithm

import (
	"algorithm/DS"
)

// LRUCache 最近最少使用算法
type LRUCache struct {
	capacity int
	cache    map[int]*DS.DLinkedNode
	DS.DLinked
}

// Constructor 构造函数
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}

	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}

	// 移至链表头
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node

	// 返回值
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.cache[key]
	if ok {
		// 移至链表头
		node.prev.next = node.next
		node.next.prev = node.prev

		node.prev = l.head
		node.next = l.head.next
		// l.head.next.prev = node
		l.head.next = node
	} else {
		if len(l.cache) == l.capacity {
			delete(l.cache, l.tail.prev.key)

			// 移除最少使用的（移除 tail 节点）
			l.tail.prev = l.tail.prev.prev
			l.tail.prev.next = l.tail
		}

		nn := &Node{
			key:   key,
			value: value,
			prev:  nil,
			next:  nil,
		}

		l.cache[key] = nn

		// 链表头插入节点
		nn.prev = l.head
		nn.next = l.head.next
		l.head.next.prev = nn
		l.head.next = nn
	}
}

func (l *LRUCache) MoveToHeader(key int) {

}
