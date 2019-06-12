package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan string)
	count := len(os.Args[1:])
	for _,url := range os.Args[1:] {
		go fetch(url, ch)
	}
	i := 0
	for str := range ch {
		fmt.Println(str)
		i++
		if i >= count {
			break
		}
	}
	fmt.Printf("total_rquest_time: %0.2f \n", time.Since(startTime).Seconds())
}

func fetch(url string, ch chan<- string) {
	stime := time.Now()
	resp,err := http.Get(url)
	if err!=nil {
		panic(err)
	}
	rsize, err := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if err!=nil {
		panic(err)
		return
	}
	ch <-fmt.Sprintf("time: %0.2f ,size:%d, url:%s \n",time.Since(stime).Seconds(),rsize, url)
}
