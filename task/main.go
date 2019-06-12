package main

import (
	"fmt"
	"net"
	"task/task"
)

func main() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {

	}
	conn.Write([]byte("nihao"))
	reqTask := task.NewRequestTask()
	reqTask.Run()
	var data []byte
	data = make([]byte, 1024)
	for {
		len, err := conn.Read(data)
		if err != nil {
			return
		}
		fmt.Println(len)
	}
}
