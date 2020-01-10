package main

import (
	"fmt"
	"net"
	"time"
)

//模拟客户端
func main(){


	fmt.Println("client start")
	//直接连接远程服务器，得到一个conn连接
	time.Sleep(1 * time.Second)

	conn ,err := net.Dial("tcp","0.0.0.0:8081")

	if err != nil {
		fmt.Println("client start err ，exit!")
		return
	}
	sequence := 0
	for{
		conn.Write([]byte("hello this is gober message num:" + fmt.Sprintf("%d",sequence)))
		sequence ++
		if err != nil{
			fmt.Println("senf data err",err )
			return
		}
		buf := make([]byte,512)
		cin,err := conn.Read(buf)
		if err != nil{
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("receviver from sever data : %s cin: %d \n" ,buf,cin)
		time.Sleep(1 * time.Second)

	}



}