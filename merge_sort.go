package Algorithm

// 归并排序
//  - 分解：将长为 N 的数组分为 N/2 的左右两段
//  - 求解：递归调用分别将左右两边排好序
//  - 合并：合并两个有序的子序列使整体有序
func MergeSort(arr []int) {
    binarySort(arr, 0, len(arr)-1)
}

// 二分排序
// 递归出口：仅有一个数时，认为有序
// 递归体：一直二分，二分到左右各一个元素，然后合并排序
func binarySort(arr []int, l, r int) {
    // 递归出口
    if l == r {
        return
    }

    // 中位
    m := l + (r-l)>>1
    // 递归将左右两边排好序
    binarySort(arr, l, m)
    binarySort(arr, m+1, r)
    // 合并两个有序子序列
    // 目前 [l, m] [m+1, r] 已分别有序
    merge(arr, l, m, r)
}

// 合并思路：
// 对比左右两个数组，将较小的数先放入有序数组
// 将有序数组替换到元素的 [l, r] 位置
func merge(arr []int, l, m, r int) {
    // 存储合并后的数组
    orderly := make([]int, r-l+1)

    // 左右数组的起始索引
    i := 0
    li := l
    ri := m + 1

    // 任何一个越界，就结束比较
    for li <= m && ri <= r {
        // 将较小的数先放入 orderly
        if arr[li] < arr[ri] {
            orderly[i] = arr[li]
            li++
        } else {
            orderly[i] = arr[ri]
            ri++
        }
        i++
    }

    // 把左边或右边剩余的部分全部放入 orderly
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

    // 将 orderly 写入原 arr
    for j := 0; j < len(orderly); j++ {
        arr[l+j] = orderly[j]
    }
}