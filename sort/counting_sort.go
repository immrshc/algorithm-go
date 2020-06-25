package sort

import "fmt"

func CountingSort() {
	items := []int{9, 2, 5, 5, 3, 5, 1, 2}

	count := make([]int, 10)
	for _, v := range items {
		count[v]++
	}
	for i := 0; i < len(count)-1; i++ {
		count[i+1] += count[i]
	}
	fmt.Printf("count: %v\n", count)

	sortedItems := make([]int, len(items))
	for i := 0; i < len(items); i++ {
		count[items[i]]--
		sortedItems[(count[items[i]])] = items[i]
	}
	fmt.Printf("items: %v\nsortedItems: %v\n", items, sortedItems)
}
