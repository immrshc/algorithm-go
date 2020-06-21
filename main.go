package main

import (
	"math/rand"
	"time"

	"github.com/shoichiimamura/algorithm-go/sort"
)

func main() {
	sort.PrintResult(randSlice(20))
}

func randSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}
