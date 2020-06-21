package sort

func BubbleSort(items []int) {
	n := len(items)
	for i := 0; i < n-1; i++ {
		for j := n - 1; i < j; j-- {
			if items[j-1] > items[j] {
				tmp := items[j-1]
				items[j-1] = items[j]
				items[j] = tmp
			}
		}
	}
}
