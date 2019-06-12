package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _,file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err!=nil {
			//panic(err)
			continue
		}
		for _,text := range strings.Split(string(data),"\n") {
			counts[text]++
		}
	}
	for line, n := range counts {
		fmt.Printf("%s,count: %v \n", line, n)
	}
}
