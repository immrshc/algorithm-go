package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintResult() {
	items := randSlice(50)
	unsortedItems := append([]int{}, items...)
	// BubbleSort(items)
	// SelectionSort(items)
	// InsertionSort(items)
	// ShellSort(items)
	// QuickSortByRecursion(items)
	// QuickSortByStack(items)
	// MergeSort(items)
	HeapSort(items)
	// BinSort()
	// CountingSort()
	// RadixSort()
	fmt.Printf("unsorted: %v\n", unsortedItems)
	fmt.Printf("sorted:   %v\n", items)
}

func randSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}
