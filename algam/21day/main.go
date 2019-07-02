package main

import (
	"sort"
	"fmt"
)

func main() {
	for v := range Merge(InMemSort(ArraySource(2,1,4)),InMemSort(ArraySource(3,9,7))) {
		fmt.Printf("%d , ", v)
	}
}

func ArraySource(n ...int) chan int {
	out := make(chan int)
	go func() {
		for _,v := range n {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		a := make([]int,0)
		for v := range in {
			a = append(a, v)
		}
		sort.Ints(a)
		for _,v := range a {
			out <- v
		}
	}()
	return out
}

func Merge(left, right chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		v1, ok1 := <-left
		v2, ok2 := <-right
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-left
			} else {
				out <- v2
				v2, ok2 = <-right
			}
		}
	}()
	return out
}