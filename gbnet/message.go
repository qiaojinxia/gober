package gbnet

import _ "gober/gbinterface"

type Message struct {
	Id uint32 //消息id
	DataLen uint32//消息长度
	Data []byte//消息内容
}

func NewMsg(msgid uint32,data []byte) *Message {
	msg := &Message{
		Id:      msgid,
		DataLen: uint32(len(data)),
		Data:    data,
	}
	return msg
}


func (m *Message) GetMsgId() uint32{
	return m.Id
}

func (m *Message)  GetMessgLen() uint32{
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message)  SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetData(data []byte){
	m.Data = data
}

func (m *Message) SetDataLen(len uint32){
	m.DataLen = len
}