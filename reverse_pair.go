package Algorithm

/**
 * 剑指 Offer 51 数组中的逆序对
 * NC118 数组中的逆序对
 *  在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
 *  输入一个数组，求出这个数组中的逆序对的总数。
 *
 * 示例 1：
 *  输入: [7,5,6,4]
 *  输出: 5
 *
 * Tags：分治 归并排序
 */

// 数组中的逆序对
//  当前值的左边有多少个数大于当前值，当前值就能产生多少个逆序对
//  同样利用归并排序的红利，不需要迭代，可立即得出数量
func ReversePair(arr []int) int {
    return binarySortAndNum(arr, 0, len(arr)-1)
}

// 二分排序并统计逆序数量
func binarySortAndNum(arr []int, l, r int) int {
    // 递归出口
    // 一个元素时，没有逆序对
    if l == r {
        return 0
    }

    // 中位
    m := l + (r-l)>>1
    return binarySortAndNum(arr, l, m) +
        binarySortAndNum(arr, m+1, r) +
        mergeAndNum(arr, l, m, r)
}

// 合并 & 统计逆序数量（核心）
//  仅当左大右小时（arr[li] > arr[ri]）产生逆序对，
//  逆序对的数量为 左侧 li ~ m 的数量
func mergeAndNum(arr []int, l, m, r int) int {
    // 有序
    orderly := make([]int, r-l+1)

    // 左右指针
    li, ri := l, m+1
    i := 0
    num := 0
    for li <= m && ri <= r {
        // 右边小，产生逆序对
        if arr[li] > arr[ri] {
            // arr[ri] 产生的逆序对数量为 li ~ m
            num += m - li + 1
            orderly[i] = arr[ri]
            ri++
        } else {
            orderly[i] = arr[li]
            li++
        }

        i++
    }

    // 左右剩余
    for li <= m {
        orderly[i] = arr[li]
        i++
        li++
    }
    for ri <= r {
        orderly[i] = arr[ri]
        i++
        ri++
    }

    // 有序回填
    for j := 0; j < len(orderly); j++ {
        arr[l+j] = orderly[j]
    }

    return num
}
