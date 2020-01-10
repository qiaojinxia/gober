package gbnet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T){

	listen,err := net.Listen("tcp","127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err",err)
		return
	}

	go func() {
		for {
			conn,err := listen.Accept()
			if err != nil{
				fmt.Println("server accept error",err)
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte,dp.GetHeanLen())
					//读满容器
					_,err := io.ReadFull(conn,headData)
					if err != nil{
						fmt.Println("read head error")
						return
					}
					msgHead,err := dp.DecodePack(headData)
					if err != nil{
						fmt.Println("server unpack err",err)
						return
					}
					if msgHead.GetMessgLen() > 0 {

						msg := msgHead.(*Message)

						msg.Data = make([]byte,msgHead.GetMessgLen())

						_,err := io.ReadFull(conn,msg.Data)

						if err != nil{

							panic(err)
						}

						fmt.Println("Receiver MsgID: ",msg.Id,"dataLen = ",msg.DataLen,"data = " ,string(msg.Data))

					}
				}

			}(conn)
		}



	}()



	//client

	conn,err := net.Dial("tcp","127.0.0.1:7777")
	if err != nil{
		fmt.Println("connection error",err)
		return
	}
	dp := NewDataPack()

	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'1','3','5','6'},
	}

	senddata1,err := dp.CodePack(msg1)

	if err != nil{
		fmt.Println("client pack msg1 error",err)
		return
	}
	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'h','e','l','l','o','g','o'},
	}
	senddata2,err := dp.CodePack(msg2)
	if err != nil{
		fmt.Println("client pack msg1 error",err)
		return
	}


	senddata1 = append(senddata1, senddata2...)

	conn.Write(senddata1)
	select {}


}
