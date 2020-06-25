package sort

import "fmt"

func BinSort() {
	// 10未満の自然数で重複がない要素である前提
	items := []int{9, 2, 3, 4, 7, 1, 8}

	// binの初期化
	bin := make([]int, 10)
	for i := range bin {
		bin[i] = -1
	}

	// binのインデックスと要素を対応させる
	for _, v := range items {
		bin[v] = v
	}

	i := 0
	for _, v := range bin {
		if i >= len(items) {
			break
		}
		if v == -1 {
			continue
		}
		items[i] = v
		i++
	}
	fmt.Printf("bin: %v\nitems: %v\n", bin, items)
}
