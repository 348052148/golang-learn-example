package main

import (
	"fmt"
)

func main() {

	//fmt.Println(countPrimes(1231231))
	fmt.Println(isIsomorphic("ca","ab"))
	fmt.Println(isIsomorphic("ab","ca"))
}



func countPrimes(n int) int {
	var maps map[int]int = make(map[int]int)
	//放入列表
	for i:=2; i < n ; i++ {
		maps[i] = 1;
	}
	//找到未画的数，进行去除
	for dr := 2; dr < n; dr++ {
		if v,ok := maps[dr]; ok && v == 1 {
			for j := dr + 1; j < n; j++ {
				if j % dr == 0 {
					maps[j] = 0
				}
			}
		}
	}
	//取出值
	count :=0
	for i:=2; i < n;i++ {
		if v, ok := maps[i]; ok && v == 1 {
			count++
		}
	}
	return count
}


func isIsomorphic(s string, t string) bool {
	var hash map[uint8] uint8 = make(map[uint8]uint8)
	length := len(s)
	for i := 0; i< length; i++ {
		if v, ok := hash[s[i]]; ok && v != t[i] {
			return false
		}
		hash[s[i]] = t[i]
	}
	return true
}