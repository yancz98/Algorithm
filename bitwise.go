package Algorithm

import "fmt"

/**
 * 位运算
 *  & 按位与
 *  | 按位或
 *  ^ 按位异或：a ^ b
 *  ^ 按位取反：^x
 *  >> 右移：M >> N = M / 2^N
 *  << 左移：M << N = M * 2^N
 */

/**
 * 按位异或的性质：
 *  0 ^ N == N
 *  N ^ N == 0
 *  异或运算满足交换律和结合律：a ^ b ^ c == a ^ c ^ b
 */

// 不用额外的变量交换两个变量
func swap(x, y int) {
    fmt.Printf("Before: x = %v, y = %v \n", x, y)
    x = x ^ y // x = x ^ y
    y = x ^ y // y = (x ^ y) ^ y = x
    x = x ^ y // x = (x ^ y) ^ x = y
    fmt.Printf("After: x = %v, y = %v \n", x, y)

    /**
     * 类似：
     * a = a + b
     * b = a - b // b = (a+b)-b = a
     * a = a - b // a = (a+b)-a = b
     */
}

// 数组中仅有一个数是奇数个，其余数都是偶数个，找出奇数个的那个数
func onlyOneOdd() {
    arr := []int{1, 2, 3, 4, 3, 1, 2, 3, 4} // 3 出现了 3次

    var eor int
    for _, v := range arr {
        eor ^= v
    }

    // eor = 1 ^ 1 ^ 2 ^ 2 ^ 3 ^ 3 ^ 3 ^ 4 ^ 4 = 3
    fmt.Println(eor)
}

// 保留 int 型二进制数中最右侧的 1
// 如：11111010 => 00000010
func retainRightmost1(x int) int {
    // x = 250
    // x = 11111010
    fmt.Printf("x = %.8b\n", x)

    y := x & (^x + 1)
    // 取反       ^x = 00000101
    // +1    ^x + 1 = 00000110
    // x & (^x + 1) = 00000010

    // y = 00000010
    fmt.Printf("y = %.8b\n", y)

    return y
}

// 数组中有两个数是奇数个，其余数都是偶数个，找出这两个数
func twoOdd() {
    arr := []int{1, 1, 4, 4, 5, 9, 9, 9}

    // 假设那两个数分别为：a、b，由题意：a ≠ b
    // 将所有元素异或结果：eor = a ^ b ≠ 0
    var eor int
    for _, v := range arr {
        eor ^= v
    }
    // eor = 00001100
    fmt.Printf("eor = %.8b \n", eor)

    // 继续求 a、b 分别是什么
    // 因为 eor ≠ 0，所以其上至少有一个位置为 1
    // 根据异或性质 eor 结果为 1 的位置，a 和 b 的同位置一定 0 ^ 1
    // 保留 eor 上最右侧的 1
    eorRightmost1 := eor & (^eor + 1)
    // eorRightmost1 = 00000100
    fmt.Printf("eorRightmost1 = %.8b \n", eorRightmost1)

    // 将数组中的元素分为：与 eorRightmost1 同位置是 1 的元素和同位置是 0 的元素
    // 对与 eorRightmost1 同位置是 1 的元素集合异或
    // 其中偶数个的那些数不影响异或结果，即结果为 a 或 b
    with0Eor := 0
    for _, v := range arr {
        // 与 eorRightmost1 同位置是 0
        if v&eorRightmost1 == 0 {
            with0Eor ^= v
        }
    }

    // with0Eor 的结果为 a 或 b
    // 则另一个结果：another = with0Eor ^ eor = with1Eor ^ a ^ b
    // 当 with0Eor = a 时，another = b
    // 当 with0Eor = b 时，another = a

    fmt.Printf("这两个数分别是：%v, %v \n", with0Eor, eor^with0Eor)
}

// 二进制数中 1 的数量
func binary1Num() {
    x := 250
    count := 0
    // x = 11111010
    fmt.Printf("x = %.8b \n", x)

    for x != 0 {
        count++

        // 抹掉最右侧的 1
        // r1   = 00000010
        // x    = 11111010
        // x^r1 = 11111000
        r1 := x & (^x + 1)
        x ^= r1
    }

    // 1 的数量为：6
    fmt.Printf("1 的数量为：%v \n", count)
}
