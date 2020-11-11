package routers

import (
	"DataCerPlatfomr/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router:路由
	//用户注册的接口请求 http://127.0.0.1:8080/register
	beego.Router("/user_register",&controllers.RegisterController{})
	////http://127.0.0.1:8080
	////beego.Router("/", &controllers.MainController{})
	//登入页面接口
	//beego.Router("/login",&controllers.Userlogin{})

	//直接登入的页面请求接口
	beego.Router("/dengru",&controllers.Userlogin{})

	////登入成功跳转主页面
	beego.Router("/userlogin",&controllers.Userlogin{})

	//文件上传的接口
	beego.Router("/home",&controllers.UploadFile{})

	//查看认证数据证书
	beego.Router("/cert_detail.html",&controllers.CertDetailController{})

	//用户实名认证
	beego.Router("kyc.html",&controllers.UserKycController{})


	//验证码登入
	beego.Router("login_sms.html",&controllers.LoginSmsController{})











}