package main

import "sync"

var c = make(chan int, 0)
var a string
func f() {
	a = "hello, world"
	<-c
}

var once sync.Once
func setup() {
	a = "hello, world"
	done = true
}
func doprint() {
	once.Do(setup)
	print(a)
}
func twoprint() {
	go doprint()
	go doprint()
}
var done bool
func main() {
	//go f()
	//c <- 0
	//print(a)
	//
	//twoprint()

	go setup()
	for !done {
	}
	print(a)
}