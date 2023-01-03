package Algorithm

/**
 * NC349：计算数组的小和
 *  数组中第 i 个数的左侧比 i 小的数的和，叫数的小和，
 *  数组中所有数的小和叫数组的小和。
 *
 * 例如：数组 s = [1, 3, 5, 2, 4, 6]
 *  s[0] 的小和为 0；
 *  s[1] 的小和为 1；
 *  s[2] 的小和为 1+3=4；
 *  s[3] 的小和为 1；
 *  s[4] 的小和为 1+3+2=6；
 *  s[5] 的小和为 1+3+5+2+4=15。
 *  所以数组 s 的小和为 0+1+4+1+6+15=27
 *
 * Tags：归并排序、分治
 */

// 小和问题的等价问题：
//  当前值的右组有多少个数大于当前值，当前值就得累加多少次。
//  若当前值的右边有序，则不用迭代即可立即得出有多少个大于当前值的数。
//  利用归并排序的红利，左右已分别有序，即可实现。
func SmallSum(arr []int) int {
    return binarySortAndSum(arr, 0, len(arr)-1)
}

// 二分排序并求和
//  先用二分排序使左右两组分别有序，并得到左右两组的小和，
//  再合并有序的左右组，并求出小和
func binarySortAndSum(arr []int, l, r int) int {
    // 递归出口
    // 一个元素的数组没有小和
    if l == r {
        return 0
    }

    m := l + (r-l)>>2
    // 左组的小和 + 右组的小和 + 合并后的小和
    return binarySortAndSum(arr, l, m) +
        binarySortAndSum(arr, m+1, r) +
        mergeAndSum(arr, l, m, r)
}

// 合并 & 求和（求小和的核心思路）
//  合并之前，左右已分别有序，依次比较左右组的元素，
//  仅当 arr[li] < arr[ri]，累加 arr[li]*右组其后的元素数量
func mergeAndSum(arr []int, l, m, r int) int {
    // 已合并
    orderly := make([]int, r-l+1)

    // 左右指针
    li, ri := l, m+1
    i := 0
    sum := 0
    for li <= m && ri <= r {
        if arr[li] < arr[ri] {
            // 当前值累加 右组大于该值的数量 次
            sum += arr[li] * (r - ri + 1)
            orderly[i] = arr[li]
            li++
        } else {
            orderly[i] = arr[ri]
            ri++
        }

        i++
    }

    // 左右剩余元素
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

    // 已排序回写原数组
    for j := 0; j < len(orderly); j++ {
        arr[l+j] = orderly[j]
    }

    return sum
}
