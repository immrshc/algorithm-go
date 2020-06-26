package sort

import (
	"fmt"
	"math/rand"
	"time"
)

const MASK = 0x0f
const MaxUint16 = ^uint16(0)

func RadixSort() {
	items := randUint16Slice(20)
	n := len(items)
	sortedItems := make([]uint16, n)

	for bit := 0; bit < 16; bit += 4 {
		count := make([]uint16, 16)

		// ビットシフトして最下位4ビットの値の数を数える
		// e.g. 1200(10010110000)を右4ビットシフトすると75(1001011)になり、15(1111)との論理積は11(1011)になる
		for i := 0; i < n; i++ {
			count[(items[i]>>bit)&MASK]++
		}
		// 累積度数分布を作る
		for i := 0; i < len(count)-1; i++ {
			count[i+1] += count[i]
		}
		for i := n - 1; i >= 0; i-- {
			// インデックスへのアクセスなので先に1引く
			count[(items[i]>>bit)&MASK]--
			sortedItems[count[(items[i]>>bit)&MASK]] = items[i]
		}
		for i := 0; i < n; i++ {
			items[i] = sortedItems[i]
		}
	}
	fmt.Println(items)
}

func randUint16Slice(size int) []uint16 {
	rand.Seed(time.Now().UnixNano())
	slice := make([]uint16, size)
	for i := 0; i < size; i++ {
		slice[i] = uint16(rand.Intn(int(MaxUint16) + 1))
	}
	return slice
}
