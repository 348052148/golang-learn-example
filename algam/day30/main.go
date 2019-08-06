package main

import "fmt"

type Line struct {
	l, r int
} 

func main() {
	fmt.Println(juli([]Line{
		Line{l:1,r:5},Line{l:1,r:2},Line{l:2,r:6},Line{l:2,r:3},Line{l:7,r:9},
		Line{l:7,r:8},Line{l:8,r:10},Line{l:11,r:12},
		}, 12))
	fmt.Println(juli01([]Line{
		Line{l:1,r:5},Line{l:1,r:2},Line{l:2,r:6},Line{l:2,r:3},Line{l:7,r:9},
		Line{l:7,r:8},Line{l:8,r:10},Line{l:11,r:12},
	}, 12))
	stoneGame([]int{3,9,1,2})

	fmt.Println(dfs01([][]int{
		{0,1,1,1},
		{0,1,0,0},
		{0,0,0,0},
	},0, 3, []int{}))
	fmt.Println(0 ^ 1)
	//fmt.Println("RobotPr", RobotPr(5, 4,4, []int{1,2}))
	fmt.Println("RobotPr", RobotPr(3, 1,2, []int{1}))
}

func RobotPr(n int, l, r int, cmd []int) float32 {
	//dp公式  dp[m][i] = 0.5 * dp[m-1][i - v + n % n] + 0.5 * dp[m-1][n + v % n]
	mlen := len(cmd)
	//初始化dp 二维结构
	dp := make([][]float32, mlen+1)
	for i := 0; i < mlen+1; i++ {
		dp[i] = make([]float32, n)
	}
	//起始位置 概率 1
	dp[0][0] = 1
	//0步能到达的其他格子，概率为0
	for i := 1; i < n; i++ {
		dp[0][i] = 0
	}
	step := 0
	for step < len(cmd) {
		v := cmd[step]
		for i := 0; i < n; i++ {
			dp[step+1][i] = 0.5 * dp[step][(i - v + n)%n] + dp[step][(i+v)%n]*0.5
		}
		step++
	}
	for i :=0; i < len(dp); i++ {
		for j:=0; j <len(dp[i]); j++ {
			fmt.Printf("%0.2f\t", dp[i][j])
		}
		fmt.Println()
	}
	//计算最终概率
	var ans float32 = 0.0
	for l = l-1 ; l < r; l++ {
		ans += dp[step][l]
	}
	return ans
}

func dfs01(maps [][]int, row int, k int, lines []int) int {
	ans := 0
	if k == 0 {
		fmt.Println("方案：", lines)
		return 1
	}
	if row >= len(maps) {
		return 0
	}
	for col := 0; col < len(maps[row]); col++ {
		if maps[row][col] == 0 {
			//检查是否占用列
			ok := true
			for j:=0; j < len(lines); j++  {
				if lines[j] == col {
					ok = false
					break
				}
			}

			if ok {
				//记录状态
				lines = append(lines, col)
				//处理下一行记录
				ans += dfs01(maps, row+1, k-1, lines)
				//回溯状态
				lines = lines[:len(lines)-1]
			}
		}
	}
	return ans + dfs01(maps, row+1, k, lines)
}
var lines []int = make([]int, 10)
var cnt = 0
var ans = 0
func dfs(maps [][]int, row int) {
	//记录已经放置的棋子
	if cnt > 4 {
		ans++
		return
	}
	if row >= len(maps) {
		return
	}
	for i := 0; i < len(maps[0]); i++ {
			//可放置
			if maps[row][i] == 0 {
				//检查是否在同一列
				ok := true
				for j := 0; j < cnt; j++ {
					if lines[j] == i {
						ok = false
						break
					}
				}
				//
				if ok {
					lines[cnt] = i
					cnt++
					dfs(maps, row)
					cnt--
				}
		}
	}
	dfs(maps, row + 1);
}

type T struct {
	Fir int
	Sec int
}

func stoneGame(piles []int) {
	//dp[i][j] = max(p[i]+dp[i+1][j], dp[i][j-1] + p[j])
	n := len(piles)
	dp := make([][]T, n)
	for i := 0; i < n; i++  {
		dp[i] = make([]T, n)
		for j := 0; j < n ; j++  {
			dp[i][j] = T{}
		}
	}
	for i := 0; i < n; i++ {
		dp[i][i] = T{Fir:piles[i], Sec:0}
	}
	fmt.Println(dp)
	for l := 2; l <= n ; l++ {
		for i := 0; i <= n - 1; i++ {
			j := l + i - 1
			// 2+ 0 -1  = (0, 1)  dp[0][1] = max(p[0] + dp[1][1], dp[0,0] + p[1])
			// 2 + 1 -1 = (1, 2) dp[1][2] = max(p[1] + dp[2][2], dp[1,1] + p[2])
			fmt.Println(i, j)
			//过滤掉超了的
			if i > n-1 || j > n-1 {
				continue
			}
			left := piles[i] + dp[i+1][j].Sec
			right := dp[i][j-1].Sec + piles[j]
			if left > right {
				dp[i][j].Fir = left
				dp[i][j].Sec = dp[i+1][j].Fir
			} else {
				dp[i][j].Fir = right
				dp[i][j].Sec =  dp[i][j-1].Fir
			}
		}
	}
	for _,rows := range dp {
		for _, v := range rows  {
			fmt.Printf("%v\t",v)
		}
		fmt.Println()
	}
}
func juli01(d []Line, m int) int  {
	begin, end, sum  := 0, 0,0
	judge:=false
	for i := 0; i < len(d); i++ {
		tt :=false
		//fmt.Println(d[i])
		for i < len(d) && d[i].l <= begin+1  {
			if end < d[i].r {
				end = d[i].r
			}
			tt= true
			i++
		}
		begin = end
		if  i < len(d) {
			fmt.Printf("$%v", d[i])
		}
		sum++
		if tt {
			i--
		}
		if(begin >= m){
			judge=true;
			break;
		}
	}
	if judge {
		return sum
	}else {
		return -1
	}
}

func juli(d []Line, m int) int {
	right := d[0].r
	z := right
	ans := 0
	fmt.Printf("\n %v", d[0])
	for i:=1; i < len(d); i++  {
		if d[i].l > right+1 {
			ans++
			fmt.Printf("%v", d[i])
			right = z
		}
		if d[i].l <= right+1 {
			//更新z
			if z < d[i].r {
				z = d[i].r
			}
			if d[i].r == m {
				fmt.Printf("$ %v", d[i])
				ans++
				right = m
				break
			}
		}
	}
	if right == m {
		return ans
	}
	return -1
}
