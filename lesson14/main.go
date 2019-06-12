package main

import (
	"fmt"
	"image"
	"image/color"
	palette2 "image/color/palette"
	"image/jpeg"
	"os"
)

func main() {
	//fmt.Println(shellSort([]int{3,2,1,5,0,6,8,7,9,4}, 3))
	//fmt.Println(insertSort([]int{3,2,1,5,0,6,8,7,9,4}, 3))
	//fmt.Println(paoSort([]int{3,4,2,1}))
	//drawPng()
	//fmt.Println(baseSort([]int{3,2,10,5,0,6,8,7,9,4,9}))
	fmt.Println(isValid01("()"))
}
//var palette = []color.Color{color.White, color.Black}
func drawPng()  {

	palette := color.Palette{
		palette2.Plan9[70],
		palette2.Plan9[10],
	}
	fmt.Println(palette.Convert(palette2.Plan9[70]))
	fmt.Println(palette.Convert(color.Black))
	file,_ := os.Create("w.jpg")
	rect := image.Rect(0,0,320,480)
	p := image.NewPaletted(rect,palette)
	p.ColorModel()
	//for i:=0;i < 320 ;i+=10  {
	//	for j:=0;j < 480 ;j++  {
	//		p.SetColorIndex(i, j, 1)
	//	}
	//}
	if err:= jpeg.Encode(file, p, nil); err != nil {
		panic(err)
	}

}

func shellSort(nums []int, n int) []int  {
	for step := len(nums) / n ; step > 0; step = step / n  {
		for i:= step;i < len(nums) ; i++  {
			tmp := nums[i]
			j := i
			for j >= step && nums[j-step] > tmp {
				fmt.Println(step,nums[j], nums[j-step])
				nums[j] = nums[j-step]
				j -= step
			}
			nums[j] = tmp
		}
	}
	return  nums
}

func insertSort(nums []int, n int) []int {
	for i:=1; i < len(nums) ; i++  {
		tmp := nums[i] //需要向前插入的数
		var j int
		//比较如果成立 比较的数往后移动
		for j=i; j > 0 && tmp < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = tmp
	}
	return nums
}
// 10 /2 = 5
// 5 / = 2
// 2 /2 = 1
// 3,4,2,1
// 3,4,2,1 3,2,4,1 3,2,1,4
// 2,3,1,4 2,1,3,4
// 1,2,3,4
func paoSort(nums []int) []int {
	for i:=0;i < len(nums) - 1 ; i++  {
		for j:=0; j < len(nums) - i -1; j++  {
			if nums[j] > nums[j+1] {
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
	return nums
}

func getMax(nums []int) int  {
	max := 0
	for _,n := range nums  {
		if n > max {
			max = n
		}
	}
	return max
}

func baseSort(nums []int) int {
	max := getMax(nums)
	fmt.Println(max/1)
	for ext := 1; max/ext > 0 ;  ext *= 10 {
		nums = sort(nums, ext)
	}
	return 0
}
func sort(nums []int, ext int) []int {
	var output []int
	output = make([]int, len(nums))
	count := [10]int{0}
	for _, num := range nums {
		count[num / ext % 10]++
	}
	fmt.Println("Scounts:", count)
	for i:=1; i < 10; i++ {
		count[i] += count[i-1]
	}
	//fmt.Println("Change counts:", count)
	for i := len(nums) - 1; i >= 0; i-- {
		output[count[(nums[i]/ext)%10] - 1] = nums[i]
		count[(nums[i]/ext)%10]--
		//fmt.Println("Change counts:", count)
		fmt.Println(output)
	}
	fmt.Println(output)
	return output
}

func isValid(s string) bool {
	stack := []rune{}
	for _,s := range []rune(s) {
		if s == '(' || s == '[' || s == '{' {
			stack = append(stack,s)
		}
		if s == ')' || s == ']' || s == '}' {
			if len(stack) > 0  {
				d := stack[len(stack) - 1]
				var c rune
				switch s {
				case ')':
					c = '('
				case ']':
					c = '['
				case '}':
					c = '{'
				}
				if (c != d) {
					return false
				}
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
//使用map
func isValid01(s string) bool {
	stack := []rune{}
	maps := make(map[rune]rune)
	maps[')'] = '('
	maps['}'] = '{'
	maps[']'] = '['
	for _,s := range []rune(s) {
		if s == '(' || s == '[' || s == '{' {
			stack = append(stack,s)
		}
		if s == ')' || s == ']' || s == '}' {
			if len(stack) < 1 {
				return false
			}
			b := stack[len(stack)-1]
			if b != maps[s] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		fmt.Println(stack)
	}
	if len(stack) == 0 {
		return true
	}
	return false
}