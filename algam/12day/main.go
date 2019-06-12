package main

import "fmt"

func main() {
	fmt.Println(Sqrtnd(16))
}
func Sqrtnd(x int) int {
	if x == 0 {
		return 0
	}
	last, res  := 0.0, 1.0
	for res != last {
		last = res
		res = (res + float64(x) / res) / 2;
	}
	return int(res)
}
func Sqrt(x int) int {
	if x <= 1 {
		return x
	}
	var mid int
	l, r := 0, x
	for l <= r {
		mid = (l + r) / 2
		if x / mid == mid {
			return mid
		}
		if x / mid > mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	l,r := 0, x
	for l <= r {
		mid := (l + r) / 2
		fmt.Println(l,r,l + r,mid)
		if mid == x/mid {
			return mid
		}
		if x/mid > mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
//9 / 4 =  2 mid = 4
//9 / 2 =  4 mid = 4