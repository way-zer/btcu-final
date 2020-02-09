package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type CopyrightHomeBlockParam struct {
	Id         int
	Name       string
	Author     string
	Press      string
	Timestamp  string
	Hash       string
	Signature  string
	Link       string
	UpdateLink string

	//记录是否登录
	IsLogin bool
}

/**
 * 分页的结构体
 */
type CopyrightHomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

//----------首页显示内容---------
func CopyrightMakeHomeBlocks(copyrights []Copyright, isLogin bool) template.HTML {
	htmlHome := ""
	for _, copyright := range copyrights {

		//将数据库model转换为首页模板所需要的model
		homeParam := CopyrightHomeBlockParam{}
		homeParam.Id = copyright.Id
		homeParam.Name = copyright.Name
		homeParam.Author = copyright.Author
		homeParam.Press = copyright.Press
		homeParam.Timestamp = strconv.FormatInt(copyright.Timestamp, 10)
		homeParam.Hash = copyright.Hash
		homeParam.Signature = copyright.Signature
		homeParam.Link = "/copyright/" + strconv.Itoa(copyright.Id)
		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/copyright_home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的变量替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}

//-----------翻页-----------
//page是当前的页数
func CopyrightConfigHomeFooterPageCode(page int) CopyrightHomeFooterPageCode {
	pageCode := CopyrightHomeFooterPageCode{}

	//查询出总的条数
	num := GetCopyrightRowsNum()

	//从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")

	//计算出总页数
	fmt.Println(num)
	allPageNum := (num-1)/pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}
