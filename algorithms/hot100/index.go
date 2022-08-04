package hot100

import (
	"sort"
)

/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
nums = [2,7,11,15], target = 9  ==>   [0,1]
*/
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for cur, val := range nums {
		if prev, exist := dict[val]; exist {
			return []int{prev, cur}
		} else {
			dict[target-val] = cur
		}
	}
	return nil
}

/**********************************************************************************************************/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。
它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
342 + 465 = 807 --- l1 = [2,4,3], l2 = [5,6,4] ==> [7,0,8]
*/
// tips: 同时遍历两个链表，注意进位
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	extra := 0
	var res, cur *ListNode
	// 考虑链表遍历完，但进位还有的情况 extra != 0
	for extra != 0 || l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + extra
		sum, extra = sum%10, sum/10

		if res == nil {
			res = &ListNode{Val: sum}
			cur = res
		} else {
			cur.Next = &ListNode{Val: sum}
			cur = cur.Next
		}
	}
	return res
}

/**********************************************************************************************************/

/*
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
abcabcbb abc ==> 3
pwwkew wke ==> 3
*/
func lengthOfLongestSubstring(s string) int {
	var res, start, tmp int
	for i := 0; i < len(s); i++ {
		tmp = i - start + 1
		for j := start; j < i; j++ {
			if s[i] == s[j] {
				tmp--
				start = j + 1
				// 跳出j的循环
				break
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

/**********************************************************************************************************/

/*
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
	***算法的时间复杂度应该为 O(log(m+n))***
*/
// tips: 第k小数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var m = len(nums1)
	var n = len(nums2)
	if n < m { // 确保nums1比nums2短，即确保m比n小
		return findMedianSortedArrays(nums2, nums1)
	}
	var midM = (m - 1) / 2
	var midN = (n - 1) / 2

	if m == 0 { // 处理长度为0的情况
		if n%2 == 1 {
			return float64(nums2[midN])
		}
		return float64(nums2[midN]+nums2[midN+1]) / 2
	}

	if m == 1 || m == 2 { // 边界条件
		if n < 3 { // n小于3的情况下，取nums2所有元素和nums1的元素进行排序
			for i := 0; i < n; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else if n%2 == 1 { // n大于2且为奇数的情况下，取nums2中间3位和nums1的元素进行排序
			for i := midN - 1; i < midN+2; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else { // 其他情况下，取nums2的中间4位和nums1的元素进行排序
			for i := midN - 1; i < midN+3; i++ {
				nums1 = append(nums1, nums2[i])
			}
		}
		sort.Ints(nums1)
		m = len(nums1)
		midM = (m - 1) / 2

		if len(nums1)%2 == 1 {
			return float64(nums1[midM])
		} else {
			return float64(nums1[midM]+nums1[midM+1]) / 2
		}
	}

	// n为奇数时，midNP==midN。n为偶数时，midNP==midN+1。
	var midNP = midN
	if n%2 == 0 {
		midNP++
	}

	if nums1[midM] == nums2[midNP] {
		return float64(nums1[midM])
	}
	if nums1[midM] < nums2[midNP] {
		//消除nums1数组0至midM-1下标的元素，和nums2数组n-midM下标之后的元素
		return findMedianSortedArrays(nums1[midM:], nums2[:n-midM])
	}
	//消除nums2数组0至midM-1下标的元素，和nums1数组n-midM下标之后的元素
	return findMedianSortedArrays(nums2[midM:], nums1[:m-midM])
}
