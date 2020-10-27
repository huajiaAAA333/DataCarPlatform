package main

import (
	"DataCerPlatfomr/blockchain"
	"DataCerPlatfomr/database"
	_ "DataCerPlatfomr/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

)



func main() {

	block0 := blockchain.CreateGenesisBlock()//创建创世区块
	fmt.Println(block0)
	fmt.Println("当前block的哈希值：%x\n",block0.Hash)
	block1 := blockchain.NewBlock(
		block0.Height+1,
		block0.Hash,
		[]byte{})
	fmt.Println("block1的哈希：%x\n",block1.Hash)
	fmt.Println("block的prevhash：%x\n",block1.PrevHash)


	block0Bytes := block0.Serialize()
	deBlcok0,err := blockchain.DeSerialize(block0Bytes)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("反序列化操作后区块高度是:",deBlcok0.Height)
	fmt.Println("反序列化后区块的哈希:%x\n",block1.PrevHash)

	return


	//序列化marshal：将数据从内存当中的形式转化为可以持久化存储在硬盘上或在网络上传输的形式
	blockJson,_ := json.Marshal(block0)
	fmt.Println("通过json序列化以后的block：",string(blockJson))
	//只有序列化以后的对象才能进行传输，

	//反序列化unmarshal：将数据从文件中或者网络中读取，然后转化到计算机内存中的过程


	return


	//连接数据库
	_,err = database.OpenDb()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(database.Db == nil)

	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	//
	////定义config变量, 接收并赋值为全局配置变量
	//config := beego.AppConfig
	////获取配置选项
	//appName := config.String("appname")
	//fmt.Println("项目应用名称：",appName)
	//port,err  := config.Int("httpport")
	//if err != nil {
	//	//配置信息解析错误
	//	panic("项目配置信息解释错误,请查验后重试")
	//}
	//fmt.Println("应用监听端口:",port)
	//
	//driver := config.String("db_driver")//数据库驱动
	//dbUser := config.String("db_user")//数据库用户名
	//dbPassword := config.String("db_password")//密码
	//dbIP := config.String("db_ip")
	//dbName := config.String("db_name")
	//
	//db, err :=sql.Open(driver,dbUser+":"+dbPassword+"@tcp("+dbIP+")/"+dbName+"?charset=utf8")
	//if err != nil {
	//	fmt.Println("错误:",err)
	//	panic("数据连接打开失败，请重试")
	//}
	//fmt.Println(db)
	beego.Run()
}

