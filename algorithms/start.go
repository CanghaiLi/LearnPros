package main

import (
	"fmt"
	"github.com/CanghaiLi/LearnPros/algorithms/lSort"
)

var Arr = []int{6, 1, 3, 4, 10, 4, 5}

func main() {
	lSort.QuickSort(Arr, 0, len(Arr)-1)
	fmt.Println(Arr)
}
