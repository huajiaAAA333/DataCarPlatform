package controllers

import (
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"github.com/astaxie/beego"
	"time"
)
type SendSmsController struct {
	beego.Controller
}

/**
 * 发送短信验证码的功能
 */
func (s *SendSmsController) Post() {
	var smsLogin models.SmsLogin
	err := s.ParseForm(&smsLogin)
	if err != nil {
		s.Ctx.WriteString("发送验证码数据解析失败")
		return
	}
	dianhua := smsLogin.Dianhua
	code := utils.GenRandCode(6)
	result,err := utils.SendSms(dianhua,code,utils.SMS_TLP_REGISTER)
	if err != nil {
		s.Ctx.WriteString("发送验证码失败，请重试")
		return
	}
	if len(result.BizId) == 0 {
		s.Ctx.WriteString(result.Message)
		return
	}
	//发送验证码成功
	smsRecord := models.SmsRecord{
		BizId: result.BizId,
		Dianhua: dianhua,
		Code: code,
		Status: result.Code,
		Message: result.Message,
		TimeStamp: time.Now().Unix(),
	}

}

