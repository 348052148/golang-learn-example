package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(rob([]int{2,7,9,3,1}))
}
func rob(nums []int) int {
	var dp []int =make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i:=2 ; i <= len(nums); i++ {
		//dp[2] = max{dp[1],dp[0]+num[1]}
		//dp[3] = max{dp[2],dp[1]+num[2]}
		dp[i] = int(math.Max(float64(dp[i-1]),float64(dp[i-2] + nums[i-1])))
	}
	fmt.Println(dp)
	return dp[len(nums)]
}
//2,7,9,3,1

// 0, 2 => 7
//2, 7 => 9
//7, 11 => 3
//11, 10 => 1