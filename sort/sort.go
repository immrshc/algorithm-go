package sort

import "fmt"

func PrintResult(items []int) {
	fmt.Printf("unsorted: %v\n", items)
	BubbleSort(items)
	fmt.Printf("sorted:   %v\n", items)
}
