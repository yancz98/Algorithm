package Algorithm

import (
    "math/rand"
    "time"
)

/**
 * 快速排序
 *  分区过程
 *  荷兰国旗问题
 *
 *  快排 1：用分区过程实现
 *  快排 2：用荷兰国旗问题实现
 *  快排 3：随机快排（选随机值分区）
 */

// 分区过程
//  在 arr[l,r] 位置上（假设 key = arr[r]），
//  将 <= key 的值放到 arr[l,r] 的左边，将 > key 的值放到 arr[l,r] 的右边，返回分区位置
//  将 <= key 的值交换到小于区的下一个位置，而 > key 的值直接进入下一轮，相当于大于区扩大
func partition(arr []int, l, r int) int {
    // 只有一个元素
    if l == r {
        return l
    }

    // 小于区
    lt := l - 1
    // 分区键
    key := arr[r]

    for i := l; i <= r; i++ {
        // 放入小于区逻辑
        // 当前值和 lt 下一个位置交换，lt 扩大，i++
        if arr[i] <= key {
            lt++
            arr[lt], arr[i] = arr[i], arr[lt]
        }

        // 放入大于区的逻辑，直接判断下一个
    }

    return lt
}

// 荷兰国旗问题
//  在 arr[l,r] 位置上（假设 key = arr[r]），
//  将小于 key 的放 arr[l, r] 左边
//  将等于 key 的放 arr[l, r] 中间
//  将大于 key 的放 arr[l, r] 右边
// 设置两个区指针，小于区、大于区，中部分就是等于区，返回等于区的左右边界
func dutchFlag(arr []int, l, r int) (lt, gt int) {
    // 只有一个元素
    if l == r {
        return
    }

    // 小于区、大于区指针
    lt, gt = l-1, r+1
    // 分区键
    key := arr[r]

    // 当 i 撞上大于区边界时，比较完成
    for i := l; i < gt; {
        if arr[i] < key {
            // 小于逻辑：与小于区下一个位置交换，lt++，比较下一个，i++
            lt++
            arr[lt], arr[i] = arr[i], arr[lt]
            i++
        } else if arr[i] == key {
            // 等于逻辑：比较下一个，i++
            i++
        } else {
            // 大于逻辑：与大于区的前一个位置交换，gt--
            // 新交换到 i 位置的元素还未比较，故留在 i 位置进行下一轮比较
            gt--
            arr[gt], arr[i] = arr[i], arr[gt]
        }
    }

    return lt, gt
}

// 快排 1.0
//  每次搞定分区位置的一个元素
//
// 时间复杂度（最差情况：原数组有序） = N-1 + N-2 + ... + 1 = (N^2 - N)/2 = O(N^2)
func QuickSort1(arr []int) {
    if len(arr) < 2 {
        return
    }

    process1(arr, 0, len(arr)-1)
}

func process1(arr []int, l, r int) {
    if l >= r {
        return
    }

    // 求分区位置，分区位置即是有序数组中的正确位置，所以本轮搞定了一个位置 p
    p := partition(arr, l, r)
    // 递归左、右部分 每次让分区位置排到正确位置
    process1(arr, l, p-1)
    process1(arr, p+1, r)
}

// 快排 2.0
//  每次搞定相等的一批 key（若没有相同元素，与快排 1.0 效率一样）
//
// 时间复杂度（最差情况：原数组有序） = N-1 + N-2 + ... + 1 = (N^2 - N)/2 = O(N^2)
func QuickSort2(arr []int) {
    if len(arr) < 2 {
        return
    }

    process2(arr, 0, len(arr)-1)
}

func process2(arr []int, l, r int) {
    if l >= r {
        return
    }

    // 小于区、大于区的范围
    lt, gt := dutchFlag(arr, l, r)

    // 递归使左右部分有序
    process2(arr, l, lt)
    process2(arr, gt, r)
}

// 快排 3.0（随机快排）
//  不使用固定位置的 arr[R] 作分区键，而是在数组中随机挑选一个与 arr[R] 交换
//
// 时间复杂度 = 最好情况和最差情况都是概率事件，最终期望 = O(N*logN)
// 空间复杂度 = O(logN)
func QuickSort3(arr []int) {
    if len(arr) < 2 {
        return
    }

    process3(arr, 0, len(arr)-1)
}

func process3(arr []int, l, r int) {
    if l >= r {
        return
    }

    // 选取随机位置，并交换到 R
    rand.Seed(time.Now().UnixNano())
    rd := l + rand.Intn(r-l+1)
    arr[rd], arr[r] = arr[r], arr[rd]

    // 小于区、大于区的范围
    lt, gt := dutchFlag(arr, l, r)

    // 递归使左右部分有序
    process3(arr, l, lt)
    process3(arr, gt, r)
}
