package utils

import (
	"bytes"
	"encoding/binary"
)

//将一个int64转化为【】byte
func Int64ToByte(num int64)([]byte,error)  {

	//buffer：缓冲区
	buff := new(bytes.Buffer)//通过new实例化一个缓冲区
	//buff.Write()//通过一系列write方法向缓冲区写入数据
	//buff.Bytes()//通过bytes方法从缓冲区当中获取数据
	/*
	大端位序排列：binary.BigEndian
	小端位序排列：binary.LittleEndian
	 */
	binary.Write(buff,binary.LittleEndian,num)

	//从缓冲区去读数据
	return  buff.Bytes(),nil
}


//将字符串转化为字节切片
func StringToBytes(data string)  []byte{
	return []byte(data)
}