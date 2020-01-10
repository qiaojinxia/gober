package gbinterface

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	GetMessageID() uint32
	GetMessageLen() uint32
}
