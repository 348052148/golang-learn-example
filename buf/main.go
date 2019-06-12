package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	//resp, err := http.Get("http://192.168.0.173/omsv2/sync/instruction/search_list?count=9999")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//reader := bufio.NewReader(resp.Body)
	//
	//for  {
	//	if r,ok  := reader.ReadString(10); ok == nil {
	//		fmt.Println(r)
	//	}else {
	//		break
	//	}
	//
	//}
	h := md5.New()
	fmt.Fprint(h, "12321312");
	fmt.Fprint(h, "12321312");
	fmt.Printf("%x",h.Sum(nil))
}
