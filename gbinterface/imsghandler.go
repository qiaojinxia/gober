package gbinterface

type IMsgHandler interface {
	//添加待处理的方法 消息id 和 处理函数
	AddRouter(uint32,IRouter)
	//消息处理
	DoRouter(IRequest)
	//启动工作池
	StartWorkPool()
	//给chan增加request
	AddReqToQueue(IRequest)

}
