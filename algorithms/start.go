package main

import (
	"fmt"
	"math"
)

var Arr = []int{6, 1, -3, 4, -9, 4, 5, 2}
var Str = "abc"

type SSS struct {
}

func main() {
	fmt.Println(Arr[:3])
	//window := [128]int{}
	fmt.Println(Str[1])
	fmt.Println(string(98))
	fmt.Println(math.Min(1.11, 2.22))

}

func ThreeSum(arr []int) [][]int {
	var res [][]int
	// test
	if arr == nil || len(arr) < 3 {
		return res
	}
	return res
}
