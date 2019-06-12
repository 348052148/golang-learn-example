package server

import (
	"bytes"
	"fmt"
)

type ServerManager struct {
	ConnList map[int]*MessageConn
}

func NewServerManager() *ServerManager {
	serverManager := new(ServerManager)
	serverManager.ConnList = make(map[int]*MessageConn)
	return serverManager
}

func (server *ServerManager) AddConn(conn *MessageConn) {
	conn.serverManager = server
	server.ConnList[conn.id] = conn
}
func (server *ServerManager) DelConn(conn *MessageConn) {
	delete(server.ConnList, conn.id)
}

func (server *ServerManager) Run() {
	server.hanldReader()
}

func (server *ServerManager) hanldReader() {
	for _, mConn := range server.ConnList {
		go connHandle(mConn)
		go mConn.HeartBeating()
	}
}

func connHandle(mConn *MessageConn) {
	var data []byte
	data = make([]byte, 1024)
	buffer := new(bytes.Buffer)
	for {
		buffer.Reset()
		for {
			len, err := mConn.conn.Read(data)
			if err != nil {
				return
			}
			buffer.Write(data[:len])
			if len < 1024 {
				break
			}
		}
		fmt.Println("READ-VALUE" + string(buffer.Bytes()))
		mConn.heartMessage <- true
	}
}

func (server *ServerManager) heart() {
	for _, conn := range server.ConnList {
		conn.HeartBeating()
	}
}
