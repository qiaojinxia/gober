package gbnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"gober/gbinterface"
	"gober/utils"
)

type DataPack struct {}


func NewDataPack() *DataPack{
	return &DataPack{}
}

func (dp *DataPack)GetHeanLen() uint32{
	//包含
	return 8
}
func (dp *DataPack) CodePack(msg gbinterface.IMessage)([]byte,error){

	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMessgLen()) ; err != nil{
		return nil,err
	}
	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgId()) ; err != nil{
		return nil,err
	}
	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetData()) ; err != nil{
		return nil,err
	}
	return dataBuff.Bytes(),nil
}
func (dp *DataPack)  DecodePack(binaryData []byte)(gbinterface.IMessage,error){
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}
	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen); err!= nil{
		return nil,err
	}

	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.Id); err!= nil{
		return nil,err
	}

	if (utils.ConfigObject.MaxPackageSize >0 && msg.DataLen > utils.ConfigObject.MaxPackageSize){
		return nil,errors.New(fmt.Sprintf("TOO LAGE DATA SIZE  %d RECEIVE! ",msg.DataLen))
	}


	return msg,nil

}