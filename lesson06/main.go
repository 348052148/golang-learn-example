package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _,url := range os.Args[1:] {
		resp,err := http.Get(url)
		if err!=nil {
			panic(err)
			os.Exit(1)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err!=nil {
			panic(err)
			os.Exit(1)
		}
		fmt.Printf("%v \n", string(data))
	}
}
