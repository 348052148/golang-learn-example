package main

import (
	"fmt"
	"reflect"
)

type A interface {
	Say()
}

type B struct {
}

func (b *B) Say() {
	fmt.Println("Say: HELLO")
}

func MapTo(val interface{}, ifacePtr interface{}) reflect.Value {
	if reflect.TypeOf(ifacePtr).Kind() != reflect.Interface {
		fmt.Errorf("%s", "ERROR NOT INTERFACE")
	}
	return reflect.ValueOf(val)
}

func main() {
	//var i int32
	t := reflect.TypeOf((*A)(nil))
	//fmt.Println(reflect.TypeOf(&B{}).Kind())
	if t.Kind() == reflect.Ptr {
		fmt.Println("PTR")
		t = t.Elem()
	}
	if t.Kind() == reflect.Interface {
		fmt.Println("Interface")
	}
	fmt.Println(t)
	v := MapTo(&B{}, (*A)(nil))
}
