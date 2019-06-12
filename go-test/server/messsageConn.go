package server

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type MessageConn struct {
	id            int
	conn          net.Conn
	serverManager *ServerManager
	heartMessage  chan bool
	timeOut       int
}

func NewMessageConn(conn net.Conn) *MessageConn {
	a := new(MessageConn)
	a.id = rand.Int()
	a.conn = conn
	a.timeOut = 10
	a.heartMessage = make(chan bool)
	return a
}

func (mConn *MessageConn) HeartBeating() {
	for {
		select {
		case v := <-mConn.heartMessage:
			mConn.conn.SetDeadline(time.Now().Add(time.Duration(mConn.timeOut) * time.Second))
			fmt.Println("XINTAO" + "ID:" + strconv.Itoa(mConn.id) + strconv.FormatBool(v))
		case <-time.After(time.Second * 10):
			fmt.Println("close conn")
			mConn.Close()
			return
		}
	}
}

func (mConn *MessageConn) Close() {
	mConn.serverManager.DelConn(mConn)
	mConn.conn.Close()
}
