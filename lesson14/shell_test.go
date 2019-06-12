package main

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T)  {
	tests := []struct {
		nums []int
		n int
		ans []int
	}{
		{
			nums:[]int{5,6,3,1,2,4,7,9},
			n: 2,
			ans:[]int{1,2,3,4,5,6,7,8,9},
		},

	}
	for _,ts := range tests{
		//if shellSort(ts.nums, ts.n) == ts.ans {
		//
		//}
		fmt.Println(ts)
	}
}
