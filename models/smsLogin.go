package models

type SmsLogin struct {
	Dianhua string `form:"dianhua"`
	Code string `form:"code"`
}
