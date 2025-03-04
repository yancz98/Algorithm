package Algorithm

import (
	"fmt"
	"testing"
)

// 位运算
func Test_binaryOperate(t *testing.T) {
	binaryOperate1()
	binaryOperate2()
}

func Test_swap(t *testing.T) {
	swap(5, 10)
}

func Test_onlyOneOdd(t *testing.T) {
	onlyOneOdd()
}

func Test_retainRightmost1(t *testing.T) {
	retainRightmost1(250)

	fmt.Println(-74 << 2)
}

func Test_twoOdd(t *testing.T) {
	twoOdd()
}

func Test_binary1Num(t *testing.T) {
	binary1Num()
}

// 计算数组的小和
func TestSmallSum(t *testing.T) {
	fmt.Println(SmallSum([]int{1, 3, 4, 2, 5}))    // 16
	fmt.Println(SmallSum([]int{1, 3, 5, 2, 4, 6})) // 27
}

// 数组中的逆序对
func TestReversePair(t *testing.T) {
	fmt.Println(ReversePair([]int{7, 5, 6, 4}))             // 5
	fmt.Println(ReversePair([]int{1, 2, 3, 4, 5, 6, 7, 0})) // 7
	fmt.Println(ReversePair([]int{1, 2, 3, 4, 5, 6, 7}))    // 0
}

func TestTrie(t *testing.T) {
	// 纯小写字母的串
	tr0 := InitTrie0()
	tr0.Insert("apple")
	fmt.Println(tr0.Search("apple"))
	fmt.Println(tr0.Search("app"))
	fmt.Println(tr0.StartWith("app"))

	tr0.Insert("app")
	fmt.Println(tr0.Search("app"))

	fmt.Println("------------------")

	// 任意字符的串
	tr := InitTrie()
	tr.Insert("YCZ")
	fmt.Println(tr.Search("YCZ"))
	fmt.Println(tr.Search("YYY"))
	tr.Insert("中华文化博大精深")
	fmt.Println(tr.StartWith("中华"))

	// 删除
	fmt.Println("------------------")
	tr.Delete("YCZ")
	fmt.Println(tr.Search("YCZ"))
	tr.Delete("123")
	fmt.Println(tr.StartWith("中华文化"))
}
