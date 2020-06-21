package sort

import "fmt"

func partition(items []int, l, r int) int {
	i := l
	j := r - 1
	pivot := items[r]

	fmt.Printf("-----r: %d-----\n", r)
	for {
		// pivotより等しいか大きくなる位置になるまで増分する
		for items[i] < pivot {
			i++
		}
		// iと等しくなるか、pivotよりも小さくなるまで減分する
		for i < j && items[j] >= pivot {
			j--
		}
		if i >= j {
			break
		}
		fmt.Printf("i: %d, j: %d, r: %d\n", i, j, r)
		fmt.Printf("items[i]: %v, items[r]: %v, items[j]: %v\n", items[i], pivot, items[j])
		items[i], items[j] = items[j], items[i]
	}
	items[i], items[r] = items[r], items[i]
	return i
}

func quickSort(items []int, l, r int) {
	if l >= r {
		return
	}
	v := partition(items, l, r)
	quickSort(items, l, v-1)
	quickSort(items, v+1, r)
}

func QuickSort(items []int) {
	quickSort(items, 0, len(items)-1)
}
