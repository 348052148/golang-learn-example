package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"math"
	"strings"
)


type Json struct {
	data interface{}
}

func newJson(bytes []byte) *Json  {
	j := &Json{
		data:make(map[string]interface{}),
	}
	err := j.UnmarshalJSON(bytes)
	if err!= nil {
		panic(err)
	}
	return j
}
func (j *Json)Print()  {
	fmt.Printf("%v \n",j.data)
}
func (j *Json)Get(key string) *Json  {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return &Json{m[key]}
	}
	return &Json{nil}
}

func (j *Json)GetIndex(index int) *Json  {
	if m, ok := (j.data).([]interface{}); ok {
		return &Json{m[index]}
	}
	return &Json{nil}
}

func (j *Json) UnmarshalJSON(p []byte) error {
	return json.Unmarshal(p, &j.data)
}
// Implements the json.Marshaler interface.
func (j *Json) MarshalJSON() ([]byte, error) {
	return json.Marshal(&j.data)
}

func (j *Json) GetPath(branch ...string) *Json {
	jin := j
	for _, p := range branch {
		jin = jin.Get(p)
	}
	return jin
}

func (j *Json)String() (string, error) {
	if s,ok := j.data.(string); ok {
		return s, nil
	}
	return "", errors.New("解析错误")
}


func main() {
	jsonStr :=
		`
      [{
          "person": [{
             "name": "piao",
             "age": 30,
             "email": "piaoyunsoft@163.com",
             "phoneNum": [
                 "13974999999",
                 "13984999999"
             ]
          }, {
             "name": "aaaaa",
             "age": 20,
             "email": "aaaaaa@163.com",
             "phoneNum": [
                 "13974998888",
                 "13984998888"
             ]
          }, {
             "name": "bbbbbb",
             "age": 10,
             "email": "bbbbbb@163.com",
             "phoneNum": [
                 "13974997777",
                 "13984997777"
             ]
          }]
      }]
      `

	jsons := newJson([]byte(jsonStr))
	jsons.Print()
	fmt.Println(jsons.GetIndex(0).Get("person"))

	jsonb, err := simplejson.NewJson([]byte(jsonStr))
	if err!=nil {
		panic(err)
	}
	fmt.Println(jsonb.GetIndex(0).Get("person"))
	type user struct {
		name int
		age int
		email string
	}
	//u := user{}
	//
	var params []interface{}
	json.NewDecoder(strings.NewReader(jsonStr)).Decode(&params)
	fmt.Println(params)
	//if v, ok := params.([]interface{}); ok {
	//	fmt.Println(v)
	//}

	fmt.Println(numJewelsInStones("ab", "AabbcCd"))

	fmt.Println(reverse(-2147483648))

	fmt.Println(isPalindrome(12121))

	fmt.Println(romanToInt("LIV"))
	var a A
	b := B{}
	a = b
	if _,ok := a.(B);ok {
		a.echo()
	}else {
		fmt.Println("FAIL")
	}

	fmt.Println("-------")
	strs := []string {"flot","flotassa","flotias"}
	fmt.Println(longestCommonPrefix02(strs))

	str := "abcdef"
	fmt.Println(strings.Index(str, "b"))
}

type A interface {
	echo()
}
type B struct {

}

func (b B)echo()  {
	fmt.Printf("ECHO")
}

func twoSum(numbers []int, target int) (int,int)  {
	maps := make(map[int]int)
	for i,v := range numbers {
		if vi,ok := maps[target - v]; ok {
			return i, vi
		}
		maps[v]=i
	}
	panic("error")
}
//O2n
func numJewelsInStones(J string, S string) int {
	count := 0
	maps := make(map[rune]string)
	for _,v := range []rune(J) {
		maps[v] = "b"
	}
	for _,k := range []rune(S) {
		if _,ok := maps[k]; ok {
			count++
		}
	}
	return count
}

func reverse(x int) int {
	res := 0
	for {
		pop := x % 10
		if x == 0 {
			break
		}
		//最大尾数是7 超过溢出
		if res > math.MaxInt32 / 10 || (res == math.MaxInt32 / 10 && pop > 7) {
			return 0
		}
		//最小尾数是8 超过溢出
		if res < math.MinInt32 / 10 || (res == math.MinInt32 / 10 && pop < -8) {
			return 0
		}
		x /= 10
		res = res * 10 + pop
	}
	return res
}

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x!=0) {
		return  false
	}
	rev := 0
	for {
		if x < rev {
			break
		}
		rev = rev * 10 + x % 10
		x /= 10
	}
	return rev == x || rev / 10 == x
}
func romanToInt(s string) int {
	maps := map[string]int {
				"I":1,
				"V":5,
				"X":10,
				"L":50,
				"C":100,
				"D":500,
				"M":1000}
	total := 0;
	for i,_ := range s {
		if i < len(s)-1 && maps[string(s[i])] < maps[string(s[i+1])] {
			total -= maps[string(s[i])]
		}else {
			total += maps[string(s[i])]
		}
	}
	return total
}
// 水平扫描
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for k:=1;k< len(strs);k++ { //3 * n ,  10 * 2 * n
			//var n int = 0
			//for i,s1 := range []rune(strs[k]) {
			//	if i >= len(prefix) || s1 != rune(prefix[i]) {
			//		break
			//	}
			//	n = i
			//}
			//prefix = prefix[:n+1]
			//
			for {
				if strings.Index(strs[k], prefix) == 0 {
					break
				}
				fmt.Println(strings.Index(strs[k], prefix))
				prefix = prefix[:len(prefix)-1]
			}
	}
	return prefix
}
// 纵向扫描
func longestCommonPrefix01(strs []string) string {
	prefix := strs[0]
	for i,s := range strs[0]{ //3 * n ,  10 * 2 * n
		for j:=1; j<len(strs); j++ {
			if i == len(strs[j]) || s != rune(strs[j][i]) {
				return prefix[:i+1]
			}
		}
	}
	return prefix
}
// 分治
func longestCommonPrefix02(strs []string) string {
	if len(strs) == 0 {
		return  ""
	}
	return commonPrefix(strs,0, len(strs))
}

func lcp (left, right string) string {
	var min  int
	if len(left) > len(right) {
		min = len(right)
	}else {
		min = len(left)
	}
	for i:=0; i < min ; i++  {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:min]
}

func commonPrefix(strs []string, l, r int) string {
	fmt.Println(l, r)
	if l == r {
		return strs[l]
	} else {
		mid := (l + r) /2
		left := commonPrefix(strs, l, mid)
		right := commonPrefix(strs, mid, r)
		return lcp(left, right)
	}
}
