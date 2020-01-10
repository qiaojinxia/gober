package main

import (
	"fmt"
	"gober/gbinterface"
	"gober/gbnet"
)

type pingRouter struct {

}

func (this *pingRouter) PreHandle(request gbinterface.IRequest) {
	return
}

func (this *pingRouter) PostHandle(request gbinterface.IRequest) {
	return
}

func (this *pingRouter) Handler(request gbinterface.IRequest){
	fmt.Println("call handler action \n")
	request.GetMessageID()
	fmt.Printf("get id from client msgid :%d msglen:%d msgcontent : %s \n",request.GetMessageID(),request.GetMessageLen(),request.GetData())
	content :=[]byte{'c','a','o','b','o','y'}
	err := request.GetConnection().SendMsg(12,content)
	if err != nil{
		fmt.Println("send msg err!")
	}


}



func main(){
	s := gbnet.NewServer("caomao")
	s.AddRouter(&pingRouter{})
	s.Server()
}

