package gbinterface

type IDataPack interface {
	GetHeanLen() uint32
	CodePack(IMessage)([]byte,error)
	DecodePack([]byte)(IMessage,error)
}
