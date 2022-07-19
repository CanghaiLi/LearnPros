package lSort

func QuickSort(arr []int, left, right int) {
	if left < right {
		mid := partition(arr, left, right)
		QuickSort(arr, left, mid-1)
		QuickSort(arr, mid+1, right)
	}
}

func partition(list []int, left, right int) int {
	temp := list[left]
	for left < right {
		//从右边找比temp小的树
		for list[right] >= temp && left < right {
			right -= 1 //右指针左移一步
		}
		list[left] = list[right] //把右边的值写到左边的空位

		for list[left] <= temp && left < right {
			left += 1
		}
		list[right] = list[left]
	}
	// left right都可以
	list[left] = temp
	return left
}
