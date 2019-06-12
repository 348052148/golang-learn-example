package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Printf("%v \n",files)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _,f := range files {
			file,err := os.Open(f)
			if err != nil {
				panic(err)
			}
			countLines(file, counts)
		}
	}
	for line, n := range counts {
		fmt.Printf(" %s, count: %v \n", line, n)
	}
}

func countLines(reader *os.File, counts map[string]int)  {
	input := bufio.NewScanner(reader)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
}
