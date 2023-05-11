package Algorithm

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
