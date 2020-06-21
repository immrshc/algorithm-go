package sort

func BubbleSort(items []int) {
	n := len(items)
	// 何番目に小さい値を見つけるか決める
	for i := 0; i < n-1; i++ {
		// 残りの範囲で、後ろから順に比較する
		for j := n - 1; i < j; j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}
