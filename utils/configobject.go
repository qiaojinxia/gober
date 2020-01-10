package utils

import (
	"encoding/json"
	"gober/gbinterface"
	"io/ioutil"
)

type ConfigObj struct {
	Logo string
	TcpServer gbinterface.IServer //全局server对象
	Host string//当前服务器主机监听IP
	TcpPort int//当前服务器主机监听端口号
	Name string//当前服务器名称

	Version string //当前软件版本号
	MaxConn int//最大连接数
	MaxPackageSize uint32 //数据包最大值

	WorkPoolSize uint32
	MaxWorkTaskLen uint32


}
var ConfigObject *ConfigObj

func init(){
	ConfigObject = &ConfigObj{
		Logo:"",
		TcpServer:nil,
		Host:           "192.168.1.0",
		TcpPort:        8999,
		Name:           "goberServer",
		Version:        "v0.1",
		MaxConn:        100,
		MaxPackageSize: 512,
		WorkPoolSize: 10,
		MaxWorkTaskLen: 1024,
	}
	ConfigObject.Reload()
}

func (g *ConfigObj) Reload(){
	data , err := ioutil.ReadFile("/Users/qiao/go/src/gober/demo/version0.4/Server/conf/gober.json")
	if err != nil{
		panic(err)
	}
	err =json.Unmarshal(data,&ConfigObject)
	logodata , err := ioutil.ReadFile("/Users/qiao/go/src/gober/utils/logo")
	ConfigObject.Logo = string(logodata)
	if err != nil{
		panic(err)
	}

}