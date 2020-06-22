package sort

import "fmt"

func mergeSort(items []int, low, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2
	mergeSort(items, low, mid)
	mergeSort(items, mid+1, high)

	// 実際に必要な配列の長さはhigh-low+1
	tmp := make([]int, len(items))
	// 1, 21, 38
	for i := low; i <= mid; i++ {
		tmp[i] = items[i]
	}
	// 1, 21, 38, 99, 12, 8
	for i, j := mid+1, high; i <= high; i, j = i+1, j-1 {
		tmp[i] = items[j]
	}

	// 1(i), 21, 38, 99, 12, 8(j) => [1]
	// 21(i), 38, 99, 12, 8(j)    => [1, 8]
	// 21(i), 38, 99, 12(j)       => [1, 8, 12]
	// 21(i), 38, 99(j)           => [1, 8, 12, 21]
	// 38(i), 99(j)               => [1, 8, 12, 21, 38]
	// 99(i, j)                   => [1, 8, 12, 21, 38, 99]
	i, j := low, high
	for k := low; k <= high; k++ {
		if tmp[i] <= tmp[j] {
			items[k] = tmp[i]
			i++
		} else {
			items[k] = tmp[j]
			j--
		}
	}
	fmt.Printf("low: %d, high: %d, items[low:high+1]: %v\n", low, high, items[low:high+1])
}

func MergeSort(items []int) {
	mergeSort(items, 0, len(items)-1)
}
