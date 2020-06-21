package sort

import "fmt"

func ShellSort(items []int) {
	n := len(items)
	h := 1
	// n/9を超えず、互いに倍数にならない最大の値を最初のhにする
	// e.g. n = 1098 (1098/9 = 122) の場合、121, 40, 13, 4, 1 より h = 121
	for ; h < n/9; h = 3*h + 1 {
	}
	// ループの最後で余分に3*h+1された分を戻す
	// 数列を辿って、h-ソートの幅を小さくしていく
	for h /= 3; h > 0; h = h / 3 {
		fmt.Printf("h: %d\n", h)
		// 何番目までh-ソートするか
		// e.g. h = 121 の場合、i = 121, 122, 123, ... 1098
		for i := h; i < n; i++ {
			j := i
			for j >= h && items[j-h] > items[j] {
				items[j], items[j-h] = items[j-h], items[j]
				j -= h
			}
		}
	}
}
