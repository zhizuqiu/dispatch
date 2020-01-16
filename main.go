package main

import (
	_ "dispatch/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.WebConfig.DirectoryIndex = true

	beego.BConfig.WebConfig.StaticDir["/dispatch"] = "data"

	beego.Run()
}
