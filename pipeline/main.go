package main

import (
	"fmt"
	"sync"
)
var i int = 0
func main() {
	done := make(chan interface{})
	ch := take(done,reaptFn(done, func() interface{} {
		i++
		return i
	}),10)
	//Tee
	//o1, o2 := tee(done, ch)
	//
	//for v := range o2 {
	//	fmt.Printf("V1: %d, V2: %d \n",v, <-o1)
	//}
	merge := make(chan interface{})
	chs := fanOut(done, ch, 5)
	for _,ch := range chs {
		merge = fanIn(done, merge, add(done,ch, 10))
	}

	for v := range merge {
		fmt.Println(v)
	}

	//for v := range ch {
	//	fmt.Println(v)
	//}
	close(done)
	fmt.Println("END")
}

func add(done chan interface{}, ich chan interface{}, val interface{}) chan interface{} {
	och := make(chan interface{})
	go func() {
		defer close(och)
		for v := range ich {
			select {
			case och <- v.(int) + val.(int):
			case <-done:
				return
			}
		}
	}()
	return och
}

func reapt(done chan interface{},nums ...interface{}) <-chan interface{}{
	och := make(chan interface{})
	go func() {
		defer close(och)
		for {
			for _, num := range nums {
				select {
				case och<-num:
				case <-done:
					return
				}
			}
		}
	}()
	return och
}

func reaptFn(done chan interface{},fn func() interface{}) <-chan interface{} {
	och := make(chan interface{})
	go func() {
		defer close(och)
		for {
			select {
			case och <- fn():
				
			case <-done:
				return 
			}
		}
	}()
	return och
}

func take(done chan interface{},ich <-chan interface{}, n int) <- chan interface{} {
	och := make(chan interface{})
	go func() {
		defer close(och)
		for i:=0 ; i<n ; i++  {
			select {
			case v := <-ich:
				och <- v
			case <-done:
				return
			}
		}
	}()
	return och
}

func tee(done chan interface{}, ich <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	o1, o2 := make(chan interface{}), make(chan interface{})
	go func() {
		defer close(o1)
		defer close(o2)
		var t1,t2 chan interface{}
		for v := range ich {
			t1,t2 = o1,o2
			select {
			case t1 <- v:
				//t1 = nil
			case t2 <- v:
				//t2 = nil
			case <-done:
				return
			}
		}
	}()
	return o1, o2
}

func fanOut(done chan interface{}, ich <-chan interface{},n int) []chan interface{} {
	var ochs []chan interface{}
	for i := 0; i < n; i++ {
		ochs = append(ochs, make(chan interface{}))
	}
	for _,och := range ochs {
		go func(c chan interface{}) {
			defer close(c)
			for v := range ich {
				select {
				case c<-v:
				case <-done:
					return
				}
			}
		}(och)
	}
	return ochs
}
//遵循创建它的负责销毁它
func fanIn(done chan interface{}, ichs ...chan interface{}) chan interface{} {
	och := make(chan interface{})
	wg := &sync.WaitGroup{}
	for _, ch := range ichs {
		wg.Add(1)
		go func(c chan interface{}) {
			wg.Done()
			//defer close(och)
			for v := range c {
				select {
				case och <- v:
				case <-done:
					return
				}
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(och)
	}()
	return och
}