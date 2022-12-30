package Algorithm

import (
    "fmt"
    "testing"
)

func TestSelectSort(t *testing.T) {
    arr := []int{9, 2, 5, 8, 3, 7, 6}
    fmt.Println("Unordered:", arr)
    SelectSort(arr)
    fmt.Println("Orderly:", arr)
}

func TestBubbleSort(t *testing.T) {
    arr := []int{9, 2, 5, 8, 3, 7, 6}
    fmt.Println("Unordered:", arr)
    BubbleSort(arr)
    fmt.Println("Orderly:", arr)
}

func TestInsertSort(t *testing.T) {
    arr := []int{9, 2, 5, 8, 3, 7, 6}
    fmt.Println("Unordered:", arr)
    InsertSort(arr)
    fmt.Println("Orderly:", arr)

    a := 10
    fmt.Println(a^0)
}