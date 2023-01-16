package Algorithm

/**
 * 堆排序
 * 
 * 1、调堆，将 arr 调整为大根堆；
 * 2、每次将堆顶交换到本轮的最后位置后，再次调堆；
 *
 * 时间复杂度：O(N*logN)
 * 空间复杂度：O(1)
 */
func HeapSort(arr []int) {
    size := len(arr)

    // 调堆（大根堆）
    // O(N*logN)
    // for r := size - 1; r >= 0; r-- {
    //     heapify(arr, r, size)
    // }
    // 优化：O(N)
    for r := size/2 - 1; r >= 0; r-- {
        heapify(arr, r, size)
    } 

    // 此时堆顶元素为最大值
    // O(N*logN)
    for sz := size-1; sz > 0; sz-- {
        // 第一个位置与每轮的最后一个位置交换
        arr[0], arr[sz] = arr[sz], arr[0]

        // 继续堆化，大小为 sz 的堆
        heapify(arr, 0, sz)
    }
}
