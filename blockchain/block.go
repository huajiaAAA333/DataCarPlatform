package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

//定义区块结构体，用于表示区块
type Block struct {
	Height int64 //区块的高度，第某个区块
	TimeStamp int64 //区块产生的时间戳
	PrevHash []byte  //上个区块的哈希值
	Data []byte //数据
	Hash []byte //当前区块的哈希
	Version string  //版本号
	Nonce int64 //区块对应的nonce值


}

//创建一个新区快
func NewBlock(height int64,prevHash []byte,data []byte) Block {
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		Version:   "0x01",
	}

	//找nonce值，通过工作量证明算法寻找
	pow := NewPow(block)
	hash,nonce := pow.Run()
	block.Nonce = nonce
	block.Hash = hash



	////1.将block结构体数据转换为【】byte类型
	//heightBytes,_ := utils.Int64ToByte(block.Height)
	//timeStampBytes,_ :=utils.Int64ToByte(block.TimeStamp)
	//versionBytes := utils.StringToBytes(block.Version)
	//nonceBytes,_ := utils.Int64ToByte(block.Nonce)
	//
	//
	//
	//var blockBytes []byte
	////bytes.Join 拼接
	//bytes.Join([][]byte{
	//	heightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//	nonceBytes,
	//},[]byte{})
	//
	////调用hash计算，对区块进行激素按
	//block.Hash = utils.SHA256HashBlock(blockBytes)
	//	//计算哈希
	//	//挖矿竞争
	//	//核心公式; SHA256（区块A + n ）《 系统B


	   return  block
}

//创建创世区块
func CreateGenesisBlock() Block  {
	genesisblock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return  genesisblock
}


//对区块进行序列化
func (b Block) Serialize() ([]byte) {
	 buff := new( bytes.Buffer)//缓冲区
	 encoder := gob.NewEncoder(buff)
	 encoder.Encode(b)//将区块b放入到序列化编码器中
	return buff.Bytes()
}

//区块反序列化操作
func DeSerialize(data []byte) (*Block ,error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err!=nil {
		return nil,err
	}
	return  &block,nil
}

