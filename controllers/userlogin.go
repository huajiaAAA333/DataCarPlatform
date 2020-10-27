package controllers

import (
	"DataCerPlatfomr/models"
	"fmt"
	"github.com/astaxie/beego"
)

type Userlogin struct {
	beego.Controller
}
//直接访问login.html页面请求
func (l *Userlogin) Get() {
	//设置logi.html为模板文件
	l.TplName = "dengru.html"
}

//处理用户登入

//1.解析客户端用户登入提交的数据
func (l*Userlogin)  Post(){
	var user models.User
	err := l.ParseForm(&user)
	if err !=nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("解析失败,请重试!")
		return
	}

	//2.返回数据库查询用户信息
	u,err := user.QueryUser()
	if err!= nil {
		fmt.Println("---------",err.Error())
		l.Ctx.WriteString("抱歉洪湖登入失败,请重试!")

	}


	//3.判断数据库的查询结果
	if err!= nil {
		//sql:no rows in result set
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉,错误!")
		return
	}



	//4.根据查询结果返回客户端相应的处理结果信息或者页面跳转
	l.Data["dianhua"] = u.Dianhua//设置动态数据
	l.TplName = "home.html"



}
