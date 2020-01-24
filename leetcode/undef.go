package leetcode

import "math"

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)

	arr := make([]int, l/2+1)
	i1 := 0
	i2 := 0
	for i := 0; i <= l/2; i++ {
		n1 := math.MaxInt32
		if i1 < len(nums1) {
			n1 = nums1[i1]
		}
		n2 := math.MaxInt32
		if i2 < len(nums2) {
			n2 = nums2[i2]
		}

		if n1 < n2 {
			arr[i] = n1
			i1++
		} else {
			arr[i] = n2
			i2++
		}
	}

	if l%2 == 0 {
		return (float64(arr[l/2] + arr[l/2-1])) / 2
	}
	return float64(arr[l/2])
}

// https://leetcode-cn.com/problems/add-two-numbers/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	p1 := l1
// 	p2 := l2
// 	r := &ListNode{}
// 	c := r
// 	carry := 0
// 	for {
// 		tmp := 0
// 		if p1 != nil {
// 			tmp = p1.Val
// 			p1 = p1.Next
// 		}
// 		if p2 != nil {
// 			tmp += p2.Val
// 			p2 = p2.Next
// 		}
// 		tmp += carry
// 		if tmp >= 10 {
// 			carry = 1
// 			c.Val = tmp - 10
// 		} else {
// 			carry = 0
// 			c.Val = tmp
// 		}

// 		if p1 == nil && p2 == nil {
// 			if carry > 0 {
// 				c.Next = &ListNode{Val: carry}
// 			}
// 			return r
// 		}
// 		c.Next = &ListNode{}
// 		c = c.Next
// 	}
// }
