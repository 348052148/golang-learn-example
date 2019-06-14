package creates

import "fmt"

type Object int

func NewObjectPool(total int) chan *Object {
	objChan := make(chan *Object, total)
	for i := 0; i < total; i++ {
		objChan <- new(Object)
	}
	return objChan
}

func Use()  {
	objPool := NewObjectPool(2)
	for {
		select {
		case obj := <- objPool:
			//obj doShan
			*obj += 1
			fmt.Println(*obj)
			objPool <- obj
		default:

		}
	}
}