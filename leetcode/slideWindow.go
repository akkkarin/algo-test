package leetcode

import "math"

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	max := 0
	tMap := make(map[byte]int, len(s))

	for ; right < len(s); right++ {
		if tMap[s[right]] > 0 {
			for ; left < right; left++ {
				if s[left] == s[right] {
					left++
					break
				} else {
					tMap[s[left]]--
				}
			}
			continue
		}

		tMap[s[right]]++
		if max < right-left+1 {
			max = right - left + 1
		}
	}
	return max
}

// https://leetcode-cn.com/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	left := 0
	right := 0
	sMap := map[byte]int{}
	min := math.MaxInt32
	res := ""
	count := len(t)
	for i := 0; i < len(s); i++ {
		sMap[s[i]] = 0
	}
	for i := 0; i < len(t); i++ {
		if _, ok := sMap[t[i]]; !ok {
			return res
		}
		sMap[t[i]]++
	}

	for ; right < len(s); right++ {
		if v := sMap[s[right]]; v > 0 {
			count--
		}
		sMap[s[right]]--
		if count == 0 {
			if right-left < min {
				min = right - left
				res = s[left : right+1]
			}

			for {
				if sMap[s[left]] == 0 {
					count++
				}
				sMap[s[left]]++
				left++
				if count == 0 {
					if right-left < min {
						min = right - left
						res = s[left : right+1]
					}
				} else {
					break
				}
			}
		}
	}
	return res
}

//https://leetcode-cn.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if nums == nil || len(nums) == 0 {
		return res
	}

	for left := 0; left <= len(nums)-k; left++ {
		max := math.MinInt32
		for right := left; right < left+k; right++ {
			if nums[right] > max {
				max = nums[right]
			}
		}
		res = append(res, max)
	}
	return res
}
