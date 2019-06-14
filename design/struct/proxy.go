package _struct

import "fmt"

type Say interface {
	Say()
} 

type A struct {
}

func (a A)Say()  {
	fmt.Println("A SAY")
}
//和装饰器模式不同在于，装饰器模式
type Proxy struct {
	a Say
}

func NewProxy(a Say) Proxy  {
	return Proxy{
		a:a,
	}
}

func (p Proxy)Say()  {
	p.a.Say()
}

