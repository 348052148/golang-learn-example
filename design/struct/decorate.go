package _struct

import (
	"log"
	"fmt"
)

type Func func(int, int) int

func LogDecorate(fn Func) Func  {

	return func(i int, i2 int) int {
		log.Print("logging")
		return fn(i, i2)
	}
}

func Add(a, b int) int  {
	return a + b
}

func Use()  {
	f := LogDecorate(Add)
	fmt.Println(f(1,2))
}