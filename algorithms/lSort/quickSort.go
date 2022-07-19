package lSort

func QuickSort(arr []int, left, right int) {
	if left < right {
		mid := partition(arr, left, right)
		QuickSort(arr, left, mid-1)
		QuickSort(arr, mid+1, right)
	}
}
func partition(arr []int, left, right int) int {
	// Pick a number at any position, for example the most left
	// Store it in a temporary variable
	temp := arr[left]
	for left < right {
		for arr[right] >= temp && left < right {
			right -= 1
		}
		arr[left] = arr[right]
		for arr[left] <= temp && left < right {
			left += 1
		}
		arr[right] = arr[left]
	}
	arr[left] = temp
	return left
}
