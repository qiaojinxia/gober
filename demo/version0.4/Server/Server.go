package main

import (
	"fmt"
	"gober/gbinterface"
	"gober/gbnet"
)

type pingRouter struct {

}
func (this *pingRouter) PreHandle(request gbinterface.IRequest){
	fmt.Println("call Router perhandler")
	_,err := request.GetConnection().GetTcpConnection().Write([]byte("before ping ....\n"))
	if err != nil{
		fmt.Println("call back ping error!")
	}

}

func (this *pingRouter) Handler(request gbinterface.IRequest){
	fmt.Println("call Router handler")
	_,err := request.GetConnection().GetTcpConnection().Write([]byte("ping ....\n"))
	if err != nil{
		fmt.Println("call back ping error!")
	}
}


func (this *pingRouter) PostHandle(request gbinterface.IRequest){
	fmt.Println("call Router afterhandler")
	_,err := request.GetConnection().GetTcpConnection().Write([]byte("afterping ....\n"))
	if err != nil{
		fmt.Println("call back ping error!")
	}
}

func main(){
	s := gbnet.NewServer("caomao")
	s.AddRouter(&pingRouter{})
	s.Server()
}

