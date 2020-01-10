package gbnet

import "gober/gbinterface"

type Request struct {
	conn gbinterface.IConnection
	message gbinterface.IMessage
}

func (r *Request) GetConnection() gbinterface.IConnection{
	return r.conn
}

func (r *Request) GetData() []byte{
	return r.message.GetData()
}

func (r *Request) GetMessageID() uint32{
	return r.message.GetMsgId()
}


func (r *Request) GetMessageLen() uint32{
	return r.message.GetMessgLen()
}


