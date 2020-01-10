package main

import (
	"fmt"
	"gober/gbnet"
	"io"
	"net"
	"time"
)

//模拟客户端
func main(){


	fmt.Println("client start")
	//直接连接远程服务器，得到一个conn连接
	time.Sleep(1 * time.Second)

	conn ,err := net.Dial("tcp","0.0.0.0:8999")

	if err != nil {
		fmt.Println("client start err ，exit!")
		return
	}
	sequence := 0
	for{
		dp := gbnet.NewDataPack()
		sequence ++
		content :=[]byte{'h','e','l','l','o','!'}
		msgbean :=gbnet.NewMsg(uint32(sequence),content)
		msg,err := dp.CodePack(msgbean)
		if err != nil{
			fmt.Println("undcode datapack err",err )
		}
		conn.Write(msg)
		if err != nil{
			fmt.Println("senf data err",err )
			return
		}


		headData := make([]byte,dp.GetHeanLen())
		//读满容器
		_,err = io.ReadFull(conn,headData)
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

			msg := msgHead.(*gbnet.Message)

			msg.Data = make([]byte,msgHead.GetMessgLen())

			_,err := io.ReadFull(conn,msg.Data)

			if err != nil{

				panic(err)
			}

			fmt.Println("Receiver MsgID: ",msg.Id,"dataLen = ",msg.DataLen,"data = " ,string(msg.Data))

		}
		if err != nil{
			fmt.Println("read buf error")
			return
		}
		time.Sleep(10 * time.Second)

	}



}