package controllers

import (
	"btcu-final/clientSDK"
	"btcu-final/server/models"
	"fmt"
	"github.com/cloudflare/cfssl/log"
)

type CopyrightAddController struct {
	BaseController
}

func (this *CopyrightAddController) Get() {
	this.TplName = "copyright_add.html"
}

func (this *CopyrightAddController) Post() {

	name := this.GetString("name")
	author := this.GetString("author")
	press := this.GetString("press")
	hash := this.GetString("hash")
	privateKey := this.GetString("privateKey")
	copyrightNum, _ := this.GetInt("copyrightNum")
	fmt.Printf("name:%s,hash:%s\n", name, hash)

	// 上传文件之前先查询链码，判断该文档是否已经被上传过，如果已经存在，返回错误
	copyright1, err := clientSDK.GetRightByHash(hash)
	if copyright1 != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "版权已经存在", "hash": hash}
		this.ServeJSON()
		return
	}

	// 根据session获取登录用户的公钥
	username := this.GetSession("loginuser").(string)
	publicKey := models.GetPublicKeyWithUsername(username)
	copyright2 := clientSDK.Copyright{
		Name:      name,
		Author:    author,
		Press:     press,
		Hash:      hash,
		PublicKey: clientSDK.PublicKey(publicKey),
	}

	var response map[string]interface{}

	// 调用链码进行版权登记
	data, err := clientSDK.Register(&copyright2, clientSDK.PrivateKey(privateKey))
	if err != nil {
		response = map[string]interface{}{"code": 0, "message": "调用链码进行版权登记失败！"}
		log.Fatal(err)
	} else {
		response = map[string]interface{}{
			"code": 1, "message": "版权登记成功！",
			"作品名称":   name,
			"作者":     author,
			"出版社":    press,
			"作品hash": hash,
			"作者公钥":   publicKey,
			"作品签名":   data.Signature,
			"作品时间戳":  data.Timestamp,
		}
	}

	// 将版权信息写入数据库
	copyright := models.Copyright{0, name, author, press, hash, publicKey, data.Signature, data.Timestamp, copyrightNum}
	_, err = models.AddCopyright(copyright)

	if err != nil {
		log.Fatal(err)
	}

	this.Data["json"] = response
	this.ServeJSON()
}
