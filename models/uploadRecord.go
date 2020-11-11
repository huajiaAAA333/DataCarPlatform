package models

import (
	"DataCerPlatfomr/database"
	"time"
)

//上传文件的记录
type UploadRecord struct {
	ID int
	UserID int
	FileName string
	FileSize int64
	FileCert string
	FileTitle string
	Time int64
	TimeFormat string
}
/*
把一条认证数据保存到表中
 */

func (u UploadRecord) SaveRecord()(int64, error ) {
	rs,err := database.Db.Exec("insert into upload_record(user_id,file_name,file_size,file_cert,file_titel,time) "+
		"values(?,?,?,?,?,?)")
	if err != nil {
		return-1,err

	}
	id,err := rs.RowsAffected()
	if err != nil {
		return-1 ,err

	}
	return id,nil
}


//根据用户ID查询符合条件的认证数据信息记录
func QueryRecordsByUserId(userId int) ([]UploadRecord,error)  {
	rs, err := database.Db.Query("select id,user_id,file_name,file_size,file_cert,file_title,time )where user_id = ?")
	if err != nil {
		return nil,err

	}
	//从rs中读取查询到的数据,返回
	records := make([]UploadRecord,0)//容器

	for rs.Next()  {
		var record UploadRecord
		rs.Scan(&record.ID,&record.UserID,&record.FileName,&record.FileSize,&record.FileCert,&record.FileTitle,&record.Time)
		if err != nil {
			return nil,err
		}
		//将time整形 变成字符串几年几月几分几秒
		//utils.TimeFormat(record.Time,"2006/01/02 15:04:05")
		t := time.Unix(record.Time,0)
		tStr := t.Format("2006/01/02 15：04：05")
		record.TimeFormat = tStr

		records = append(records,record)
	}
	return records,nil



}







