package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(5))
}


func countAndSay(n int) string {
	l := []rune("1")
	for i := 1; i < n; i++ {
		c := '0'
		p := l[0]
		var tmp []rune
		for _, v := range l {
			if v != p {
				tmp = append(tmp, []rune{rune(c),p}...)
				p = v
				c = '0'
			}
			c++
		}
		tmp = append(tmp, []rune{rune(c),p}...)
		fmt.Println(string(tmp))
		l = tmp
	}
	return string(l)
}

