package models

import (
	"DataCerPlatfomr/database"
	"DataCerPlatfomr/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)
//用户注册信息

type User struct {
	UserId int `form:"user_id"`
	Name string `form:"name"`
	Password string `form:"password"`
	Time string `form:"time"`
	Dianhua string `form:"dianhua"`
	Card     string `form:"card"` //身份证号
	Sex      string `form:"sex"`//性别
}

//该方法用于更新数据库中用户记录的实名认证信息
func (u User) UpdateUser() (int64, error) {
	rs, err := database.Db.Exec("update user set name = ?, card = ?, sex = ? where dianhua = ?", u.Name, u.Card, u.Dianhua,u.Sex)
	if err != nil {
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}




/*
保存用信息的方法,保存用户星系到数据库当中
 */

func (u *User) SaveUser(user *User) {
	//1.密码哈希
	md5Hash := md5.New()
	md5Hash.Write([]byte(user.Password))
	passwordBytes := md5Hash.Sum(nil)
	user.Password = hex.EncodeToString(passwordBytes)


	//2执行数据库操作
	row,err :=database.Db.Exec("insert into username(dianhua,password)"+"values(?,?)",u.Dianhua,u.Password)
	fmt.Println(row,err)



}




//用户登入信息
type UserLogin struct {
	UserId int
	PassWord string
}


func (u *User) AddUser() {
	//脱敏
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	padBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(padBytes)
	//_,err := database.Db.Exce("insert into user(dianhua,passward) values(?,?)",u.Dianhua,u.Password)
	_,err :=database.Db.Exec("")
	//错误早发现早解决
	if err != nil{
		fmt.Println(err)
	}




	database.Db.Exec("insert into user(dianhua,passward) values(?,?)",u.Dianhua,u.Password)

}

func (u *User) QueryUser() (*User,error) {


	u.Password = utils.MD5HashString(u.Password)


	row,err := database.Db.Query("select dianhua from username where dianhua=? and password=? ",u.Dianhua,u.Password)
	if err!=nil {
		return nil,err
	}

	//有问题的地方
	//err = row.Scan(u)
	//if err!=nil {
	//	return nil,err
	//}
	for row.Next() {
		row.Scan(u.Dianhua)
	}
	return u,nil

}



func (u User) QueryUserByDianhua() (*User,error){
	row := database.Db.QueryRow("select id from user where dianhua = ?",u.Dianhua,u.Password,u.Name,u.UserId,u.Time)
	err:= row.Scan(&u.UserId)
	if err != nil {
		return nil,err

	}
	return &u,nil
}









