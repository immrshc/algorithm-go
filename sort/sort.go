package sort

import "fmt"

func PrintResult(items []int) {
	unsortedItems := append([]int{}, items...)
	// BubbleSort(items)
	// SelectionSort(items)
	// InsertionSort(items)
	// ShellSort(items)
	// QuickSortByRecursion(items)
	// QuickSortByStack(items)
	// MergeSort(items)
	HeapSort(items)
	fmt.Printf("unsorted: %v\n", unsortedItems)
	fmt.Printf("sorted:   %v\n", items)
}
