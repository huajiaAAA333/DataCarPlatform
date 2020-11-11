package controllers

import ("DataCerPlatfomr/models"
"github.com/astaxie/beego")

type UserKycController struct {
	beego.Controller
}

//浏览器的get请求,用于转跳到实名认证页面
func (u *UserKycController)Get()  {
	u.TplName = "kyc.html"
}


//for表单的post请求,用于处理实名认证业务
func (u *UserKycController)Post()  {
	//1.解析前段的数据
	var user models.User
	err := u.ParseForm(&user)
	if err!=nil {
		u.Ctx.WriteString("抱歉,数据解析错误")
		return
	}
	//2.把用户的实名认证的更新到数据库的用户表中
	_,err = user.UpdateUser()

	//3.判断实名认证操作结果
	if err!=nil {
		u.Ctx.WriteString("抱歉,数据解析错误")
		return
	}
	//4.转跳到上传页面
	u.TplName = "home.html"

}
