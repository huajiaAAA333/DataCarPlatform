package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)


type SmsCode struct {
	Code string `json:"code"`
}

type SmsResult struct {
	BizId string
	Code string
	Message string
	RequestId string
}

const SMS_TLP_REGISTER  = "SMS_205393604"//注册业务的短信模板
const SMS_TLP_LOGIN  = "SMS_205398654"//用户登入的短信模板
const SMS_TLP_KYC  = ""//实名认证的短信模板

//该函数用于发送一条短信
//电话:dianhua   发送验证码数字:code  模板数字:templateType
func SendSms(Dianhua string,code string,templateType string)(*SmsResult ,error) {
	config := beego.AppConfig
	//获取配置文件中国的accesskey
	accessKey := config.String("sms_access_ket")
	fmt.Println(accessKey)
	//secret
	accessKeySecret := config.String("sms_access_secret")
	client,err := dysmsapi.NewClientWithAccessKey("cn-hangzhou",accessKey,accessKeySecret)
	if err!=nil {
		return nil, err
	}
	//batch:批量.批次
	request := dysmsapi.CreateSendSmsRequest()
    request.DianhuaNubers = dianhua//指定要发送的目标手机号
    request.SignName = "线上餐厅"//签名信息
    request.TemplateCode = templateType//指定短信模板
    //{"code" : "xxxx"}json格式
    SmsCode := SmsCode{
    	Code : code,
	}
	smsbytes,_ := json.Marshal(SmsCode)

    request.TemplateParam = string(smsbytes)//指定要发送的验证码
    response,err := client.SendSms(request)
	if err!= nil {
		return nil,err
	}
    //Biz: business,商业,业务
   SmsResult := &SmsResult{
	   BizId:     response.BizId,
	   Code:      response.Code,
	   Message:   response.Message,
	   RequestId: response.RequestId,
   }
   return SmsResult,nil
}


//生成一个味素为width的随机验证码，并将该验证码返回
func CenRandCode(width int)string  {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0;i <width ;i++  {
		fmt.Fprintf(&sb,"%d",numeric[rand.Intn(r)])
	}
	return sb.String()
}