package main

import (
	"fmt"
	"math"
)

func main() {
	for i,v := range generate(5) {
		for j := 5 -i ;j > 0; j-- {
			fmt.Print(" ")
		}
		fmt.Println(v)
	}

	fmt.Println(getRow(4))

	fmt.Println(maxProfit1([]int{7,1,5,3,6,4}))
}

func generate(numRows int) [][]int {
	var sanj = [][]int{{1},{1,1}}
	for i := 2; i < numRows; i++ {
		a := []int{1}
		for j := 1; j < i; j++ {
			a = append(a, sanj[i-1][j-1] + sanj[i-1][j])
		}
		a = append(a, 1)

		sanj = append(sanj, a)
	}
	return sanj
}

func getRow(rowIndex int) []int {
	var preRow []int
	for i:=0; i <= rowIndex; i++ {
		var curRow []int
		curRow = []int{1}
		if i > 0 {
			curRow = []int{1}
			for j:=1; j < i ; j++  {
				curRow = append(curRow, preRow[j-1] + preRow[j])
			}
			curRow = append(curRow, 1)
		}
		preRow = curRow
	}
	return preRow
}

func maxProfit(prices []int) int {
	var lr int
	var buy int
	l := len(prices)
	for i := 0; i < l; i++ {
		buy = prices[i]
		for j := i + 1; j < l; j++ {
			if lr < prices[j] - buy  {
				lr = prices[j] - buy
			}
		}
	}
	return lr
}


func maxProfit1(prices []int) int {
	var profilePrice int = 0
	var minPrice int = math.MaxInt32
	l := len(prices)
	for i :=0 ;i < l;i++  {
		if prices[i] < minPrice  {
			minPrice = prices[i]
		} else if prices[i] - minPrice > profilePrice {
			profilePrice = prices[i] - minPrice
		}
	}
	return profilePrice
}

