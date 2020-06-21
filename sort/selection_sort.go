package sort

func SelectionSort(items []int) {
	n := len(items)
	var lowestIndex int
	// 何番目に小さい値を見つけるか決める
	for i := 0; i < n; i++ {
		lowestIndex = i
		// 残りの範囲で最も小さい値を見つける
		for j := i; j < n; j++ {
			if items[j] < items[lowestIndex] {
				lowestIndex = j
			}
		}
		// 順番を入れ替える
		tmp := items[i]
		items[i] = items[lowestIndex]
		items[lowestIndex] = tmp
	}
}
