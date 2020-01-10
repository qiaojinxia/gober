package gbinterface

type IServer interface {
	Start()
	Server()
	Stop()
	AddRouter(IRouter)

}
