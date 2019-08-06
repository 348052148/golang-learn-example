package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(findLongestPlalindromeString("adaelele"))
	/*
	s = "mississippi"
p = "mis*is*p*."
	 */
	 //a := "123";
	fmt.Println(isMatch("aab", "c*a*b"))
}

func isMatch(s, p string) bool  {
	if s == p  {
		return true
	}
	isFlush := false
	if len(s) > 0 && len(p) > 0 && (s[0] == p[0] || p[0] == '.') {
		isFlush = true
	}
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (isFlush && isMatch(s[1:], p))
	}
	return isFlush && isMatch(s[1:], p[1:])
}

func preHandleString(str string) string  {
	buf := bytes.Buffer{}
	buf.WriteByte('#')
	for i:=0; i < len(str); i++ {
		buf.WriteByte(str[i])
		buf.WriteByte('#')
	}
	return buf.String()
}
func findLongestPlalindromeString(str string) string{
	//预处理字符串加上#分割
	str = preHandleString(str)
	//各个回文覆盖范围
	var halfLen []int =make([]int, len(str))
	//左边界范围
	rightSide := 0
	//边界中心点
	rightSideCenter := 0
	//记录最长回文中心
	center := 0
	//记录最长回文半长
	longestHalf := 0
	for i := 0; i < len(str); i++ {
		//
		needCalc := true
		//如果i在回文边界内 必须是内 边缘也不行
		if i < rightSide {
			// 计算相对rightSideCenter的对称位置
			//找到关于其中心对称的点的位置
			leftCenter := 2 * rightSideCenter - i;
			// 根据回文性质得到的结论
			// 关于回文对称的点的回文半长度一致
			halfLen[i] = halfLen[leftCenter];
			// 如果超过了右边界，进行调整
			if(i + halfLen[i] > rightSide) {
				halfLen[i] = rightSide - i;
			}
			// 如果根据已知条件计算得出的最长回文小于右边界，则不需要扩展了
			if(i + halfLen[leftCenter] < rightSide) {
				// 直接推出结论
				needCalc = false
			}
		}
		if needCalc {
			for i - 1 - halfLen[i] >= 0 && i + 1 + halfLen[i] < len(str) {
				if str[i + 1 + halfLen[i]] == str[i - 1 - halfLen[i]] {
					halfLen[i]++;
				} else {
					break
				}
			}
			rightSide = i + halfLen[i]
			rightSideCenter = i
			//
			if(halfLen[i] > longestHalf) {
				center = i;
				longestHalf = halfLen[i];
			}
		}
	}
	buf := bytes.Buffer{}
	for i := center - longestHalf + 1; i <= center + longestHalf; i += 2  {
		buf.WriteByte(str[i])
	}

	return buf.String()
}