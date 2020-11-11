package main

import (
	"DataCerPlatfomr/blockchain"
	"DataCerPlatfomr/database"
	_ "DataCerPlatfomr/routers"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

)



func main() {


	//1.创世区块
	bc := blockchain.NewBlockChain()//封装
	bc.SaveData([]byte("区块链学院2019"))
	blocks,err := bc.QueryAllBlocks()
	if err!=nil {
		fmt.Println(err.Error())

		return
	}
	fmt.Println("hello world")
	//blocks是一个切片
	for index,block := range blocks  {
		fmt.Println("区块高度:%d,区块hash:%x,PrevHash:",index,block.Height,block.PrevHash)
	}
	fmt.Println("创世区块的哈希值:%x\n",bc.LastHash)
	block1,err := bc.SaveData([]byte("用户要保存到区块的的数据"))
	if err!= nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("区块高度:%d\n",block1.Height)
	fmt.Println("区块的hash值:%d\n",block1.Hash)
	fmt.Println("区块前一个hash值:%d\n",block1.PrevHash)


	return


	//连接数据库
	//_,err = database.OpenDb()
	//if err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println(database.Db == nil)

	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

