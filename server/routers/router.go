package routers

import (
	"btcu-final/server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.CopyrightShowController{})
	//注册
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/copyright/add", &controllers.AddCopyrightController{})
	beego.Router("/copyright/:hash", &controllers.ShowCopyrightController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/copyright/home", &controllers.CopyrightShowController{})
}
