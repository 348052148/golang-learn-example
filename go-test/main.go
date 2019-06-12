package main

import (
	"fmt"
	"go-test/server"
	"net"
)

func main() {
	fmt.Println("SERVER")
	serverConn, _ := net.Listen("tcp", "127.0.0.1:3000")
	// message := make(chan []byte)
	serverManager := server.NewServerManager()
	for {
		clientConn, _ := serverConn.Accept()
		serverManager.AddConn(server.NewMessageConn(clientConn))

		serverManager.Run()

		fmt.Println(serverManager.ConnList)
	}

}
