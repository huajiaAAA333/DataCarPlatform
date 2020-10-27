package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

//对一个字符串数据进行MD5哈希计算
func MD5HashString(data string) string {
	md5Hash := md5.New()

	md5Hash.Write([]byte(data))


	bytes := md5Hash.Sum(nil)

	return hex.EncodeToString(bytes)




}
//读取io流的数值
func SHA256HashReader(reader io.Reader) (string,error)  {
	hash256 := md5.New()
	readerBytes,err :=ioutil.ReadAll(reader)
	if err!=nil {
		return"",err
	}
	hash256.Write(readerBytes)
	hashBytes := hash256.Sum(nil)
	return hex.EncodeToString(hashBytes),nil

}


//
func MD5HashReader (reader io.Reader) (string,error) {
	md5Hash := md5.New()
	readerBytes,err:=ioutil.ReadAll(reader)
	//fmt.Println("读取到的文件：",readerBytes)
	if err!=nil {
		return "",err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil


}


//对区块数据进行哈希计算
func SHA256HashBlock(bs []byte ) []byte{
	////1.将block结构体数据转换为【】byte类型
	//heightBytes,_ := Int64ToByte(block.Height)
	//timeStampBytes,_ := Int64ToByte(block.TimeStamp)
	//versionBytes := StringToBytes(block.Version)
	//
	//var blockBytes []byte
	////bytes.Join 拼接
	//bytes.Join([][]byte{
	//	heightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//},[]byte{})




	//2.将转换后的byte字节切片输入wite方法r


	sha256Hash := sha256.New()
	sha256Hash.Write(bs)
	hash := sha256Hash.Sum(nil)
	return hash
}









