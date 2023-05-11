package Algorithm

/**
 * 计数排序
 * 局限于非负数的小范围样本，如年龄 [0 ~ 150]
 *
 * 1.找到样本中的最大值
 * 2.创建 max + 1 个桶 bucket[0, ..., max]
 * 3.统计每个值的数量，写入 bucket，下标对应值，bucket 值对应数量
 * 4.输出 bucket 中的所有 idx （val 个）
 *
 * 时间复杂度：O(N)
 * 空间复杂度：O(M)
 */
func CountingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// 最大值
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// 创建 bucket
	bucket := make([]int, max+1)

	// 计数
	for _, v := range arr {
		bucket[v]++
	}

	// 输出
	i := 0
	for idx, cnt := range bucket {
		// 写入 cnt 个 idx
		for ; cnt > 0; cnt-- {
			arr[i] = idx
			i++
		}
	}
}
