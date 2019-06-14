package behavioral

import "fmt"

type Observer interface {
	HandleNotify() error
}

type ClickEvent struct {

}

func (c ClickEvent)HandleNotify() error  {
	fmt.Println("CLICK EVENT")
	return nil
}

type OpenEvent struct {
}

func (c OpenEvent)HandleNotify() error  {
	fmt.Println("OPEN EVENT")
	return nil
}

type ObserverListener struct {
	observers map[Observer]interface{}
}

func (lst ObserverListener)addListener(l Observer)  {
	lst.observers[l] =interface{}(nil)
}

func (lst ObserverListener)Notify()  {
	for ob, _ := range lst.observers {
		ob.HandleNotify()
	}
}
