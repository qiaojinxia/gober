package gbnet

import (
	"fmt"
	"gober/gbinterface"
	"gober/utils"
	"net"
)

type Server struct {
	logo string
	//服务器名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听ip
	Ip string
	//服务器端口
	Port int
	//多路由处理器
	msgHandler gbinterface.IMsgHandler
}

func (s *Server) Start(){
	fmt.Printf("[Gober]Server Name : %s ,listener at IP  %s ,Port %d \n",utils.ConfigObject.Name,utils.ConfigObject.Host,utils.ConfigObject.TcpPort)
	//获取一个tcp的addr
	fmt.Printf("Accpet Client Conn from IP  %s ,Port %d \n",s.Ip,s.Port)
	s.msgHandler.StartWorkPool()
	mh := NewMsgHandler()
	mh.StartWorkPool()
	go func() {
		addr ,err := net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.Ip,s.Port))
		if err != nil{
			fmt.Println("resolve tcp address error \n",err)
			return
		}
		//监听服务器地址
		listener,err := net.ListenTCP(s.IPVersion,addr)
		if err != nil {
			fmt.Printf("listen to Port %s:%d error: %v \n" ,s.Ip,s.Port,err)

		}
		fmt.Printf("Start server success listening in: %s:%d \n",s.Ip,s.Port)

		for  {
			//阻塞等待客户端连接
			conn,err := listener.AcceptTCP()
			if err != nil{
				fmt.Printf("Accept err %v",err)
				continue
			}
			var cid uint32
			cid = 0

			dealConn := NewConnection(conn,cid,s.msgHandler)
			go dealConn.Start()
			cid ++

		}


	}()
}

func (s *Server) Server(){
	//启动 server
	s.Start()
	select {}
}


func (s *Server) AddRouter(msgid uint32,router gbinterface.IRouter){
	s.msgHandler.AddRouter(msgid,router)
}


func (s *Server) Stop(){

}


func NewServer(name string) *Server{
	s := &Server{
		logo:utils.ConfigObject.Logo,
		Name:utils.ConfigObject.Name,
		IPVersion: "tcp4",
		Ip:        utils.ConfigObject.Host,
		Port:      utils.ConfigObject.TcpPort,
		msgHandler: NewMsgHandler(),
	}

	fmt.Println(s.logo)
	if len(name) > 0  {
		s.Name = name
	}
	return s
}
