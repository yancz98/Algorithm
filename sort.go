package Algorithm

/**
 * 比较排序算法
 *  选择排序
 *  冒泡排序
 *  插入排序
 */

// 选择排序
//
// 每轮目标：选出一个最小的元素与该轮第一个位置交换
// 假设最小元素在第一个位置，与后面元素依次比较，最后交换
// 第 1 轮：从 [1:N) 中选出一个最小的元素与第一个位置交换，需比较 N-1 次，还可能存在一次交换
// 第 2 轮：从 [2:N) 中选出一个最小的元素与第二个位置交换，需比较 N-2 次
// ...
// 第 N-1 轮：从 [N-1:N) 需比较 1 次
//
// 时间复杂度 = N-1 + N-2 + ... + 1 = (N^2 - N)/2 = O(N^2)
func SelectSort(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        // 每轮假设最小的元素为第一个
        min := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[min] {
                min = j
            }
        }

        // 本轮查找结束，交换位置
        if min != i {
            // 假设不成立，最小元素不在本轮的第一个位置，则交换
            arr[i], arr[min] = arr[min], arr[i]
        }
    }
}

// 冒泡排序
//
// 每轮目标：将最大元素交换到该轮的最后一个位置
// 第 1 轮：从 [0:N-1) 依次比较并交换，需比较+交换 2*(N-1) 次
// 第 2 轮：从 [0:N-2) 依次比较并交换，需比较+交换 2*(N-2) 次
// ...
// 第 N-1 轮：从 [0:1] 依次比较并交换，需比较+交换 2*1 次
//
// 时间复杂度 = 2*(N-1 + N-2 + ... + 1) = N^2 - N = O(N^2)
func BubbleSort(arr []int) {
    // 控制每轮比较的范围：[0:N-round) （round 表示轮次）
    for i := len(arr) - 1; i > 0; i-- {
        for j := 0; j < i; j++ {
            // 如果前一个大于后一个，则交换
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

// 插入排序
//
// 每轮目标：把该轮的值插入到左边已排好序的队列中，保证 [0:round] 上有序（round 表示轮次）
// 第 1 轮：将 arr[1] 依次与前 1 个值 [0] 比较，小于则交换
// 第 2 轮：将 arr[2] 依次与前 2 个值 [1,0] 比较，小于则交换
// ...
// 第 N-1 轮：将 arr[N-1] 依次与前 N-1 个值 (N-1:0] 比较，小于则交换
//
// 最好情况：原数组已排好序 = O(N)
// 最坏情况：原数组逆序 = 2*(N-1 + N-2 + ... + 1) = N^2 - N = O(N^2)
// 时间复杂度 = O(N^2)，取最坏情况
func InsertSort(arr []int) {
    // 假设 arr[0] 已排好序，从 arr[1] 开始处理
    for i := 1; i < len(arr); i++ {
        for j := i; j > 0; j-- {
            if arr[j] < arr[j-1] {
                arr[j], arr[j-1] = arr[j-1], arr[j]
            }
        }
    }
}
