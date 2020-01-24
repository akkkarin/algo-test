package leetcode

// f(1) = nums[1]
// f(2) = max{num[1],num[2]}
// f(3) = max{f(1) + num[3], f(2)}
// https://leetcode-cn.com/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	f := make([]int, len(nums))
	f[0] = nums[0]
	max := f[0]
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			f[i] = max
			if nums[i] > max {
				f[i] = nums[i]
				max = f[i]
			}
			continue
		}
		f[i] = f[i-2] + nums[i]
		if f[i] > f[i-1] {
			if f[i] > max {
				max = f[i]
			}
		} else {
			f[i] = f[i-1]
		}
	}
	return max
}

// step = 1 or 2
// f(1) = 1
// f(2) = 2
// f(n) = f(n-1) + f(n-2)
// https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	f := make([]int, n)
	f[0] = 1
	f[1] = 2

	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n-1]
}

// https://leetcode-cn.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	mp := make([]int, len(nums))
	max := nums[0]
	mp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		if tmp < mp[i-1]+nums[i] {
			tmp = mp[i-1] + nums[i]
		}
		mp[i] = tmp
		if max < tmp {
			max = tmp
		}
	}
	return max
}
