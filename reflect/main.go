package main

import (
	"reflect"
	"fmt"
)

type a interface {
	fn()
}

type TestStruct struct {
	Field string `json:"ffinfo"`
	ffinfo FieldInfo
}
func (t *TestStruct)fn() {
	return 
}
func (t TestStruct)Say()  {
	fmt.Println("Say Hello")
}
type FieldInfo struct {

}

func main() {
	typ := reflect.TypeOf(&TestStruct{})
	val := reflect.ValueOf(&TestStruct{})
	fmt.Println(typ.Elem().Field(0))
	fmt.Println("Fields",typ.Elem().NumField())
	fmt.Println("Methods",typ.Elem().NumMethod())
	method, ok := typ.Elem().MethodByName("Say")
	//直接调用Func
	if ok {
		fmt.Println("typeFunc",method.Func.Call([]reflect.Value{val.Elem()}))
	}
	val.Elem().Field(0).SetString("abc")
	fmt.Println(val.Elem().Field(0).String())
	//调用结构附带类
	fmt.Println(val.Elem().Method(0).Call([]reflect.Value{}))
}
