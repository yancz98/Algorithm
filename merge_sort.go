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

// 归并排序 - 迭代实现
//
// 每轮将 eg 个数当作一组，迭代合并左右两组使有序
// 第一轮：eg = 1，依次迭代，使相邻的 2 个数有序
// 第二轮：eg = eg<<1，依次迭代，使相邻的 4 个数有序
// ...
// 当 eg >= len(arr) 时，结束迭代
func MergeSortByIterate(arr []int) {
    // 每组数量
    eg := 1
    len := len(arr)

    // 左+右两组的数组小于 len 说明还未完全合并
    for 2*eg < len {
        // 按每组 eg 个迭代，每次会合并掉 2*eg
        for l := 0; l+eg < len; l += 2 * eg {
            // 右边界，注意不能越界
            r := l + eg
            if r >= len {
                r = len
            }
            // 中位
            m := l + (r-l)>>1
            // 合并有序
            merge(arr, l, m, int(r))
        }

        // eg * 2
        eg <<= 1
    }
}
