package sort

import "fmt"

func partition(items []int, l, r int) int {
	i := l
	j := r - 1
	pivot := items[r]

	// fmt.Printf("-----r: %d-----\n", r)
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
		// fmt.Printf("i: %d, j: %d, r: %d\n", i, j, r)
		// fmt.Printf("items[i]: %v, items[r]: %v, items[j]: %v\n", items[i], pivot, items[j])
		items[i], items[j] = items[j], items[i]
	}
	items[i], items[r] = items[r], items[i]
	return i
}

func quickSortByRecursion(items []int, l, r int) {
	if l >= r {
		return
	}
	v := partition(items, l, r)
	quickSortByRecursion(items, l, v-1)
	quickSortByRecursion(items, v+1, r)
}

func QuickSortByRecursion(items []int) {
	quickSortByRecursion(items, 0, len(items)-1)
}

func QuickSortByStack(items []int) {
	n := len(items)
	low := make([]int, n)
	high := make([]int, n)

	low[0] = 0
	high[0] = n - 1
	sp := 1

	for sp > 0 {
		sp--
		l := low[sp]
		r := high[sp]
		fmt.Printf("sp: %d, l: %d, r: %d\n", sp, l, r)

		if l >= r {
			continue
		} else {
			v := partition(items, l, r)
			// 左右の小さい方のスライスを先に処理する
			if v-l < r-v {
				// 左の方が小さい
				low[sp], high[sp] = v+1, r
				sp++
				low[sp], high[sp] = l, v-1
			} else {
				// 右の方が小さい
				low[sp], high[sp] = l, v-1
				sp++
				low[sp], high[sp] = v+1, r
			}
			sp++
		}
	}
}
