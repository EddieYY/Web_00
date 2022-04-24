package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"web_00/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/project7", &controllers.Project7Controller{})
}
