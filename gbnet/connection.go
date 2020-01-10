package gbnet

import (
	"errors"
	"fmt"
	"gober/gbinterface"
	"gober/utils"
	"io"
	"net"
)

type Connection struct {

	Conn *net.TCPConn

	ConnID uint32

	isClosed bool
	//handler
	ExitChan chan bool

	msgHandler gbinterface.IMsgHandler
	//将要发送的数据 存储的管道
	SendDataChan chan []byte


}

func NewConnection(conn *net.TCPConn,connID uint32,msghandler gbinterface.IMsgHandler) *Connection{
	
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		msgHandler: msghandler,
		ExitChan:  make(chan bool ,1),
		SendDataChan:make(chan []byte,256),
	}

}



func (c *Connection) Start(){
	fmt.Println("Conn Start() ... ConnID = ",c.ConnID)
	go c.connReader()
	go c.connWrite()

}

//读取管道内容并发送
func (c *Connection) connWrite(){
	fmt.Println("[Gober Write Gortine is running!]")
	defer fmt.Println("[Gober Connection Write is Closed! ConnID = ",c.ConnID,"Remote addr :",c.RemoteAddr().String() +"]")
	for{
		select {
		case data := <-c.SendDataChan:
			if _,err :=c.Conn.Write(data) ;err != nil{
				fmt.Println("[Gober Send Data error",err,"Conn Write exit]")
				return
			}
		case  <- c.ExitChan:
			return

		}
	}

}

func (c *Connection) connReader(){
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("[Gober Connection Reader is Exited! ConnID = ",c.ConnID,"Remote addr :",c.RemoteAddr().String()+"]")
	defer c.Stop()
	for{
		dp := NewDataPack()
		buff := make([]byte,dp.GetHeanLen())
		if _,err :=io.ReadFull(c.GetTcpConnection(),buff); err != nil{
			fmt.Println("read conn  error",err)
			break
		}
		msgHead,err:= dp.DecodePack(buff)

		if err != nil{
			fmt.Println("read msghead error",err)
			break
		}
		if msgHead.GetMessgLen() > 0 {
			msg := msgHead.(*Message)
			dbuf := make([]byte,msgHead.GetMessgLen())


			_,err := io.ReadFull(c.GetTcpConnection(),dbuf)
			if err != nil  {
				fmt.Println("read msgdata  error",err)
				break
			}

			msg.Data = dbuf

			req := Request{
				conn: c,
				message:msg,
			}


			if utils.ConfigObject.MaxWorkTaskLen > 1{
				c.msgHandler.AddReqToQueue(&req)
			}else{
				c.msgHandler.DoRouter(&req)
			}
		}
	}

}

func (c *Connection) Stop(){
	fmt.Println("Conn Stop()....ConnID = ",c.ConnID)
	if c.isClosed  {
		return
	}
	c.isClosed = true

	c.Conn.Close()

	c.ExitChan <- true

	close(c.ExitChan)

}
func (c *Connection) GetTcpConnection() *net.TCPConn{
	return c.Conn
}
func (c *Connection) GetConnId() uint32{
	return c.ConnID
}
func (c *Connection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}
func (c *Connection) end(data []byte) error{
	return nil
}

func (c *Connection) SendMsg(msgid uint32,data []byte) error{
	if c.isClosed {
		return errors.New("Conn Is Closed Can't Send MSG!")
	}
	dp := NewDataPack()
	msgbean := NewMsg(msgid,data)
	msg, err := dp.CodePack(msgbean)
	if err != nil {
		fmt.Println("unpack msg error",err)
	}
	//交给管道发送消息
	c.SendDataChan <- msg
	return nil
}