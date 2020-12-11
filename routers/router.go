package routers

import (
	"dispatch/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.MainController{}, "GET:Upload")
	beego.Router("/upload2", &controllers.MainController{}, "GET:Upload2")

	beego.Router("/file/List", &controllers.MainController{}, "GET:List")
	beego.Router("/file/Delete", &controllers.MainController{}, "DELETE:Delete")
	beego.Router("/file/newDir", &controllers.MainController{}, "PUT:NewDir")
	beego.Router("/file/UploadFile", &controllers.MainController{}, "POST:UploadFile")

	beego.Router("/getDanmuServer", &controllers.MainController{}, "GET:GetDanmuServer")
}
