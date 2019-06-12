package main

import "fmt"

type T1 struct {
	name string
}
type T2 struct {
	name string
}
type I1 interface {
	M1()
}
type I2 interface {
	M1()
} 
type T struct {
	name string
	age int
}

func main() {
	vs := []interface{}{T2(T1{"foo"}),string(322),[]byte("abc")}
	for _,v := range vs {
		fmt.Printf("%v %T\n", v, v)
	}
	v1 := struct {
		name string
	}{}
	fmt.Printf("%T\n", v1)
	var v2 T1
	v2 = v1
	fmt.Printf("%T \n", v2)
}
