package sort

func InsertionSort(items []int) {
	n := len(items)
	// 何番目までソートするか
	for i := 0; i < n-1; i++ {
		j := i
		// 指定の範囲内がソートされるまで入れ替える
		for j >= 0 && items[j] > items[j+1] {
			tmp := items[j]
			items[j] = items[j+1]
			items[j+1] = tmp
			j--
		}
	}
}
