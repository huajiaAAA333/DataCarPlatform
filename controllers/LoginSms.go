package controllers


type LoginSmsController struct {
	beego.Controller
}

//浏览器发起短信验证码的登入请求

func (l*LoginSmsController)Get(){
	l.TplName = "login_sms.html"
}