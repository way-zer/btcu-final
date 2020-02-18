package controllers

import (
	"btcu-final/clientSDK"
	"fmt"
)

type CopyrightQueryController struct {
	BaseController
}

func (this *CopyrightQueryController) Get() {
	this.TplName = "copyright_query.html"
}

func (this *CopyrightQueryController) Post() {

	name := this.GetString("name")
	author := this.GetString("author")
	press := this.GetString("press")
	hash := this.GetString("hash")

	fmt.Printf("name:%s,author:%s,press:%s,hash:%s\n", name, author, press, hash)

	// 查询版权是否存在，若不存在则返回错误信息
	// 如果填写了hash，就按照hash查找,否则按作品名、作者名和出版社查找
	if hash != "" {
		copyright, err := clientSDK.GetRightByHash(hash)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "按照hash查找的版权不存在"}
			this.ServeJSON()
			return
		}
		if copyright != nil {
			response := map[string]interface{}{
				"code": 1, "message": "版权查询成功！",
				"name": copyright.Name, "author": copyright.Author, "press": copyright.Press,
				"timestamp": copyright.Timestamp, "hash": copyright.Hash, "signature": copyright.Signature}

			this.Data["json"] = response
			this.ServeJSON()
			return
		}
	} else if (name != "") && (author != "") && (press != "") {
		copyright, err := clientSDK.GetRightByInfo(name, author, press)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "按照name,author,press查找的版权不存在"}
			this.ServeJSON()
			return
		}
		if copyright != nil {
			response := map[string]interface{}{
				"code": 1, "message": "版权查询成功！",
				"name": copyright.Name, "author": copyright.Author, "press": copyright.Press,
				"timestamp": copyright.Timestamp, "hash": copyright.Hash, "signature": copyright.Signature}

			this.Data["json"] = response
			this.ServeJSON()
			return
		}
	}
}
