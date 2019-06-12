package main

import (
	"fmt"
	"net/http"
)

func Processor(seq chan int, wait chan struct{})  {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			fmt.Println("QUIT")
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for {
			num, ok:= <- seq
			if !ok {
				break
			}
			if num % prime != 0 {
				out <- num
			}
		}
		close(out)
		//for num := range seq {
		//	if num % prime != 0 {
		//		out <- num
		//	}
		//}
	}()
}
func main() {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < 100; num++ {
		origin <- num
	}
	close(origin)
	<-wait

	http.Header{}
}

//Wait
/*
  processor
			 Wait
  processor
 */
// Cancel 利用context
/*
		processor
 Cancel
		processor
 */

 //超时 利用context