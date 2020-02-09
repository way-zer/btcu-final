package controllers

import (
	"btcu-final/clientSDK"
	"btcu-final/server/models"
	"fmt"
	"github.com/cloudflare/cfssl/log"
)

type AddCopyrightController struct {
	BaseController
}

func (this *AddCopyrightController) Get() {
	this.TplName = "copyright_add.html"
}

func (this *AddCopyrightController) Post() {

	name := this.GetString("name")
	author := this.GetString("author")
	press := this.GetString("press")
	hash := this.GetString("hash")
	privateKey := this.GetString("privateKey")
	copyrightNum, _ := this.GetInt("copyrightNum")
	fmt.Printf("name:%s,hash:%s\n", name, hash)

	//上传文件之前先判断该文档是否已经被上传过，如果已经存在，返回错误
	copyright1 := models.QueryCopyrightWithHash(hash)
	fmt.Println(copyright1)
	if copyright1.Id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "版权已经存在", "hash": hash}
		this.ServeJSON()
		return
	}

	username := this.GetSession("loginuser").(string)
	publicKey := models.GetPublicKeyWithUsername(username)
	copyright2 := clientSDK.Copyright{
		Name:      name,
		Author:    author,
		Press:     press,
		Hash:      hash,
		PublicKey: clientSDK.PublicKey(publicKey),
	}
	data, err := clientSDK.Register(&copyright2, clientSDK.PrivateKey(privateKey))
	if err != nil {
		log.Fatal(err)
	}

	copyright := models.Copyright{0, name, author, press, hash, publicKey, data.Signature, data.Timestamp, copyrightNum}
	_, err = models.AddCopyright(copyright)

	var response map[string]interface{}
	if err == nil {
		response = map[string]interface{}{
			"code": 1, "message": "版权登记成功！",
			"作品名称":   copyright.Name,
			"作者":     copyright.Author,
			"出版社":    copyright.Press,
			"作品hash": copyright.Hash,
			"作者公钥":   copyright.PublicKey,
			"作品签名":   copyright.Signature,
			"作品时间戳":  copyright.Timestamp,
		}
	} else {
		response = map[string]interface{}{"code": 0, "message": "版权登记失败！"}
	}

	this.Data["json"] = response
	this.ServeJSON()
}
