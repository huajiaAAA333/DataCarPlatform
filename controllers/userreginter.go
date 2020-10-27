package controllers

import (
	"DataCerPlatfomr/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}


//注册post请求



func (R *RegisterController) Get() {

	R.TplName = "zuce.html"
}

func (r *RegisterController) Post() {
	//第一步解析请求数据
	var user models.User
	err :=r.ParseForm(&user)
	if err !=nil {
		//返回错误页面
		r.Ctx.WriteString("抱歉,小脑袋瓜停顿了一下,请重试QAQ")
		return
	}
	//var User models.User
	//r.ParseForm(&User)
	//if err!=nil {
	//	r.Ctx.WriteString("抱歉,数据解析失败,请重试!")
	//return
	//}


	//第二步保存用户信息到数据库
	user.SaveUser(&user)








	//第三步返回前段结果(成功跳登入页面,失败弹出错误信息)
	if err!= nil {
		r.Ctx.WriteString("抱歉,用户注册失败,请重试!")
		return

	}
	//用户注册成功
	//tpl:模板template
	r.TplName = "dengru.html"

}


