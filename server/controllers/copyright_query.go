package controllers

import (
	"btcu-final/server/models"
	"fmt"
	"strconv"
)

type ShowCopyrightController struct {
	BaseController
}

func (this *ShowCopyrightController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id", id)
	art := models.QueryCopyrightWithId(id)
	this.Data["name"] = art.Name
	this.Data["author"] = art.Author
	this.Data["press"] = art.Press
	this.Data["hash"] = art.Hash
	this.Data["signature"] = art.Signature
	this.TplName = "copyright_query.html"
}

func (this *ShowCopyrightController) Post() {

	name := this.GetString("name")
	author := this.GetString("author")
	press := this.GetString("press")
	hash := this.GetString("hash")

	fmt.Printf("name:%s,author:%s,press:%s,hash:%s\n", name, author, press, hash)

	var copyright models.Copyright
	//查询版权是否存在，若不存在则返回错误信息
	// 如果填写了hash，就按照hash查找
	if hash != "" {
		copyright = models.QueryCopyrightWithHash(hash)
		fmt.Println(copyright)
		if copyright.Id == 0 {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "按照hash查找的版权不存在"}
			this.ServeJSON()
			return
		}
	} else if (name != "") && (author != "") {
		copyright = models.QueryCopyrightWithName(name, author)
		fmt.Println(copyright)
		if copyright.Id == 0 {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "按作品名和作者名查找的版权不存在"}
			this.ServeJSON()
			return
		}
	}

	response := map[string]interface{}{
		"code": 1, "message": "版权查询成功！",
		"name": copyright.Name, "author": copyright.Author, "press": copyright.Press,
		"timestamp": copyright.Timestamp, "hash": copyright.Hash, "signature": copyright.Signature}

	this.Data["json"] = response
	this.ServeJSON()

}
