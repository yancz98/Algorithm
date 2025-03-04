package DS

import (
	"testing"
)

func Test_StackWithDLinked(t *testing.T) {
	stack := StackWithDLinked{}
	stack.Push(1)
	stack.Push(2)
	t.Log("Pop:", stack.Pop())
	stack.Push(3)
	stack.Push(4)
	t.Log("Pop:", stack.Pop())
	stack.Push(5)

	t.Log("Foreach:")
	for !stack.Empty() {
		t.Log("    ", stack.Pop())
	}
}

func Test_StackWithArray(t *testing.T) {
	stack := StackWithArray{}
	stack.Push(1)
	stack.Push(2)
	t.Log("Pop:", stack.Pop())
	stack.Push(3)
	stack.Push(4)
	t.Log("Pop:", stack.Pop())
	stack.Push(5)

	t.Log("Foreach:")
	for !stack.Empty() {
		t.Log("    ", stack.Pop())
	}
}

func Test_StackWithQueue(t *testing.T) {
	stack := StackWithTwoQueues{}
	stack.Push(1)
	stack.Push(2)
	t.Log("Pop:", stack.Pop())
	stack.Push(3)
	stack.Push(4)
	t.Log("Pop:", stack.Pop())
	stack.Push(5)

	t.Log("Foreach:")
	for !stack.Empty() {
		t.Log("    ", stack.Pop())
	}
}
