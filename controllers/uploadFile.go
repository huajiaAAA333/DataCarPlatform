package controllers

import (
	"DataCerPlatfomr/models"
	"DataCerPlatfomr/utils"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strings"
	"time"
)

//改控制器解构体用于处理文件上传功能

type UploadFile struct {
	beego.Controller
}

//该post方法用于处理用户在客户端提交认证文件

func (u *UploadFile) Post() {

	dianhua := u.Ctx.Request.PostFormValue("dianhua") //获取用户的电话信息

	//1.解析用户上传的数据
	//用户上传的自定义的标题
	title := u.Ctx.Request.PostFormValue("upload_title") //用户输入的标题
	fmt.Println("电子数据标签:", title)

	//用户上传的文件
	file, header, err := u.GetFile("dong")
	if err != nil { //解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉,,文件解析失败")
		return

	}
	defer file.Close() //延迟执行,

	//使用io包提供保存文件的方法
	//io.Copy: dst :数据复制的目的地 src:数据源
	//	//返回值,复制的长度
	saveFilePath := "static/upload/" + header.Filename
	//saveFile,err := os.OpenFile(saveFilePath,os.O_CREATE|os.O_RDWR,777)
	//if err!= nil {
	//	u.Ctx.WriteString("文件认证保存失败")
	//	return
	//}
	////saveFile.Write()
	//_,err = io.Copy(saveFile,file)
	//if err!= nil  {
	//	u.Ctx.WriteString("抱歉,,电子认证数据是啊比,请重新尝试!")
	//	return
	//}
	_, err = utils.SaveFile(saveFilePath, file)
	if err != nil {
		u.Ctx.WriteString("文件数据认证失败,请重试!")
		return
	}

	//2.计算文件的sha256值
	fileHash, err := utils.SHA256HashReader(file)
	if err != nil {
		return
	}

	fmt.Println(fileHash)

	//先查询ID
	user1, err := models.User{Dianhua: dianhua}.UserByDianhua()
	if err != nil {
		u.Ctx.WriteString("wenjiancuowu")

		return
	}
	//把上传的文件作为记录保存到数据库中1
	//计算mad5值
	saveFile,err := os.Open(saveFilePath)
	md5String, err := utils.MD5HashReader(saveFile)
	if err != nil {
		u.Ctx.WriteString("抱歉,电子数据认证失败")
		return
	}
	record := models.UploadRecord{
		UserID:    user1.UserId,
		FileName:  header.Filename,
		FileSize:  header.Size,
		FileCert:  md5String,
		FileTitle: title,
		Time:      time.Now().Unix(),
	}

	//保存认证数据到数据库中
	_, err = record.SaveRecord()
	if err != nil {
		u.Ctx.WriteString("点整认证数据保存失败")
		return
	}
	//上传文件保存到数据库中成功
	records, err := models.QueryRecordsByUserId(user1.UserId)
	if err != nil {
		u.Ctx.WriteString("获取电子数据列表失败")
		return

	}
	u.Data["Records"] = records
	u.TplName = "list_record.html"
}

//
//
//	//文件上传列表(专跳 )
//
//
//	u.Ctx.WriteString("恭喜已经接收到上传文件!")
//}

/*
	post方法用于吃力用户在客户端提交的认证文件
*/
func (u *UploadFile) Post1() {
	//1、解析用户上传的数据及文件内容
	//用户上传的自定义的标题
	title := u.Ctx.Request.PostFormValue("upload_title") //用户输入的标题

	//用户上传的文件
	file, header, err := u.GetFile("dong")
	if err != nil { //解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}
	defer file.Close()

	fmt.Println("自定义的标题:", title)
	//获得了上传文件
	fmt.Println("上传的文件昵称:", header.Filename)
	///截取文件格式,判断后缀格式是否支持上传
	fileNameSlice := strings.Split(header.Filename, ".")
	fileType := fileNameSlice[1]
	fmt.Println(":", strings.TrimSpace(fileType))
	isJpg := strings.HasSuffix(header.Filename, ".jpg")
	isPng := strings.HasSuffix(header.Filename, ".png")
	if !isJpg && !isPng {
		//文件类型不支持
		u.Ctx.WriteString("抱歉,文件类型不符合,请重新上传符合的文件格式")
		return
	}
	//文件大小 判断文件大小 200kb
	config := beego.AppConfig
	fileSize, err := config.Int64("file_size")
	if header.Size/1024 > 200 {
		u.Ctx.WriteString("抱歉,文件大小不符合,请重新上传符合的文件格式")
		return
	}
	fmt.Println(fileSize)
	fmt.Println("上传文件大小:", header.Size) //字节  大小

	//fromFile:文件
	//toFile:要保存的文件路径
	//
	//perm:权限
	//权限的组成: a+b+c
	//a:文件所有者对文件的操作权限,读4,写2,执行1
	//b:文件所有者所在组的用户操作权限,读4,写2,执行1
	//c:其他用户的操作权限,读4,写2,执行1
	//eg:m文件,权限是:651.
	//判断:文件所有者对该m文件有写权限(对),,,解析:文件所有者对该文件对应到数字第一位,6=4+2,对应有读权限和写权限

	saveDir := "static/upload"
	//1.打开文件
	//尝试打开文件
	//_,err = os.Open(saveDir)
	F, err := os.Open(saveDir)
	if err != nil {
		err = os.Mkdir(saveDir, 777)
		if err != nil { //打开失败
			//2.自己手动创建文件夹
			fmt.Println(err.Error())
			u.Ctx.WriteString("抱歉打开文件夹失败")
			return
		}
	}

	fmt.Println("打开的文件夹:", F.Name())
	err = os.Mkdir("static/upload1", 777)

	if err != nil {
		u.Ctx.WriteString("抱歉,文件夹遇到错误,请重试!")
		return
	}

	//文件名: 文件路径 + 文件名+ "." + 文件拓展名
	saveName := "static/upload/" + header.Filename
	fmt.Println("要保存的文件名:", saveName)
	//formfile:文件
	//tofile:要保存的文件路径
	err = u.SaveToFile("dong", saveName)
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉,文件夹认证失败,请重试!")
		return

	}

	//u.saveToFile("file",saveName)

	fmt.Println("上传的文件:", file)
	u.Ctx.WriteString("以获取上传文件")

}
