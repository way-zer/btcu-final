package routers

import (
	"btcu-final/server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.CopyrightShowController{})                 //默认跳转到版权展示界面
	beego.Router("/register", &controllers.RegisterController{})              //注册
	beego.Router("/login", &controllers.LoginController{})                    //登录
	beego.Router("/exit", &controllers.ExitController{})                      //退出
	beego.Router("/copyright/add", &controllers.CopyrightAddController{})     //版权登记
	beego.Router("/copyright/:hash", &controllers.CopyrightQueryController{}) //版权查询
	beego.Router("/copyright/home", &controllers.CopyrightShowController{})   //版权展示
	beego.Router("/upload", &controllers.UploadController{})                  //文件上传
}
