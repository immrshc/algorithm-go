package sort

func downHeap(heap []int, from, to int) {
	// 0から始まる降順のヒープ（根が最大値）を構築する
	// fromは並び替え開始の節のインデックスで、toは並び替え対象の最後の節のインデックス
	val := heap[from]
	i := from
	for {
		// 子の節
		// e.g. 2の子の節は5, 6
		j := i*2 + 1
		if j > to || j < 0 { // j < 0 after int overflow
			break
		}
		// 二分ヒープなので兄弟は二つで、その中の最大の節の位置（j）を見つける
		if j+1 <= to && heap[j] < heap[j+1] {
			j++
		}
		// 子の最大の節より大きいか等しければヒープの条件を満たすので終了
		if val >= heap[j] {
			break
		}
		// 子の最大の節を親にする
		i, heap[i] = j, heap[j]
	}
	heap[i] = val
}

func HeapSort(items []int) {
	n := len(items)
	// 最下層の親から順に上へとヒープを構築する
	// e.g. 7個の節のヒープの最後の親のインデックスは、2
	// e.g. 8個の節のヒープの最後の親のインデックスは、3
	for i := n/2 - 1; i >= 0; i-- {
		downHeap(items, i, n-1)
	}
	// ヒープの根（最大値）を取り出して、残りの要素でヒープを再構築する
	for i := n - 1; i >= 0; i-- {
		items[0], items[i] = items[i], items[0]
		downHeap(items, 0, i-1)
	}
}
