package controllers

import (
	"DataCerPlatfomr/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
)

//注册 post请求

type Register struct {
	beego.Controller
}

//处理用户注册时json格式数据的post请求

func (c *Register) Post(){
	//body := c.Ctx.Request.Body
	dataByes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败,请重试")
		return
	}
	//json包解析
	var user models.User
	err  = json.Unmarshal(dataByes,&user)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return


	}
	////解析后得到user的数据保存到数据库中
	//userId,_,err := user.SaveUser(&user)
	//if err != nil {
	//	c.Ctx.WriteString("将解析后得到的user数据保存到数据库中时出错,请重试!")
	//	return
	//}
	//
	////反馈注册成功与否
	//c.Ctx.WriteString("注册成功!!!")





}


