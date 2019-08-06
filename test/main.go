package main

import (
	"fmt"
	"unsafe"
)
type T struct {
	name string
}
type I interface {
	say()
}
type A struct {
	name string
}

func (a A)say()  {
}
func main() {
	//builder := NewCarBuilrer()
	//
	//fmt.Println(unsafe.Pointer(&builder))
	//b := builder.Color(1).Wheels(2).TopSpeed(300)
	//fmt.Println(unsafe.Pointer(&b))
	//car := b.Build()
	//fmt.Println(car)
	//
	//var u I
	//a := A{}
	//u = a
	//var u1 I
	//u1 = u
	//inspect(&u,&a)
	//inspect(&u1,&a)
	//wg := &sync.WaitGroup{}
	//wg.Add(2)
	fmt.Println(s(26))
}

func s(n int) string  {
	if n <= 26 {
		return fmt.Sprintf("%c", 65 + (n-1) % 26 )
	}
	return s((n-1) / 26 ) + fmt.Sprintf("%c", 65 + (n-1) % 26 )
}

func inspect(n *I, u *A) {
	word := uintptr(unsafe.Pointer(n)) + uintptr(unsafe.Sizeof(&u))
	value := (**I)(unsafe.Pointer(word))
	fmt.Printf("Addr User: %p  Word Value: %p  Ptr Value: %v\n", u, *value, **value)
}
func change(arr []int,offst int) int {
	for i,v := range arr[offst:] {
		if i > 2 {
			offst +=2
			break
		}
		fmt.Println(v)
	}
	return offst
}


type (
	Color int
	Wheels int
	Speed int
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build()CarInterface
}

type CarBuilder struct {
	Colors Color
	Wheelss Wheels
	Speed Speed
}

func NewCarBuilrer() CarBuilder {
	return CarBuilder{}
}
func (b CarBuilder)Color( c Color) Builder{
	b.Colors = c
	return b
}
func (b CarBuilder)Wheels(w Wheels) Builder {
	b.Wheelss = w
	return b
}
func (b CarBuilder)TopSpeed(s Speed) Builder {
	b.Speed = s
	return b
}
func (b  CarBuilder)Build() CarInterface {
	return &Car{
		Colors: b.Colors,
		Wheelss: b.Wheelss,
		Speed: b.Speed,
	}
}

type CarInterface interface {
	Drive() error
	Stop() error
}

type Car struct {
	Colors Color
	Wheelss Wheels
	Speed Speed
}

func (car * Car)Drive()error  {
	return nil
}
func (car * Car)Stop()error  {
	return nil
}