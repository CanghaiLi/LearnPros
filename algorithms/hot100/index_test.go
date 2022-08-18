package hot100

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	l2 := &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}
	l3 := addTwoNumbers(l1, l2)

	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcdabcef..."))
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-1, 3}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

// [1,2,3,4] 4
// [2,4,5,6,7] 5
// 4 2

// [1,2,3] 3
// [2,4,5,6] 4
//
