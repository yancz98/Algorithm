package Algorithm

import (
    "fmt"
    "math"
)

/**
 * 基数排序
 *  思想：按 value 各位的值进行排序
 *  要求样本为非负数（这里为十进制数）
 *
 * 1.准备一个桶 bucket[10]，用来存储位为 [0, ..., 9] 的值数量
 * 2.按个位、十位、百位... 分别排序
 *  3.按个位排序后，保留相对次序
 *  4.按十位排序后，保留相对次序
 *    ...
 *  5.按最高位排序后，保留相对次序
 */
func RadixSort(arr []int) {
    if len(arr)  <= 1 {
        return
    }

    max := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    // 最高位数
    hd := getDigit(max)

    fmt.Println("最大值：", max, "最高位数：", hd)

    // 按 个位、十位、百位... 分别排序
    for i := 0; i < hd; i++ {
        fmt.Printf("~~ 第 %d 轮 ~~\n", i + 1)
        
        // 记录每轮（个、十、百...）位分别为 [0, ..., 9] 的值数量
        var bucket [10]int
        for a := 0; a < len(arr); a++ {
            // 各位基数
            radix := getRadix(arr[a], int(math.Pow(10, float64(i))))
            // 对应桶位置值加 1
            bucket[radix]++
        }

        fmt.Println("bucket:", bucket)

        // bucket 统计出了（个、十、百...）位的各值的对应数量
        // 将 bucket 中的当前位置与前一位置的值累加
        // 即得到相对位的最终位置，注意：最大位置 == len(arr)
        for b := 1; b < len(bucket); b++ {
            bucket[b] += bucket[b-1] 
        }

        fmt.Println("Bucket:", bucket)

        // 临时数组
        temp := make([]int, len(arr))
        for c := len(arr) - 1; c >= 0; c-- {
            // 各位基数
            radix := getRadix(arr[c], int(math.Pow(10, float64(i))))
            // 1 ~ len => 0 ~ len - 1
            bucket[radix]--
            // 写入临时数组的最终位置
            temp[bucket[radix]] = arr[c]
        }

        // 回写到原数组，保留本轮按位排序的相对次序
        for d := 0; d < len(arr); d++ {
            arr[d] = temp[d]
        }

        fmt.Println("本轮结果：", temp)
    }
}

// 获取 x 的最高位数
func getDigit(x int) int {
    hd := 1

    for x/10 > 0 {
        hd++
        x = x/10
    }

    return hd
}

// 获取 bit 位的基数
//   1   个位
//   10  十位
//   100 百位
func getRadix(x, bit int) int {
    return x/bit % 10
}
