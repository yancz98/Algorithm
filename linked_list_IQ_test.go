package Algorithm

import (
	"fmt"
	"testing"

	"algorithm/DS"
)

func Test_ListMidpoint(t *testing.T) {
	list := &DS.SingleNode{
		Data: 1,
		Next: nil,
	}
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)
	// list.Push(6)
	list.Print()

	fmt.Println("midpoint-1:", ListMidpointByArray(list, 1))
	fmt.Println("midpoint-2:", ListMidpointByArray(list, 2))
	fmt.Println("midpoint-3:", ListMidpointByArray(list, 3))
	fmt.Println("midpoint-4:", ListMidpointByArray(list, 4))

	fmt.Println("       upper-midpoint:", ListMidpointOrUpperMidpoint(list))
	fmt.Println("       lower-midpoint:", ListMidpointOrLowerMidpoint(list))
	fmt.Println("before-upper-midpoint:", ListBeforeMidpointOrUpperMidpoint(list))
	fmt.Println("before-lower-midpoint:", ListBeforeMidpointOrLowerMidpoint(list))
}
