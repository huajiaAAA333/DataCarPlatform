package models

import "DataCerPlatfomr/database"

type SmsRecord struct {
	BizId string
	Dianhua string
	Code string
	Status string
	Message string
	TimeStamp int64
}

func (s SmsRecord) SaveSmsRecord() {
	database.Db.Exec("insert into sms_record(...")
}
