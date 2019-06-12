package main

import "fmt"

func main() {
	fmt.Println(climbStairs2(5))
}
func climbStairs2(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	} else {
		tmp := 0
		a,b := 1, 2
		for i:=3; i < n+1 ; i++  {
			tmp = a + b
			a = b
			b = tmp
		}
		return tmp
	}
}
func climbStairs1(n int) int  {
	mem := make([]int,n+1)
	return climbStep(0, n, mem)
}
func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
func climbStep(i, n int, mem []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if mem[i] > 0 {
		return mem[i]
	}
	mem[i] = climbStep(i+1, n, mem) + climbStep(i+2, n, mem)
	return mem[i]
}