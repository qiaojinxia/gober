package gbinterface

import "net"

type IConnection interface{
	Start()
	Stop()
	GetTcpConnection() *net.TCPConn
	GetConnId() uint32
	RemoteAddr() net.Addr
	SendMsg(uint32,[]byte) error
}

type HandlerFunc func (*net.TCPConn,[]byte,int) error