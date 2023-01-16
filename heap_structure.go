package Algorithm

/**
 * 堆结构
 *  堆结构就是用数组实现的完全二叉树结构
 *  完全二叉树中如果每颗子树的最大值都在顶部就是大根堆（最大堆）
 *  完全二叉树中如果每颗子树的最小值都在顶部就是小根堆（最小堆）
 *  堆结构的 heapInsert 与 heapify 操作
 *  堆结构的增大和减少
 *  优先级队列结构，就是堆结构
 *
 * 将 arr[0] 作为根节点：
 *  左孩子 = 2 * i + 1
 *  右孩子 = 2 * i + 2
 *  父节点 = (i - 1)/2
 *
 * 将 arr[1] 作为根节点：
 *  左孩子 = 2 * i     = i << 1
 *  右孩子 = 2 * i + 1 = (i << 1) | 1
 *  父节点 = i / 2     = i >> 1
 */

type Heap struct {
    heap []int
}

// 往堆中插入元素（最大堆）
//  将元素 x 追加到数组中（最后一个位置）
//  将最后位置的元素与其父级比较，比父级大则交换到父级，依次进行比较
//  插入完成后，依然是堆结构
func heapInsert(arr []int, x int) []int {
    arr = append(arr, x)
    // 从新插入位置开始比较
    cur := len(arr) - 1
    // 如果 arr[cur] > parent，则交换
    for arr[cur] > arr[(cur - 1) / 2] {
        arr[(cur - 1) / 2], arr[cur] = arr[cur], arr[(cur - 1) / 2]
        // 拿当前父节点继续向上比较
	    cur = (cur - 1) / 2
    }

    return arr
}

// 堆化（最大堆）
//  将以 root 为根的节点调整成堆结构
//  从 root 开始，依次将左、右子树堆化
//  size 为堆的长度
func heapify(arr []int, root, size int) {
    for {
        // 左子
        l := 2*root + 1
        // 没有左子，右子更不存在
        if l >= size {
            break
        }

        // 左右子中较大的
        large := l
        r := l + 1
        // 右子没越界且比左子大
        if r < size && arr[r] > arr[l] {
            large = r
        }

        // 父节点比左右子都大，不用交换
        if arr[root] > arr[large] {
            break
        }

        // large > root 则交换
        arr[large], arr[root] = arr[root], arr[large]
        
        // 将交换的子树继续堆化
        root = large
    }
}

// 判空
func (h *Heap) Empty() bool {
    if len(h.heap) == 0 {
        return true
    }

    return false
}

// 入堆（最大堆）
//  保证任意元素入堆后，自动调整成之前的堆结构（最大堆）
func (h *Heap) Push(x int) {
    h.heap = heapInsert(h.heap, x)
}

// 根节点出堆
//  先用最后一个位置顶替根节点
//  再拿根节点与左右子节点比，将较大的换上来
func (h *Heap) Pop() int {
    if h.Empty() {
        panic("Heap Empty...")
    }
    
    // 根节点
    res := h.heap[0]

    // 将最后一个节点换到第一个
    h.heap[0] = h.heap[len(h.heap)-1]
    // 数组切短
    h.heap = h.heap[:len(h.heap)-1]

    // 将 0 位置的元素放到堆中合适的位置
    heapify(h.heap, 0, len(h.heap))

    return res
}

// 将数组初识化为一个堆
func Init(source []int) (h Heap) {
    // 方式 1：遍历插入
    // 从叶子节点起，最多向上比较树的高度次：O(logN)
    // 时间复杂度：O(N*logN)
    /*for _, v := range source {        // O(N)
        h.heap = heapInsert(h.heap, v)  // O(logN)
    }*/

    // 方式 2：将每个子树堆化
    // 从后往前，依次将其作为根节点，实现堆化
    // 根据二叉树的性质：后 N/2 个节点为叶子节点，无需堆化
    // 最后一层：N/2 + 1 个节点，比较 0 次
    // 倒数第二层：N/4 + 1 个节点，最多向下比较 1 次
    // 倒数第三层：N/8 + 1 个节点，最多向下比较 2 次
    // ...
    // 最后一层：1 个节点，最多向下比较树的高度 logN 次
    // 
    // 时间复杂度 = N/4 + N/8*2 + N/16*3 + ... + logN（等比数列）= O(N)
    for r := len(source)/2 - 1; r >= 0; r-- {
        heapify(source, r, len(source))
    }
    h.heap = source

    return
}

