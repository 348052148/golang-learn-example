package main

import (
	"strconv"
	"fmt"
	"strings"
	"bytes"
	"time"
)
const interval  = 12
func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Add(interval*time.Second))
	fmt.Println(addBinary("1010","1011"))
	fmt.Println(addBinary1("1000","10"))
}

func addBinary1(a string, b string) string {
	bl := len(b) - 1
	al := len(a) - 1
	buf := bytes.Buffer{}
	carry := 0
	for bl >= 0 || al >= 0 {
		bnum := 0
		anum := 0
		if bl >= 0 {
			bnum,_ = strconv.Atoi(string(b[bl]))
			bl--
		}
		if al >= 0 {
			anum,_ = strconv.Atoi(string(a[al]))
			al--
		}
		buf.WriteString(strconv.Itoa((anum + bnum + carry) % 2))
		carry = (anum + bnum + carry) / 2
	}
	if carry != 0 {
		//res = strconv.Itoa(carry) + res
		buf.WriteString(strconv.Itoa(carry))
	}
	return reverseString(buf.String())
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func addBinary(a string, b string) string {
	bl := len(b) - 1
	al := len(a) - 1
	var strs  []string
	carry := 0
	for bl >= 0 || al >= 0 {
		bnum := 0
		anum := 0
		if bl >= 0 {
			bnum,_ = strconv.Atoi(string(b[bl]))
			bl--
		}
		if al >= 0 {
			anum,_ = strconv.Atoi(string(a[al]))
			al--
		}
		strs = append(strs,strconv.Itoa((anum + bnum + carry) % 2))
		carry = (anum + bnum + carry) / 2
	}
	if carry != 0 {
		//res = strconv.Itoa(carry) + res
		strs = append(strs, strconv.Itoa(carry))
	}
	return strings.Join(reverse(strs),"")
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}