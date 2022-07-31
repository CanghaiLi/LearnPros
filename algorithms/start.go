package main

import "fmt"

var Arr = []int{6, 1, -3, 4, -9, 4, 5, 2}

func main() {
	fmt.Println(Arr[:3])
}

func ThreeSum(arr []int) [][]int {
	var res [][]int
	// test
	if arr == nil || len(arr) < 3 {
		return res
	}
	return res
}
