package main

import (
	_ "btcu-final/server/routers"
	"btcu-final/server/utils"
	"github.com/astaxie/beego"
)

func main() {

	utils.InitMysql()
	beego.Run()

}
