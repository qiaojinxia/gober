package gbnet

import (
	"fmt"
	"gober/gbinterface"
	"gober/utils"
)

type MsgHandler struct {
	//用于存放 id 路由
	reflectMap map[uint32]gbinterface.IRouter
	TaskQueue []chan gbinterface.IRequest
	WorkPoolSize uint32
}

func NewMsgHandler() *MsgHandler{
	return &MsgHandler{
		reflectMap:make(map[uint32]gbinterface.IRouter,0),
		TaskQueue:make([]chan gbinterface.IRequest,utils.ConfigObject.WorkPoolSize),
		WorkPoolSize:utils.ConfigObject.WorkPoolSize,
	}
}

//添加待处理的方法 消息id 和 处理函数
func (mh *MsgHandler) AddRouter(msgid uint32,router gbinterface.IRouter){
	if ok := mh.reflectMap[msgid]; ok == nil{
		mh.reflectMap[msgid] = router
	}else{
		fmt.Println("handler func alerday exists!")
	}

}

func (mh *MsgHandler) StartWorkPool(){

		for i:= 0;i<int(mh.WorkPoolSize);i++ {

			mh.TaskQueue[i] = make(chan gbinterface.IRequest,utils.ConfigObject.MaxWorkTaskLen)

			go mh.startOneWorker(i,mh.TaskQueue[i])


		}

}
func (mh *MsgHandler) startOneWorker(worknum int,requestchan chan gbinterface.IRequest) {
	fmt.Printf("Work Pool Num : %d Start!\n",worknum)
	for{

		select{

			case req := <- requestchan:
				mh.DoRouter(req)
		}
	}

}



func (mh *MsgHandler) AddReqToQueue(requset gbinterface.IRequest){

	workId := requset.GetMessageID() % mh.WorkPoolSize

	fmt.Printf("Add Request To WorkPool Num: %d  \n",workId)

	mh.TaskQueue[workId] <- requset


}
//消息处理

func (mh *MsgHandler) DoRouter(request gbinterface.IRequest){
	msgid := request.GetMessageID()
	if hd := mh.reflectMap[msgid];  hd != nil{
		hd.PreHandle(request)
		hd.Handler(request)
		hd.PostHandle(request)
	}else{
		fmt.Println("have no handler get!")
	}
}

