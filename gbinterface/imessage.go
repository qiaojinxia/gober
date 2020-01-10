package gbinterface

type IMessage interface {

	GetMsgId() uint32

	GetMessgLen() uint32

	GetData() []byte

	SetMsgId(uint32)

	SetData([]byte)

	SetDataLen(uint32)

}
