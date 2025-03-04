package DS

import (
	"fmt"
	"testing"
)

func Test_QueueWithDLinked(t *testing.T) {
	queue := QueueWithDLinked{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	t.Log("Dequeue:", queue.Dequeue())
	queue.Enqueue(3)

	t.Log("Foreach:")
	for !queue.Empty() {
		t.Log(queue.Dequeue())
	}
}

func Test_QueueWithArray(t *testing.T) {
	queue2 := InitQueue(5)
	queue2.Enqueue(5)
	queue2.Enqueue(6)
	queue2.Enqueue(7)
	t.Log("Dequeue:", queue2.Dequeue())
	queue2.Enqueue(8)
	queue2.Enqueue(9)
	// queue2.Enqueue(10) // Full...

	t.Log("Foreach:")
	for !queue2.Empty() {
		t.Log(queue2.Dequeue())
	}

	t.Log("Empty:", queue2.Dequeue()) // Empty...
}

func TestQueue(t *testing.T) {
	fmt.Println("QueueWithDLinked...")

	// fmt.Println(queue.Dequeue())

	fmt.Println("QueueWithArray...")

	fmt.Println("QueueWithTwoStacks...")
	queue3 := QueueWithTwoStacks{}
	queue3.Enqueue(5)
	queue3.Enqueue(6)
	queue3.Enqueue(7)
	fmt.Println(queue3.Dequeue())
	queue3.Enqueue(8)
	queue3.Enqueue(9)
	for !queue3.Empty() {
		fmt.Println(queue3.Dequeue())
	}
	// fmt.Println(queue3.Dequeue()) // Empty...
}
