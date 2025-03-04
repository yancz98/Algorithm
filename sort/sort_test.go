package sort

import (
	"math/rand"
	"testing"
	"time"
)

// 选择排序
func TestSelectSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)
	t.Log("Unordered:", arr)
	SelectSort(arr)
	t.Log("Orderly:", arr)
}

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	arr := []int{9, 2, 5, 8, 3, 7, 6}
	t.Log("Unordered:", arr)
	BubbleSort(arr)
	t.Log("Orderly:", arr)
}

// 插入排序
func TestInsertSort(t *testing.T) {
	arr := []int{9, 2, 5, 8, 3, 7, 6}
	t.Log("Unordered:", arr)
	InsertSort(arr)
	t.Log("Orderly:", arr)

	a := 10
	t.Log(a ^ 0)
}

// 归并排序
func TestMergeSort(t *testing.T) {
	arr := []int{3, 1, 2, 4, 6, 8, 5, 7}
	MergeSort(arr)
	t.Log(arr)

	arr2 := []int{3, 1, 2, 6, 4, 8, 5, 7, 9}
	MergeSortByIterate(arr2)
	t.Log(arr2)
}

// 快速排序
func TestQuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	arr := rand.Perm(6)
	t.Log("arr:", arr)
	t.Log("partition:", partition(arr, 0, len(arr)-1))
	t.Log("after:", arr)

	arr = rand.Perm(7)
	t.Log("arr:", arr)
	t.Log(dutchFlag(arr, 0, len(arr)-1))
	t.Log("after:", arr)

	arr1 := rand.Perm(8)
	t.Log("arr1:", arr1)
	QuickSort1(arr1)
	t.Log("sort1:", arr1)

	arr2 := rand.Perm(9)
	t.Log("arr2:", arr2)
	QuickSort2(arr2)
	t.Log("sort2:", arr2)

	arr3 := rand.Perm(10)
	t.Log("arr3:", arr3)
	QuickSort3(arr3)
	t.Log("sort3:", arr3)
}

// 堆结构
func TestHeapStructure(t *testing.T) {
	// heapInsert
	var arr []int
	arr = heapInsert(arr, 1)
	arr = heapInsert(arr, 9)
	arr = heapInsert(arr, 6)
	arr = heapInsert(arr, 8)
	arr = heapInsert(arr, 2)
	arr = heapInsert(arr, 3)
	arr = heapInsert(arr, 5)
	t.Log("heapInsert:", arr)

	// heapify
	arr = []int{1, 4, 7, 3, 0, 2, 5}
	t.Log("heapify before:", arr)
	for i := 0; i < len(arr); i++ {
		heapify(arr, i, len(arr))
	}
	t.Log("heapify after:", arr)

	// Push
	heap := Heap{}
	heap.Push(2)
	heap.Push(5)
	heap.Push(8)
	heap.Push(3)
	heap.Push(6)
	heap.Push(9)
	t.Log("Push End...", heap.heap)

	// Pop
	for !heap.Empty() {
		t.Log(heap.Pop())
	}

	// t.Log(heap.Pop())

	// 将数组初始化为堆
	rand.Seed(time.Now().UnixNano())
	arr = rand.Perm(8)
	t.Log("arr:", arr)
	h := Init(arr)
	t.Log("heap:", h.heap)
}

// 堆排序
func TestHeapSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(8)
	t.Log("arr:", arr)
	HeapSort(arr)
	t.Log("heapSort:", arr)
}

// 计数排序
func TestCountingSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// arr := rand.Perm(100)
	arr := []int{5, 5, 5, 3, 3, 7}
	t.Log("Before:", arr)
	CountingSort(arr)
	t.Log("After", arr)
}

// 桶排序
func TestRadixSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// x := rand.Intn(999999)
	// digit := getDigit(x)
	// t.Log(x ,"的位数：", digit)
	// t.Log(x, "从低到高位分别为：")
	// for i := 0; i < digit; i++ {
	//     t.Log(getRadix(x, int(math.Pow(10, float64(i)))))
	// }

	// 基数排序
	arr := []int{rand.Intn(100), rand.Intn(1000), rand.Intn(10000), rand.Intn(1000),
		rand.Intn(1000), rand.Intn(1000), rand.Intn(100), rand.Intn(10)}

	t.Log("Before:", arr)
	RadixSort(arr)
	t.Log("After:", arr)
}
