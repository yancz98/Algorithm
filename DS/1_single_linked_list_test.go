package DS

import (
	"testing"
)

func Test_SingleLinked(t *testing.T) {
	head := NewSingleLinked()
	head.Push(1)
	head.Push(2)
	head.Push(3)
	head.Print()

	// 第一个为空头结点
	dataHead := head.Next

	t.Log(dataHead.Find(2))

	t.Log(dataHead.Insert(3, 4))
	dataHead.Print()

	t.Log(dataHead.Delete(2))
	dataHead.Print()
}

func Test_ReverseSingleLinkedByRecursive(t *testing.T) {
	head := &SingleNode{
		Data: 1,
		Next: &SingleNode{
			Data: 2,
			Next: &SingleNode{
				Data: 3,
				Next: &SingleNode{
					Data: 4,
				},
			},
		},
	}

	t.Log("Original:")
	head.Print()

	head = ReverseSingleLinkedByRecursive(head)
	t.Log("Reverse:")
	head.Print()
}

func Test_ReverseSingleLinkedByIterative(t *testing.T) {
	head := &SingleNode{
		Data: 1,
		Next: &SingleNode{
			Data: 2,
			Next: &SingleNode{
				Data: 3,
				Next: &SingleNode{
					Data: 4,
				},
			},
		},
	}

	t.Log("Original:")
	head.Print()

	// 反转
	head = ReverseSingleLinkedByIterative(head)
	t.Log("Reverse:")
	head.Print()
}

func Test_DeleteGiveValueForSingleList(t *testing.T) {
	head := &SingleNode{
		Data: 1,
		Next: &SingleNode{
			Data: 2,
			Next: &SingleNode{
				Data: 3,
				Next: &SingleNode{
					Data: 4,
				},
			},
		},
	}

	deleted := DeleteGiveValueForSingleList(head, 1)
	t.Log("Deleted:")
	deleted.Print()
}
