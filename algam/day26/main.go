package main

import "math"

func main() {
}

func MaxSubSequeue(nums []int) int {
	l := len(nums)
	var dp [l+1]int
	dp[0] = 0
	dp[1] = nums[0]
	for i:=2 ; i < l+1 ; i++ {
		dp[i] = math.Max(dp[i-1],dp[i-2])+dp[i]
	}
	return dp[l-1];
}