package controllers

import (
	"DataCerPlatfomr/blockchain"
	"DataCerPlatfomr/models"
	"DataCerPlatfomr/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetailController struct {
	beego.Controller
}


/*
该get方法用于处理留恋其的get请求,网查看证书详情页面跳转
 */

func (c *CertDetailController) Get()  {
	//1.先解析和接受前段页面传递的数据
	cert_id := c.GetString("cert_id")
	//2.到区块链上查询区块数据
	block,err := blockchain.CHAIN.QueryBlockByCertId(cert_id)
	if err!=nil {
		c.Ctx.WriteString("抱歉,查询连上数据遇到错误请重试")
		return
	}
	if block == nil {//遍历整条区块链,但是未查到数据
		c.Ctx.WriteString("抱歉,未查询到连上数据")
		return

	}
	fmt.Println("查询到的区块的高度:",block.Height)
	//反序列化

	certRecord,err := models.DeserializeCertRecord(block.Data)
	certRecord.CertIdHex = strings.ToUpper(string(certRecord.CertId))
	certRecord.CertHashHex = string(certRecord.CertHash)
	certRecord.CertTimeFormat = utils.TimeFormat(certRecord.CertTime,utils.TIME_FORMAT_ONE)
	//结构体
	c.Data["CertRecord"] = certRecord

	//3.跳转证书详情页面
	c.TplName = "cert_detail.html"
}