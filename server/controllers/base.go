package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Islogin   bool
	Loginuser interface{}
}

func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser---->", loginuser)
	if loginuser != nil {
		this.Islogin = true
		this.Loginuser = loginuser
	} else {
		this.Islogin = false
	}
	this.Data["IsLogin"] = this.Islogin
}
