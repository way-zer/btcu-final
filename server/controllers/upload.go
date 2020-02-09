package controllers

import (
	"btcu-final/server/models"
	"btcu-final/server/utils"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	BaseController
}

func (this *UploadController) Post() {
	fmt.Println("file upload...")
	fileData, fileHeader, err := this.GetFile("upload")
	if err != nil {
		this.responseErr(err)
		return
	}
	fmt.Println("filename:", fileHeader.Filename, "filesize:", fileHeader.Size)
	fmt.Println(fileData)
	now := time.Now()
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileType := "other"

	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		this.responseErr(err)
		return
	}
	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		this.responseErr(err)
		return
	}
	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		this.responseErr(err)
		return
	}

	// 文件上传之后，对文件进行hash，和signture
	hash := ""
	owner := ""
	signature := ""
	hash = utils.GetHash(filePathStr)
	fmt.Println("hash:", hash)
	document := models.Document{0, fileName, filePathStr, hash, owner, timeStamp, signature}

	//上传文件之前先判断该文档是否已经被上传过，如果已经存在，返回错误
	document1 := models.QueryDocumentWithHash(hash)
	fmt.Println(document1)
	if document1.Id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "文件已经存在", "hash": hash}
		this.ServeJSON()
		return
	}
	_, err = models.InsertDocument(document)
	if err != nil {
		this.responseErr(err)
		return
	}
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功", "hash": hash}
	this.ServeJSON()
}

func (this *UploadController) responseErr(err error) {
	this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	this.ServeJSON()
}
