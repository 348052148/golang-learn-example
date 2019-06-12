package main

import (
	"fmt"
	"io"
	"os"
)

type a interface {
	echo()
}
type inter struct {

}

func (i inter)echo()  {

}

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	var wr io.Writer
	k := wr.(*os.File)
	fmt.Println(k)
	fmt.Println(s)
}
